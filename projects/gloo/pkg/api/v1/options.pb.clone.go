// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_duration "github.com/golang/protobuf/ptypes/duration"

	github_com_golang_protobuf_ptypes_struct "github.com/golang/protobuf/ptypes/struct"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_filter_http_gzip_v2 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/filter/http/gzip/v2"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_buffer_v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/filters/http/buffer/v3"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_csrf_v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/filters/http/csrf/v3"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_proxylatency "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/proxylatency"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_type_matcher_v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/matcher/v3"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_dlp "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/dlp"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_jwt "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/jwt"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ratelimit"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_rbac "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/rbac"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_waf "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/waf"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_als "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/als"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_aws "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/aws"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_azure "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/azure"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_cors "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/cors"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_dynamic_forward_proxy "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/dynamic_forward_proxy"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_faultinjection "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/faultinjection"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_grpc "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/grpc"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_grpc_json "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/grpc_json"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_grpc_web "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/grpc_web"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_hcm "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/hcm"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_headers "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/headers"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_healthcheck "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/healthcheck"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_lbhash "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/lbhash"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/protocol_upgrade"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_proxy_protocol "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/proxy_protocol"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_rest "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/rest"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_retries "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/retries"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_shadowing "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/shadowing"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_stats "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/stats"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_tcp "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/tcp"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_tracing "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/tracing"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_wasm "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/wasm"

	github_com_solo_io_solo_kit_pkg_api_external_envoy_api_v2_core "github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *ListenerOptions) Clone() proto.Message {
	var target *ListenerOptions
	if m == nil {
		return target
	}
	target = &ListenerOptions{}

	if h, ok := interface{}(m.GetAccessLoggingService()).(clone.Cloner); ok {
		target.AccessLoggingService = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_als.AccessLoggingService)
	} else {
		target.AccessLoggingService = proto.Clone(m.GetAccessLoggingService()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_als.AccessLoggingService)
	}

	if h, ok := interface{}(m.GetExtensions()).(clone.Cloner); ok {
		target.Extensions = h.Clone().(*Extensions)
	} else {
		target.Extensions = proto.Clone(m.GetExtensions()).(*Extensions)
	}

	if h, ok := interface{}(m.GetPerConnectionBufferLimitBytes()).(clone.Cloner); ok {
		target.PerConnectionBufferLimitBytes = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.PerConnectionBufferLimitBytes = proto.Clone(m.GetPerConnectionBufferLimitBytes()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if m.GetSocketOptions() != nil {
		target.SocketOptions = make([]*github_com_solo_io_solo_kit_pkg_api_external_envoy_api_v2_core.SocketOption, len(m.GetSocketOptions()))
		for idx, v := range m.GetSocketOptions() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SocketOptions[idx] = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_external_envoy_api_v2_core.SocketOption)
			} else {
				target.SocketOptions[idx] = proto.Clone(v).(*github_com_solo_io_solo_kit_pkg_api_external_envoy_api_v2_core.SocketOption)
			}

		}
	}

	if h, ok := interface{}(m.GetProxyProtocol()).(clone.Cloner); ok {
		target.ProxyProtocol = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_proxy_protocol.ProxyProtocol)
	} else {
		target.ProxyProtocol = proto.Clone(m.GetProxyProtocol()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_proxy_protocol.ProxyProtocol)
	}

	return target
}

