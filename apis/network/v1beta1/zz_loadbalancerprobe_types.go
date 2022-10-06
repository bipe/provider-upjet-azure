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

type LoadBalancerProbeObservation struct {

	// The ID of the Load Balancer Probe.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LoadBalancerRules []*string `json:"loadBalancerRules,omitempty" tf:"load_balancer_rules,omitempty"`
}

type LoadBalancerProbeParameters struct {

	// The interval, in seconds between probes to the backend endpoint for health status. The default value is 15, the minimum value is 5.
	// +kubebuilder:validation:Optional
	IntervalInSeconds *float64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`

	// The ID of the LoadBalancer in which to create the NAT Rule.
	// +crossplane:generate:reference:type=LoadBalancer
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LoadbalancerID *string `json:"loadbalancerId,omitempty" tf:"loadbalancer_id,omitempty"`

	// Reference to a LoadBalancer to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDRef *v1.Reference `json:"loadbalancerIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancer to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDSelector *v1.Selector `json:"loadbalancerIdSelector,omitempty" tf:"-"`

	// The number of failed probe attempts after which the backend endpoint is removed from rotation. The default value is 2. NumberOfProbes multiplied by intervalInSeconds value must be greater or equal to 10.Endpoints are returned to rotation when at least one probe is successful.
	// +kubebuilder:validation:Optional
	NumberOfProbes *float64 `json:"numberOfProbes,omitempty" tf:"number_of_probes,omitempty"`

	// Port on which the Probe queries the backend endpoint. Possible values range from 1 to 65535, inclusive.
	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// Specifies the protocol of the end point. Possible values are Http, Https or Tcp. If TCP is specified, a received ACK is required for the probe to be successful. If HTTP is specified, a 200 OK response from the specified URI is required for the probe to be successful.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The URI used for requesting health status from the backend endpoint. Required if protocol is set to Http or Https. Otherwise, it is not allowed.
	// +kubebuilder:validation:Optional
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`
}

// LoadBalancerProbeSpec defines the desired state of LoadBalancerProbe
type LoadBalancerProbeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerProbeParameters `json:"forProvider"`
}

// LoadBalancerProbeStatus defines the observed state of LoadBalancerProbe.
type LoadBalancerProbeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerProbeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerProbe is the Schema for the LoadBalancerProbes API. Manages a Load Balancer Probe Resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LoadBalancerProbe struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerProbeSpec   `json:"spec"`
	Status            LoadBalancerProbeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerProbeList contains a list of LoadBalancerProbes
type LoadBalancerProbeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancerProbe `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancerProbe_Kind             = "LoadBalancerProbe"
	LoadBalancerProbe_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancerProbe_Kind}.String()
	LoadBalancerProbe_KindAPIVersion   = LoadBalancerProbe_Kind + "." + CRDGroupVersion.String()
	LoadBalancerProbe_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancerProbe_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancerProbe{}, &LoadBalancerProbeList{})
}
