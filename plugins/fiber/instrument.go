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

package fiber

import (
	"embed"
	"strings"
	
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
	return "fiber"
}

func (i *Instrument) BasePackage() string {
	return "github.com/gofiber/fiber/v2"
}

func (i *Instrument) VersionChecker(version string) bool {
	return strings.HasPrefix(version, "v2.")
}

func (i *Instrument) Points() []*instrument.Point {
	return []*instrument.Point{
		{
			PackageName: "fiber",
			PackagePath: "",
			At: instrument.NewMethodEnhance("*App", "handler",
				instrument.WithArgsCount(1),
				instrument.WithArgType(0, "*fasthttp.RequestCtx")),
			Interceptor: "HTTPInterceptor",
		},
	}
}

func (i *Instrument) FS() *embed.FS {
	return &fs
}
