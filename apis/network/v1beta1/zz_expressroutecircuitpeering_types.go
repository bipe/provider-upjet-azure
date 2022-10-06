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

type ExpressRouteCircuitPeeringObservation struct {

	// The ASN used by Azure.
	AzureAsn *float64 `json:"azureAsn,omitempty" tf:"azure_asn,omitempty"`

	// The ID of the ExpressRoute Circuit Peering.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Primary Port used by Azure for this Peering.
	PrimaryAzurePort *string `json:"primaryAzurePort,omitempty" tf:"primary_azure_port,omitempty"`

	// The Secondary Port used by Azure for this Peering.
	SecondaryAzurePort *string `json:"secondaryAzurePort,omitempty" tf:"secondary_azure_port,omitempty"`
}

type ExpressRouteCircuitPeeringParameters struct {

	// The name of the ExpressRoute Circuit in which to create the Peering.
	// +crossplane:generate:reference:type=ExpressRouteCircuit
	// +kubebuilder:validation:Optional
	ExpressRouteCircuitName *string `json:"expressRouteCircuitName,omitempty" tf:"express_route_circuit_name,omitempty"`

	// Reference to a ExpressRouteCircuit to populate expressRouteCircuitName.
	// +kubebuilder:validation:Optional
	ExpressRouteCircuitNameRef *v1.Reference `json:"expressRouteCircuitNameRef,omitempty" tf:"-"`

	// Selector for a ExpressRouteCircuit to populate expressRouteCircuitName.
	// +kubebuilder:validation:Optional
	ExpressRouteCircuitNameSelector *v1.Selector `json:"expressRouteCircuitNameSelector,omitempty" tf:"-"`

	// A ipv6 block as defined below.
	// +kubebuilder:validation:Optional
	IPv6 []IPv6Parameters `json:"ipv6,omitempty" tf:"ipv6,omitempty"`

	// A microsoft_peering_config block as defined below. Required when peering_type is set to MicrosoftPeering.
	// +kubebuilder:validation:Optional
	MicrosoftPeeringConfig []MicrosoftPeeringConfigParameters `json:"microsoftPeeringConfig,omitempty" tf:"microsoft_peering_config,omitempty"`

	// The Either a 16-bit or a 32-bit ASN. Can either be public or private.
	// +kubebuilder:validation:Optional
	PeerAsn *float64 `json:"peerAsn,omitempty" tf:"peer_asn,omitempty"`

	// A /30 subnet for the primary link.
	// +kubebuilder:validation:Required
	PrimaryPeerAddressPrefix *string `json:"primaryPeerAddressPrefix" tf:"primary_peer_address_prefix,omitempty"`

	// The name of the resource group in which to create the Express Route Circuit Peering. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The ID of the Route Filter. Only available when peering_type is set to MicrosoftPeering.
	// +kubebuilder:validation:Optional
	RouteFilterID *string `json:"routeFilterId,omitempty" tf:"route_filter_id,omitempty"`

	// A /30 subnet for the secondary link.
	// +kubebuilder:validation:Required
	SecondaryPeerAddressPrefix *string `json:"secondaryPeerAddressPrefix" tf:"secondary_peer_address_prefix,omitempty"`

	// The shared key. Can be a maximum of 25 characters.
	// +kubebuilder:validation:Optional
	SharedKeySecretRef *v1.SecretKeySelector `json:"sharedKeySecretRef,omitempty" tf:"-"`

	// A valid VLAN ID to establish this peering on.
	// +kubebuilder:validation:Required
	VlanID *float64 `json:"vlanId" tf:"vlan_id,omitempty"`
}

type IPv6Observation struct {
}

