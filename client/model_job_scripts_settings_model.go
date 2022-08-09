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

// JobScriptsSettingsModel Script settings.
type JobScriptsSettingsModel struct {
	PreCommand *ScriptCommand `json:"preCommand,omitempty"`
	PostCommand *ScriptCommand `json:"postCommand,omitempty"`
	PeriodicityType *EScriptPeriodicityType `json:"periodicityType,omitempty"`
	// Number of the backup job session after which the scripts must be executed.
	RunScriptEvery *int32 `json:"runScriptEvery,omitempty"`
	// Days of the week when the scripts must be executed.
	DayOfWeek *[]EDayOfWeek `json:"dayOfWeek,omitempty"`
}

// NewJobScriptsSettingsModel instantiates a new JobScriptsSettingsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobScriptsSettingsModel() *JobScriptsSettingsModel {
	this := JobScriptsSettingsModel{}
	return &this
}

// NewJobScriptsSettingsModelWithDefaults instantiates a new JobScriptsSettingsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobScriptsSettingsModelWithDefaults() *JobScriptsSettingsModel {
	this := JobScriptsSettingsModel{}
	return &this
}

// GetPreCommand returns the PreCommand field value if set, zero value otherwise.
func (o *JobScriptsSettingsModel) GetPreCommand() ScriptCommand {
	if o == nil || o.PreCommand == nil {
		var ret ScriptCommand
		return ret
	}
	return *o.PreCommand
}

// GetPreCommandOk returns a tuple with the PreCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobScriptsSettingsModel) GetPreCommandOk() (*ScriptCommand, bool) {
	if o == nil || o.PreCommand == nil {
		return nil, false
	}
	return o.PreCommand, true
}

// HasPreCommand returns a boolean if a field has been set.
func (o *JobScriptsSettingsModel) HasPreCommand() bool {
	if o != nil && o.PreCommand != nil {
		return true
	}

	return false
}

// SetPreCommand gets a reference to the given ScriptCommand and assigns it to the PreCommand field.
func (o *JobScriptsSettingsModel) SetPreCommand(v ScriptCommand) {
	o.PreCommand = &v
}

// GetPostCommand returns the PostCommand field value if set, zero value otherwise.
func (o *JobScriptsSettingsModel) GetPostCommand() ScriptCommand {
	if o == nil || o.PostCommand == nil {
		var ret ScriptCommand
		return ret
	}
	return *o.PostCommand
}

// GetPostCommandOk returns a tuple with the PostCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobScriptsSettingsModel) GetPostCommandOk() (*ScriptCommand, bool) {
	if o == nil || o.PostCommand == nil {
		return nil, false
	}
	return o.PostCommand, true
}

// HasPostCommand returns a boolean if a field has been set.
func (o *JobScriptsSettingsModel) HasPostCommand() bool {
	if o != nil && o.PostCommand != nil {
		return true
	}

	return false
}

// SetPostCommand gets a reference to the given ScriptCommand and assigns it to the PostCommand field.
func (o *JobScriptsSettingsModel) SetPostCommand(v ScriptCommand) {
	o.PostCommand = &v
}

// GetPeriodicityType returns the PeriodicityType field value if set, zero value otherwise.
func (o *JobScriptsSettingsModel) GetPeriodicityType() EScriptPeriodicityType {
	if o == nil || o.PeriodicityType == nil {
		var ret EScriptPeriodicityType
		return ret
	}
	return *o.PeriodicityType
}

// GetPeriodicityTypeOk returns a tuple with the PeriodicityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobScriptsSettingsModel) GetPeriodicityTypeOk() (*EScriptPeriodicityType, bool) {
	if o == nil || o.PeriodicityType == nil {
		return nil, false
	}
	return o.PeriodicityType, true
}

// HasPeriodicityType returns a boolean if a field has been set.
func (o *JobScriptsSettingsModel) HasPeriodicityType() bool {
	if o != nil && o.PeriodicityType != nil {
		return true
	}

	return false
}

// SetPeriodicityType gets a reference to the given EScriptPeriodicityType and assigns it to the PeriodicityType field.
func (o *JobScriptsSettingsModel) SetPeriodicityType(v EScriptPeriodicityType) {
	o.PeriodicityType = &v
}

// GetRunScriptEvery returns the RunScriptEvery field value if set, zero value otherwise.
func (o *JobScriptsSettingsModel) GetRunScriptEvery() int32 {
	if o == nil || o.RunScriptEvery == nil {
		var ret int32
		return ret
	}
	return *o.RunScriptEvery
}

// GetRunScriptEveryOk returns a tuple with the RunScriptEvery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobScriptsSettingsModel) GetRunScriptEveryOk() (*int32, bool) {
	if o == nil || o.RunScriptEvery == nil {
		return nil, false
	}
	return o.RunScriptEvery, true
}

// HasRunScriptEvery returns a boolean if a field has been set.
func (o *JobScriptsSettingsModel) HasRunScriptEvery() bool {
	if o != nil && o.RunScriptEvery != nil {
		return true
	}

	return false
}

// SetRunScriptEvery gets a reference to the given int32 and assigns it to the RunScriptEvery field.
func (o *JobScriptsSettingsModel) SetRunScriptEvery(v int32) {
	o.RunScriptEvery = &v
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise.
func (o *JobScriptsSettingsModel) GetDayOfWeek() []EDayOfWeek {
	if o == nil || o.DayOfWeek == nil {
		var ret []EDayOfWeek
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobScriptsSettingsModel) GetDayOfWeekOk() (*[]EDayOfWeek, bool) {
	if o == nil || o.DayOfWeek == nil {
		return nil, false
	}
	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *JobScriptsSettingsModel) HasDayOfWeek() bool {
	if o != nil && o.DayOfWeek != nil {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given []EDayOfWeek and assigns it to the DayOfWeek field.
func (o *JobScriptsSettingsModel) SetDayOfWeek(v []EDayOfWeek) {
	o.DayOfWeek = &v
}

func (o JobScriptsSettingsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PreCommand != nil {
		toSerialize["preCommand"] = o.PreCommand
	}
	if o.PostCommand != nil {
		toSerialize["postCommand"] = o.PostCommand
	}
	if o.PeriodicityType != nil {
		toSerialize["periodicityType"] = o.PeriodicityType
	}
	if o.RunScriptEvery != nil {
		toSerialize["runScriptEvery"] = o.RunScriptEvery
	}
	if o.DayOfWeek != nil {
		toSerialize["dayOfWeek"] = o.DayOfWeek
	}
	return json.Marshal(toSerialize)
}

type NullableJobScriptsSettingsModel struct {
	value *JobScriptsSettingsModel
	isSet bool
}

func (v NullableJobScriptsSettingsModel) Get() *JobScriptsSettingsModel {
	return v.value
}

func (v *NullableJobScriptsSettingsModel) Set(val *JobScriptsSettingsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableJobScriptsSettingsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableJobScriptsSettingsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobScriptsSettingsModel(val *JobScriptsSettingsModel) *NullableJobScriptsSettingsModel {
	return &NullableJobScriptsSettingsModel{value: val, isSet: true}
}

func (v NullableJobScriptsSettingsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobScriptsSettingsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

