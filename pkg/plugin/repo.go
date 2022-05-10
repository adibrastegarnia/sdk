// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"context"
	"github.com/atomix/sdk/pkg/errors"
	"io"
	"os"
	"plugin"
	"sync"
)

func NewRepository[T any](cache *Cache, opts ...RepoOption) Repository[T] {
	var options RepoOptions
	options.apply(opts...)
	return &cacheRepository[T]{
		options: options,
		cache:   cache,
		plugins: make(map[string]T),
	}
}

type DownloadFunc func(ctx context.Context, name string, version string, writer io.Writer) error

type Repository[T any] interface {
	Load(ctx context.Context, name string, version string) (T, error)
}

type cacheRepository[T any] struct {
	options RepoOptions
	cache   *Cache
	plugins map[string]T
	mu      sync.RWMutex
}

func (r *cacheRepository[T]) Load(ctx context.Context, name string, version string) (T, error) {
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

	path, err := r.check(ctx, name, version)
	if err != nil {
		return t, err
	}

	plugin, err := plugin.Open(path)
	if err != nil {
		return t, err
	}

	symbol, err := plugin.Lookup(r.options.Symbol)
	if err != nil {
		return t, err
	}

	t = symbol.(T)
	r.plugins[key] = t
	return t, nil
}

func (r *cacheRepository[T]) check(ctx context.Context, name string, version string) (string, error) {
	plugin := r.cache.Get(name, version)
	if _, err := os.Stat(plugin.Path); err == nil {
		return plugin.Path, nil
	}
	if err := r.download(ctx, plugin); err != nil {
		return "", err
	}
	return plugin.Path, nil
}

func (r *cacheRepository[T]) download(ctx context.Context, plugin *Plugin) error {
	downloader := r.options.Downloader
	if downloader == nil {
		return errors.NewNotSupported("plugin repository does not support downloads")
	}
	writer, err := plugin.Create()
	if err != nil {
		return err
	}
	err = r.options.Downloader(ctx, plugin.Name, plugin.Version, writer)
	if err != nil {
		return err
	}
	return nil
}