type IPv6Parameters struct {

	// A microsoft_peering block as defined below.
	// +kubebuilder:validation:Required
	MicrosoftPeering []MicrosoftPeeringParameters `json:"microsoftPeering" tf:"microsoft_peering,omitempty"`

	// A subnet for the primary link.
	// +kubebuilder:validation:Required
	PrimaryPeerAddressPrefix *string `json:"primaryPeerAddressPrefix" tf:"primary_peer_address_prefix,omitempty"`

	// The ID of the Route Filter. Only available when peering_type is set to MicrosoftPeering.
	// +kubebuilder:validation:Optional
	RouteFilterID *string `json:"routeFilterId,omitempty" tf:"route_filter_id,omitempty"`

	// A subnet for the secondary link.
	// +kubebuilder:validation:Required
	SecondaryPeerAddressPrefix *string `json:"secondaryPeerAddressPrefix" tf:"secondary_peer_address_prefix,omitempty"`
}

type MicrosoftPeeringConfigObservation struct {
}

type MicrosoftPeeringConfigParameters struct {

	// A list of Advertised Public Prefixes.
	// +kubebuilder:validation:Required
	AdvertisedPublicPrefixes []*string `json:"advertisedPublicPrefixes" tf:"advertised_public_prefixes,omitempty"`

	// The CustomerASN of the peering.
	// +kubebuilder:validation:Optional
	CustomerAsn *float64 `json:"customerAsn,omitempty" tf:"customer_asn,omitempty"`

	// The Routing Registry against which the AS number and prefixes are registered.  For example:  ARIN, RIPE, AFRINIC etc.
	// +kubebuilder:validation:Optional
	RoutingRegistryName *string `json:"routingRegistryName,omitempty" tf:"routing_registry_name,omitempty"`
}

type MicrosoftPeeringObservation struct {
}

type MicrosoftPeeringParameters struct {

	// A list of Advertised Public Prefixes.
	// +kubebuilder:validation:Optional
	AdvertisedPublicPrefixes []*string `json:"advertisedPublicPrefixes,omitempty" tf:"advertised_public_prefixes,omitempty"`

	// The CustomerASN of the peering.
	// +kubebuilder:validation:Optional
	CustomerAsn *float64 `json:"customerAsn,omitempty" tf:"customer_asn,omitempty"`

	// The Routing Registry against which the AS number and prefixes are registered.  For example:  ARIN, RIPE, AFRINIC etc.
	// +kubebuilder:validation:Optional
	RoutingRegistryName *string `json:"routingRegistryName,omitempty" tf:"routing_registry_name,omitempty"`
}

// ExpressRouteCircuitPeeringSpec defines the desired state of ExpressRouteCircuitPeering
type ExpressRouteCircuitPeeringSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExpressRouteCircuitPeeringParameters `json:"forProvider"`
}

// ExpressRouteCircuitPeeringStatus defines the observed state of ExpressRouteCircuitPeering.
type ExpressRouteCircuitPeeringStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExpressRouteCircuitPeeringObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteCircuitPeering is the Schema for the ExpressRouteCircuitPeerings API. Manages an ExpressRoute Circuit Peering.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ExpressRouteCircuitPeering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRouteCircuitPeeringSpec   `json:"spec"`
	Status            ExpressRouteCircuitPeeringStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteCircuitPeeringList contains a list of ExpressRouteCircuitPeerings
type ExpressRouteCircuitPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExpressRouteCircuitPeering `json:"items"`
}

// Repository type metadata.
var (
	ExpressRouteCircuitPeering_Kind             = "ExpressRouteCircuitPeering"
	ExpressRouteCircuitPeering_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExpressRouteCircuitPeering_Kind}.String()
	ExpressRouteCircuitPeering_KindAPIVersion   = ExpressRouteCircuitPeering_Kind + "." + CRDGroupVersion.String()
	ExpressRouteCircuitPeering_GroupVersionKind = CRDGroupVersion.WithKind(ExpressRouteCircuitPeering_Kind)
)

func init() {
	SchemeBuilder.Register(&ExpressRouteCircuitPeering{}, &ExpressRouteCircuitPeeringList{})
}
