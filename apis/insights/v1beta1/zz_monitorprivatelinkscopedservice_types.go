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

type MonitorPrivateLinkScopedServiceInitParameters struct {

	// The ID of the linked resource. It must be the Log Analytics workspace or the Application Insights component or the Data Collection endpoint. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=ApplicationInsights
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	LinkedResourceID *string `json:"linkedResourceId,omitempty" tf:"linked_resource_id,omitempty"`

	// Reference to a ApplicationInsights to populate linkedResourceId.
	// +kubebuilder:validation:Optional
	LinkedResourceIDRef *v1.Reference `json:"linkedResourceIdRef,omitempty" tf:"-"`

	// Selector for a ApplicationInsights to populate linkedResourceId.
	// +kubebuilder:validation:Optional
	LinkedResourceIDSelector *v1.Selector `json:"linkedResourceIdSelector,omitempty" tf:"-"`
}

type MonitorPrivateLinkScopedServiceObservation struct {

	// The ID of the Azure Monitor Private Link Scoped Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the linked resource. It must be the Log Analytics workspace or the Application Insights component or the Data Collection endpoint. Changing this forces a new resource to be created.
	LinkedResourceID *string `json:"linkedResourceId,omitempty" tf:"linked_resource_id,omitempty"`

	// The name of the Resource Group where the Azure Monitor Private Link Scoped Service should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The name of the Azure Monitor Private Link Scope. Changing this forces a new resource to be created.
	ScopeName *string `json:"scopeName,omitempty" tf:"scope_name,omitempty"`
}

type MonitorPrivateLinkScopedServiceParameters struct {

	// The ID of the linked resource. It must be the Log Analytics workspace or the Application Insights component or the Data Collection endpoint. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=ApplicationInsights
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LinkedResourceID *string `json:"linkedResourceId,omitempty" tf:"linked_resource_id,omitempty"`

	// Reference to a ApplicationInsights to populate linkedResourceId.
	// +kubebuilder:validation:Optional
	LinkedResourceIDRef *v1.Reference `json:"linkedResourceIdRef,omitempty" tf:"-"`

	// Selector for a ApplicationInsights to populate linkedResourceId.
	// +kubebuilder:validation:Optional
	LinkedResourceIDSelector *v1.Selector `json:"linkedResourceIdSelector,omitempty" tf:"-"`

	// The name of the Resource Group where the Azure Monitor Private Link Scoped Service should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the Azure Monitor Private Link Scope. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=MonitorPrivateLinkScope
	// +kubebuilder:validation:Optional
	ScopeName *string `json:"scopeName,omitempty" tf:"scope_name,omitempty"`

	// Reference to a MonitorPrivateLinkScope to populate scopeName.
	// +kubebuilder:validation:Optional
	ScopeNameRef *v1.Reference `json:"scopeNameRef,omitempty" tf:"-"`

	// Selector for a MonitorPrivateLinkScope to populate scopeName.
	// +kubebuilder:validation:Optional
	ScopeNameSelector *v1.Selector `json:"scopeNameSelector,omitempty" tf:"-"`
}

// MonitorPrivateLinkScopedServiceSpec defines the desired state of MonitorPrivateLinkScopedService
type MonitorPrivateLinkScopedServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorPrivateLinkScopedServiceParameters `json:"forProvider"`
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
	InitProvider MonitorPrivateLinkScopedServiceInitParameters `json:"initProvider,omitempty"`
}

// MonitorPrivateLinkScopedServiceStatus defines the observed state of MonitorPrivateLinkScopedService.
type MonitorPrivateLinkScopedServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorPrivateLinkScopedServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MonitorPrivateLinkScopedService is the Schema for the MonitorPrivateLinkScopedServices API. Manages an Azure Monitor Private Link Scoped Service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorPrivateLinkScopedService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorPrivateLinkScopedServiceSpec   `json:"spec"`
	Status            MonitorPrivateLinkScopedServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorPrivateLinkScopedServiceList contains a list of MonitorPrivateLinkScopedServices
type MonitorPrivateLinkScopedServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorPrivateLinkScopedService `json:"items"`
}

// Repository type metadata.
var (
	MonitorPrivateLinkScopedService_Kind             = "MonitorPrivateLinkScopedService"
	MonitorPrivateLinkScopedService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorPrivateLinkScopedService_Kind}.String()
	MonitorPrivateLinkScopedService_KindAPIVersion   = MonitorPrivateLinkScopedService_Kind + "." + CRDGroupVersion.String()
	MonitorPrivateLinkScopedService_GroupVersionKind = CRDGroupVersion.WithKind(MonitorPrivateLinkScopedService_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorPrivateLinkScopedService{}, &MonitorPrivateLinkScopedServiceList{})
}
