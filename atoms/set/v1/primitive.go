// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	setv1 "github.com/atomix/runtime-api/api/atomix/set/v1"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/runtime/proxy"
)

func newSetV1Server(proxies *proxy.Registry[SetProxy]) setv1.SetServer {
	return &setV1Server{
		proxies: proxies,
	}
}

type setV1Server struct {
	proxies *proxy.Registry[SetProxy]
}

func (s *setV1Server) Size(ctx context.Context, request *setv1.SizeRequest) (*setv1.SizeResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Size(ctx, request)
}

func (s *setV1Server) Contains(ctx context.Context, request *setv1.ContainsRequest) (*setv1.ContainsResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Contains(ctx, request)
}

func (s *setV1Server) Add(ctx context.Context, request *setv1.AddRequest) (*setv1.AddResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Add(ctx, request)
}

func (s *setV1Server) Remove(ctx context.Context, request *setv1.RemoveRequest) (*setv1.RemoveResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Remove(ctx, request)
}

func (s *setV1Server) Clear(ctx context.Context, request *setv1.ClearRequest) (*setv1.ClearResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Clear(ctx, request)
}

func (s *setV1Server) Events(request *setv1.EventsRequest, server setv1.Set_EventsServer) error {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Events(request, server)
}

func (s *setV1Server) Elements(request *setv1.ElementsRequest, server setv1.Set_ElementsServer) error {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Elements(request, server)
}

var _ setv1.SetServer = (*setV1Server)(nil)
