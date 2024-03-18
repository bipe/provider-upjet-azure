// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConnectionMonitorInitParameters struct {

	// A endpoint block as defined below.
	Endpoint []EndpointInitParameters `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The Azure Region where the Network Connection Monitor should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The description of the Network Connection Monitor.
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`

	// A list of IDs of the Log Analytics Workspace which will accept the output from the Network Connection Monitor.
	// +listType=set
	OutputWorkspaceResourceIds []*string `json:"outputWorkspaceResourceIds,omitempty" tf:"output_workspace_resource_ids,omitempty"`

	// A mapping of tags which should be assigned to the Network Connection Monitor.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A test_configuration block as defined below.
	TestConfiguration []TestConfigurationInitParameters `json:"testConfiguration,omitempty" tf:"test_configuration,omitempty"`

	// A test_group block as defined below.
	TestGroup []TestGroupInitParameters `json:"testGroup,omitempty" tf:"test_group,omitempty"`
}

type ConnectionMonitorObservation struct {

	// A endpoint block as defined below.
	Endpoint []EndpointObservation `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The ID of the Network Connection Monitor.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where the Network Connection Monitor should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherID *string `json:"networkWatcherId,omitempty" tf:"network_watcher_id,omitempty"`

	// The description of the Network Connection Monitor.
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`

	// A list of IDs of the Log Analytics Workspace which will accept the output from the Network Connection Monitor.
	// +listType=set
	OutputWorkspaceResourceIds []*string `json:"outputWorkspaceResourceIds,omitempty" tf:"output_workspace_resource_ids,omitempty"`

	// A mapping of tags which should be assigned to the Network Connection Monitor.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A test_configuration block as defined below.
	TestConfiguration []TestConfigurationObservation `json:"testConfiguration,omitempty" tf:"test_configuration,omitempty"`

	// A test_group block as defined below.
	TestGroup []TestGroupObservation `json:"testGroup,omitempty" tf:"test_group,omitempty"`
}

type ConnectionMonitorParameters struct {

	// A endpoint block as defined below.
	// +kubebuilder:validation:Optional
	Endpoint []EndpointParameters `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The Azure Region where the Network Connection Monitor should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Network Watcher. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Watcher
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkWatcherID *string `json:"networkWatcherId,omitempty" tf:"network_watcher_id,omitempty"`

	// Reference to a Watcher to populate networkWatcherId.
	// +kubebuilder:validation:Optional
	NetworkWatcherIDRef *v1.Reference `json:"networkWatcherIdRef,omitempty" tf:"-"`

	// Selector for a Watcher to populate networkWatcherId.
	// +kubebuilder:validation:Optional
	NetworkWatcherIDSelector *v1.Selector `json:"networkWatcherIdSelector,omitempty" tf:"-"`

	// The description of the Network Connection Monitor.
	// +kubebuilder:validation:Optional
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`

	// A list of IDs of the Log Analytics Workspace which will accept the output from the Network Connection Monitor.
	// +kubebuilder:validation:Optional
	// +listType=set
	OutputWorkspaceResourceIds []*string `json:"outputWorkspaceResourceIds,omitempty" tf:"output_workspace_resource_ids,omitempty"`

	// A mapping of tags which should be assigned to the Network Connection Monitor.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A test_configuration block as defined below.
	// +kubebuilder:validation:Optional
	TestConfiguration []TestConfigurationParameters `json:"testConfiguration,omitempty" tf:"test_configuration,omitempty"`

	// A test_group block as defined below.
	// +kubebuilder:validation:Optional
	TestGroup []TestGroupParameters `json:"testGroup,omitempty" tf:"test_group,omitempty"`
}

