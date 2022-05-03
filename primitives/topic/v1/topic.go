// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	topicv1 "github.com/atomix/runtime-api/api/atomix/topic/v1"
	"github.com/atomix/runtime-api/pkg/runtime"
	"google.golang.org/grpc"
)

// Register registers the primitive with the given runtime
func Register(server *grpc.Server, rt runtime.Runtime) {
	proxies := runtime.NewProxyRegistry[TopicProxy]()
	topicv1.RegisterTopicManagerServer(server, newTopicV1ManagerServer(runtime.NewProxyService[TopicProxy](rt, PrimitiveType, proxies)))
	topicv1.RegisterTopicServer(server, newTopicV1Server(proxies))
}

// PrimitiveType is the topic/v1 primitive type
var PrimitiveType = runtime.NewPrimitiveType[TopicProxy](func(client runtime.Client) (runtime.PrimitiveClient[TopicProxy], bool) {
	if topicClient, ok := client.(TopicClient); ok {
		return runtime.NewPrimitiveClient[TopicProxy](topicClient.GetTopic), true
	}
	return nil, false
})

type TopicClient interface {
	GetTopic(ctx context.Context, name string) (TopicProxy, error)
}

type TopicProxy interface {
	runtime.Proxy
	topicv1.TopicServer
}
