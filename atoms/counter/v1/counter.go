// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	counterv1 "github.com/atomix/runtime-api/api/atomix/counter/v1"
	"github.com/atomix/runtime-api/pkg/runtime"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"github.com/atomix/runtime-api/pkg/runtime/proxy"
	"google.golang.org/grpc"
)

// Register registers the primitive with the given runtime
func Register(server *grpc.Server, rt *runtime.Runtime) {
	proxies := proxy.NewRegistry[CounterProxy]()
	counterv1.RegisterCounterManagerServer(server, newCounterV1ManagerServer(proxy.NewService[CounterProxy](rt, PrimitiveType, proxies)))
	counterv1.RegisterCounterServer(server, newCounterV1Server(proxies))
}

// PrimitiveType is the counter/v1 primitive type
var PrimitiveType = atom.NewType[CounterProxy](func(client driver.Client) (*atom.Client[CounterProxy], bool) {
	if counterClient, ok := client.(CounterClient); ok {
		return atom.NewClient[CounterProxy](counterClient.GetCounter), true
	}
	return nil, false
})

type CounterClient interface {
	GetCounter(ctx context.Context, name string) (CounterProxy, error)
}

type CounterProxy interface {
	atom.Atom
	counterv1.CounterServer
}
