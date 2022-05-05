// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	"github.com/atomix/sdk/pkg/errors"
	"github.com/atomix/sdk/pkg/runtime/atom"
)

func newTopicV1ManagerServer(proxies *atom.Service[Topic]) v1.TopicManagerServer {
	return &topicV1ManagerServer{
		proxies: proxies,
	}
}

type topicV1ManagerServer struct {
	proxies *atom.Service[Topic]
}

func (s *topicV1ManagerServer) Create(ctx context.Context, request *v1.CreateRequest) (*v1.CreateResponse, error) {
	namespace, err := s.proxies.GetCluster(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = namespace.CreateProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &v1.CreateResponse{}, nil
}

func (s *topicV1ManagerServer) Close(ctx context.Context, request *v1.CloseRequest) (*v1.CloseResponse, error) {
	namespace, err := s.proxies.GetCluster(ctx, request.Cluster.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = namespace.CloseProxy(ctx, request.Headers.Primitive.Name)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &v1.CloseResponse{}, nil
}

var _ v1.TopicManagerServer = (*topicV1ManagerServer)(nil)
