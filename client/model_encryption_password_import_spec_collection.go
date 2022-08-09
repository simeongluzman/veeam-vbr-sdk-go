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

// EncryptionPasswordImportSpecCollection struct for EncryptionPasswordImportSpecCollection
type EncryptionPasswordImportSpecCollection struct {
	// Array of encryption passwords.
	EncryptionPasswords *[]EncryptionPasswordImportSpec `json:"encryptionPasswords,omitempty"`
}

// NewEncryptionPasswordImportSpecCollection instantiates a new EncryptionPasswordImportSpecCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptionPasswordImportSpecCollection() *EncryptionPasswordImportSpecCollection {
	this := EncryptionPasswordImportSpecCollection{}
	return &this
}

// NewEncryptionPasswordImportSpecCollectionWithDefaults instantiates a new EncryptionPasswordImportSpecCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionPasswordImportSpecCollectionWithDefaults() *EncryptionPasswordImportSpecCollection {
	this := EncryptionPasswordImportSpecCollection{}
	return &this
}

// GetEncryptionPasswords returns the EncryptionPasswords field value if set, zero value otherwise.
func (o *EncryptionPasswordImportSpecCollection) GetEncryptionPasswords() []EncryptionPasswordImportSpec {
	if o == nil || o.EncryptionPasswords == nil {
		var ret []EncryptionPasswordImportSpec
		return ret
	}
	return *o.EncryptionPasswords
}

// GetEncryptionPasswordsOk returns a tuple with the EncryptionPasswords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionPasswordImportSpecCollection) GetEncryptionPasswordsOk() (*[]EncryptionPasswordImportSpec, bool) {
	if o == nil || o.EncryptionPasswords == nil {
		return nil, false
	}
	return o.EncryptionPasswords, true
}

// HasEncryptionPasswords returns a boolean if a field has been set.
func (o *EncryptionPasswordImportSpecCollection) HasEncryptionPasswords() bool {
	if o != nil && o.EncryptionPasswords != nil {
		return true
	}

	return false
}

// SetEncryptionPasswords gets a reference to the given []EncryptionPasswordImportSpec and assigns it to the EncryptionPasswords field.
func (o *EncryptionPasswordImportSpecCollection) SetEncryptionPasswords(v []EncryptionPasswordImportSpec) {
	o.EncryptionPasswords = &v
}

func (o EncryptionPasswordImportSpecCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EncryptionPasswords != nil {
		toSerialize["encryptionPasswords"] = o.EncryptionPasswords
	}
	return json.Marshal(toSerialize)
}

type NullableEncryptionPasswordImportSpecCollection struct {
	value *EncryptionPasswordImportSpecCollection
	isSet bool
}

func (v NullableEncryptionPasswordImportSpecCollection) Get() *EncryptionPasswordImportSpecCollection {
	return v.value
}

func (v *NullableEncryptionPasswordImportSpecCollection) Set(val *EncryptionPasswordImportSpecCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptionPasswordImportSpecCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptionPasswordImportSpecCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptionPasswordImportSpecCollection(val *EncryptionPasswordImportSpecCollection) *NullableEncryptionPasswordImportSpecCollection {
	return &NullableEncryptionPasswordImportSpecCollection{value: val, isSet: true}
}

func (v NullableEncryptionPasswordImportSpecCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptionPasswordImportSpecCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

