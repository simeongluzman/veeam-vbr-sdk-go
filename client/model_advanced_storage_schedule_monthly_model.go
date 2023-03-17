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

// AdvancedStorageScheduleMonthlyModel Monthly schedule settings.
type AdvancedStorageScheduleMonthlyModel struct {
	// If *true*, the monthly schedule is enabled.
	IsEnabled bool `json:"isEnabled"`
	DayOfWeek *EDayOfWeek `json:"dayOfWeek,omitempty"`
	DayNumberInMonth *EDayNumberInMonth `json:"dayNumberInMonth,omitempty"`
	// Day of the month when the operation is performed.
	DayOfMonths *int32 `json:"dayOfMonths,omitempty"`
	// Months when the operation is performed.
	Months []EMonth `json:"months,omitempty"`
}

// NewAdvancedStorageScheduleMonthlyModel instantiates a new AdvancedStorageScheduleMonthlyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvancedStorageScheduleMonthlyModel(isEnabled bool) *AdvancedStorageScheduleMonthlyModel {
	this := AdvancedStorageScheduleMonthlyModel{}
	this.IsEnabled = isEnabled
	return &this
}

// NewAdvancedStorageScheduleMonthlyModelWithDefaults instantiates a new AdvancedStorageScheduleMonthlyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvancedStorageScheduleMonthlyModelWithDefaults() *AdvancedStorageScheduleMonthlyModel {
	this := AdvancedStorageScheduleMonthlyModel{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *AdvancedStorageScheduleMonthlyModel) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *AdvancedStorageScheduleMonthlyModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *AdvancedStorageScheduleMonthlyModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise.
func (o *AdvancedStorageScheduleMonthlyModel) GetDayOfWeek() EDayOfWeek {
	if o == nil || isNil(o.DayOfWeek) {
		var ret EDayOfWeek
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedStorageScheduleMonthlyModel) GetDayOfWeekOk() (*EDayOfWeek, bool) {
	if o == nil || isNil(o.DayOfWeek) {
    return nil, false
	}
	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *AdvancedStorageScheduleMonthlyModel) HasDayOfWeek() bool {
	if o != nil && !isNil(o.DayOfWeek) {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given EDayOfWeek and assigns it to the DayOfWeek field.
func (o *AdvancedStorageScheduleMonthlyModel) SetDayOfWeek(v EDayOfWeek) {
	o.DayOfWeek = &v
}

// GetDayNumberInMonth returns the DayNumberInMonth field value if set, zero value otherwise.
func (o *AdvancedStorageScheduleMonthlyModel) GetDayNumberInMonth() EDayNumberInMonth {
	if o == nil || isNil(o.DayNumberInMonth) {
		var ret EDayNumberInMonth
		return ret
	}
	return *o.DayNumberInMonth
}

// GetDayNumberInMonthOk returns a tuple with the DayNumberInMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedStorageScheduleMonthlyModel) GetDayNumberInMonthOk() (*EDayNumberInMonth, bool) {
	if o == nil || isNil(o.DayNumberInMonth) {
    return nil, false
	}
	return o.DayNumberInMonth, true
}

// HasDayNumberInMonth returns a boolean if a field has been set.
func (o *AdvancedStorageScheduleMonthlyModel) HasDayNumberInMonth() bool {
	if o != nil && !isNil(o.DayNumberInMonth) {
		return true
	}

	return false
}

// SetDayNumberInMonth gets a reference to the given EDayNumberInMonth and assigns it to the DayNumberInMonth field.
func (o *AdvancedStorageScheduleMonthlyModel) SetDayNumberInMonth(v EDayNumberInMonth) {
	o.DayNumberInMonth = &v
}

// GetDayOfMonths returns the DayOfMonths field value if set, zero value otherwise.
func (o *AdvancedStorageScheduleMonthlyModel) GetDayOfMonths() int32 {
	if o == nil || isNil(o.DayOfMonths) {
		var ret int32
		return ret
	}
	return *o.DayOfMonths
}

// GetDayOfMonthsOk returns a tuple with the DayOfMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedStorageScheduleMonthlyModel) GetDayOfMonthsOk() (*int32, bool) {
	if o == nil || isNil(o.DayOfMonths) {
    return nil, false
	}
	return o.DayOfMonths, true
}

// HasDayOfMonths returns a boolean if a field has been set.
func (o *AdvancedStorageScheduleMonthlyModel) HasDayOfMonths() bool {
	if o != nil && !isNil(o.DayOfMonths) {
		return true
	}

	return false
}

// SetDayOfMonths gets a reference to the given int32 and assigns it to the DayOfMonths field.
func (o *AdvancedStorageScheduleMonthlyModel) SetDayOfMonths(v int32) {
	o.DayOfMonths = &v
}

// GetMonths returns the Months field value if set, zero value otherwise.
func (o *AdvancedStorageScheduleMonthlyModel) GetMonths() []EMonth {
	if o == nil || isNil(o.Months) {
		var ret []EMonth
		return ret
	}
	return o.Months
}

// GetMonthsOk returns a tuple with the Months field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedStorageScheduleMonthlyModel) GetMonthsOk() ([]EMonth, bool) {
	if o == nil || isNil(o.Months) {
    return nil, false
	}
	return o.Months, true
}

// HasMonths returns a boolean if a field has been set.
func (o *AdvancedStorageScheduleMonthlyModel) HasMonths() bool {
	if o != nil && !isNil(o.Months) {
		return true
	}

	return false
}

// SetMonths gets a reference to the given []EMonth and assigns it to the Months field.
func (o *AdvancedStorageScheduleMonthlyModel) SetMonths(v []EMonth) {
	o.Months = v
}

func (o AdvancedStorageScheduleMonthlyModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !isNil(o.DayOfWeek) {
		toSerialize["dayOfWeek"] = o.DayOfWeek
	}
	if !isNil(o.DayNumberInMonth) {
		toSerialize["dayNumberInMonth"] = o.DayNumberInMonth
	}
	if !isNil(o.DayOfMonths) {
		toSerialize["dayOfMonths"] = o.DayOfMonths
	}
	if !isNil(o.Months) {
		toSerialize["months"] = o.Months
	}
	return json.Marshal(toSerialize)
}

type NullableAdvancedStorageScheduleMonthlyModel struct {
	value *AdvancedStorageScheduleMonthlyModel
	isSet bool
}

func (v NullableAdvancedStorageScheduleMonthlyModel) Get() *AdvancedStorageScheduleMonthlyModel {
	return v.value
}

func (v *NullableAdvancedStorageScheduleMonthlyModel) Set(val *AdvancedStorageScheduleMonthlyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvancedStorageScheduleMonthlyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvancedStorageScheduleMonthlyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvancedStorageScheduleMonthlyModel(val *AdvancedStorageScheduleMonthlyModel) *NullableAdvancedStorageScheduleMonthlyModel {
	return &NullableAdvancedStorageScheduleMonthlyModel{value: val, isSet: true}
}

func (v NullableAdvancedStorageScheduleMonthlyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvancedStorageScheduleMonthlyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


