// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	valuev1 "github.com/atomix/runtime-api/api/atomix/value/v1"
	"github.com/atomix/runtime-api/pkg/runtime"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"github.com/atomix/runtime-api/pkg/runtime/proxy"
	"google.golang.org/grpc"
)

// Register registers the primitive with the given runtime
func Register(server *grpc.Server, rt *runtime.Runtime) {
	proxies := proxy.NewRegistry[ValueProxy]()
	valuev1.RegisterValueManagerServer(server, newValueV1ManagerServer(proxy.NewService[ValueProxy](rt, PrimitiveType, proxies)))
	valuev1.RegisterValueServer(server, newValueV1Server(proxies))
}

// PrimitiveType is the value/v1 primitive type
var PrimitiveType = atom.NewType[ValueProxy](func(client driver.Client) (*atom.Client[ValueProxy], bool) {
	if valueClient, ok := client.(ValueClient); ok {
		return atom.NewClient[ValueProxy](valueClient.GetValue), true
	}
	return nil, false
})

type ValueClient interface {
	GetValue(ctx context.Context, name string) (ValueProxy, error)
}

type ValueProxy interface {
	atom.Atom
	valuev1.ValueServer
}
