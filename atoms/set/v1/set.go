// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	setv1 "github.com/atomix/runtime-api/api/atomix/set/v1"
	"github.com/atomix/runtime-api/pkg/runtime"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"github.com/atomix/runtime-api/pkg/runtime/proxy"
	"google.golang.org/grpc"
)

// Register registers the primitive with the given runtime
func Register(server *grpc.Server, rt *runtime.Runtime) {
	proxies := proxy.NewRegistry[SetProxy]()
	setv1.RegisterSetManagerServer(server, newSetV1ManagerServer(proxy.NewService[SetProxy](rt, PrimitiveType, proxies)))
	setv1.RegisterSetServer(server, newSetV1Server(proxies))
}

// PrimitiveType is the set/v1 primitive type
var PrimitiveType = atom.NewType[SetProxy](func(client driver.Client) (*atom.Client[SetProxy], bool) {
	if setClient, ok := client.(SetClient); ok {
		return atom.NewClient[SetProxy](setClient.GetSet), true
	}
	return nil, false
})

type SetClient interface {
	GetSet(ctx context.Context, name string) (SetProxy, error)
}

type SetProxy interface {
	atom.Atom
	setv1.SetServer
}
