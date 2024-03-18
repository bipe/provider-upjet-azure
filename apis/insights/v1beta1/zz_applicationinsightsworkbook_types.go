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

type ApplicationInsightsWorkbookInitParameters struct {

	// Workbook category, as defined by the user at creation time. There may be additional category types beyond the following: workbook, sentinel. Defaults to workbook.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// Configuration of this particular workbook. Configuration data is a string containing valid JSON.
	DataJSON *string `json:"dataJson,omitempty" tf:"data_json,omitempty"`

	// Specifies the description of the workbook.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the user-defined name (display name) of the workbook.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// An identity block as defined below. Changing this forces a new Workbook to be created.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the Workbook should exist. Changing this forces a new Workbook to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of this Workbook as a UUID/GUID. It should not contain any uppercase letters. Changing this forces a new Workbook to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the name of the Resource Group where the Workbook should exist. Changing this forces a new Workbook to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Resource ID for a source resource. It should not contain any uppercase letters. Defaults to azure monitor.
	SourceID *string `json:"sourceId,omitempty" tf:"source_id,omitempty"`

	// Specifies the Resource Manager ID of the Storage Container when bring your own storage is used. Changing this forces a new Workbook to be created.
	StorageContainerID *string `json:"storageContainerId,omitempty" tf:"storage_container_id,omitempty"`

	// A mapping of tags which should be assigned to the Workbook.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ApplicationInsightsWorkbookObservation struct {

	// Workbook category, as defined by the user at creation time. There may be additional category types beyond the following: workbook, sentinel. Defaults to workbook.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// Configuration of this particular workbook. Configuration data is a string containing valid JSON.
	DataJSON *string `json:"dataJson,omitempty" tf:"data_json,omitempty"`

	// Specifies the description of the workbook.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the user-defined name (display name) of the workbook.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the Workbook.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below. Changing this forces a new Workbook to be created.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the Workbook should exist. Changing this forces a new Workbook to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of this Workbook as a UUID/GUID. It should not contain any uppercase letters. Changing this forces a new Workbook to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the name of the Resource Group where the Workbook should exist. Changing this forces a new Workbook to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Resource ID for a source resource. It should not contain any uppercase letters. Defaults to azure monitor.
	SourceID *string `json:"sourceId,omitempty" tf:"source_id,omitempty"`

	// Specifies the Resource Manager ID of the Storage Container when bring your own storage is used. Changing this forces a new Workbook to be created.
	StorageContainerID *string `json:"storageContainerId,omitempty" tf:"storage_container_id,omitempty"`

	// A mapping of tags which should be assigned to the Workbook.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ApplicationInsightsWorkbookParameters struct {

	// Workbook category, as defined by the user at creation time. There may be additional category types beyond the following: workbook, sentinel. Defaults to workbook.
	// +kubebuilder:validation:Optional
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// Configuration of this particular workbook. Configuration data is a string containing valid JSON.
	// +kubebuilder:validation:Optional
	DataJSON *string `json:"dataJson,omitempty" tf:"data_json,omitempty"`

	// Specifies the description of the workbook.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the user-defined name (display name) of the workbook.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// An identity block as defined below. Changing this forces a new Workbook to be created.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the Workbook should exist. Changing this forces a new Workbook to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of this Workbook as a UUID/GUID. It should not contain any uppercase letters. Changing this forces a new Workbook to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the name of the Resource Group where the Workbook should exist. Changing this forces a new Workbook to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Resource ID for a source resource. It should not contain any uppercase letters. Defaults to azure monitor.
	// +kubebuilder:validation:Optional
	SourceID *string `json:"sourceId,omitempty" tf:"source_id,omitempty"`

	// Specifies the Resource Manager ID of the Storage Container when bring your own storage is used. Changing this forces a new Workbook to be created.
	// +kubebuilder:validation:Optional
	StorageContainerID *string `json:"storageContainerId,omitempty" tf:"storage_container_id,omitempty"`

	// A mapping of tags which should be assigned to the Workbook.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IdentityInitParameters struct {

	// The list of User Assigned Managed Identity IDs assigned to this Workbook. Changing this forces a new resource to be created.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The type of Managed Service Identity that is configured on this Workbook. Possible values are UserAssigned, SystemAssigned and SystemAssigned, UserAssigned. Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// The list of User Assigned Managed Identity IDs assigned to this Workbook. Changing this forces a new resource to be created.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID of the System Assigned Managed Service Identity that is configured on this Workbook.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID of the System Assigned Managed Service Identity that is configured on this Workbook.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// The type of Managed Service Identity that is configured on this Workbook. Possible values are UserAssigned, SystemAssigned and SystemAssigned, UserAssigned. Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// The list of User Assigned Managed Identity IDs assigned to this Workbook. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The type of Managed Service Identity that is configured on this Workbook. Possible values are UserAssigned, SystemAssigned and SystemAssigned, UserAssigned. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// ApplicationInsightsWorkbookSpec defines the desired state of ApplicationInsightsWorkbook
type ApplicationInsightsWorkbookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationInsightsWorkbookParameters `json:"forProvider"`
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
	InitProvider ApplicationInsightsWorkbookInitParameters `json:"initProvider,omitempty"`
}

// ApplicationInsightsWorkbookStatus defines the observed state of ApplicationInsightsWorkbook.
type ApplicationInsightsWorkbookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationInsightsWorkbookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ApplicationInsightsWorkbook is the Schema for the ApplicationInsightsWorkbooks API. Manages an Azure Workbook.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApplicationInsightsWorkbook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataJson) || (has(self.initProvider) && has(self.initProvider.dataJson))",message="spec.forProvider.dataJson is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ApplicationInsightsWorkbookSpec   `json:"spec"`
	Status ApplicationInsightsWorkbookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationInsightsWorkbookList contains a list of ApplicationInsightsWorkbooks
type ApplicationInsightsWorkbookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationInsightsWorkbook `json:"items"`
}

// Repository type metadata.
var (
	ApplicationInsightsWorkbook_Kind             = "ApplicationInsightsWorkbook"
	ApplicationInsightsWorkbook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationInsightsWorkbook_Kind}.String()
	ApplicationInsightsWorkbook_KindAPIVersion   = ApplicationInsightsWorkbook_Kind + "." + CRDGroupVersion.String()
	ApplicationInsightsWorkbook_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationInsightsWorkbook_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationInsightsWorkbook{}, &ApplicationInsightsWorkbookList{})
}
