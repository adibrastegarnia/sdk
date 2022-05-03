// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	counterv1 "github.com/atomix/runtime-api/api/atomix/counter/v1"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/runtime"
)

func newCounterV1Server(proxies *runtime.ProxyRegistry[CounterProxy]) counterv1.CounterServer {
	return &counterV1Server{
		proxies: proxies,
	}
}

type counterV1Server struct {
	proxies *runtime.ProxyRegistry[CounterProxy]
}

func (s *counterV1Server) Set(ctx context.Context, request *counterv1.SetRequest) (*counterv1.SetResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.Proto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Set(ctx, request)
}

func (s *counterV1Server) Get(ctx context.Context, request *counterv1.GetRequest) (*counterv1.GetResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.Proto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Get(ctx, request)
}

func (s *counterV1Server) Increment(ctx context.Context, request *counterv1.IncrementRequest) (*counterv1.IncrementResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.Proto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Increment(ctx, request)
}

func (s *counterV1Server) Decrement(ctx context.Context, request *counterv1.DecrementRequest) (*counterv1.DecrementResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.Proto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Decrement(ctx, request)
}

var _ counterv1.CounterServer = (*counterV1Server)(nil)
