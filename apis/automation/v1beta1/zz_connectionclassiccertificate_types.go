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

type ConnectionClassicCertificateInitParameters struct {

	// The name of the certificate asset.
	CertificateAssetName *string `json:"certificateAssetName,omitempty" tf:"certificate_asset_name,omitempty"`

	// A description for this Connection.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The id of subscription.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// The name of subscription.
	SubscriptionName *string `json:"subscriptionName,omitempty" tf:"subscription_name,omitempty"`
}

type ConnectionClassicCertificateObservation struct {

	// The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// The name of the certificate asset.
	CertificateAssetName *string `json:"certificateAssetName,omitempty" tf:"certificate_asset_name,omitempty"`

	// A description for this Connection.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Automation Connection ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The id of subscription.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// The name of subscription.
	SubscriptionName *string `json:"subscriptionName,omitempty" tf:"subscription_name,omitempty"`
}

type ConnectionClassicCertificateParameters struct {

	// The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/automation/v1beta1.Account
	// +kubebuilder:validation:Optional
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// Reference to a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameRef *v1.Reference `json:"automationAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameSelector *v1.Selector `json:"automationAccountNameSelector,omitempty" tf:"-"`

	// The name of the certificate asset.
	// +kubebuilder:validation:Optional
	CertificateAssetName *string `json:"certificateAssetName,omitempty" tf:"certificate_asset_name,omitempty"`

	// A description for this Connection.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the resource group in which the Connection is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The id of subscription.
	// +kubebuilder:validation:Optional
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// The name of subscription.
	// +kubebuilder:validation:Optional
	SubscriptionName *string `json:"subscriptionName,omitempty" tf:"subscription_name,omitempty"`
}

// ConnectionClassicCertificateSpec defines the desired state of ConnectionClassicCertificate
type ConnectionClassicCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionClassicCertificateParameters `json:"forProvider"`
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
	InitProvider ConnectionClassicCertificateInitParameters `json:"initProvider,omitempty"`
}

// ConnectionClassicCertificateStatus defines the observed state of ConnectionClassicCertificate.
type ConnectionClassicCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionClassicCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ConnectionClassicCertificate is the Schema for the ConnectionClassicCertificates API. Manages an Automation Connection with type
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ConnectionClassicCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.certificateAssetName) || (has(self.initProvider) && has(self.initProvider.certificateAssetName))",message="spec.forProvider.certificateAssetName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subscriptionId) || (has(self.initProvider) && has(self.initProvider.subscriptionId))",message="spec.forProvider.subscriptionId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subscriptionName) || (has(self.initProvider) && has(self.initProvider.subscriptionName))",message="spec.forProvider.subscriptionName is a required parameter"
	Spec   ConnectionClassicCertificateSpec   `json:"spec"`
	Status ConnectionClassicCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionClassicCertificateList contains a list of ConnectionClassicCertificates
type ConnectionClassicCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectionClassicCertificate `json:"items"`
}

// Repository type metadata.
var (
	ConnectionClassicCertificate_Kind             = "ConnectionClassicCertificate"
	ConnectionClassicCertificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConnectionClassicCertificate_Kind}.String()
	ConnectionClassicCertificate_KindAPIVersion   = ConnectionClassicCertificate_Kind + "." + CRDGroupVersion.String()
	ConnectionClassicCertificate_GroupVersionKind = CRDGroupVersion.WithKind(ConnectionClassicCertificate_Kind)
)

func init() {
	SchemeBuilder.Register(&ConnectionClassicCertificate{}, &ConnectionClassicCertificateList{})
}