type EndpointInitParameters struct {

	// The IP address or domain name of the Network Connection Monitor endpoint.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The test coverage for the Network Connection Monitor endpoint. Possible values are AboveAverage, Average, BelowAverage, Default, Full and Low.
	CoverageLevel *string `json:"coverageLevel,omitempty" tf:"coverage_level,omitempty"`

	// A list of IPv4/IPv6 subnet masks or IPv4/IPv6 IP addresses to be excluded to the Network Connection Monitor endpoint.
	// +listType=set
	ExcludedIPAddresses []*string `json:"excludedIpAddresses,omitempty" tf:"excluded_ip_addresses,omitempty"`

	// A filter block as defined below.
	Filter []FilterInitParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// A list of IPv4/IPv6 subnet masks or IPv4/IPv6 IP addresses to be included to the Network Connection Monitor endpoint.
	// +listType=set
	IncludedIPAddresses []*string `json:"includedIpAddresses,omitempty" tf:"included_ip_addresses,omitempty"`

	// The name of the endpoint for the Network Connection Monitor .
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The resource ID which is used as the endpoint by the Network Connection Monitor.
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// The endpoint type of the Network Connection Monitor. Possible values are AzureSubnet, AzureVM, AzureVNet, ExternalAddress, MMAWorkspaceMachine and MMAWorkspaceNetwork.
	TargetResourceType *string `json:"targetResourceType,omitempty" tf:"target_resource_type,omitempty"`
}

type EndpointObservation struct {

	// The IP address or domain name of the Network Connection Monitor endpoint.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The test coverage for the Network Connection Monitor endpoint. Possible values are AboveAverage, Average, BelowAverage, Default, Full and Low.
	CoverageLevel *string `json:"coverageLevel,omitempty" tf:"coverage_level,omitempty"`

	// A list of IPv4/IPv6 subnet masks or IPv4/IPv6 IP addresses to be excluded to the Network Connection Monitor endpoint.
	// +listType=set
	ExcludedIPAddresses []*string `json:"excludedIpAddresses,omitempty" tf:"excluded_ip_addresses,omitempty"`

	// A filter block as defined below.
	Filter []FilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	// A list of IPv4/IPv6 subnet masks or IPv4/IPv6 IP addresses to be included to the Network Connection Monitor endpoint.
	// +listType=set
	IncludedIPAddresses []*string `json:"includedIpAddresses,omitempty" tf:"included_ip_addresses,omitempty"`

	// The name of the endpoint for the Network Connection Monitor .
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The resource ID which is used as the endpoint by the Network Connection Monitor.
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// The endpoint type of the Network Connection Monitor. Possible values are AzureSubnet, AzureVM, AzureVNet, ExternalAddress, MMAWorkspaceMachine and MMAWorkspaceNetwork.
	TargetResourceType *string `json:"targetResourceType,omitempty" tf:"target_resource_type,omitempty"`
}

type EndpointParameters struct {

	// The IP address or domain name of the Network Connection Monitor endpoint.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The test coverage for the Network Connection Monitor endpoint. Possible values are AboveAverage, Average, BelowAverage, Default, Full and Low.
	// +kubebuilder:validation:Optional
	CoverageLevel *string `json:"coverageLevel,omitempty" tf:"coverage_level,omitempty"`

	// A list of IPv4/IPv6 subnet masks or IPv4/IPv6 IP addresses to be excluded to the Network Connection Monitor endpoint.
	// +kubebuilder:validation:Optional
	// +listType=set
	ExcludedIPAddresses []*string `json:"excludedIpAddresses,omitempty" tf:"excluded_ip_addresses,omitempty"`

	// A filter block as defined below.
	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// A list of IPv4/IPv6 subnet masks or IPv4/IPv6 IP addresses to be included to the Network Connection Monitor endpoint.
	// +kubebuilder:validation:Optional
	// +listType=set
	IncludedIPAddresses []*string `json:"includedIpAddresses,omitempty" tf:"included_ip_addresses,omitempty"`

	// The name of the endpoint for the Network Connection Monitor .
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The resource ID which is used as the endpoint by the Network Connection Monitor.
	// +kubebuilder:validation:Optional
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// The endpoint type of the Network Connection Monitor. Possible values are AzureSubnet, AzureVM, AzureVNet, ExternalAddress, MMAWorkspaceMachine and MMAWorkspaceNetwork.
	// +kubebuilder:validation:Optional
	TargetResourceType *string `json:"targetResourceType,omitempty" tf:"target_resource_type,omitempty"`
}

