// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package driver

type Options struct {
	Repository RepositoryOptions `yaml:"repository,omitempty"`
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

type RepositoryOptions struct {
	Registry RegistryOptions `yaml:"registry"`
	Path     string          `yaml:"path"`
}

type RegistryOptions struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (o RepositoryOptions) apply(opts ...RepositoryOption) {
	for _, opt := range opts {
		opt(&o)
	}
}

type RepositoryOption func(*RepositoryOptions)

func WithRepositoryOptions(options RepositoryOptions) RepositoryOption {
	return func(opts *RepositoryOptions) {
		*opts = options
	}
}

func WithRegistryHost(host string) RepositoryOption {
	return func(options *RepositoryOptions) {
		options.Registry.Host = host
	}
}

func WithRegistryPort(port int) RepositoryOption {
	return func(options *RepositoryOptions) {
		options.Registry.Port = port
	}
}

func WithPath(path string) RepositoryOption {
	return func(options *RepositoryOptions) {
		options.Path = path
	}
}
