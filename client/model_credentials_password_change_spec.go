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

// CredentialsPasswordChangeSpec struct for CredentialsPasswordChangeSpec
type CredentialsPasswordChangeSpec struct {
	// New password.
	Password string `json:"password"`
}

// NewCredentialsPasswordChangeSpec instantiates a new CredentialsPasswordChangeSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsPasswordChangeSpec(password string, ) *CredentialsPasswordChangeSpec {
	this := CredentialsPasswordChangeSpec{}
	this.Password = password
	return &this
}

// NewCredentialsPasswordChangeSpecWithDefaults instantiates a new CredentialsPasswordChangeSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsPasswordChangeSpecWithDefaults() *CredentialsPasswordChangeSpec {
	this := CredentialsPasswordChangeSpec{}
	return &this
}

// GetPassword returns the Password field value
func (o *CredentialsPasswordChangeSpec) GetPassword() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CredentialsPasswordChangeSpec) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CredentialsPasswordChangeSpec) SetPassword(v string) {
	o.Password = v
}

func (o CredentialsPasswordChangeSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableCredentialsPasswordChangeSpec struct {
	value *CredentialsPasswordChangeSpec
	isSet bool
}

func (v NullableCredentialsPasswordChangeSpec) Get() *CredentialsPasswordChangeSpec {
	return v.value
}

func (v *NullableCredentialsPasswordChangeSpec) Set(val *CredentialsPasswordChangeSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsPasswordChangeSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsPasswordChangeSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsPasswordChangeSpec(val *CredentialsPasswordChangeSpec) *NullableCredentialsPasswordChangeSpec {
	return &NullableCredentialsPasswordChangeSpec{value: val, isSet: true}
}

func (v NullableCredentialsPasswordChangeSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsPasswordChangeSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

