// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"context"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
)

type Atom interface {
	Close(ctx context.Context) error
}

// NewAtomType creates a new type for the given atom type
func NewAtomType[T Atom](getter func(client driver.Client) (*AtomClient[T], bool)) *AtomType[T] {
	return &AtomType[T]{
		getter: getter,
	}
}

type AtomType[T Atom] struct {
	getter func(client driver.Client) (*AtomClient[T], bool)
}

func (t *AtomType[T]) Client(client driver.Client) (*AtomClient[T], bool) {
	return t.getter(client)
}

// NewAtomClient creates a new client for the given atom type
func NewAtomClient[T Atom](getter func(ctx context.Context, name string) (T, error)) *AtomClient[T] {
	return &AtomClient[T]{
		getter: getter,
	}
}

type AtomClient[T Atom] struct {
	getter func(ctx context.Context, name string) (T, error)
}

func (c *AtomClient[T]) GetProxy(ctx context.Context, name string) (T, error) {
	return c.getter(ctx, name)
}
