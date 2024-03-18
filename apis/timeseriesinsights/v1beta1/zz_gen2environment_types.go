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

type Gen2EnvironmentInitParameters struct {

	// A list of property ids for the Azure IoT Time Series Insights Gen2 Environment. Changing this forces a new resource to be created.
	IDProperties []*string `json:"idProperties,omitempty" tf:"id_properties,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the SKU Name for this IoT Time Series Insights Gen2 Environment. Currently it supports only L1. For gen2, capacity cannot be specified. Changing this forces a new resource to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A storage block as defined below.
	Storage []StorageInitParameters `json:"storage,omitempty" tf:"storage,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query.
	WarmStoreDataRetentionTime *string `json:"warmStoreDataRetentionTime,omitempty" tf:"warm_store_data_retention_time,omitempty"`
}

type Gen2EnvironmentObservation struct {

	// The FQDN used to access the environment data.
	DataAccessFqdn *string `json:"dataAccessFqdn,omitempty" tf:"data_access_fqdn,omitempty"`

	// The ID of the IoT Time Series Insights Gen2 Environment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of property ids for the Azure IoT Time Series Insights Gen2 Environment. Changing this forces a new resource to be created.
	IDProperties []*string `json:"idProperties,omitempty" tf:"id_properties,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the Azure IoT Time Series Insights Gen2 Environment. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the SKU Name for this IoT Time Series Insights Gen2 Environment. Currently it supports only L1. For gen2, capacity cannot be specified. Changing this forces a new resource to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A storage block as defined below.
	Storage []StorageObservation `json:"storage,omitempty" tf:"storage,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query.
	WarmStoreDataRetentionTime *string `json:"warmStoreDataRetentionTime,omitempty" tf:"warm_store_data_retention_time,omitempty"`
}

type Gen2EnvironmentParameters struct {

	// A list of property ids for the Azure IoT Time Series Insights Gen2 Environment. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	IDProperties []*string `json:"idProperties,omitempty" tf:"id_properties,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the Azure IoT Time Series Insights Gen2 Environment. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the SKU Name for this IoT Time Series Insights Gen2 Environment. Currently it supports only L1. For gen2, capacity cannot be specified. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A storage block as defined below.
	// +kubebuilder:validation:Optional
	Storage []StorageParameters `json:"storage,omitempty" tf:"storage,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query.
	// +kubebuilder:validation:Optional
	WarmStoreDataRetentionTime *string `json:"warmStoreDataRetentionTime,omitempty" tf:"warm_store_data_retention_time,omitempty"`
}

type StorageInitParameters struct {

	// Name of storage account for Azure IoT Time Series Insights Gen2 Environment. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Account in storage to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`
}

type StorageObservation struct {

	// Name of storage account for Azure IoT Time Series Insights Gen2 Environment. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type StorageParameters struct {

	// Access key of storage account for Azure IoT Time Series Insights Gen2 Environment
	// +kubebuilder:validation:Required
	KeySecretRef v1.SecretKeySelector `json:"keySecretRef" tf:"-"`

	// Name of storage account for Azure IoT Time Series Insights Gen2 Environment. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Account in storage to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`
}

// Gen2EnvironmentSpec defines the desired state of Gen2Environment
type Gen2EnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     Gen2EnvironmentParameters `json:"forProvider"`
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
	InitProvider Gen2EnvironmentInitParameters `json:"initProvider,omitempty"`
}

// Gen2EnvironmentStatus defines the observed state of Gen2Environment.
type Gen2EnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        Gen2EnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Gen2Environment is the Schema for the Gen2Environments API. Manages an Azure IoT Time Series Insights Gen2 Environment.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Gen2Environment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.idProperties) || (has(self.initProvider) && has(self.initProvider.idProperties))",message="spec.forProvider.idProperties is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuName) || (has(self.initProvider) && has(self.initProvider.skuName))",message="spec.forProvider.skuName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storage) || (has(self.initProvider) && has(self.initProvider.storage))",message="spec.forProvider.storage is a required parameter"
	Spec   Gen2EnvironmentSpec   `json:"spec"`
	Status Gen2EnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Gen2EnvironmentList contains a list of Gen2Environments
type Gen2EnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gen2Environment `json:"items"`
}

// Repository type metadata.
var (
	Gen2Environment_Kind             = "Gen2Environment"
	Gen2Environment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Gen2Environment_Kind}.String()
	Gen2Environment_KindAPIVersion   = Gen2Environment_Kind + "." + CRDGroupVersion.String()
	Gen2Environment_GroupVersionKind = CRDGroupVersion.WithKind(Gen2Environment_Kind)
)

func init() {
	SchemeBuilder.Register(&Gen2Environment{}, &Gen2EnvironmentList{})
}
