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

func newIndexedMapV1Server(proxies *runtime.ProxyRegistry[IndexedMapProxy]) indexed_mapv1.IndexedMapServer {
	return &indexedMapV1Server{
		proxies: proxies,
	}
}

type indexedMapV1Server struct {
	proxies *runtime.ProxyRegistry[IndexedMapProxy]
}

func (s *indexedMapV1Server) Size(ctx context.Context, request *indexed_mapv1.SizeRequest) (*indexed_mapv1.SizeResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Size(ctx, request)
}

func (s *indexedMapV1Server) Put(ctx context.Context, request *indexed_mapv1.PutRequest) (*indexed_mapv1.PutResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Put(ctx, request)
}

func (s *indexedMapV1Server) Get(ctx context.Context, request *indexed_mapv1.GetRequest) (*indexed_mapv1.GetResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Get(ctx, request)
}

func (s *indexedMapV1Server) FirstEntry(ctx context.Context, request *indexed_mapv1.FirstEntryRequest) (*indexed_mapv1.FirstEntryResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.FirstEntry(ctx, request)
}

func (s *indexedMapV1Server) LastEntry(ctx context.Context, request *indexed_mapv1.LastEntryRequest) (*indexed_mapv1.LastEntryResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.LastEntry(ctx, request)
}

func (s *indexedMapV1Server) PrevEntry(ctx context.Context, request *indexed_mapv1.PrevEntryRequest) (*indexed_mapv1.PrevEntryResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.PrevEntry(ctx, request)
}

func (s *indexedMapV1Server) NextEntry(ctx context.Context, request *indexed_mapv1.NextEntryRequest) (*indexed_mapv1.NextEntryResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.NextEntry(ctx, request)
}

func (s *indexedMapV1Server) Remove(ctx context.Context, request *indexed_mapv1.RemoveRequest) (*indexed_mapv1.RemoveResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Remove(ctx, request)
}

func (s *indexedMapV1Server) Clear(ctx context.Context, request *indexed_mapv1.ClearRequest) (*indexed_mapv1.ClearResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Clear(ctx, request)
}

func (s *indexedMapV1Server) Events(request *indexed_mapv1.EventsRequest, server indexed_mapv1.IndexedMap_EventsServer) error {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Events(request, server)
}

func (s *indexedMapV1Server) Entries(request *indexed_mapv1.EntriesRequest, server indexed_mapv1.IndexedMap_EntriesServer) error {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Entries(request, server)
}

var _ indexed_mapv1.IndexedMapServer = (*indexedMapV1Server)(nil)
