// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	counterv1 "github.com/atomix/runtime-api/api/atomix/counter/v1"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"google.golang.org/grpc"
)

var Atom = atom.New[Counter](clientFactory, func(server *grpc.Server, service *atom.Service[Counter], registry *atom.Registry[Counter]) {
	counterv1.RegisterCounterManagerServer(server, newCounterV1ManagerServer(service))
	counterv1.RegisterCounterServer(server, newCounterV1Server(registry))
})

// clientFactory is the counter/v1 client factory
var clientFactory = atom.NewClientFactory[Counter](func(client driver.Client) (*atom.Client[Counter], bool) {
	if counterClient, ok := client.(CounterClient); ok {
		return atom.NewClient[Counter](counterClient.GetCounter), true
	}
	return nil, false
})

type CounterClient interface {
	GetCounter(ctx context.Context, name string) (Counter, error)
}

type Counter interface {
	atom.Atom
	counterv1.CounterServer
}
