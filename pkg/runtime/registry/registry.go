// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package registry

import (
	"github.com/atomix/sdk/pkg/logging"
	"io"
	"path/filepath"
	"sync"
)

var log = logging.GetLogger()

func newRegistry(opts ...Option) *Registry {
	var options Options
	options.apply(opts...)
	return &Registry{
		Options: options,
		plugins: make(map[string]*Plugin),
	}
}

type Registry struct {
	Options
	plugins map[string]*Plugin
	mu      sync.Mutex
}

func (r *Registry) Create(name, version, apiVersion string) (io.WriteCloser, error) {
	return r.get(name, version, apiVersion).Create()
}

func (r *Registry) Open(name, version, apiVersion string) (io.ReadCloser, error) {
	return r.get(name, version, apiVersion).Open()
}

func (r *Registry) get(name, version, apiVersion string) *Plugin {
	r.mu.Lock()
	defer r.mu.Unlock()
	path := filepath.Join(r.Path, getPluginFilename(name, version, apiVersion))
	plugin, ok := r.plugins[path]
	if !ok {
		plugin = &Plugin{
			registry:   r,
			Name:       name,
			Version:    version,
			APIVersion: apiVersion,
			Path:       path,
		}
	}
	return plugin
}
