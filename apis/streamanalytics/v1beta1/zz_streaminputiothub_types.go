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

type StreamInputIOTHubInitParameters struct {

	// The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.ConsumerGroup
	EventHubConsumerGroupName *string `json:"eventhubConsumerGroupName,omitempty" tf:"eventhub_consumer_group_name,omitempty"`

	// Reference to a ConsumerGroup in eventhub to populate eventhubConsumerGroupName.
	// +kubebuilder:validation:Optional
	EventHubConsumerGroupNameRef *v1.Reference `json:"eventhubConsumerGroupNameRef,omitempty" tf:"-"`

	// Selector for a ConsumerGroup in eventhub to populate eventhubConsumerGroupName.
	// +kubebuilder:validation:Optional
	EventHubConsumerGroupNameSelector *v1.Selector `json:"eventhubConsumerGroupNameSelector,omitempty" tf:"-"`

	// The name or the URI of the IoT Hub.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	IOTHubNamespace *string `json:"iothubNamespace,omitempty" tf:"iothub_namespace,omitempty"`

	// Reference to a IOTHub in devices to populate iothubNamespace.
	// +kubebuilder:validation:Optional
	IOTHubNamespaceRef *v1.Reference `json:"iothubNamespaceRef,omitempty" tf:"-"`

	// Selector for a IOTHub in devices to populate iothubNamespace.
	// +kubebuilder:validation:Optional
	IOTHubNamespaceSelector *v1.Selector `json:"iothubNamespaceSelector,omitempty" tf:"-"`

	// The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A serialization block as defined below.
	Serialization []StreamInputIOTHubSerializationInitParameters `json:"serialization,omitempty" tf:"serialization,omitempty"`

	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName *string `json:"sharedAccessPolicyName,omitempty" tf:"shared_access_policy_name,omitempty"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Job
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// Reference to a Job to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameRef *v1.Reference `json:"streamAnalyticsJobNameRef,omitempty" tf:"-"`

	// Selector for a Job to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameSelector *v1.Selector `json:"streamAnalyticsJobNameSelector,omitempty" tf:"-"`
}

type StreamInputIOTHubObservation struct {

	// The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
	EventHubConsumerGroupName *string `json:"eventhubConsumerGroupName,omitempty" tf:"eventhub_consumer_group_name,omitempty"`

	// The ID of the Stream Analytics Stream Input IoTHub.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name or the URI of the IoT Hub.
	IOTHubNamespace *string `json:"iothubNamespace,omitempty" tf:"iothub_namespace,omitempty"`

	// The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A serialization block as defined below.
	Serialization []StreamInputIOTHubSerializationObservation `json:"serialization,omitempty" tf:"serialization,omitempty"`

	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName *string `json:"sharedAccessPolicyName,omitempty" tf:"shared_access_policy_name,omitempty"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`
}

type StreamInputIOTHubParameters struct {

	// The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.ConsumerGroup
	// +kubebuilder:validation:Optional
	EventHubConsumerGroupName *string `json:"eventhubConsumerGroupName,omitempty" tf:"eventhub_consumer_group_name,omitempty"`

	// Reference to a ConsumerGroup in eventhub to populate eventhubConsumerGroupName.
	// +kubebuilder:validation:Optional
	EventHubConsumerGroupNameRef *v1.Reference `json:"eventhubConsumerGroupNameRef,omitempty" tf:"-"`

	// Selector for a ConsumerGroup in eventhub to populate eventhubConsumerGroupName.
	// +kubebuilder:validation:Optional
	EventHubConsumerGroupNameSelector *v1.Selector `json:"eventhubConsumerGroupNameSelector,omitempty" tf:"-"`

	// The name or the URI of the IoT Hub.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	// +kubebuilder:validation:Optional
	IOTHubNamespace *string `json:"iothubNamespace,omitempty" tf:"iothub_namespace,omitempty"`

	// Reference to a IOTHub in devices to populate iothubNamespace.
	// +kubebuilder:validation:Optional
	IOTHubNamespaceRef *v1.Reference `json:"iothubNamespaceRef,omitempty" tf:"-"`

	// Selector for a IOTHub in devices to populate iothubNamespace.
	// +kubebuilder:validation:Optional
	IOTHubNamespaceSelector *v1.Selector `json:"iothubNamespaceSelector,omitempty" tf:"-"`

	// The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

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
	Serialization []StreamInputIOTHubSerializationParameters `json:"serialization,omitempty" tf:"serialization,omitempty"`

	// The shared access policy key for the specified shared access policy. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SharedAccessPolicyKeySecretRef v1.SecretKeySelector `json:"sharedAccessPolicyKeySecretRef" tf:"-"`

	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	// +kubebuilder:validation:Optional
	SharedAccessPolicyName *string `json:"sharedAccessPolicyName,omitempty" tf:"shared_access_policy_name,omitempty"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Job
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// Reference to a Job to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameRef *v1.Reference `json:"streamAnalyticsJobNameRef,omitempty" tf:"-"`

	// Selector for a Job to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameSelector *v1.Selector `json:"streamAnalyticsJobNameSelector,omitempty" tf:"-"`
}

