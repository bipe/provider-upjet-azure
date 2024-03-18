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

type ActionInitParameters struct {

	// Specifies the endpoint of the action.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Specifies the name of the action.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ActionObservation struct {

	// Specifies the endpoint of the action.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Specifies the name of the action.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ActionParameters struct {

	// Specifies the endpoint of the action.
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint" tf:"endpoint,omitempty"`

	// Specifies the name of the action.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type CustomProviderInitParameters struct {

	// Any number of action block as defined below. One of resource_type or action must be specified.
	Action []ActionInitParameters `json:"action,omitempty" tf:"action,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Any number of resource_type block as defined below. One of resource_type or action must be specified.
	ResourceType []ResourceTypeInitParameters `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Any number of validation block as defined below.
	Validation []ValidationInitParameters `json:"validation,omitempty" tf:"validation,omitempty"`
}

type CustomProviderObservation struct {

	// Any number of action block as defined below. One of resource_type or action must be specified.
	Action []ActionObservation `json:"action,omitempty" tf:"action,omitempty"`

	// The ID of the Custom Provider.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the Custom Provider. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Any number of resource_type block as defined below. One of resource_type or action must be specified.
	ResourceType []ResourceTypeObservation `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Any number of validation block as defined below.
	Validation []ValidationObservation `json:"validation,omitempty" tf:"validation,omitempty"`
}

type CustomProviderParameters struct {

	// Any number of action block as defined below. One of resource_type or action must be specified.
	// +kubebuilder:validation:Optional
	Action []ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the Custom Provider. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Any number of resource_type block as defined below. One of resource_type or action must be specified.
	// +kubebuilder:validation:Optional
	ResourceType []ResourceTypeParameters `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Any number of validation block as defined below.
	// +kubebuilder:validation:Optional
	Validation []ValidationParameters `json:"validation,omitempty" tf:"validation,omitempty"`
}

type ResourceTypeInitParameters struct {

	// Specifies the endpoint of the route definition.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Specifies the name of the route definition.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The routing type that is supported for the resource request. Valid values are Proxy and Proxy,Cache. Defaults to Proxy.
	RoutingType *string `json:"routingType,omitempty" tf:"routing_type,omitempty"`
}

type ResourceTypeObservation struct {

	// Specifies the endpoint of the route definition.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Specifies the name of the route definition.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The routing type that is supported for the resource request. Valid values are Proxy and Proxy,Cache. Defaults to Proxy.
	RoutingType *string `json:"routingType,omitempty" tf:"routing_type,omitempty"`
}

type ResourceTypeParameters struct {

	// Specifies the endpoint of the route definition.
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint" tf:"endpoint,omitempty"`

	// Specifies the name of the route definition.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The routing type that is supported for the resource request. Valid values are Proxy and Proxy,Cache. Defaults to Proxy.
	// +kubebuilder:validation:Optional
	RoutingType *string `json:"routingType,omitempty" tf:"routing_type,omitempty"`
}

type ValidationInitParameters struct {

	// The endpoint where the validation specification is located.
	Specification *string `json:"specification,omitempty" tf:"specification,omitempty"`
}

type ValidationObservation struct {

	// The endpoint where the validation specification is located.
	Specification *string `json:"specification,omitempty" tf:"specification,omitempty"`
}

type ValidationParameters struct {

	// The endpoint where the validation specification is located.
	// +kubebuilder:validation:Optional
	Specification *string `json:"specification" tf:"specification,omitempty"`
}

// CustomProviderSpec defines the desired state of CustomProvider
type CustomProviderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomProviderParameters `json:"forProvider"`
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
	InitProvider CustomProviderInitParameters `json:"initProvider,omitempty"`
}

// CustomProviderStatus defines the observed state of CustomProvider.
type CustomProviderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CustomProvider is the Schema for the CustomProviders API. Manages an Azure Custom Provider.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CustomProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   CustomProviderSpec   `json:"spec"`
	Status CustomProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomProviderList contains a list of CustomProviders
type CustomProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomProvider `json:"items"`
}

// Repository type metadata.
var (
	CustomProvider_Kind             = "CustomProvider"
	CustomProvider_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomProvider_Kind}.String()
	CustomProvider_KindAPIVersion   = CustomProvider_Kind + "." + CRDGroupVersion.String()
	CustomProvider_GroupVersionKind = CRDGroupVersion.WithKind(CustomProvider_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomProvider{}, &CustomProviderList{})
}
