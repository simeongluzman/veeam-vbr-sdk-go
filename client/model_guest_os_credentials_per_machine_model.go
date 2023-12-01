/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Some request and response bodies refer to reusable schema objects that are defined in the `schemas` section of the REST API specification. Schemas, in turn, may inherit a part of their properties by referring to other schemas, and be polymorphic by referring to multiple alternate schemas.

API version: 1.1-rev0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GuestOsCredentialsPerMachineModel struct for GuestOsCredentialsPerMachineModel
type GuestOsCredentialsPerMachineModel struct {
	// Credentials ID for a Microsoft Windows VM.
	WindowsCredsId *string `json:"windowsCredsId,omitempty"`
	// Credentials ID for a Linux VM.
	LinuxCredsId *string `json:"linuxCredsId,omitempty"`
	VmObject VmwareObjectModel `json:"vmObject"`
}

// NewGuestOsCredentialsPerMachineModel instantiates a new GuestOsCredentialsPerMachineModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuestOsCredentialsPerMachineModel(vmObject VmwareObjectModel) *GuestOsCredentialsPerMachineModel {
	this := GuestOsCredentialsPerMachineModel{}
	this.VmObject = vmObject
	return &this
}

// NewGuestOsCredentialsPerMachineModelWithDefaults instantiates a new GuestOsCredentialsPerMachineModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuestOsCredentialsPerMachineModelWithDefaults() *GuestOsCredentialsPerMachineModel {
	this := GuestOsCredentialsPerMachineModel{}
	return &this
}

// GetWindowsCredsId returns the WindowsCredsId field value if set, zero value otherwise.
func (o *GuestOsCredentialsPerMachineModel) GetWindowsCredsId() string {
	if o == nil || isNil(o.WindowsCredsId) {
		var ret string
		return ret
	}
	return *o.WindowsCredsId
}

// GetWindowsCredsIdOk returns a tuple with the WindowsCredsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestOsCredentialsPerMachineModel) GetWindowsCredsIdOk() (*string, bool) {
	if o == nil || isNil(o.WindowsCredsId) {
    return nil, false
	}
	return o.WindowsCredsId, true
}

// HasWindowsCredsId returns a boolean if a field has been set.
func (o *GuestOsCredentialsPerMachineModel) HasWindowsCredsId() bool {
	if o != nil && !isNil(o.WindowsCredsId) {
		return true
	}

	return false
}

// SetWindowsCredsId gets a reference to the given string and assigns it to the WindowsCredsId field.
func (o *GuestOsCredentialsPerMachineModel) SetWindowsCredsId(v string) {
	o.WindowsCredsId = &v
}

// GetLinuxCredsId returns the LinuxCredsId field value if set, zero value otherwise.
func (o *GuestOsCredentialsPerMachineModel) GetLinuxCredsId() string {
	if o == nil || isNil(o.LinuxCredsId) {
		var ret string
		return ret
	}
	return *o.LinuxCredsId
}

// GetLinuxCredsIdOk returns a tuple with the LinuxCredsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestOsCredentialsPerMachineModel) GetLinuxCredsIdOk() (*string, bool) {
	if o == nil || isNil(o.LinuxCredsId) {
    return nil, false
	}
	return o.LinuxCredsId, true
}

// HasLinuxCredsId returns a boolean if a field has been set.
func (o *GuestOsCredentialsPerMachineModel) HasLinuxCredsId() bool {
	if o != nil && !isNil(o.LinuxCredsId) {
		return true
	}

	return false
}

// SetLinuxCredsId gets a reference to the given string and assigns it to the LinuxCredsId field.
func (o *GuestOsCredentialsPerMachineModel) SetLinuxCredsId(v string) {
	o.LinuxCredsId = &v
}

// GetVmObject returns the VmObject field value
func (o *GuestOsCredentialsPerMachineModel) GetVmObject() VmwareObjectModel {
	if o == nil {
		var ret VmwareObjectModel
		return ret
	}

	return o.VmObject
}

// GetVmObjectOk returns a tuple with the VmObject field value
// and a boolean to check if the value has been set.
func (o *GuestOsCredentialsPerMachineModel) GetVmObjectOk() (*VmwareObjectModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VmObject, true
}

// SetVmObject sets field value
func (o *GuestOsCredentialsPerMachineModel) SetVmObject(v VmwareObjectModel) {
	o.VmObject = v
}

func (o GuestOsCredentialsPerMachineModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WindowsCredsId) {
		toSerialize["windowsCredsId"] = o.WindowsCredsId
	}
	if !isNil(o.LinuxCredsId) {
		toSerialize["linuxCredsId"] = o.LinuxCredsId
	}
	if true {
		toSerialize["vmObject"] = o.VmObject
	}
	return json.Marshal(toSerialize)
}

type NullableGuestOsCredentialsPerMachineModel struct {
	value *GuestOsCredentialsPerMachineModel
	isSet bool
}

func (v NullableGuestOsCredentialsPerMachineModel) Get() *GuestOsCredentialsPerMachineModel {
	return v.value
}

func (v *NullableGuestOsCredentialsPerMachineModel) Set(val *GuestOsCredentialsPerMachineModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGuestOsCredentialsPerMachineModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGuestOsCredentialsPerMachineModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuestOsCredentialsPerMachineModel(val *GuestOsCredentialsPerMachineModel) *NullableGuestOsCredentialsPerMachineModel {
	return &NullableGuestOsCredentialsPerMachineModel{value: val, isSet: true}
}

func (v NullableGuestOsCredentialsPerMachineModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuestOsCredentialsPerMachineModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

