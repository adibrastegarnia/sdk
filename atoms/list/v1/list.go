// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	listv1 "github.com/atomix/runtime-api/api/atomix/list/v1"
	"github.com/atomix/runtime-api/pkg/runtime"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"github.com/atomix/runtime-api/pkg/runtime/proxy"
	"google.golang.org/grpc"
)

// Register registers the primitive with the given runtime
func Register(server *grpc.Server, rt *runtime.Runtime) {
	proxies := proxy.NewRegistry[ListProxy]()
	listv1.RegisterListManagerServer(server, newListV1ManagerServer(proxy.NewService[ListProxy](rt, PrimitiveType, proxies)))
	listv1.RegisterListServer(server, newListV1Server(proxies))
}

// PrimitiveType is the list/v1 primitive type
var PrimitiveType = atom.NewType[ListProxy](func(client driver.Client) (*atom.Client[ListProxy], bool) {
	if listClient, ok := client.(ListClient); ok {
		return atom.NewClient[ListProxy](listClient.GetList), true
	}
	return nil, false
})

type ListClient interface {
	GetList(ctx context.Context, name string) (ListProxy, error)
}

type ListProxy interface {
	atom.Atom
	listv1.ListServer
}
