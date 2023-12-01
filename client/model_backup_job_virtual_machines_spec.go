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

// BackupJobVirtualMachinesSpec Arrays of objects that you want to back up or exclude from the backup.
type BackupJobVirtualMachinesSpec struct {
	// Array of objects that you want to back up.
	Includes []VmwareObjectModel `json:"includes"`
	Excludes *BackupJobExclusionsSpec `json:"excludes,omitempty"`
}

// NewBackupJobVirtualMachinesSpec instantiates a new BackupJobVirtualMachinesSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupJobVirtualMachinesSpec(includes []VmwareObjectModel) *BackupJobVirtualMachinesSpec {
	this := BackupJobVirtualMachinesSpec{}
	this.Includes = includes
	return &this
}

// NewBackupJobVirtualMachinesSpecWithDefaults instantiates a new BackupJobVirtualMachinesSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupJobVirtualMachinesSpecWithDefaults() *BackupJobVirtualMachinesSpec {
	this := BackupJobVirtualMachinesSpec{}
	return &this
}

// GetIncludes returns the Includes field value
func (o *BackupJobVirtualMachinesSpec) GetIncludes() []VmwareObjectModel {
	if o == nil {
		var ret []VmwareObjectModel
		return ret
	}

	return o.Includes
}

// GetIncludesOk returns a tuple with the Includes field value
// and a boolean to check if the value has been set.
func (o *BackupJobVirtualMachinesSpec) GetIncludesOk() ([]VmwareObjectModel, bool) {
	if o == nil {
    return nil, false
	}
	return o.Includes, true
}

// SetIncludes sets field value
func (o *BackupJobVirtualMachinesSpec) SetIncludes(v []VmwareObjectModel) {
	o.Includes = v
}

// GetExcludes returns the Excludes field value if set, zero value otherwise.
func (o *BackupJobVirtualMachinesSpec) GetExcludes() BackupJobExclusionsSpec {
	if o == nil || isNil(o.Excludes) {
		var ret BackupJobExclusionsSpec
		return ret
	}
	return *o.Excludes
}

// GetExcludesOk returns a tuple with the Excludes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobVirtualMachinesSpec) GetExcludesOk() (*BackupJobExclusionsSpec, bool) {
	if o == nil || isNil(o.Excludes) {
    return nil, false
	}
	return o.Excludes, true
}

// HasExcludes returns a boolean if a field has been set.
func (o *BackupJobVirtualMachinesSpec) HasExcludes() bool {
	if o != nil && !isNil(o.Excludes) {
		return true
	}

	return false
}

// SetExcludes gets a reference to the given BackupJobExclusionsSpec and assigns it to the Excludes field.
func (o *BackupJobVirtualMachinesSpec) SetExcludes(v BackupJobExclusionsSpec) {
	o.Excludes = &v
}

func (o BackupJobVirtualMachinesSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["includes"] = o.Includes
	}
	if !isNil(o.Excludes) {
		toSerialize["excludes"] = o.Excludes
	}
	return json.Marshal(toSerialize)
}

type NullableBackupJobVirtualMachinesSpec struct {
	value *BackupJobVirtualMachinesSpec
	isSet bool
}

func (v NullableBackupJobVirtualMachinesSpec) Get() *BackupJobVirtualMachinesSpec {
	return v.value
}

func (v *NullableBackupJobVirtualMachinesSpec) Set(val *BackupJobVirtualMachinesSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupJobVirtualMachinesSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupJobVirtualMachinesSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupJobVirtualMachinesSpec(val *BackupJobVirtualMachinesSpec) *NullableBackupJobVirtualMachinesSpec {
	return &NullableBackupJobVirtualMachinesSpec{value: val, isSet: true}
}

func (v NullableBackupJobVirtualMachinesSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupJobVirtualMachinesSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

