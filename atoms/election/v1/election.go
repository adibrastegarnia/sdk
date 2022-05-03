// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	electionv1 "github.com/atomix/runtime-api/api/atomix/election/v1"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"google.golang.org/grpc"
)

var Atom = atom.New[LeaderElection](clientFactory, func(server *grpc.Server, service *atom.Service[LeaderElection], registry *atom.Registry[LeaderElection]) {
	electionv1.RegisterLeaderElectionManagerServer(server, newLeaderElectionV1ManagerServer(service))
	electionv1.RegisterLeaderElectionServer(server, newLeaderElectionV1Server(registry))
})

// clientFactory is the election/v1 client factory
var clientFactory = atom.NewClientFactory[LeaderElection](func(client driver.Client) (*atom.Client[LeaderElection], bool) {
	if electionClient, ok := client.(LeaderElectionClient); ok {
		return atom.NewClient[LeaderElection](electionClient.GetLeaderElection), true
	}
	return nil, false
})

type LeaderElectionClient interface {
	GetLeaderElection(ctx context.Context, name string) (LeaderElection, error)
}

type LeaderElection interface {
	atom.Atom
	electionv1.LeaderElectionServer
}