// Clone function
func (m *RouteConfigurationOptions) Clone() proto.Message {
	var target *RouteConfigurationOptions
	if m == nil {
		return target
	}
	target = &RouteConfigurationOptions{}

	if h, ok := interface{}(m.GetMaxDirectResponseBodySizeBytes()).(clone.Cloner); ok {
		target.MaxDirectResponseBodySizeBytes = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.MaxDirectResponseBodySizeBytes = proto.Clone(m.GetMaxDirectResponseBodySizeBytes()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	return target
}

// Clone function
func (m *HttpListenerOptions) Clone() proto.Message {
	var target *HttpListenerOptions
	if m == nil {
		return target
	}
	target = &HttpListenerOptions{}

	if h, ok := interface{}(m.GetGrpcWeb()).(clone.Cloner); ok {
		target.GrpcWeb = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_grpc_web.GrpcWeb)
	} else {
		target.GrpcWeb = proto.Clone(m.GetGrpcWeb()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_grpc_web.GrpcWeb)
	}

	if h, ok := interface{}(m.GetHttpConnectionManagerSettings()).(clone.Cloner); ok {
		target.HttpConnectionManagerSettings = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_hcm.HttpConnectionManagerSettings)
	} else {
		target.HttpConnectionManagerSettings = proto.Clone(m.GetHttpConnectionManagerSettings()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_hcm.HttpConnectionManagerSettings)
	}

	if h, ok := interface{}(m.GetHealthCheck()).(clone.Cloner); ok {
		target.HealthCheck = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_healthcheck.HealthCheck)
	} else {
		target.HealthCheck = proto.Clone(m.GetHealthCheck()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_healthcheck.HealthCheck)
	}

	if h, ok := interface{}(m.GetExtensions()).(clone.Cloner); ok {
		target.Extensions = h.Clone().(*Extensions)
	} else {
		target.Extensions = proto.Clone(m.GetExtensions()).(*Extensions)
	}

	if h, ok := interface{}(m.GetWaf()).(clone.Cloner); ok {
		target.Waf = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_waf.Settings)
	} else {
		target.Waf = proto.Clone(m.GetWaf()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_waf.Settings)
	}

	if h, ok := interface{}(m.GetDlp()).(clone.Cloner); ok {
		target.Dlp = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_dlp.FilterConfig)
	} else {
		target.Dlp = proto.Clone(m.GetDlp()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_dlp.FilterConfig)
	}

	if h, ok := interface{}(m.GetWasm()).(clone.Cloner); ok {
		target.Wasm = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_wasm.PluginSource)
	} else {
		target.Wasm = proto.Clone(m.GetWasm()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_wasm.PluginSource)
	}

	if h, ok := interface{}(m.GetExtauth()).(clone.Cloner); ok {
		target.Extauth = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.Settings)
	} else {
		target.Extauth = proto.Clone(m.GetExtauth()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.Settings)
	}

	if h, ok := interface{}(m.GetRatelimitServer()).(clone.Cloner); ok {
		target.RatelimitServer = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.Settings)
	} else {
		target.RatelimitServer = proto.Clone(m.GetRatelimitServer()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.Settings)
	}

	if h, ok := interface{}(m.GetGzip()).(clone.Cloner); ok {
		target.Gzip = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_filter_http_gzip_v2.Gzip)
	} else {
		target.Gzip = proto.Clone(m.GetGzip()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_filter_http_gzip_v2.Gzip)
	}

	if h, ok := interface{}(m.GetProxyLatency()).(clone.Cloner); ok {
		target.ProxyLatency = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_proxylatency.ProxyLatency)
	} else {
		target.ProxyLatency = proto.Clone(m.GetProxyLatency()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_proxylatency.ProxyLatency)
	}

	if h, ok := interface{}(m.GetBuffer()).(clone.Cloner); ok {
		target.Buffer = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_buffer_v3.Buffer)
	} else {
		target.Buffer = proto.Clone(m.GetBuffer()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_buffer_v3.Buffer)
	}

	if h, ok := interface{}(m.GetCsrf()).(clone.Cloner); ok {
		target.Csrf = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_csrf_v3.CsrfPolicy)
	} else {
		target.Csrf = proto.Clone(m.GetCsrf()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_csrf_v3.CsrfPolicy)
	}

	if h, ok := interface{}(m.GetGrpcJsonTranscoder()).(clone.Cloner); ok {
		target.GrpcJsonTranscoder = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_grpc_json.GrpcJsonTranscoder)
	} else {
		target.GrpcJsonTranscoder = proto.Clone(m.GetGrpcJsonTranscoder()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_grpc_json.GrpcJsonTranscoder)
	}

	if h, ok := interface{}(m.GetSanitizeClusterHeader()).(clone.Cloner); ok {
		target.SanitizeClusterHeader = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.SanitizeClusterHeader = proto.Clone(m.GetSanitizeClusterHeader()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetLeftmostXffAddress()).(clone.Cloner); ok {
		target.LeftmostXffAddress = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.LeftmostXffAddress = proto.Clone(m.GetLeftmostXffAddress()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetDynamicForwardProxy()).(clone.Cloner); ok {
		target.DynamicForwardProxy = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_dynamic_forward_proxy.FilterConfig)
	} else {
		target.DynamicForwardProxy = proto.Clone(m.GetDynamicForwardProxy()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_dynamic_forward_proxy.FilterConfig)
	}

	return target
}

