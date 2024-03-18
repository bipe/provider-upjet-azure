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

type WorkspaceExtendedAuditingPolicyInitParameters struct {

	// Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its master database audit events to Azure Monitor. Defaults to true.
	LogMonitoringEnabled *bool `json:"logMonitoringEnabled,omitempty" tf:"log_monitoring_enabled,omitempty"`

	// The number of days to retain logs for in the storage account. Defaults to 0.
	RetentionInDays *float64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

	// Is storage_account_access_key value the storage's secondary key?
	StorageAccountAccessKeyIsSecondary *bool `json:"storageAccountAccessKeyIsSecondary,omitempty" tf:"storage_account_access_key_is_secondary,omitempty"`

	// The blob storage endpoint (e.g. https://example.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("primary_blob_endpoint",true)
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`

	// Reference to a Account in storage to populate storageEndpoint.
	// +kubebuilder:validation:Optional
	StorageEndpointRef *v1.Reference `json:"storageEndpointRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageEndpoint.
	// +kubebuilder:validation:Optional
	StorageEndpointSelector *v1.Selector `json:"storageEndpointSelector,omitempty" tf:"-"`
}

type WorkspaceExtendedAuditingPolicyObservation struct {

	// The ID of the Synapse Workspace Extended Auditing Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its master database audit events to Azure Monitor. Defaults to true.
	LogMonitoringEnabled *bool `json:"logMonitoringEnabled,omitempty" tf:"log_monitoring_enabled,omitempty"`

	// The number of days to retain logs for in the storage account. Defaults to 0.
	RetentionInDays *float64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

	// Is storage_account_access_key value the storage's secondary key?
	StorageAccountAccessKeyIsSecondary *bool `json:"storageAccountAccessKeyIsSecondary,omitempty" tf:"storage_account_access_key_is_secondary,omitempty"`

	// The blob storage endpoint (e.g. https://example.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`

	// The ID of the Synapse workspace to set the extended auditing policy. Changing this forces a new resource to be created.
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`
}

type WorkspaceExtendedAuditingPolicyParameters struct {

	// Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its master database audit events to Azure Monitor. Defaults to true.
	// +kubebuilder:validation:Optional
	LogMonitoringEnabled *bool `json:"logMonitoringEnabled,omitempty" tf:"log_monitoring_enabled,omitempty"`

	// The number of days to retain logs for in the storage account. Defaults to 0.
	// +kubebuilder:validation:Optional
	RetentionInDays *float64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

	// Is storage_account_access_key value the storage's secondary key?
	// +kubebuilder:validation:Optional
	StorageAccountAccessKeyIsSecondary *bool `json:"storageAccountAccessKeyIsSecondary,omitempty" tf:"storage_account_access_key_is_secondary,omitempty"`

	// The access key to use for the auditing storage account.
	// +kubebuilder:validation:Optional
	StorageAccountAccessKeySecretRef *v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

	// The blob storage endpoint (e.g. https://example.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("primary_blob_endpoint",true)
	// +kubebuilder:validation:Optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`

	// Reference to a Account in storage to populate storageEndpoint.
	// +kubebuilder:validation:Optional
	StorageEndpointRef *v1.Reference `json:"storageEndpointRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageEndpoint.
	// +kubebuilder:validation:Optional
	StorageEndpointSelector *v1.Selector `json:"storageEndpointSelector,omitempty" tf:"-"`

	// The ID of the Synapse workspace to set the extended auditing policy. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/synapse/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`

	// Reference to a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDRef *v1.Reference `json:"synapseWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDSelector *v1.Selector `json:"synapseWorkspaceIdSelector,omitempty" tf:"-"`
}

// WorkspaceExtendedAuditingPolicySpec defines the desired state of WorkspaceExtendedAuditingPolicy
type WorkspaceExtendedAuditingPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceExtendedAuditingPolicyParameters `json:"forProvider"`
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
	InitProvider WorkspaceExtendedAuditingPolicyInitParameters `json:"initProvider,omitempty"`
}

// WorkspaceExtendedAuditingPolicyStatus defines the observed state of WorkspaceExtendedAuditingPolicy.
type WorkspaceExtendedAuditingPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceExtendedAuditingPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// WorkspaceExtendedAuditingPolicy is the Schema for the WorkspaceExtendedAuditingPolicys API. Manages a Synapse Workspace Extended Auditing Policy.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type WorkspaceExtendedAuditingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceExtendedAuditingPolicySpec   `json:"spec"`
	Status            WorkspaceExtendedAuditingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceExtendedAuditingPolicyList contains a list of WorkspaceExtendedAuditingPolicys
type WorkspaceExtendedAuditingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspaceExtendedAuditingPolicy `json:"items"`
}

// Repository type metadata.
var (
	WorkspaceExtendedAuditingPolicy_Kind             = "WorkspaceExtendedAuditingPolicy"
	WorkspaceExtendedAuditingPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WorkspaceExtendedAuditingPolicy_Kind}.String()
	WorkspaceExtendedAuditingPolicy_KindAPIVersion   = WorkspaceExtendedAuditingPolicy_Kind + "." + CRDGroupVersion.String()
	WorkspaceExtendedAuditingPolicy_GroupVersionKind = CRDGroupVersion.WithKind(WorkspaceExtendedAuditingPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&WorkspaceExtendedAuditingPolicy{}, &WorkspaceExtendedAuditingPolicyList{})
}
