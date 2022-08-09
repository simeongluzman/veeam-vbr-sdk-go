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

// BackupOracleSettingsModel Oracle archived log settings.
type BackupOracleSettingsModel struct {
	// If *true*, Veeam Backup & Replication uses credentials specified in the guest processing settings.
	UseGuestCredentials bool `json:"useGuestCredentials"`
	// ID of the credentials record that is used if `useGuestCredentials` is *false*.
	CredentialsId *string `json:"credentialsId,omitempty"`
	ArchiveLogs EBackupOracleLogsSettings `json:"archiveLogs"`
	// Time period in hours to keep archived logs. This parameter should be specified if the `EBackupOracleLogsSettings` value is *deleteExpiredHours*.
	DeleteHoursCount *int32 `json:"deleteHoursCount,omitempty"`
	// Maximum size for archived logs in GB. This parameter should be specified if the `EBackupOracleLogsSettings` value is *deleteExpiredGBs*.
	DeleteGBsCount *int32 `json:"deleteGBsCount,omitempty"`
	// If *true*, archived logs are backed up.
	BackupLogs *bool `json:"backupLogs,omitempty"`
	// Frequency of archived log backup, in minutes.
	BackupMinsCount *int32 `json:"backupMinsCount,omitempty"`
	RetainLogBackups *ERetainLogBackupsType `json:"retainLogBackups,omitempty"`
	// Number of days to keep archived logs.
	KeepDaysCount *int32 `json:"keepDaysCount,omitempty"`
	LogShippingServers *BackupLogShippingServersModel `json:"logShippingServers,omitempty"`
}

// NewBackupOracleSettingsModel instantiates a new BackupOracleSettingsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupOracleSettingsModel(useGuestCredentials bool, archiveLogs EBackupOracleLogsSettings, ) *BackupOracleSettingsModel {
	this := BackupOracleSettingsModel{}
	this.UseGuestCredentials = useGuestCredentials
	this.ArchiveLogs = archiveLogs
	return &this
}

// NewBackupOracleSettingsModelWithDefaults instantiates a new BackupOracleSettingsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupOracleSettingsModelWithDefaults() *BackupOracleSettingsModel {
	this := BackupOracleSettingsModel{}
	return &this
}

// GetUseGuestCredentials returns the UseGuestCredentials field value
func (o *BackupOracleSettingsModel) GetUseGuestCredentials() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.UseGuestCredentials
}

// GetUseGuestCredentialsOk returns a tuple with the UseGuestCredentials field value
// and a boolean to check if the value has been set.
func (o *BackupOracleSettingsModel) GetUseGuestCredentialsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UseGuestCredentials, true
}

// SetUseGuestCredentials sets field value
func (o *BackupOracleSettingsModel) SetUseGuestCredentials(v bool) {
	o.UseGuestCredentials = v
}

// GetCredentialsId returns the CredentialsId field value if set, zero value otherwise.
func (o *BackupOracleSettingsModel) GetCredentialsId() string {
	if o == nil || o.CredentialsId == nil {
		var ret string
		return ret
	}
	return *o.CredentialsId
}

// GetCredentialsIdOk returns a tuple with the CredentialsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOracleSettingsModel) GetCredentialsIdOk() (*string, bool) {
	if o == nil || o.CredentialsId == nil {
		return nil, false
	}
	return o.CredentialsId, true
}

// HasCredentialsId returns a boolean if a field has been set.
func (o *BackupOracleSettingsModel) HasCredentialsId() bool {
	if o != nil && o.CredentialsId != nil {
		return true
	}

	return false
}

// SetCredentialsId gets a reference to the given string and assigns it to the CredentialsId field.
func (o *BackupOracleSettingsModel) SetCredentialsId(v string) {
	o.CredentialsId = &v
}

// GetArchiveLogs returns the ArchiveLogs field value
func (o *BackupOracleSettingsModel) GetArchiveLogs() EBackupOracleLogsSettings {
	if o == nil  {
		var ret EBackupOracleLogsSettings
		return ret
	}

	return o.ArchiveLogs
}

// GetArchiveLogsOk returns a tuple with the ArchiveLogs field value
// and a boolean to check if the value has been set.
func (o *BackupOracleSettingsModel) GetArchiveLogsOk() (*EBackupOracleLogsSettings, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ArchiveLogs, true
}

// SetArchiveLogs sets field value
func (o *BackupOracleSettingsModel) SetArchiveLogs(v EBackupOracleLogsSettings) {
	o.ArchiveLogs = v
}

