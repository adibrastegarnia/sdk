// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package runtime

func NewService(runtime *Runtime, opts ...Option) *Service {
	var options Options
	options.apply(opts...)
	return &Service{
		Options: options,
	}
}

type Service struct {
	Options
	runtime *Runtime
	server  *Server
}

func (s *Service) Start() error {
	s.runtime = New(WithOptions(s.Options))
	if err := s.runtime.Run(); err != nil {
		return err
	}
	s.server = newServer(s.Server)
	for _, f := range s.Atoms {
		f(s.server.Server, s.runtime)
	}
	if err := s.server.start(); err != nil {
		return err
	}
	return nil
}

func (s *Service) Stop() error {
	return s.runtime.Shutdown()
}
