// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package proxy

import (
	"context"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"sync"
)

type Client interface {
	Connect(ctx context.Context, name string) (driver.Conn, error)
}

func NewService[T atom.Atom](client Client, primitiveType *atom.Type[T], proxies *Registry[T]) *Service[T] {
	return &Service[T]{
		client:        client,
		primitiveType: primitiveType,
		proxies:       proxies,
		clusters:      make(map[string]*Cluster[T]),
	}
}

type Service[T atom.Atom] struct {
	client        Client
	primitiveType *atom.Type[T]
	proxies       *Registry[T]
	clusters      map[string]*Cluster[T]
	mu            sync.RWMutex
}

func (m *Service[T]) GetCluster(ctx context.Context, name string) (*Cluster[T], error) {
	namespace, ok := m.getCluster(name)
	if ok {
		return namespace, nil
	}
	return m.newCluster(ctx, name)
}

func (m *Service[T]) getCluster(name string) (*Cluster[T], bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	namespace, ok := m.clusters[name]
	return namespace, ok
}

func (m *Service[T]) newCluster(ctx context.Context, name string) (*Cluster[T], error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	namespace, ok := m.clusters[name]
	if ok {
		return namespace, nil
	}

	conn, err := m.client.Connect(ctx, name)
	if err != nil {
		return nil, err
	}

	client, ok := m.primitiveType.Client(conn.Client())
	if !ok {
		return nil, errors.NewNotSupported("primitive type not supported by client for store '%s'", namespace)
	}

	namespace = newCluster(m.proxies, client)
	m.clusters[name] = namespace
	return namespace, nil
}
