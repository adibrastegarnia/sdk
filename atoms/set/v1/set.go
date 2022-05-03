// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	setv1 "github.com/atomix/runtime-api/api/atomix/set/v1"
	"github.com/atomix/runtime-api/pkg/runtime"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"google.golang.org/grpc"
)

// Register registers the primitive with the given runtime
func Register(server *grpc.Server, rt *runtime.Runtime) {
	proxies := runtime.NewProxyRegistry[SetProxy]()
	setv1.RegisterSetManagerServer(server, newSetV1ManagerServer(runtime.NewProxyService[SetProxy](rt, PrimitiveType, proxies)))
	setv1.RegisterSetServer(server, newSetV1Server(proxies))
}

// PrimitiveType is the set/v1 primitive type
var PrimitiveType = runtime.NewAtomType[SetProxy](func(client driver.Client) (*runtime.AtomClient[SetProxy], bool) {
	if setClient, ok := client.(SetClient); ok {
		return runtime.NewAtomClient[SetProxy](setClient.GetSet), true
	}
	return nil, false
})

type SetClient interface {
	GetSet(ctx context.Context, name string) (SetProxy, error)
}

type SetProxy interface {
	runtime.Atom
	setv1.SetServer
}
