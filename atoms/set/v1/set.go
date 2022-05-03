// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	setv1 "github.com/atomix/runtime-api/api/atomix/set/v1"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"google.golang.org/grpc"
)

var Atom = atom.New[Set](clientFactory, func(server *grpc.Server, service *atom.Service[Set], registry *atom.Registry[Set]) {
	setv1.RegisterSetManagerServer(server, newSetV1ManagerServer(service))
	setv1.RegisterSetServer(server, newSetV1Server(registry))
})

// clientFactory is the set/v1 client factory
var clientFactory = atom.NewClientFactory[Set](func(client driver.Client) (*atom.Client[Set], bool) {
	if setClient, ok := client.(SetClient); ok {
		return atom.NewClient[Set](setClient.GetSet), true
	}
	return nil, false
})

type SetClient interface {
	GetSet(ctx context.Context, name string) (Set, error)
}

type Set interface {
	atom.Atom
	setv1.SetServer
}
