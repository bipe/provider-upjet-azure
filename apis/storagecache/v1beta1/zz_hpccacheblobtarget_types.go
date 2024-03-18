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

type HPCCacheBlobTargetInitParameters struct {

	// The name of the access policy applied to this target. Defaults to default.
	AccessPolicyName *string `json:"accessPolicyName,omitempty" tf:"access_policy_name,omitempty"`

	// The name HPC Cache, which the HPC Cache Blob Target will be added to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storagecache/v1beta1.HPCCache
	CacheName *string `json:"cacheName,omitempty" tf:"cache_name,omitempty"`

	// Reference to a HPCCache in storagecache to populate cacheName.
	// +kubebuilder:validation:Optional
	CacheNameRef *v1.Reference `json:"cacheNameRef,omitempty" tf:"-"`

	// Selector for a HPCCache in storagecache to populate cacheName.
	// +kubebuilder:validation:Optional
	CacheNameSelector *v1.Selector `json:"cacheNameSelector,omitempty" tf:"-"`

	// The client-facing file path of the HPC Cache Blob Target.
	NamespacePath *string `json:"namespacePath,omitempty" tf:"namespace_path,omitempty"`

	// The Resource Manager ID of the Storage Container used as the HPC Cache Blob Target. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Container
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("resource_manager_id",true)
	StorageContainerID *string `json:"storageContainerId,omitempty" tf:"storage_container_id,omitempty"`

	// Reference to a Container in storage to populate storageContainerId.
	// +kubebuilder:validation:Optional
	StorageContainerIDRef *v1.Reference `json:"storageContainerIdRef,omitempty" tf:"-"`

	// Selector for a Container in storage to populate storageContainerId.
	// +kubebuilder:validation:Optional
	StorageContainerIDSelector *v1.Selector `json:"storageContainerIdSelector,omitempty" tf:"-"`
}

type HPCCacheBlobTargetObservation struct {

	// The name of the access policy applied to this target. Defaults to default.
	AccessPolicyName *string `json:"accessPolicyName,omitempty" tf:"access_policy_name,omitempty"`

	// The name HPC Cache, which the HPC Cache Blob Target will be added to. Changing this forces a new resource to be created.
	CacheName *string `json:"cacheName,omitempty" tf:"cache_name,omitempty"`

	// The ID of the HPC Cache Blob Target.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The client-facing file path of the HPC Cache Blob Target.
	NamespacePath *string `json:"namespacePath,omitempty" tf:"namespace_path,omitempty"`

	// The name of the Resource Group in which to create the HPC Cache Blob Target. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The Resource Manager ID of the Storage Container used as the HPC Cache Blob Target. Changing this forces a new resource to be created.
	StorageContainerID *string `json:"storageContainerId,omitempty" tf:"storage_container_id,omitempty"`
}

type HPCCacheBlobTargetParameters struct {

	// The name of the access policy applied to this target. Defaults to default.
	// +kubebuilder:validation:Optional
	AccessPolicyName *string `json:"accessPolicyName,omitempty" tf:"access_policy_name,omitempty"`

	// The name HPC Cache, which the HPC Cache Blob Target will be added to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storagecache/v1beta1.HPCCache
	// +kubebuilder:validation:Optional
	CacheName *string `json:"cacheName,omitempty" tf:"cache_name,omitempty"`

	// Reference to a HPCCache in storagecache to populate cacheName.
	// +kubebuilder:validation:Optional
	CacheNameRef *v1.Reference `json:"cacheNameRef,omitempty" tf:"-"`

	// Selector for a HPCCache in storagecache to populate cacheName.
	// +kubebuilder:validation:Optional
	CacheNameSelector *v1.Selector `json:"cacheNameSelector,omitempty" tf:"-"`

	// The client-facing file path of the HPC Cache Blob Target.
	// +kubebuilder:validation:Optional
	NamespacePath *string `json:"namespacePath,omitempty" tf:"namespace_path,omitempty"`

	// The name of the Resource Group in which to create the HPC Cache Blob Target. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Resource Manager ID of the Storage Container used as the HPC Cache Blob Target. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Container
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("resource_manager_id",true)
	// +kubebuilder:validation:Optional
	StorageContainerID *string `json:"storageContainerId,omitempty" tf:"storage_container_id,omitempty"`

	// Reference to a Container in storage to populate storageContainerId.
	// +kubebuilder:validation:Optional
	StorageContainerIDRef *v1.Reference `json:"storageContainerIdRef,omitempty" tf:"-"`

	// Selector for a Container in storage to populate storageContainerId.
	// +kubebuilder:validation:Optional
	StorageContainerIDSelector *v1.Selector `json:"storageContainerIdSelector,omitempty" tf:"-"`
}

// HPCCacheBlobTargetSpec defines the desired state of HPCCacheBlobTarget
type HPCCacheBlobTargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HPCCacheBlobTargetParameters `json:"forProvider"`
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
	InitProvider HPCCacheBlobTargetInitParameters `json:"initProvider,omitempty"`
}

// HPCCacheBlobTargetStatus defines the observed state of HPCCacheBlobTarget.
type HPCCacheBlobTargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HPCCacheBlobTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HPCCacheBlobTarget is the Schema for the HPCCacheBlobTargets API. Manages a Blob Target within a HPC Cache.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HPCCacheBlobTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.namespacePath) || (has(self.initProvider) && has(self.initProvider.namespacePath))",message="spec.forProvider.namespacePath is a required parameter"
	Spec   HPCCacheBlobTargetSpec   `json:"spec"`
	Status HPCCacheBlobTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HPCCacheBlobTargetList contains a list of HPCCacheBlobTargets
type HPCCacheBlobTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HPCCacheBlobTarget `json:"items"`
}

// Repository type metadata.
var (
	HPCCacheBlobTarget_Kind             = "HPCCacheBlobTarget"
	HPCCacheBlobTarget_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HPCCacheBlobTarget_Kind}.String()
	HPCCacheBlobTarget_KindAPIVersion   = HPCCacheBlobTarget_Kind + "." + CRDGroupVersion.String()
	HPCCacheBlobTarget_GroupVersionKind = CRDGroupVersion.WithKind(HPCCacheBlobTarget_Kind)
)

func init() {
	SchemeBuilder.Register(&HPCCacheBlobTarget{}, &HPCCacheBlobTargetList{})
}