// Clone function
func (m *TcpListenerOptions) Clone() proto.Message {
	var target *TcpListenerOptions
	if m == nil {
		return target
	}
	target = &TcpListenerOptions{}

	if h, ok := interface{}(m.GetTcpProxySettings()).(clone.Cloner); ok {
		target.TcpProxySettings = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_tcp.TcpProxySettings)
	} else {
		target.TcpProxySettings = proto.Clone(m.GetTcpProxySettings()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_tcp.TcpProxySettings)
	}

	return target
}

// Clone function
func (m *VirtualHostOptions) Clone() proto.Message {
	var target *VirtualHostOptions
	if m == nil {
		return target
	}
	target = &VirtualHostOptions{}

	if h, ok := interface{}(m.GetExtensions()).(clone.Cloner); ok {
		target.Extensions = h.Clone().(*Extensions)
	} else {
		target.Extensions = proto.Clone(m.GetExtensions()).(*Extensions)
	}

	if h, ok := interface{}(m.GetRetries()).(clone.Cloner); ok {
		target.Retries = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_retries.RetryPolicy)
	} else {
		target.Retries = proto.Clone(m.GetRetries()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_retries.RetryPolicy)
	}

	if h, ok := interface{}(m.GetStats()).(clone.Cloner); ok {
		target.Stats = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_stats.Stats)
	} else {
		target.Stats = proto.Clone(m.GetStats()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_stats.Stats)
	}

	if h, ok := interface{}(m.GetHeaderManipulation()).(clone.Cloner); ok {
		target.HeaderManipulation = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_headers.HeaderManipulation)
	} else {
		target.HeaderManipulation = proto.Clone(m.GetHeaderManipulation()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_headers.HeaderManipulation)
	}

	if h, ok := interface{}(m.GetCors()).(clone.Cloner); ok {
		target.Cors = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_cors.CorsPolicy)
	} else {
		target.Cors = proto.Clone(m.GetCors()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_cors.CorsPolicy)
	}

	if h, ok := interface{}(m.GetTransformations()).(clone.Cloner); ok {
		target.Transformations = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.Transformations)
	} else {
		target.Transformations = proto.Clone(m.GetTransformations()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.Transformations)
	}

	if h, ok := interface{}(m.GetRatelimitBasic()).(clone.Cloner); ok {
		target.RatelimitBasic = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.IngressRateLimit)
	} else {
		target.RatelimitBasic = proto.Clone(m.GetRatelimitBasic()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.IngressRateLimit)
	}

	if h, ok := interface{}(m.GetWaf()).(clone.Cloner); ok {
		target.Waf = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_waf.Settings)
	} else {
		target.Waf = proto.Clone(m.GetWaf()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_waf.Settings)
	}

	if h, ok := interface{}(m.GetRbac()).(clone.Cloner); ok {
		target.Rbac = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_rbac.ExtensionSettings)
	} else {
		target.Rbac = proto.Clone(m.GetRbac()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_rbac.ExtensionSettings)
	}

	if h, ok := interface{}(m.GetExtauth()).(clone.Cloner); ok {
		target.Extauth = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.ExtAuthExtension)
	} else {
		target.Extauth = proto.Clone(m.GetExtauth()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.ExtAuthExtension)
	}

	if h, ok := interface{}(m.GetDlp()).(clone.Cloner); ok {
		target.Dlp = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_dlp.Config)
	} else {
		target.Dlp = proto.Clone(m.GetDlp()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_dlp.Config)
	}

	if h, ok := interface{}(m.GetBufferPerRoute()).(clone.Cloner); ok {
		target.BufferPerRoute = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_buffer_v3.BufferPerRoute)
	} else {
		target.BufferPerRoute = proto.Clone(m.GetBufferPerRoute()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_buffer_v3.BufferPerRoute)
	}

	if h, ok := interface{}(m.GetCsrf()).(clone.Cloner); ok {
		target.Csrf = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_csrf_v3.CsrfPolicy)
	} else {
		target.Csrf = proto.Clone(m.GetCsrf()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_csrf_v3.CsrfPolicy)
	}

	if h, ok := interface{}(m.GetIncludeRequestAttemptCount()).(clone.Cloner); ok {
		target.IncludeRequestAttemptCount = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.IncludeRequestAttemptCount = proto.Clone(m.GetIncludeRequestAttemptCount()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetIncludeAttemptCountInResponse()).(clone.Cloner); ok {
		target.IncludeAttemptCountInResponse = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	} else {
		target.IncludeAttemptCountInResponse = proto.Clone(m.GetIncludeAttemptCountInResponse()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue)
	}

	if h, ok := interface{}(m.GetStagedTransformations()).(clone.Cloner); ok {
		target.StagedTransformations = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.TransformationStages)
	} else {
		target.StagedTransformations = proto.Clone(m.GetStagedTransformations()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.TransformationStages)
	}

	switch m.RateLimitConfigType.(type) {

	case *VirtualHostOptions_Ratelimit:

		if h, ok := interface{}(m.GetRatelimit()).(clone.Cloner); ok {
			target.RateLimitConfigType = &VirtualHostOptions_Ratelimit{
				Ratelimit: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitVhostExtension),
			}
		} else {
			target.RateLimitConfigType = &VirtualHostOptions_Ratelimit{
				Ratelimit: proto.Clone(m.GetRatelimit()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitVhostExtension),
			}
		}

	case *VirtualHostOptions_RatelimitStaged:

		if h, ok := interface{}(m.GetRatelimitStaged()).(clone.Cloner); ok {
			target.RateLimitConfigType = &VirtualHostOptions_RatelimitStaged{
				RatelimitStaged: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitStagedVhostExtension),
			}
		} else {
			target.RateLimitConfigType = &VirtualHostOptions_RatelimitStaged{
				RatelimitStaged: proto.Clone(m.GetRatelimitStaged()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitStagedVhostExtension),
			}
		}

	case *VirtualHostOptions_RateLimitConfigs:

		if h, ok := interface{}(m.GetRateLimitConfigs()).(clone.Cloner); ok {
			target.RateLimitConfigType = &VirtualHostOptions_RateLimitConfigs{
				RateLimitConfigs: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitConfigRefs),
			}
		} else {
			target.RateLimitConfigType = &VirtualHostOptions_RateLimitConfigs{
				RateLimitConfigs: proto.Clone(m.GetRateLimitConfigs()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitConfigRefs),
			}
		}

	case *VirtualHostOptions_RateLimitStagedConfigs:

		if h, ok := interface{}(m.GetRateLimitStagedConfigs()).(clone.Cloner); ok {
			target.RateLimitConfigType = &VirtualHostOptions_RateLimitStagedConfigs{
				RateLimitStagedConfigs: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitStagedConfigRefs),
			}
		} else {
			target.RateLimitConfigType = &VirtualHostOptions_RateLimitStagedConfigs{
				RateLimitStagedConfigs: proto.Clone(m.GetRateLimitStagedConfigs()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitStagedConfigRefs),
			}
		}

	}

	switch m.JwtConfig.(type) {

	case *VirtualHostOptions_Jwt:

		if h, ok := interface{}(m.GetJwt()).(clone.Cloner); ok {
			target.JwtConfig = &VirtualHostOptions_Jwt{
				Jwt: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_jwt.VhostExtension),
			}
		} else {
			target.JwtConfig = &VirtualHostOptions_Jwt{
				Jwt: proto.Clone(m.GetJwt()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_jwt.VhostExtension),
			}
		}

	case *VirtualHostOptions_JwtStaged:

		if h, ok := interface{}(m.GetJwtStaged()).(clone.Cloner); ok {
			target.JwtConfig = &VirtualHostOptions_JwtStaged{
				JwtStaged: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_jwt.JwtStagedVhostExtension),
			}
		} else {
			target.JwtConfig = &VirtualHostOptions_JwtStaged{
				JwtStaged: proto.Clone(m.GetJwtStaged()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_jwt.JwtStagedVhostExtension),
			}
		}

	}

	return target
}

