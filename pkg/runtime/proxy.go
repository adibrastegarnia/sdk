// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"context"
	"github.com/atomix/runtime-api/pkg/errors"
	"sync"
)

func newProxyCluster[T Atom](proxies *ProxyRegistry[T], client *AtomClient[T]) *ProxyCluster[T] {
	return &ProxyCluster[T]{
		proxies: proxies,
		client:  client,
	}
}

type ProxyCluster[T Atom] struct {
	proxies *ProxyRegistry[T]
	client  *AtomClient[T]
}

func (n *ProxyCluster[T]) CreateProxy(ctx context.Context, name string) error {
	proxy, err := n.client.GetProxy(ctx, name)
	if err != nil {
		return err
	}
	n.proxies.register(name, proxy)
	return nil
}

func (n *ProxyCluster[T]) CloseProxy(ctx context.Context, name string) error {
	proxy, ok := n.proxies.unregister(name)
	if !ok {
		return errors.NewForbidden("proxy '%s' not found", name)
	}
	return proxy.Close(ctx)
}

func NewProxyService[T Atom](runtime *Runtime, primitiveType *AtomType[T], proxies *ProxyRegistry[T]) *ProxyService[T] {
	return &ProxyService[T]{
		runtime:       runtime,
		primitiveType: primitiveType,
		proxies:       proxies,
		clusters:      make(map[string]*ProxyCluster[T]),
	}
}

type ProxyService[T Atom] struct {
	runtime       *Runtime
	primitiveType *AtomType[T]
	proxies       *ProxyRegistry[T]
	clusters      map[string]*ProxyCluster[T]
	mu            sync.RWMutex
}

func (m *ProxyService[T]) GetCluster(ctx context.Context, name string) (*ProxyCluster[T], error) {
	namespace, ok := m.getCluster(name)
	if ok {
		return namespace, nil
	}
	return m.newCluster(ctx, name)
}

func (m *ProxyService[T]) getCluster(name string) (*ProxyCluster[T], bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	namespace, ok := m.clusters[name]
	return namespace, ok
}

func (m *ProxyService[T]) newCluster(ctx context.Context, name string) (*ProxyCluster[T], error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	namespace, ok := m.clusters[name]
	if ok {
		return namespace, nil
	}

	conn, err := m.runtime.Connect(ctx, name)
	if err != nil {
		return nil, err
	}

	client, ok := m.primitiveType.Client(conn.Client())
	if !ok {
		return nil, errors.NewNotSupported("primitive type not supported by client for store '%s'", namespace)
	}

	namespace = newProxyCluster(m.proxies, client)
	m.clusters[name] = namespace
	return namespace, nil
}

func NewProxyRegistry[T Atom]() *ProxyRegistry[T] {
	return &ProxyRegistry[T]{
		proxies: make(map[string]T),
	}
}

type ProxyRegistry[T Atom] struct {
	proxies map[string]T
	mu      sync.RWMutex
}

func (r *ProxyRegistry[T]) GetProxy(name string) (T, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	proxy, ok := r.proxies[name]
	return proxy, ok
}

func (r *ProxyRegistry[T]) register(name string, proxy T) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.proxies[name] = proxy
}

func (r *ProxyRegistry[T]) unregister(name string) (T, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var proxy T
	proxy, ok := r.proxies[name]
	if !ok {
		return proxy, false
	}
	delete(r.proxies, name)
	return proxy, true
}
