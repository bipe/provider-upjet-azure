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

type MonthlyInitParameters struct {

	// The occurrence of the specified day during the month. For example, a monthly property with weekday and week values of Sunday, -1 means the last Sunday of the month.
	Week *float64 `json:"week,omitempty" tf:"week,omitempty"`

	// The day of the week on which the trigger runs. For example, a monthly property with a weekday value of Sunday means every Sunday of the month.
	Weekday *string `json:"weekday,omitempty" tf:"weekday,omitempty"`
}

type MonthlyObservation struct {

	// The occurrence of the specified day during the month. For example, a monthly property with weekday and week values of Sunday, -1 means the last Sunday of the month.
	Week *float64 `json:"week,omitempty" tf:"week,omitempty"`

	// The day of the week on which the trigger runs. For example, a monthly property with a weekday value of Sunday means every Sunday of the month.
	Weekday *string `json:"weekday,omitempty" tf:"weekday,omitempty"`
}

type MonthlyParameters struct {

	// The occurrence of the specified day during the month. For example, a monthly property with weekday and week values of Sunday, -1 means the last Sunday of the month.
	// +kubebuilder:validation:Optional
	Week *float64 `json:"week,omitempty" tf:"week,omitempty"`

	// The day of the week on which the trigger runs. For example, a monthly property with a weekday value of Sunday means every Sunday of the month.
	// +kubebuilder:validation:Optional
	Weekday *string `json:"weekday" tf:"weekday,omitempty"`
}

type ScheduleInitParameters struct {

	// Day(s) of the month on which the trigger is scheduled. This value can be specified with a monthly frequency only.
	DaysOfMonth []*float64 `json:"daysOfMonth,omitempty" tf:"days_of_month,omitempty"`

	// Days of the week on which the trigger is scheduled. This value can be specified only with a weekly frequency.
	DaysOfWeek []*string `json:"daysOfWeek,omitempty" tf:"days_of_week,omitempty"`

	// Hours of the day on which the trigger is scheduled.
	Hours []*float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// Minutes of the hour on which the trigger is scheduled.
	Minutes []*float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// A monthly block as documented below, which specifies the days of the month on which the trigger is scheduled. The value can be specified only with a monthly frequency.
	Monthly []MonthlyInitParameters `json:"monthly,omitempty" tf:"monthly,omitempty"`
}

type ScheduleObservation struct {

	// Day(s) of the month on which the trigger is scheduled. This value can be specified with a monthly frequency only.
	DaysOfMonth []*float64 `json:"daysOfMonth,omitempty" tf:"days_of_month,omitempty"`

	// Days of the week on which the trigger is scheduled. This value can be specified only with a weekly frequency.
	DaysOfWeek []*string `json:"daysOfWeek,omitempty" tf:"days_of_week,omitempty"`

	// Hours of the day on which the trigger is scheduled.
	Hours []*float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// Minutes of the hour on which the trigger is scheduled.
	Minutes []*float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// A monthly block as documented below, which specifies the days of the month on which the trigger is scheduled. The value can be specified only with a monthly frequency.
	Monthly []MonthlyObservation `json:"monthly,omitempty" tf:"monthly,omitempty"`
}

type ScheduleParameters struct {

	// Day(s) of the month on which the trigger is scheduled. This value can be specified with a monthly frequency only.
	// +kubebuilder:validation:Optional
	DaysOfMonth []*float64 `json:"daysOfMonth,omitempty" tf:"days_of_month,omitempty"`

	// Days of the week on which the trigger is scheduled. This value can be specified only with a weekly frequency.
	// +kubebuilder:validation:Optional
	DaysOfWeek []*string `json:"daysOfWeek,omitempty" tf:"days_of_week,omitempty"`

	// Hours of the day on which the trigger is scheduled.
	// +kubebuilder:validation:Optional
	Hours []*float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// Minutes of the hour on which the trigger is scheduled.
	// +kubebuilder:validation:Optional
	Minutes []*float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// A monthly block as documented below, which specifies the days of the month on which the trigger is scheduled. The value can be specified only with a monthly frequency.
	// +kubebuilder:validation:Optional
	Monthly []MonthlyParameters `json:"monthly,omitempty" tf:"monthly,omitempty"`
}