// Clone function
func (m *RouteOptions) Clone() proto.Message {
	var target *RouteOptions
	if m == nil {
		return target
	}
	target = &RouteOptions{}

	if h, ok := interface{}(m.GetTransformations()).(clone.Cloner); ok {
		target.Transformations = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.Transformations)
	} else {
		target.Transformations = proto.Clone(m.GetTransformations()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.Transformations)
	}

	if h, ok := interface{}(m.GetFaults()).(clone.Cloner); ok {
		target.Faults = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_faultinjection.RouteFaults)
	} else {
		target.Faults = proto.Clone(m.GetFaults()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_faultinjection.RouteFaults)
	}

	if h, ok := interface{}(m.GetPrefixRewrite()).(clone.Cloner); ok {
		target.PrefixRewrite = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	} else {
		target.PrefixRewrite = proto.Clone(m.GetPrefixRewrite()).(*github_com_golang_protobuf_ptypes_wrappers.StringValue)
	}

	if h, ok := interface{}(m.GetTimeout()).(clone.Cloner); ok {
		target.Timeout = h.Clone().(*github_com_golang_protobuf_ptypes_duration.Duration)
	} else {
		target.Timeout = proto.Clone(m.GetTimeout()).(*github_com_golang_protobuf_ptypes_duration.Duration)
	}

	if h, ok := interface{}(m.GetRetries()).(clone.Cloner); ok {
		target.Retries = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_retries.RetryPolicy)
	} else {
		target.Retries = proto.Clone(m.GetRetries()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_retries.RetryPolicy)
	}

	if h, ok := interface{}(m.GetExtensions()).(clone.Cloner); ok {
		target.Extensions = h.Clone().(*Extensions)
	} else {
		target.Extensions = proto.Clone(m.GetExtensions()).(*Extensions)
	}

	if h, ok := interface{}(m.GetTracing()).(clone.Cloner); ok {
		target.Tracing = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_tracing.RouteTracingSettings)
	} else {
		target.Tracing = proto.Clone(m.GetTracing()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_tracing.RouteTracingSettings)
	}

	if h, ok := interface{}(m.GetShadowing()).(clone.Cloner); ok {
		target.Shadowing = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_shadowing.RouteShadowing)
	} else {
		target.Shadowing = proto.Clone(m.GetShadowing()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_shadowing.RouteShadowing)
	}

	if h, ok := interface{}(m.GetHeaderManipulation()).(clone.Cloner); ok {
		target.HeaderManipulation = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_headers.HeaderManipulation)
	} else {
		target.HeaderManipulation = proto.Clone(m.GetHeaderManipulation()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_headers.HeaderManipulation)
	}

	if h, ok := interface{}(m.GetCors()).(clone.Cloner); ok {
		target.Cors = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_cors.CorsPolicy)
	} else {
		target.Cors = proto.Clone(m.GetCors()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_cors.CorsPolicy)
	}

	if h, ok := interface{}(m.GetLbHash()).(clone.Cloner); ok {
		target.LbHash = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_lbhash.RouteActionHashConfig)
	} else {
		target.LbHash = proto.Clone(m.GetLbHash()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_lbhash.RouteActionHashConfig)
	}

	if m.GetUpgrades() != nil {
		target.Upgrades = make([]*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade.ProtocolUpgradeConfig, len(m.GetUpgrades()))
		for idx, v := range m.GetUpgrades() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Upgrades[idx] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade.ProtocolUpgradeConfig)
			} else {
				target.Upgrades[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade.ProtocolUpgradeConfig)
			}

		}
	}

	if h, ok := interface{}(m.GetRatelimitBasic()).(clone.Cloner); ok {
		target.RatelimitBasic = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.IngressRateLimit)
	} else {
		target.RatelimitBasic = proto.Clone(m.GetRatelimitBasic()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.IngressRateLimit)
	}

	if h, ok := interface{}(m.GetWaf()).(clone.Cloner); ok {
		target.Waf = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_waf.Settings)
	} else {
		target.Waf = proto.Clone(m.GetWaf()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_waf.Settings)
	}

	if h, ok := interface{}(m.GetRbac()).(clone.Cloner); ok {
		target.Rbac = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_rbac.ExtensionSettings)
	} else {
		target.Rbac = proto.Clone(m.GetRbac()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_rbac.ExtensionSettings)
	}

	if h, ok := interface{}(m.GetExtauth()).(clone.Cloner); ok {
		target.Extauth = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.ExtAuthExtension)
	} else {
		target.Extauth = proto.Clone(m.GetExtauth()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.ExtAuthExtension)
	}

	if h, ok := interface{}(m.GetDlp()).(clone.Cloner); ok {
		target.Dlp = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_dlp.Config)
	} else {
		target.Dlp = proto.Clone(m.GetDlp()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_dlp.Config)
	}

	if h, ok := interface{}(m.GetBufferPerRoute()).(clone.Cloner); ok {
		target.BufferPerRoute = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_buffer_v3.BufferPerRoute)
	} else {
		target.BufferPerRoute = proto.Clone(m.GetBufferPerRoute()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_buffer_v3.BufferPerRoute)
	}

	if h, ok := interface{}(m.GetCsrf()).(clone.Cloner); ok {
		target.Csrf = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_csrf_v3.CsrfPolicy)
	} else {
		target.Csrf = proto.Clone(m.GetCsrf()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_csrf_v3.CsrfPolicy)
	}

	if h, ok := interface{}(m.GetStagedTransformations()).(clone.Cloner); ok {
		target.StagedTransformations = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.TransformationStages)
	} else {
		target.StagedTransformations = proto.Clone(m.GetStagedTransformations()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.TransformationStages)
	}

	if m.GetEnvoyMetadata() != nil {
		target.EnvoyMetadata = make(map[string]*github_com_golang_protobuf_ptypes_struct.Struct, len(m.GetEnvoyMetadata()))
		for k, v := range m.GetEnvoyMetadata() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.EnvoyMetadata[k] = h.Clone().(*github_com_golang_protobuf_ptypes_struct.Struct)
			} else {
				target.EnvoyMetadata[k] = proto.Clone(v).(*github_com_golang_protobuf_ptypes_struct.Struct)
			}

		}
	}

	if h, ok := interface{}(m.GetRegexRewrite()).(clone.Cloner); ok {
		target.RegexRewrite = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_type_matcher_v3.RegexMatchAndSubstitute)
	} else {
		target.RegexRewrite = proto.Clone(m.GetRegexRewrite()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_type_matcher_v3.RegexMatchAndSubstitute)
	}

	switch m.HostRewriteType.(type) {

	case *RouteOptions_HostRewrite:

		target.HostRewriteType = &RouteOptions_HostRewrite{
			HostRewrite: m.GetHostRewrite(),
		}

	case *RouteOptions_AutoHostRewrite:

		if h, ok := interface{}(m.GetAutoHostRewrite()).(clone.Cloner); ok {
			target.HostRewriteType = &RouteOptions_AutoHostRewrite{
				AutoHostRewrite: h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.BoolValue),
			}
		} else {
			target.HostRewriteType = &RouteOptions_AutoHostRewrite{
				AutoHostRewrite: proto.Clone(m.GetAutoHostRewrite()).(*github_com_golang_protobuf_ptypes_wrappers.BoolValue),
			}
		}

	}

	switch m.RateLimitConfigType.(type) {

	case *RouteOptions_Ratelimit:

		if h, ok := interface{}(m.GetRatelimit()).(clone.Cloner); ok {
			target.RateLimitConfigType = &RouteOptions_Ratelimit{
				Ratelimit: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitRouteExtension),
			}
		} else {
			target.RateLimitConfigType = &RouteOptions_Ratelimit{
				Ratelimit: proto.Clone(m.GetRatelimit()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitRouteExtension),
			}
		}

	case *RouteOptions_RatelimitStaged:

		if h, ok := interface{}(m.GetRatelimitStaged()).(clone.Cloner); ok {
			target.RateLimitConfigType = &RouteOptions_RatelimitStaged{
				RatelimitStaged: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitStagedRouteExtension),
			}
		} else {
			target.RateLimitConfigType = &RouteOptions_RatelimitStaged{
				RatelimitStaged: proto.Clone(m.GetRatelimitStaged()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitStagedRouteExtension),
			}
		}

	case *RouteOptions_RateLimitConfigs:

		if h, ok := interface{}(m.GetRateLimitConfigs()).(clone.Cloner); ok {
			target.RateLimitConfigType = &RouteOptions_RateLimitConfigs{
				RateLimitConfigs: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitConfigRefs),
			}
		} else {
			target.RateLimitConfigType = &RouteOptions_RateLimitConfigs{
				RateLimitConfigs: proto.Clone(m.GetRateLimitConfigs()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitConfigRefs),
			}
		}

	case *RouteOptions_RateLimitStagedConfigs:

		if h, ok := interface{}(m.GetRateLimitStagedConfigs()).(clone.Cloner); ok {
			target.RateLimitConfigType = &RouteOptions_RateLimitStagedConfigs{
				RateLimitStagedConfigs: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitStagedConfigRefs),
			}
		} else {
			target.RateLimitConfigType = &RouteOptions_RateLimitStagedConfigs{
				RateLimitStagedConfigs: proto.Clone(m.GetRateLimitStagedConfigs()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_ratelimit.RateLimitStagedConfigRefs),
			}
		}

	}

	switch m.JwtConfig.(type) {

	case *RouteOptions_Jwt:

		if h, ok := interface{}(m.GetJwt()).(clone.Cloner); ok {
			target.JwtConfig = &RouteOptions_Jwt{
				Jwt: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_jwt.RouteExtension),
			}
		} else {
			target.JwtConfig = &RouteOptions_Jwt{
				Jwt: proto.Clone(m.GetJwt()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_jwt.RouteExtension),
			}
		}

	case *RouteOptions_JwtStaged:

		if h, ok := interface{}(m.GetJwtStaged()).(clone.Cloner); ok {
			target.JwtConfig = &RouteOptions_JwtStaged{
				JwtStaged: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_jwt.JwtStagedRouteExtension),
			}
		} else {
			target.JwtConfig = &RouteOptions_JwtStaged{
				JwtStaged: proto.Clone(m.GetJwtStaged()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_jwt.JwtStagedRouteExtension),
			}
		}

	}

	return target
}

