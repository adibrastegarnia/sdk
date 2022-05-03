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

var Atom = atom.New[IndexedMap](clientFactory, func(server *grpc.Server, service *atom.Service[IndexedMap], registry *atom.Registry[IndexedMap]) {
	indexed_mapv1.RegisterIndexedMapManagerServer(server, newIndexedMapV1ManagerServer(service))
	indexed_mapv1.RegisterIndexedMapServer(server, newIndexedMapV1Server(registry))
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
	indexed_mapv1.IndexedMapServer
}
