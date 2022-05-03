// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"context"
)

// NewAtomType creates a new type for the given atom type
func NewAtomType[P Proxy](getter func(client Client) (*AtomClient[P], bool)) *AtomType[P] {
	return &AtomType[P]{
		getter: getter,
	}
}

type AtomType[P Proxy] struct {
	getter func(client Client) (*AtomClient[P], bool)
}

func (t *AtomType[P]) Client(client Client) (*AtomClient[P], bool) {
	return t.getter(client)
}

// NewAtomClient creates a new client for the given atom type
func NewAtomClient[P Proxy](getter func(ctx context.Context, name string) (P, error)) *AtomClient[P] {
	return &AtomClient[P]{
		getter: getter,
	}
}

type AtomClient[P Proxy] struct {
	getter func(ctx context.Context, name string) (P, error)
}

func (c *AtomClient[P]) GetProxy(ctx context.Context, name string) (P, error) {
	return c.getter(ctx, name)
}
