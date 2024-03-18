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

type NetworkInterfaceIPConfigurationInitParameters struct {

	// The Frontend IP Configuration ID of a Gateway SKU Load Balancer.
	GatewayLoadBalancerFrontendIPConfigurationID *string `json:"gatewayLoadBalancerFrontendIpConfigurationId,omitempty" tf:"gateway_load_balancer_frontend_ip_configuration_id,omitempty"`

	// A name used for this IP Configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Is this the Primary IP Configuration? Must be true for the first ip_configuration when multiple are specified. Defaults to false.
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`

	// The Static IP Address which should be used.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The allocation method used for the Private IP Address. Possible values are Dynamic and Static.
	PrivateIPAddressAllocation *string `json:"privateIpAddressAllocation,omitempty" tf:"private_ip_address_allocation,omitempty"`

	// The IP Version to use. Possible values are IPv4 or IPv6. Defaults to IPv4.
	PrivateIPAddressVersion *string `json:"privateIpAddressVersion,omitempty" tf:"private_ip_address_version,omitempty"`

	// Reference to a Public IP Address to associate with this NIC
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.PublicIP
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to a PublicIP in network to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDRef *v1.Reference `json:"publicIpAddressIdRef,omitempty" tf:"-"`

	// Selector for a PublicIP in network to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDSelector *v1.Selector `json:"publicIpAddressIdSelector,omitempty" tf:"-"`

	// The ID of the Subnet where this Network Interface should be located in.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type NetworkInterfaceIPConfigurationObservation struct {

	// The Frontend IP Configuration ID of a Gateway SKU Load Balancer.
	GatewayLoadBalancerFrontendIPConfigurationID *string `json:"gatewayLoadBalancerFrontendIpConfigurationId,omitempty" tf:"gateway_load_balancer_frontend_ip_configuration_id,omitempty"`

	// A name used for this IP Configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Is this the Primary IP Configuration? Must be true for the first ip_configuration when multiple are specified. Defaults to false.
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`

	// The Static IP Address which should be used.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The allocation method used for the Private IP Address. Possible values are Dynamic and Static.
	PrivateIPAddressAllocation *string `json:"privateIpAddressAllocation,omitempty" tf:"private_ip_address_allocation,omitempty"`

	// The IP Version to use. Possible values are IPv4 or IPv6. Defaults to IPv4.
	PrivateIPAddressVersion *string `json:"privateIpAddressVersion,omitempty" tf:"private_ip_address_version,omitempty"`

	// Reference to a Public IP Address to associate with this NIC
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// The ID of the Subnet where this Network Interface should be located in.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type NetworkInterfaceIPConfigurationParameters struct {

	// The Frontend IP Configuration ID of a Gateway SKU Load Balancer.
	// +kubebuilder:validation:Optional
	GatewayLoadBalancerFrontendIPConfigurationID *string `json:"gatewayLoadBalancerFrontendIpConfigurationId,omitempty" tf:"gateway_load_balancer_frontend_ip_configuration_id,omitempty"`

	// A name used for this IP Configuration.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Is this the Primary IP Configuration? Must be true for the first ip_configuration when multiple are specified. Defaults to false.
	// +kubebuilder:validation:Optional
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`

	// The Static IP Address which should be used.
	// +kubebuilder:validation:Optional
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The allocation method used for the Private IP Address. Possible values are Dynamic and Static.
	// +kubebuilder:validation:Optional
	PrivateIPAddressAllocation *string `json:"privateIpAddressAllocation" tf:"private_ip_address_allocation,omitempty"`

	// The IP Version to use. Possible values are IPv4 or IPv6. Defaults to IPv4.
	// +kubebuilder:validation:Optional
	PrivateIPAddressVersion *string `json:"privateIpAddressVersion,omitempty" tf:"private_ip_address_version,omitempty"`

	// Reference to a Public IP Address to associate with this NIC
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.PublicIP
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to a PublicIP in network to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDRef *v1.Reference `json:"publicIpAddressIdRef,omitempty" tf:"-"`

	// Selector for a PublicIP in network to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDSelector *v1.Selector `json:"publicIpAddressIdSelector,omitempty" tf:"-"`

	// The ID of the Subnet where this Network Interface should be located in.
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
}

type NetworkInterfaceInitParameters struct {

	// Specifies the auxiliary mode used to enable network high-performance feature on Network Virtual Appliances (NVAs). This feature offers competitive performance in Connections Per Second (CPS) optimization, along with improvements to handling large amounts of simultaneous connections. Possible values are AcceleratedConnections, Floating, MaxConnections and None.
	AuxiliaryMode *string `json:"auxiliaryMode,omitempty" tf:"auxiliary_mode,omitempty"`

	// Specifies the SKU used for the network high-performance feature on Network Virtual Appliances (NVAs). Possible values are A8, A4, A1, A2 and None.
	AuxiliarySku *string `json:"auxiliarySku,omitempty" tf:"auxiliary_sku,omitempty"`

	// A list of IP Addresses defining the DNS Servers which should be used for this Network Interface.
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// Specifies the Edge Zone within the Azure Region where this Network Interface should exist. Changing this forces a new Network Interface to be created.
	EdgeZone *string `json:"edgeZone,omitempty" tf:"edge_zone,omitempty"`

	// Should Accelerated Networking be enabled? Defaults to false.
	EnableAcceleratedNetworking *bool `json:"enableAcceleratedNetworking,omitempty" tf:"enable_accelerated_networking,omitempty"`

	// Should IP Forwarding be enabled? Defaults to false.
	EnableIPForwarding *bool `json:"enableIpForwarding,omitempty" tf:"enable_ip_forwarding,omitempty"`

	// One or more ip_configuration blocks as defined below.
	IPConfiguration []NetworkInterfaceIPConfigurationInitParameters `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// The (relative) DNS Name used for internal communications between Virtual Machines in the same Virtual Network.
	InternalDNSNameLabel *string `json:"internalDnsNameLabel,omitempty" tf:"internal_dns_name_label,omitempty"`

	// The location where the Network Interface should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NetworkInterfaceObservation struct {

	// If the Virtual Machine using this Network Interface is part of an Availability Set, then this list will have the union of all DNS servers from all Network Interfaces that are part of the Availability Set.
	AppliedDNSServers []*string `json:"appliedDnsServers,omitempty" tf:"applied_dns_servers,omitempty"`

	// Specifies the auxiliary mode used to enable network high-performance feature on Network Virtual Appliances (NVAs). This feature offers competitive performance in Connections Per Second (CPS) optimization, along with improvements to handling large amounts of simultaneous connections. Possible values are AcceleratedConnections, Floating, MaxConnections and None.
	AuxiliaryMode *string `json:"auxiliaryMode,omitempty" tf:"auxiliary_mode,omitempty"`

	// Specifies the SKU used for the network high-performance feature on Network Virtual Appliances (NVAs). Possible values are A8, A4, A1, A2 and None.
	AuxiliarySku *string `json:"auxiliarySku,omitempty" tf:"auxiliary_sku,omitempty"`

	// A list of IP Addresses defining the DNS Servers which should be used for this Network Interface.
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// Specifies the Edge Zone within the Azure Region where this Network Interface should exist. Changing this forces a new Network Interface to be created.
	EdgeZone *string `json:"edgeZone,omitempty" tf:"edge_zone,omitempty"`

	// Should Accelerated Networking be enabled? Defaults to false.
	EnableAcceleratedNetworking *bool `json:"enableAcceleratedNetworking,omitempty" tf:"enable_accelerated_networking,omitempty"`

	// Should IP Forwarding be enabled? Defaults to false.
	EnableIPForwarding *bool `json:"enableIpForwarding,omitempty" tf:"enable_ip_forwarding,omitempty"`

	// The ID of the Network Interface.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more ip_configuration blocks as defined below.
	IPConfiguration []NetworkInterfaceIPConfigurationObservation `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// The (relative) DNS Name used for internal communications between Virtual Machines in the same Virtual Network.
	InternalDNSNameLabel *string `json:"internalDnsNameLabel,omitempty" tf:"internal_dns_name_label,omitempty"`

	// Even if internal_dns_name_label is not specified, a DNS entry is created for the primary NIC of the VM. This DNS name can be constructed by concatenating the VM name with the value of internal_domain_name_suffix.
	InternalDomainNameSuffix *string `json:"internalDomainNameSuffix,omitempty" tf:"internal_domain_name_suffix,omitempty"`

	// The location where the Network Interface should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Media Access Control (MAC) Address of the Network Interface.
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	// The first private IP address of the network interface.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The private IP addresses of the network interface.
	PrivateIPAddresses []*string `json:"privateIpAddresses,omitempty" tf:"private_ip_addresses,omitempty"`

	// The name of the Resource Group in which to create the Network Interface. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the Virtual Machine which this Network Interface is connected to.
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`
}

