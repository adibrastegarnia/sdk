// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func NewService(runtime *Runtime, opts ...Option) *Service {
	var options Options
	options.apply(opts...)
	return &Service{
		Options: options,
		runtime: runtime,
		server:  grpc.NewServer(),
	}
}

type Service struct {
	Options
	runtime *Runtime
	server  *grpc.Server
}

func (s *Service) Start() error {
	if err := s.runtime.Run(); err != nil {
		return err
	}
	for _, f := range s.Atoms {
		f(s.server, s.runtime)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Server.Host, s.Server.Port))
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
	return s.runtime.Shutdown()
}
