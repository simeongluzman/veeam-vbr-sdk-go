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

// BackupSQLSettingsModel Microsoft SQL Server transaction log settings.
type BackupSQLSettingsModel struct {
	LogsProcessing ESQLLogsProcessing `json:"logsProcessing"`
	// Frequency of transaction log backup, in minutes.
	BackupMinsCount *int32 `json:"backupMinsCount,omitempty"`
	RetainLogBackups *ERetainLogBackupsType `json:"retainLogBackups,omitempty"`
	// Number of days to keep transaction logs in the backup repository.
	KeepDaysCount *int32 `json:"keepDaysCount,omitempty"`
	LogShippingServers *BackupLogShippingServersModel `json:"logShippingServers,omitempty"`
}

// NewBackupSQLSettingsModel instantiates a new BackupSQLSettingsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupSQLSettingsModel(logsProcessing ESQLLogsProcessing) *BackupSQLSettingsModel {
	this := BackupSQLSettingsModel{}
	this.LogsProcessing = logsProcessing
	return &this
}

// NewBackupSQLSettingsModelWithDefaults instantiates a new BackupSQLSettingsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupSQLSettingsModelWithDefaults() *BackupSQLSettingsModel {
	this := BackupSQLSettingsModel{}
	return &this
}

// GetLogsProcessing returns the LogsProcessing field value
func (o *BackupSQLSettingsModel) GetLogsProcessing() ESQLLogsProcessing {
	if o == nil {
		var ret ESQLLogsProcessing
		return ret
	}

	return o.LogsProcessing
}

// GetLogsProcessingOk returns a tuple with the LogsProcessing field value
// and a boolean to check if the value has been set.
func (o *BackupSQLSettingsModel) GetLogsProcessingOk() (*ESQLLogsProcessing, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LogsProcessing, true
}

// SetLogsProcessing sets field value
func (o *BackupSQLSettingsModel) SetLogsProcessing(v ESQLLogsProcessing) {
	o.LogsProcessing = v
}

// GetBackupMinsCount returns the BackupMinsCount field value if set, zero value otherwise.
func (o *BackupSQLSettingsModel) GetBackupMinsCount() int32 {
	if o == nil || isNil(o.BackupMinsCount) {
		var ret int32
		return ret
	}
	return *o.BackupMinsCount
}

// GetBackupMinsCountOk returns a tuple with the BackupMinsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSQLSettingsModel) GetBackupMinsCountOk() (*int32, bool) {
	if o == nil || isNil(o.BackupMinsCount) {
    return nil, false
	}
	return o.BackupMinsCount, true
}

// HasBackupMinsCount returns a boolean if a field has been set.
func (o *BackupSQLSettingsModel) HasBackupMinsCount() bool {
	if o != nil && !isNil(o.BackupMinsCount) {
		return true
	}

	return false
}

// SetBackupMinsCount gets a reference to the given int32 and assigns it to the BackupMinsCount field.
func (o *BackupSQLSettingsModel) SetBackupMinsCount(v int32) {
	o.BackupMinsCount = &v
}

// GetRetainLogBackups returns the RetainLogBackups field value if set, zero value otherwise.
func (o *BackupSQLSettingsModel) GetRetainLogBackups() ERetainLogBackupsType {
	if o == nil || isNil(o.RetainLogBackups) {
		var ret ERetainLogBackupsType
		return ret
	}
	return *o.RetainLogBackups
}

// GetRetainLogBackupsOk returns a tuple with the RetainLogBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSQLSettingsModel) GetRetainLogBackupsOk() (*ERetainLogBackupsType, bool) {
	if o == nil || isNil(o.RetainLogBackups) {
    return nil, false
	}
	return o.RetainLogBackups, true
}

// HasRetainLogBackups returns a boolean if a field has been set.
func (o *BackupSQLSettingsModel) HasRetainLogBackups() bool {
	if o != nil && !isNil(o.RetainLogBackups) {
		return true
	}

	return false
}

// SetRetainLogBackups gets a reference to the given ERetainLogBackupsType and assigns it to the RetainLogBackups field.
func (o *BackupSQLSettingsModel) SetRetainLogBackups(v ERetainLogBackupsType) {
	o.RetainLogBackups = &v
}

// GetKeepDaysCount returns the KeepDaysCount field value if set, zero value otherwise.
func (o *BackupSQLSettingsModel) GetKeepDaysCount() int32 {
	if o == nil || isNil(o.KeepDaysCount) {
		var ret int32
		return ret
	}
	return *o.KeepDaysCount
}

// GetKeepDaysCountOk returns a tuple with the KeepDaysCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSQLSettingsModel) GetKeepDaysCountOk() (*int32, bool) {
	if o == nil || isNil(o.KeepDaysCount) {
    return nil, false
	}
	return o.KeepDaysCount, true
}

// HasKeepDaysCount returns a boolean if a field has been set.
func (o *BackupSQLSettingsModel) HasKeepDaysCount() bool {
	if o != nil && !isNil(o.KeepDaysCount) {
		return true
	}

	return false
}

// SetKeepDaysCount gets a reference to the given int32 and assigns it to the KeepDaysCount field.
func (o *BackupSQLSettingsModel) SetKeepDaysCount(v int32) {
	o.KeepDaysCount = &v
}

// GetLogShippingServers returns the LogShippingServers field value if set, zero value otherwise.
func (o *BackupSQLSettingsModel) GetLogShippingServers() BackupLogShippingServersModel {
	if o == nil || isNil(o.LogShippingServers) {
		var ret BackupLogShippingServersModel
		return ret
	}
	return *o.LogShippingServers
}

// GetLogShippingServersOk returns a tuple with the LogShippingServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSQLSettingsModel) GetLogShippingServersOk() (*BackupLogShippingServersModel, bool) {
	if o == nil || isNil(o.LogShippingServers) {
    return nil, false
	}
	return o.LogShippingServers, true
}

// HasLogShippingServers returns a boolean if a field has been set.
func (o *BackupSQLSettingsModel) HasLogShippingServers() bool {
	if o != nil && !isNil(o.LogShippingServers) {
		return true
	}

	return false
}

// SetLogShippingServers gets a reference to the given BackupLogShippingServersModel and assigns it to the LogShippingServers field.
func (o *BackupSQLSettingsModel) SetLogShippingServers(v BackupLogShippingServersModel) {
	o.LogShippingServers = &v
}

func (o BackupSQLSettingsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["logsProcessing"] = o.LogsProcessing
	}
	if !isNil(o.BackupMinsCount) {
		toSerialize["backupMinsCount"] = o.BackupMinsCount
	}
	if !isNil(o.RetainLogBackups) {
		toSerialize["retainLogBackups"] = o.RetainLogBackups
	}
	if !isNil(o.KeepDaysCount) {
		toSerialize["keepDaysCount"] = o.KeepDaysCount
	}
	if !isNil(o.LogShippingServers) {
		toSerialize["logShippingServers"] = o.LogShippingServers
	}
	return json.Marshal(toSerialize)
}

type NullableBackupSQLSettingsModel struct {
	value *BackupSQLSettingsModel
	isSet bool
}

func (v NullableBackupSQLSettingsModel) Get() *BackupSQLSettingsModel {
	return v.value
}

func (v *NullableBackupSQLSettingsModel) Set(val *BackupSQLSettingsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupSQLSettingsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupSQLSettingsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupSQLSettingsModel(val *BackupSQLSettingsModel) *NullableBackupSQLSettingsModel {
	return &NullableBackupSQLSettingsModel{value: val, isSet: true}
}

func (v NullableBackupSQLSettingsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupSQLSettingsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


