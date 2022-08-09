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

// ViBackupObjectModelAllOf struct for ViBackupObjectModelAllOf
type ViBackupObjectModelAllOf struct {
	// ID of the virtual infrastructure object: mo-ref or ID, depending on the virtualization platform. 
	ObjectId string `json:"objectId"`
	ViType EVmwareInventoryType `json:"viType"`
	// Path to the object.
	Path *string `json:"path,omitempty"`
}

// NewViBackupObjectModelAllOf instantiates a new ViBackupObjectModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViBackupObjectModelAllOf(objectId string, viType EVmwareInventoryType, ) *ViBackupObjectModelAllOf {
	this := ViBackupObjectModelAllOf{}
	this.ObjectId = objectId
	this.ViType = viType
	return &this
}

// NewViBackupObjectModelAllOfWithDefaults instantiates a new ViBackupObjectModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViBackupObjectModelAllOfWithDefaults() *ViBackupObjectModelAllOf {
	this := ViBackupObjectModelAllOf{}
	return &this
}

// GetObjectId returns the ObjectId field value
func (o *ViBackupObjectModelAllOf) GetObjectId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *ViBackupObjectModelAllOf) GetObjectIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *ViBackupObjectModelAllOf) SetObjectId(v string) {
	o.ObjectId = v
}

// GetViType returns the ViType field value
func (o *ViBackupObjectModelAllOf) GetViType() EVmwareInventoryType {
	if o == nil  {
		var ret EVmwareInventoryType
		return ret
	}

	return o.ViType
}

// GetViTypeOk returns a tuple with the ViType field value
// and a boolean to check if the value has been set.
func (o *ViBackupObjectModelAllOf) GetViTypeOk() (*EVmwareInventoryType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ViType, true
}

// SetViType sets field value
func (o *ViBackupObjectModelAllOf) SetViType(v EVmwareInventoryType) {
	o.ViType = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ViBackupObjectModelAllOf) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViBackupObjectModelAllOf) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ViBackupObjectModelAllOf) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ViBackupObjectModelAllOf) SetPath(v string) {
	o.Path = &v
}

func (o ViBackupObjectModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objectId"] = o.ObjectId
	}
	if true {
		toSerialize["viType"] = o.ViType
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableViBackupObjectModelAllOf struct {
	value *ViBackupObjectModelAllOf
	isSet bool
}

func (v NullableViBackupObjectModelAllOf) Get() *ViBackupObjectModelAllOf {
	return v.value
}

func (v *NullableViBackupObjectModelAllOf) Set(val *ViBackupObjectModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableViBackupObjectModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableViBackupObjectModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViBackupObjectModelAllOf(val *ViBackupObjectModelAllOf) *NullableViBackupObjectModelAllOf {
	return &NullableViBackupObjectModelAllOf{value: val, isSet: true}
}

func (v NullableViBackupObjectModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViBackupObjectModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

