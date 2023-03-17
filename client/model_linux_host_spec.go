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

// LinuxHostSpec struct for LinuxHostSpec
type LinuxHostSpec struct {
	SshSettings *LinuxHostSSHSettingsModel `json:"sshSettings,omitempty"`
	// SSH key fingerprint used to verify the server identity. For details on how to get the fingerprint, see [Request TLS Certificate or SSH Fingerprint](#tag/Connection/operation/GetConnectionCertificate).
	SshFingerprint string `json:"sshFingerprint"`
}

// NewLinuxHostSpec instantiates a new LinuxHostSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinuxHostSpec(sshFingerprint string) *LinuxHostSpec {
	this := LinuxHostSpec{}
	this.SshFingerprint = sshFingerprint
	return &this
}

// NewLinuxHostSpecWithDefaults instantiates a new LinuxHostSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinuxHostSpecWithDefaults() *LinuxHostSpec {
	this := LinuxHostSpec{}
	return &this
}

// GetSshSettings returns the SshSettings field value if set, zero value otherwise.
func (o *LinuxHostSpec) GetSshSettings() LinuxHostSSHSettingsModel {
	if o == nil || isNil(o.SshSettings) {
		var ret LinuxHostSSHSettingsModel
		return ret
	}
	return *o.SshSettings
}

// GetSshSettingsOk returns a tuple with the SshSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxHostSpec) GetSshSettingsOk() (*LinuxHostSSHSettingsModel, bool) {
	if o == nil || isNil(o.SshSettings) {
    return nil, false
	}
	return o.SshSettings, true
}

// HasSshSettings returns a boolean if a field has been set.
func (o *LinuxHostSpec) HasSshSettings() bool {
	if o != nil && !isNil(o.SshSettings) {
		return true
	}

	return false
}

// SetSshSettings gets a reference to the given LinuxHostSSHSettingsModel and assigns it to the SshSettings field.
func (o *LinuxHostSpec) SetSshSettings(v LinuxHostSSHSettingsModel) {
	o.SshSettings = &v
}

// GetSshFingerprint returns the SshFingerprint field value
func (o *LinuxHostSpec) GetSshFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SshFingerprint
}

// GetSshFingerprintOk returns a tuple with the SshFingerprint field value
// and a boolean to check if the value has been set.
func (o *LinuxHostSpec) GetSshFingerprintOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SshFingerprint, true
}

// SetSshFingerprint sets field value
func (o *LinuxHostSpec) SetSshFingerprint(v string) {
	o.SshFingerprint = v
}

func (o LinuxHostSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SshSettings) {
		toSerialize["sshSettings"] = o.SshSettings
	}
	if true {
		toSerialize["sshFingerprint"] = o.SshFingerprint
	}
	return json.Marshal(toSerialize)
}

type NullableLinuxHostSpec struct {
	value *LinuxHostSpec
	isSet bool
}

func (v NullableLinuxHostSpec) Get() *LinuxHostSpec {
	return v.value
}

func (v *NullableLinuxHostSpec) Set(val *LinuxHostSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableLinuxHostSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableLinuxHostSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinuxHostSpec(val *LinuxHostSpec) *NullableLinuxHostSpec {
	return &NullableLinuxHostSpec{value: val, isSet: true}
}

func (v NullableLinuxHostSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinuxHostSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


