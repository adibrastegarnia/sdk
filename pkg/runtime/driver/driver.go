// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"context"
	"github.com/atomix/runtime-api/pkg/config"
)

type Driver interface {
	Connect(ctx context.Context, config []byte) (Conn, error)
}

func NewDriver[C config.Config](connector Connector[C], codec config.Codec[C]) Driver {
	return newConfigurableDriver[C](connector, codec)
}

func newConfigurableDriver[C config.Config](connector Connector[C], codec config.Codec[C]) Driver {
	return &configurableDriver[C]{
		connector: connector,
		codec:     codec,
	}
}

type configurableDriver[C config.Config] struct {
	connector Connector[C]
	codec     config.Codec[C]
}

func (d *configurableDriver[C]) Connect(ctx context.Context, bytes []byte) (Conn, error) {
	config, err := d.codec.Decode(bytes)
	if err != nil {
		return nil, err
	}
	conn, err := d.connector.Connect(ctx, config)
	if err != nil {
		return nil, err
	}
	return newConfigurableConn[C](conn, d.codec), nil
}

var _ Driver = (*configurableDriver[config.Config])(nil)
