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

type SubAccountTagRuleInitParameters struct {

	// The ID of the Logz Sub Account. Changing this forces a new Logz Sub Account Tag Rule to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/logz/v1beta1.SubAccount
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	LogzSubAccountID *string `json:"logzSubAccountId,omitempty" tf:"logz_sub_account_id,omitempty"`

	// Reference to a SubAccount in logz to populate logzSubAccountId.
	// +kubebuilder:validation:Optional
	LogzSubAccountIDRef *v1.Reference `json:"logzSubAccountIdRef,omitempty" tf:"-"`

	// Selector for a SubAccount in logz to populate logzSubAccountId.
	// +kubebuilder:validation:Optional
	LogzSubAccountIDSelector *v1.Selector `json:"logzSubAccountIdSelector,omitempty" tf:"-"`

	// Whether AAD logs should be sent to the Monitor resource?
	SendAADLogs *bool `json:"sendAadLogs,omitempty" tf:"send_aad_logs,omitempty"`

	// Whether activity logs from this Logz Sub Account Tag Rule should be sent to the Monitor resource?
	SendActivityLogs *bool `json:"sendActivityLogs,omitempty" tf:"send_activity_logs,omitempty"`

	// Whether subscription logs should be sent to the Monitor resource?
	SendSubscriptionLogs *bool `json:"sendSubscriptionLogs,omitempty" tf:"send_subscription_logs,omitempty"`

	// One or more (up to 10) tag_filter blocks as defined below.
	TagFilter []TagFilterInitParameters `json:"tagFilter,omitempty" tf:"tag_filter,omitempty"`
}

type SubAccountTagRuleObservation struct {

	// The ID of the Logz Sub Account Tag Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Logz Sub Account. Changing this forces a new Logz Sub Account Tag Rule to be created.
	LogzSubAccountID *string `json:"logzSubAccountId,omitempty" tf:"logz_sub_account_id,omitempty"`

	// Whether AAD logs should be sent to the Monitor resource?
	SendAADLogs *bool `json:"sendAadLogs,omitempty" tf:"send_aad_logs,omitempty"`

	// Whether activity logs from this Logz Sub Account Tag Rule should be sent to the Monitor resource?
	SendActivityLogs *bool `json:"sendActivityLogs,omitempty" tf:"send_activity_logs,omitempty"`

	// Whether subscription logs should be sent to the Monitor resource?
	SendSubscriptionLogs *bool `json:"sendSubscriptionLogs,omitempty" tf:"send_subscription_logs,omitempty"`

	// One or more (up to 10) tag_filter blocks as defined below.
	TagFilter []TagFilterObservation `json:"tagFilter,omitempty" tf:"tag_filter,omitempty"`
}

type SubAccountTagRuleParameters struct {

	// The ID of the Logz Sub Account. Changing this forces a new Logz Sub Account Tag Rule to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/logz/v1beta1.SubAccount
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LogzSubAccountID *string `json:"logzSubAccountId,omitempty" tf:"logz_sub_account_id,omitempty"`

	// Reference to a SubAccount in logz to populate logzSubAccountId.
	// +kubebuilder:validation:Optional
	LogzSubAccountIDRef *v1.Reference `json:"logzSubAccountIdRef,omitempty" tf:"-"`

	// Selector for a SubAccount in logz to populate logzSubAccountId.
	// +kubebuilder:validation:Optional
	LogzSubAccountIDSelector *v1.Selector `json:"logzSubAccountIdSelector,omitempty" tf:"-"`

	// Whether AAD logs should be sent to the Monitor resource?
	// +kubebuilder:validation:Optional
	SendAADLogs *bool `json:"sendAadLogs,omitempty" tf:"send_aad_logs,omitempty"`

	// Whether activity logs from this Logz Sub Account Tag Rule should be sent to the Monitor resource?
	// +kubebuilder:validation:Optional
	SendActivityLogs *bool `json:"sendActivityLogs,omitempty" tf:"send_activity_logs,omitempty"`

	// Whether subscription logs should be sent to the Monitor resource?
	// +kubebuilder:validation:Optional
	SendSubscriptionLogs *bool `json:"sendSubscriptionLogs,omitempty" tf:"send_subscription_logs,omitempty"`

	// One or more (up to 10) tag_filter blocks as defined below.
	// +kubebuilder:validation:Optional
	TagFilter []TagFilterParameters `json:"tagFilter,omitempty" tf:"tag_filter,omitempty"`
}

type TagFilterInitParameters struct {

	// The action is used to limit logs collection to include or exclude Azure resources with specific tags. Possible values are Include and Exclude. Note that the Exclude takes priority over the Include.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The name of the tag to match.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the tag to match.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagFilterObservation struct {

	// The action is used to limit logs collection to include or exclude Azure resources with specific tags. Possible values are Include and Exclude. Note that the Exclude takes priority over the Include.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The name of the tag to match.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the tag to match.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagFilterParameters struct {

	// The action is used to limit logs collection to include or exclude Azure resources with specific tags. Possible values are Include and Exclude. Note that the Exclude takes priority over the Include.
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// The name of the tag to match.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The value of the tag to match.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// SubAccountTagRuleSpec defines the desired state of SubAccountTagRule
type SubAccountTagRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubAccountTagRuleParameters `json:"forProvider"`
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
	InitProvider SubAccountTagRuleInitParameters `json:"initProvider,omitempty"`
}

// SubAccountTagRuleStatus defines the observed state of SubAccountTagRule.
type SubAccountTagRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubAccountTagRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SubAccountTagRule is the Schema for the SubAccountTagRules API. Manages a Logz Sub Account Tag Rule.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SubAccountTagRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubAccountTagRuleSpec   `json:"spec"`
	Status            SubAccountTagRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubAccountTagRuleList contains a list of SubAccountTagRules
type SubAccountTagRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubAccountTagRule `json:"items"`
}

// Repository type metadata.
var (
	SubAccountTagRule_Kind             = "SubAccountTagRule"
	SubAccountTagRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubAccountTagRule_Kind}.String()
	SubAccountTagRule_KindAPIVersion   = SubAccountTagRule_Kind + "." + CRDGroupVersion.String()
	SubAccountTagRule_GroupVersionKind = CRDGroupVersion.WithKind(SubAccountTagRule_Kind)
)

func init() {
	SchemeBuilder.Register(&SubAccountTagRule{}, &SubAccountTagRuleList{})
}
