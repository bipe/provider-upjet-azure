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

type HealthcareMedtechServiceIdentityInitParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Healthcare Med Tech Service.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Healthcare Med Tech Service. Possible values are SystemAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HealthcareMedtechServiceIdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Healthcare Med Tech Service.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID associated with this System Assigned Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this System Assigned Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Healthcare Med Tech Service. Possible values are SystemAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HealthcareMedtechServiceIdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Healthcare Med Tech Service.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Healthcare Med Tech Service. Possible values are SystemAssigned.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type HealthcareMedtechServiceInitParameters struct {

	// Specifies the Device Mappings of the Med Tech Service.
	DeviceMappingJSON *string `json:"deviceMappingJson,omitempty" tf:"device_mapping_json,omitempty"`

	// Specifies the Consumer Group of the Event Hub to connect to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.ConsumerGroup
	EventHubConsumerGroupName *string `json:"eventhubConsumerGroupName,omitempty" tf:"eventhub_consumer_group_name,omitempty"`

	// Reference to a ConsumerGroup in eventhub to populate eventhubConsumerGroupName.
	// +kubebuilder:validation:Optional
	EventHubConsumerGroupNameRef *v1.Reference `json:"eventhubConsumerGroupNameRef,omitempty" tf:"-"`

	// Selector for a ConsumerGroup in eventhub to populate eventhubConsumerGroupName.
	// +kubebuilder:validation:Optional
	EventHubConsumerGroupNameSelector *v1.Selector `json:"eventhubConsumerGroupNameSelector,omitempty" tf:"-"`

	// Specifies the name of the Event Hub to connect to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.EventHub
	EventHubName *string `json:"eventhubName,omitempty" tf:"eventhub_name,omitempty"`

	// Reference to a EventHub in eventhub to populate eventhubName.
	// +kubebuilder:validation:Optional
	EventHubNameRef *v1.Reference `json:"eventhubNameRef,omitempty" tf:"-"`

	// Selector for a EventHub in eventhub to populate eventhubName.
	// +kubebuilder:validation:Optional
	EventHubNameSelector *v1.Selector `json:"eventhubNameSelector,omitempty" tf:"-"`

	// Specifies the namespace name of the Event Hub to connect to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.EventHubNamespace
	EventHubNamespaceName *string `json:"eventhubNamespaceName,omitempty" tf:"eventhub_namespace_name,omitempty"`

	// Reference to a EventHubNamespace in eventhub to populate eventhubNamespaceName.
	// +kubebuilder:validation:Optional
	EventHubNamespaceNameRef *v1.Reference `json:"eventhubNamespaceNameRef,omitempty" tf:"-"`

	// Selector for a EventHubNamespace in eventhub to populate eventhubNamespaceName.
	// +kubebuilder:validation:Optional
	EventHubNamespaceNameSelector *v1.Selector `json:"eventhubNamespaceNameSelector,omitempty" tf:"-"`

	// An identity block as defined below.
	Identity []HealthcareMedtechServiceIdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the Healthcare Med Tech Service should be created. Changing this forces a new Healthcare Med Tech Service to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags to assign to the Healthcare Med Tech Service.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type HealthcareMedtechServiceObservation struct {

	// Specifies the Device Mappings of the Med Tech Service.
	DeviceMappingJSON *string `json:"deviceMappingJson,omitempty" tf:"device_mapping_json,omitempty"`

	// Specifies the Consumer Group of the Event Hub to connect to.
	EventHubConsumerGroupName *string `json:"eventhubConsumerGroupName,omitempty" tf:"eventhub_consumer_group_name,omitempty"`

	// Specifies the name of the Event Hub to connect to.
	EventHubName *string `json:"eventhubName,omitempty" tf:"eventhub_name,omitempty"`

	// Specifies the namespace name of the Event Hub to connect to.
	EventHubNamespaceName *string `json:"eventhubNamespaceName,omitempty" tf:"eventhub_namespace_name,omitempty"`

	// The ID of the Healthcare Med Tech Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []HealthcareMedtechServiceIdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the Healthcare Med Tech Service should be created. Changing this forces a new Healthcare Med Tech Service to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags to assign to the Healthcare Med Tech Service.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the id of the Healthcare Workspace where the Healthcare Med Tech Service should exist. Changing this forces a new Healthcare Med Tech Service to be created.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type HealthcareMedtechServiceParameters struct {

	// Specifies the Device Mappings of the Med Tech Service.
	// +kubebuilder:validation:Optional
	DeviceMappingJSON *string `json:"deviceMappingJson,omitempty" tf:"device_mapping_json,omitempty"`

	// Specifies the Consumer Group of the Event Hub to connect to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.ConsumerGroup
	// +kubebuilder:validation:Optional
	EventHubConsumerGroupName *string `json:"eventhubConsumerGroupName,omitempty" tf:"eventhub_consumer_group_name,omitempty"`

	// Reference to a ConsumerGroup in eventhub to populate eventhubConsumerGroupName.
	// +kubebuilder:validation:Optional
	EventHubConsumerGroupNameRef *v1.Reference `json:"eventhubConsumerGroupNameRef,omitempty" tf:"-"`

	// Selector for a ConsumerGroup in eventhub to populate eventhubConsumerGroupName.
	// +kubebuilder:validation:Optional
	EventHubConsumerGroupNameSelector *v1.Selector `json:"eventhubConsumerGroupNameSelector,omitempty" tf:"-"`

	// Specifies the name of the Event Hub to connect to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.EventHub
	// +kubebuilder:validation:Optional
	EventHubName *string `json:"eventhubName,omitempty" tf:"eventhub_name,omitempty"`

	// Reference to a EventHub in eventhub to populate eventhubName.
	// +kubebuilder:validation:Optional
	EventHubNameRef *v1.Reference `json:"eventhubNameRef,omitempty" tf:"-"`

	// Selector for a EventHub in eventhub to populate eventhubName.
	// +kubebuilder:validation:Optional
	EventHubNameSelector *v1.Selector `json:"eventhubNameSelector,omitempty" tf:"-"`

	// Specifies the namespace name of the Event Hub to connect to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.EventHubNamespace
	// +kubebuilder:validation:Optional
	EventHubNamespaceName *string `json:"eventhubNamespaceName,omitempty" tf:"eventhub_namespace_name,omitempty"`

	// Reference to a EventHubNamespace in eventhub to populate eventhubNamespaceName.
	// +kubebuilder:validation:Optional
	EventHubNamespaceNameRef *v1.Reference `json:"eventhubNamespaceNameRef,omitempty" tf:"-"`

	// Selector for a EventHubNamespace in eventhub to populate eventhubNamespaceName.
	// +kubebuilder:validation:Optional
	EventHubNamespaceNameSelector *v1.Selector `json:"eventhubNamespaceNameSelector,omitempty" tf:"-"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []HealthcareMedtechServiceIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the Healthcare Med Tech Service should be created. Changing this forces a new Healthcare Med Tech Service to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags to assign to the Healthcare Med Tech Service.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the id of the Healthcare Workspace where the Healthcare Med Tech Service should exist. Changing this forces a new Healthcare Med Tech Service to be created.
	// +crossplane:generate:reference:type=HealthcareWorkspace
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a HealthcareWorkspace to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a HealthcareWorkspace to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`
}

// HealthcareMedtechServiceSpec defines the desired state of HealthcareMedtechService
type HealthcareMedtechServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HealthcareMedtechServiceParameters `json:"forProvider"`
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
	InitProvider HealthcareMedtechServiceInitParameters `json:"initProvider,omitempty"`
}

