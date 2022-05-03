// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package registry

type Options struct {
	Server ServerOptions `yaml:"server,omitempty"`
	Path   string        `yaml:"path"`
}

type ServerOptions struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (o Options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(&o)
	}
}

type Option func(*Options)

func WithOptions(options Options) Option {
	return func(opts *Options) {
		*opts = options
	}
}

func WithHost(host string) Option {
	return func(options *Options) {
		options.Server.Host = host
	}
}

func WithPort(port int) Option {
	return func(options *Options) {
		options.Server.Port = port
	}
}

func WithPath(path string) Option {
	return func(options *Options) {
		options.Path = path
	}
}
