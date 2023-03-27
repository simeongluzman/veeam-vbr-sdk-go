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

// SyntheticFullSettingsModel Synthetic full backup settings.
type SyntheticFullSettingsModel struct {
	// If *true*, active full backups are enabled.
	IsEnabled bool `json:"isEnabled"`
	Weekly *AdvancedStorageScheduleWeeklyModel `json:"weekly,omitempty"`
	Monthly *AdvancedStorageScheduleMonthlyModel `json:"monthly,omitempty"`
}

// NewSyntheticFullSettingsModel instantiates a new SyntheticFullSettingsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticFullSettingsModel(isEnabled bool) *SyntheticFullSettingsModel {
	this := SyntheticFullSettingsModel{}
	this.IsEnabled = isEnabled
	return &this
}

// NewSyntheticFullSettingsModelWithDefaults instantiates a new SyntheticFullSettingsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticFullSettingsModelWithDefaults() *SyntheticFullSettingsModel {
	this := SyntheticFullSettingsModel{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *SyntheticFullSettingsModel) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *SyntheticFullSettingsModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *SyntheticFullSettingsModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetWeekly returns the Weekly field value if set, zero value otherwise.
func (o *SyntheticFullSettingsModel) GetWeekly() AdvancedStorageScheduleWeeklyModel {
	if o == nil || isNil(o.Weekly) {
		var ret AdvancedStorageScheduleWeeklyModel
		return ret
	}
	return *o.Weekly
}

// GetWeeklyOk returns a tuple with the Weekly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticFullSettingsModel) GetWeeklyOk() (*AdvancedStorageScheduleWeeklyModel, bool) {
	if o == nil || isNil(o.Weekly) {
    return nil, false
	}
	return o.Weekly, true
}

// HasWeekly returns a boolean if a field has been set.
func (o *SyntheticFullSettingsModel) HasWeekly() bool {
	if o != nil && !isNil(o.Weekly) {
		return true
	}

	return false
}

// SetWeekly gets a reference to the given AdvancedStorageScheduleWeeklyModel and assigns it to the Weekly field.
func (o *SyntheticFullSettingsModel) SetWeekly(v AdvancedStorageScheduleWeeklyModel) {
	o.Weekly = &v
}

// GetMonthly returns the Monthly field value if set, zero value otherwise.
func (o *SyntheticFullSettingsModel) GetMonthly() AdvancedStorageScheduleMonthlyModel {
	if o == nil || isNil(o.Monthly) {
		var ret AdvancedStorageScheduleMonthlyModel
		return ret
	}
	return *o.Monthly
}

// GetMonthlyOk returns a tuple with the Monthly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticFullSettingsModel) GetMonthlyOk() (*AdvancedStorageScheduleMonthlyModel, bool) {
	if o == nil || isNil(o.Monthly) {
    return nil, false
	}
	return o.Monthly, true
}

// HasMonthly returns a boolean if a field has been set.
func (o *SyntheticFullSettingsModel) HasMonthly() bool {
	if o != nil && !isNil(o.Monthly) {
		return true
	}

	return false
}

// SetMonthly gets a reference to the given AdvancedStorageScheduleMonthlyModel and assigns it to the Monthly field.
func (o *SyntheticFullSettingsModel) SetMonthly(v AdvancedStorageScheduleMonthlyModel) {
	o.Monthly = &v
}

func (o SyntheticFullSettingsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !isNil(o.Weekly) {
		toSerialize["weekly"] = o.Weekly
	}
	if !isNil(o.Monthly) {
		toSerialize["monthly"] = o.Monthly
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticFullSettingsModel struct {
	value *SyntheticFullSettingsModel
	isSet bool
}

func (v NullableSyntheticFullSettingsModel) Get() *SyntheticFullSettingsModel {
	return v.value
}

func (v *NullableSyntheticFullSettingsModel) Set(val *SyntheticFullSettingsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticFullSettingsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticFullSettingsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticFullSettingsModel(val *SyntheticFullSettingsModel) *NullableSyntheticFullSettingsModel {
	return &NullableSyntheticFullSettingsModel{value: val, isSet: true}
}

func (v NullableSyntheticFullSettingsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticFullSettingsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


