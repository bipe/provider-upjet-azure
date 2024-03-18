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

type SpringCloudContainerDeploymentInitParameters struct {

	// A JSON object that contains the addon configurations of the Spring Cloud Container Deployment.
	AddonJSON *string `json:"addonJson,omitempty" tf:"addon_json,omitempty"`

	// Specifies a list of Spring Cloud Application Performance Monitoring IDs.
	ApplicationPerformanceMonitoringIds []*string `json:"applicationPerformanceMonitoringIds,omitempty" tf:"application_performance_monitoring_ids,omitempty"`

	// Specifies the arguments to the entrypoint. The docker image's CMD is used if not specified.
	Arguments []*string `json:"arguments,omitempty" tf:"arguments,omitempty"`

	// Specifies the entrypoint array. It will not be executed within a shell. The docker image's ENTRYPOINT is used if not specified.
	Commands []*string `json:"commands,omitempty" tf:"commands,omitempty"`

	// Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
	// +mapType=granular
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// Container image of the custom container. This should be in the form of <repository>:<tag> without the server name of the registry.
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between 1 and 500. Defaults to 1 if not specified.
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// Specifies the language framework of the container image. The only possible value is springboot.
	LanguageFramework *string `json:"languageFramework,omitempty" tf:"language_framework,omitempty"`

	// A quota block as defined below.
	Quota []SpringCloudContainerDeploymentQuotaInitParameters `json:"quota,omitempty" tf:"quota,omitempty"`

	// The name of the registry that contains the container image.
	Server *string `json:"server,omitempty" tf:"server,omitempty"`
}

type SpringCloudContainerDeploymentObservation struct {

	// A JSON object that contains the addon configurations of the Spring Cloud Container Deployment.
	AddonJSON *string `json:"addonJson,omitempty" tf:"addon_json,omitempty"`

	// Specifies a list of Spring Cloud Application Performance Monitoring IDs.
	ApplicationPerformanceMonitoringIds []*string `json:"applicationPerformanceMonitoringIds,omitempty" tf:"application_performance_monitoring_ids,omitempty"`

	// Specifies the arguments to the entrypoint. The docker image's CMD is used if not specified.
	Arguments []*string `json:"arguments,omitempty" tf:"arguments,omitempty"`

	// Specifies the entrypoint array. It will not be executed within a shell. The docker image's ENTRYPOINT is used if not specified.
	Commands []*string `json:"commands,omitempty" tf:"commands,omitempty"`

	// Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
	// +mapType=granular
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// The ID of the Spring Cloud Container Deployment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Container image of the custom container. This should be in the form of <repository>:<tag> without the server name of the registry.
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between 1 and 500. Defaults to 1 if not specified.
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// Specifies the language framework of the container image. The only possible value is springboot.
	LanguageFramework *string `json:"languageFramework,omitempty" tf:"language_framework,omitempty"`

	// A quota block as defined below.
	Quota []SpringCloudContainerDeploymentQuotaObservation `json:"quota,omitempty" tf:"quota,omitempty"`

	// The name of the registry that contains the container image.
	Server *string `json:"server,omitempty" tf:"server,omitempty"`

	// The ID of the Spring Cloud Service. Changing this forces a new Spring Cloud Container Deployment to be created.
	SpringCloudAppID *string `json:"springCloudAppId,omitempty" tf:"spring_cloud_app_id,omitempty"`
}

