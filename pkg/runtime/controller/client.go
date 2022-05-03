// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"context"
	"fmt"
	runtimev1 "github.com/atomix/runtime-api/api/atomix/runtime/v1"
	"github.com/atomix/runtime-api/pkg/errors"
	"github.com/atomix/runtime-api/pkg/grpc/retry"
	"github.com/atomix/runtime-api/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
)

var log = logging.GetLogger()

func NewClient(opts ...Option) (*Client, error) {
	var options Options
	options.apply(opts...)
	controller := &Client{
		Options: options,
	}
	if err := controller.connect(); err != nil {
		return nil, err
	}
	return controller, nil
}

type Client struct {
	Options
	conn *grpc.ClientConn
}

func (c *Client) connect() error {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", c.Host, c.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(retry.RetryingUnaryClientInterceptor()),
		grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor()))
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *Client) GetCluster(ctx context.Context, name string) (runtimev1.Cluster, error) {
	client := runtimev1.NewControllerClient(c.conn)
	request := &runtimev1.GetClusterRequest{
		Name: name,
	}
	response, err := client.GetCluster(ctx, request)
	if err != nil {
		return runtimev1.Cluster{}, errors.From(err)
	}
	return response.Cluster, nil
}

func (c *Client) WatchCluster(ctx context.Context, name string, ch chan<- runtimev1.Cluster) error {
	client := runtimev1.NewControllerClient(c.conn)
	request := &runtimev1.ListClustersRequest{
		Name:  name,
		Watch: true,
	}
	stream, err := client.ListClusters(ctx, request)
	if err != nil {
		return errors.From(err)
	}
	go func() {
		defer close(ch)
		for {
			response, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					log.Error(err)
				}
				return
			}
			ch <- response.Cluster
		}
	}()
	return nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
