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

type CustomPersistentDiskInitParameters struct {

	// These are the mount options for a persistent disk.
	// +listType=set
	MountOptions []*string `json:"mountOptions,omitempty" tf:"mount_options,omitempty"`

	// The mount path of the persistent disk.
	MountPath *string `json:"mountPath,omitempty" tf:"mount_path,omitempty"`

	// Indicates whether the persistent disk is a readOnly one.
	ReadOnlyEnabled *bool `json:"readOnlyEnabled,omitempty" tf:"read_only_enabled,omitempty"`

	// The share name of the Azure File share.
	ShareName *string `json:"shareName,omitempty" tf:"share_name,omitempty"`

	// The name of the Spring Cloud Storage.
	StorageName *string `json:"storageName,omitempty" tf:"storage_name,omitempty"`
}

type CustomPersistentDiskObservation struct {

	// These are the mount options for a persistent disk.
	// +listType=set
	MountOptions []*string `json:"mountOptions,omitempty" tf:"mount_options,omitempty"`

	// The mount path of the persistent disk.
	MountPath *string `json:"mountPath,omitempty" tf:"mount_path,omitempty"`

	// Indicates whether the persistent disk is a readOnly one.
	ReadOnlyEnabled *bool `json:"readOnlyEnabled,omitempty" tf:"read_only_enabled,omitempty"`

	// The share name of the Azure File share.
	ShareName *string `json:"shareName,omitempty" tf:"share_name,omitempty"`

	// The name of the Spring Cloud Storage.
	StorageName *string `json:"storageName,omitempty" tf:"storage_name,omitempty"`
}

type CustomPersistentDiskParameters struct {

	// These are the mount options for a persistent disk.
	// +kubebuilder:validation:Optional
	// +listType=set
	MountOptions []*string `json:"mountOptions,omitempty" tf:"mount_options,omitempty"`

	// The mount path of the persistent disk.
	// +kubebuilder:validation:Optional
	MountPath *string `json:"mountPath" tf:"mount_path,omitempty"`

	// Indicates whether the persistent disk is a readOnly one.
	// +kubebuilder:validation:Optional
	ReadOnlyEnabled *bool `json:"readOnlyEnabled,omitempty" tf:"read_only_enabled,omitempty"`

	// The share name of the Azure File share.
	// +kubebuilder:validation:Optional
	ShareName *string `json:"shareName" tf:"share_name,omitempty"`

	// The name of the Spring Cloud Storage.
	// +kubebuilder:validation:Optional
	StorageName *string `json:"storageName" tf:"storage_name,omitempty"`
}

type IdentityInitParameters struct {

	// A list of User Assigned Managed Identity IDs to be assigned to this Spring Cloud Application.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Spring Cloud Application. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// A list of User Assigned Managed Identity IDs to be assigned to this Spring Cloud Application.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID for the Service Principal associated with the Managed Service Identity of this Spring Cloud Application.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID for the Service Principal associated with the Managed Service Identity of this Spring Cloud Application.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Spring Cloud Application. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// A list of User Assigned Managed Identity IDs to be assigned to this Spring Cloud Application.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Spring Cloud Application. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type IngressSettingsInitParameters struct {

	// Specifies how ingress should communicate with this app backend service. Allowed values are GRPC and Default. Defaults to Default.
	BackendProtocol *string `json:"backendProtocol,omitempty" tf:"backend_protocol,omitempty"`

	// Specifies the ingress read time out in seconds. Defaults to 300.
	ReadTimeoutInSeconds *float64 `json:"readTimeoutInSeconds,omitempty" tf:"read_timeout_in_seconds,omitempty"`

	// Specifies the ingress send time out in seconds. Defaults to 60.
	SendTimeoutInSeconds *float64 `json:"sendTimeoutInSeconds,omitempty" tf:"send_timeout_in_seconds,omitempty"`

	// Specifies the type of the affinity, set this to Cookie to enable session affinity. Allowed values are Cookie and None. Defaults to None.
	SessionAffinity *string `json:"sessionAffinity,omitempty" tf:"session_affinity,omitempty"`

	// Specifies the time in seconds until the cookie expires.
	SessionCookieMaxAge *float64 `json:"sessionCookieMaxAge,omitempty" tf:"session_cookie_max_age,omitempty"`
}

