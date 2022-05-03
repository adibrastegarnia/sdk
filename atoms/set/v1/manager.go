// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	setv1 "github.com/atomix/runtime-api/api/atomix/set/v1"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/runtime"
)

func newSetV1ManagerServer(proxies *runtime.ProxyService[SetProxy]) setv1.SetManagerServer {
	return &setV1ManagerServer{
		proxies: proxies,
	}
}

type setV1ManagerServer struct {
	proxies *runtime.ProxyService[SetProxy]
}

func (s *setV1ManagerServer) Create(ctx context.Context, request *setv1.CreateRequest) (*setv1.CreateResponse, error) {
	namespace, err := s.proxies.GetNamespace(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	err = namespace.CreateProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	return &setv1.CreateResponse{}, nil
}

func (s *setV1ManagerServer) Close(ctx context.Context, request *setv1.CloseRequest) (*setv1.CloseResponse, error) {
	namespace, err := s.proxies.GetNamespace(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	err = namespace.CloseProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	return &setv1.CloseResponse{}, nil
}

var _ setv1.SetManagerServer = (*setV1ManagerServer)(nil)