type FilterInitParameters struct {

	// A item block as defined below.
	Item []ItemInitParameters `json:"item,omitempty" tf:"item,omitempty"`

	// The type of items included in the filter. Possible values are AgentAddress. Defaults to AgentAddress.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FilterObservation struct {

	// A item block as defined below.
	Item []ItemObservation `json:"item,omitempty" tf:"item,omitempty"`

	// The type of items included in the filter. Possible values are AgentAddress. Defaults to AgentAddress.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FilterParameters struct {

	// A item block as defined below.
	// +kubebuilder:validation:Optional
	Item []ItemParameters `json:"item,omitempty" tf:"item,omitempty"`

	// The type of items included in the filter. Possible values are AgentAddress. Defaults to AgentAddress.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HTTPConfigurationInitParameters struct {

	// The HTTP method for the HTTP request. Possible values are Get and Post. Defaults to Get.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// The path component of the URI. It only accepts the absolute path.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The port for the TCP connection.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Should HTTPS be preferred over HTTP in cases where the choice is not explicit? Defaults to false.
	PreferHTTPS *bool `json:"preferHttps,omitempty" tf:"prefer_https,omitempty"`

	// A request_header block as defined below.
	RequestHeader []HTTPConfigurationRequestHeaderInitParameters `json:"requestHeader,omitempty" tf:"request_header,omitempty"`

	// The HTTP status codes to consider successful. For instance, 2xx, 301-304 and 418.
	// +listType=set
	ValidStatusCodeRanges []*string `json:"validStatusCodeRanges,omitempty" tf:"valid_status_code_ranges,omitempty"`
}

type HTTPConfigurationObservation struct {

	// The HTTP method for the HTTP request. Possible values are Get and Post. Defaults to Get.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// The path component of the URI. It only accepts the absolute path.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The port for the TCP connection.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Should HTTPS be preferred over HTTP in cases where the choice is not explicit? Defaults to false.
	PreferHTTPS *bool `json:"preferHttps,omitempty" tf:"prefer_https,omitempty"`

	// A request_header block as defined below.
	RequestHeader []HTTPConfigurationRequestHeaderObservation `json:"requestHeader,omitempty" tf:"request_header,omitempty"`

	// The HTTP status codes to consider successful. For instance, 2xx, 301-304 and 418.
	// +listType=set
	ValidStatusCodeRanges []*string `json:"validStatusCodeRanges,omitempty" tf:"valid_status_code_ranges,omitempty"`
}

type HTTPConfigurationParameters struct {

	// The HTTP method for the HTTP request. Possible values are Get and Post. Defaults to Get.
	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// The path component of the URI. It only accepts the absolute path.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The port for the TCP connection.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Should HTTPS be preferred over HTTP in cases where the choice is not explicit? Defaults to false.
	// +kubebuilder:validation:Optional
	PreferHTTPS *bool `json:"preferHttps,omitempty" tf:"prefer_https,omitempty"`

	// A request_header block as defined below.
	// +kubebuilder:validation:Optional
	RequestHeader []HTTPConfigurationRequestHeaderParameters `json:"requestHeader,omitempty" tf:"request_header,omitempty"`

	// The HTTP status codes to consider successful. For instance, 2xx, 301-304 and 418.
	// +kubebuilder:validation:Optional
	// +listType=set
	ValidStatusCodeRanges []*string `json:"validStatusCodeRanges,omitempty" tf:"valid_status_code_ranges,omitempty"`
}

type HTTPConfigurationRequestHeaderInitParameters struct {

	// The name of the test group for the Network Connection Monitor.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the HTTP header.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type HTTPConfigurationRequestHeaderObservation struct {

	// The name of the test group for the Network Connection Monitor.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the HTTP header.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type HTTPConfigurationRequestHeaderParameters struct {

	// The name of the test group for the Network Connection Monitor.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The value of the HTTP header.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type IcmpConfigurationInitParameters struct {

	// Should path evaluation with trace route be enabled? Defaults to true.
	TraceRouteEnabled *bool `json:"traceRouteEnabled,omitempty" tf:"trace_route_enabled,omitempty"`
}

