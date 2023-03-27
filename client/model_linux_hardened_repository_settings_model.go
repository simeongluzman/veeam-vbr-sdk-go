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

// LinuxHardenedRepositorySettingsModel Repository settings.
type LinuxHardenedRepositorySettingsModel struct {
	// Path to the folder where backup files are stored.
	Path *string `json:"path,omitempty"`
	// If *true*, the maximum number of concurrent tasks is limited.
	EnableTaskLimit *bool `json:"enableTaskLimit,omitempty"`
	// Maximum number of concurrent tasks.
	MaxTaskCount *int32 `json:"maxTaskCount,omitempty"`
	// If *true*, reading and writing speed is limited.
	EnableReadWriteLimit *bool `json:"enableReadWriteLimit,omitempty"`
	// Maximum rate that restricts the total speed of reading and writing data to the backup repository disk.
	ReadWriteRate *int32 `json:"readWriteRate,omitempty"`
	// If *true*, fast cloning on XFS volumes is used.
	UseFastCloningOnXFSVolumes *bool `json:"useFastCloningOnXFSVolumes,omitempty"`
	// Number of days to keep immutable backups.
	MakeRecentBackupsImmutableDays *int32 `json:"makeRecentBackupsImmutableDays,omitempty"`
	AdvancedSettings *RepositoryAdvancedSettingsModel `json:"advancedSettings,omitempty"`
}

// NewLinuxHardenedRepositorySettingsModel instantiates a new LinuxHardenedRepositorySettingsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinuxHardenedRepositorySettingsModel() *LinuxHardenedRepositorySettingsModel {
	this := LinuxHardenedRepositorySettingsModel{}
	return &this
}

// NewLinuxHardenedRepositorySettingsModelWithDefaults instantiates a new LinuxHardenedRepositorySettingsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinuxHardenedRepositorySettingsModelWithDefaults() *LinuxHardenedRepositorySettingsModel {
	this := LinuxHardenedRepositorySettingsModel{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *LinuxHardenedRepositorySettingsModel) GetPath() string {
	if o == nil || isNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxHardenedRepositorySettingsModel) GetPathOk() (*string, bool) {
	if o == nil || isNil(o.Path) {
    return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *LinuxHardenedRepositorySettingsModel) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *LinuxHardenedRepositorySettingsModel) SetPath(v string) {
	o.Path = &v
}

// GetEnableTaskLimit returns the EnableTaskLimit field value if set, zero value otherwise.
func (o *LinuxHardenedRepositorySettingsModel) GetEnableTaskLimit() bool {
	if o == nil || isNil(o.EnableTaskLimit) {
		var ret bool
		return ret
	}
	return *o.EnableTaskLimit
}

// GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxHardenedRepositorySettingsModel) GetEnableTaskLimitOk() (*bool, bool) {
	if o == nil || isNil(o.EnableTaskLimit) {
    return nil, false
	}
	return o.EnableTaskLimit, true
}

// HasEnableTaskLimit returns a boolean if a field has been set.
func (o *LinuxHardenedRepositorySettingsModel) HasEnableTaskLimit() bool {
	if o != nil && !isNil(o.EnableTaskLimit) {
		return true
	}

	return false
}

// SetEnableTaskLimit gets a reference to the given bool and assigns it to the EnableTaskLimit field.
func (o *LinuxHardenedRepositorySettingsModel) SetEnableTaskLimit(v bool) {
	o.EnableTaskLimit = &v
}

// GetMaxTaskCount returns the MaxTaskCount field value if set, zero value otherwise.
func (o *LinuxHardenedRepositorySettingsModel) GetMaxTaskCount() int32 {
	if o == nil || isNil(o.MaxTaskCount) {
		var ret int32
		return ret
	}
	return *o.MaxTaskCount
}

// GetMaxTaskCountOk returns a tuple with the MaxTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxHardenedRepositorySettingsModel) GetMaxTaskCountOk() (*int32, bool) {
	if o == nil || isNil(o.MaxTaskCount) {
    return nil, false
	}
	return o.MaxTaskCount, true
}

// HasMaxTaskCount returns a boolean if a field has been set.
func (o *LinuxHardenedRepositorySettingsModel) HasMaxTaskCount() bool {
	if o != nil && !isNil(o.MaxTaskCount) {
		return true
	}

	return false
}

