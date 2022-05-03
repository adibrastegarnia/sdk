// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	counterv1 "github.com/atomix/runtime-api/api/atomix/counter/v1"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/runtime/proxy"
)

func newCounterV1ManagerServer(proxies *proxy.Service[CounterProxy]) counterv1.CounterManagerServer {
	return &counterV1ManagerServer{
		proxies: proxies,
	}
}

type counterV1ManagerServer struct {
	proxies *proxy.Service[CounterProxy]
}

func (s *counterV1ManagerServer) Create(ctx context.Context, request *counterv1.CreateRequest) (*counterv1.CreateResponse, error) {
	namespace, err := s.proxies.GetCluster(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = namespace.CreateProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &counterv1.CreateResponse{}, nil
}

func (s *counterV1ManagerServer) Close(ctx context.Context, request *counterv1.CloseRequest) (*counterv1.CloseResponse, error) {
	namespace, err := s.proxies.GetCluster(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = namespace.CloseProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &counterv1.CloseResponse{}, nil
}

var _ counterv1.CounterManagerServer = (*counterV1ManagerServer)(nil)