type IcmpConfigurationObservation struct {

	// Should path evaluation with trace route be enabled? Defaults to true.
	TraceRouteEnabled *bool `json:"traceRouteEnabled,omitempty" tf:"trace_route_enabled,omitempty"`
}

type IcmpConfigurationParameters struct {

	// Should path evaluation with trace route be enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	TraceRouteEnabled *bool `json:"traceRouteEnabled,omitempty" tf:"trace_route_enabled,omitempty"`
}

type ItemInitParameters struct {

	// The address of the filter item.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The type of items included in the filter. Possible values are AgentAddress. Defaults to AgentAddress.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ItemObservation struct {

	// The address of the filter item.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The type of items included in the filter. Possible values are AgentAddress. Defaults to AgentAddress.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ItemParameters struct {

	// The address of the filter item.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The type of items included in the filter. Possible values are AgentAddress. Defaults to AgentAddress.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SuccessThresholdInitParameters struct {

	// The maximum percentage of failed checks permitted for a test to be successful.
	ChecksFailedPercent *float64 `json:"checksFailedPercent,omitempty" tf:"checks_failed_percent,omitempty"`

	// The maximum round-trip time in milliseconds permitted for a test to be successful.
	RoundTripTimeMS *float64 `json:"roundTripTimeMs,omitempty" tf:"round_trip_time_ms,omitempty"`
}

type SuccessThresholdObservation struct {

	// The maximum percentage of failed checks permitted for a test to be successful.
	ChecksFailedPercent *float64 `json:"checksFailedPercent,omitempty" tf:"checks_failed_percent,omitempty"`

	// The maximum round-trip time in milliseconds permitted for a test to be successful.
	RoundTripTimeMS *float64 `json:"roundTripTimeMs,omitempty" tf:"round_trip_time_ms,omitempty"`
}

type SuccessThresholdParameters struct {

	// The maximum percentage of failed checks permitted for a test to be successful.
	// +kubebuilder:validation:Optional
	ChecksFailedPercent *float64 `json:"checksFailedPercent,omitempty" tf:"checks_failed_percent,omitempty"`

	// The maximum round-trip time in milliseconds permitted for a test to be successful.
	// +kubebuilder:validation:Optional
	RoundTripTimeMS *float64 `json:"roundTripTimeMs,omitempty" tf:"round_trip_time_ms,omitempty"`
}

type TCPConfigurationInitParameters struct {

	// The destination port behavior for the TCP connection. Possible values are None and ListenIfAvailable.
	DestinationPortBehavior *string `json:"destinationPortBehavior,omitempty" tf:"destination_port_behavior,omitempty"`

	// The port for the TCP connection.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Should path evaluation with trace route be enabled? Defaults to true.
	TraceRouteEnabled *bool `json:"traceRouteEnabled,omitempty" tf:"trace_route_enabled,omitempty"`
}

type TCPConfigurationObservation struct {

	// The destination port behavior for the TCP connection. Possible values are None and ListenIfAvailable.
	DestinationPortBehavior *string `json:"destinationPortBehavior,omitempty" tf:"destination_port_behavior,omitempty"`

	// The port for the TCP connection.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Should path evaluation with trace route be enabled? Defaults to true.
	TraceRouteEnabled *bool `json:"traceRouteEnabled,omitempty" tf:"trace_route_enabled,omitempty"`
}

type TCPConfigurationParameters struct {

	// The destination port behavior for the TCP connection. Possible values are None and ListenIfAvailable.
	// +kubebuilder:validation:Optional
	DestinationPortBehavior *string `json:"destinationPortBehavior,omitempty" tf:"destination_port_behavior,omitempty"`

	// The port for the TCP connection.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port" tf:"port,omitempty"`

	// Should path evaluation with trace route be enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	TraceRouteEnabled *bool `json:"traceRouteEnabled,omitempty" tf:"trace_route_enabled,omitempty"`
}

