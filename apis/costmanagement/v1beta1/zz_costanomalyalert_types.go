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

type CostAnomalyAlertInitParameters struct {

	// The display name which should be used for this Cost Anomaly Alert.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Specifies a list of email addresses which the Anomaly Alerts are send to.
	// +listType=set
	EmailAddresses []*string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`

	// The email subject of the Cost Anomaly Alerts. Maximum length of the subject is 70.
	EmailSubject *string `json:"emailSubject,omitempty" tf:"email_subject,omitempty"`

	// The message of the Cost Anomaly Alert. Maximum length of the message is 250.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// The ID of the Subscription this Cost Anomaly Alert is scoped to. Changing this forces a new resource to be created. When not supplied this defaults to the subscription configured in the provider.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

type CostAnomalyAlertObservation struct {

	// The display name which should be used for this Cost Anomaly Alert.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Specifies a list of email addresses which the Anomaly Alerts are send to.
	// +listType=set
	EmailAddresses []*string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`

	// The email subject of the Cost Anomaly Alerts. Maximum length of the subject is 70.
	EmailSubject *string `json:"emailSubject,omitempty" tf:"email_subject,omitempty"`

	// The ID of the Cost Anomaly Alert.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The message of the Cost Anomaly Alert. Maximum length of the message is 250.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// The ID of the Subscription this Cost Anomaly Alert is scoped to. Changing this forces a new resource to be created. When not supplied this defaults to the subscription configured in the provider.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

type CostAnomalyAlertParameters struct {

	// The display name which should be used for this Cost Anomaly Alert.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Specifies a list of email addresses which the Anomaly Alerts are send to.
	// +kubebuilder:validation:Optional
	// +listType=set
	EmailAddresses []*string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`

	// The email subject of the Cost Anomaly Alerts. Maximum length of the subject is 70.
	// +kubebuilder:validation:Optional
	EmailSubject *string `json:"emailSubject,omitempty" tf:"email_subject,omitempty"`

	// The message of the Cost Anomaly Alert. Maximum length of the message is 250.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// The ID of the Subscription this Cost Anomaly Alert is scoped to. Changing this forces a new resource to be created. When not supplied this defaults to the subscription configured in the provider.
	// +kubebuilder:validation:Optional
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

// CostAnomalyAlertSpec defines the desired state of CostAnomalyAlert
type CostAnomalyAlertSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CostAnomalyAlertParameters `json:"forProvider"`
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
	InitProvider CostAnomalyAlertInitParameters `json:"initProvider,omitempty"`
}

// CostAnomalyAlertStatus defines the observed state of CostAnomalyAlert.
type CostAnomalyAlertStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CostAnomalyAlertObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CostAnomalyAlert is the Schema for the CostAnomalyAlerts API. Manages a Cost Anomaly Alert.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CostAnomalyAlert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.emailAddresses) || (has(self.initProvider) && has(self.initProvider.emailAddresses))",message="spec.forProvider.emailAddresses is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.emailSubject) || (has(self.initProvider) && has(self.initProvider.emailSubject))",message="spec.forProvider.emailSubject is a required parameter"
	Spec   CostAnomalyAlertSpec   `json:"spec"`
	Status CostAnomalyAlertStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CostAnomalyAlertList contains a list of CostAnomalyAlerts
type CostAnomalyAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CostAnomalyAlert `json:"items"`
}

// Repository type metadata.
var (
	CostAnomalyAlert_Kind             = "CostAnomalyAlert"
	CostAnomalyAlert_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CostAnomalyAlert_Kind}.String()
	CostAnomalyAlert_KindAPIVersion   = CostAnomalyAlert_Kind + "." + CRDGroupVersion.String()
	CostAnomalyAlert_GroupVersionKind = CRDGroupVersion.WithKind(CostAnomalyAlert_Kind)
)

func init() {
	SchemeBuilder.Register(&CostAnomalyAlert{}, &CostAnomalyAlertList{})
}
