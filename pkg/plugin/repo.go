// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"plugin"
	"sync"
)

func NewRepo[T any](cache *Cache, opts ...RepoOption) *Repository[T] {
	var options RepoOptions
	options.apply(opts...)
	return &Repository[T]{
		RepoOptions: options,
		cache:       cache,
		plugins:     make(map[string]T),
	}
}

type Repository[T any] struct {
	RepoOptions
	cache   *Cache
	plugins map[string]T
	mu      sync.RWMutex
}

func (r *Repository[T]) Load(name string, version string) (T, error) {
	var t T

	key := getPluginName(name, version)
	r.mu.RLock()
	t, ok := r.plugins[key]
	r.mu.RUnlock()
	if ok {
		return t, nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	t, ok = r.plugins[key]
	if ok {
		return t, nil
	}

	path := r.cache.Get(name, version).Path
	plugin, err := plugin.Open(path)
	if err != nil {
		return t, err
	}

	symbol, err := plugin.Lookup(r.Symbol)
	if err != nil {
		return t, err
	}

	t = symbol.(T)
	r.plugins[key] = t
	return t, nil
}
