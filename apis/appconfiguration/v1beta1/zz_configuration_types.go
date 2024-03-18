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

type ConfigurationInitParameters struct {

	// An encryption block as defined below.
	Encryption []EncryptionInitParameters `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// An identity block as defined below.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether local authentication methods is enabled. Defaults to true.
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Public Network Access setting of the App Configuration. Possible values are Enabled and Disabled.
	PublicNetworkAccess *string `json:"publicNetworkAccess,omitempty" tf:"public_network_access,omitempty"`

	// Whether Purge Protection is enabled. This field only works for standard sku. Defaults to false.
	PurgeProtectionEnabled *bool `json:"purgeProtectionEnabled,omitempty" tf:"purge_protection_enabled,omitempty"`

	// One or more replica blocks as defined below.
	Replica []ReplicaInitParameters `json:"replica,omitempty" tf:"replica,omitempty"`

	// The SKU name of the App Configuration. Possible values are free and standard. Defaults to free.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// The number of days that items should be retained for once soft-deleted. This field only works for standard sku. This value can be between 1 and 7 days. Defaults to 7. Changing this forces a new resource to be created.
	SoftDeleteRetentionDays *float64 `json:"softDeleteRetentionDays,omitempty" tf:"soft_delete_retention_days,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ConfigurationObservation struct {

	// An encryption block as defined below.
	Encryption []EncryptionObservation `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// The URL of the App Configuration.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The App Configuration ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether local authentication methods is enabled. Defaults to true.
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A primary_read_key block as defined below containing the primary read access key.
	PrimaryReadKey []PrimaryReadKeyObservation `json:"primaryReadKey,omitempty" tf:"primary_read_key,omitempty"`

	// A primary_write_key block as defined below containing the primary write access key.
	PrimaryWriteKey []PrimaryWriteKeyObservation `json:"primaryWriteKey,omitempty" tf:"primary_write_key,omitempty"`

	// The Public Network Access setting of the App Configuration. Possible values are Enabled and Disabled.
	PublicNetworkAccess *string `json:"publicNetworkAccess,omitempty" tf:"public_network_access,omitempty"`

	// Whether Purge Protection is enabled. This field only works for standard sku. Defaults to false.
	PurgeProtectionEnabled *bool `json:"purgeProtectionEnabled,omitempty" tf:"purge_protection_enabled,omitempty"`

	// One or more replica blocks as defined below.
	Replica []ReplicaObservation `json:"replica,omitempty" tf:"replica,omitempty"`

	// The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A secondary_read_key block as defined below containing the secondary read access key.
	SecondaryReadKey []SecondaryReadKeyObservation `json:"secondaryReadKey,omitempty" tf:"secondary_read_key,omitempty"`

	// A secondary_write_key block as defined below containing the secondary write access key.
	SecondaryWriteKey []SecondaryWriteKeyObservation `json:"secondaryWriteKey,omitempty" tf:"secondary_write_key,omitempty"`

	// The SKU name of the App Configuration. Possible values are free and standard. Defaults to free.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// The number of days that items should be retained for once soft-deleted. This field only works for standard sku. This value can be between 1 and 7 days. Defaults to 7. Changing this forces a new resource to be created.
	SoftDeleteRetentionDays *float64 `json:"softDeleteRetentionDays,omitempty" tf:"soft_delete_retention_days,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ConfigurationParameters struct {

	// An encryption block as defined below.
	// +kubebuilder:validation:Optional
	Encryption []EncryptionParameters `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether local authentication methods is enabled. Defaults to true.
	// +kubebuilder:validation:Optional
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Public Network Access setting of the App Configuration. Possible values are Enabled and Disabled.
	// +kubebuilder:validation:Optional
	PublicNetworkAccess *string `json:"publicNetworkAccess,omitempty" tf:"public_network_access,omitempty"`

	// Whether Purge Protection is enabled. This field only works for standard sku. Defaults to false.
	// +kubebuilder:validation:Optional
	PurgeProtectionEnabled *bool `json:"purgeProtectionEnabled,omitempty" tf:"purge_protection_enabled,omitempty"`

	// One or more replica blocks as defined below.
	// +kubebuilder:validation:Optional
	Replica []ReplicaParameters `json:"replica,omitempty" tf:"replica,omitempty"`

	// The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The SKU name of the App Configuration. Possible values are free and standard. Defaults to free.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// The number of days that items should be retained for once soft-deleted. This field only works for standard sku. This value can be between 1 and 7 days. Defaults to 7. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SoftDeleteRetentionDays *float64 `json:"softDeleteRetentionDays,omitempty" tf:"soft_delete_retention_days,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EncryptionInitParameters struct {

	// Specifies the client id of the identity which will be used to access key vault.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/managedidentity/v1beta1.UserAssignedIdentity
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("client_id",true)
	IdentityClientID *string `json:"identityClientId,omitempty" tf:"identity_client_id,omitempty"`

	// Reference to a UserAssignedIdentity in managedidentity to populate identityClientId.
	// +kubebuilder:validation:Optional
	IdentityClientIDRef *v1.Reference `json:"identityClientIdRef,omitempty" tf:"-"`

	// Selector for a UserAssignedIdentity in managedidentity to populate identityClientId.
	// +kubebuilder:validation:Optional
	IdentityClientIDSelector *v1.Selector `json:"identityClientIdSelector,omitempty" tf:"-"`

	// Specifies the URI of the key vault key used to encrypt data.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	KeyVaultKeyIdentifier *string `json:"keyVaultKeyIdentifier,omitempty" tf:"key_vault_key_identifier,omitempty"`

	// Reference to a Key in keyvault to populate keyVaultKeyIdentifier.
	// +kubebuilder:validation:Optional
	KeyVaultKeyIdentifierRef *v1.Reference `json:"keyVaultKeyIdentifierRef,omitempty" tf:"-"`

	// Selector for a Key in keyvault to populate keyVaultKeyIdentifier.
	// +kubebuilder:validation:Optional
	KeyVaultKeyIdentifierSelector *v1.Selector `json:"keyVaultKeyIdentifierSelector,omitempty" tf:"-"`
}