// SetMaxTaskCount gets a reference to the given int32 and assigns it to the MaxTaskCount field.
func (o *LinuxHardenedRepositorySettingsModel) SetMaxTaskCount(v int32) {
	o.MaxTaskCount = &v
}

// GetEnableReadWriteLimit returns the EnableReadWriteLimit field value if set, zero value otherwise.
func (o *LinuxHardenedRepositorySettingsModel) GetEnableReadWriteLimit() bool {
	if o == nil || isNil(o.EnableReadWriteLimit) {
		var ret bool
		return ret
	}
	return *o.EnableReadWriteLimit
}

// GetEnableReadWriteLimitOk returns a tuple with the EnableReadWriteLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxHardenedRepositorySettingsModel) GetEnableReadWriteLimitOk() (*bool, bool) {
	if o == nil || isNil(o.EnableReadWriteLimit) {
    return nil, false
	}
	return o.EnableReadWriteLimit, true
}

// HasEnableReadWriteLimit returns a boolean if a field has been set.
func (o *LinuxHardenedRepositorySettingsModel) HasEnableReadWriteLimit() bool {
	if o != nil && !isNil(o.EnableReadWriteLimit) {
		return true
	}

	return false
}

// SetEnableReadWriteLimit gets a reference to the given bool and assigns it to the EnableReadWriteLimit field.
func (o *LinuxHardenedRepositorySettingsModel) SetEnableReadWriteLimit(v bool) {
	o.EnableReadWriteLimit = &v
}

// GetReadWriteRate returns the ReadWriteRate field value if set, zero value otherwise.
func (o *LinuxHardenedRepositorySettingsModel) GetReadWriteRate() int32 {
	if o == nil || isNil(o.ReadWriteRate) {
		var ret int32
		return ret
	}
	return *o.ReadWriteRate
}

// GetReadWriteRateOk returns a tuple with the ReadWriteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxHardenedRepositorySettingsModel) GetReadWriteRateOk() (*int32, bool) {
	if o == nil || isNil(o.ReadWriteRate) {
    return nil, false
	}
	return o.ReadWriteRate, true
}

// HasReadWriteRate returns a boolean if a field has been set.
func (o *LinuxHardenedRepositorySettingsModel) HasReadWriteRate() bool {
	if o != nil && !isNil(o.ReadWriteRate) {
		return true
	}

	return false
}

// SetReadWriteRate gets a reference to the given int32 and assigns it to the ReadWriteRate field.
func (o *LinuxHardenedRepositorySettingsModel) SetReadWriteRate(v int32) {
	o.ReadWriteRate = &v
}

// GetUseFastCloningOnXFSVolumes returns the UseFastCloningOnXFSVolumes field value if set, zero value otherwise.
func (o *LinuxHardenedRepositorySettingsModel) GetUseFastCloningOnXFSVolumes() bool {
	if o == nil || isNil(o.UseFastCloningOnXFSVolumes) {
		var ret bool
		return ret
	}
	return *o.UseFastCloningOnXFSVolumes
}

// GetUseFastCloningOnXFSVolumesOk returns a tuple with the UseFastCloningOnXFSVolumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxHardenedRepositorySettingsModel) GetUseFastCloningOnXFSVolumesOk() (*bool, bool) {
	if o == nil || isNil(o.UseFastCloningOnXFSVolumes) {
    return nil, false
	}
	return o.UseFastCloningOnXFSVolumes, true
}

// HasUseFastCloningOnXFSVolumes returns a boolean if a field has been set.
func (o *LinuxHardenedRepositorySettingsModel) HasUseFastCloningOnXFSVolumes() bool {
	if o != nil && !isNil(o.UseFastCloningOnXFSVolumes) {
		return true
	}

	return false
}

// SetUseFastCloningOnXFSVolumes gets a reference to the given bool and assigns it to the UseFastCloningOnXFSVolumes field.
func (o *LinuxHardenedRepositorySettingsModel) SetUseFastCloningOnXFSVolumes(v bool) {
	o.UseFastCloningOnXFSVolumes = &v
}

// GetMakeRecentBackupsImmutableDays returns the MakeRecentBackupsImmutableDays field value if set, zero value otherwise.
func (o *LinuxHardenedRepositorySettingsModel) GetMakeRecentBackupsImmutableDays() int32 {
	if o == nil || isNil(o.MakeRecentBackupsImmutableDays) {
		var ret int32
		return ret
	}
	return *o.MakeRecentBackupsImmutableDays
}

