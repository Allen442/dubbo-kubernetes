/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package runtime

import (
	dubbo_cp "github.com/apache/dubbo-admin/pkg/config/app/dubbo-cp"
	"github.com/apache/dubbo-admin/pkg/core/cert/provider"
	"github.com/apache/dubbo-admin/pkg/core/runtime/component"
	"github.com/apache/dubbo-admin/pkg/cp-server/server"
	"sync"
	"time"
)

type Runtime interface {
	RuntimeInfo
	RuntimeContext
	component.Manager
}

type RuntimeInfo interface {
	GetInstanceId() string
	SetClusterId(clusterId string)
	GetClusterId() string
	GetStartTime() time.Time
}

type runtimeInfo struct {
	mtx sync.RWMutex

	instanceId string
	clusterId  string
	startTime  time.Time
}

func (i *runtimeInfo) GetInstanceId() string {
	return i.instanceId
}

func (i *runtimeInfo) SetClusterId(clusterId string) {
	i.mtx.Lock()
	defer i.mtx.Unlock()
	i.clusterId = clusterId
}

func (i *runtimeInfo) GetClusterId() string {
	i.mtx.RLock()
	defer i.mtx.RUnlock()
	return i.clusterId
}

func (i *runtimeInfo) GetStartTime() time.Time {
	return i.startTime
}

type RuntimeContext interface {
	Config() *dubbo_cp.Config
	GrpcServer() *server.GrpcServer
	CertStorage() provider.Storage
	KubuClient() provider.Client
}

type runtime struct {
	RuntimeInfo
	RuntimeContext
	component.Manager
}

var _ RuntimeInfo = &runtimeInfo{}

var _ RuntimeContext = &runtimeContext{}

type runtimeContext struct {
	cfg         *dubbo_cp.Config
	grpcServer  *server.GrpcServer
	certStorage provider.Storage
	kubuClient  provider.Client
}

func (rc *runtimeContext) KubuClient() provider.Client {
	return rc.kubuClient
}

func (rc *runtimeContext) CertStorage() provider.Storage {
	return rc.certStorage
}

func (rc *runtimeContext) Config() *dubbo_cp.Config {
	return rc.cfg
}

func (rc *runtimeContext) GrpcServer() *server.GrpcServer {
	return rc.grpcServer
}