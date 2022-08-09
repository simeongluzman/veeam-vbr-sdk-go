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

// BackupWindowSettingModel Time scheme.
type BackupWindowSettingModel struct {
	// Array of per-day schemes.
	Days []BackupWindowDayHoursModel `json:"days"`
}

// NewBackupWindowSettingModel instantiates a new BackupWindowSettingModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupWindowSettingModel(days []BackupWindowDayHoursModel, ) *BackupWindowSettingModel {
	this := BackupWindowSettingModel{}
	this.Days = days
	return &this
}

// NewBackupWindowSettingModelWithDefaults instantiates a new BackupWindowSettingModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupWindowSettingModelWithDefaults() *BackupWindowSettingModel {
	this := BackupWindowSettingModel{}
	return &this
}

// GetDays returns the Days field value
func (o *BackupWindowSettingModel) GetDays() []BackupWindowDayHoursModel {
	if o == nil  {
		var ret []BackupWindowDayHoursModel
		return ret
	}

	return o.Days
}

// GetDaysOk returns a tuple with the Days field value
// and a boolean to check if the value has been set.
func (o *BackupWindowSettingModel) GetDaysOk() (*[]BackupWindowDayHoursModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Days, true
}

// SetDays sets field value
func (o *BackupWindowSettingModel) SetDays(v []BackupWindowDayHoursModel) {
	o.Days = v
}

func (o BackupWindowSettingModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["days"] = o.Days
	}
	return json.Marshal(toSerialize)
}

type NullableBackupWindowSettingModel struct {
	value *BackupWindowSettingModel
	isSet bool
}

func (v NullableBackupWindowSettingModel) Get() *BackupWindowSettingModel {
	return v.value
}

func (v *NullableBackupWindowSettingModel) Set(val *BackupWindowSettingModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupWindowSettingModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupWindowSettingModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupWindowSettingModel(val *BackupWindowSettingModel) *NullableBackupWindowSettingModel {
	return &NullableBackupWindowSettingModel{value: val, isSet: true}
}

func (v NullableBackupWindowSettingModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupWindowSettingModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