type TriggerScheduleInitParameters struct {

	// Specifies if the Data Factory Schedule Trigger is activated. Defaults to true.
	Activated *bool `json:"activated,omitempty" tf:"activated,omitempty"`

	// List of tags that can be used for describing the Data Factory Schedule Trigger.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The Schedule Trigger's description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The time the Schedule Trigger should end. The time will be represented in UTC.
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The trigger frequency. Valid values include Minute, Hour, Day, Week, Month. Defaults to Minute.
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// The interval for how often the trigger occurs. This defaults to 1.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// A pipeline block as defined below.
	Pipeline []TriggerSchedulePipelineInitParameters `json:"pipeline,omitempty" tf:"pipeline,omitempty"`

	// The Data Factory Pipeline name that the trigger will act on.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Pipeline
	PipelineName *string `json:"pipelineName,omitempty" tf:"pipeline_name,omitempty"`

	// Reference to a Pipeline in datafactory to populate pipelineName.
	// +kubebuilder:validation:Optional
	PipelineNameRef *v1.Reference `json:"pipelineNameRef,omitempty" tf:"-"`

	// Selector for a Pipeline in datafactory to populate pipelineName.
	// +kubebuilder:validation:Optional
	PipelineNameSelector *v1.Selector `json:"pipelineNameSelector,omitempty" tf:"-"`

	// The pipeline parameters that the trigger will act upon.
	// +mapType=granular
	PipelineParameters map[string]*string `json:"pipelineParameters,omitempty" tf:"pipeline_parameters,omitempty"`

	// A schedule block as defined below, which further specifies the recurrence schedule for the trigger. A schedule is capable of limiting or increasing the number of trigger executions specified by the frequency and interval properties.
	Schedule []ScheduleInitParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The time the Schedule Trigger will start. This defaults to the current time. The time will be represented in UTC.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// The timezone of the start/end time.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type TriggerScheduleObservation struct {

	// Specifies if the Data Factory Schedule Trigger is activated. Defaults to true.
	Activated *bool `json:"activated,omitempty" tf:"activated,omitempty"`

	// List of tags that can be used for describing the Data Factory Schedule Trigger.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The Schedule Trigger's description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The time the Schedule Trigger should end. The time will be represented in UTC.
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The trigger frequency. Valid values include Minute, Hour, Day, Week, Month. Defaults to Minute.
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// The ID of the Data Factory Schedule Trigger.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The interval for how often the trigger occurs. This defaults to 1.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// A pipeline block as defined below.
	Pipeline []TriggerSchedulePipelineObservation `json:"pipeline,omitempty" tf:"pipeline,omitempty"`

	// The Data Factory Pipeline name that the trigger will act on.
	PipelineName *string `json:"pipelineName,omitempty" tf:"pipeline_name,omitempty"`

	// The pipeline parameters that the trigger will act upon.
	// +mapType=granular
	PipelineParameters map[string]*string `json:"pipelineParameters,omitempty" tf:"pipeline_parameters,omitempty"`

	// A schedule block as defined below, which further specifies the recurrence schedule for the trigger. A schedule is capable of limiting or increasing the number of trigger executions specified by the frequency and interval properties.
	Schedule []ScheduleObservation `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The time the Schedule Trigger will start. This defaults to the current time. The time will be represented in UTC.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// The timezone of the start/end time.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type TriggerScheduleParameters struct {

	// Specifies if the Data Factory Schedule Trigger is activated. Defaults to true.
	// +kubebuilder:validation:Optional
	Activated *bool `json:"activated,omitempty" tf:"activated,omitempty"`

	// List of tags that can be used for describing the Data Factory Schedule Trigger.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The Schedule Trigger's description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The time the Schedule Trigger should end. The time will be represented in UTC.
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The trigger frequency. Valid values include Minute, Hour, Day, Week, Month. Defaults to Minute.
	// +kubebuilder:validation:Optional
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// The interval for how often the trigger occurs. This defaults to 1.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// A pipeline block as defined below.
	// +kubebuilder:validation:Optional
	Pipeline []TriggerSchedulePipelineParameters `json:"pipeline,omitempty" tf:"pipeline,omitempty"`

	// The Data Factory Pipeline name that the trigger will act on.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Pipeline
	// +kubebuilder:validation:Optional
	PipelineName *string `json:"pipelineName,omitempty" tf:"pipeline_name,omitempty"`

	// Reference to a Pipeline in datafactory to populate pipelineName.
	// +kubebuilder:validation:Optional
	PipelineNameRef *v1.Reference `json:"pipelineNameRef,omitempty" tf:"-"`

	// Selector for a Pipeline in datafactory to populate pipelineName.
	// +kubebuilder:validation:Optional
	PipelineNameSelector *v1.Selector `json:"pipelineNameSelector,omitempty" tf:"-"`

	// The pipeline parameters that the trigger will act upon.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	PipelineParameters map[string]*string `json:"pipelineParameters,omitempty" tf:"pipeline_parameters,omitempty"`

	// A schedule block as defined below, which further specifies the recurrence schedule for the trigger. A schedule is capable of limiting or increasing the number of trigger executions specified by the frequency and interval properties.
	// +kubebuilder:validation:Optional
	Schedule []ScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The time the Schedule Trigger will start. This defaults to the current time. The time will be represented in UTC.
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// The timezone of the start/end time.
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type TriggerSchedulePipelineInitParameters struct {

	// Reference pipeline name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The pipeline parameters that the trigger will act upon.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type TriggerSchedulePipelineObservation struct {

	// Reference pipeline name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The pipeline parameters that the trigger will act upon.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type TriggerSchedulePipelineParameters struct {

	// Reference pipeline name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The pipeline parameters that the trigger will act upon.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

// TriggerScheduleSpec defines the desired state of TriggerSchedule
type TriggerScheduleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TriggerScheduleParameters `json:"forProvider"`
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
	InitProvider TriggerScheduleInitParameters `json:"initProvider,omitempty"`
}

// TriggerScheduleStatus defines the observed state of TriggerSchedule.
type TriggerScheduleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TriggerScheduleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TriggerSchedule is the Schema for the TriggerSchedules API. Manages a Trigger Schedule inside a Azure Data Factory.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type TriggerSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TriggerScheduleSpec   `json:"spec"`
	Status            TriggerScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerScheduleList contains a list of TriggerSchedules
type TriggerScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TriggerSchedule `json:"items"`
}

// Repository type metadata.
var (
	TriggerSchedule_Kind             = "TriggerSchedule"
	TriggerSchedule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TriggerSchedule_Kind}.String()
	TriggerSchedule_KindAPIVersion   = TriggerSchedule_Kind + "." + CRDGroupVersion.String()
	TriggerSchedule_GroupVersionKind = CRDGroupVersion.WithKind(TriggerSchedule_Kind)
)

func init() {
	SchemeBuilder.Register(&TriggerSchedule{}, &TriggerScheduleList{})
}
