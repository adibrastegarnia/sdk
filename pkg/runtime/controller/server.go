// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"context"
	runtimev1 "github.com/atomix/runtime-api/api/atomix/runtime/v1"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/logging"
)

func newServer(controller Controller) *Server {
	return &Server{
		controller: controller,
	}
}

type Server struct {
	controller Controller
}

func (s *Server) GetCluster(ctx context.Context, request *runtimev1.GetClusterRequest) (*runtimev1.GetClusterResponse, error) {
	log.Debugw("GetCluster", logging.Stringer("GetClusterRequest", request))
	cluster, err := s.controller.GetCluster(ctx, request.Name, GetClusterOptions{})
	if err != nil {
		return nil, errors.ToProto(err)
	}

	response := &runtimev1.GetClusterResponse{
		Cluster: cluster,
	}
	log.Debugw("GetCluster", logging.Stringer("GetClusterResponse", response))
	return response, nil
}

func (s *Server) ListClusters(request *runtimev1.ListClustersRequest, server runtimev1.Controller_ListClustersServer) error {
	log.Debugw("ListClusters", logging.Stringer("ListClustersRequest", request))
	options := ListClustersOptions{
		Cluster: request.Name,
		Watch:   request.Watch,
	}
	ch := make(chan runtimev1.Cluster)
	err := s.controller.ListClusters(server.Context(), ch, options)
	if err != nil {
		return errors.ToProto(err)
	}

	for cluster := range ch {
		response := &runtimev1.ListClustersResponse{
			Cluster: cluster,
		}
		log.Debugw("ListClusters", logging.Stringer("ListClustersResponse", response))
		err := server.Send(response)
		if err != nil {
			return err
		}
	}
	return nil
}

var _ runtimev1.ControllerServer = (*Server)(nil)
