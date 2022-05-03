// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	counterv1 "github.com/atomix/runtime-api/api/atomix/counter/v1"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/runtime"
)

func newCounterV1ManagerServer(proxies *runtime.ProxyService[CounterProxy]) counterv1.CounterManagerServer {
	return &counterV1ManagerServer{
		proxies: proxies,
	}
}

type counterV1ManagerServer struct {
	proxies *runtime.ProxyService[CounterProxy]
}

func (s *counterV1ManagerServer) Create(ctx context.Context, request *counterv1.CreateRequest) (*counterv1.CreateResponse, error) {
	namespace, err := s.proxies.GetNamespace(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	err = namespace.CreateProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	return &counterv1.CreateResponse{}, nil
}

func (s *counterV1ManagerServer) Close(ctx context.Context, request *counterv1.CloseRequest) (*counterv1.CloseResponse, error) {
	namespace, err := s.proxies.GetNamespace(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	err = namespace.CloseProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	return &counterv1.CloseResponse{}, nil
}

var _ counterv1.CounterManagerServer = (*counterV1ManagerServer)(nil)
