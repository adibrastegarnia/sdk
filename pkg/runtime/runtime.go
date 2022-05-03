// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"context"
	runtimev1 "github.com/atomix/runtime-api/api/atomix/runtime/v1"
	"github.com/atomix/runtime-api/pkg/logging"
	"github.com/atomix/runtime-api/pkg/runtime/controller"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"github.com/atomix/runtime-api/pkg/version"
	"sync"
)

var log = logging.GetLogger()

func New(options Options) *Runtime {
	return &Runtime{
		Options: options,
		conns:   make(map[string]driver.Conn),
	}
}

type Runtime struct {
	Options
	controller *controller.Client
	drivers    *driver.Repository
	conns      map[string]driver.Conn
	mu         sync.RWMutex
}

func (r *Runtime) Run() error {
	controller, err := controller.NewClient(controller.WithOptions(r.Controller))
	if err != nil {
		return err
	}
	r.controller = controller

	drivers, err := driver.NewRepository(driver.WithRepositoryOptions(r.Repository))
	if err != nil {
		return err
	}
	r.drivers = drivers
	return nil
}

func (r *Runtime) Connect(ctx context.Context, cluster string) (driver.Conn, error) {
	conn, ok := r.getConn(cluster)
	if ok {
		return conn, nil
	}
	return r.connect(ctx, cluster)
}

func (r *Runtime) getConn(store string) (driver.Conn, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	conn, ok := r.conns[store]
	return conn, ok
}

func (r *Runtime) connect(ctx context.Context, name string) (driver.Conn, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	conn, ok := r.conns[name]
	if ok {
		return conn, nil
	}

	watchCh := make(chan runtimev1.Cluster)
	if err := r.controller.WatchCluster(context.Background(), name, watchCh); err != nil {
		return nil, err
	}

	select {
	case cluster, ok := <-watchCh:
		if !ok {
			return nil, context.Canceled
		}

		pluginInfo := driver.PluginInfo{
			Name:       cluster.Driver,
			Version:    cluster.Version,
			APIVersion: version.Version(),
		}
		driver, err := r.drivers.Load(ctx, pluginInfo)
		if err != nil {
			return nil, err
		}

		conn, err := driver.Connect(ctx, cluster.Data)
		if err != nil {
			return nil, err
		}

		go func() {
			for configuration := range watchCh {
				if err := conn.Configure(context.Background(), configuration.Data); err != nil {
					log.Error(err)
				}
			}

			r.mu.Lock()
			delete(r.conns, name)
			r.mu.Unlock()

			if err := conn.Close(context.Background()); err != nil {
				log.Error(err)
			}
		}()

		r.conns[name] = conn
		return conn, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (r *Runtime) Shutdown() error {
	if err := r.drivers.Close(); err != nil {
		return err
	}
	if err := r.controller.Close(); err != nil {
		return err
	}
	return nil
}
