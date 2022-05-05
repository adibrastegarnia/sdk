// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	"github.com/atomix/api/pkg/atomix/map/v1"
	"github.com/atomix/sdk/pkg/runtime/atom"
	"github.com/atomix/sdk/pkg/runtime/driver"
	"google.golang.org/grpc"
)

var Atom = atom.New[Map](clientFactory, func(server *grpc.Server, service *atom.Service[Map], registry *atom.Registry[Map]) {
	v1.RegisterMapManagerServer(server, newMapV1ManagerServer(service))
	v1.RegisterMapServer(server, newMapV1Server(registry))
})

// clientFactory is the map/v1 client factory
var clientFactory = atom.NewClientFactory[Map](func(client driver.Client) (*atom.Client[Map], bool) {
	if mapClient, ok := client.(MapClient); ok {
		return atom.NewClient[Map](mapClient.GetMap), true
	}
	return nil, false
})

type MapClient interface {
	GetMap(ctx context.Context, name string) (Map, error)
}

type Map interface {
	atom.Atom
	v1.MapServer
}
