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

package v3

import (
	envoy_core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoy_hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"

	"github.com/pkg/errors"
)

import (
	envoy_routes "github.com/apache/dubbo-kubernetes/pkg/xds/envoy/routes"
)

// HttpStaticRouteConfigurer configures a static set of routes into the
// HttpConnectionManager in the filter chain.
type HttpStaticRouteConfigurer struct {
	Builder *envoy_routes.RouteConfigurationBuilder
}

var _ FilterChainConfigurer = &HttpStaticRouteConfigurer{}

func (c *HttpStaticRouteConfigurer) Configure(filterChain *envoy_listener.FilterChain) error {
	routeConfig, err := c.Builder.Build()
	if err != nil {
		return err
	}

	return UpdateHTTPConnectionManager(filterChain, func(hcm *envoy_hcm.HttpConnectionManager) error {
		hcm.RouteSpecifier = &envoy_hcm.HttpConnectionManager_RouteConfig{
			RouteConfig: routeConfig.(*envoy_route.RouteConfiguration),
		}
		return nil
	})
}

// HttpDynamicRouteConfigurer configures the HttpConnectionManager in the
// filter chain to accept its routes dynamically via ADS.
type HttpDynamicRouteConfigurer struct {
	// RouteName is the globally unique name for the RouteConfiguration
	// that this configures xDS client to request.
	RouteName string
}

var _ FilterChainConfigurer = &HttpDynamicRouteConfigurer{}

func (c *HttpDynamicRouteConfigurer) Configure(filterChain *envoy_listener.FilterChain) error {
	return UpdateHTTPConnectionManager(filterChain, func(hcm *envoy_hcm.HttpConnectionManager) error {
		hcm.RouteSpecifier = &envoy_hcm.HttpConnectionManager_Rds{
			Rds: &envoy_hcm.Rds{
				RouteConfigName: c.RouteName,
				ConfigSource: &envoy_core.ConfigSource{
					ResourceApiVersion: envoy_core.ApiVersion_V3,
					ConfigSourceSpecifier: &envoy_core.ConfigSource_Ads{
						Ads: &envoy_core.AggregatedConfigSource{},
					},
				},
			},
		}

		return nil
	})
}

// HttpScopedRouteConfigurer configures a set of scoped routes into the
// HttpConnectionManager in the filter chain.
type HttpScopedRouteConfigurer struct{}

var _ FilterChainConfigurer = &HttpScopedRouteConfigurer{}

func (c *HttpScopedRouteConfigurer) Configure(filterChain *envoy_listener.FilterChain) error {
	return errors.New("scoped routes not implemented")
}
