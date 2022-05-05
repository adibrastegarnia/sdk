// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"github.com/atomix/sdk/pkg/runtime/atom"
	"github.com/atomix/sdk/pkg/runtime/controller"
	"github.com/atomix/sdk/pkg/runtime/driver"
	"google.golang.org/grpc"
)

type Options struct {
	Controller controller.Options       `yaml:"controller"`
	Repository driver.RepositoryOptions `yaml:"repository"`
}

func (o Options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(&o)
	}
}

type Option func(*Options)

type ServiceOptions struct {
	Server ServerOptions `yaml:"server"`
	Atoms  []AtomFunc
}

func (o ServiceOptions) apply(opts ...ServiceOption) {
	for _, opt := range opts {
		opt(&o)
	}
}

type ServiceOption func(*ServiceOptions)

type AtomFunc func(*grpc.Server, *Runtime)

type ServerOptions struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func WithOptions(opts Options) Option {
	return func(options *Options) {
		*options = opts
	}
}

func WithServiceOptions(opts ServiceOptions) ServiceOption {
	return func(options *ServiceOptions) {
		*options = opts
	}
}

func WithHost(host string) ServiceOption {
	return func(options *ServiceOptions) {
		options.Server.Host = host
	}
}

func WithPort(port int) ServiceOption {
	return func(options *ServiceOptions) {
		options.Server.Port = port
	}
}

func WithAtom[T atom.Atom](atom *atom.Type[T]) ServiceOption {
	return func(options *ServiceOptions) {
		options.Atoms = append(options.Atoms, func(server *grpc.Server, runtime *Runtime) {
			atom.Register(server, runtime)
		})
	}
}

func WithControllerHost(host string) Option {
	return func(options *Options) {
		options.Controller.Host = host
	}
}

func WithControllerPort(port int) Option {
	return func(options *Options) {
		options.Controller.Port = port
	}
}

func WithRegistryHost(host string) Option {
	return func(options *Options) {
		options.Repository.Registry.Host = host
	}
}

func WithRegistryPort(port int) Option {
	return func(options *Options) {
		options.Repository.Registry.Port = port
	}
}

func WithPlugins(path string) Option {
	return func(options *Options) {
		options.Repository.Path = path
	}
}
