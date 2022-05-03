// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	lockv1 "github.com/atomix/runtime-api/api/atomix/lock/v1"
	"github.com/atomix/runtime-api/pkg/runtime"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"github.com/atomix/runtime-api/pkg/runtime/proxy"
	"google.golang.org/grpc"
)

// Register registers the primitive with the given runtime
func Register(server *grpc.Server, rt *runtime.Runtime) {
	proxies := proxy.NewRegistry[LockProxy]()
	lockv1.RegisterLockManagerServer(server, newLockV1ManagerServer(proxy.NewService[LockProxy](rt, PrimitiveType, proxies)))
	lockv1.RegisterLockServer(server, newLockV1Server(proxies))
}

// PrimitiveType is the lock/v1 primitive type
var PrimitiveType = atom.NewType[LockProxy](func(client driver.Client) (*atom.Client[LockProxy], bool) {
	if lockClient, ok := client.(LockClient); ok {
		return atom.NewClient[LockProxy](lockClient.GetLock), true
	}
	return nil, false
})

type LockClient interface {
	GetLock(ctx context.Context, name string) (LockProxy, error)
}

type LockProxy interface {
	atom.Atom
	lockv1.LockServer
}
