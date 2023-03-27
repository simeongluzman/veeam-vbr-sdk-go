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

// ProxyImportSpec struct for ProxyImportSpec
type ProxyImportSpec struct {
	// Description of the backup proxy.
	Description string `json:"description"`
	Type EProxyType `json:"type"`
	Server ProxyServerSettingsImportSpec `json:"server"`
}

// NewProxyImportSpec instantiates a new ProxyImportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyImportSpec(description string, type_ EProxyType, server ProxyServerSettingsImportSpec) *ProxyImportSpec {
	this := ProxyImportSpec{}
	this.Description = description
	this.Type = type_
	this.Server = server
	return &this
}

// NewProxyImportSpecWithDefaults instantiates a new ProxyImportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyImportSpecWithDefaults() *ProxyImportSpec {
	this := ProxyImportSpec{}
	return &this
}

// GetDescription returns the Description field value
func (o *ProxyImportSpec) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ProxyImportSpec) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ProxyImportSpec) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value
func (o *ProxyImportSpec) GetType() EProxyType {
	if o == nil {
		var ret EProxyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProxyImportSpec) GetTypeOk() (*EProxyType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ProxyImportSpec) SetType(v EProxyType) {
	o.Type = v
}

// GetServer returns the Server field value
func (o *ProxyImportSpec) GetServer() ProxyServerSettingsImportSpec {
	if o == nil {
		var ret ProxyServerSettingsImportSpec
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *ProxyImportSpec) GetServerOk() (*ProxyServerSettingsImportSpec, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *ProxyImportSpec) SetServer(v ProxyServerSettingsImportSpec) {
	o.Server = v
}

func (o ProxyImportSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["server"] = o.Server
	}
	return json.Marshal(toSerialize)
}

type NullableProxyImportSpec struct {
	value *ProxyImportSpec
	isSet bool
}

func (v NullableProxyImportSpec) Get() *ProxyImportSpec {
	return v.value
}

func (v *NullableProxyImportSpec) Set(val *ProxyImportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyImportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyImportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyImportSpec(val *ProxyImportSpec) *NullableProxyImportSpec {
	return &NullableProxyImportSpec{value: val, isSet: true}
}

func (v NullableProxyImportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyImportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


