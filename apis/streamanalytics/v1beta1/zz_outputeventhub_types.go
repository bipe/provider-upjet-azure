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

type OutputEventHubInitParameters struct {

	// The authentication mode for the Stream Output. Possible values are Msi and ConnectionString. Defaults to ConnectionString.
	AuthenticationMode *string `json:"authenticationMode,omitempty" tf:"authentication_mode,omitempty"`

	// The name of the Event Hub.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.EventHub
	EventHubName *string `json:"eventhubName,omitempty" tf:"eventhub_name,omitempty"`

	// Reference to a EventHub in eventhub to populate eventhubName.
	// +kubebuilder:validation:Optional
	EventHubNameRef *v1.Reference `json:"eventhubNameRef,omitempty" tf:"-"`

	// Selector for a EventHub in eventhub to populate eventhubName.
	// +kubebuilder:validation:Optional
	EventHubNameSelector *v1.Selector `json:"eventhubNameSelector,omitempty" tf:"-"`

	// The column that is used for the Event Hub partition key.
	PartitionKey *string `json:"partitionKey,omitempty" tf:"partition_key,omitempty"`

	// A list of property columns to add to the Event Hub output.
	PropertyColumns []*string `json:"propertyColumns,omitempty" tf:"property_columns,omitempty"`

	// A serialization block as defined below.
	Serialization []OutputEventHubSerializationInitParameters `json:"serialization,omitempty" tf:"serialization,omitempty"`

	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.EventHubNamespace
	ServiceBusNamespace *string `json:"servicebusNamespace,omitempty" tf:"servicebus_namespace,omitempty"`

	// Reference to a EventHubNamespace in eventhub to populate servicebusNamespace.
	// +kubebuilder:validation:Optional
	ServiceBusNamespaceRef *v1.Reference `json:"servicebusNamespaceRef,omitempty" tf:"-"`

	// Selector for a EventHubNamespace in eventhub to populate servicebusNamespace.
	// +kubebuilder:validation:Optional
	ServiceBusNamespaceSelector *v1.Selector `json:"servicebusNamespaceSelector,omitempty" tf:"-"`

	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc. Required when authentication_mode is set to ConnectionString.
	SharedAccessPolicyName *string `json:"sharedAccessPolicyName,omitempty" tf:"shared_access_policy_name,omitempty"`
}

type OutputEventHubObservation struct {

	// The authentication mode for the Stream Output. Possible values are Msi and ConnectionString. Defaults to ConnectionString.
	AuthenticationMode *string `json:"authenticationMode,omitempty" tf:"authentication_mode,omitempty"`

	// The name of the Event Hub.
	EventHubName *string `json:"eventhubName,omitempty" tf:"eventhub_name,omitempty"`

	// The ID of the Stream Analytics Output EventHub.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The column that is used for the Event Hub partition key.
	PartitionKey *string `json:"partitionKey,omitempty" tf:"partition_key,omitempty"`

	// A list of property columns to add to the Event Hub output.
	PropertyColumns []*string `json:"propertyColumns,omitempty" tf:"property_columns,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A serialization block as defined below.
	Serialization []OutputEventHubSerializationObservation `json:"serialization,omitempty" tf:"serialization,omitempty"`

	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	ServiceBusNamespace *string `json:"servicebusNamespace,omitempty" tf:"servicebus_namespace,omitempty"`

	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc. Required when authentication_mode is set to ConnectionString.
	SharedAccessPolicyName *string `json:"sharedAccessPolicyName,omitempty" tf:"shared_access_policy_name,omitempty"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`
}

type OutputEventHubParameters struct {

	// The authentication mode for the Stream Output. Possible values are Msi and ConnectionString. Defaults to ConnectionString.
	// +kubebuilder:validation:Optional
	AuthenticationMode *string `json:"authenticationMode,omitempty" tf:"authentication_mode,omitempty"`

	// The name of the Event Hub.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.EventHub
	// +kubebuilder:validation:Optional
	EventHubName *string `json:"eventhubName,omitempty" tf:"eventhub_name,omitempty"`

	// Reference to a EventHub in eventhub to populate eventhubName.
	// +kubebuilder:validation:Optional
	EventHubNameRef *v1.Reference `json:"eventhubNameRef,omitempty" tf:"-"`

	// Selector for a EventHub in eventhub to populate eventhubName.
	// +kubebuilder:validation:Optional
	EventHubNameSelector *v1.Selector `json:"eventhubNameSelector,omitempty" tf:"-"`

	// The column that is used for the Event Hub partition key.
	// +kubebuilder:validation:Optional
	PartitionKey *string `json:"partitionKey,omitempty" tf:"partition_key,omitempty"`

	// A list of property columns to add to the Event Hub output.
	// +kubebuilder:validation:Optional
	PropertyColumns []*string `json:"propertyColumns,omitempty" tf:"property_columns,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A serialization block as defined below.
	// +kubebuilder:validation:Optional
	Serialization []OutputEventHubSerializationParameters `json:"serialization,omitempty" tf:"serialization,omitempty"`

	// The namespace that is associated with the desired Event Hub, Service Bus Queue, Service Bus Topic, etc.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.EventHubNamespace
	// +kubebuilder:validation:Optional
	ServiceBusNamespace *string `json:"servicebusNamespace,omitempty" tf:"servicebus_namespace,omitempty"`

	// Reference to a EventHubNamespace in eventhub to populate servicebusNamespace.
	// +kubebuilder:validation:Optional
	ServiceBusNamespaceRef *v1.Reference `json:"servicebusNamespaceRef,omitempty" tf:"-"`

	// Selector for a EventHubNamespace in eventhub to populate servicebusNamespace.
	// +kubebuilder:validation:Optional
	ServiceBusNamespaceSelector *v1.Selector `json:"servicebusNamespaceSelector,omitempty" tf:"-"`

	// The shared access policy key for the specified shared access policy. Required when authentication_mode is set to ConnectionString.
	// +kubebuilder:validation:Optional
	SharedAccessPolicyKeySecretRef *v1.SecretKeySelector `json:"sharedAccessPolicyKeySecretRef,omitempty" tf:"-"`

	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc. Required when authentication_mode is set to ConnectionString.
	// +kubebuilder:validation:Optional
	SharedAccessPolicyName *string `json:"sharedAccessPolicyName,omitempty" tf:"shared_access_policy_name,omitempty"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name,omitempty"`
}

