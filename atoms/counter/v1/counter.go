// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	counterv1 "github.com/atomix/runtime-api/api/atomix/counter/v1"
	"github.com/atomix/runtime-api/pkg/runtime"
	"google.golang.org/grpc"
)

// Register registers the primitive with the given runtime
func Register(server *grpc.Server, rt runtime.Runtime) {
	proxies := runtime.NewProxyRegistry[CounterProxy]()
	counterv1.RegisterCounterManagerServer(server, newCounterV1ManagerServer(runtime.NewProxyService[CounterProxy](rt, PrimitiveType, proxies)))
	counterv1.RegisterCounterServer(server, newCounterV1Server(proxies))
}

// PrimitiveType is the counter/v1 primitive type
var PrimitiveType = runtime.NewPrimitiveType[CounterProxy](func(client runtime.Client) (runtime.PrimitiveClient[CounterProxy], bool) {
	if counterClient, ok := client.(CounterClient); ok {
		return runtime.NewPrimitiveClient[CounterProxy](counterClient.GetCounter), true
	}
	return nil, false
})

type CounterClient interface {
	GetCounter(ctx context.Context, name string) (CounterProxy, error)
}

type CounterProxy interface {
	runtime.Proxy
	counterv1.CounterServer
}