type TestConfigurationInitParameters struct {

	// A http_configuration block as defined below.
	HTTPConfiguration []HTTPConfigurationInitParameters `json:"httpConfiguration,omitempty" tf:"http_configuration,omitempty"`

	// A icmp_configuration block as defined below.
	IcmpConfiguration []IcmpConfigurationInitParameters `json:"icmpConfiguration,omitempty" tf:"icmp_configuration,omitempty"`

	// The name of test configuration for the Network Connection Monitor.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The preferred IP version which is used in the test evaluation. Possible values are IPv4 and IPv6.
	PreferredIPVersion *string `json:"preferredIpVersion,omitempty" tf:"preferred_ip_version,omitempty"`

	// The protocol used to evaluate tests. Possible values are Tcp, Http and Icmp.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// A success_threshold block as defined below.
	SuccessThreshold []SuccessThresholdInitParameters `json:"successThreshold,omitempty" tf:"success_threshold,omitempty"`

	// A tcp_configuration block as defined below.
	TCPConfiguration []TCPConfigurationInitParameters `json:"tcpConfiguration,omitempty" tf:"tcp_configuration,omitempty"`

	// The time interval in seconds at which the test evaluation will happen. Defaults to 60.
	TestFrequencyInSeconds *float64 `json:"testFrequencyInSeconds,omitempty" tf:"test_frequency_in_seconds,omitempty"`
}

type TestConfigurationObservation struct {

	// A http_configuration block as defined below.
	HTTPConfiguration []HTTPConfigurationObservation `json:"httpConfiguration,omitempty" tf:"http_configuration,omitempty"`

	// A icmp_configuration block as defined below.
	IcmpConfiguration []IcmpConfigurationObservation `json:"icmpConfiguration,omitempty" tf:"icmp_configuration,omitempty"`

	// The name of test configuration for the Network Connection Monitor.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The preferred IP version which is used in the test evaluation. Possible values are IPv4 and IPv6.
	PreferredIPVersion *string `json:"preferredIpVersion,omitempty" tf:"preferred_ip_version,omitempty"`

	// The protocol used to evaluate tests. Possible values are Tcp, Http and Icmp.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// A success_threshold block as defined below.
	SuccessThreshold []SuccessThresholdObservation `json:"successThreshold,omitempty" tf:"success_threshold,omitempty"`

	// A tcp_configuration block as defined below.
	TCPConfiguration []TCPConfigurationObservation `json:"tcpConfiguration,omitempty" tf:"tcp_configuration,omitempty"`

	// The time interval in seconds at which the test evaluation will happen. Defaults to 60.
	TestFrequencyInSeconds *float64 `json:"testFrequencyInSeconds,omitempty" tf:"test_frequency_in_seconds,omitempty"`
}

type TestConfigurationParameters struct {

	// A http_configuration block as defined below.
	// +kubebuilder:validation:Optional
	HTTPConfiguration []HTTPConfigurationParameters `json:"httpConfiguration,omitempty" tf:"http_configuration,omitempty"`

	// A icmp_configuration block as defined below.
	// +kubebuilder:validation:Optional
	IcmpConfiguration []IcmpConfigurationParameters `json:"icmpConfiguration,omitempty" tf:"icmp_configuration,omitempty"`

	// The name of test configuration for the Network Connection Monitor.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The preferred IP version which is used in the test evaluation. Possible values are IPv4 and IPv6.
	// +kubebuilder:validation:Optional
	PreferredIPVersion *string `json:"preferredIpVersion,omitempty" tf:"preferred_ip_version,omitempty"`

	// The protocol used to evaluate tests. Possible values are Tcp, Http and Icmp.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// A success_threshold block as defined below.
	// +kubebuilder:validation:Optional
	SuccessThreshold []SuccessThresholdParameters `json:"successThreshold,omitempty" tf:"success_threshold,omitempty"`

	// A tcp_configuration block as defined below.
	// +kubebuilder:validation:Optional
	TCPConfiguration []TCPConfigurationParameters `json:"tcpConfiguration,omitempty" tf:"tcp_configuration,omitempty"`

	// The time interval in seconds at which the test evaluation will happen. Defaults to 60.
	// +kubebuilder:validation:Optional
	TestFrequencyInSeconds *float64 `json:"testFrequencyInSeconds,omitempty" tf:"test_frequency_in_seconds,omitempty"`
}

