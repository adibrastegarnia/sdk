// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	lockv1 "github.com/atomix/runtime-api/api/atomix/lock/v1"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"google.golang.org/grpc"
)

var Atom = atom.New[LockProxy](clientFactory, func(server *grpc.Server, service *atom.Service[LockProxy], registry *atom.Registry[LockProxy]) {
	lockv1.RegisterLockManagerServer(server, newLockV1ManagerServer(service))
	lockv1.RegisterLockServer(server, newLockV1Server(registry))
})

// clientFactory is the counter/v1 client factory
var clientFactory = atom.NewClientFactory[LockProxy](func(client driver.Client) (*atom.Client[LockProxy], bool) {
	if counterClient, ok := client.(LockClient); ok {
		return atom.NewClient[LockProxy](counterClient.GetLock), true
	}
	return nil, false
})

type LockClient interface {
	GetLock(ctx context.Context, name string) (LockProxy, error)
}

type LockProxy interface {
	atom.Atom
	lockv1.LockServer
}
