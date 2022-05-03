// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	indexed_mapv1 "github.com/atomix/runtime-api/api/atomix/indexed_map/v1"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"google.golang.org/grpc"
)

var Atom = atom.New[IndexedMapProxy](clientFactory, func(server *grpc.Server, service *atom.Service[IndexedMapProxy], registry *atom.Registry[IndexedMapProxy]) {
	indexed_mapv1.RegisterIndexedMapManagerServer(server, newIndexedMapV1ManagerServer(service))
	indexed_mapv1.RegisterIndexedMapServer(server, newIndexedMapV1Server(registry))
})

// clientFactory is the counter/v1 client factory
var clientFactory = atom.NewClientFactory[IndexedMapProxy](func(client driver.Client) (*atom.Client[IndexedMapProxy], bool) {
	if counterClient, ok := client.(IndexedMapClient); ok {
		return atom.NewClient[IndexedMapProxy](counterClient.GetIndexedMap), true
	}
	return nil, false
})

type IndexedMapClient interface {
	GetIndexedMap(ctx context.Context, name string) (IndexedMapProxy, error)
}

type IndexedMapProxy interface {
	atom.Atom
	indexed_mapv1.IndexedMapServer
}
