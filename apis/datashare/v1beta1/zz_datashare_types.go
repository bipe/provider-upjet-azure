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

type DataShareObservation struct {

	// The ID of the Data Share.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DataShareParameters struct {

	// The ID of the Data Share account in which the Data Share is created. Changing this forces a new Data Share to be created.
	// +crossplane:generate:reference:type=Account
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// The Data Share's description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The kind of the Data Share. Possible values are CopyBased and InPlace. Changing this forces a new Data Share to be created.
	// +kubebuilder:validation:Required
	Kind *string `json:"kind" tf:"kind,omitempty"`

	// A snapshot_schedule block as defined below.
	// +kubebuilder:validation:Optional
	SnapshotSchedule []SnapshotScheduleParameters `json:"snapshotSchedule,omitempty" tf:"snapshot_schedule,omitempty"`

	// The terms of the Data Share.
	// +kubebuilder:validation:Optional
	Terms *string `json:"terms,omitempty" tf:"terms,omitempty"`
}

type SnapshotScheduleObservation struct {
}

type SnapshotScheduleParameters struct {

	// The name of the snapshot schedule.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The interval of the synchronization with the source data. Possible values are Hour and Day.
	// +kubebuilder:validation:Required
	Recurrence *string `json:"recurrence" tf:"recurrence,omitempty"`

	// The synchronization with the source data's start time.
	// +kubebuilder:validation:Required
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

// DataShareSpec defines the desired state of DataShare
type DataShareSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataShareParameters `json:"forProvider"`
}

// DataShareStatus defines the observed state of DataShare.
type DataShareStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataShareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataShare is the Schema for the DataShares API. Manages a Data Share.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataShareSpec   `json:"spec"`
	Status            DataShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataShareList contains a list of DataShares
type DataShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataShare `json:"items"`
}

// Repository type metadata.
var (
	DataShare_Kind             = "DataShare"
	DataShare_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataShare_Kind}.String()
	DataShare_KindAPIVersion   = DataShare_Kind + "." + CRDGroupVersion.String()
	DataShare_GroupVersionKind = CRDGroupVersion.WithKind(DataShare_Kind)
)

func init() {
	SchemeBuilder.Register(&DataShare{}, &DataShareList{})
}
