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

// JobExportSpec struct for JobExportSpec
type JobExportSpec struct {
	// Array of job IDs.
	Ids []string `json:"ids,omitempty"`
	// Array of job types.
	Types []string `json:"types,omitempty"`
	// Array of job names. Wildcard characters are supported.
	Names []string `json:"names,omitempty"`
}

// NewJobExportSpec instantiates a new JobExportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobExportSpec() *JobExportSpec {
	this := JobExportSpec{}
	return &this
}

// NewJobExportSpecWithDefaults instantiates a new JobExportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobExportSpecWithDefaults() *JobExportSpec {
	this := JobExportSpec{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *JobExportSpec) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobExportSpec) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *JobExportSpec) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *JobExportSpec) SetIds(v []string) {
	o.Ids = v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *JobExportSpec) GetTypes() []string {
	if o == nil || isNil(o.Types) {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobExportSpec) GetTypesOk() ([]string, bool) {
	if o == nil || isNil(o.Types) {
    return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *JobExportSpec) HasTypes() bool {
	if o != nil && !isNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *JobExportSpec) SetTypes(v []string) {
	o.Types = v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *JobExportSpec) GetNames() []string {
	if o == nil || isNil(o.Names) {
		var ret []string
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobExportSpec) GetNamesOk() ([]string, bool) {
	if o == nil || isNil(o.Names) {
    return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *JobExportSpec) HasNames() bool {
	if o != nil && !isNil(o.Names) {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *JobExportSpec) SetNames(v []string) {
	o.Names = v
}

func (o JobExportSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !isNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	if !isNil(o.Names) {
		toSerialize["names"] = o.Names
	}
	return json.Marshal(toSerialize)
}

type NullableJobExportSpec struct {
	value *JobExportSpec
	isSet bool
}

func (v NullableJobExportSpec) Get() *JobExportSpec {
	return v.value
}

func (v *NullableJobExportSpec) Set(val *JobExportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableJobExportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableJobExportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobExportSpec(val *JobExportSpec) *NullableJobExportSpec {
	return &NullableJobExportSpec{value: val, isSet: true}
}

func (v NullableJobExportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobExportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