type TestGroupInitParameters struct {

	// A list of destination endpoint names.
	// +listType=set
	DestinationEndpoints []*string `json:"destinationEndpoints,omitempty" tf:"destination_endpoints,omitempty"`

	// Should the test group be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name of the test group for the Network Connection Monitor.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of source endpoint names.
	// +listType=set
	SourceEndpoints []*string `json:"sourceEndpoints,omitempty" tf:"source_endpoints,omitempty"`

	// A list of test configuration names.
	// +listType=set
	TestConfigurationNames []*string `json:"testConfigurationNames,omitempty" tf:"test_configuration_names,omitempty"`
}

type TestGroupObservation struct {

	// A list of destination endpoint names.
	// +listType=set
	DestinationEndpoints []*string `json:"destinationEndpoints,omitempty" tf:"destination_endpoints,omitempty"`

	// Should the test group be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name of the test group for the Network Connection Monitor.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of source endpoint names.
	// +listType=set
	SourceEndpoints []*string `json:"sourceEndpoints,omitempty" tf:"source_endpoints,omitempty"`

	// A list of test configuration names.
	// +listType=set
	TestConfigurationNames []*string `json:"testConfigurationNames,omitempty" tf:"test_configuration_names,omitempty"`
}

type TestGroupParameters struct {

	// A list of destination endpoint names.
	// +kubebuilder:validation:Optional
	// +listType=set
	DestinationEndpoints []*string `json:"destinationEndpoints" tf:"destination_endpoints,omitempty"`

	// Should the test group be enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name of the test group for the Network Connection Monitor.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A list of source endpoint names.
	// +kubebuilder:validation:Optional
	// +listType=set
	SourceEndpoints []*string `json:"sourceEndpoints" tf:"source_endpoints,omitempty"`

	// A list of test configuration names.
	// +kubebuilder:validation:Optional
	// +listType=set
	TestConfigurationNames []*string `json:"testConfigurationNames" tf:"test_configuration_names,omitempty"`
}

// ConnectionMonitorSpec defines the desired state of ConnectionMonitor
type ConnectionMonitorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionMonitorParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ConnectionMonitorInitParameters `json:"initProvider,omitempty"`
}

// ConnectionMonitorStatus defines the observed state of ConnectionMonitor.
type ConnectionMonitorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionMonitorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ConnectionMonitor is the Schema for the ConnectionMonitors API. Manages a Network Connection Monitor.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ConnectionMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.endpoint) || (has(self.initProvider) && has(self.initProvider.endpoint))",message="spec.forProvider.endpoint is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.testConfiguration) || (has(self.initProvider) && has(self.initProvider.testConfiguration))",message="spec.forProvider.testConfiguration is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.testGroup) || (has(self.initProvider) && has(self.initProvider.testGroup))",message="spec.forProvider.testGroup is a required parameter"
	Spec   ConnectionMonitorSpec   `json:"spec"`
	Status ConnectionMonitorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionMonitorList contains a list of ConnectionMonitors
type ConnectionMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectionMonitor `json:"items"`
}

// Repository type metadata.
var (
	ConnectionMonitor_Kind             = "ConnectionMonitor"
	ConnectionMonitor_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConnectionMonitor_Kind}.String()
	ConnectionMonitor_KindAPIVersion   = ConnectionMonitor_Kind + "." + CRDGroupVersion.String()
	ConnectionMonitor_GroupVersionKind = CRDGroupVersion.WithKind(ConnectionMonitor_Kind)
)

func init() {
	SchemeBuilder.Register(&ConnectionMonitor{}, &ConnectionMonitorList{})
}
