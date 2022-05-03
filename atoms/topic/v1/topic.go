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

var Atom = atom.New[Topic](clientFactory, func(server *grpc.Server, service *atom.Service[Topic], registry *atom.Registry[Topic]) {
	topicv1.RegisterTopicManagerServer(server, newTopicV1ManagerServer(service))
	topicv1.RegisterTopicServer(server, newTopicV1Server(registry))
})

// clientFactory is the topic/v1 client factory
var clientFactory = atom.NewClientFactory[Topic](func(client driver.Client) (*atom.Client[Topic], bool) {
	if topicClient, ok := client.(TopicClient); ok {
		return atom.NewClient[Topic](topicClient.GetTopic), true
	}
	return nil, false
})

type TopicClient interface {
	GetTopic(ctx context.Context, name string) (Topic, error)
}

type Topic interface {
	atom.Atom
	topicv1.TopicServer
}
