// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	"github.com/atomix/api/pkg/atomix/topic/v1"
	"github.com/atomix/sdk/pkg/atom"
	"github.com/atomix/sdk/pkg/errors"
)

func newTopicV1Server(proxies *atom.Registry[Topic]) v1.TopicServer {
	return &topicV1Server{
		proxies: proxies,
	}
}

type topicV1Server struct {
	proxies *atom.Registry[Topic]
}

func (s *topicV1Server) Publish(ctx context.Context, request *v1.PublishRequest) (*v1.PublishResponse, error) {
	proxy, ok := s.proxies.GetProxy(request.Headers.Atom)
	if !ok {
		return nil, errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Atom))
	}
	return proxy.Publish(ctx, request)
}

func (s *topicV1Server) Subscribe(request *v1.SubscribeRequest, server v1.Topic_SubscribeServer) error {
	proxy, ok := s.proxies.GetProxy(request.Headers.Atom)
	if !ok {
		return errors.ToProto(errors.NewForbidden("proxy '%s' not open", request.Headers.Atom))
	}
	return proxy.Subscribe(request, server)
}

var _ v1.TopicServer = (*topicV1Server)(nil)
