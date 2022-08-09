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

// ConfigBackupNotificationsModel Notification settings.
type ConfigBackupNotificationsModel struct {
	// If *true*, SNMP traps are enabled for this job.
	SNMPEnabled bool `json:"SNMPEnabled"`
	SMTPSettings *ConfigBackupSMTPSettigsModel `json:"SMTPSettings,omitempty"`
}

// NewConfigBackupNotificationsModel instantiates a new ConfigBackupNotificationsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigBackupNotificationsModel(sNMPEnabled bool, ) *ConfigBackupNotificationsModel {
	this := ConfigBackupNotificationsModel{}
	this.SNMPEnabled = sNMPEnabled
	return &this
}

// NewConfigBackupNotificationsModelWithDefaults instantiates a new ConfigBackupNotificationsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigBackupNotificationsModelWithDefaults() *ConfigBackupNotificationsModel {
	this := ConfigBackupNotificationsModel{}
	return &this
}

// GetSNMPEnabled returns the SNMPEnabled field value
func (o *ConfigBackupNotificationsModel) GetSNMPEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.SNMPEnabled
}

// GetSNMPEnabledOk returns a tuple with the SNMPEnabled field value
// and a boolean to check if the value has been set.
func (o *ConfigBackupNotificationsModel) GetSNMPEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SNMPEnabled, true
}

// SetSNMPEnabled sets field value
func (o *ConfigBackupNotificationsModel) SetSNMPEnabled(v bool) {
	o.SNMPEnabled = v
}

// GetSMTPSettings returns the SMTPSettings field value if set, zero value otherwise.
func (o *ConfigBackupNotificationsModel) GetSMTPSettings() ConfigBackupSMTPSettigsModel {
	if o == nil || o.SMTPSettings == nil {
		var ret ConfigBackupSMTPSettigsModel
		return ret
	}
	return *o.SMTPSettings
}

// GetSMTPSettingsOk returns a tuple with the SMTPSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigBackupNotificationsModel) GetSMTPSettingsOk() (*ConfigBackupSMTPSettigsModel, bool) {
	if o == nil || o.SMTPSettings == nil {
		return nil, false
	}
	return o.SMTPSettings, true
}

// HasSMTPSettings returns a boolean if a field has been set.
func (o *ConfigBackupNotificationsModel) HasSMTPSettings() bool {
	if o != nil && o.SMTPSettings != nil {
		return true
	}

	return false
}

// SetSMTPSettings gets a reference to the given ConfigBackupSMTPSettigsModel and assigns it to the SMTPSettings field.
func (o *ConfigBackupNotificationsModel) SetSMTPSettings(v ConfigBackupSMTPSettigsModel) {
	o.SMTPSettings = &v
}

func (o ConfigBackupNotificationsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["SNMPEnabled"] = o.SNMPEnabled
	}
	if o.SMTPSettings != nil {
		toSerialize["SMTPSettings"] = o.SMTPSettings
	}
	return json.Marshal(toSerialize)
}

type NullableConfigBackupNotificationsModel struct {
	value *ConfigBackupNotificationsModel
	isSet bool
}

func (v NullableConfigBackupNotificationsModel) Get() *ConfigBackupNotificationsModel {
	return v.value
}

func (v *NullableConfigBackupNotificationsModel) Set(val *ConfigBackupNotificationsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigBackupNotificationsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigBackupNotificationsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigBackupNotificationsModel(val *ConfigBackupNotificationsModel) *NullableConfigBackupNotificationsModel {
	return &NullableConfigBackupNotificationsModel{value: val, isSet: true}
}

func (v NullableConfigBackupNotificationsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigBackupNotificationsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

