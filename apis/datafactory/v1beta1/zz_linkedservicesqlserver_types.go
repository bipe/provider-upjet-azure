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

type LinkedServiceSQLServerInitParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string in which to authenticate with the SQL Server. Exactly one of either connection_string or key_vault_connection_string is required.
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The description for the Data Factory Linked Service SQL Server.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A key_vault_connection_string block as defined below. Use this argument to store SQL Server connection string in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service. Exactly one of either connection_string or key_vault_connection_string is required.
	KeyVaultConnectionString []LinkedServiceSQLServerKeyVaultConnectionStringInitParameters `json:"keyVaultConnectionString,omitempty" tf:"key_vault_connection_string,omitempty"`

	// A key_vault_password block as defined below. Use this argument to store SQL Server password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword []LinkedServiceSQLServerKeyVaultPasswordInitParameters `json:"keyVaultPassword,omitempty" tf:"key_vault_password,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The on-premises Windows authentication user name.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type LinkedServiceSQLServerKeyVaultConnectionStringInitParameters struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Specifies the secret name in Azure Key Vault that stores SQL Server connection string.
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`
}

type LinkedServiceSQLServerKeyVaultConnectionStringObservation struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Specifies the secret name in Azure Key Vault that stores SQL Server connection string.
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`
}

type LinkedServiceSQLServerKeyVaultConnectionStringParameters struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	LinkedServiceName *string `json:"linkedServiceName" tf:"linked_service_name,omitempty"`

	// Specifies the secret name in Azure Key Vault that stores SQL Server connection string.
	// +kubebuilder:validation:Optional
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type LinkedServiceSQLServerKeyVaultPasswordInitParameters struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.LinkedServiceKeyVault
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Reference to a LinkedServiceKeyVault in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameRef *v1.Reference `json:"linkedServiceNameRef,omitempty" tf:"-"`

	// Selector for a LinkedServiceKeyVault in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameSelector *v1.Selector `json:"linkedServiceNameSelector,omitempty" tf:"-"`

	// Specifies the secret name in Azure Key Vault that stores SQL Server password.
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`
}

type LinkedServiceSQLServerKeyVaultPasswordObservation struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Specifies the secret name in Azure Key Vault that stores SQL Server password.
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`
}

type LinkedServiceSQLServerKeyVaultPasswordParameters struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.LinkedServiceKeyVault
	// +kubebuilder:validation:Optional
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Reference to a LinkedServiceKeyVault in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameRef *v1.Reference `json:"linkedServiceNameRef,omitempty" tf:"-"`

	// Selector for a LinkedServiceKeyVault in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameSelector *v1.Selector `json:"linkedServiceNameSelector,omitempty" tf:"-"`

	// Specifies the secret name in Azure Key Vault that stores SQL Server password.
	// +kubebuilder:validation:Optional
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type LinkedServiceSQLServerObservation struct {

	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string in which to authenticate with the SQL Server. Exactly one of either connection_string or key_vault_connection_string is required.
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Linked Service SQL Server.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Data Factory SQL Server Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A key_vault_connection_string block as defined below. Use this argument to store SQL Server connection string in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service. Exactly one of either connection_string or key_vault_connection_string is required.
	KeyVaultConnectionString []LinkedServiceSQLServerKeyVaultConnectionStringObservation `json:"keyVaultConnectionString,omitempty" tf:"key_vault_connection_string,omitempty"`

	// A key_vault_password block as defined below. Use this argument to store SQL Server password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	KeyVaultPassword []LinkedServiceSQLServerKeyVaultPasswordObservation `json:"keyVaultPassword,omitempty" tf:"key_vault_password,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The on-premises Windows authentication user name.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type LinkedServiceSQLServerParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string in which to authenticate with the SQL Server. Exactly one of either connection_string or key_vault_connection_string is required.
	// +kubebuilder:validation:Optional
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
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

	// The description for the Data Factory Linked Service SQL Server.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A key_vault_connection_string block as defined below. Use this argument to store SQL Server connection string in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service. Exactly one of either connection_string or key_vault_connection_string is required.
	// +kubebuilder:validation:Optional
	KeyVaultConnectionString []LinkedServiceSQLServerKeyVaultConnectionStringParameters `json:"keyVaultConnectionString,omitempty" tf:"key_vault_connection_string,omitempty"`

	// A key_vault_password block as defined below. Use this argument to store SQL Server password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	KeyVaultPassword []LinkedServiceSQLServerKeyVaultPasswordParameters `json:"keyVaultPassword,omitempty" tf:"key_vault_password,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The on-premises Windows authentication user name.
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

// LinkedServiceSQLServerSpec defines the desired state of LinkedServiceSQLServer
type LinkedServiceSQLServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceSQLServerParameters `json:"forProvider"`
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
	InitProvider LinkedServiceSQLServerInitParameters `json:"initProvider,omitempty"`
}

// LinkedServiceSQLServerStatus defines the observed state of LinkedServiceSQLServer.
type LinkedServiceSQLServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceSQLServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LinkedServiceSQLServer is the Schema for the LinkedServiceSQLServers API. Manages a Linked Service (connection) between a SQL Server and Azure Data Factory.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedServiceSQLServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinkedServiceSQLServerSpec   `json:"spec"`
	Status            LinkedServiceSQLServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceSQLServerList contains a list of LinkedServiceSQLServers
type LinkedServiceSQLServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceSQLServer `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceSQLServer_Kind             = "LinkedServiceSQLServer"
	LinkedServiceSQLServer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceSQLServer_Kind}.String()
	LinkedServiceSQLServer_KindAPIVersion   = LinkedServiceSQLServer_Kind + "." + CRDGroupVersion.String()
	LinkedServiceSQLServer_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceSQLServer_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceSQLServer{}, &LinkedServiceSQLServerList{})
}
