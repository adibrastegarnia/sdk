// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	indexed_mapv1 "github.com/atomix/runtime-api/api/atomix/indexed_map/v1"
	"github.com/atomix/runtime-api/pkg/runtime"
	"google.golang.org/grpc"
)

// Register registers the primitive with the given runtime
func Register(server *grpc.Server, rt runtime.Runtime) {
	proxies := runtime.NewProxyRegistry[IndexedMapProxy]()
	indexed_mapv1.RegisterIndexedMapManagerServer(server, newIndexedMapV1ManagerServer(runtime.NewProxyService[IndexedMapProxy](rt, PrimitiveType, proxies)))
	indexed_mapv1.RegisterIndexedMapServer(server, newIndexedMapV1Server(proxies))
}

// PrimitiveType is the indexed_map/v1 primitive type
var PrimitiveType = runtime.NewAtomType[IndexedMapProxy](func(client runtime.Client) (*runtime.AtomClient[IndexedMapProxy], bool) {
	if indexed_mapClient, ok := client.(IndexedMapClient); ok {
		return runtime.NewAtomClient[IndexedMapProxy](indexed_mapClient.GetIndexedMap), true
	}
	return nil, false
})

type IndexedMapClient interface {
	GetIndexedMap(ctx context.Context, name string) (IndexedMapProxy, error)
}

type IndexedMapProxy interface {
	runtime.Proxy
	indexed_mapv1.IndexedMapServer
}
