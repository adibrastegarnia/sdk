// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	mapv1 "github.com/atomix/runtime-api/api/atomix/map/v1"
	"github.com/atomix/runtime-api/pkg/runtime"
	"google.golang.org/grpc"
)

// Register registers the primitive with the given runtime
func Register(server *grpc.Server, rt runtime.Runtime) {
	proxies := runtime.NewProxyRegistry[MapProxy]()
	mapv1.RegisterMapManagerServer(server, newMapV1ManagerServer(runtime.NewProxyService[MapProxy](rt, PrimitiveType, proxies)))
	mapv1.RegisterMapServer(server, newMapV1Server(proxies))
}

// PrimitiveType is the map/v1 primitive type
var PrimitiveType = runtime.NewAtomType[MapProxy](func(client runtime.Client) (runtime.AtomClient[MapProxy], bool) {
	if mapClient, ok := client.(MapClient); ok {
		return runtime.NewAtomClient[MapProxy](mapClient.GetMap), true
	}
	return nil, false
})

type MapClient interface {
	GetMap(ctx context.Context, name string) (MapProxy, error)
}

type MapProxy interface {
	runtime.Proxy
	mapv1.MapServer
}
