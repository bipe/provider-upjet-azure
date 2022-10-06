/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NATIPConfigurationObservation struct {
}

type NATIPConfigurationParameters struct {

	// Specifies the name which should be used for the NAT IP Configuration. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Is this is the Primary IP Configuration? Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Primary *bool `json:"primary" tf:"primary,omitempty"`

	// Specifies a Private Static IP Address for this IP Configuration.
	// +kubebuilder:validation:Optional
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The version of the IP Protocol which should be used. At this time the only supported value is IPv4. Defaults to IPv4.
	// +kubebuilder:validation:Optional
	PrivateIPAddressVersion *string `json:"privateIpAddressVersion,omitempty" tf:"private_ip_address_version,omitempty"`

	// Specifies the ID of the Subnet which should be used for the Private Link Service.
	// +crossplane:generate:reference:type=Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type PrivateLinkServiceObservation struct {

	// A globally unique DNS Name for your Private Link Service. You can use this alias to request a connection to your Private Link Service.
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PrivateLinkServiceParameters struct {

	// A list of Subscription UUID/GUID's that will be automatically be able to use this Private Link Service.
	// +kubebuilder:validation:Optional
	AutoApprovalSubscriptionIds []*string `json:"autoApprovalSubscriptionIds,omitempty" tf:"auto_approval_subscription_ids,omitempty"`

	// Should the Private Link Service support the Proxy Protocol? Defaults to false.
	// +kubebuilder:validation:Optional
	EnableProxyProtocol *bool `json:"enableProxyProtocol,omitempty" tf:"enable_proxy_protocol,omitempty"`

	// A list of Frontend IP Configuration IDs from a Standard Load Balancer, where traffic from the Private Link Service should be routed. You can use Load Balancer Rules to direct this traffic to appropriate backend pools where your applications are running.
	// +kubebuilder:validation:Required
	LoadBalancerFrontendIPConfigurationIds []*string `json:"loadBalancerFrontendIpConfigurationIds" tf:"load_balancer_frontend_ip_configuration_ids,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// One or more (up to 8) nat_ip_configuration block as defined below.
	// +kubebuilder:validation:Required
	NATIPConfiguration []NATIPConfigurationParameters `json:"natIpConfiguration" tf:"nat_ip_configuration,omitempty"`

	// The name of the Resource Group where the Private Link Service should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A list of Subscription UUID/GUID's that will be able to see this Private Link Service.
	// +kubebuilder:validation:Optional
	VisibilitySubscriptionIds []*string `json:"visibilitySubscriptionIds,omitempty" tf:"visibility_subscription_ids,omitempty"`
}

// PrivateLinkServiceSpec defines the desired state of PrivateLinkService
type PrivateLinkServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateLinkServiceParameters `json:"forProvider"`
}

// PrivateLinkServiceStatus defines the observed state of PrivateLinkService.
type PrivateLinkServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateLinkServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateLinkService is the Schema for the PrivateLinkServices API. Manages a Private Link Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PrivateLinkService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateLinkServiceSpec   `json:"spec"`
	Status            PrivateLinkServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateLinkServiceList contains a list of PrivateLinkServices
type PrivateLinkServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateLinkService `json:"items"`
}

// Repository type metadata.
var (
	PrivateLinkService_Kind             = "PrivateLinkService"
	PrivateLinkService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateLinkService_Kind}.String()
	PrivateLinkService_KindAPIVersion   = PrivateLinkService_Kind + "." + CRDGroupVersion.String()
	PrivateLinkService_GroupVersionKind = CRDGroupVersion.WithKind(PrivateLinkService_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateLinkService{}, &PrivateLinkServiceList{})
}
