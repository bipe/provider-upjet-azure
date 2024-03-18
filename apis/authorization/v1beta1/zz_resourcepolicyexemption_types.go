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

type ResourcePolicyExemptionInitParameters struct {

	// A description to use for this Policy Exemption.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A friendly display name to use for this Policy Exemption.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The category of this policy exemption. Possible values are Waiver and Mitigated.
	ExemptionCategory *string `json:"exemptionCategory,omitempty" tf:"exemption_category,omitempty"`

	// The expiration date and time in UTC ISO 8601 format of this policy exemption.
	ExpiresOn *string `json:"expiresOn,omitempty" tf:"expires_on,omitempty"`

	// The metadata for this policy exemption. This is a JSON string representing additional metadata that should be stored with the policy exemption.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the Policy Exemption. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Policy Assignment to be exempted at the specified Scope. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/authorization/v1beta1.ResourcePolicyAssignment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	PolicyAssignmentID *string `json:"policyAssignmentId,omitempty" tf:"policy_assignment_id,omitempty"`

	// Reference to a ResourcePolicyAssignment in authorization to populate policyAssignmentId.
	// +kubebuilder:validation:Optional
	PolicyAssignmentIDRef *v1.Reference `json:"policyAssignmentIdRef,omitempty" tf:"-"`

	// Selector for a ResourcePolicyAssignment in authorization to populate policyAssignmentId.
	// +kubebuilder:validation:Optional
	PolicyAssignmentIDSelector *v1.Selector `json:"policyAssignmentIdSelector,omitempty" tf:"-"`

	// The policy definition reference ID list when the associated policy assignment is an assignment of a policy set definition.
	PolicyDefinitionReferenceIds []*string `json:"policyDefinitionReferenceIds,omitempty" tf:"policy_definition_reference_ids,omitempty"`

	// The Resource ID where the Policy Exemption should be applied. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/authorization/v1beta1.ResourcePolicyAssignment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("resource_id",false)
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a ResourcePolicyAssignment in authorization to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a ResourcePolicyAssignment in authorization to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`
}

type ResourcePolicyExemptionObservation struct {

	// A description to use for this Policy Exemption.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A friendly display name to use for this Policy Exemption.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The category of this policy exemption. Possible values are Waiver and Mitigated.
	ExemptionCategory *string `json:"exemptionCategory,omitempty" tf:"exemption_category,omitempty"`

	// The expiration date and time in UTC ISO 8601 format of this policy exemption.
	ExpiresOn *string `json:"expiresOn,omitempty" tf:"expires_on,omitempty"`

	// The Policy Exemption id.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The metadata for this policy exemption. This is a JSON string representing additional metadata that should be stored with the policy exemption.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the Policy Exemption. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Policy Assignment to be exempted at the specified Scope. Changing this forces a new resource to be created.
	PolicyAssignmentID *string `json:"policyAssignmentId,omitempty" tf:"policy_assignment_id,omitempty"`

	// The policy definition reference ID list when the associated policy assignment is an assignment of a policy set definition.
	PolicyDefinitionReferenceIds []*string `json:"policyDefinitionReferenceIds,omitempty" tf:"policy_definition_reference_ids,omitempty"`

	// The Resource ID where the Policy Exemption should be applied. Changing this forces a new resource to be created.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`
}

type ResourcePolicyExemptionParameters struct {

	// A description to use for this Policy Exemption.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A friendly display name to use for this Policy Exemption.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The category of this policy exemption. Possible values are Waiver and Mitigated.
	// +kubebuilder:validation:Optional
	ExemptionCategory *string `json:"exemptionCategory,omitempty" tf:"exemption_category,omitempty"`

	// The expiration date and time in UTC ISO 8601 format of this policy exemption.
	// +kubebuilder:validation:Optional
	ExpiresOn *string `json:"expiresOn,omitempty" tf:"expires_on,omitempty"`

	// The metadata for this policy exemption. This is a JSON string representing additional metadata that should be stored with the policy exemption.
	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the Policy Exemption. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Policy Assignment to be exempted at the specified Scope. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/authorization/v1beta1.ResourcePolicyAssignment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PolicyAssignmentID *string `json:"policyAssignmentId,omitempty" tf:"policy_assignment_id,omitempty"`

	// Reference to a ResourcePolicyAssignment in authorization to populate policyAssignmentId.
	// +kubebuilder:validation:Optional
	PolicyAssignmentIDRef *v1.Reference `json:"policyAssignmentIdRef,omitempty" tf:"-"`

	// Selector for a ResourcePolicyAssignment in authorization to populate policyAssignmentId.
	// +kubebuilder:validation:Optional
	PolicyAssignmentIDSelector *v1.Selector `json:"policyAssignmentIdSelector,omitempty" tf:"-"`

	// The policy definition reference ID list when the associated policy assignment is an assignment of a policy set definition.
	// +kubebuilder:validation:Optional
	PolicyDefinitionReferenceIds []*string `json:"policyDefinitionReferenceIds,omitempty" tf:"policy_definition_reference_ids,omitempty"`

	// The Resource ID where the Policy Exemption should be applied. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/authorization/v1beta1.ResourcePolicyAssignment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("resource_id",false)
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a ResourcePolicyAssignment in authorization to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a ResourcePolicyAssignment in authorization to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`
}

// ResourcePolicyExemptionSpec defines the desired state of ResourcePolicyExemption
type ResourcePolicyExemptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourcePolicyExemptionParameters `json:"forProvider"`
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
	InitProvider ResourcePolicyExemptionInitParameters `json:"initProvider,omitempty"`
}

// ResourcePolicyExemptionStatus defines the observed state of ResourcePolicyExemption.
type ResourcePolicyExemptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourcePolicyExemptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ResourcePolicyExemption is the Schema for the ResourcePolicyExemptions API. Manages a Resource Policy Exemption.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ResourcePolicyExemption struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.exemptionCategory) || (has(self.initProvider) && has(self.initProvider.exemptionCategory))",message="spec.forProvider.exemptionCategory is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ResourcePolicyExemptionSpec   `json:"spec"`
	Status ResourcePolicyExemptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourcePolicyExemptionList contains a list of ResourcePolicyExemptions
type ResourcePolicyExemptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourcePolicyExemption `json:"items"`
}

// Repository type metadata.
var (
	ResourcePolicyExemption_Kind             = "ResourcePolicyExemption"
	ResourcePolicyExemption_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourcePolicyExemption_Kind}.String()
	ResourcePolicyExemption_KindAPIVersion   = ResourcePolicyExemption_Kind + "." + CRDGroupVersion.String()
	ResourcePolicyExemption_GroupVersionKind = CRDGroupVersion.WithKind(ResourcePolicyExemption_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourcePolicyExemption{}, &ResourcePolicyExemptionList{})
}
