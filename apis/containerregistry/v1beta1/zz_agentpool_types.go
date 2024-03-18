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

type AgentPoolInitParameters struct {

	// VMSS instance count. Defaults to 1.
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// The Azure Region where the Azure Container Registry Agent Pool should exist. Changing this forces a new Azure Container Registry Agent Pool to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags which should be assigned to the Azure Container Registry Agent Pool.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Sets the VM your agent pool will run on. Valid values are: S1 (2 vCPUs, 3 GiB RAM), S2 (4 vCPUs, 8 GiB RAM), S3 (8 vCPUs, 16 GiB RAM) or I6 (64 vCPUs, 216 GiB RAM, Isolated). Defaults to S1. Changing this forces a new Azure Container Registry Agent Pool to be created.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// The ID of the Virtual Network Subnet Resource where the agent machines will be running. Changing this forces a new Azure Container Registry Agent Pool to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	VirtualNetworkSubnetID *string `json:"virtualNetworkSubnetId,omitempty" tf:"virtual_network_subnet_id,omitempty"`

	// Reference to a Subnet in network to populate virtualNetworkSubnetId.
	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIDRef *v1.Reference `json:"virtualNetworkSubnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate virtualNetworkSubnetId.
	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIDSelector *v1.Selector `json:"virtualNetworkSubnetIdSelector,omitempty" tf:"-"`
}

type AgentPoolObservation struct {

	// Name of Azure Container Registry to create an Agent Pool for. Changing this forces a new Azure Container Registry Agent Pool to be created.
	ContainerRegistryName *string `json:"containerRegistryName,omitempty" tf:"container_registry_name,omitempty"`

	// The ID of the Azure Container Registry Agent Pool.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// VMSS instance count. Defaults to 1.
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// The Azure Region where the Azure Container Registry Agent Pool should exist. Changing this forces a new Azure Container Registry Agent Pool to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Resource Group where the Azure Container Registry Agent Pool should exist. Changing this forces a new Azure Container Registry Agent Pool to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags which should be assigned to the Azure Container Registry Agent Pool.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Sets the VM your agent pool will run on. Valid values are: S1 (2 vCPUs, 3 GiB RAM), S2 (4 vCPUs, 8 GiB RAM), S3 (8 vCPUs, 16 GiB RAM) or I6 (64 vCPUs, 216 GiB RAM, Isolated). Defaults to S1. Changing this forces a new Azure Container Registry Agent Pool to be created.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// The ID of the Virtual Network Subnet Resource where the agent machines will be running. Changing this forces a new Azure Container Registry Agent Pool to be created.
	VirtualNetworkSubnetID *string `json:"virtualNetworkSubnetId,omitempty" tf:"virtual_network_subnet_id,omitempty"`
}

type AgentPoolParameters struct {

	// Name of Azure Container Registry to create an Agent Pool for. Changing this forces a new Azure Container Registry Agent Pool to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/containerregistry/v1beta1.Registry
	// +kubebuilder:validation:Optional
	ContainerRegistryName *string `json:"containerRegistryName,omitempty" tf:"container_registry_name,omitempty"`

	// Reference to a Registry in containerregistry to populate containerRegistryName.
	// +kubebuilder:validation:Optional
	ContainerRegistryNameRef *v1.Reference `json:"containerRegistryNameRef,omitempty" tf:"-"`

	// Selector for a Registry in containerregistry to populate containerRegistryName.
	// +kubebuilder:validation:Optional
	ContainerRegistryNameSelector *v1.Selector `json:"containerRegistryNameSelector,omitempty" tf:"-"`

	// VMSS instance count. Defaults to 1.
	// +kubebuilder:validation:Optional
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// The Azure Region where the Azure Container Registry Agent Pool should exist. Changing this forces a new Azure Container Registry Agent Pool to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Resource Group where the Azure Container Registry Agent Pool should exist. Changing this forces a new Azure Container Registry Agent Pool to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Azure Container Registry Agent Pool.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Sets the VM your agent pool will run on. Valid values are: S1 (2 vCPUs, 3 GiB RAM), S2 (4 vCPUs, 8 GiB RAM), S3 (8 vCPUs, 16 GiB RAM) or I6 (64 vCPUs, 216 GiB RAM, Isolated). Defaults to S1. Changing this forces a new Azure Container Registry Agent Pool to be created.
	// +kubebuilder:validation:Optional
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// The ID of the Virtual Network Subnet Resource where the agent machines will be running. Changing this forces a new Azure Container Registry Agent Pool to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetID *string `json:"virtualNetworkSubnetId,omitempty" tf:"virtual_network_subnet_id,omitempty"`

	// Reference to a Subnet in network to populate virtualNetworkSubnetId.
	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIDRef *v1.Reference `json:"virtualNetworkSubnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate virtualNetworkSubnetId.
	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIDSelector *v1.Selector `json:"virtualNetworkSubnetIdSelector,omitempty" tf:"-"`
}

// AgentPoolSpec defines the desired state of AgentPool
type AgentPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AgentPoolParameters `json:"forProvider"`
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
	InitProvider AgentPoolInitParameters `json:"initProvider,omitempty"`
}

// AgentPoolStatus defines the observed state of AgentPool.
type AgentPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AgentPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AgentPool is the Schema for the AgentPools API. Manages an Azure Container Registry Agent Pool.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AgentPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   AgentPoolSpec   `json:"spec"`
	Status AgentPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AgentPoolList contains a list of AgentPools
type AgentPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AgentPool `json:"items"`
}

// Repository type metadata.
var (
	AgentPool_Kind             = "AgentPool"
	AgentPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AgentPool_Kind}.String()
	AgentPool_KindAPIVersion   = AgentPool_Kind + "." + CRDGroupVersion.String()
	AgentPool_GroupVersionKind = CRDGroupVersion.WithKind(AgentPool_Kind)
)

func init() {
	SchemeBuilder.Register(&AgentPool{}, &AgentPoolList{})
}
