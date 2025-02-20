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

type AccessPolicyInitParameters_2 struct {

	// The object ID of an Application in Azure Active Directory. Changing this forces a new resource to be created.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// List of certificate permissions, must be one or more from the following: Backup, Create, Delete, DeleteIssuers, Get, GetIssuers, Import, List, ListIssuers, ManageContacts, ManageIssuers, Purge, Recover, Restore, SetIssuers and Update.
	CertificatePermissions []*string `json:"certificatePermissions,omitempty" tf:"certificate_permissions,omitempty"`

	// List of key permissions, must be one or more from the following: Backup, Create, Decrypt, Delete, Encrypt, Get, Import, List, Purge, Recover, Restore, Sign, UnwrapKey, Update, Verify, WrapKey, Release, Rotate, GetRotationPolicy and SetRotationPolicy.
	KeyPermissions []*string `json:"keyPermissions,omitempty" tf:"key_permissions,omitempty"`

	// Specifies the id of the Key Vault resource. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Vault
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Reference to a Vault in keyvault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDRef *v1.Reference `json:"keyVaultIdRef,omitempty" tf:"-"`

	// Selector for a Vault in keyvault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDSelector *v1.Selector `json:"keyVaultIdSelector,omitempty" tf:"-"`

	// The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault. The object ID of a service principal can be fetched from azuread_service_principal.object_id. The object ID must be unique for the list of access policies. Changing this forces a new resource to be created.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// List of secret permissions, must be one or more from the following: Backup, Delete, Get, List, Purge, Recover, Restore and Set.
	SecretPermissions []*string `json:"secretPermissions,omitempty" tf:"secret_permissions,omitempty"`

	// List of storage permissions, must be one or more from the following: Backup, Delete, DeleteSAS, Get, GetSAS, List, ListSAS, Purge, Recover, RegenerateKey, Restore, Set, SetSAS and Update.
	StoragePermissions []*string `json:"storagePermissions,omitempty" tf:"storage_permissions,omitempty"`

	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault. Changing this forces a new resource to be created.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AccessPolicyObservation_2 struct {

	// The object ID of an Application in Azure Active Directory. Changing this forces a new resource to be created.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// List of certificate permissions, must be one or more from the following: Backup, Create, Delete, DeleteIssuers, Get, GetIssuers, Import, List, ListIssuers, ManageContacts, ManageIssuers, Purge, Recover, Restore, SetIssuers and Update.
	CertificatePermissions []*string `json:"certificatePermissions,omitempty" tf:"certificate_permissions,omitempty"`

	// Key Vault Access Policy ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of key permissions, must be one or more from the following: Backup, Create, Decrypt, Delete, Encrypt, Get, Import, List, Purge, Recover, Restore, Sign, UnwrapKey, Update, Verify, WrapKey, Release, Rotate, GetRotationPolicy and SetRotationPolicy.
	KeyPermissions []*string `json:"keyPermissions,omitempty" tf:"key_permissions,omitempty"`

	// Specifies the id of the Key Vault resource. Changing this forces a new resource to be created.
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault. The object ID of a service principal can be fetched from azuread_service_principal.object_id. The object ID must be unique for the list of access policies. Changing this forces a new resource to be created.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// List of secret permissions, must be one or more from the following: Backup, Delete, Get, List, Purge, Recover, Restore and Set.
	SecretPermissions []*string `json:"secretPermissions,omitempty" tf:"secret_permissions,omitempty"`

	// List of storage permissions, must be one or more from the following: Backup, Delete, DeleteSAS, Get, GetSAS, List, ListSAS, Purge, Recover, RegenerateKey, Restore, Set, SetSAS and Update.
	StoragePermissions []*string `json:"storagePermissions,omitempty" tf:"storage_permissions,omitempty"`

	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault. Changing this forces a new resource to be created.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AccessPolicyParameters_2 struct {

	// The object ID of an Application in Azure Active Directory. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// List of certificate permissions, must be one or more from the following: Backup, Create, Delete, DeleteIssuers, Get, GetIssuers, Import, List, ListIssuers, ManageContacts, ManageIssuers, Purge, Recover, Restore, SetIssuers and Update.
	// +kubebuilder:validation:Optional
	CertificatePermissions []*string `json:"certificatePermissions,omitempty" tf:"certificate_permissions,omitempty"`

	// List of key permissions, must be one or more from the following: Backup, Create, Decrypt, Delete, Encrypt, Get, Import, List, Purge, Recover, Restore, Sign, UnwrapKey, Update, Verify, WrapKey, Release, Rotate, GetRotationPolicy and SetRotationPolicy.
	// +kubebuilder:validation:Optional
	KeyPermissions []*string `json:"keyPermissions,omitempty" tf:"key_permissions,omitempty"`

	// Specifies the id of the Key Vault resource. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Vault
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Reference to a Vault in keyvault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDRef *v1.Reference `json:"keyVaultIdRef,omitempty" tf:"-"`

	// Selector for a Vault in keyvault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDSelector *v1.Selector `json:"keyVaultIdSelector,omitempty" tf:"-"`

	// The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault. The object ID of a service principal can be fetched from azuread_service_principal.object_id. The object ID must be unique for the list of access policies. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// List of secret permissions, must be one or more from the following: Backup, Delete, Get, List, Purge, Recover, Restore and Set.
	// +kubebuilder:validation:Optional
	SecretPermissions []*string `json:"secretPermissions,omitempty" tf:"secret_permissions,omitempty"`

	// List of storage permissions, must be one or more from the following: Backup, Delete, DeleteSAS, Get, GetSAS, List, ListSAS, Purge, Recover, RegenerateKey, Restore, Set, SetSAS and Update.
	// +kubebuilder:validation:Optional
	StoragePermissions []*string `json:"storagePermissions,omitempty" tf:"storage_permissions,omitempty"`

	// The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

// AccessPolicySpec defines the desired state of AccessPolicy
type AccessPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessPolicyParameters_2 `json:"forProvider"`
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
	InitProvider AccessPolicyInitParameters_2 `json:"initProvider,omitempty"`
}

// AccessPolicyStatus defines the observed state of AccessPolicy.
type AccessPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessPolicyObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AccessPolicy is the Schema for the AccessPolicys API. Manages a Key Vault Access Policy.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.objectId) || (has(self.initProvider) && has(self.initProvider.objectId))",message="spec.forProvider.objectId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tenantId) || (has(self.initProvider) && has(self.initProvider.tenantId))",message="spec.forProvider.tenantId is a required parameter"
	Spec   AccessPolicySpec   `json:"spec"`
	Status AccessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessPolicyList contains a list of AccessPolicys
type AccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessPolicy `json:"items"`
}

// Repository type metadata.
var (
	AccessPolicy_Kind             = "AccessPolicy"
	AccessPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessPolicy_Kind}.String()
	AccessPolicy_KindAPIVersion   = AccessPolicy_Kind + "." + CRDGroupVersion.String()
	AccessPolicy_GroupVersionKind = CRDGroupVersion.WithKind(AccessPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessPolicy{}, &AccessPolicyList{})
}
