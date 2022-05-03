// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package registry

import (
	runtimev1 "github.com/atomix/runtime-api/api/atomix/runtime/v1"
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
	var writer io.WriteCloser
	defer writer.Close()
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		if writer == nil {
			writer, err = s.registry.Create(request.Driver.Name, request.Driver.Version, request.Runtime.Version)
			if err != nil {
				return err
			}
		}

		_, err = writer.Write(request.Data)
		if err != nil {
			return err
		}
	}
}

func (s *Server) PullDriver(request *runtimev1.PullDriverRequest, stream runtimev1.Registry_PullDriverServer) error {
	reader, err := s.registry.Open(request.Driver.Name, request.Driver.Version, request.Runtime.Version)
	if err != nil {
		return err
	}
	defer reader.Close()

	bytes := make([]byte, bufferSize)
	for {
		_, err := reader.Read(bytes)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		response := &runtimev1.PullDriverResponse{
			Data: bytes,
		}
		if err := stream.Send(response); err != nil {
			return err
		}
	}
}

var _ runtimev1.RegistryServer = (*Server)(nil)
