// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"github.com/atomix/runtime-api/pkg/runtime/controller"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"google.golang.org/grpc"
)

type Options struct {
	Server     ServerOptions            `yaml:"server"`
	Controller controller.Options       `yaml:"controller"`
	Repository driver.RepositoryOptions `yaml:"repository"`
	Atoms      []AtomFunc
}

type AtomFunc func(*grpc.Server, *Runtime)

func (o Options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(&o)
	}
}

type Option func(*Options)

type ServerOptions struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func WithOptions(opts Options) Option {
	return func(options *Options) {
		*options = opts
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

func WithAtom(atom AtomFunc) Option {
	return func(options *Options) {
		options.Atoms = append(options.Atoms, atom)
	}
}

func WithAtoms(atom ...AtomFunc) Option {
	return func(options *Options) {
		options.Atoms = append(options.Atoms, atom...)
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
