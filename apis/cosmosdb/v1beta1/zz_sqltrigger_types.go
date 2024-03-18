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

type SQLTriggerInitParameters struct {

	// Body of the Trigger.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// The operation the trigger is associated with. Possible values are All, Create, Update, Delete and Replace.
	Operation *string `json:"operation,omitempty" tf:"operation,omitempty"`

	// Type of the Trigger. Possible values are Pre and Post.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SQLTriggerObservation struct {

	// Body of the Trigger.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// The id of the Cosmos DB SQL Container to create the SQL Trigger within. Changing this forces a new SQL Trigger to be created.
	ContainerID *string `json:"containerId,omitempty" tf:"container_id,omitempty"`

	// The ID of the SQL Trigger.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The operation the trigger is associated with. Possible values are All, Create, Update, Delete and Replace.
	Operation *string `json:"operation,omitempty" tf:"operation,omitempty"`

	// Type of the Trigger. Possible values are Pre and Post.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SQLTriggerParameters struct {

	// Body of the Trigger.
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// The id of the Cosmos DB SQL Container to create the SQL Trigger within. Changing this forces a new SQL Trigger to be created.
	// +crossplane:generate:reference:type=SQLContainer
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ContainerID *string `json:"containerId,omitempty" tf:"container_id,omitempty"`

	// Reference to a SQLContainer to populate containerId.
	// +kubebuilder:validation:Optional
	ContainerIDRef *v1.Reference `json:"containerIdRef,omitempty" tf:"-"`

	// Selector for a SQLContainer to populate containerId.
	// +kubebuilder:validation:Optional
	ContainerIDSelector *v1.Selector `json:"containerIdSelector,omitempty" tf:"-"`

	// The operation the trigger is associated with. Possible values are All, Create, Update, Delete and Replace.
	// +kubebuilder:validation:Optional
	Operation *string `json:"operation,omitempty" tf:"operation,omitempty"`

	// Type of the Trigger. Possible values are Pre and Post.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// SQLTriggerSpec defines the desired state of SQLTrigger
type SQLTriggerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLTriggerParameters `json:"forProvider"`
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
	InitProvider SQLTriggerInitParameters `json:"initProvider,omitempty"`
}

// SQLTriggerStatus defines the observed state of SQLTrigger.
type SQLTriggerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLTriggerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SQLTrigger is the Schema for the SQLTriggers API. Manages an SQL Trigger.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SQLTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.body) || (has(self.initProvider) && has(self.initProvider.body))",message="spec.forProvider.body is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.operation) || (has(self.initProvider) && has(self.initProvider.operation))",message="spec.forProvider.operation is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   SQLTriggerSpec   `json:"spec"`
	Status SQLTriggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLTriggerList contains a list of SQLTriggers
type SQLTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLTrigger `json:"items"`
}

// Repository type metadata.
var (
	SQLTrigger_Kind             = "SQLTrigger"
	SQLTrigger_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLTrigger_Kind}.String()
	SQLTrigger_KindAPIVersion   = SQLTrigger_Kind + "." + CRDGroupVersion.String()
	SQLTrigger_GroupVersionKind = CRDGroupVersion.WithKind(SQLTrigger_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLTrigger{}, &SQLTriggerList{})
}
