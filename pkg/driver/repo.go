// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"context"
	"fmt"
	"github.com/atomix/sdk/pkg/controller"
	"os"
	"path/filepath"
	"plugin"
	"sync"
)

const driverSymName = "Driver"

func NewRepository(controller *controller.Client, opts ...RepoOption) *Repository {
	var options RepoOptions
	options.apply(opts...)
	return &Repository{
		RepoOptions: options,
		controller:  controller,
		drivers:     make(map[string]Driver),
	}
}

// Repository is a driver registry
type Repository struct {
	RepoOptions
	controller *controller.Client
	drivers    map[string]Driver
	mu         sync.RWMutex
}

func (r *Repository) file(name, version string) string {
	return fmt.Sprintf("%s@%s.so", name, version)
}

func (r *Repository) path(name, version string) string {
	return filepath.Join(r.Path, r.file(name, version))
}

func (r *Repository) GetDriver(ctx context.Context, name string, version string) (Driver, error) {
	driver, ok := r.getDriver(name, version)
	if ok {
		return driver, nil
	}
	return r.fetchDriver(ctx, name, version)
}

func (r *Repository) getDriver(name, version string) (Driver, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	driver, ok := r.drivers[r.file(name, version)]
	return driver, ok
}

func (r *Repository) fetchDriver(ctx context.Context, name string, version string) (Driver, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	path, err := r.downloadDriver(ctx, name, version)
	if err != nil {
		return nil, err
	}

	plugin, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	symbol, err := plugin.Lookup(driverSymName)
	if err != nil {
		return nil, err
	}
	return symbol.(Driver), nil
}

func (r *Repository) downloadDriver(ctx context.Context, name string, version string) (string, error) {
	path := r.path(name, version)
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return path, err
	}

	file, err := os.Create(path)
	if err != nil {
		if os.IsExist(err) {
			file, err = os.Open(path)
			if err != nil {
				return path, err
			}
		}
		return path, err
	}
	defer file.Close()
	return path, r.controller.GetDriver(ctx, name, version, file)
}