// GetMakeRecentBackupsImmutableDaysOk returns a tuple with the MakeRecentBackupsImmutableDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxHardenedRepositorySettingsModel) GetMakeRecentBackupsImmutableDaysOk() (*int32, bool) {
	if o == nil || isNil(o.MakeRecentBackupsImmutableDays) {
    return nil, false
	}
	return o.MakeRecentBackupsImmutableDays, true
}

// HasMakeRecentBackupsImmutableDays returns a boolean if a field has been set.
func (o *LinuxHardenedRepositorySettingsModel) HasMakeRecentBackupsImmutableDays() bool {
	if o != nil && !isNil(o.MakeRecentBackupsImmutableDays) {
		return true
	}

	return false
}

// SetMakeRecentBackupsImmutableDays gets a reference to the given int32 and assigns it to the MakeRecentBackupsImmutableDays field.
func (o *LinuxHardenedRepositorySettingsModel) SetMakeRecentBackupsImmutableDays(v int32) {
	o.MakeRecentBackupsImmutableDays = &v
}

// GetAdvancedSettings returns the AdvancedSettings field value if set, zero value otherwise.
func (o *LinuxHardenedRepositorySettingsModel) GetAdvancedSettings() RepositoryAdvancedSettingsModel {
	if o == nil || isNil(o.AdvancedSettings) {
		var ret RepositoryAdvancedSettingsModel
		return ret
	}
	return *o.AdvancedSettings
}

// GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxHardenedRepositorySettingsModel) GetAdvancedSettingsOk() (*RepositoryAdvancedSettingsModel, bool) {
	if o == nil || isNil(o.AdvancedSettings) {
    return nil, false
	}
	return o.AdvancedSettings, true
}

// HasAdvancedSettings returns a boolean if a field has been set.
func (o *LinuxHardenedRepositorySettingsModel) HasAdvancedSettings() bool {
	if o != nil && !isNil(o.AdvancedSettings) {
		return true
	}

	return false
}

// SetAdvancedSettings gets a reference to the given RepositoryAdvancedSettingsModel and assigns it to the AdvancedSettings field.
func (o *LinuxHardenedRepositorySettingsModel) SetAdvancedSettings(v RepositoryAdvancedSettingsModel) {
	o.AdvancedSettings = &v
}

func (o LinuxHardenedRepositorySettingsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !isNil(o.EnableTaskLimit) {
		toSerialize["enableTaskLimit"] = o.EnableTaskLimit
	}
	if !isNil(o.MaxTaskCount) {
		toSerialize["maxTaskCount"] = o.MaxTaskCount
	}
	if !isNil(o.EnableReadWriteLimit) {
		toSerialize["enableReadWriteLimit"] = o.EnableReadWriteLimit
	}
	if !isNil(o.ReadWriteRate) {
		toSerialize["readWriteRate"] = o.ReadWriteRate
	}
	if !isNil(o.UseFastCloningOnXFSVolumes) {
		toSerialize["useFastCloningOnXFSVolumes"] = o.UseFastCloningOnXFSVolumes
	}
	if !isNil(o.MakeRecentBackupsImmutableDays) {
		toSerialize["makeRecentBackupsImmutableDays"] = o.MakeRecentBackupsImmutableDays
	}
	if !isNil(o.AdvancedSettings) {
		toSerialize["advancedSettings"] = o.AdvancedSettings
	}
	return json.Marshal(toSerialize)
}

type NullableLinuxHardenedRepositorySettingsModel struct {
	value *LinuxHardenedRepositorySettingsModel
	isSet bool
}

func (v NullableLinuxHardenedRepositorySettingsModel) Get() *LinuxHardenedRepositorySettingsModel {
	return v.value
}

func (v *NullableLinuxHardenedRepositorySettingsModel) Set(val *LinuxHardenedRepositorySettingsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLinuxHardenedRepositorySettingsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLinuxHardenedRepositorySettingsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinuxHardenedRepositorySettingsModel(val *LinuxHardenedRepositorySettingsModel) *NullableLinuxHardenedRepositorySettingsModel {
	return &NullableLinuxHardenedRepositorySettingsModel{value: val, isSet: true}
}

func (v NullableLinuxHardenedRepositorySettingsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinuxHardenedRepositorySettingsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

