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

type NamespaceNetworkRuleSetInitParameters struct {

	// Specifies the default action for the ServiceBus Namespace Network Rule Set. Possible values are Allow and Deny. Defaults to Allow.
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// One or more IP Addresses, or CIDR Blocks which should be able to access the ServiceBus Namespace.
	// +listType=set
	IPRules []*string `json:"ipRules,omitempty" tf:"ip_rules,omitempty"`

	// Specifies the ServiceBus Namespace ID to which to attach the ServiceBus Namespace Network Rule Set. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/servicebus/v1beta1.ServiceBusNamespace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Reference to a ServiceBusNamespace in servicebus to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDRef *v1.Reference `json:"namespaceIdRef,omitempty" tf:"-"`

	// Selector for a ServiceBusNamespace in servicebus to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDSelector *v1.Selector `json:"namespaceIdSelector,omitempty" tf:"-"`

	// One or more network_rules blocks as defined below.
	NetworkRules []NamespaceNetworkRuleSetNetworkRulesInitParameters `json:"networkRules,omitempty" tf:"network_rules,omitempty"`

	// Whether to allow traffic over public network. Possible values are true and false. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// If True, then Azure Services that are known and trusted for this resource type are allowed to bypass firewall configuration. See Trusted Microsoft Services
	TrustedServicesAllowed *bool `json:"trustedServicesAllowed,omitempty" tf:"trusted_services_allowed,omitempty"`
}

type NamespaceNetworkRuleSetNetworkRulesInitParameters struct {

	// Should the ServiceBus Namespace Network Rule Set ignore missing Virtual Network Service Endpoint option in the Subnet? Defaults to false.
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint,omitempty"`

	// The Subnet ID which should be able to access this ServiceBus Namespace.
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

type NamespaceNetworkRuleSetNetworkRulesObservation struct {

	// Should the ServiceBus Namespace Network Rule Set ignore missing Virtual Network Service Endpoint option in the Subnet? Defaults to false.
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint,omitempty"`

	// The Subnet ID which should be able to access this ServiceBus Namespace.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type NamespaceNetworkRuleSetNetworkRulesParameters struct {

	// Should the ServiceBus Namespace Network Rule Set ignore missing Virtual Network Service Endpoint option in the Subnet? Defaults to false.
	// +kubebuilder:validation:Optional
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint,omitempty"`

	// The Subnet ID which should be able to access this ServiceBus Namespace.
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

type NamespaceNetworkRuleSetObservation struct {

	// Specifies the default action for the ServiceBus Namespace Network Rule Set. Possible values are Allow and Deny. Defaults to Allow.
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// The ID of the ServiceBus Namespace Network Rule Set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more IP Addresses, or CIDR Blocks which should be able to access the ServiceBus Namespace.
	// +listType=set
	IPRules []*string `json:"ipRules,omitempty" tf:"ip_rules,omitempty"`

	// Specifies the ServiceBus Namespace ID to which to attach the ServiceBus Namespace Network Rule Set. Changing this forces a new resource to be created.
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// One or more network_rules blocks as defined below.
	NetworkRules []NamespaceNetworkRuleSetNetworkRulesObservation `json:"networkRules,omitempty" tf:"network_rules,omitempty"`

	// Whether to allow traffic over public network. Possible values are true and false. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// If True, then Azure Services that are known and trusted for this resource type are allowed to bypass firewall configuration. See Trusted Microsoft Services
	TrustedServicesAllowed *bool `json:"trustedServicesAllowed,omitempty" tf:"trusted_services_allowed,omitempty"`
}

type NamespaceNetworkRuleSetParameters struct {

	// Specifies the default action for the ServiceBus Namespace Network Rule Set. Possible values are Allow and Deny. Defaults to Allow.
	// +kubebuilder:validation:Optional
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// One or more IP Addresses, or CIDR Blocks which should be able to access the ServiceBus Namespace.
	// +kubebuilder:validation:Optional
	// +listType=set
	IPRules []*string `json:"ipRules,omitempty" tf:"ip_rules,omitempty"`

	// Specifies the ServiceBus Namespace ID to which to attach the ServiceBus Namespace Network Rule Set. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/servicebus/v1beta1.ServiceBusNamespace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Reference to a ServiceBusNamespace in servicebus to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDRef *v1.Reference `json:"namespaceIdRef,omitempty" tf:"-"`

	// Selector for a ServiceBusNamespace in servicebus to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDSelector *v1.Selector `json:"namespaceIdSelector,omitempty" tf:"-"`

	// One or more network_rules blocks as defined below.
	// +kubebuilder:validation:Optional
	NetworkRules []NamespaceNetworkRuleSetNetworkRulesParameters `json:"networkRules,omitempty" tf:"network_rules,omitempty"`

	// Whether to allow traffic over public network. Possible values are true and false. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// If True, then Azure Services that are known and trusted for this resource type are allowed to bypass firewall configuration. See Trusted Microsoft Services
	// +kubebuilder:validation:Optional
	TrustedServicesAllowed *bool `json:"trustedServicesAllowed,omitempty" tf:"trusted_services_allowed,omitempty"`
}

// NamespaceNetworkRuleSetSpec defines the desired state of NamespaceNetworkRuleSet
type NamespaceNetworkRuleSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NamespaceNetworkRuleSetParameters `json:"forProvider"`
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
	InitProvider NamespaceNetworkRuleSetInitParameters `json:"initProvider,omitempty"`
}

// NamespaceNetworkRuleSetStatus defines the observed state of NamespaceNetworkRuleSet.
type NamespaceNetworkRuleSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NamespaceNetworkRuleSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NamespaceNetworkRuleSet is the Schema for the NamespaceNetworkRuleSets API. Manages a ServiceBus Namespace Network Rule Set.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NamespaceNetworkRuleSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespaceNetworkRuleSetSpec   `json:"spec"`
	Status            NamespaceNetworkRuleSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NamespaceNetworkRuleSetList contains a list of NamespaceNetworkRuleSets
type NamespaceNetworkRuleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespaceNetworkRuleSet `json:"items"`
}

// Repository type metadata.
var (
	NamespaceNetworkRuleSet_Kind             = "NamespaceNetworkRuleSet"
	NamespaceNetworkRuleSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NamespaceNetworkRuleSet_Kind}.String()
	NamespaceNetworkRuleSet_KindAPIVersion   = NamespaceNetworkRuleSet_Kind + "." + CRDGroupVersion.String()
	NamespaceNetworkRuleSet_GroupVersionKind = CRDGroupVersion.WithKind(NamespaceNetworkRuleSet_Kind)
)

func init() {
	SchemeBuilder.Register(&NamespaceNetworkRuleSet{}, &NamespaceNetworkRuleSetList{})
}
