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

// AmazonCloudCredentialsModel struct for AmazonCloudCredentialsModel
type AmazonCloudCredentialsModel struct {
	CloudCredentialsModel
	// ID of the access key.
	AccessKey string `json:"accessKey"`
	// Tag used to identify the cloud credentials record.
	Tag *string `json:"tag,omitempty"`
}

// NewAmazonCloudCredentialsModel instantiates a new AmazonCloudCredentialsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonCloudCredentialsModel(accessKey string, id string, type_ ECloudCredentialsType) *AmazonCloudCredentialsModel {
	this := AmazonCloudCredentialsModel{}
	this.Id = id
	this.Type = type_
	this.AccessKey = accessKey
	return &this
}

// NewAmazonCloudCredentialsModelWithDefaults instantiates a new AmazonCloudCredentialsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonCloudCredentialsModelWithDefaults() *AmazonCloudCredentialsModel {
	this := AmazonCloudCredentialsModel{}
	return &this
}

// GetAccessKey returns the AccessKey field value
func (o *AmazonCloudCredentialsModel) GetAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value
// and a boolean to check if the value has been set.
func (o *AmazonCloudCredentialsModel) GetAccessKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccessKey, true
}

// SetAccessKey sets field value
func (o *AmazonCloudCredentialsModel) SetAccessKey(v string) {
	o.AccessKey = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *AmazonCloudCredentialsModel) GetTag() string {
	if o == nil || isNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonCloudCredentialsModel) GetTagOk() (*string, bool) {
	if o == nil || isNil(o.Tag) {
    return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *AmazonCloudCredentialsModel) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *AmazonCloudCredentialsModel) SetTag(v string) {
	o.Tag = &v
}

func (o AmazonCloudCredentialsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudCredentialsModel, errCloudCredentialsModel := json.Marshal(o.CloudCredentialsModel)
	if errCloudCredentialsModel != nil {
		return []byte{}, errCloudCredentialsModel
	}
	errCloudCredentialsModel = json.Unmarshal([]byte(serializedCloudCredentialsModel), &toSerialize)
	if errCloudCredentialsModel != nil {
		return []byte{}, errCloudCredentialsModel
	}
	if true {
		toSerialize["accessKey"] = o.AccessKey
	}
	if !isNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonCloudCredentialsModel struct {
	value *AmazonCloudCredentialsModel
	isSet bool
}

func (v NullableAmazonCloudCredentialsModel) Get() *AmazonCloudCredentialsModel {
	return v.value
}

func (v *NullableAmazonCloudCredentialsModel) Set(val *AmazonCloudCredentialsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonCloudCredentialsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonCloudCredentialsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonCloudCredentialsModel(val *AmazonCloudCredentialsModel) *NullableAmazonCloudCredentialsModel {
	return &NullableAmazonCloudCredentialsModel{value: val, isSet: true}
}

func (v NullableAmazonCloudCredentialsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonCloudCredentialsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


