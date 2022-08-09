/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev2
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// VmwareObjectDiskModel struct for VmwareObjectDiskModel
type VmwareObjectDiskModel struct {
	VmObject VmwareObjectModel `json:"vmObject"`
	DisksToProcess EVmwareDisksTypeToProcess `json:"disksToProcess"`
	// Array of disks.
	Disks []string `json:"disks"`
	// If *true*, the disk is removed from VM configuration.
	RemoveFromVMConfiguration *bool `json:"removeFromVMConfiguration,omitempty"`
}

// NewVmwareObjectDiskModel instantiates a new VmwareObjectDiskModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareObjectDiskModel(vmObject VmwareObjectModel, disksToProcess EVmwareDisksTypeToProcess, disks []string, ) *VmwareObjectDiskModel {
	this := VmwareObjectDiskModel{}
	this.VmObject = vmObject
	this.DisksToProcess = disksToProcess
	this.Disks = disks
	return &this
}

// NewVmwareObjectDiskModelWithDefaults instantiates a new VmwareObjectDiskModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareObjectDiskModelWithDefaults() *VmwareObjectDiskModel {
	this := VmwareObjectDiskModel{}
	return &this
}

// GetVmObject returns the VmObject field value
func (o *VmwareObjectDiskModel) GetVmObject() VmwareObjectModel {
	if o == nil  {
		var ret VmwareObjectModel
		return ret
	}

	return o.VmObject
}

// GetVmObjectOk returns a tuple with the VmObject field value
// and a boolean to check if the value has been set.
func (o *VmwareObjectDiskModel) GetVmObjectOk() (*VmwareObjectModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VmObject, true
}

// SetVmObject sets field value
func (o *VmwareObjectDiskModel) SetVmObject(v VmwareObjectModel) {
	o.VmObject = v
}

// GetDisksToProcess returns the DisksToProcess field value
func (o *VmwareObjectDiskModel) GetDisksToProcess() EVmwareDisksTypeToProcess {
	if o == nil  {
		var ret EVmwareDisksTypeToProcess
		return ret
	}

	return o.DisksToProcess
}

// GetDisksToProcessOk returns a tuple with the DisksToProcess field value
// and a boolean to check if the value has been set.
func (o *VmwareObjectDiskModel) GetDisksToProcessOk() (*EVmwareDisksTypeToProcess, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisksToProcess, true
}

// SetDisksToProcess sets field value
func (o *VmwareObjectDiskModel) SetDisksToProcess(v EVmwareDisksTypeToProcess) {
	o.DisksToProcess = v
}

// GetDisks returns the Disks field value
func (o *VmwareObjectDiskModel) GetDisks() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value
// and a boolean to check if the value has been set.
func (o *VmwareObjectDiskModel) GetDisksOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Disks, true
}

// SetDisks sets field value
func (o *VmwareObjectDiskModel) SetDisks(v []string) {
	o.Disks = v
}

// GetRemoveFromVMConfiguration returns the RemoveFromVMConfiguration field value if set, zero value otherwise.
func (o *VmwareObjectDiskModel) GetRemoveFromVMConfiguration() bool {
	if o == nil || o.RemoveFromVMConfiguration == nil {
		var ret bool
		return ret
	}
	return *o.RemoveFromVMConfiguration
}

// GetRemoveFromVMConfigurationOk returns a tuple with the RemoveFromVMConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareObjectDiskModel) GetRemoveFromVMConfigurationOk() (*bool, bool) {
	if o == nil || o.RemoveFromVMConfiguration == nil {
		return nil, false
	}
	return o.RemoveFromVMConfiguration, true
}

// HasRemoveFromVMConfiguration returns a boolean if a field has been set.
func (o *VmwareObjectDiskModel) HasRemoveFromVMConfiguration() bool {
	if o != nil && o.RemoveFromVMConfiguration != nil {
		return true
	}

	return false
}

// SetRemoveFromVMConfiguration gets a reference to the given bool and assigns it to the RemoveFromVMConfiguration field.
func (o *VmwareObjectDiskModel) SetRemoveFromVMConfiguration(v bool) {
	o.RemoveFromVMConfiguration = &v
}

func (o VmwareObjectDiskModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["vmObject"] = o.VmObject
	}
	if true {
		toSerialize["disksToProcess"] = o.DisksToProcess
	}
	if true {
		toSerialize["disks"] = o.Disks
	}
	if o.RemoveFromVMConfiguration != nil {
		toSerialize["removeFromVMConfiguration"] = o.RemoveFromVMConfiguration
	}
	return json.Marshal(toSerialize)
}

type NullableVmwareObjectDiskModel struct {
	value *VmwareObjectDiskModel
	isSet bool
}

func (v NullableVmwareObjectDiskModel) Get() *VmwareObjectDiskModel {
	return v.value
}

func (v *NullableVmwareObjectDiskModel) Set(val *VmwareObjectDiskModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareObjectDiskModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareObjectDiskModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareObjectDiskModel(val *VmwareObjectDiskModel) *NullableVmwareObjectDiskModel {
	return &NullableVmwareObjectDiskModel{value: val, isSet: true}
}

func (v NullableVmwareObjectDiskModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareObjectDiskModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