// Clone function
func (m *DestinationSpec) Clone() proto.Message {
	var target *DestinationSpec
	if m == nil {
		return target
	}
	target = &DestinationSpec{}

	switch m.DestinationType.(type) {

	case *DestinationSpec_Aws:

		if h, ok := interface{}(m.GetAws()).(clone.Cloner); ok {
			target.DestinationType = &DestinationSpec_Aws{
				Aws: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_aws.DestinationSpec),
			}
		} else {
			target.DestinationType = &DestinationSpec_Aws{
				Aws: proto.Clone(m.GetAws()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_aws.DestinationSpec),
			}
		}

	case *DestinationSpec_Azure:

		if h, ok := interface{}(m.GetAzure()).(clone.Cloner); ok {
			target.DestinationType = &DestinationSpec_Azure{
				Azure: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_azure.DestinationSpec),
			}
		} else {
			target.DestinationType = &DestinationSpec_Azure{
				Azure: proto.Clone(m.GetAzure()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_azure.DestinationSpec),
			}
		}

	case *DestinationSpec_Rest:

		if h, ok := interface{}(m.GetRest()).(clone.Cloner); ok {
			target.DestinationType = &DestinationSpec_Rest{
				Rest: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_rest.DestinationSpec),
			}
		} else {
			target.DestinationType = &DestinationSpec_Rest{
				Rest: proto.Clone(m.GetRest()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_rest.DestinationSpec),
			}
		}

	case *DestinationSpec_Grpc:

		if h, ok := interface{}(m.GetGrpc()).(clone.Cloner); ok {
			target.DestinationType = &DestinationSpec_Grpc{
				Grpc: h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_grpc.DestinationSpec),
			}
		} else {
			target.DestinationType = &DestinationSpec_Grpc{
				Grpc: proto.Clone(m.GetGrpc()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_grpc.DestinationSpec),
			}
		}

	}

	return target
}

