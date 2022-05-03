// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	listv1 "github.com/atomix/runtime-api/api/atomix/list/v1"
	"github.com/atomix/runtime-api/pkg/runtime"
	"google.golang.org/grpc"
)

// Register registers the primitive with the given runtime
func Register(server *grpc.Server, rt runtime.Runtime) {
	proxies := runtime.NewProxyRegistry[ListProxy]()
	listv1.RegisterListManagerServer(server, newListV1ManagerServer(runtime.NewProxyService[ListProxy](rt, PrimitiveType, proxies)))
	listv1.RegisterListServer(server, newListV1Server(proxies))
}

// PrimitiveType is the list/v1 primitive type
var PrimitiveType = runtime.NewAtomType[ListProxy](func(client runtime.Client) (runtime.AtomClient[ListProxy], bool) {
	if listClient, ok := client.(ListClient); ok {
		return runtime.NewAtomClient[ListProxy](listClient.GetList), true
	}
	return nil, false
})

type ListClient interface {
	GetList(ctx context.Context, name string) (ListProxy, error)
}

type ListProxy interface {
	runtime.Proxy
	listv1.ListServer
}
