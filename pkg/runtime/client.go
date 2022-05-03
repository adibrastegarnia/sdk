// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"context"
)

type Client interface {
	Closer
}

type Configurator[C any] interface {
	Configure(ctx context.Context, config C) error
}

type Creator interface {
	Create(ctx context.Context) error
}

type Closer interface {
	Close(ctx context.Context) error
}

type Connector[C any] interface {
	Connect(ctx context.Context, config C) (Client, error)
}

func NewConnector[C any](connector func(ctx context.Context, config C) (Client, error)) Connector[C] {
	return &funcConnector[C]{
		connector: connector,
	}
}

type funcConnector[C any] struct {
	connector func(ctx context.Context, config C) (Client, error)
}

func (c *funcConnector[C]) Connect(ctx context.Context, config C) (Client, error) {
	return c.connector(ctx, config)
}

var _ Connector[any] = (*funcConnector[any])(nil)
