// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"context"
	"fmt"
	runtimev1 "github.com/atomix/runtime-api/api/atomix/runtime/v1"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/grpc/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"os"
	"path/filepath"
	"plugin"
	"sync"
)

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

func (r *Repository) Load(ctx context.Context, name, version, apiVersion string) (Driver, error) {
	plugin, err := r.get(name, version, apiVersion).Load(ctx)
	if err != nil {
		return nil, err
	}
	symbol, err := plugin.Lookup("Driver")
	if err != nil {
		return nil, err
	}
	return symbol.(Driver), nil
}

func (r *Repository) get(name, version, apiVersion string) *Plugin {
	r.mu.Lock()
	defer r.mu.Unlock()

	path := filepath.Join(r.Path, fmt.Sprintf("%s-%s.%s.so", name, version, apiVersion))
	plugin, ok := r.plugins[path]
	if !ok {
		plugin = &Plugin{
			registry:   r,
			Name:       name,
			Version:    version,
			APIVersion: apiVersion,
			Path:       path,
		}
		r.plugins[path] = plugin
	}
	return plugin
}

func (r *Repository) Close() error {
	return r.conn.Close()
}

type Plugin struct {
	registry   *Repository
	Name       string
	Version    string
	APIVersion string
	Path       string
	plugin     *plugin.Plugin
	mu         sync.RWMutex
}

func (p *Plugin) Load(ctx context.Context) (*plugin.Plugin, error) {
	p.mu.RLock()
	loaded := p.plugin
	p.mu.RUnlock()
	if loaded != nil {
		return loaded, nil
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.plugin != nil {
		return p.plugin, nil
	}

	client := runtimev1.NewRegistryClient(p.registry.conn)
	request := &runtimev1.PullDriverRequest{
		Driver: runtimev1.DriverInfo{
			Name:    p.Name,
			Version: p.Version,
		},
		Runtime: runtimev1.RuntimeInfo{
			Version: p.APIVersion,
		},
	}
	stream, err := client.PullDriver(ctx, request)
	if err != nil {
		return nil, errors.From(err)
	}

	writer, err := os.Create(p.Path)
	if err != nil {
		return nil, err
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			err = writer.Close()
			if err != nil {
				return nil, err
			}
			loaded, err := plugin.Open(p.Path)
			if err != nil {
				return nil, err
			}
			p.plugin = loaded
			return loaded, nil
		}
		if err != nil {
			_ = writer.Close()
			_ = os.Remove(p.Path)
			return nil, errors.From(err)
		}

		_, err = writer.Write(response.Data)
		if err != nil {
			_ = writer.Close()
			_ = os.Remove(p.Path)
			return nil, err
		}
	}
}