// Clone function
func (m *WeightedDestinationOptions) Clone() proto.Message {
	var target *WeightedDestinationOptions
	if m == nil {
		return target
	}
	target = &WeightedDestinationOptions{}

	if h, ok := interface{}(m.GetHeaderManipulation()).(clone.Cloner); ok {
		target.HeaderManipulation = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_headers.HeaderManipulation)
	} else {
		target.HeaderManipulation = proto.Clone(m.GetHeaderManipulation()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_headers.HeaderManipulation)
	}

	if h, ok := interface{}(m.GetTransformations()).(clone.Cloner); ok {
		target.Transformations = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.Transformations)
	} else {
		target.Transformations = proto.Clone(m.GetTransformations()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.Transformations)
	}

	if h, ok := interface{}(m.GetExtensions()).(clone.Cloner); ok {
		target.Extensions = h.Clone().(*Extensions)
	} else {
		target.Extensions = proto.Clone(m.GetExtensions()).(*Extensions)
	}

	if h, ok := interface{}(m.GetExtauth()).(clone.Cloner); ok {
		target.Extauth = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.ExtAuthExtension)
	} else {
		target.Extauth = proto.Clone(m.GetExtauth()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_enterprise_options_extauth_v1.ExtAuthExtension)
	}

	if h, ok := interface{}(m.GetBufferPerRoute()).(clone.Cloner); ok {
		target.BufferPerRoute = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_buffer_v3.BufferPerRoute)
	} else {
		target.BufferPerRoute = proto.Clone(m.GetBufferPerRoute()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_buffer_v3.BufferPerRoute)
	}

	if h, ok := interface{}(m.GetCsrf()).(clone.Cloner); ok {
		target.Csrf = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_csrf_v3.CsrfPolicy)
	} else {
		target.Csrf = proto.Clone(m.GetCsrf()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_filters_http_csrf_v3.CsrfPolicy)
	}

	if h, ok := interface{}(m.GetStagedTransformations()).(clone.Cloner); ok {
		target.StagedTransformations = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.TransformationStages)
	} else {
		target.StagedTransformations = proto.Clone(m.GetStagedTransformations()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_transformation.TransformationStages)
	}

	return target
}
