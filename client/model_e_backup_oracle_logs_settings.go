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
	"fmt"
)

// EBackupOracleLogsSettings Type of archived logs processing.
type EBackupOracleLogsSettings string

// List of EBackupOracleLogsSettings
const (
	EBACKUPORACLELOGSSETTINGS_PRESERVE EBackupOracleLogsSettings = "preserve"
	EBACKUPORACLELOGSSETTINGS_DELETE_EXPIRED_HOURS EBackupOracleLogsSettings = "deleteExpiredHours"
	EBACKUPORACLELOGSSETTINGS_DELETE_EXPIRED_GBS EBackupOracleLogsSettings = "deleteExpiredGBs"
)

func (v *EBackupOracleLogsSettings) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EBackupOracleLogsSettings(value)
	for _, existing := range []EBackupOracleLogsSettings{ "preserve", "deleteExpiredHours", "deleteExpiredGBs",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EBackupOracleLogsSettings", value)
}

// Ptr returns reference to EBackupOracleLogsSettings value
func (v EBackupOracleLogsSettings) Ptr() *EBackupOracleLogsSettings {
	return &v
}

type NullableEBackupOracleLogsSettings struct {
	value *EBackupOracleLogsSettings
	isSet bool
}

func (v NullableEBackupOracleLogsSettings) Get() *EBackupOracleLogsSettings {
	return v.value
}

func (v *NullableEBackupOracleLogsSettings) Set(val *EBackupOracleLogsSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableEBackupOracleLogsSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableEBackupOracleLogsSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEBackupOracleLogsSettings(val *EBackupOracleLogsSettings) *NullableEBackupOracleLogsSettings {
	return &NullableEBackupOracleLogsSettings{value: val, isSet: true}
}

func (v NullableEBackupOracleLogsSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEBackupOracleLogsSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
