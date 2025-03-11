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

package traceactivation

import (
	"github.com/gaulzhw/skywalking-go-agent/plugins/core/operator"
	"github.com/gaulzhw/skywalking-go-agent/plugins/core/tracing"
)

type AsyncLogInterceptor struct {
}

func (h *AsyncLogInterceptor) BeforeInvoke(invocation operator.Invocation) error {
	return nil
}

func (h *AsyncLogInterceptor) AfterInvoke(invocation operator.Invocation, result ...interface{}) error {
	enhancced, ok := invocation.CallerInstance().(operator.EnhancedInstance)
	if !ok {
		return nil
	}
	s := enhancced.GetSkyWalkingDynamicField().(tracing.Span)
	logs := invocation.Args()[0].([]string)
	s.Log(logs...)
	enhancced.SetSkyWalkingDynamicField(s)
	return nil
}
