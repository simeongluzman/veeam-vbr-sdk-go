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

// ServicesModel struct for ServicesModel
type ServicesModel struct {
	// Name of the service.
	Name string `json:"name"`
	// Port used to communicate with the service.
	Port int32 `json:"port"`
}

// NewServicesModel instantiates a new ServicesModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesModel(name string, port int32) *ServicesModel {
	this := ServicesModel{}
	this.Name = name
	this.Port = port
	return &this
}

// NewServicesModelWithDefaults instantiates a new ServicesModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesModelWithDefaults() *ServicesModel {
	this := ServicesModel{}
	return &this
}

// GetName returns the Name field value
func (o *ServicesModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServicesModel) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServicesModel) SetName(v string) {
	o.Name = v
}

// GetPort returns the Port field value
func (o *ServicesModel) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ServicesModel) GetPortOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ServicesModel) SetPort(v int32) {
	o.Port = v
}

func (o ServicesModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableServicesModel struct {
	value *ServicesModel
	isSet bool
}

func (v NullableServicesModel) Get() *ServicesModel {
	return v.value
}

func (v *NullableServicesModel) Set(val *ServicesModel) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesModel) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesModel(val *ServicesModel) *NullableServicesModel {
	return &NullableServicesModel{value: val, isSet: true}
}

func (v NullableServicesModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


