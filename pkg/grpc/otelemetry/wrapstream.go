// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package otelemetry

import (
	"context"
	"google.golang.org/grpc"
)

type wrappedServerStream struct {
	grpc.ServerStream
	ctx                    context.Context
	receivedMessageCounter int32
	sentMessageCounter     int32
}

func (w *wrappedServerStream) RecvMsg(m interface{}) error {
	err := w.ServerStream.RecvMsg(m)

	if err == nil {
		w.receivedMessageCounter++
	}

	return err
}

func (w *wrappedServerStream) SendMsg(m interface{}) error {
	err := w.ServerStream.SendMsg(m)
	w.sentMessageCounter++
	return err
}

func wrapServerStream(ctx context.Context, stream grpc.ServerStream) *wrappedServerStream {
	return &wrappedServerStream{
		ServerStream: stream,
		ctx:          ctx,
	}
}
