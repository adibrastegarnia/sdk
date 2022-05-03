// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package registry

import (
	"crypto/sha256"
	"encoding/hex"
	runtimev1 "github.com/atomix/runtime-api/api/atomix/runtime/v1"
	"github.com/atomix/runtime-api/pkg/errors"
	"io"
)

const bufferSize = 1024 * 1024

func newServer(registry *Registry) *Server {
	return &Server{
		registry: registry,
	}
}

type Server struct {
	registry *Registry
}

func (s *Server) PushDriver(stream runtimev1.Registry_PushDriverServer) error {
	request, err := stream.Recv()
	if err != nil {
		return err
	}

	var writer io.WriteCloser
	switch r := request.Request.(type) {
	case *runtimev1.PushDriverRequest_Header:
		writer, err = s.registry.Create(r.Header.Driver.Name, r.Header.Driver.Version, r.Header.Runtime.Version)
		if err != nil {
			return errors.ToProto(errors.NewInternal(err.Error()))
		}
	case *runtimev1.PushDriverRequest_Chunk:
		return errors.ToProto(errors.NewForbidden("received Chunk request; expected Header"))
	case *runtimev1.PushDriverRequest_Trailer:
		return errors.ToProto(errors.NewForbidden("received Trailer request; expected Header"))
	}

	sha := sha256.New()
	defer writer.Close()
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		switch r := request.Request.(type) {
		case *runtimev1.PushDriverRequest_Header:
			return errors.ToProto(errors.NewForbidden("received Chunk request; expected Chunk or Trailer"))
		case *runtimev1.PushDriverRequest_Chunk:
			_, err = writer.Write(r.Chunk.Data)
			if err != nil {
				return errors.ToProto(errors.NewInternal(err.Error()))
			}

			_, err = sha.Write(r.Chunk.Data)
			if err != nil {
				return errors.ToProto(errors.NewInternal(err.Error()))
			}
		case *runtimev1.PushDriverRequest_Trailer:
			checksum := hex.EncodeToString(sha.Sum(nil))
			if r.Trailer.Checksum != checksum {
				return errors.ToProto(errors.NewFault(""))
			}
			return nil
		}
	}
}

func (s *Server) PullDriver(request *runtimev1.PullDriverRequest, stream runtimev1.Registry_PullDriverServer) error {
	reader, err := s.registry.Open(request.Header.Driver.Name, request.Header.Driver.Version, request.Header.Runtime.Version)
	if err != nil {
		return errors.ToProto(errors.NewInternal(err.Error()))
	}
	defer reader.Close()

	sha := sha256.New()
	buf := make([]byte, bufferSize)
	for {
		i, err := reader.Read(buf)
		if err == io.EOF {
			checksum := hex.EncodeToString(sha.Sum(nil))
			response := &runtimev1.PullDriverResponse{
				Response: &runtimev1.PullDriverResponse_Trailer{
					Trailer: &runtimev1.PluginTrailer{
						Checksum: checksum,
					},
				},
			}
			if err := stream.Send(response); err != nil {
				return err
			}
			return nil
		}
		if err != nil {
			return err
		}

		response := &runtimev1.PullDriverResponse{
			Response: &runtimev1.PullDriverResponse_Chunk{
				Chunk: &runtimev1.PluginChunk{
					Data: buf[:i+1],
				},
			},
		}
		if err := stream.Send(response); err != nil {
			return err
		}

		_, err = sha.Write(buf[:i+1])
		if err != nil {
			return errors.ToProto(errors.NewInternal(err.Error()))
		}
	}
}

var _ runtimev1.RegistryServer = (*Server)(nil)
