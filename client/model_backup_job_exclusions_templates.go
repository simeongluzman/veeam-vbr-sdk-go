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

// BackupJobExclusionsTemplates Array of VM templates excluded from the backup.
type BackupJobExclusionsTemplates struct {
	// If *true*, the template is included into the backup.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// If *true*, the template is excluded from the incremental backup.
	ExcludeFromIncremental *bool `json:"excludeFromIncremental,omitempty"`
}

// NewBackupJobExclusionsTemplates instantiates a new BackupJobExclusionsTemplates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupJobExclusionsTemplates() *BackupJobExclusionsTemplates {
	this := BackupJobExclusionsTemplates{}
	return &this
}

// NewBackupJobExclusionsTemplatesWithDefaults instantiates a new BackupJobExclusionsTemplates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupJobExclusionsTemplatesWithDefaults() *BackupJobExclusionsTemplates {
	this := BackupJobExclusionsTemplates{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *BackupJobExclusionsTemplates) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobExclusionsTemplates) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *BackupJobExclusionsTemplates) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *BackupJobExclusionsTemplates) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetExcludeFromIncremental returns the ExcludeFromIncremental field value if set, zero value otherwise.
func (o *BackupJobExclusionsTemplates) GetExcludeFromIncremental() bool {
	if o == nil || o.ExcludeFromIncremental == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeFromIncremental
}

// GetExcludeFromIncrementalOk returns a tuple with the ExcludeFromIncremental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobExclusionsTemplates) GetExcludeFromIncrementalOk() (*bool, bool) {
	if o == nil || o.ExcludeFromIncremental == nil {
		return nil, false
	}
	return o.ExcludeFromIncremental, true
}

// HasExcludeFromIncremental returns a boolean if a field has been set.
func (o *BackupJobExclusionsTemplates) HasExcludeFromIncremental() bool {
	if o != nil && o.ExcludeFromIncremental != nil {
		return true
	}

	return false
}

// SetExcludeFromIncremental gets a reference to the given bool and assigns it to the ExcludeFromIncremental field.
func (o *BackupJobExclusionsTemplates) SetExcludeFromIncremental(v bool) {
	o.ExcludeFromIncremental = &v
}

func (o BackupJobExclusionsTemplates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsEnabled != nil {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.ExcludeFromIncremental != nil {
		toSerialize["excludeFromIncremental"] = o.ExcludeFromIncremental
	}
	return json.Marshal(toSerialize)
}

type NullableBackupJobExclusionsTemplates struct {
	value *BackupJobExclusionsTemplates
	isSet bool
}

func (v NullableBackupJobExclusionsTemplates) Get() *BackupJobExclusionsTemplates {
	return v.value
}

func (v *NullableBackupJobExclusionsTemplates) Set(val *BackupJobExclusionsTemplates) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupJobExclusionsTemplates) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupJobExclusionsTemplates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupJobExclusionsTemplates(val *BackupJobExclusionsTemplates) *NullableBackupJobExclusionsTemplates {
	return &NullableBackupJobExclusionsTemplates{value: val, isSet: true}
}

func (v NullableBackupJobExclusionsTemplates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupJobExclusionsTemplates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