type IngressSettingsObservation struct {

	// Specifies how ingress should communicate with this app backend service. Allowed values are GRPC and Default. Defaults to Default.
	BackendProtocol *string `json:"backendProtocol,omitempty" tf:"backend_protocol,omitempty"`

	// Specifies the ingress read time out in seconds. Defaults to 300.
	ReadTimeoutInSeconds *float64 `json:"readTimeoutInSeconds,omitempty" tf:"read_timeout_in_seconds,omitempty"`

	// Specifies the ingress send time out in seconds. Defaults to 60.
	SendTimeoutInSeconds *float64 `json:"sendTimeoutInSeconds,omitempty" tf:"send_timeout_in_seconds,omitempty"`

	// Specifies the type of the affinity, set this to Cookie to enable session affinity. Allowed values are Cookie and None. Defaults to None.
	SessionAffinity *string `json:"sessionAffinity,omitempty" tf:"session_affinity,omitempty"`

	// Specifies the time in seconds until the cookie expires.
	SessionCookieMaxAge *float64 `json:"sessionCookieMaxAge,omitempty" tf:"session_cookie_max_age,omitempty"`
}

type IngressSettingsParameters struct {

	// Specifies how ingress should communicate with this app backend service. Allowed values are GRPC and Default. Defaults to Default.
	// +kubebuilder:validation:Optional
	BackendProtocol *string `json:"backendProtocol,omitempty" tf:"backend_protocol,omitempty"`

	// Specifies the ingress read time out in seconds. Defaults to 300.
	// +kubebuilder:validation:Optional
	ReadTimeoutInSeconds *float64 `json:"readTimeoutInSeconds,omitempty" tf:"read_timeout_in_seconds,omitempty"`

	// Specifies the ingress send time out in seconds. Defaults to 60.
	// +kubebuilder:validation:Optional
	SendTimeoutInSeconds *float64 `json:"sendTimeoutInSeconds,omitempty" tf:"send_timeout_in_seconds,omitempty"`

	// Specifies the type of the affinity, set this to Cookie to enable session affinity. Allowed values are Cookie and None. Defaults to None.
	// +kubebuilder:validation:Optional
	SessionAffinity *string `json:"sessionAffinity,omitempty" tf:"session_affinity,omitempty"`

	// Specifies the time in seconds until the cookie expires.
	// +kubebuilder:validation:Optional
	SessionCookieMaxAge *float64 `json:"sessionCookieMaxAge,omitempty" tf:"session_cookie_max_age,omitempty"`
}

type PersistentDiskInitParameters struct {

	// Specifies the mount path of the persistent disk. Defaults to /persistent.
	MountPath *string `json:"mountPath,omitempty" tf:"mount_path,omitempty"`

	// Specifies the size of the persistent disk in GB. Possible values are between 0 and 50.
	SizeInGb *float64 `json:"sizeInGb,omitempty" tf:"size_in_gb,omitempty"`
}

type PersistentDiskObservation struct {

	// Specifies the mount path of the persistent disk. Defaults to /persistent.
	MountPath *string `json:"mountPath,omitempty" tf:"mount_path,omitempty"`

	// Specifies the size of the persistent disk in GB. Possible values are between 0 and 50.
	SizeInGb *float64 `json:"sizeInGb,omitempty" tf:"size_in_gb,omitempty"`
}

type PersistentDiskParameters struct {

	// Specifies the mount path of the persistent disk. Defaults to /persistent.
	// +kubebuilder:validation:Optional
	MountPath *string `json:"mountPath,omitempty" tf:"mount_path,omitempty"`

	// Specifies the size of the persistent disk in GB. Possible values are between 0 and 50.
	// +kubebuilder:validation:Optional
	SizeInGb *float64 `json:"sizeInGb" tf:"size_in_gb,omitempty"`
}

type SpringCloudAppInitParameters struct {

	// A JSON object that contains the addon configurations of the Spring Cloud Service.
	AddonJSON *string `json:"addonJson,omitempty" tf:"addon_json,omitempty"`

	// A custom_persistent_disk block as defined below.
	CustomPersistentDisk []CustomPersistentDiskInitParameters `json:"customPersistentDisk,omitempty" tf:"custom_persistent_disk,omitempty"`

	// Is only HTTPS allowed? Defaults to false.
	HTTPSOnly *bool `json:"httpsOnly,omitempty" tf:"https_only,omitempty"`

	// An identity block as defined below.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// An ingress_settings block as defined below.
	IngressSettings []IngressSettingsInitParameters `json:"ingressSettings,omitempty" tf:"ingress_settings,omitempty"`

	// Does the Spring Cloud Application have public endpoint? Defaults to false.
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// An persistent_disk block as defined below.
	PersistentDisk []PersistentDiskInitParameters `json:"persistentDisk,omitempty" tf:"persistent_disk,omitempty"`

	// Should the App in vnet injection instance exposes endpoint which could be accessed from Internet?
	PublicEndpointEnabled *bool `json:"publicEndpointEnabled,omitempty" tf:"public_endpoint_enabled,omitempty"`

	// Is End to End TLS Enabled? Defaults to false.
	TLSEnabled *bool `json:"tlsEnabled,omitempty" tf:"tls_enabled,omitempty"`
}

