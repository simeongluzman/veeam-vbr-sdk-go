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

// GuestFileSystemIndexingModel VM guest OS file indexing.
type GuestFileSystemIndexingModel struct {
	IsEnabled bool `json:"isEnabled"`
	// Array of VMs with guest OS file indexing options.
	IndexingSettings *[]BackupIndexingSettingsModel `json:"indexingSettings,omitempty"`
}

// NewGuestFileSystemIndexingModel instantiates a new GuestFileSystemIndexingModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuestFileSystemIndexingModel(isEnabled bool, ) *GuestFileSystemIndexingModel {
	this := GuestFileSystemIndexingModel{}
	this.IsEnabled = isEnabled
	return &this
}

// NewGuestFileSystemIndexingModelWithDefaults instantiates a new GuestFileSystemIndexingModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuestFileSystemIndexingModelWithDefaults() *GuestFileSystemIndexingModel {
	this := GuestFileSystemIndexingModel{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *GuestFileSystemIndexingModel) GetIsEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *GuestFileSystemIndexingModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *GuestFileSystemIndexingModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetIndexingSettings returns the IndexingSettings field value if set, zero value otherwise.
func (o *GuestFileSystemIndexingModel) GetIndexingSettings() []BackupIndexingSettingsModel {
	if o == nil || o.IndexingSettings == nil {
		var ret []BackupIndexingSettingsModel
		return ret
	}
	return *o.IndexingSettings
}

// GetIndexingSettingsOk returns a tuple with the IndexingSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestFileSystemIndexingModel) GetIndexingSettingsOk() (*[]BackupIndexingSettingsModel, bool) {
	if o == nil || o.IndexingSettings == nil {
		return nil, false
	}
	return o.IndexingSettings, true
}

// HasIndexingSettings returns a boolean if a field has been set.
func (o *GuestFileSystemIndexingModel) HasIndexingSettings() bool {
	if o != nil && o.IndexingSettings != nil {
		return true
	}

	return false
}

// SetIndexingSettings gets a reference to the given []BackupIndexingSettingsModel and assigns it to the IndexingSettings field.
func (o *GuestFileSystemIndexingModel) SetIndexingSettings(v []BackupIndexingSettingsModel) {
	o.IndexingSettings = &v
}

func (o GuestFileSystemIndexingModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.IndexingSettings != nil {
		toSerialize["indexingSettings"] = o.IndexingSettings
	}
	return json.Marshal(toSerialize)
}

type NullableGuestFileSystemIndexingModel struct {
	value *GuestFileSystemIndexingModel
	isSet bool
}

func (v NullableGuestFileSystemIndexingModel) Get() *GuestFileSystemIndexingModel {
	return v.value
}

func (v *NullableGuestFileSystemIndexingModel) Set(val *GuestFileSystemIndexingModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGuestFileSystemIndexingModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGuestFileSystemIndexingModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuestFileSystemIndexingModel(val *GuestFileSystemIndexingModel) *NullableGuestFileSystemIndexingModel {
	return &NullableGuestFileSystemIndexingModel{value: val, isSet: true}
}

func (v NullableGuestFileSystemIndexingModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuestFileSystemIndexingModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

