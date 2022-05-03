// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	listv1 "github.com/atomix/runtime-api/api/atomix/list/v1"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/runtime"
)

func newListV1ManagerServer(proxies *runtime.ProxyService[ListProxy]) listv1.ListManagerServer {
	return &listV1ManagerServer{
		proxies: proxies,
	}
}

type listV1ManagerServer struct {
	proxies *runtime.ProxyService[ListProxy]
}

func (s *listV1ManagerServer) Create(ctx context.Context, request *listv1.CreateRequest) (*listv1.CreateResponse, error) {
	namespace, err := s.proxies.GetCluster(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = namespace.CreateProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &listv1.CreateResponse{}, nil
}

func (s *listV1ManagerServer) Close(ctx context.Context, request *listv1.CloseRequest) (*listv1.CloseResponse, error) {
	namespace, err := s.proxies.GetCluster(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = namespace.CloseProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &listv1.CloseResponse{}, nil
}

var _ listv1.ListManagerServer = (*listV1ManagerServer)(nil)
