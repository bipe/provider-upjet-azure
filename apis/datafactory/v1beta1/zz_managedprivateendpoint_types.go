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

type ManagedPrivateEndpointInitParameters struct {

	// The ID of the Data Factory on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// Fully qualified domain names. Changing this forces a new resource to be created.
	Fqdns []*string `json:"fqdns,omitempty" tf:"fqdns,omitempty"`

	// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the sub resource name which the Data Factory Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	SubresourceName *string `json:"subresourceName,omitempty" tf:"subresource_name,omitempty"`

	// The ID of the Private Link Enabled Remote Resource which this Data Factory Private Endpoint should be connected to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// Reference to a Account in storage to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDRef *v1.Reference `json:"targetResourceIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDSelector *v1.Selector `json:"targetResourceIdSelector,omitempty" tf:"-"`
}

type ManagedPrivateEndpointObservation struct {

	// The ID of the Data Factory on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Fully qualified domain names. Changing this forces a new resource to be created.
	Fqdns []*string `json:"fqdns,omitempty" tf:"fqdns,omitempty"`

	// The ID of the Data Factory Managed Private Endpoint.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the sub resource name which the Data Factory Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	SubresourceName *string `json:"subresourceName,omitempty" tf:"subresource_name,omitempty"`

	// The ID of the Private Link Enabled Remote Resource which this Data Factory Private Endpoint should be connected to. Changing this forces a new resource to be created.
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`
}

type ManagedPrivateEndpointParameters struct {

	// The ID of the Data Factory on which to create the Managed Private Endpoint. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// Fully qualified domain names. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Fqdns []*string `json:"fqdns,omitempty" tf:"fqdns,omitempty"`

	// Specifies the name which should be used for this Managed Private Endpoint. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the sub resource name which the Data Factory Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SubresourceName *string `json:"subresourceName,omitempty" tf:"subresource_name,omitempty"`

	// The ID of the Private Link Enabled Remote Resource which this Data Factory Private Endpoint should be connected to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// Reference to a Account in storage to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDRef *v1.Reference `json:"targetResourceIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDSelector *v1.Selector `json:"targetResourceIdSelector,omitempty" tf:"-"`
}

// ManagedPrivateEndpointSpec defines the desired state of ManagedPrivateEndpoint
type ManagedPrivateEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedPrivateEndpointParameters `json:"forProvider"`
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
	InitProvider ManagedPrivateEndpointInitParameters `json:"initProvider,omitempty"`
}

// ManagedPrivateEndpointStatus defines the observed state of ManagedPrivateEndpoint.
type ManagedPrivateEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedPrivateEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ManagedPrivateEndpoint is the Schema for the ManagedPrivateEndpoints API. Manages a Data Factory Managed Private Endpoint.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ManagedPrivateEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ManagedPrivateEndpointSpec   `json:"spec"`
	Status ManagedPrivateEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedPrivateEndpointList contains a list of ManagedPrivateEndpoints
type ManagedPrivateEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedPrivateEndpoint `json:"items"`
}

// Repository type metadata.
var (
	ManagedPrivateEndpoint_Kind             = "ManagedPrivateEndpoint"
	ManagedPrivateEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedPrivateEndpoint_Kind}.String()
	ManagedPrivateEndpoint_KindAPIVersion   = ManagedPrivateEndpoint_Kind + "." + CRDGroupVersion.String()
	ManagedPrivateEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(ManagedPrivateEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedPrivateEndpoint{}, &ManagedPrivateEndpointList{})
}
