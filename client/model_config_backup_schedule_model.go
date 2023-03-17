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

// ConfigBackupScheduleModel Scheduling settings.
type ConfigBackupScheduleModel struct {
	// If *true*, backup scheduling is enabled.
	IsEnabled bool `json:"isEnabled"`
	Daily *ScheduleDailyModel `json:"daily,omitempty"`
	Monthly *ScheduleMonthlyModel `json:"monthly,omitempty"`
}

// NewConfigBackupScheduleModel instantiates a new ConfigBackupScheduleModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigBackupScheduleModel(isEnabled bool) *ConfigBackupScheduleModel {
	this := ConfigBackupScheduleModel{}
	this.IsEnabled = isEnabled
	return &this
}

// NewConfigBackupScheduleModelWithDefaults instantiates a new ConfigBackupScheduleModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigBackupScheduleModelWithDefaults() *ConfigBackupScheduleModel {
	this := ConfigBackupScheduleModel{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *ConfigBackupScheduleModel) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *ConfigBackupScheduleModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *ConfigBackupScheduleModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetDaily returns the Daily field value if set, zero value otherwise.
func (o *ConfigBackupScheduleModel) GetDaily() ScheduleDailyModel {
	if o == nil || isNil(o.Daily) {
		var ret ScheduleDailyModel
		return ret
	}
	return *o.Daily
}

// GetDailyOk returns a tuple with the Daily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigBackupScheduleModel) GetDailyOk() (*ScheduleDailyModel, bool) {
	if o == nil || isNil(o.Daily) {
    return nil, false
	}
	return o.Daily, true
}

// HasDaily returns a boolean if a field has been set.
func (o *ConfigBackupScheduleModel) HasDaily() bool {
	if o != nil && !isNil(o.Daily) {
		return true
	}

	return false
}

// SetDaily gets a reference to the given ScheduleDailyModel and assigns it to the Daily field.
func (o *ConfigBackupScheduleModel) SetDaily(v ScheduleDailyModel) {
	o.Daily = &v
}

// GetMonthly returns the Monthly field value if set, zero value otherwise.
func (o *ConfigBackupScheduleModel) GetMonthly() ScheduleMonthlyModel {
	if o == nil || isNil(o.Monthly) {
		var ret ScheduleMonthlyModel
		return ret
	}
	return *o.Monthly
}

// GetMonthlyOk returns a tuple with the Monthly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigBackupScheduleModel) GetMonthlyOk() (*ScheduleMonthlyModel, bool) {
	if o == nil || isNil(o.Monthly) {
    return nil, false
	}
	return o.Monthly, true
}

// HasMonthly returns a boolean if a field has been set.
func (o *ConfigBackupScheduleModel) HasMonthly() bool {
	if o != nil && !isNil(o.Monthly) {
		return true
	}

	return false
}

// SetMonthly gets a reference to the given ScheduleMonthlyModel and assigns it to the Monthly field.
func (o *ConfigBackupScheduleModel) SetMonthly(v ScheduleMonthlyModel) {
	o.Monthly = &v
}

func (o ConfigBackupScheduleModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !isNil(o.Daily) {
		toSerialize["daily"] = o.Daily
	}
	if !isNil(o.Monthly) {
		toSerialize["monthly"] = o.Monthly
	}
	return json.Marshal(toSerialize)
}

type NullableConfigBackupScheduleModel struct {
	value *ConfigBackupScheduleModel
	isSet bool
}

func (v NullableConfigBackupScheduleModel) Get() *ConfigBackupScheduleModel {
	return v.value
}

func (v *NullableConfigBackupScheduleModel) Set(val *ConfigBackupScheduleModel) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigBackupScheduleModel) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigBackupScheduleModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigBackupScheduleModel(val *ConfigBackupScheduleModel) *NullableConfigBackupScheduleModel {
	return &NullableConfigBackupScheduleModel{value: val, isSet: true}
}

func (v NullableConfigBackupScheduleModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigBackupScheduleModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


