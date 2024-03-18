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

type StorageSyncInitParameters struct {

	// Incoming traffic policy. Possible values are AllowAllTraffic and AllowVirtualNetworksOnly. Defaults to AllowAllTraffic.
	IncomingTrafficPolicy *string `json:"incomingTrafficPolicy,omitempty" tf:"incoming_traffic_policy,omitempty"`

	// The Azure Region where the Storage Sync should exist. Changing this forces a new Storage Sync to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags which should be assigned to the Storage Sync.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StorageSyncObservation struct {

	// The ID of the Storage Sync.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Incoming traffic policy. Possible values are AllowAllTraffic and AllowVirtualNetworksOnly. Defaults to AllowAllTraffic.
	IncomingTrafficPolicy *string `json:"incomingTrafficPolicy,omitempty" tf:"incoming_traffic_policy,omitempty"`

	// The Azure Region where the Storage Sync should exist. Changing this forces a new Storage Sync to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Resource Group where the Storage Sync should exist. Changing this forces a new Storage Sync to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags which should be assigned to the Storage Sync.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StorageSyncParameters struct {

	// Incoming traffic policy. Possible values are AllowAllTraffic and AllowVirtualNetworksOnly. Defaults to AllowAllTraffic.
	// +kubebuilder:validation:Optional
	IncomingTrafficPolicy *string `json:"incomingTrafficPolicy,omitempty" tf:"incoming_traffic_policy,omitempty"`

	// The Azure Region where the Storage Sync should exist. Changing this forces a new Storage Sync to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Resource Group where the Storage Sync should exist. Changing this forces a new Storage Sync to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Storage Sync.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// StorageSyncSpec defines the desired state of StorageSync
type StorageSyncSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StorageSyncParameters `json:"forProvider"`
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
	InitProvider StorageSyncInitParameters `json:"initProvider,omitempty"`
}

// StorageSyncStatus defines the observed state of StorageSync.
type StorageSyncStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StorageSyncObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// StorageSync is the Schema for the StorageSyncs API. Manages a Storage Sync.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type StorageSync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   StorageSyncSpec   `json:"spec"`
	Status StorageSyncStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageSyncList contains a list of StorageSyncs
type StorageSyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageSync `json:"items"`
}

// Repository type metadata.
var (
	StorageSync_Kind             = "StorageSync"
	StorageSync_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StorageSync_Kind}.String()
	StorageSync_KindAPIVersion   = StorageSync_Kind + "." + CRDGroupVersion.String()
	StorageSync_GroupVersionKind = CRDGroupVersion.WithKind(StorageSync_Kind)
)

func init() {
	SchemeBuilder.Register(&StorageSync{}, &StorageSyncList{})
}
