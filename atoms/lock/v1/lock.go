// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	lockv1 "github.com/atomix/runtime-api/api/atomix/lock/v1"
	"github.com/atomix/runtime-api/pkg/runtime"
	"google.golang.org/grpc"
)

// Register registers the primitive with the given runtime
func Register(server *grpc.Server, rt runtime.Runtime) {
	proxies := runtime.NewProxyRegistry[LockProxy]()
	lockv1.RegisterLockManagerServer(server, newLockV1ManagerServer(runtime.NewProxyService[LockProxy](rt, PrimitiveType, proxies)))
	lockv1.RegisterLockServer(server, newLockV1Server(proxies))
}

// PrimitiveType is the lock/v1 primitive type
var PrimitiveType = runtime.NewAtomType[LockProxy](func(client runtime.Client) (runtime.AtomClient[LockProxy], bool) {
	if lockClient, ok := client.(LockClient); ok {
		return runtime.NewAtomClient[LockProxy](lockClient.GetLock), true
	}
	return nil, false
})

type LockClient interface {
	GetLock(ctx context.Context, name string) (LockProxy, error)
}

type LockProxy interface {
	runtime.Proxy
	lockv1.LockServer
}
