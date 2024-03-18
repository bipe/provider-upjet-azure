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

type FederatedIdentityCredentialInitParameters struct {

	// Specifies the audience for this Federated Identity Credential.
	Audience []*string `json:"audience,omitempty" tf:"audience,omitempty"`

	// Specifies the issuer of this Federated Identity Credential.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// Specifies the name of the Resource Group within which this Federated Identity Credential should exist. Changing this forces a new Federated Identity Credential to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the subject for this Federated Identity Credential.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`
}

type FederatedIdentityCredentialObservation struct {

	// Specifies the audience for this Federated Identity Credential.
	Audience []*string `json:"audience,omitempty" tf:"audience,omitempty"`

	// The ID of the Federated Identity Credential.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the issuer of this Federated Identity Credential.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// Specifies parent ID of User Assigned Identity for this Federated Identity Credential. Changing this forces a new Federated Identity Credential to be created.
	ParentID *string `json:"parentId,omitempty" tf:"parent_id,omitempty"`

	// Specifies the name of the Resource Group within which this Federated Identity Credential should exist. Changing this forces a new Federated Identity Credential to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the subject for this Federated Identity Credential.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`
}

type FederatedIdentityCredentialParameters struct {

	// Specifies the audience for this Federated Identity Credential.
	// +kubebuilder:validation:Optional
	Audience []*string `json:"audience,omitempty" tf:"audience,omitempty"`

	// Specifies the issuer of this Federated Identity Credential.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// Specifies parent ID of User Assigned Identity for this Federated Identity Credential. Changing this forces a new Federated Identity Credential to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/managedidentity/v1beta1.UserAssignedIdentity
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ParentID *string `json:"parentId,omitempty" tf:"parent_id,omitempty"`

	// Reference to a UserAssignedIdentity in managedidentity to populate parentId.
	// +kubebuilder:validation:Optional
	ParentIDRef *v1.Reference `json:"parentIdRef,omitempty" tf:"-"`

	// Selector for a UserAssignedIdentity in managedidentity to populate parentId.
	// +kubebuilder:validation:Optional
	ParentIDSelector *v1.Selector `json:"parentIdSelector,omitempty" tf:"-"`

	// Specifies the name of the Resource Group within which this Federated Identity Credential should exist. Changing this forces a new Federated Identity Credential to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the subject for this Federated Identity Credential.
	// +kubebuilder:validation:Optional
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`
}

// FederatedIdentityCredentialSpec defines the desired state of FederatedIdentityCredential
type FederatedIdentityCredentialSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FederatedIdentityCredentialParameters `json:"forProvider"`
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
	InitProvider FederatedIdentityCredentialInitParameters `json:"initProvider,omitempty"`
}

// FederatedIdentityCredentialStatus defines the observed state of FederatedIdentityCredential.
type FederatedIdentityCredentialStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FederatedIdentityCredentialObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FederatedIdentityCredential is the Schema for the FederatedIdentityCredentials API. Manages a Federated Identity Credential.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FederatedIdentityCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.audience) || (has(self.initProvider) && has(self.initProvider.audience))",message="spec.forProvider.audience is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.issuer) || (has(self.initProvider) && has(self.initProvider.issuer))",message="spec.forProvider.issuer is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subject) || (has(self.initProvider) && has(self.initProvider.subject))",message="spec.forProvider.subject is a required parameter"
	Spec   FederatedIdentityCredentialSpec   `json:"spec"`
	Status FederatedIdentityCredentialStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FederatedIdentityCredentialList contains a list of FederatedIdentityCredentials
type FederatedIdentityCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedIdentityCredential `json:"items"`
}

// Repository type metadata.
var (
	FederatedIdentityCredential_Kind             = "FederatedIdentityCredential"
	FederatedIdentityCredential_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FederatedIdentityCredential_Kind}.String()
	FederatedIdentityCredential_KindAPIVersion   = FederatedIdentityCredential_Kind + "." + CRDGroupVersion.String()
	FederatedIdentityCredential_GroupVersionKind = CRDGroupVersion.WithKind(FederatedIdentityCredential_Kind)
)

func init() {
	SchemeBuilder.Register(&FederatedIdentityCredential{}, &FederatedIdentityCredentialList{})
}
