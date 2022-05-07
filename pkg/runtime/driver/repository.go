// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	runtimev1 "github.com/atomix/api/pkg/atomix/runtime/v1"
	"github.com/atomix/sdk/pkg/errors"
	"github.com/atomix/sdk/pkg/grpc/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"os"
	"path/filepath"
	"plugin"
	"sync"
)

const bufferSize = 1024 * 1024

func NewRepository(opts ...RepositoryOption) (*Repository, error) {
	var options RepositoryOptions
	options.apply(opts...)
	repo := &Repository{
		RepositoryOptions: options,
		plugins:           make(map[string]*Plugin),
	}
	if err := repo.connect(); err != nil {
		return nil, err
	}
	return repo, nil
}

type Repository struct {
	RepositoryOptions
	conn    *grpc.ClientConn
	plugins map[string]*Plugin
	mu      sync.Mutex
}

func (r *Repository) connect() error {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", r.Registry.Host, r.Registry.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(retry.RetryingUnaryClientInterceptor()),
		grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor()))
	if err != nil {
		return err
	}
	r.conn = conn
	return nil
}

func (r *Repository) Load(ctx context.Context, info PluginInfo) (Driver, error) {
	plugin := r.get(info)
	driver, err := plugin.Load()
	if err == nil {
		return driver, nil
	}

	if err := plugin.Pull(ctx); err != nil {
		return nil, err
	}
	return plugin.Load()
}

func (r *Repository) Pull(ctx context.Context, info PluginInfo) (*Plugin, error) {
	plugin := r.get(info)
	if err := plugin.Pull(ctx); err != nil {
		return nil, err
	}
	return plugin, nil
}

func (r *Repository) Push(ctx context.Context, info PluginInfo) error {
	return r.get(info).Push(ctx)
}

func (r *Repository) get(info PluginInfo) *Plugin {
	r.mu.Lock()
	defer r.mu.Unlock()

	if info.Path == "" {
		info.Path = filepath.Join(r.Path, fmt.Sprintf("%s-%s.%s.so", info.Name, info.Version, info.APIVersion))
	}

	plugin, ok := r.plugins[info.Path]
	if !ok {
		plugin = &Plugin{
			PluginInfo: info,
			repository: r,
		}
		r.plugins[info.Path] = plugin
	}
	return plugin
}

func (r *Repository) Close() error {
	return r.conn.Close()
}

type PluginInfo struct {
	Name       string
	Version    string
	APIVersion string
	Path       string
}

type Plugin struct {
	PluginInfo
	repository *Repository
	driver     Driver
	mu         sync.RWMutex
}

func (p *Plugin) Load() (Driver, error) {
	p.mu.RLock()
	driver := p.driver
	p.mu.RUnlock()
	if driver != nil {
		return driver, nil
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	if p.driver != nil {
		return p.driver, nil
	}

	plugin, err := plugin.Open(p.Path)
	if err != nil {
		return nil, err
	}
	symbol, err := plugin.Lookup("Driver")
	if err != nil {
		return nil, err
	}
	p.driver = symbol.(Driver)
	return driver, nil
}

func (p *Plugin) Push(ctx context.Context) error {
	client := runtimev1.NewRegistryClient(p.repository.conn)

	reader, err := os.Open(p.Path)
	if err != nil {
		return err
	}
	defer reader.Close()

	stream, err := client.PushDriver(ctx)
	if err != nil {
		return errors.FromProto(err)
	}

	request := &runtimev1.PushDriverRequest{
		Request: &runtimev1.PushDriverRequest_Header{
			Header: &runtimev1.PluginHeader{
				Driver: runtimev1.DriverMeta{
					Name:    p.Name,
					Version: p.Version,
				},
				Runtime: runtimev1.RuntimeMeta{
					Version: p.APIVersion,
				},
			},
		},
	}
	if err := stream.Send(request); err != nil {
		return errors.FromProto(err)
	}

	sha := sha256.New()
	buf := make([]byte, bufferSize)
	for {
		i, err := reader.Read(buf)
		if err == io.EOF {
			checksum := hex.EncodeToString(sha.Sum(nil))
			request := &runtimev1.PushDriverRequest{
				Request: &runtimev1.PushDriverRequest_Trailer{
					Trailer: &runtimev1.PluginTrailer{
						Checksum: checksum,
					},
				},
			}
			if err := stream.Send(request); err != nil {
				return errors.FromProto(err)
			}
			return nil
		}
		if err != nil {
			return errors.NewInternal(err.Error())
		}

		request := &runtimev1.PushDriverRequest{
			Request: &runtimev1.PushDriverRequest_Chunk{
				Chunk: &runtimev1.PluginChunk{
					Data: buf[:i+1],
				},
			},
		}
		if err := stream.Send(request); err != nil {
			return errors.FromProto(err)
		}

		_, err = sha.Write(buf[:i+1])
		if err != nil {
			return errors.NewInternal(err.Error())
		}
	}
}

func (p *Plugin) Pull(ctx context.Context) error {
	client := runtimev1.NewRegistryClient(p.repository.conn)
	request := &runtimev1.PullDriverRequest{
		Header: &runtimev1.PluginHeader{
			Driver: runtimev1.DriverMeta{
				Name:    p.Name,
				Version: p.Version,
			},
			Runtime: runtimev1.RuntimeMeta{
				Version: p.APIVersion,
			},
		},
	}
	stream, err := client.PullDriver(ctx, request)
	if err != nil {
		return errors.FromProto(err)
	}

	writer, err := os.Create(p.Path)
	if err != nil {
		return err
	}

	sha := sha256.New()
	for {
		response, err := stream.Recv()
		if err != nil {
			_ = writer.Close()
			_ = os.Remove(p.Path)
			return errors.FromProto(err)
		}

		switch r := response.Response.(type) {
		case *runtimev1.PullDriverResponse_Chunk:
			_, err = writer.Write(r.Chunk.Data)
			if err != nil {
				_ = writer.Close()
				_ = os.Remove(p.Path)
				return errors.NewInternal(err.Error())
			}

			_, err = sha.Write(r.Chunk.Data)
			if err != nil {
				_ = writer.Close()
				_ = os.Remove(p.Path)
				return errors.NewInternal(err.Error())
			}
		case *runtimev1.PullDriverResponse_Trailer:
			err = writer.Close()
			if err != nil {
				return errors.NewInternal(err.Error())
			}
			return p.validate(r.Trailer.Checksum)
		}
	}
}

func (p *Plugin) validate(checksum string) error {
	filesum, err := p.checksum()
	if err != nil {
		return err
	}
	if filesum != checksum {
		return errors.NewFault("checksum for plugin %s did not match", p.Path)
	}
	return nil
}

func (p *Plugin) checksum() (string, error) {
	reader, err := os.Open(p.Path)
	if err != nil {
		return "", err
	}
	defer reader.Close()

	sha := sha256.New()
	buf := make([]byte, bufferSize)
	for {
		_, err := reader.Read(buf)
		if err == io.EOF {
			return hex.EncodeToString(sha.Sum(nil)), nil
		}
		if err != nil {
			return "", err
		}

		_, err = sha.Write(buf)
		if err != nil {
			return "", err
		}
	}
}
