// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package client

import (
	"context"
	"fmt"
	
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/metadata"
	
	"github.com/gaulzhw/skywalking-go-agent/plugins/core/tracing"
)

//skywalking:public
func NewClientWrapper(cli client.Client) client.Client {
	return &clientWrapper{cli}
}

type clientWrapper struct {
	client.Client
}

// Call is used for client calls
func (s *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	span, err := tracing.CreateExitSpan(fmt.Sprintf("%s.%s", req.Service(), req.Endpoint()), req.Service(), func(k, v string) error {
		mda, _ := metadata.FromContext(ctx)
		md := metadata.Copy(mda)
		md[k] = v
		ctx = metadata.NewContext(ctx, md)
		return nil
	}, tracing.WithComponent(5008),
		tracing.WithLayer(tracing.SpanLayerRPCFramework))
	if err != nil {
		return err
	}
	
	defer span.End()
	if err = s.Client.Call(ctx, req, rsp, opts...); err != nil {
		span.Error(err.Error())
	}
	return err
}

// Stream is used streaming
func (s *clientWrapper) Stream(ctx context.Context, req client.Request, opts ...client.CallOption) (client.Stream, error) {
	span, err := tracing.CreateExitSpan(fmt.Sprintf("%s.%s", req.Service(), req.Endpoint()), req.Service(), func(k, v string) error {
		mda, _ := metadata.FromContext(ctx)
		md := metadata.Copy(mda)
		md[k] = v
		ctx = metadata.NewContext(ctx, md)
		return nil
	}, tracing.WithComponent(5008),
		tracing.WithLayer(tracing.SpanLayerRPCFramework))
	if err != nil {
		return nil, err
	}
	
	defer span.End()
	stream, err := s.Client.Stream(ctx, req, opts...)
	if err != nil {
		span.Error(err.Error())
	}
	return stream, err
}

// Publish is used publish message to subscriber
func (s *clientWrapper) Publish(ctx context.Context, p client.Message, opts ...client.PublishOption) error {
	span, err := tracing.CreateExitSpan(fmt.Sprintf("Pub to %s", p.Topic()), p.ContentType(), func(k, v string) error {
		mda, _ := metadata.FromContext(ctx)
		md := metadata.Copy(mda)
		md[k] = v
		ctx = metadata.NewContext(ctx, md)
		return nil
	}, tracing.WithComponent(5008),
		tracing.WithLayer(tracing.SpanLayerRPCFramework))
	if err != nil {
		return err
	}
	
	defer span.End()
	if err = s.Client.Publish(ctx, p, opts...); err != nil {
		span.Error(err.Error())
	}
	return err
}
