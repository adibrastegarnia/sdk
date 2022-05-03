// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	listv1 "github.com/atomix/runtime-api/api/atomix/list/v1"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"google.golang.org/grpc"
)

var Atom = atom.New[ListProxy](clientFactory, func(server *grpc.Server, service *atom.Service[ListProxy], registry *atom.Registry[ListProxy]) {
	listv1.RegisterListManagerServer(server, newListV1ManagerServer(service))
	listv1.RegisterListServer(server, newListV1Server(registry))
})

// clientFactory is the counter/v1 client factory
var clientFactory = atom.NewClientFactory[ListProxy](func(client driver.Client) (*atom.Client[ListProxy], bool) {
	if counterClient, ok := client.(ListClient); ok {
		return atom.NewClient[ListProxy](counterClient.GetList), true
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