type EncryptionObservation struct {

	// Specifies the client id of the identity which will be used to access key vault.
	IdentityClientID *string `json:"identityClientId,omitempty" tf:"identity_client_id,omitempty"`

	// Specifies the URI of the key vault key used to encrypt data.
	KeyVaultKeyIdentifier *string `json:"keyVaultKeyIdentifier,omitempty" tf:"key_vault_key_identifier,omitempty"`
}

type EncryptionParameters struct {

	// Specifies the client id of the identity which will be used to access key vault.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/managedidentity/v1beta1.UserAssignedIdentity
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("client_id",true)
	// +kubebuilder:validation:Optional
	IdentityClientID *string `json:"identityClientId,omitempty" tf:"identity_client_id,omitempty"`

	// Reference to a UserAssignedIdentity in managedidentity to populate identityClientId.
	// +kubebuilder:validation:Optional
	IdentityClientIDRef *v1.Reference `json:"identityClientIdRef,omitempty" tf:"-"`

	// Selector for a UserAssignedIdentity in managedidentity to populate identityClientId.
	// +kubebuilder:validation:Optional
	IdentityClientIDSelector *v1.Selector `json:"identityClientIdSelector,omitempty" tf:"-"`

	// Specifies the URI of the key vault key used to encrypt data.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyVaultKeyIdentifier *string `json:"keyVaultKeyIdentifier,omitempty" tf:"key_vault_key_identifier,omitempty"`

	// Reference to a Key in keyvault to populate keyVaultKeyIdentifier.
	// +kubebuilder:validation:Optional
	KeyVaultKeyIdentifierRef *v1.Reference `json:"keyVaultKeyIdentifierRef,omitempty" tf:"-"`

	// Selector for a Key in keyvault to populate keyVaultKeyIdentifier.
	// +kubebuilder:validation:Optional
	KeyVaultKeyIdentifierSelector *v1.Selector `json:"keyVaultKeyIdentifierSelector,omitempty" tf:"-"`
}

type IdentityInitParameters struct {

	// A list of User Assigned Managed Identity IDs to be assigned to this App Configuration.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this App Configuration. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// A list of User Assigned Managed Identity IDs to be assigned to this App Configuration.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this App Configuration. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// A list of User Assigned Managed Identity IDs to be assigned to this App Configuration.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this App Configuration. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type PrimaryReadKeyInitParameters struct {
}

type PrimaryReadKeyObservation struct {

	// The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The ID of the Access Key.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Secret of the Access Key.
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`
}

type PrimaryReadKeyParameters struct {
}

type PrimaryWriteKeyInitParameters struct {
}

type PrimaryWriteKeyObservation struct {

	// The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The ID of the Access Key.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Secret of the Access Key.
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`
}

type PrimaryWriteKeyParameters struct {
}

type ReplicaInitParameters struct {

	// Specifies the supported Azure location where the replica exists.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the replica.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ReplicaObservation struct {

	// The URL of the App Configuration Replica.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The ID of the App Configuration Replica.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the replica exists.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the replica.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ReplicaParameters struct {

	// Specifies the supported Azure location where the replica exists.
	// +kubebuilder:validation:Optional
	Location *string `json:"location" tf:"location,omitempty"`

	// Specifies the name of the replica.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type SecondaryReadKeyInitParameters struct {
}

type SecondaryReadKeyObservation struct {

	// The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The ID of the Access Key.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Secret of the Access Key.
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`
}

type SecondaryReadKeyParameters struct {
}

type SecondaryWriteKeyInitParameters struct {
}

type SecondaryWriteKeyObservation struct {

	// The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The ID of the Access Key.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Secret of the Access Key.
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`
}

type SecondaryWriteKeyParameters struct {
}

// ConfigurationSpec defines the desired state of Configuration
type ConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigurationParameters `json:"forProvider"`
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
	InitProvider ConfigurationInitParameters `json:"initProvider,omitempty"`
}

// ConfigurationStatus defines the observed state of Configuration.
type ConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Configuration is the Schema for the Configurations API. Manages an Azure App Configuration.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Configuration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   ConfigurationSpec   `json:"spec"`
	Status ConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationList contains a list of Configurations
type ConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Configuration `json:"items"`
}

// Repository type metadata.
var (
	Configuration_Kind             = "Configuration"
	Configuration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Configuration_Kind}.String()
	Configuration_KindAPIVersion   = Configuration_Kind + "." + CRDGroupVersion.String()
	Configuration_GroupVersionKind = CRDGroupVersion.WithKind(Configuration_Kind)
)

func init() {
	SchemeBuilder.Register(&Configuration{}, &ConfigurationList{})
}