type NetworkInterfaceParameters struct {

	// Specifies the auxiliary mode used to enable network high-performance feature on Network Virtual Appliances (NVAs). This feature offers competitive performance in Connections Per Second (CPS) optimization, along with improvements to handling large amounts of simultaneous connections. Possible values are AcceleratedConnections, Floating, MaxConnections and None.
	// +kubebuilder:validation:Optional
	AuxiliaryMode *string `json:"auxiliaryMode,omitempty" tf:"auxiliary_mode,omitempty"`

	// Specifies the SKU used for the network high-performance feature on Network Virtual Appliances (NVAs). Possible values are A8, A4, A1, A2 and None.
	// +kubebuilder:validation:Optional
	AuxiliarySku *string `json:"auxiliarySku,omitempty" tf:"auxiliary_sku,omitempty"`

	// A list of IP Addresses defining the DNS Servers which should be used for this Network Interface.
	// +kubebuilder:validation:Optional
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// Specifies the Edge Zone within the Azure Region where this Network Interface should exist. Changing this forces a new Network Interface to be created.
	// +kubebuilder:validation:Optional
	EdgeZone *string `json:"edgeZone,omitempty" tf:"edge_zone,omitempty"`

	// Should Accelerated Networking be enabled? Defaults to false.
	// +kubebuilder:validation:Optional
	EnableAcceleratedNetworking *bool `json:"enableAcceleratedNetworking,omitempty" tf:"enable_accelerated_networking,omitempty"`

	// Should IP Forwarding be enabled? Defaults to false.
	// +kubebuilder:validation:Optional
	EnableIPForwarding *bool `json:"enableIpForwarding,omitempty" tf:"enable_ip_forwarding,omitempty"`

	// One or more ip_configuration blocks as defined below.
	// +kubebuilder:validation:Optional
	IPConfiguration []NetworkInterfaceIPConfigurationParameters `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// The (relative) DNS Name used for internal communications between Virtual Machines in the same Virtual Network.
	// +kubebuilder:validation:Optional
	InternalDNSNameLabel *string `json:"internalDnsNameLabel,omitempty" tf:"internal_dns_name_label,omitempty"`

	// The location where the Network Interface should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Resource Group in which to create the Network Interface. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// NetworkInterfaceSpec defines the desired state of NetworkInterface
type NetworkInterfaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkInterfaceParameters `json:"forProvider"`
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
	InitProvider NetworkInterfaceInitParameters `json:"initProvider,omitempty"`
}

// NetworkInterfaceStatus defines the observed state of NetworkInterface.
type NetworkInterfaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NetworkInterface is the Schema for the NetworkInterfaces API. Manages a Network Interface.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NetworkInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipConfiguration) || (has(self.initProvider) && has(self.initProvider.ipConfiguration))",message="spec.forProvider.ipConfiguration is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   NetworkInterfaceSpec   `json:"spec"`
	Status NetworkInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceList contains a list of NetworkInterfaces
type NetworkInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterface `json:"items"`
}

// Repository type metadata.
var (
	NetworkInterface_Kind             = "NetworkInterface"
	NetworkInterface_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkInterface_Kind}.String()
	NetworkInterface_KindAPIVersion   = NetworkInterface_Kind + "." + CRDGroupVersion.String()
	NetworkInterface_GroupVersionKind = CRDGroupVersion.WithKind(NetworkInterface_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkInterface{}, &NetworkInterfaceList{})
}
