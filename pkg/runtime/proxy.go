// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"context"
	"github.com/atomix/runtime-api/pkg/errors"
	"sync"
)

type Proxy interface {
	Closer
}

func newProxyCluster[P Proxy](proxies *ProxyRegistry[P], client *AtomClient[P]) *ProxyCluster[P] {
	return &ProxyCluster[P]{
		proxies: proxies,
		client:  client,
	}
}

type ProxyCluster[P Proxy] struct {
	proxies *ProxyRegistry[P]
	client  *AtomClient[P]
}

func (n *ProxyCluster[P]) CreateProxy(ctx context.Context, name string) error {
	proxy, err := n.client.GetProxy(ctx, name)
	if err != nil {
		return err
	}
	n.proxies.register(name, proxy)
	return nil
}

func (n *ProxyCluster[P]) CloseProxy(ctx context.Context, name string) error {
	proxy, ok := n.proxies.unregister(name)
	if !ok {
		return errors.NewForbidden("proxy '%s' not found", name)
	}
	return proxy.Close(ctx)
}

func NewProxyService[P Proxy](runtime Runtime, primitiveType *AtomType[P], proxies *ProxyRegistry[P]) *ProxyService[P] {
	return &ProxyService[P]{
		runtime:       runtime,
		primitiveType: primitiveType,
		proxies:       proxies,
		clusters:      make(map[string]*ProxyCluster[P]),
	}
}

type ProxyService[P Proxy] struct {
	runtime       Runtime
	primitiveType *AtomType[P]
	proxies       *ProxyRegistry[P]
	clusters      map[string]*ProxyCluster[P]
	mu            sync.RWMutex
}

func (m *ProxyService[P]) GetNamespace(ctx context.Context, name string) (*ProxyCluster[P], error) {
	namespace, ok := m.getNamespace(name)
	if ok {
		return namespace, nil
	}
	return m.newNamespace(ctx, name)
}

func (m *ProxyService[P]) getNamespace(name string) (*ProxyCluster[P], bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	namespace, ok := m.clusters[name]
	return namespace, ok
}

func (m *ProxyService[P]) newNamespace(ctx context.Context, name string) (*ProxyCluster[P], error) {
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

func NewProxyRegistry[P Proxy]() *ProxyRegistry[P] {
	return &ProxyRegistry[P]{
		proxies: make(map[string]P),
	}
}

type ProxyRegistry[P Proxy] struct {
	proxies map[string]P
	mu      sync.RWMutex
}

func (r *ProxyRegistry[P]) GetProxy(name string) (P, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	proxy, ok := r.proxies[name]
	return proxy, ok
}

func (r *ProxyRegistry[P]) register(name string, proxy P) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.proxies[name] = proxy
}

func (r *ProxyRegistry[P]) unregister(name string) (P, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var proxy P
	proxy, ok := r.proxies[name]
	if !ok {
		return proxy, false
	}
	delete(r.proxies, name)
	return proxy, true
}
