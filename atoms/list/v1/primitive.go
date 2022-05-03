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

func newListV1Server(proxies *runtime.ProxyRegistry[ListProxy]) listv1.ListServer {
	return &listV1Server{
		proxies: proxies,
	}
}

type listV1Server struct {
	proxies *runtime.ProxyRegistry[ListProxy]
}

func (s *listV1Server) Size(ctx context.Context, request *listv1.SizeRequest) (*listv1.SizeResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.Proto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Size(ctx, request)
}

func (s *listV1Server) Append(ctx context.Context, request *listv1.AppendRequest) (*listv1.AppendResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.Proto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Append(ctx, request)
}

func (s *listV1Server) Insert(ctx context.Context, request *listv1.InsertRequest) (*listv1.InsertResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.Proto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Insert(ctx, request)
}

func (s *listV1Server) Get(ctx context.Context, request *listv1.GetRequest) (*listv1.GetResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.Proto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Get(ctx, request)
}

func (s *listV1Server) Set(ctx context.Context, request *listv1.SetRequest) (*listv1.SetResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.Proto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Set(ctx, request)
}

func (s *listV1Server) Remove(ctx context.Context, request *listv1.RemoveRequest) (*listv1.RemoveResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.Proto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Remove(ctx, request)
}

func (s *listV1Server) Clear(ctx context.Context, request *listv1.ClearRequest) (*listv1.ClearResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.Proto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Clear(ctx, request)
}

func (s *listV1Server) Events(request *listv1.EventsRequest, server listv1.List_EventsServer) error {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return errors.Proto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Events(request, server)
}

func (s *listV1Server) Elements(request *listv1.ElementsRequest, server listv1.List_ElementsServer) error {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return errors.Proto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Elements(request, server)
}

var _ listv1.ListServer = (*listV1Server)(nil)