type StreamInputIOTHubSerializationInitParameters struct {

	// The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to UTF8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are   (space), , (comma), 	 (tab), | (pipe) and ;.
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// The serialization format used for incoming data streams. Possible values are Avro, Csv and Json.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StreamInputIOTHubSerializationObservation struct {

	// The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to UTF8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are   (space), , (comma), 	 (tab), | (pipe) and ;.
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// The serialization format used for incoming data streams. Possible values are Avro, Csv and Json.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StreamInputIOTHubSerializationParameters struct {

	// The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to UTF8.
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are   (space), , (comma), 	 (tab), | (pipe) and ;.
	// +kubebuilder:validation:Optional
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// The serialization format used for incoming data streams. Possible values are Avro, Csv and Json.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// StreamInputIOTHubSpec defines the desired state of StreamInputIOTHub
type StreamInputIOTHubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StreamInputIOTHubParameters `json:"forProvider"`
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
	InitProvider StreamInputIOTHubInitParameters `json:"initProvider,omitempty"`
}

// StreamInputIOTHubStatus defines the observed state of StreamInputIOTHub.
type StreamInputIOTHubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StreamInputIOTHubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// StreamInputIOTHub is the Schema for the StreamInputIOTHubs API. Manages a Stream Analytics Stream Input IoTHub.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type StreamInputIOTHub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.endpoint) || (has(self.initProvider) && has(self.initProvider.endpoint))",message="spec.forProvider.endpoint is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serialization) || (has(self.initProvider) && has(self.initProvider.serialization))",message="spec.forProvider.serialization is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sharedAccessPolicyKeySecretRef)",message="spec.forProvider.sharedAccessPolicyKeySecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sharedAccessPolicyName) || (has(self.initProvider) && has(self.initProvider.sharedAccessPolicyName))",message="spec.forProvider.sharedAccessPolicyName is a required parameter"
	Spec   StreamInputIOTHubSpec   `json:"spec"`
	Status StreamInputIOTHubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StreamInputIOTHubList contains a list of StreamInputIOTHubs
type StreamInputIOTHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StreamInputIOTHub `json:"items"`
}

// Repository type metadata.
var (
	StreamInputIOTHub_Kind             = "StreamInputIOTHub"
	StreamInputIOTHub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StreamInputIOTHub_Kind}.String()
	StreamInputIOTHub_KindAPIVersion   = StreamInputIOTHub_Kind + "." + CRDGroupVersion.String()
	StreamInputIOTHub_GroupVersionKind = CRDGroupVersion.WithKind(StreamInputIOTHub_Kind)
)

func init() {
	SchemeBuilder.Register(&StreamInputIOTHub{}, &StreamInputIOTHubList{})
}