type SpringCloudAppObservation struct {

	// A JSON object that contains the addon configurations of the Spring Cloud Service.
	AddonJSON *string `json:"addonJson,omitempty" tf:"addon_json,omitempty"`

	// A custom_persistent_disk block as defined below.
	CustomPersistentDisk []CustomPersistentDiskObservation `json:"customPersistentDisk,omitempty" tf:"custom_persistent_disk,omitempty"`

	// The Fully Qualified DNS Name of the Spring Application in the service.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// Is only HTTPS allowed? Defaults to false.
	HTTPSOnly *bool `json:"httpsOnly,omitempty" tf:"https_only,omitempty"`

	// The ID of the Spring Cloud Application.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// An ingress_settings block as defined below.
	IngressSettings []IngressSettingsObservation `json:"ingressSettings,omitempty" tf:"ingress_settings,omitempty"`

	// Does the Spring Cloud Application have public endpoint? Defaults to false.
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// An persistent_disk block as defined below.
	PersistentDisk []PersistentDiskObservation `json:"persistentDisk,omitempty" tf:"persistent_disk,omitempty"`

	// Should the App in vnet injection instance exposes endpoint which could be accessed from Internet?
	PublicEndpointEnabled *bool `json:"publicEndpointEnabled,omitempty" tf:"public_endpoint_enabled,omitempty"`

	// Specifies the name of the resource group in which to create the Spring Cloud Application. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Is End to End TLS Enabled? Defaults to false.
	TLSEnabled *bool `json:"tlsEnabled,omitempty" tf:"tls_enabled,omitempty"`

	// The public endpoint of the Spring Cloud Application.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type SpringCloudAppParameters struct {

	// A JSON object that contains the addon configurations of the Spring Cloud Service.
	// +kubebuilder:validation:Optional
	AddonJSON *string `json:"addonJson,omitempty" tf:"addon_json,omitempty"`

	// A custom_persistent_disk block as defined below.
	// +kubebuilder:validation:Optional
	CustomPersistentDisk []CustomPersistentDiskParameters `json:"customPersistentDisk,omitempty" tf:"custom_persistent_disk,omitempty"`

	// Is only HTTPS allowed? Defaults to false.
	// +kubebuilder:validation:Optional
	HTTPSOnly *bool `json:"httpsOnly,omitempty" tf:"https_only,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// An ingress_settings block as defined below.
	// +kubebuilder:validation:Optional
	IngressSettings []IngressSettingsParameters `json:"ingressSettings,omitempty" tf:"ingress_settings,omitempty"`

	// Does the Spring Cloud Application have public endpoint? Defaults to false.
	// +kubebuilder:validation:Optional
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// An persistent_disk block as defined below.
	// +kubebuilder:validation:Optional
	PersistentDisk []PersistentDiskParameters `json:"persistentDisk,omitempty" tf:"persistent_disk,omitempty"`

	// Should the App in vnet injection instance exposes endpoint which could be accessed from Internet?
	// +kubebuilder:validation:Optional
	PublicEndpointEnabled *bool `json:"publicEndpointEnabled,omitempty" tf:"public_endpoint_enabled,omitempty"`

	// Specifies the name of the resource group in which to create the Spring Cloud Application. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudService
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Reference to a SpringCloudService in appplatform to populate serviceName.
	// +kubebuilder:validation:Optional
	ServiceNameRef *v1.Reference `json:"serviceNameRef,omitempty" tf:"-"`

	// Selector for a SpringCloudService in appplatform to populate serviceName.
	// +kubebuilder:validation:Optional
	ServiceNameSelector *v1.Selector `json:"serviceNameSelector,omitempty" tf:"-"`

	// Is End to End TLS Enabled? Defaults to false.
	// +kubebuilder:validation:Optional
	TLSEnabled *bool `json:"tlsEnabled,omitempty" tf:"tls_enabled,omitempty"`
}

// SpringCloudAppSpec defines the desired state of SpringCloudApp
type SpringCloudAppSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudAppParameters `json:"forProvider"`
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
	InitProvider SpringCloudAppInitParameters `json:"initProvider,omitempty"`
}

// SpringCloudAppStatus defines the observed state of SpringCloudApp.
type SpringCloudAppStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudAppObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SpringCloudApp is the Schema for the SpringCloudApps API. Manage an Azure Spring Cloud Application.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpringCloudAppSpec   `json:"spec"`
	Status            SpringCloudAppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudAppList contains a list of SpringCloudApps
type SpringCloudAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudApp `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudApp_Kind             = "SpringCloudApp"
	SpringCloudApp_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudApp_Kind}.String()
	SpringCloudApp_KindAPIVersion   = SpringCloudApp_Kind + "." + CRDGroupVersion.String()
	SpringCloudApp_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudApp_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudApp{}, &SpringCloudAppList{})
}
