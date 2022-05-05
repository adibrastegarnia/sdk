// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	"github.com/atomix/api/pkg/atomix/election/v1"
	"github.com/atomix/sdk/pkg/errors"
	"github.com/atomix/sdk/pkg/runtime/atom"
)

func newLeaderElectionV1Server(proxies *atom.Registry[LeaderElection]) v1.LeaderElectionServer {
	return &leaderElectionV1Server{
		proxies: proxies,
	}
}

type leaderElectionV1Server struct {
	proxies *atom.Registry[LeaderElection]
}

func (s *leaderElectionV1Server) Enter(ctx context.Context, request *v1.EnterRequest) (*v1.EnterResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Enter(ctx, request)
}

func (s *leaderElectionV1Server) Withdraw(ctx context.Context, request *v1.WithdrawRequest) (*v1.WithdrawResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Withdraw(ctx, request)
}

func (s *leaderElectionV1Server) Anoint(ctx context.Context, request *v1.AnointRequest) (*v1.AnointResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Anoint(ctx, request)
}

func (s *leaderElectionV1Server) Promote(ctx context.Context, request *v1.PromoteRequest) (*v1.PromoteResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Promote(ctx, request)
}

func (s *leaderElectionV1Server) Evict(ctx context.Context, request *v1.EvictRequest) (*v1.EvictResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Evict(ctx, request)
}

func (s *leaderElectionV1Server) GetTerm(ctx context.Context, request *v1.GetTermRequest) (*v1.GetTermResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.GetTerm(ctx, request)
}

func (s *leaderElectionV1Server) Events(request *v1.EventsRequest, server v1.LeaderElection_EventsServer) error {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Events(request, server)
}

var _ v1.LeaderElectionServer = (*leaderElectionV1Server)(nil)
