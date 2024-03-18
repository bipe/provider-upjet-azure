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

type SpringCloudAppRedisAssociationInitParameters struct {

	// Specifies the Redis Cache access key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cache/v1beta1.RedisCache
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("primary_access_key",true)
	RedisAccessKey *string `json:"redisAccessKey,omitempty" tf:"redis_access_key,omitempty"`

	// Reference to a RedisCache in cache to populate redisAccessKey.
	// +kubebuilder:validation:Optional
	RedisAccessKeyRef *v1.Reference `json:"redisAccessKeyRef,omitempty" tf:"-"`

	// Selector for a RedisCache in cache to populate redisAccessKey.
	// +kubebuilder:validation:Optional
	RedisAccessKeySelector *v1.Selector `json:"redisAccessKeySelector,omitempty" tf:"-"`

	// Specifies the Redis Cache resource ID. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cache/v1beta1.RedisCache
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	RedisCacheID *string `json:"redisCacheId,omitempty" tf:"redis_cache_id,omitempty"`

	// Reference to a RedisCache in cache to populate redisCacheId.
	// +kubebuilder:validation:Optional
	RedisCacheIDRef *v1.Reference `json:"redisCacheIdRef,omitempty" tf:"-"`

	// Selector for a RedisCache in cache to populate redisCacheId.
	// +kubebuilder:validation:Optional
	RedisCacheIDSelector *v1.Selector `json:"redisCacheIdSelector,omitempty" tf:"-"`

	// Should SSL be used when connecting to Redis? Defaults to true.
	SSLEnabled *bool `json:"sslEnabled,omitempty" tf:"ssl_enabled,omitempty"`
}

type SpringCloudAppRedisAssociationObservation struct {

	// The ID of the Spring Cloud Application Redis Association.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the Redis Cache access key.
	RedisAccessKey *string `json:"redisAccessKey,omitempty" tf:"redis_access_key,omitempty"`

	// Specifies the Redis Cache resource ID. Changing this forces a new resource to be created.
	RedisCacheID *string `json:"redisCacheId,omitempty" tf:"redis_cache_id,omitempty"`

	// Should SSL be used when connecting to Redis? Defaults to true.
	SSLEnabled *bool `json:"sslEnabled,omitempty" tf:"ssl_enabled,omitempty"`

	// Specifies the Spring Cloud Application resource ID in which the Association is created. Changing this forces a new resource to be created.
	SpringCloudAppID *string `json:"springCloudAppId,omitempty" tf:"spring_cloud_app_id,omitempty"`
}

type SpringCloudAppRedisAssociationParameters struct {

	// Specifies the Redis Cache access key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cache/v1beta1.RedisCache
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("primary_access_key",true)
	// +kubebuilder:validation:Optional
	RedisAccessKey *string `json:"redisAccessKey,omitempty" tf:"redis_access_key,omitempty"`

	// Reference to a RedisCache in cache to populate redisAccessKey.
	// +kubebuilder:validation:Optional
	RedisAccessKeyRef *v1.Reference `json:"redisAccessKeyRef,omitempty" tf:"-"`

	// Selector for a RedisCache in cache to populate redisAccessKey.
	// +kubebuilder:validation:Optional
	RedisAccessKeySelector *v1.Selector `json:"redisAccessKeySelector,omitempty" tf:"-"`

	// Specifies the Redis Cache resource ID. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cache/v1beta1.RedisCache
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RedisCacheID *string `json:"redisCacheId,omitempty" tf:"redis_cache_id,omitempty"`

	// Reference to a RedisCache in cache to populate redisCacheId.
	// +kubebuilder:validation:Optional
	RedisCacheIDRef *v1.Reference `json:"redisCacheIdRef,omitempty" tf:"-"`

	// Selector for a RedisCache in cache to populate redisCacheId.
	// +kubebuilder:validation:Optional
	RedisCacheIDSelector *v1.Selector `json:"redisCacheIdSelector,omitempty" tf:"-"`

	// Should SSL be used when connecting to Redis? Defaults to true.
	// +kubebuilder:validation:Optional
	SSLEnabled *bool `json:"sslEnabled,omitempty" tf:"ssl_enabled,omitempty"`

	// Specifies the Spring Cloud Application resource ID in which the Association is created. Changing this forces a new resource to be created.
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

// SpringCloudAppRedisAssociationSpec defines the desired state of SpringCloudAppRedisAssociation
type SpringCloudAppRedisAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudAppRedisAssociationParameters `json:"forProvider"`
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
	InitProvider SpringCloudAppRedisAssociationInitParameters `json:"initProvider,omitempty"`
}

// SpringCloudAppRedisAssociationStatus defines the observed state of SpringCloudAppRedisAssociation.
type SpringCloudAppRedisAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudAppRedisAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SpringCloudAppRedisAssociation is the Schema for the SpringCloudAppRedisAssociations API. Associates a
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudAppRedisAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpringCloudAppRedisAssociationSpec   `json:"spec"`
	Status            SpringCloudAppRedisAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudAppRedisAssociationList contains a list of SpringCloudAppRedisAssociations
type SpringCloudAppRedisAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudAppRedisAssociation `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudAppRedisAssociation_Kind             = "SpringCloudAppRedisAssociation"
	SpringCloudAppRedisAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudAppRedisAssociation_Kind}.String()
	SpringCloudAppRedisAssociation_KindAPIVersion   = SpringCloudAppRedisAssociation_Kind + "." + CRDGroupVersion.String()
	SpringCloudAppRedisAssociation_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudAppRedisAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudAppRedisAssociation{}, &SpringCloudAppRedisAssociationList{})
}
