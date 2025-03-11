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
	traceactivation "github.com/gaulzhw/skywalking-go-agent/plugins/toolkit-activation"
	
	"github.com/gaulzhw/skywalking-go-agent/plugins/amqp"
	"github.com/gaulzhw/skywalking-go-agent/plugins/core/instrument"
	"github.com/gaulzhw/skywalking-go-agent/plugins/dubbo"
	"github.com/gaulzhw/skywalking-go-agent/plugins/echov4"
	fasthttp_client "github.com/gaulzhw/skywalking-go-agent/plugins/fasthttp/hostclient"
	fasthttp_router "github.com/gaulzhw/skywalking-go-agent/plugins/fasthttp/router"
	"github.com/gaulzhw/skywalking-go-agent/plugins/fiber"
	"github.com/gaulzhw/skywalking-go-agent/plugins/gin"
	goelasticsearchv8 "github.com/gaulzhw/skywalking-go-agent/plugins/go-elasticsearchv8"
	goredisv9 "github.com/gaulzhw/skywalking-go-agent/plugins/go-redisv9"
	"github.com/gaulzhw/skywalking-go-agent/plugins/go-restfulv3"
	"github.com/gaulzhw/skywalking-go-agent/plugins/goframe"
	gorm_entry "github.com/gaulzhw/skywalking-go-agent/plugins/gorm/entry"
	gorm_mysql "github.com/gaulzhw/skywalking-go-agent/plugins/gorm/mysql"
	"github.com/gaulzhw/skywalking-go-agent/plugins/grpc"
	"github.com/gaulzhw/skywalking-go-agent/plugins/http"
	"github.com/gaulzhw/skywalking-go-agent/plugins/irisv12"
	"github.com/gaulzhw/skywalking-go-agent/plugins/kratosv2"
	"github.com/gaulzhw/skywalking-go-agent/plugins/microv4"
	"github.com/gaulzhw/skywalking-go-agent/plugins/mongo"
	"github.com/gaulzhw/skywalking-go-agent/plugins/mux"
	"github.com/gaulzhw/skywalking-go-agent/plugins/pulsar"
	"github.com/gaulzhw/skywalking-go-agent/plugins/rocketmq"
	runtime_metrics "github.com/gaulzhw/skywalking-go-agent/plugins/runtimemetrics"
	segmentiokafka "github.com/gaulzhw/skywalking-go-agent/plugins/segmentio-kafka"
	sql_entry "github.com/gaulzhw/skywalking-go-agent/plugins/sql/entry"
	sql_mysql "github.com/gaulzhw/skywalking-go-agent/plugins/sql/mysql"
)

var instruments = make([]instrument.Instrument, 0)

func init() {
	// register the plugins instrument
	registerFramework(gin.NewInstrument())
	registerFramework(http.NewInstrument())
	registerFramework(dubbo.NewInstrument())
	registerFramework(restfulv3.NewInstrument())
	registerFramework(kratosv2.NewInstrument())
	registerFramework(microv4.NewInstrument())
	registerFramework(goredisv9.NewInstrument())
	registerFramework(mongo.NewInstrument())
	registerFramework(runtime_metrics.NewInstrument())
	registerFramework(mux.NewInstrument())
	registerFramework(grpc.NewInstrument())
	registerFramework(irisv12.NewInstrument())
	registerFramework(traceactivation.NewInstrument())
	registerFramework(fiber.NewInstrument())
	registerFramework(rocketmq.NewInstrument())
	registerFramework(amqp.NewInstrument())
	registerFramework(pulsar.NewInstrument())
	registerFramework(segmentiokafka.NewInstrument())
	registerFramework(goelasticsearchv8.NewInstrument())
	
	// fasthttp related instruments
	registerFramework(fasthttp_client.NewInstrument())
	registerFramework(fasthttp_router.NewInstrument())
	
	// gorm related instruments
	registerFramework(gorm_entry.NewInstrument())
	registerFramework(gorm_mysql.NewInstrument())
	
	// sql related instruments
	registerFramework(sql_entry.NewInstrument())
	registerFramework(sql_mysql.NewInstrument())
	
	// echov4 related instruments
	registerFramework(echov4.NewInstrument())
	
	// goframe
	registerFramework(goframe.NewInstrument())
}

func registerFramework(ins instrument.Instrument) {
	instruments = append(instruments, ins)
}
