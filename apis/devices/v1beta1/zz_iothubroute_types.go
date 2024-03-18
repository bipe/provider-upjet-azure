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

type IOTHubRouteInitParameters struct {

	// The condition that is evaluated to apply the routing rule. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language. Defaults to true.
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// Specifies whether a route is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The list of endpoints to which messages that satisfy the condition are routed. Currently only one endpoint is allowed.
	EndpointNames []*string `json:"endpointNames,omitempty" tf:"endpoint_names,omitempty"`

	// The source that the routing rule is to be applied to. Possible values include: DeviceConnectionStateEvents, DeviceJobLifecycleEvents, DeviceLifecycleEvents, DeviceMessages, DigitalTwinChangeEvents, Invalid, TwinChangeEvents.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type IOTHubRouteObservation struct {

	// The condition that is evaluated to apply the routing rule. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language. Defaults to true.
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// Specifies whether a route is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The list of endpoints to which messages that satisfy the condition are routed. Currently only one endpoint is allowed.
	EndpointNames []*string `json:"endpointNames,omitempty" tf:"endpoint_names,omitempty"`

	// The ID of the IoTHub Route.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the IoTHub to which this Route belongs. Changing this forces a new resource to be created.
	IOTHubName *string `json:"iothubName,omitempty" tf:"iothub_name,omitempty"`

	// The name of the resource group under which the IotHub Route resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The source that the routing rule is to be applied to. Possible values include: DeviceConnectionStateEvents, DeviceJobLifecycleEvents, DeviceLifecycleEvents, DeviceMessages, DigitalTwinChangeEvents, Invalid, TwinChangeEvents.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type IOTHubRouteParameters struct {

	// The condition that is evaluated to apply the routing rule. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language. Defaults to true.
	// +kubebuilder:validation:Optional
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// Specifies whether a route is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The list of endpoints to which messages that satisfy the condition are routed. Currently only one endpoint is allowed.
	// +kubebuilder:validation:Optional
	EndpointNames []*string `json:"endpointNames,omitempty" tf:"endpoint_names,omitempty"`

	// The name of the IoTHub to which this Route belongs. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	// +kubebuilder:validation:Optional
	IOTHubName *string `json:"iothubName,omitempty" tf:"iothub_name,omitempty"`

	// Reference to a IOTHub in devices to populate iothubName.
	// +kubebuilder:validation:Optional
	IOTHubNameRef *v1.Reference `json:"iothubNameRef,omitempty" tf:"-"`

	// Selector for a IOTHub in devices to populate iothubName.
	// +kubebuilder:validation:Optional
	IOTHubNameSelector *v1.Selector `json:"iothubNameSelector,omitempty" tf:"-"`

	// The name of the resource group under which the IotHub Route resource has to be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The source that the routing rule is to be applied to. Possible values include: DeviceConnectionStateEvents, DeviceJobLifecycleEvents, DeviceLifecycleEvents, DeviceMessages, DigitalTwinChangeEvents, Invalid, TwinChangeEvents.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

// IOTHubRouteSpec defines the desired state of IOTHubRoute
type IOTHubRouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubRouteParameters `json:"forProvider"`
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
	InitProvider IOTHubRouteInitParameters `json:"initProvider,omitempty"`
}

// IOTHubRouteStatus defines the observed state of IOTHubRoute.
type IOTHubRouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubRouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IOTHubRoute is the Schema for the IOTHubRoutes API. Manages an IotHub Route
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHubRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.endpointNames) || (has(self.initProvider) && has(self.initProvider.endpointNames))",message="spec.forProvider.endpointNames is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.source) || (has(self.initProvider) && has(self.initProvider.source))",message="spec.forProvider.source is a required parameter"
	Spec   IOTHubRouteSpec   `json:"spec"`
	Status IOTHubRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubRouteList contains a list of IOTHubRoutes
type IOTHubRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubRoute `json:"items"`
}

// Repository type metadata.
var (
	IOTHubRoute_Kind             = "IOTHubRoute"
	IOTHubRoute_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubRoute_Kind}.String()
	IOTHubRoute_KindAPIVersion   = IOTHubRoute_Kind + "." + CRDGroupVersion.String()
	IOTHubRoute_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubRoute_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubRoute{}, &IOTHubRouteList{})
}
