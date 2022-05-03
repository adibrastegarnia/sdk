// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	lockv1 "github.com/atomix/runtime-api/api/atomix/lock/v1"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/runtime"
)

func newLockV1ManagerServer(proxies *runtime.ProxyService[LockProxy]) lockv1.LockManagerServer {
	return &lockV1ManagerServer{
		proxies: proxies,
	}
}

type lockV1ManagerServer struct {
	proxies *runtime.ProxyService[LockProxy]
}

func (s *lockV1ManagerServer) Create(ctx context.Context, request *lockv1.CreateRequest) (*lockv1.CreateResponse, error) {
	namespace, err := s.proxies.GetNamespace(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	err = namespace.CreateProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	return &lockv1.CreateResponse{}, nil
}

func (s *lockV1ManagerServer) Close(ctx context.Context, request *lockv1.CloseRequest) (*lockv1.CloseResponse, error) {
	namespace, err := s.proxies.GetNamespace(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	err = namespace.CloseProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	return &lockv1.CloseResponse{}, nil
}

var _ lockv1.LockManagerServer = (*lockV1ManagerServer)(nil)
