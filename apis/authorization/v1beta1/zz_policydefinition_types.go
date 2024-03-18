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

type PolicyDefinitionInitParameters struct {

	// The description of the policy definition.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of the policy definition.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The id of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupID *string `json:"managementGroupId,omitempty" tf:"management_group_id,omitempty"`

	// The metadata for the policy definition. This is a JSON string representing additional metadata that should be stored with the policy definition.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The policy resource manager mode that allows you to specify which resource types will be evaluated. Possible values are All, Indexed, Microsoft.ContainerService.Data, Microsoft.CustomerLockbox.Data, Microsoft.DataCatalog.Data, Microsoft.KeyVault.Data, Microsoft.Kubernetes.Data, Microsoft.MachineLearningServices.Data, Microsoft.Network.Data and Microsoft.Synapse.Data.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Parameters for the policy definition. This field is a JSON string that allows you to parameterize your policy definition.
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The policy rule for the policy definition. This is a JSON string representing the rule that contains an if and a then block.
	PolicyRule *string `json:"policyRule,omitempty" tf:"policy_rule,omitempty"`

	// The policy type. Possible values are BuiltIn, Custom, NotSpecified and Static. Changing this forces a new resource to be created.
	PolicyType *string `json:"policyType,omitempty" tf:"policy_type,omitempty"`
}

type PolicyDefinitionObservation struct {

	// The description of the policy definition.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of the policy definition.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the Policy Definition.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The id of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupID *string `json:"managementGroupId,omitempty" tf:"management_group_id,omitempty"`

	// The metadata for the policy definition. This is a JSON string representing additional metadata that should be stored with the policy definition.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The policy resource manager mode that allows you to specify which resource types will be evaluated. Possible values are All, Indexed, Microsoft.ContainerService.Data, Microsoft.CustomerLockbox.Data, Microsoft.DataCatalog.Data, Microsoft.KeyVault.Data, Microsoft.Kubernetes.Data, Microsoft.MachineLearningServices.Data, Microsoft.Network.Data and Microsoft.Synapse.Data.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Parameters for the policy definition. This field is a JSON string that allows you to parameterize your policy definition.
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The policy rule for the policy definition. This is a JSON string representing the rule that contains an if and a then block.
	PolicyRule *string `json:"policyRule,omitempty" tf:"policy_rule,omitempty"`

	// The policy type. Possible values are BuiltIn, Custom, NotSpecified and Static. Changing this forces a new resource to be created.
	PolicyType *string `json:"policyType,omitempty" tf:"policy_type,omitempty"`

	// A list of role definition id extracted from policy_rule required for remediation.
	RoleDefinitionIds []*string `json:"roleDefinitionIds,omitempty" tf:"role_definition_ids,omitempty"`
}

type PolicyDefinitionParameters struct {

	// The description of the policy definition.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of the policy definition.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The id of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ManagementGroupID *string `json:"managementGroupId,omitempty" tf:"management_group_id,omitempty"`

	// The metadata for the policy definition. This is a JSON string representing additional metadata that should be stored with the policy definition.
	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The policy resource manager mode that allows you to specify which resource types will be evaluated. Possible values are All, Indexed, Microsoft.ContainerService.Data, Microsoft.CustomerLockbox.Data, Microsoft.DataCatalog.Data, Microsoft.KeyVault.Data, Microsoft.Kubernetes.Data, Microsoft.MachineLearningServices.Data, Microsoft.Network.Data and Microsoft.Synapse.Data.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Parameters for the policy definition. This field is a JSON string that allows you to parameterize your policy definition.
	// +kubebuilder:validation:Optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The policy rule for the policy definition. This is a JSON string representing the rule that contains an if and a then block.
	// +kubebuilder:validation:Optional
	PolicyRule *string `json:"policyRule,omitempty" tf:"policy_rule,omitempty"`

	// The policy type. Possible values are BuiltIn, Custom, NotSpecified and Static. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PolicyType *string `json:"policyType,omitempty" tf:"policy_type,omitempty"`
}

// PolicyDefinitionSpec defines the desired state of PolicyDefinition
type PolicyDefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyDefinitionParameters `json:"forProvider"`
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
	InitProvider PolicyDefinitionInitParameters `json:"initProvider,omitempty"`
}

// PolicyDefinitionStatus defines the observed state of PolicyDefinition.
type PolicyDefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PolicyDefinition is the Schema for the PolicyDefinitions API. Manages a policy rule definition. Policy definitions do not take effect until they are assigned to a scope using a Policy Assignment.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PolicyDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.mode) || (has(self.initProvider) && has(self.initProvider.mode))",message="spec.forProvider.mode is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyType) || (has(self.initProvider) && has(self.initProvider.policyType))",message="spec.forProvider.policyType is a required parameter"
	Spec   PolicyDefinitionSpec   `json:"spec"`
	Status PolicyDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyDefinitionList contains a list of PolicyDefinitions
type PolicyDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyDefinition `json:"items"`
}

// Repository type metadata.
var (
	PolicyDefinition_Kind             = "PolicyDefinition"
	PolicyDefinition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyDefinition_Kind}.String()
	PolicyDefinition_KindAPIVersion   = PolicyDefinition_Kind + "." + CRDGroupVersion.String()
	PolicyDefinition_GroupVersionKind = CRDGroupVersion.WithKind(PolicyDefinition_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyDefinition{}, &PolicyDefinitionList{})
}
