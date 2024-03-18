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

type NotificationHubNamespaceInitParameters struct {

	// Is this Notification Hub Namespace enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The Azure Region in which this Notification Hub Namespace should be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Type of Namespace - possible values are Messaging or NotificationHub.
	NamespaceType *string `json:"namespaceType,omitempty" tf:"namespace_type,omitempty"`

	// The name of the SKU to use for this Notification Hub Namespace. Possible values are Free, Basic or Standard.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NotificationHubNamespaceObservation struct {

	// Is this Notification Hub Namespace enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the Notification Hub Namespace.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region in which this Notification Hub Namespace should be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Type of Namespace - possible values are Messaging or NotificationHub.
	NamespaceType *string `json:"namespaceType,omitempty" tf:"namespace_type,omitempty"`

	// The name of the Resource Group in which the Notification Hub Namespace should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The ServiceBus Endpoint for this Notification Hub Namespace.
	ServiceBusEndpoint *string `json:"servicebusEndpoint,omitempty" tf:"servicebus_endpoint,omitempty"`

	// The name of the SKU to use for this Notification Hub Namespace. Possible values are Free, Basic or Standard.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NotificationHubNamespaceParameters struct {

	// Is this Notification Hub Namespace enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The Azure Region in which this Notification Hub Namespace should be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Type of Namespace - possible values are Messaging or NotificationHub.
	// +kubebuilder:validation:Optional
	NamespaceType *string `json:"namespaceType,omitempty" tf:"namespace_type,omitempty"`

	// The name of the Resource Group in which the Notification Hub Namespace should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the SKU to use for this Notification Hub Namespace. Possible values are Free, Basic or Standard.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// NotificationHubNamespaceSpec defines the desired state of NotificationHubNamespace
type NotificationHubNamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotificationHubNamespaceParameters `json:"forProvider"`
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
	InitProvider NotificationHubNamespaceInitParameters `json:"initProvider,omitempty"`
}

// NotificationHubNamespaceStatus defines the observed state of NotificationHubNamespace.
type NotificationHubNamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotificationHubNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NotificationHubNamespace is the Schema for the NotificationHubNamespaces API. Manages a Notification Hub Namespace.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NotificationHubNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.namespaceType) || (has(self.initProvider) && has(self.initProvider.namespaceType))",message="spec.forProvider.namespaceType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuName) || (has(self.initProvider) && has(self.initProvider.skuName))",message="spec.forProvider.skuName is a required parameter"
	Spec   NotificationHubNamespaceSpec   `json:"spec"`
	Status NotificationHubNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationHubNamespaceList contains a list of NotificationHubNamespaces
type NotificationHubNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotificationHubNamespace `json:"items"`
}

// Repository type metadata.
var (
	NotificationHubNamespace_Kind             = "NotificationHubNamespace"
	NotificationHubNamespace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NotificationHubNamespace_Kind}.String()
	NotificationHubNamespace_KindAPIVersion   = NotificationHubNamespace_Kind + "." + CRDGroupVersion.String()
	NotificationHubNamespace_GroupVersionKind = CRDGroupVersion.WithKind(NotificationHubNamespace_Kind)
)

func init() {
	SchemeBuilder.Register(&NotificationHubNamespace{}, &NotificationHubNamespaceList{})
}
