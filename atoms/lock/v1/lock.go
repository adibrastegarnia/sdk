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

var Atom = atom.New[Lock](clientFactory, func(server *grpc.Server, service *atom.Service[Lock], registry *atom.Registry[Lock]) {
	lockv1.RegisterLockManagerServer(server, newLockV1ManagerServer(service))
	lockv1.RegisterLockServer(server, newLockV1Server(registry))
})

// clientFactory is the lock/v1 client factory
var clientFactory = atom.NewClientFactory[Lock](func(client driver.Client) (*atom.Client[Lock], bool) {
	if lockClient, ok := client.(LockClient); ok {
		return atom.NewClient[Lock](lockClient.GetLock), true
	}
	return nil, false
})

type LockClient interface {
	GetLock(ctx context.Context, name string) (Lock, error)
}

type Lock interface {
	atom.Atom
	lockv1.LockServer
}