type OutputEventHubSerializationInitParameters struct {

	// The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to UTF8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are   (space), , (comma), 	 (tab), | (pipe) and ;.
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// Specifies the format of the JSON the output will be written in. Possible values are Array and LineSeparated.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The serialization format used for outgoing data streams. Possible values are Avro, Csv, Json and Parquet.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type OutputEventHubSerializationObservation struct {

	// The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to UTF8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are   (space), , (comma), 	 (tab), | (pipe) and ;.
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// Specifies the format of the JSON the output will be written in. Possible values are Array and LineSeparated.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The serialization format used for outgoing data streams. Possible values are Avro, Csv, Json and Parquet.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type OutputEventHubSerializationParameters struct {

	// The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to UTF8.
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are   (space), , (comma), 	 (tab), | (pipe) and ;.
	// +kubebuilder:validation:Optional
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// Specifies the format of the JSON the output will be written in. Possible values are Array and LineSeparated.
	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The serialization format used for outgoing data streams. Possible values are Avro, Csv, Json and Parquet.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// OutputEventHubSpec defines the desired state of OutputEventHub
type OutputEventHubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OutputEventHubParameters `json:"forProvider"`
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
	InitProvider OutputEventHubInitParameters `json:"initProvider,omitempty"`
}

// OutputEventHubStatus defines the observed state of OutputEventHub.
type OutputEventHubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OutputEventHubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OutputEventHub is the Schema for the OutputEventHubs API. Manages a Stream Analytics Output to an EventHub.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type OutputEventHub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serialization) || (has(self.initProvider) && has(self.initProvider.serialization))",message="spec.forProvider.serialization is a required parameter"
	Spec   OutputEventHubSpec   `json:"spec"`
	Status OutputEventHubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OutputEventHubList contains a list of OutputEventHubs
type OutputEventHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OutputEventHub `json:"items"`
}

// Repository type metadata.
var (
	OutputEventHub_Kind             = "OutputEventHub"
	OutputEventHub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OutputEventHub_Kind}.String()
	OutputEventHub_KindAPIVersion   = OutputEventHub_Kind + "." + CRDGroupVersion.String()
	OutputEventHub_GroupVersionKind = CRDGroupVersion.WithKind(OutputEventHub_Kind)
)

func init() {
	SchemeBuilder.Register(&OutputEventHub{}, &OutputEventHubList{})
}
