// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	mapv1 "github.com/atomix/runtime-api/api/atomix/map/v1"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"google.golang.org/grpc"
)

var Atom = atom.New[MapProxy](clientFactory, func(server *grpc.Server, service *atom.Service[MapProxy], registry *atom.Registry[MapProxy]) {
	mapv1.RegisterMapManagerServer(server, newMapV1ManagerServer(service))
	mapv1.RegisterMapServer(server, newMapV1Server(registry))
})

// clientFactory is the counter/v1 client factory
var clientFactory = atom.NewClientFactory[MapProxy](func(client driver.Client) (*atom.Client[MapProxy], bool) {
	if counterClient, ok := client.(MapClient); ok {
		return atom.NewClient[MapProxy](counterClient.GetMap), true
	}
	return nil, false
})

type MapClient interface {
	GetMap(ctx context.Context, name string) (MapProxy, error)
}

type MapProxy interface {
	atom.Atom
	mapv1.MapServer
}
