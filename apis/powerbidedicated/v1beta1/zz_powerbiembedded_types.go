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

type PowerBIEmbeddedInitParameters struct {

	// A set of administrator user identities, which manages the Power BI Embedded and must be a member user or a service principal in your AAD tenant.
	// +listType=set
	Administrators []*string `json:"administrators,omitempty" tf:"administrators,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Sets the PowerBI Embedded's mode. Possible values include: Gen1, Gen2. Defaults to Gen1. Changing this forces a new resource to be created.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Sets the PowerBI Embedded's pricing level's SKU. Possible values include: A1, A2, A3, A4, A5, A6.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PowerBIEmbeddedObservation struct {

	// A set of administrator user identities, which manages the Power BI Embedded and must be a member user or a service principal in your AAD tenant.
	// +listType=set
	Administrators []*string `json:"administrators,omitempty" tf:"administrators,omitempty"`

	// The ID of the PowerBI Embedded.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Sets the PowerBI Embedded's mode. Possible values include: Gen1, Gen2. Defaults to Gen1. Changing this forces a new resource to be created.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the Resource Group where the PowerBI Embedded should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Sets the PowerBI Embedded's pricing level's SKU. Possible values include: A1, A2, A3, A4, A5, A6.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PowerBIEmbeddedParameters struct {

	// A set of administrator user identities, which manages the Power BI Embedded and must be a member user or a service principal in your AAD tenant.
	// +kubebuilder:validation:Optional
	// +listType=set
	Administrators []*string `json:"administrators,omitempty" tf:"administrators,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Sets the PowerBI Embedded's mode. Possible values include: Gen1, Gen2. Defaults to Gen1. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the Resource Group where the PowerBI Embedded should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Sets the PowerBI Embedded's pricing level's SKU. Possible values include: A1, A2, A3, A4, A5, A6.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// PowerBIEmbeddedSpec defines the desired state of PowerBIEmbedded
type PowerBIEmbeddedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PowerBIEmbeddedParameters `json:"forProvider"`
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
	InitProvider PowerBIEmbeddedInitParameters `json:"initProvider,omitempty"`
}

// PowerBIEmbeddedStatus defines the observed state of PowerBIEmbedded.
type PowerBIEmbeddedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PowerBIEmbeddedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PowerBIEmbedded is the Schema for the PowerBIEmbeddeds API. Manages a PowerBI Embedded.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PowerBIEmbedded struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.administrators) || (has(self.initProvider) && has(self.initProvider.administrators))",message="spec.forProvider.administrators is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuName) || (has(self.initProvider) && has(self.initProvider.skuName))",message="spec.forProvider.skuName is a required parameter"
	Spec   PowerBIEmbeddedSpec   `json:"spec"`
	Status PowerBIEmbeddedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PowerBIEmbeddedList contains a list of PowerBIEmbeddeds
type PowerBIEmbeddedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PowerBIEmbedded `json:"items"`
}

// Repository type metadata.
var (
	PowerBIEmbedded_Kind             = "PowerBIEmbedded"
	PowerBIEmbedded_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PowerBIEmbedded_Kind}.String()
	PowerBIEmbedded_KindAPIVersion   = PowerBIEmbedded_Kind + "." + CRDGroupVersion.String()
	PowerBIEmbedded_GroupVersionKind = CRDGroupVersion.WithKind(PowerBIEmbedded_Kind)
)

func init() {
	SchemeBuilder.Register(&PowerBIEmbedded{}, &PowerBIEmbeddedList{})
}
