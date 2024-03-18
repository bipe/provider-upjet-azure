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

type CustomDNSConfigsInitParameters struct {
}

type CustomDNSConfigsObservation struct {

	// The fully qualified domain name to the private_endpoint.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// A list of all IP Addresses that map to the private_endpoint fqdn.
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`
}

type CustomDNSConfigsParameters struct {
}

type PrivateDNSZoneConfigsInitParameters struct {
}

type PrivateDNSZoneConfigsObservation struct {

	// The ID of the Private DNS Zone Config.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Private DNS Zone that the config belongs to.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of IP Addresses
	PrivateDNSZoneID *string `json:"privateDnsZoneId,omitempty" tf:"private_dns_zone_id,omitempty"`

	// A record_sets block as defined below.
	RecordSets []RecordSetsObservation `json:"recordSets,omitempty" tf:"record_sets,omitempty"`
}

type PrivateDNSZoneConfigsParameters struct {
}

type PrivateDNSZoneGroupInitParameters struct {

	// Specifies the Name of the Private DNS Zone Group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the list of Private DNS Zones to include within the private_dns_zone_group.
	// +crossplane:generate:reference:type=PrivateDNSZone
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	PrivateDNSZoneIds []*string `json:"privateDnsZoneIds,omitempty" tf:"private_dns_zone_ids,omitempty"`

	// References to PrivateDNSZone to populate privateDnsZoneIds.
	// +kubebuilder:validation:Optional
	PrivateDNSZoneIdsRefs []v1.Reference `json:"privateDnsZoneIdsRefs,omitempty" tf:"-"`

	// Selector for a list of PrivateDNSZone to populate privateDnsZoneIds.
	// +kubebuilder:validation:Optional
	PrivateDNSZoneIdsSelector *v1.Selector `json:"privateDnsZoneIdsSelector,omitempty" tf:"-"`
}

type PrivateDNSZoneGroupObservation struct {

	// The ID of the Private DNS Zone Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the Name of the Private DNS Zone Group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the list of Private DNS Zones to include within the private_dns_zone_group.
	PrivateDNSZoneIds []*string `json:"privateDnsZoneIds,omitempty" tf:"private_dns_zone_ids,omitempty"`
}

type PrivateDNSZoneGroupParameters struct {

	// Specifies the Name of the Private DNS Zone Group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the list of Private DNS Zones to include within the private_dns_zone_group.
	// +crossplane:generate:reference:type=PrivateDNSZone
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PrivateDNSZoneIds []*string `json:"privateDnsZoneIds,omitempty" tf:"private_dns_zone_ids,omitempty"`

	// References to PrivateDNSZone to populate privateDnsZoneIds.
	// +kubebuilder:validation:Optional
	PrivateDNSZoneIdsRefs []v1.Reference `json:"privateDnsZoneIdsRefs,omitempty" tf:"-"`

	// Selector for a list of PrivateDNSZone to populate privateDnsZoneIds.
	// +kubebuilder:validation:Optional
	PrivateDNSZoneIdsSelector *v1.Selector `json:"privateDnsZoneIdsSelector,omitempty" tf:"-"`
}

type PrivateEndpointIPConfigurationInitParameters struct {

	// Specifies the member name this IP address applies to. If it is not specified, it will use the value of subresource_name. Changing this forces a new resource to be created.
	MemberName *string `json:"memberName,omitempty" tf:"member_name,omitempty"`

	// Specifies the Name of the IP Configuration. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the static IP address within the private endpoint's subnet to be used. Changing this forces a new resource to be created.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// Specifies the subresource this IP address applies to. subresource_names corresponds to group_id. Changing this forces a new resource to be created.
	SubresourceName *string `json:"subresourceName,omitempty" tf:"subresource_name,omitempty"`
}

type PrivateEndpointIPConfigurationObservation struct {

	// Specifies the member name this IP address applies to. If it is not specified, it will use the value of subresource_name. Changing this forces a new resource to be created.
	MemberName *string `json:"memberName,omitempty" tf:"member_name,omitempty"`

	// Specifies the Name of the IP Configuration. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the static IP address within the private endpoint's subnet to be used. Changing this forces a new resource to be created.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// Specifies the subresource this IP address applies to. subresource_names corresponds to group_id. Changing this forces a new resource to be created.
	SubresourceName *string `json:"subresourceName,omitempty" tf:"subresource_name,omitempty"`
}

type PrivateEndpointIPConfigurationParameters struct {

	// Specifies the member name this IP address applies to. If it is not specified, it will use the value of subresource_name. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	MemberName *string `json:"memberName,omitempty" tf:"member_name,omitempty"`

	// Specifies the Name of the IP Configuration. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the static IP address within the private endpoint's subnet to be used. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PrivateIPAddress *string `json:"privateIpAddress" tf:"private_ip_address,omitempty"`

	// Specifies the subresource this IP address applies to. subresource_names corresponds to group_id. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SubresourceName *string `json:"subresourceName,omitempty" tf:"subresource_name,omitempty"`
}

type PrivateEndpointInitParameters struct {

	// The custom name of the network interface attached to the private endpoint. Changing this forces a new resource to be created.
	CustomNetworkInterfaceName *string `json:"customNetworkInterfaceName,omitempty" tf:"custom_network_interface_name,omitempty"`

	// One or more ip_configuration blocks as defined below. This allows a static IP address to be set for this Private Endpoint, otherwise an address is dynamically allocated from the Subnet.
	IPConfiguration []PrivateEndpointIPConfigurationInitParameters `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A private_dns_zone_group block as defined below.
	PrivateDNSZoneGroup []PrivateDNSZoneGroupInitParameters `json:"privateDnsZoneGroup,omitempty" tf:"private_dns_zone_group,omitempty"`

	// A private_service_connection block as defined below.
	PrivateServiceConnection []PrivateServiceConnectionInitParameters `json:"privateServiceConnection,omitempty" tf:"private_service_connection,omitempty"`

	// The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateEndpointNetworkInterfaceInitParameters struct {
}

