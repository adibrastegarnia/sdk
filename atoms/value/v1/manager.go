// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	valuev1 "github.com/atomix/runtime-api/api/atomix/value/v1"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
)

func newValueV1ManagerServer(proxies *atom.Service[ValueProxy]) valuev1.ValueManagerServer {
	return &valueV1ManagerServer{
		proxies: proxies,
	}
}

type valueV1ManagerServer struct {
	proxies *atom.Service[ValueProxy]
}

func (s *valueV1ManagerServer) Create(ctx context.Context, request *valuev1.CreateRequest) (*valuev1.CreateResponse, error) {
	namespace, err := s.proxies.GetCluster(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = namespace.CreateProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &valuev1.CreateResponse{}, nil
}

func (s *valueV1ManagerServer) Close(ctx context.Context, request *valuev1.CloseRequest) (*valuev1.CloseResponse, error) {
	namespace, err := s.proxies.GetCluster(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = namespace.CloseProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &valuev1.CloseResponse{}, nil
}

var _ valuev1.ValueManagerServer = (*valueV1ManagerServer)(nil)
