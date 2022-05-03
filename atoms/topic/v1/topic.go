// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	topicv1 "github.com/atomix/runtime-api/api/atomix/topic/v1"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
	"github.com/atomix/runtime-api/pkg/runtime/driver"
	"google.golang.org/grpc"
)

var Atom = atom.New[TopicProxy](clientFactory, func(server *grpc.Server, service *atom.Service[TopicProxy], registry *atom.Registry[TopicProxy]) {
	topicv1.RegisterTopicManagerServer(server, newTopicV1ManagerServer(service))
	topicv1.RegisterTopicServer(server, newTopicV1Server(registry))
})

// clientFactory is the counter/v1 client factory
var clientFactory = atom.NewClientFactory[TopicProxy](func(client driver.Client) (*atom.Client[TopicProxy], bool) {
	if counterClient, ok := client.(TopicClient); ok {
		return atom.NewClient[TopicProxy](counterClient.GetTopic), true
	}
	return nil, false
})

type TopicClient interface {
	GetTopic(ctx context.Context, name string) (TopicProxy, error)
}

type TopicProxy interface {
	atom.Atom
	topicv1.TopicServer
}
