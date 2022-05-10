// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"
	"path/filepath"
	"sync"
)

func NewCache(opts ...CacheOption) *Cache {
	var options CacheOptions
	options.apply(opts...)
	return &Cache{
		CacheOptions: options,
	}
}

type Cache struct {
	CacheOptions
	plugins map[string]*Plugin
	mu      sync.RWMutex
}

func (c *Cache) Get(name, version string) *Plugin {
	key := getPluginName(name, version)
	c.mu.RLock()
	plugin, ok := c.plugins[key]
	c.mu.RUnlock()
	if ok {
		return plugin
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	plugin, ok = c.plugins[key]
	if ok {
		return plugin
	}

	path := filepath.Join(c.Path, fmt.Sprintf("%s@%s.so", name, version))
	plugin = newPlugin(name, version, path)
	c.plugins[key] = plugin
	return plugin
}
