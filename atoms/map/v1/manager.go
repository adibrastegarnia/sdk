// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	mapv1 "github.com/atomix/runtime-api/api/atomix/map/v1"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/runtime"
)

func newMapV1ManagerServer(proxies *runtime.ProxyService[MapProxy]) mapv1.MapManagerServer {
	return &mapV1ManagerServer{
		proxies: proxies,
	}
}

type mapV1ManagerServer struct {
	proxies *runtime.ProxyService[MapProxy]
}

func (s *mapV1ManagerServer) Create(ctx context.Context, request *mapv1.CreateRequest) (*mapv1.CreateResponse, error) {
	namespace, err := s.proxies.GetCluster(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	err = namespace.CreateProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	return &mapv1.CreateResponse{}, nil
}

func (s *mapV1ManagerServer) Close(ctx context.Context, request *mapv1.CloseRequest) (*mapv1.CloseResponse, error) {
	namespace, err := s.proxies.GetCluster(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	err = namespace.CloseProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	return &mapv1.CloseResponse{}, nil
}

var _ mapv1.MapManagerServer = (*mapV1ManagerServer)(nil)