// HealthcareMedtechServiceStatus defines the observed state of HealthcareMedtechService.
type HealthcareMedtechServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HealthcareMedtechServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HealthcareMedtechService is the Schema for the HealthcareMedtechServices API. Manages a Healthcare MedTech (Internet of Medical Things) devices Service.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HealthcareMedtechService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.deviceMappingJson) || (has(self.initProvider) && has(self.initProvider.deviceMappingJson))",message="spec.forProvider.deviceMappingJson is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   HealthcareMedtechServiceSpec   `json:"spec"`
	Status HealthcareMedtechServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HealthcareMedtechServiceList contains a list of HealthcareMedtechServices
type HealthcareMedtechServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HealthcareMedtechService `json:"items"`
}

// Repository type metadata.
var (
	HealthcareMedtechService_Kind             = "HealthcareMedtechService"
	HealthcareMedtechService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HealthcareMedtechService_Kind}.String()
	HealthcareMedtechService_KindAPIVersion   = HealthcareMedtechService_Kind + "." + CRDGroupVersion.String()
	HealthcareMedtechService_GroupVersionKind = CRDGroupVersion.WithKind(HealthcareMedtechService_Kind)
)

func init() {
	SchemeBuilder.Register(&HealthcareMedtechService{}, &HealthcareMedtechServiceList{})
}
