// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	"github.com/atomix/api/pkg/atomix/indexed_map/v1"
	"github.com/atomix/sdk/pkg/atom"
	"github.com/atomix/sdk/pkg/errors"
)

func newIndexedMapV1ManagerServer(proxies *atom.Service[IndexedMap]) v1.IndexedMapManagerServer {
	return &indexedMapV1ManagerServer{
		proxies: proxies,
	}
}

type indexedMapV1ManagerServer struct {
	proxies *atom.Service[IndexedMap]
}

func (s *indexedMapV1ManagerServer) Create(ctx context.Context, request *v1.CreateRequest) (*v1.CreateResponse, error) {
	namespace, err := s.proxies.GetCluster(ctx, request.Headers.Cluster)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = namespace.CreateProxy(ctx, request.Atom)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &v1.CreateResponse{}, nil
}

func (s *indexedMapV1ManagerServer) Close(ctx context.Context, request *v1.CloseRequest) (*v1.CloseResponse, error) {
	namespace, err := s.proxies.GetCluster(ctx, request.Headers.Cluster)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = namespace.CloseProxy(ctx, request.Atom)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &v1.CloseResponse{}, nil
}

var _ v1.IndexedMapManagerServer = (*indexedMapV1ManagerServer)(nil)
