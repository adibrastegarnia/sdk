// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func newServer(options ServerOptions) *Server {
	return &Server{
		ServerOptions: options,
		Server:        grpc.NewServer(),
	}
}

type Server struct {
	ServerOptions
	*grpc.Server
}

func (s *Server) start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		return err
	}

	go func() {
		if err := s.Server.Serve(lis); err != nil {
			panic(err)
		}
	}()
	return nil
}

func (s *Server) stop() error {
	s.Server.Stop()
	return nil
}
