// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	valuev1 "github.com/atomix/runtime-api/api/atomix/value/v1"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/runtime"
)

func newValueV1Server(proxies *runtime.ProxyRegistry[ValueProxy]) valuev1.ValueServer {
	return &valueV1Server{
		proxies: proxies,
	}
}

type valueV1Server struct {
	proxies *runtime.ProxyRegistry[ValueProxy]
}

func (s *valueV1Server) Set(ctx context.Context, request *valuev1.SetRequest) (*valuev1.SetResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.Proto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Set(ctx, request)
}

func (s *valueV1Server) Get(ctx context.Context, request *valuev1.GetRequest) (*valuev1.GetResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.Proto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Get(ctx, request)
}

func (s *valueV1Server) Events(request *valuev1.EventsRequest, server valuev1.Value_EventsServer) error {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return errors.Proto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Events(request, server)
}

var _ valuev1.ValueServer = (*valueV1Server)(nil)