// GetDeleteHoursCount returns the DeleteHoursCount field value if set, zero value otherwise.
func (o *BackupOracleSettingsModel) GetDeleteHoursCount() int32 {
	if o == nil || o.DeleteHoursCount == nil {
		var ret int32
		return ret
	}
	return *o.DeleteHoursCount
}

// GetDeleteHoursCountOk returns a tuple with the DeleteHoursCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOracleSettingsModel) GetDeleteHoursCountOk() (*int32, bool) {
	if o == nil || o.DeleteHoursCount == nil {
		return nil, false
	}
	return o.DeleteHoursCount, true
}

// HasDeleteHoursCount returns a boolean if a field has been set.
func (o *BackupOracleSettingsModel) HasDeleteHoursCount() bool {
	if o != nil && o.DeleteHoursCount != nil {
		return true
	}

	return false
}

// SetDeleteHoursCount gets a reference to the given int32 and assigns it to the DeleteHoursCount field.
func (o *BackupOracleSettingsModel) SetDeleteHoursCount(v int32) {
	o.DeleteHoursCount = &v
}

// GetDeleteGBsCount returns the DeleteGBsCount field value if set, zero value otherwise.
func (o *BackupOracleSettingsModel) GetDeleteGBsCount() int32 {
	if o == nil || o.DeleteGBsCount == nil {
		var ret int32
		return ret
	}
	return *o.DeleteGBsCount
}

// GetDeleteGBsCountOk returns a tuple with the DeleteGBsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOracleSettingsModel) GetDeleteGBsCountOk() (*int32, bool) {
	if o == nil || o.DeleteGBsCount == nil {
		return nil, false
	}
	return o.DeleteGBsCount, true
}

// HasDeleteGBsCount returns a boolean if a field has been set.
func (o *BackupOracleSettingsModel) HasDeleteGBsCount() bool {
	if o != nil && o.DeleteGBsCount != nil {
		return true
	}

	return false
}

// SetDeleteGBsCount gets a reference to the given int32 and assigns it to the DeleteGBsCount field.
func (o *BackupOracleSettingsModel) SetDeleteGBsCount(v int32) {
	o.DeleteGBsCount = &v
}

// GetBackupLogs returns the BackupLogs field value if set, zero value otherwise.
func (o *BackupOracleSettingsModel) GetBackupLogs() bool {
	if o == nil || o.BackupLogs == nil {
		var ret bool
		return ret
	}
	return *o.BackupLogs
}

// GetBackupLogsOk returns a tuple with the BackupLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOracleSettingsModel) GetBackupLogsOk() (*bool, bool) {
	if o == nil || o.BackupLogs == nil {
		return nil, false
	}
	return o.BackupLogs, true
}

// HasBackupLogs returns a boolean if a field has been set.
func (o *BackupOracleSettingsModel) HasBackupLogs() bool {
	if o != nil && o.BackupLogs != nil {
		return true
	}

	return false
}

// SetBackupLogs gets a reference to the given bool and assigns it to the BackupLogs field.
func (o *BackupOracleSettingsModel) SetBackupLogs(v bool) {
	o.BackupLogs = &v
}

// GetBackupMinsCount returns the BackupMinsCount field value if set, zero value otherwise.
func (o *BackupOracleSettingsModel) GetBackupMinsCount() int32 {
	if o == nil || o.BackupMinsCount == nil {
		var ret int32
		return ret
	}
	return *o.BackupMinsCount
}

// GetBackupMinsCountOk returns a tuple with the BackupMinsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOracleSettingsModel) GetBackupMinsCountOk() (*int32, bool) {
	if o == nil || o.BackupMinsCount == nil {
		return nil, false
	}
	return o.BackupMinsCount, true
}

// HasBackupMinsCount returns a boolean if a field has been set.
func (o *BackupOracleSettingsModel) HasBackupMinsCount() bool {
	if o != nil && o.BackupMinsCount != nil {
		return true
	}

	return false
}

// SetBackupMinsCount gets a reference to the given int32 and assigns it to the BackupMinsCount field.
func (o *BackupOracleSettingsModel) SetBackupMinsCount(v int32) {
	o.BackupMinsCount = &v
}

// GetRetainLogBackups returns the RetainLogBackups field value if set, zero value otherwise.
func (o *BackupOracleSettingsModel) GetRetainLogBackups() ERetainLogBackupsType {
	if o == nil || o.RetainLogBackups == nil {
		var ret ERetainLogBackupsType
		return ret
	}
	return *o.RetainLogBackups
}

// GetRetainLogBackupsOk returns a tuple with the RetainLogBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOracleSettingsModel) GetRetainLogBackupsOk() (*ERetainLogBackupsType, bool) {
	if o == nil || o.RetainLogBackups == nil {
		return nil, false
	}
	return o.RetainLogBackups, true
}

