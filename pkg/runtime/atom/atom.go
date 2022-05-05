// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package atom

import (
	"context"
	"github.com/atomix/sdk/pkg/runtime/driver"
	"google.golang.org/grpc"
)

type Registrar[T Atom] func(*grpc.Server, *Service[T], *Registry[T])

func New[T Atom](t *ClientFactory[T], registrar Registrar[T]) *Type[T] {
	return &Type[T]{
		factory:   t,
		registrar: registrar,
		registry:  NewRegistry[T](),
	}
}

type Type[T Atom] struct {
	factory   *ClientFactory[T]
	registrar Registrar[T]
	registry  *Registry[T]
}

func (a *Type[T]) Register(server *grpc.Server, connector Connector) {
	a.registrar(server, NewService[T](connector, a.factory, a.registry), a.registry)
}

type Atom interface {
	Close(ctx context.Context) error
}

// NewClientFactory creates a new ClientFactory for an atom
func NewClientFactory[T Atom](factory func(client driver.Client) (*Client[T], bool)) *ClientFactory[T] {
	return &ClientFactory[T]{
		factory: factory,
	}
}

type ClientFactory[T Atom] struct {
	factory func(client driver.Client) (*Client[T], bool)
}

func (t *ClientFactory[T]) GetClient(client driver.Client) (*Client[T], bool) {
	return t.factory(client)
}

// NewClient creates a new client for the given atom type
func NewClient[T Atom](getter func(ctx context.Context, name string) (T, error)) *Client[T] {
	return &Client[T]{
		getter: getter,
	}
}

type Client[T Atom] struct {
	getter func(ctx context.Context, name string) (T, error)
}

func (c *Client[T]) GetAtom(ctx context.Context, name string) (T, error) {
	return c.getter(ctx, name)
}
