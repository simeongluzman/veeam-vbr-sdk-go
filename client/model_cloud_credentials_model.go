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

// CloudCredentialsModel struct for CloudCredentialsModel
type CloudCredentialsModel struct {
	// ID of the cloud credentials record.
	Id string `json:"id"`
	// Description of the cloud credentials record.
	Description *string `json:"description,omitempty"`
	Type ECloudCredentialsType `json:"type"`
}

// NewCloudCredentialsModel instantiates a new CloudCredentialsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudCredentialsModel(id string, type_ ECloudCredentialsType) *CloudCredentialsModel {
	this := CloudCredentialsModel{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewCloudCredentialsModelWithDefaults instantiates a new CloudCredentialsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudCredentialsModelWithDefaults() *CloudCredentialsModel {
	this := CloudCredentialsModel{}
	return &this
}

// GetId returns the Id field value
func (o *CloudCredentialsModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CloudCredentialsModel) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CloudCredentialsModel) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudCredentialsModel) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCredentialsModel) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudCredentialsModel) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudCredentialsModel) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *CloudCredentialsModel) GetType() ECloudCredentialsType {
	if o == nil {
		var ret ECloudCredentialsType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CloudCredentialsModel) GetTypeOk() (*ECloudCredentialsType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CloudCredentialsModel) SetType(v ECloudCredentialsType) {
	o.Type = v
}

func (o CloudCredentialsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCloudCredentialsModel struct {
	value *CloudCredentialsModel
	isSet bool
}

func (v NullableCloudCredentialsModel) Get() *CloudCredentialsModel {
	return v.value
}

func (v *NullableCloudCredentialsModel) Set(val *CloudCredentialsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudCredentialsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudCredentialsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudCredentialsModel(val *CloudCredentialsModel) *NullableCloudCredentialsModel {
	return &NullableCloudCredentialsModel{value: val, isSet: true}
}

func (v NullableCloudCredentialsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudCredentialsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


