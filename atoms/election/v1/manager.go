// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	electionv1 "github.com/atomix/runtime-api/api/atomix/election/v1"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/runtime"
)

func newLeaderElectionV1ManagerServer(proxies runtime.ProxyService) electionv1.LeaderElectionManagerServer {
	return &leaderElectionV1ManagerServer{
		proxies: proxies,
	}
}

type leaderElectionV1ManagerServer struct {
	proxies runtime.ProxyService
}

func (s *leaderElectionV1ManagerServer) Create(ctx context.Context, request *electionv1.CreateRequest) (*electionv1.CreateResponse, error) {
	namespace, err := s.proxies.GetNamespace(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	err = namespace.CreateProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	return &electionv1.CreateResponse{}, nil
}

func (s *leaderElectionV1ManagerServer) Close(ctx context.Context, request *electionv1.CloseRequest) (*electionv1.CloseResponse, error) {
	namespace, err := s.proxies.GetNamespace(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	err = namespace.CloseProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.Proto(err)
	}
	return &electionv1.CloseResponse{}, nil
}

var _ electionv1.LeaderElectionManagerServer = (*leaderElectionV1ManagerServer)(nil)
