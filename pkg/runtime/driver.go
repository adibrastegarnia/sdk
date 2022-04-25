// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"context"
	"github.com/atomix/atomix-runtime/pkg/config"
)

type Driver[C config.Config] interface {
	Connect(ctx context.Context, config C) (Conn, error)
}

func NewDriver[C config.Config](connector Connector[C]) Driver[C] {
	return &connectorDriver[C]{
		connector: connector,
	}
}

type connectorDriver[C config.Config] struct {
	connector Connector[C]
}

func (d *connectorDriver[C]) Connect(ctx context.Context, config C) (Conn, error) {
	return d.connector(ctx, config)
}

var _ Driver[config.Config] = (*connectorDriver[config.Config])(nil)

type Conn interface {
	Closer
}

type Closer interface {
	Close(ctx context.Context) error
}

type Connector[C any] func(ctx context.Context, config C) (Conn, error)
