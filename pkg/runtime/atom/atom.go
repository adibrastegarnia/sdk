// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package atom

import (
	"context"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
)

type Atom interface {
	Close(ctx context.Context) error
}

// NewType creates a new Type for the given atom type
func NewType[T Atom](getter func(client driver.Client) (*Client[T], bool)) *Type[T] {
	return &Type[T]{
		getter: getter,
	}
}

type Type[T Atom] struct {
	getter func(client driver.Client) (*Client[T], bool)
}

func (t *Type[T]) Client(client driver.Client) (*Client[T], bool) {
	return t.getter(client)
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

func (c *Client[T]) GetProxy(ctx context.Context, name string) (T, error) {
	return c.getter(ctx, name)
}
