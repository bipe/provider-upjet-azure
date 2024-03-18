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

type PoolInitParameters struct {

	// The encryption type of the pool. Valid values include Single, and Double. Defaults to Single. Changing this forces a new resource to be created.
	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// QoS Type of the pool. Valid values include Auto or Manual.
	QosType *string `json:"qosType,omitempty" tf:"qos_type,omitempty"`

	// The service level of the file system. Valid values include Premium, Standard, and Ultra. Changing this forces a new resource to be created.
	ServiceLevel *string `json:"serviceLevel,omitempty" tf:"service_level,omitempty"`

	// Provisioned size of the pool in TB. Value must be between 2 and 500.
	SizeInTb *float64 `json:"sizeInTb,omitempty" tf:"size_in_tb,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PoolObservation struct {

	// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// The encryption type of the pool. Valid values include Single, and Double. Defaults to Single. Changing this forces a new resource to be created.
	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`

	// The ID of the NetApp Pool.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// QoS Type of the pool. Valid values include Auto or Manual.
	QosType *string `json:"qosType,omitempty" tf:"qos_type,omitempty"`

	// The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The service level of the file system. Valid values include Premium, Standard, and Ultra. Changing this forces a new resource to be created.
	ServiceLevel *string `json:"serviceLevel,omitempty" tf:"service_level,omitempty"`

	// Provisioned size of the pool in TB. Value must be between 2 and 500.
	SizeInTb *float64 `json:"sizeInTb,omitempty" tf:"size_in_tb,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PoolParameters struct {

	// The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Reference to a Account to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// Selector for a Account to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// The encryption type of the pool. Valid values include Single, and Double. Defaults to Single. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// QoS Type of the pool. Valid values include Auto or Manual.
	// +kubebuilder:validation:Optional
	QosType *string `json:"qosType,omitempty" tf:"qos_type,omitempty"`

	// The name of the resource group where the NetApp Pool should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The service level of the file system. Valid values include Premium, Standard, and Ultra. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ServiceLevel *string `json:"serviceLevel,omitempty" tf:"service_level,omitempty"`

	// Provisioned size of the pool in TB. Value must be between 2 and 500.
	// +kubebuilder:validation:Optional
	SizeInTb *float64 `json:"sizeInTb,omitempty" tf:"size_in_tb,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// PoolSpec defines the desired state of Pool
type PoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PoolParameters `json:"forProvider"`
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
	InitProvider PoolInitParameters `json:"initProvider,omitempty"`
}

// PoolStatus defines the observed state of Pool.
type PoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Pool is the Schema for the Pools API. Manages a Pool within a NetApp Account.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Pool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceLevel) || (has(self.initProvider) && has(self.initProvider.serviceLevel))",message="spec.forProvider.serviceLevel is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sizeInTb) || (has(self.initProvider) && has(self.initProvider.sizeInTb))",message="spec.forProvider.sizeInTb is a required parameter"
	Spec   PoolSpec   `json:"spec"`
	Status PoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PoolList contains a list of Pools
type PoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pool `json:"items"`
}

// Repository type metadata.
var (
	Pool_Kind             = "Pool"
	Pool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Pool_Kind}.String()
	Pool_KindAPIVersion   = Pool_Kind + "." + CRDGroupVersion.String()
	Pool_GroupVersionKind = CRDGroupVersion.WithKind(Pool_Kind)
)

func init() {
	SchemeBuilder.Register(&Pool{}, &PoolList{})
}
