// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"context"
	"plugin"
	"sync"
)

// Plugin is a runtime plugin
type Plugin interface {
	Name() string
	Version() string
}

type Loader[T Plugin] interface {
	Load(*plugin.Plugin) (T, error)
}

func NewRegistry[T Plugin](repo Repository, loader Loader[T]) *Registry[T] {
	return &Registry[T]{
		repo:    repo,
		loader:  loader,
		plugins: make(map[string]map[string]T),
	}
}

// Registry is a plugin registry
type Registry[T Plugin] struct {
	repo    Repository
	loader  Loader[T]
	plugins map[string]map[string]T
	mu      sync.RWMutex
}

func (r *Registry[T]) Get(ctx context.Context, name, version string) (Plugin, error) {
	plugin, ok := r.get(name, version)
	if ok {
		return plugin, nil
	}
	return r.lookup(ctx, name, version)
}

func (r *Registry[T]) get(name, version string) (Plugin, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	versions, ok := r.plugins[name]
	if !ok {
		return nil, false
	}
	plugin, ok := versions[version]
	if !ok {
		return nil, false
	}
	return plugin, true
}

func (r *Registry[T]) lookup(ctx context.Context, name, version string) (Plugin, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	versions, ok := r.plugins[name]
	if ok {
		plugin, ok := versions[version]
		if ok {
			return plugin, nil
		}
	}

	goPlugin, err := r.repo.Lookup(ctx, name, version)
	if err != nil {
		return nil, err
	}

	plugin, err := r.loader.Load(goPlugin)
	if err != nil {
		return nil, err
	}

	versions, ok = r.plugins[name]
	if !ok {
		versions = make(map[string]T)
		r.plugins[name] = versions
	}
	versions[version] = plugin
	return plugin, nil
}

// Repository is a plugin repository
type Repository interface {
	Lookup(ctx context.Context, name string, version string) (*plugin.Plugin, error)
}
