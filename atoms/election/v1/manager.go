// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	electionv1 "github.com/atomix/runtime-api/api/atomix/election/v1"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
)

func newLeaderElectionV1ManagerServer(proxies *atom.Service[LeaderElection]) electionv1.LeaderElectionManagerServer {
	return &leaderElectionV1ManagerServer{
		proxies: proxies,
	}
}

type leaderElectionV1ManagerServer struct {
	proxies *atom.Service[LeaderElection]
}

func (s *leaderElectionV1ManagerServer) Create(ctx context.Context, request *electionv1.CreateRequest) (*electionv1.CreateResponse, error) {
	namespace, err := s.proxies.GetCluster(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = namespace.CreateProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &electionv1.CreateResponse{}, nil
}

func (s *leaderElectionV1ManagerServer) Close(ctx context.Context, request *electionv1.CloseRequest) (*electionv1.CloseResponse, error) {
	namespace, err := s.proxies.GetCluster(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = namespace.CloseProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &electionv1.CloseResponse{}, nil
}

var _ electionv1.LeaderElectionManagerServer = (*leaderElectionV1ManagerServer)(nil)
