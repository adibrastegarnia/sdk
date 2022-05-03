// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	setv1 "github.com/atomix/runtime-api/api/atomix/set/v1"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"google.golang.org/grpc"
)

var Atom = atom.New[SetProxy](clientFactory, func(server *grpc.Server, service *atom.Service[SetProxy], registry *atom.Registry[SetProxy]) {
	setv1.RegisterSetManagerServer(server, newSetV1ManagerServer(service))
	setv1.RegisterSetServer(server, newSetV1Server(registry))
})

// clientFactory is the counter/v1 client factory
var clientFactory = atom.NewClientFactory[SetProxy](func(client driver.Client) (*atom.Client[SetProxy], bool) {
	if counterClient, ok := client.(SetClient); ok {
		return atom.NewClient[SetProxy](counterClient.GetSet), true
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
