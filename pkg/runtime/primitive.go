// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"context"
)

// PrimitiveType identifies and operates on a specific type of primitive
type PrimitiveType[P Proxy] interface {
	Client(conn Conn) (PrimitiveClient[P], bool)
}

// NewPrimitiveType creates a new type for the given primitive type
func NewPrimitiveType[P Proxy](getter func(conn Conn) (PrimitiveClient[P], bool)) PrimitiveType[P] {
	return &primitiveType[P]{
		getter: getter,
	}
}

type primitiveType[P Proxy] struct {
	getter func(conn Conn) (PrimitiveClient[P], bool)
}

func (t *primitiveType[P]) Client(conn Conn) (PrimitiveClient[P], bool) {
	return t.getter(conn)
}

var _ PrimitiveType[Proxy] = (*primitiveType[Proxy])(nil)

// PrimitiveClient is a client for a specific type of primitive
type PrimitiveClient[P Proxy] interface {
	GetProxy(ctx context.Context, name string) (P, error)
}

// NewPrimitiveClient creates a new client for the given primitive type
func NewPrimitiveClient[P Proxy](getter func(ctx context.Context, name string) (P, error)) PrimitiveClient[P] {
	return &primitiveClient[P]{
		getter: getter,
	}
}

type primitiveClient[P Proxy] struct {
	getter func(ctx context.Context, name string) (P, error)
}

func (c *primitiveClient[P]) GetProxy(ctx context.Context, name string) (P, error) {
	return c.getter(ctx, name)
}

var _ PrimitiveClient[Proxy] = (*primitiveClient[Proxy])(nil)