// HasRetainLogBackups returns a boolean if a field has been set.
func (o *BackupOracleSettingsModel) HasRetainLogBackups() bool {
	if o != nil && o.RetainLogBackups != nil {
		return true
	}

	return false
}

// SetRetainLogBackups gets a reference to the given ERetainLogBackupsType and assigns it to the RetainLogBackups field.
func (o *BackupOracleSettingsModel) SetRetainLogBackups(v ERetainLogBackupsType) {
	o.RetainLogBackups = &v
}

// GetKeepDaysCount returns the KeepDaysCount field value if set, zero value otherwise.
func (o *BackupOracleSettingsModel) GetKeepDaysCount() int32 {
	if o == nil || o.KeepDaysCount == nil {
		var ret int32
		return ret
	}
	return *o.KeepDaysCount
}

// GetKeepDaysCountOk returns a tuple with the KeepDaysCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOracleSettingsModel) GetKeepDaysCountOk() (*int32, bool) {
	if o == nil || o.KeepDaysCount == nil {
		return nil, false
	}
	return o.KeepDaysCount, true
}

// HasKeepDaysCount returns a boolean if a field has been set.
func (o *BackupOracleSettingsModel) HasKeepDaysCount() bool {
	if o != nil && o.KeepDaysCount != nil {
		return true
	}

	return false
}

// SetKeepDaysCount gets a reference to the given int32 and assigns it to the KeepDaysCount field.
func (o *BackupOracleSettingsModel) SetKeepDaysCount(v int32) {
	o.KeepDaysCount = &v
}

// GetLogShippingServers returns the LogShippingServers field value if set, zero value otherwise.
func (o *BackupOracleSettingsModel) GetLogShippingServers() BackupLogShippingServersModel {
	if o == nil || o.LogShippingServers == nil {
		var ret BackupLogShippingServersModel
		return ret
	}
	return *o.LogShippingServers
}

// GetLogShippingServersOk returns a tuple with the LogShippingServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOracleSettingsModel) GetLogShippingServersOk() (*BackupLogShippingServersModel, bool) {
	if o == nil || o.LogShippingServers == nil {
		return nil, false
	}
	return o.LogShippingServers, true
}

// HasLogShippingServers returns a boolean if a field has been set.
func (o *BackupOracleSettingsModel) HasLogShippingServers() bool {
	if o != nil && o.LogShippingServers != nil {
		return true
	}

	return false
}

// SetLogShippingServers gets a reference to the given BackupLogShippingServersModel and assigns it to the LogShippingServers field.
func (o *BackupOracleSettingsModel) SetLogShippingServers(v BackupLogShippingServersModel) {
	o.LogShippingServers = &v
}

func (o BackupOracleSettingsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["useGuestCredentials"] = o.UseGuestCredentials
	}
	if o.CredentialsId != nil {
		toSerialize["credentialsId"] = o.CredentialsId
	}
	if true {
		toSerialize["archiveLogs"] = o.ArchiveLogs
	}
	if o.DeleteHoursCount != nil {
		toSerialize["deleteHoursCount"] = o.DeleteHoursCount
	}
	if o.DeleteGBsCount != nil {
		toSerialize["deleteGBsCount"] = o.DeleteGBsCount
	}
	if o.BackupLogs != nil {
		toSerialize["backupLogs"] = o.BackupLogs
	}
	if o.BackupMinsCount != nil {
		toSerialize["backupMinsCount"] = o.BackupMinsCount
	}
	if o.RetainLogBackups != nil {
		toSerialize["retainLogBackups"] = o.RetainLogBackups
	}
	if o.KeepDaysCount != nil {
		toSerialize["keepDaysCount"] = o.KeepDaysCount
	}
	if o.LogShippingServers != nil {
		toSerialize["logShippingServers"] = o.LogShippingServers
	}
	return json.Marshal(toSerialize)
}

type NullableBackupOracleSettingsModel struct {
	value *BackupOracleSettingsModel
	isSet bool
}

func (v NullableBackupOracleSettingsModel) Get() *BackupOracleSettingsModel {
	return v.value
}

func (v *NullableBackupOracleSettingsModel) Set(val *BackupOracleSettingsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupOracleSettingsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupOracleSettingsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupOracleSettingsModel(val *BackupOracleSettingsModel) *NullableBackupOracleSettingsModel {
	return &NullableBackupOracleSettingsModel{value: val, isSet: true}
}

func (v NullableBackupOracleSettingsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupOracleSettingsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

