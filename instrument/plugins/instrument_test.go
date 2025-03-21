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

package plugins

import (
	"embed"
	"testing"
	
	"github.com/stretchr/testify/require"
	
	"github.com/gaulzhw/skywalking-go-agent/instrument/api"
	"github.com/gaulzhw/skywalking-go-agent/plugins/core/instrument"
)

func TestInstrument_tryToFindThePluginVersion(t *testing.T) {
	tests := []struct {
		name string
		opts *api.CompileOptions
		ins  instrument.Instrument
		want string
	}{
		{
			"normal plugin path",
			&api.CompileOptions{
				AllArgs: []string{
					"github.com/gin-gonic/gin@1.1.1/gin.go",
				},
			},
			NewTestInstrument("github.com/gin-gonic/gin"),
			"1.1.1",
		},
		{
			"plugin with upper-case path",
			&api.CompileOptions{
				AllArgs: []string{
					"github.com/!shopify/sarama@1.34.1/acl.go",
				},
			},
			NewTestInstrument("github.com/Shopify/sarama"),
			"1.34.1",
		},
		{
			"plugin for go stdlib",
			&api.CompileOptions{
				AllArgs: []string{
					"/opt/homebrew/Cellar/go/1.21.4/libexec/src/runtime/metrics/sample.go",
				},
			},
			NewTestInstrument("runtime/metrics"),
			"",
		},
		{
			"plugin for replaced module",
			&api.CompileOptions{
				AllArgs: []string{
					"/home/user/skywalking-go/toolkit/trace/api.go",
				},
			},
			NewTestInstrument("github.com/gaulzhw/skywalking-go-agent/toolkit"),
			"",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Instrument{}
			got, err := i.tryToFindThePluginVersion(tt.opts, tt.ins)
			if err != nil {
				require.NoError(t, err)
			}
			if got != tt.want {
				t.Errorf("tryToFindThePluginVersion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type TestInstrument struct {
	basePackage string
}

func NewTestInstrument(basePackage string) *TestInstrument {
	return &TestInstrument{basePackage: basePackage}
}

func (i *TestInstrument) Name() string {
	return ""
}

func (i *TestInstrument) BasePackage() string {
	return i.basePackage
}

func (i *TestInstrument) VersionChecker(version string) bool {
	return true
}

func (i *TestInstrument) Points() []*instrument.Point {
	return []*instrument.Point{}
}

func (i *TestInstrument) FS() *embed.FS {
	return nil
}
