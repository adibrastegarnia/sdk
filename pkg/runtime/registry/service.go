// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package registry

import (
	"fmt"
	runtimev1 "github.com/atomix/api/pkg/atomix/runtime/v1"
	"google.golang.org/grpc"
	"net"
)

func NewService(opts ...Option) *Service {
	var options Options
	options.apply(opts...)
	server := grpc.NewServer()
	registry := newRegistry(opts...)
	runtimev1.RegisterRegistryServer(server, newServer(registry))
	return &Service{
		ServerOptions: options.Server,
		server:        server,
	}
}

type Service struct {
	ServerOptions
	server *grpc.Server
}

func (s *Service) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		return err
	}
	go func() {
		if err := s.server.Serve(lis); err != nil {
			panic(err)
		}
	}()
	return nil
}

func (s *Service) Stop() error {
	s.server.Stop()
	return nil
}
