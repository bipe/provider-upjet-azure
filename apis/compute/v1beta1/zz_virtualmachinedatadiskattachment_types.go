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

type VirtualMachineDataDiskAttachmentInitParameters struct {

	// Specifies the caching requirements for this Data Disk. Possible values include None, ReadOnly and ReadWrite.
	Caching *string `json:"caching,omitempty" tf:"caching,omitempty"`

	// The Create Option of the Data Disk, such as Empty or Attach. Defaults to Attach. Changing this forces a new resource to be created.
	CreateOption *string `json:"createOption,omitempty" tf:"create_option,omitempty"`

	// The Logical Unit Number of the Data Disk, which needs to be unique within the Virtual Machine. Changing this forces a new resource to be created.
	Lun *float64 `json:"lun,omitempty" tf:"lun,omitempty"`

	// The ID of an existing Managed Disk which should be attached. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.ManagedDisk
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ManagedDiskID *string `json:"managedDiskId,omitempty" tf:"managed_disk_id,omitempty"`

	// Reference to a ManagedDisk in compute to populate managedDiskId.
	// +kubebuilder:validation:Optional
	ManagedDiskIDRef *v1.Reference `json:"managedDiskIdRef,omitempty" tf:"-"`

	// Selector for a ManagedDisk in compute to populate managedDiskId.
	// +kubebuilder:validation:Optional
	ManagedDiskIDSelector *v1.Selector `json:"managedDiskIdSelector,omitempty" tf:"-"`

	// The ID of the Virtual Machine to which the Data Disk should be attached. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.LinuxVirtualMachine
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`

	// Reference to a LinuxVirtualMachine in compute to populate virtualMachineId.
	// +kubebuilder:validation:Optional
	VirtualMachineIDRef *v1.Reference `json:"virtualMachineIdRef,omitempty" tf:"-"`

	// Selector for a LinuxVirtualMachine in compute to populate virtualMachineId.
	// +kubebuilder:validation:Optional
	VirtualMachineIDSelector *v1.Selector `json:"virtualMachineIdSelector,omitempty" tf:"-"`

	// Specifies if Write Accelerator is enabled on the disk. This can only be enabled on Premium_LRS managed disks with no caching and M-Series VMs. Defaults to false.
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled,omitempty"`
}

type VirtualMachineDataDiskAttachmentObservation struct {

	// Specifies the caching requirements for this Data Disk. Possible values include None, ReadOnly and ReadWrite.
	Caching *string `json:"caching,omitempty" tf:"caching,omitempty"`

	// The Create Option of the Data Disk, such as Empty or Attach. Defaults to Attach. Changing this forces a new resource to be created.
	CreateOption *string `json:"createOption,omitempty" tf:"create_option,omitempty"`

	// The ID of the Virtual Machine Data Disk attachment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Logical Unit Number of the Data Disk, which needs to be unique within the Virtual Machine. Changing this forces a new resource to be created.
	Lun *float64 `json:"lun,omitempty" tf:"lun,omitempty"`

	// The ID of an existing Managed Disk which should be attached. Changing this forces a new resource to be created.
	ManagedDiskID *string `json:"managedDiskId,omitempty" tf:"managed_disk_id,omitempty"`

	// The ID of the Virtual Machine to which the Data Disk should be attached. Changing this forces a new resource to be created.
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`

	// Specifies if Write Accelerator is enabled on the disk. This can only be enabled on Premium_LRS managed disks with no caching and M-Series VMs. Defaults to false.
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled,omitempty"`
}

type VirtualMachineDataDiskAttachmentParameters struct {

	// Specifies the caching requirements for this Data Disk. Possible values include None, ReadOnly and ReadWrite.
	// +kubebuilder:validation:Optional
	Caching *string `json:"caching,omitempty" tf:"caching,omitempty"`

	// The Create Option of the Data Disk, such as Empty or Attach. Defaults to Attach. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	CreateOption *string `json:"createOption,omitempty" tf:"create_option,omitempty"`

	// The Logical Unit Number of the Data Disk, which needs to be unique within the Virtual Machine. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Lun *float64 `json:"lun,omitempty" tf:"lun,omitempty"`

	// The ID of an existing Managed Disk which should be attached. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.ManagedDisk
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ManagedDiskID *string `json:"managedDiskId,omitempty" tf:"managed_disk_id,omitempty"`

	// Reference to a ManagedDisk in compute to populate managedDiskId.
	// +kubebuilder:validation:Optional
	ManagedDiskIDRef *v1.Reference `json:"managedDiskIdRef,omitempty" tf:"-"`

	// Selector for a ManagedDisk in compute to populate managedDiskId.
	// +kubebuilder:validation:Optional
	ManagedDiskIDSelector *v1.Selector `json:"managedDiskIdSelector,omitempty" tf:"-"`

	// The ID of the Virtual Machine to which the Data Disk should be attached. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.LinuxVirtualMachine
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`

	// Reference to a LinuxVirtualMachine in compute to populate virtualMachineId.
	// +kubebuilder:validation:Optional
	VirtualMachineIDRef *v1.Reference `json:"virtualMachineIdRef,omitempty" tf:"-"`

	// Selector for a LinuxVirtualMachine in compute to populate virtualMachineId.
	// +kubebuilder:validation:Optional
	VirtualMachineIDSelector *v1.Selector `json:"virtualMachineIdSelector,omitempty" tf:"-"`

	// Specifies if Write Accelerator is enabled on the disk. This can only be enabled on Premium_LRS managed disks with no caching and M-Series VMs. Defaults to false.
	// +kubebuilder:validation:Optional
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled,omitempty"`
}

// VirtualMachineDataDiskAttachmentSpec defines the desired state of VirtualMachineDataDiskAttachment
type VirtualMachineDataDiskAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualMachineDataDiskAttachmentParameters `json:"forProvider"`
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
	InitProvider VirtualMachineDataDiskAttachmentInitParameters `json:"initProvider,omitempty"`
}

// VirtualMachineDataDiskAttachmentStatus defines the observed state of VirtualMachineDataDiskAttachment.
type VirtualMachineDataDiskAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualMachineDataDiskAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VirtualMachineDataDiskAttachment is the Schema for the VirtualMachineDataDiskAttachments API. Manages attaching a Disk to a Virtual Machine.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualMachineDataDiskAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.caching) || (has(self.initProvider) && has(self.initProvider.caching))",message="spec.forProvider.caching is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.lun) || (has(self.initProvider) && has(self.initProvider.lun))",message="spec.forProvider.lun is a required parameter"
	Spec   VirtualMachineDataDiskAttachmentSpec   `json:"spec"`
	Status VirtualMachineDataDiskAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineDataDiskAttachmentList contains a list of VirtualMachineDataDiskAttachments
type VirtualMachineDataDiskAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineDataDiskAttachment `json:"items"`
}

// Repository type metadata.
var (
	VirtualMachineDataDiskAttachment_Kind             = "VirtualMachineDataDiskAttachment"
	VirtualMachineDataDiskAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualMachineDataDiskAttachment_Kind}.String()
	VirtualMachineDataDiskAttachment_KindAPIVersion   = VirtualMachineDataDiskAttachment_Kind + "." + CRDGroupVersion.String()
	VirtualMachineDataDiskAttachment_GroupVersionKind = CRDGroupVersion.WithKind(VirtualMachineDataDiskAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualMachineDataDiskAttachment{}, &VirtualMachineDataDiskAttachmentList{})
}
