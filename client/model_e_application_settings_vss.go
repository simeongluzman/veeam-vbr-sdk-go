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

// EApplicationSettingsVSS Behavior scenario for application-aware processing.
type EApplicationSettingsVSS string

// List of EApplicationSettingsVSS
const (
	EAPPLICATIONSETTINGSVSS_REQUIRE_SUCCESS EApplicationSettingsVSS = "requireSuccess"
	EAPPLICATIONSETTINGSVSS_IGNORE_FAILURES EApplicationSettingsVSS = "ignoreFailures"
	EAPPLICATIONSETTINGSVSS_DISABLED EApplicationSettingsVSS = "disabled"
)

func (v *EApplicationSettingsVSS) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EApplicationSettingsVSS(value)
	for _, existing := range []EApplicationSettingsVSS{ "requireSuccess", "ignoreFailures", "disabled",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EApplicationSettingsVSS", value)
}

// Ptr returns reference to EApplicationSettingsVSS value
func (v EApplicationSettingsVSS) Ptr() *EApplicationSettingsVSS {
	return &v
}

type NullableEApplicationSettingsVSS struct {
	value *EApplicationSettingsVSS
	isSet bool
}

func (v NullableEApplicationSettingsVSS) Get() *EApplicationSettingsVSS {
	return v.value
}

func (v *NullableEApplicationSettingsVSS) Set(val *EApplicationSettingsVSS) {
	v.value = val
	v.isSet = true
}

func (v NullableEApplicationSettingsVSS) IsSet() bool {
	return v.isSet
}

func (v *NullableEApplicationSettingsVSS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEApplicationSettingsVSS(val *EApplicationSettingsVSS) *NullableEApplicationSettingsVSS {
	return &NullableEApplicationSettingsVSS{value: val, isSet: true}
}

func (v NullableEApplicationSettingsVSS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEApplicationSettingsVSS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
