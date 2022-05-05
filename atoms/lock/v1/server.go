// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	"github.com/atomix/sdk/pkg/errors"
	"github.com/atomix/sdk/pkg/runtime/atom"
)

func newLockV1Server(proxies *atom.Registry[Lock]) v1.LockServer {
	return &lockV1Server{
		proxies: proxies,
	}
}

type lockV1Server struct {
	proxies *atom.Registry[Lock]
}

func (s *lockV1Server) Lock(ctx context.Context, request *v1.LockRequest) (*v1.LockResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Lock(ctx, request)
}

func (s *lockV1Server) Unlock(ctx context.Context, request *v1.UnlockRequest) (*v1.UnlockResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Unlock(ctx, request)
}

func (s *lockV1Server) GetLock(ctx context.Context, request *v1.GetLockRequest) (*v1.GetLockResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.GetLock(ctx, request)
}

var _ v1.LockServer = (*lockV1Server)(nil)
