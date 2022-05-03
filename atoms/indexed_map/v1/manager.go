// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	indexed_mapv1 "github.com/atomix/runtime-api/api/atomix/indexed_map/v1"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/runtime"
)

func newIndexedMapV1ManagerServer(proxies *runtime.ProxyService[IndexedMapProxy]) indexed_mapv1.IndexedMapManagerServer {
	return &indexedMapV1ManagerServer{
		proxies: proxies,
	}
}

type indexedMapV1ManagerServer struct {
	proxies *runtime.ProxyService[IndexedMapProxy]
}

func (s *indexedMapV1ManagerServer) Create(ctx context.Context, request *indexed_mapv1.CreateRequest) (*indexed_mapv1.CreateResponse, error) {
	namespace, err := s.proxies.GetCluster(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = namespace.CreateProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &indexed_mapv1.CreateResponse{}, nil
}

func (s *indexedMapV1ManagerServer) Close(ctx context.Context, request *indexed_mapv1.CloseRequest) (*indexed_mapv1.CloseResponse, error) {
	namespace, err := s.proxies.GetCluster(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = namespace.CloseProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &indexed_mapv1.CloseResponse{}, nil
}

var _ indexed_mapv1.IndexedMapManagerServer = (*indexedMapV1ManagerServer)(nil)
