// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"context"
)

// AtomType identifies and operates on a specific type of atom
type AtomType[P Proxy] interface {
	Client(client Client) (AtomClient[P], bool)
}

// NewAtomType creates a new type for the given atom type
func NewAtomType[P Proxy](getter func(client Client) (AtomClient[P], bool)) AtomType[P] {
	return &atomType[P]{
		getter: getter,
	}
}

type atomType[P Proxy] struct {
	getter func(client Client) (AtomClient[P], bool)
}

func (t *atomType[P]) Client(client Client) (AtomClient[P], bool) {
	return t.getter(client)
}

var _ AtomType[Proxy] = (*atomType[Proxy])(nil)

// AtomClient is a client for a specific type of atom
type AtomClient[P Proxy] interface {
	GetProxy(ctx context.Context, name string) (P, error)
}

// NewAtomClient creates a new client for the given atom type
func NewAtomClient[P Proxy](getter func(ctx context.Context, name string) (P, error)) AtomClient[P] {
	return &atomClient[P]{
		getter: getter,
	}
}

type atomClient[P Proxy] struct {
	getter func(ctx context.Context, name string) (P, error)
}

func (c *atomClient[P]) GetProxy(ctx context.Context, name string) (P, error) {
	return c.getter(ctx, name)
}

var _ AtomClient[Proxy] = (*atomClient[Proxy])(nil)
