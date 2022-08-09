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

// ScheduleRetryModel Retry options.
type ScheduleRetryModel struct {
	// If *true*, retry options are enabled.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Number of retries set for the job.
	RetryCount *int32 `json:"retryCount,omitempty"`
	// Time interval between job retries in minutes.
	AwaitMinutes *int32 `json:"awaitMinutes,omitempty"`
}

// NewScheduleRetryModel instantiates a new ScheduleRetryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleRetryModel() *ScheduleRetryModel {
	this := ScheduleRetryModel{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	return &this
}

// NewScheduleRetryModelWithDefaults instantiates a new ScheduleRetryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleRetryModelWithDefaults() *ScheduleRetryModel {
	this := ScheduleRetryModel{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *ScheduleRetryModel) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleRetryModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *ScheduleRetryModel) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *ScheduleRetryModel) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetRetryCount returns the RetryCount field value if set, zero value otherwise.
func (o *ScheduleRetryModel) GetRetryCount() int32 {
	if o == nil || o.RetryCount == nil {
		var ret int32
		return ret
	}
	return *o.RetryCount
}

// GetRetryCountOk returns a tuple with the RetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleRetryModel) GetRetryCountOk() (*int32, bool) {
	if o == nil || o.RetryCount == nil {
		return nil, false
	}
	return o.RetryCount, true
}

// HasRetryCount returns a boolean if a field has been set.
func (o *ScheduleRetryModel) HasRetryCount() bool {
	if o != nil && o.RetryCount != nil {
		return true
	}

	return false
}

// SetRetryCount gets a reference to the given int32 and assigns it to the RetryCount field.
func (o *ScheduleRetryModel) SetRetryCount(v int32) {
	o.RetryCount = &v
}

// GetAwaitMinutes returns the AwaitMinutes field value if set, zero value otherwise.
func (o *ScheduleRetryModel) GetAwaitMinutes() int32 {
	if o == nil || o.AwaitMinutes == nil {
		var ret int32
		return ret
	}
	return *o.AwaitMinutes
}

// GetAwaitMinutesOk returns a tuple with the AwaitMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleRetryModel) GetAwaitMinutesOk() (*int32, bool) {
	if o == nil || o.AwaitMinutes == nil {
		return nil, false
	}
	return o.AwaitMinutes, true
}

// HasAwaitMinutes returns a boolean if a field has been set.
func (o *ScheduleRetryModel) HasAwaitMinutes() bool {
	if o != nil && o.AwaitMinutes != nil {
		return true
	}

	return false
}

// SetAwaitMinutes gets a reference to the given int32 and assigns it to the AwaitMinutes field.
func (o *ScheduleRetryModel) SetAwaitMinutes(v int32) {
	o.AwaitMinutes = &v
}

func (o ScheduleRetryModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsEnabled != nil {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.RetryCount != nil {
		toSerialize["retryCount"] = o.RetryCount
	}
	if o.AwaitMinutes != nil {
		toSerialize["awaitMinutes"] = o.AwaitMinutes
	}
	return json.Marshal(toSerialize)
}

type NullableScheduleRetryModel struct {
	value *ScheduleRetryModel
	isSet bool
}

func (v NullableScheduleRetryModel) Get() *ScheduleRetryModel {
	return v.value
}

func (v *NullableScheduleRetryModel) Set(val *ScheduleRetryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleRetryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleRetryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleRetryModel(val *ScheduleRetryModel) *NullableScheduleRetryModel {
	return &NullableScheduleRetryModel{value: val, isSet: true}
}

func (v NullableScheduleRetryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleRetryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

