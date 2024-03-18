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

type LabInitParameters struct {

	// Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The type of storage used by the Dev Test Lab. Possible values are Standard and Premium. Defaults to Premium.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LabObservation struct {

	// The ID of the Storage Account used for Artifact Storage.
	ArtifactsStorageAccountID *string `json:"artifactsStorageAccountId,omitempty" tf:"artifacts_storage_account_id,omitempty"`

	// The ID of the Default Premium Storage Account for this Dev Test Lab.
	DefaultPremiumStorageAccountID *string `json:"defaultPremiumStorageAccountId,omitempty" tf:"default_premium_storage_account_id,omitempty"`

	// The ID of the Default Storage Account for this Dev Test Lab.
	DefaultStorageAccountID *string `json:"defaultStorageAccountId,omitempty" tf:"default_storage_account_id,omitempty"`

	// The ID of the Dev Test Lab.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Key used for this Dev Test Lab.
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Storage Account used for Storage of Premium Data Disk.
	PremiumDataDiskStorageAccountID *string `json:"premiumDataDiskStorageAccountId,omitempty" tf:"premium_data_disk_storage_account_id,omitempty"`

	// The name of the resource group under which the Dev Test Lab resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The type of storage used by the Dev Test Lab. Possible values are Standard and Premium. Defaults to Premium.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The unique immutable identifier of the Dev Test Lab.
	UniqueIdentifier *string `json:"uniqueIdentifier,omitempty" tf:"unique_identifier,omitempty"`
}

type LabParameters struct {

	// Specifies the supported Azure location where the Dev Test Lab should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group under which the Dev Test Lab resource has to be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The type of storage used by the Dev Test Lab. Possible values are Standard and Premium. Defaults to Premium.
	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LabSpec defines the desired state of Lab
type LabSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LabParameters `json:"forProvider"`
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
	InitProvider LabInitParameters `json:"initProvider,omitempty"`
}

// LabStatus defines the observed state of Lab.
type LabStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LabObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Lab is the Schema for the Labs API. Manages a Dev Test Lab.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Lab struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   LabSpec   `json:"spec"`
	Status LabStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LabList contains a list of Labs
type LabList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Lab `json:"items"`
}

// Repository type metadata.
var (
	Lab_Kind             = "Lab"
	Lab_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Lab_Kind}.String()
	Lab_KindAPIVersion   = Lab_Kind + "." + CRDGroupVersion.String()
	Lab_GroupVersionKind = CRDGroupVersion.WithKind(Lab_Kind)
)

func init() {
	SchemeBuilder.Register(&Lab{}, &LabList{})
}
