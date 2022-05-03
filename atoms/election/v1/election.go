// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	electionv1 "github.com/atomix/runtime-api/api/atomix/election/v1"
	"github.com/atomix/runtime-api/pkg/runtime"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"github.com/atomix/runtime-api/pkg/runtime/proxy"
	"google.golang.org/grpc"
)

// Register registers the primitive with the given runtime
func Register(server *grpc.Server, rt *runtime.Runtime) {
	proxies := proxy.NewRegistry[LeaderElectionProxy]()
	electionv1.RegisterLeaderElectionManagerServer(server, newLeaderElectionV1ManagerServer(proxy.NewService[LeaderElectionProxy](rt, PrimitiveType, proxies)))
	electionv1.RegisterLeaderElectionServer(server, newLeaderElectionV1Server(proxies))
}

// PrimitiveType is the election/v1 primitive type
var PrimitiveType = atom.NewType[LeaderElectionProxy](func(client driver.Client) (*atom.Client[LeaderElectionProxy], bool) {
	if electionClient, ok := client.(LeaderElectionClient); ok {
		return atom.NewClient[LeaderElectionProxy](electionClient.GetLeaderElection), true
	}
	return nil, false
})

type LeaderElectionClient interface {
	GetLeaderElection(ctx context.Context, name string) (LeaderElectionProxy, error)
}

type LeaderElectionProxy interface {
	atom.Atom
	electionv1.LeaderElectionServer
}
