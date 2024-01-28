package v3

import (
	util_proto "github.com/apache/dubbo-kubernetes/pkg/util/proto"
	util_xds "github.com/apache/dubbo-kubernetes/pkg/util/xds"
	envoy_listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
)

type HttpConnectionManagerConfigurer struct {
	StatsName                string
	ForwardClientCertDetails bool
	NormalizePath            bool
}

func (c *HttpConnectionManagerConfigurer) Configure(filterChain *envoy_listener.FilterChain) error {
	config := &envoy_hcm.HttpConnectionManager{
		StatPrefix:  util_xds.SanitizeMetric(c.StatsName),
		CodecType:   envoy_hcm.HttpConnectionManager_AUTO,
		HttpFilters: []*envoy_hcm.HttpFilter{},
		// notice that route configuration is left up to other configurers
	}

	if c.ForwardClientCertDetails {
		config.ForwardClientCertDetails = envoy_hcm.HttpConnectionManager_SANITIZE_SET
		config.SetCurrentClientCertDetails = &envoy_hcm.HttpConnectionManager_SetCurrentClientCertDetails{
			Uri: true,
		}
	}

	if c.NormalizePath {
		config.NormalizePath = util_proto.Bool(true)
	}

	pbst, err := util_proto.MarshalAnyDeterministic(config)
	if err != nil {
		return err
	}

	filterChain.Filters = append(filterChain.Filters, &envoy_listener.Filter{
		Name: "envoy.filters.network.http_connection_manager",
		ConfigType: &envoy_listener.Filter_TypedConfig{
			TypedConfig: pbst,
		},
	})
	return nil
}