type SpringCloudContainerDeploymentParameters struct {

	// A JSON object that contains the addon configurations of the Spring Cloud Container Deployment.
	// +kubebuilder:validation:Optional
	AddonJSON *string `json:"addonJson,omitempty" tf:"addon_json,omitempty"`

	// Specifies a list of Spring Cloud Application Performance Monitoring IDs.
	// +kubebuilder:validation:Optional
	ApplicationPerformanceMonitoringIds []*string `json:"applicationPerformanceMonitoringIds,omitempty" tf:"application_performance_monitoring_ids,omitempty"`

	// Specifies the arguments to the entrypoint. The docker image's CMD is used if not specified.
	// +kubebuilder:validation:Optional
	Arguments []*string `json:"arguments,omitempty" tf:"arguments,omitempty"`

	// Specifies the entrypoint array. It will not be executed within a shell. The docker image's ENTRYPOINT is used if not specified.
	// +kubebuilder:validation:Optional
	Commands []*string `json:"commands,omitempty" tf:"commands,omitempty"`

	// Specifies the environment variables of the Spring Cloud Deployment as a map of key-value pairs.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// Container image of the custom container. This should be in the form of <repository>:<tag> without the server name of the registry.
	// +kubebuilder:validation:Optional
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Specifies the required instance count of the Spring Cloud Deployment. Possible Values are between 1 and 500. Defaults to 1 if not specified.
	// +kubebuilder:validation:Optional
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// Specifies the language framework of the container image. The only possible value is springboot.
	// +kubebuilder:validation:Optional
	LanguageFramework *string `json:"languageFramework,omitempty" tf:"language_framework,omitempty"`

	// A quota block as defined below.
	// +kubebuilder:validation:Optional
	Quota []SpringCloudContainerDeploymentQuotaParameters `json:"quota,omitempty" tf:"quota,omitempty"`

	// The name of the registry that contains the container image.
	// +kubebuilder:validation:Optional
	Server *string `json:"server,omitempty" tf:"server,omitempty"`

	// The ID of the Spring Cloud Service. Changing this forces a new Spring Cloud Container Deployment to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudApp
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudAppID *string `json:"springCloudAppId,omitempty" tf:"spring_cloud_app_id,omitempty"`

	// Reference to a SpringCloudApp in appplatform to populate springCloudAppId.
	// +kubebuilder:validation:Optional
	SpringCloudAppIDRef *v1.Reference `json:"springCloudAppIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudApp in appplatform to populate springCloudAppId.
	// +kubebuilder:validation:Optional
	SpringCloudAppIDSelector *v1.Selector `json:"springCloudAppIdSelector,omitempty" tf:"-"`
}

type SpringCloudContainerDeploymentQuotaInitParameters struct {

	// Specifies the required cpu of the Spring Cloud Deployment. Possible Values are 500m, 1, 2, 3 and 4. Defaults to 1 if not specified.
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// Specifies the required memory size of the Spring Cloud Deployment. Possible Values are 512Mi, 1Gi, 2Gi, 3Gi, 4Gi, 5Gi, 6Gi, 7Gi, and 8Gi. Defaults to 1Gi if not specified.
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

type SpringCloudContainerDeploymentQuotaObservation struct {

	// Specifies the required cpu of the Spring Cloud Deployment. Possible Values are 500m, 1, 2, 3 and 4. Defaults to 1 if not specified.
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// Specifies the required memory size of the Spring Cloud Deployment. Possible Values are 512Mi, 1Gi, 2Gi, 3Gi, 4Gi, 5Gi, 6Gi, 7Gi, and 8Gi. Defaults to 1Gi if not specified.
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

type SpringCloudContainerDeploymentQuotaParameters struct {

	// Specifies the required cpu of the Spring Cloud Deployment. Possible Values are 500m, 1, 2, 3 and 4. Defaults to 1 if not specified.
	// +kubebuilder:validation:Optional
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// Specifies the required memory size of the Spring Cloud Deployment. Possible Values are 512Mi, 1Gi, 2Gi, 3Gi, 4Gi, 5Gi, 6Gi, 7Gi, and 8Gi. Defaults to 1Gi if not specified.
	// +kubebuilder:validation:Optional
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

// SpringCloudContainerDeploymentSpec defines the desired state of SpringCloudContainerDeployment
type SpringCloudContainerDeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudContainerDeploymentParameters `json:"forProvider"`
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
	InitProvider SpringCloudContainerDeploymentInitParameters `json:"initProvider,omitempty"`
}

// SpringCloudContainerDeploymentStatus defines the observed state of SpringCloudContainerDeployment.
type SpringCloudContainerDeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudContainerDeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SpringCloudContainerDeployment is the Schema for the SpringCloudContainerDeployments API. Manages a Spring Cloud Container Deployment.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudContainerDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.image) || (has(self.initProvider) && has(self.initProvider.image))",message="spec.forProvider.image is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.server) || (has(self.initProvider) && has(self.initProvider.server))",message="spec.forProvider.server is a required parameter"
	Spec   SpringCloudContainerDeploymentSpec   `json:"spec"`
	Status SpringCloudContainerDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudContainerDeploymentList contains a list of SpringCloudContainerDeployments
type SpringCloudContainerDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudContainerDeployment `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudContainerDeployment_Kind             = "SpringCloudContainerDeployment"
	SpringCloudContainerDeployment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudContainerDeployment_Kind}.String()
	SpringCloudContainerDeployment_KindAPIVersion   = SpringCloudContainerDeployment_Kind + "." + CRDGroupVersion.String()
	SpringCloudContainerDeployment_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudContainerDeployment_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudContainerDeployment{}, &SpringCloudContainerDeploymentList{})
}
