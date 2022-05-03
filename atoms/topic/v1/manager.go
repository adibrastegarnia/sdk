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

func newTopicV1ManagerServer(proxies *atom.Service[Topic]) topicv1.TopicManagerServer {
	return &topicV1ManagerServer{
		proxies: proxies,
	}
}

type topicV1ManagerServer struct {
	proxies *atom.Service[Topic]
}

func (s *topicV1ManagerServer) Create(ctx context.Context, request *topicv1.CreateRequest) (*topicv1.CreateResponse, error) {
	namespace, err := s.proxies.GetCluster(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = namespace.CreateProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &topicv1.CreateResponse{}, nil
}

func (s *topicV1ManagerServer) Close(ctx context.Context, request *topicv1.CloseRequest) (*topicv1.CloseResponse, error) {
	namespace, err := s.proxies.GetCluster(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = namespace.CloseProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &topicv1.CloseResponse{}, nil
}

var _ topicv1.TopicManagerServer = (*topicV1ManagerServer)(nil)
