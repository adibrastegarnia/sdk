// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"context"
	runtimev1 "github.com/atomix/api/pkg/atomix/runtime/v1"
	"github.com/atomix/sdk/pkg/controller"
	"github.com/atomix/sdk/pkg/driver"
	"github.com/atomix/sdk/pkg/logging"
	"sync"
)

var log = logging.GetLogger()

func New(opts ...Option) *Runtime {
	var options Options
	options.apply(opts...)
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
	r.drivers = driver.NewRepository(controller, driver.WithRepoOptions(r.Repository))
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

	watchCh := make(chan *runtimev1.ConnectionInfo)
	if err := r.controller.Connect(context.Background(), name, watchCh); err != nil {
		return nil, err
	}

	select {
	case info, ok := <-watchCh:
		if !ok {
			return nil, context.Canceled
		}

		driver, err := r.drivers.GetDriver(ctx, info.Driver.Name, info.Driver.Version)
		if err != nil {
			return nil, err
		}

		conn, err := driver.Connect(ctx, info.Config)
		if err != nil {
			return nil, err
		}

		go func() {
			for configuration := range watchCh {
				if err := conn.Configure(context.Background(), configuration.Config); err != nil {
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
	if err := r.controller.Close(); err != nil {
		return err
	}
	return nil
}
