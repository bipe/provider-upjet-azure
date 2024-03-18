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

type IOTHubCertificateInitParameters struct {

	// Is the certificate verified? Defaults to false.
	IsVerified *bool `json:"isVerified,omitempty" tf:"is_verified,omitempty"`
}

type IOTHubCertificateObservation struct {

	// The ID of the IoTHub Certificate.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the IoTHub that this certificate will be attached to. Changing this forces a new resource to be created.
	IOTHubName *string `json:"iothubName,omitempty" tf:"iothub_name,omitempty"`

	// Is the certificate verified? Defaults to false.
	IsVerified *bool `json:"isVerified,omitempty" tf:"is_verified,omitempty"`

	// The name of the resource group under which the IotHub Certificate resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type IOTHubCertificateParameters struct {

	// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
	// +kubebuilder:validation:Optional
	CertificateContentSecretRef v1.SecretKeySelector `json:"certificateContentSecretRef" tf:"-"`

	// The name of the IoTHub that this certificate will be attached to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	// +kubebuilder:validation:Optional
	IOTHubName *string `json:"iothubName,omitempty" tf:"iothub_name,omitempty"`

	// Reference to a IOTHub in devices to populate iothubName.
	// +kubebuilder:validation:Optional
	IOTHubNameRef *v1.Reference `json:"iothubNameRef,omitempty" tf:"-"`

	// Selector for a IOTHub in devices to populate iothubName.
	// +kubebuilder:validation:Optional
	IOTHubNameSelector *v1.Selector `json:"iothubNameSelector,omitempty" tf:"-"`

	// Is the certificate verified? Defaults to false.
	// +kubebuilder:validation:Optional
	IsVerified *bool `json:"isVerified,omitempty" tf:"is_verified,omitempty"`

	// The name of the resource group under which the IotHub Certificate resource has to be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// IOTHubCertificateSpec defines the desired state of IOTHubCertificate
type IOTHubCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubCertificateParameters `json:"forProvider"`
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
	InitProvider IOTHubCertificateInitParameters `json:"initProvider,omitempty"`
}

// IOTHubCertificateStatus defines the observed state of IOTHubCertificate.
type IOTHubCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IOTHubCertificate is the Schema for the IOTHubCertificates API. Manages an IoTHub Certificate.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHubCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.certificateContentSecretRef)",message="spec.forProvider.certificateContentSecretRef is a required parameter"
	Spec   IOTHubCertificateSpec   `json:"spec"`
	Status IOTHubCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubCertificateList contains a list of IOTHubCertificates
type IOTHubCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubCertificate `json:"items"`
}

// Repository type metadata.
var (
	IOTHubCertificate_Kind             = "IOTHubCertificate"
	IOTHubCertificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubCertificate_Kind}.String()
	IOTHubCertificate_KindAPIVersion   = IOTHubCertificate_Kind + "." + CRDGroupVersion.String()
	IOTHubCertificate_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubCertificate_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubCertificate{}, &IOTHubCertificateList{})
}
