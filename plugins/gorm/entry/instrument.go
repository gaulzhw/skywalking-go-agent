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

package entry

import (
	"embed"
	
	"github.com/gaulzhw/skywalking-go-agent/plugins/core/instrument"
)

//go:embed *
var fs embed.FS

//skywalking:nocopy
type Instrument struct {
}

func NewInstrument() *Instrument {
	return &Instrument{}
}

func (i *Instrument) Name() string {
	return "gorm"
}

func (i *Instrument) BasePackage() string {
	return "gorm.io/gorm"
}

func (i *Instrument) VersionChecker(version string) bool {
	return true
}

func (i *Instrument) Points() []*instrument.Point {
	return []*instrument.Point{
		{
			PackagePath: "",
			At: instrument.NewStaticMethodEnhance("Open",
				instrument.WithArgsCount(2),
				instrument.WithArgType(0, "Dialector"), instrument.WithArgType(1, "...Option"),
				instrument.WithResultCount(2),
				instrument.WithResultType(0, "*DB"), instrument.WithResultType(1, "error")),
			Interceptor: "OpenInterceptor",
		},
	}
}

func (i *Instrument) PluginSourceCodePath() string {
	return "entry"
}

func (i *Instrument) FS() *embed.FS {
	return &fs
}
