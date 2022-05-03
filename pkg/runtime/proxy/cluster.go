// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package proxy

import (
	"context"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
)

func newCluster[T atom.Atom](proxies *Registry[T], client *atom.Client[T]) *Cluster[T] {
	return &Cluster[T]{
		proxies: proxies,
		client:  client,
	}
}

type Cluster[T atom.Atom] struct {
	proxies *Registry[T]
	client  *atom.Client[T]
}

func (n *Cluster[T]) CreateProxy(ctx context.Context, name string) error {
	proxy, err := n.client.GetProxy(ctx, name)
	if err != nil {
		return err
	}
	n.proxies.register(name, proxy)
	return nil
}

func (n *Cluster[T]) CloseProxy(ctx context.Context, name string) error {
	proxy, ok := n.proxies.unregister(name)
	if !ok {
		return errors.NewForbidden("proxy '%s' not found", name)
	}
	return proxy.Close(ctx)
}
