// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	topicv1 "github.com/atomix/runtime-api/api/atomix/topic/v1"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/runtime/atom"
)

func newTopicV1Server(proxies *atom.Registry[Topic]) topicv1.TopicServer {
	return &topicV1Server{
		proxies: proxies,
	}
}

type topicV1Server struct {
	proxies *atom.Registry[Topic]
}

func (s *topicV1Server) Publish(ctx context.Context, request *topicv1.PublishRequest) (*topicv1.PublishResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Publish(ctx, request)
}

func (s *topicV1Server) Subscribe(request *topicv1.SubscribeRequest, server topicv1.Topic_SubscribeServer) error {
	proxy, ok := s.proxies.GetProxy(request.Headers.Primitive.Name)
	if !ok {
		return errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Primitive.Name))
	}
	return proxy.Subscribe(request, server)
}

var _ topicv1.TopicServer = (*topicV1Server)(nil)
