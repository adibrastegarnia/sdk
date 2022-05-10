// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package runtime

type Options struct {
	Host  string
	Port  int
	Atoms []AtomOptions
}

func (o Options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(&o)
	}
}

type Option func(*Options)

type AtomOptions struct {
	Name    string
	Version string
}

func WithOptions(opts Options) Option {
	return func(options *Options) {
		*options = opts
	}
}

func WithHost(host string) Option {
	return func(options *Options) {
		options.Host = host
	}
}

func WithPort(port int) Option {
	return func(options *Options) {
		options.Port = port
	}
}

func WithAtom(name, version string) Option {
	return func(options *Options) {
		options.Atoms = append(options.Atoms, AtomOptions{
			Name:    name,
			Version: version,
		})
	}
}