type PrivateEndpointNetworkInterfaceObservation struct {

	// The ID of the network interface associated with the private_endpoint.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the network interface associated with the private_endpoint.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type PrivateEndpointNetworkInterfaceParameters struct {
}

type PrivateEndpointObservation struct {

	// A custom_dns_configs block as defined below.
	CustomDNSConfigs []CustomDNSConfigsObservation `json:"customDnsConfigs,omitempty" tf:"custom_dns_configs,omitempty"`

	// The custom name of the network interface attached to the private endpoint. Changing this forces a new resource to be created.
	CustomNetworkInterfaceName *string `json:"customNetworkInterfaceName,omitempty" tf:"custom_network_interface_name,omitempty"`

	// The ID of the Private Endpoint.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more ip_configuration blocks as defined below. This allows a static IP address to be set for this Private Endpoint, otherwise an address is dynamically allocated from the Subnet.
	IPConfiguration []PrivateEndpointIPConfigurationObservation `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A network_interface block as defined below.
	NetworkInterface []PrivateEndpointNetworkInterfaceObservation `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// A private_dns_zone_configs block as defined below.
	PrivateDNSZoneConfigs []PrivateDNSZoneConfigsObservation `json:"privateDnsZoneConfigs,omitempty" tf:"private_dns_zone_configs,omitempty"`

	// A private_dns_zone_group block as defined below.
	PrivateDNSZoneGroup []PrivateDNSZoneGroupObservation `json:"privateDnsZoneGroup,omitempty" tf:"private_dns_zone_group,omitempty"`

	// A private_service_connection block as defined below.
	PrivateServiceConnection []PrivateServiceConnectionObservation `json:"privateServiceConnection,omitempty" tf:"private_service_connection,omitempty"`

	// Specifies the Name of the Resource Group within which the Private Endpoint should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateEndpointParameters struct {

	// The custom name of the network interface attached to the private endpoint. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	CustomNetworkInterfaceName *string `json:"customNetworkInterfaceName,omitempty" tf:"custom_network_interface_name,omitempty"`

	// One or more ip_configuration blocks as defined below. This allows a static IP address to be set for this Private Endpoint, otherwise an address is dynamically allocated from the Subnet.
	// +kubebuilder:validation:Optional
	IPConfiguration []PrivateEndpointIPConfigurationParameters `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A private_dns_zone_group block as defined below.
	// +kubebuilder:validation:Optional
	PrivateDNSZoneGroup []PrivateDNSZoneGroupParameters `json:"privateDnsZoneGroup,omitempty" tf:"private_dns_zone_group,omitempty"`

	// A private_service_connection block as defined below.
	// +kubebuilder:validation:Optional
	PrivateServiceConnection []PrivateServiceConnectionParameters `json:"privateServiceConnection,omitempty" tf:"private_service_connection,omitempty"`

	// Specifies the Name of the Resource Group within which the Private Endpoint should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateServiceConnectionInitParameters struct {

	// Does the Private Endpoint require Manual Approval from the remote resource owner? Changing this forces a new resource to be created.
	IsManualConnection *bool `json:"isManualConnection,omitempty" tf:"is_manual_connection,omitempty"`

	// Specifies the Name of the Private Service Connection. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Service Alias of the Private Link Enabled Remote Resource which this Private Endpoint should be connected to. One of private_connection_resource_id or private_connection_resource_alias must be specified. Changing this forces a new resource to be created.
	PrivateConnectionResourceAlias *string `json:"privateConnectionResourceAlias,omitempty" tf:"private_connection_resource_alias,omitempty"`

	// The ID of the Private Link Enabled Remote Resource which this Private Endpoint should be connected to. One of private_connection_resource_id or private_connection_resource_alias must be specified. Changing this forces a new resource to be created. For a web app or function app slot, the parent web app should be used in this field instead of a reference to the slot itself.
	PrivateConnectionResourceID *string `json:"privateConnectionResourceId,omitempty" tf:"private_connection_resource_id,omitempty"`

	// A message passed to the owner of the remote resource when the private endpoint attempts to establish the connection to the remote resource. The request message can be a maximum of 140 characters in length. Only valid if is_manual_connection is set to true.
	RequestMessage *string `json:"requestMessage,omitempty" tf:"request_message,omitempty"`

	// A list of subresource names which the Private Endpoint is able to connect to. subresource_names corresponds to group_id. Possible values are detailed in the product documentation in the Subresources column. Changing this forces a new resource to be created.
	SubresourceNames []*string `json:"subresourceNames,omitempty" tf:"subresource_names,omitempty"`
}

type PrivateServiceConnectionObservation struct {

	// Does the Private Endpoint require Manual Approval from the remote resource owner? Changing this forces a new resource to be created.
	IsManualConnection *bool `json:"isManualConnection,omitempty" tf:"is_manual_connection,omitempty"`

	// Specifies the Name of the Private Service Connection. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Service Alias of the Private Link Enabled Remote Resource which this Private Endpoint should be connected to. One of private_connection_resource_id or private_connection_resource_alias must be specified. Changing this forces a new resource to be created.
	PrivateConnectionResourceAlias *string `json:"privateConnectionResourceAlias,omitempty" tf:"private_connection_resource_alias,omitempty"`

	// The ID of the Private Link Enabled Remote Resource which this Private Endpoint should be connected to. One of private_connection_resource_id or private_connection_resource_alias must be specified. Changing this forces a new resource to be created. For a web app or function app slot, the parent web app should be used in this field instead of a reference to the slot itself.
	PrivateConnectionResourceID *string `json:"privateConnectionResourceId,omitempty" tf:"private_connection_resource_id,omitempty"`

	// (Computed) The private IP address associated with the private endpoint, note that you will have a private IP address assigned to the private endpoint even if the connection request was Rejected.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// A message passed to the owner of the remote resource when the private endpoint attempts to establish the connection to the remote resource. The request message can be a maximum of 140 characters in length. Only valid if is_manual_connection is set to true.
	RequestMessage *string `json:"requestMessage,omitempty" tf:"request_message,omitempty"`

	// A list of subresource names which the Private Endpoint is able to connect to. subresource_names corresponds to group_id. Possible values are detailed in the product documentation in the Subresources column. Changing this forces a new resource to be created.
	SubresourceNames []*string `json:"subresourceNames,omitempty" tf:"subresource_names,omitempty"`
}

type PrivateServiceConnectionParameters struct {

	// Does the Private Endpoint require Manual Approval from the remote resource owner? Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	IsManualConnection *bool `json:"isManualConnection" tf:"is_manual_connection,omitempty"`

	// Specifies the Name of the Private Service Connection. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The Service Alias of the Private Link Enabled Remote Resource which this Private Endpoint should be connected to. One of private_connection_resource_id or private_connection_resource_alias must be specified. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PrivateConnectionResourceAlias *string `json:"privateConnectionResourceAlias,omitempty" tf:"private_connection_resource_alias,omitempty"`

	// The ID of the Private Link Enabled Remote Resource which this Private Endpoint should be connected to. One of private_connection_resource_id or private_connection_resource_alias must be specified. Changing this forces a new resource to be created. For a web app or function app slot, the parent web app should be used in this field instead of a reference to the slot itself.
	// +kubebuilder:validation:Optional
	PrivateConnectionResourceID *string `json:"privateConnectionResourceId,omitempty" tf:"private_connection_resource_id,omitempty"`

	// A message passed to the owner of the remote resource when the private endpoint attempts to establish the connection to the remote resource. The request message can be a maximum of 140 characters in length. Only valid if is_manual_connection is set to true.
	// +kubebuilder:validation:Optional
	RequestMessage *string `json:"requestMessage,omitempty" tf:"request_message,omitempty"`

	// A list of subresource names which the Private Endpoint is able to connect to. subresource_names corresponds to group_id. Possible values are detailed in the product documentation in the Subresources column. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SubresourceNames []*string `json:"subresourceNames,omitempty" tf:"subresource_names,omitempty"`
}

type RecordSetsInitParameters struct {
}

type RecordSetsObservation struct {

	// The fully qualified domain name to the private_dns_zone.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// A list of all IP Addresses that map to the private_dns_zone fqdn.
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	// The name of the Private DNS Zone that the config belongs to.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The time to live for each connection to the private_dns_zone.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The type of DNS record.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RecordSetsParameters struct {
}

// PrivateEndpointSpec defines the desired state of PrivateEndpoint
type PrivateEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateEndpointParameters `json:"forProvider"`
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
	InitProvider PrivateEndpointInitParameters `json:"initProvider,omitempty"`
}

// PrivateEndpointStatus defines the observed state of PrivateEndpoint.
type PrivateEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PrivateEndpoint is the Schema for the PrivateEndpoints API. Manages a Private Endpoint.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PrivateEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.privateServiceConnection) || (has(self.initProvider) && has(self.initProvider.privateServiceConnection))",message="spec.forProvider.privateServiceConnection is a required parameter"
	Spec   PrivateEndpointSpec   `json:"spec"`
	Status PrivateEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateEndpointList contains a list of PrivateEndpoints
type PrivateEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateEndpoint `json:"items"`
}

// Repository type metadata.
var (
	PrivateEndpoint_Kind             = "PrivateEndpoint"
	PrivateEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateEndpoint_Kind}.String()
	PrivateEndpoint_KindAPIVersion   = PrivateEndpoint_Kind + "." + CRDGroupVersion.String()
	PrivateEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(PrivateEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateEndpoint{}, &PrivateEndpointList{})
}
