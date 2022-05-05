// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	"github.com/atomix/api/pkg/atomix/indexed_map/v1"
	"github.com/atomix/sdk/pkg/runtime/atom"
	"github.com/atomix/sdk/pkg/runtime/driver"
	"google.golang.org/grpc"
)

var Atom = atom.New[IndexedMap](clientFactory, func(server *grpc.Server, service *atom.Service[IndexedMap], registry *atom.Registry[IndexedMap]) {
	v1.RegisterIndexedMapManagerServer(server, newIndexedMapV1ManagerServer(service))
	v1.RegisterIndexedMapServer(server, newIndexedMapV1Server(registry))
})

// clientFactory is the indexed_map/v1 client factory
var clientFactory = atom.NewClientFactory[IndexedMap](func(client driver.Client) (*atom.Client[IndexedMap], bool) {
	if indexedMapClient, ok := client.(IndexedMapClient); ok {
		return atom.NewClient[IndexedMap](indexedMapClient.GetIndexedMap), true
	}
	return nil, false
})

type IndexedMapClient interface {
	GetIndexedMap(ctx context.Context, name string) (IndexedMap, error)
}

type IndexedMap interface {
	atom.Atom
	v1.IndexedMapServer
}
