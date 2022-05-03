// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"context"
	"fmt"
	runtimev1 "github.com/atomix/runtime-api/api/atomix/runtime/v1"
	"google.golang.org/grpc"
	"net"
)

type GetClusterOptions struct{}

type ListClustersOptions struct {
	Cluster string
	Watch   bool
}

type Controller interface {
	GetCluster(ctx context.Context, name string, options GetClusterOptions) (runtimev1.Cluster, error)
	ListClusters(ctx context.Context, ch chan<- runtimev1.Cluster, options ListClustersOptions) error
}

func NewService(controller Controller, opts ...Option) *Service {
	var options Options
	options.apply(opts...)
	server := grpc.NewServer()
	runtimev1.RegisterControllerServer(server, newServer(controller))
	return &Service{
		Options: options,
		server:  server,
	}
}

type Service struct {
	Options
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
