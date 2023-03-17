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

// AzureStorageCloudCredentialsSpecAllOf struct for AzureStorageCloudCredentialsSpecAllOf
type AzureStorageCloudCredentialsSpecAllOf struct {
	// Name of the Azure storage account.
	Account string `json:"account"`
	// Shared key of the Azure storage account.
	SharedKey string `json:"sharedKey"`
	// Tag used to identify the cloud credentials record.
	Tag *string `json:"tag,omitempty"`
}

// NewAzureStorageCloudCredentialsSpecAllOf instantiates a new AzureStorageCloudCredentialsSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureStorageCloudCredentialsSpecAllOf(account string, sharedKey string) *AzureStorageCloudCredentialsSpecAllOf {
	this := AzureStorageCloudCredentialsSpecAllOf{}
	this.Account = account
	this.SharedKey = sharedKey
	return &this
}

// NewAzureStorageCloudCredentialsSpecAllOfWithDefaults instantiates a new AzureStorageCloudCredentialsSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureStorageCloudCredentialsSpecAllOfWithDefaults() *AzureStorageCloudCredentialsSpecAllOf {
	this := AzureStorageCloudCredentialsSpecAllOf{}
	return &this
}

// GetAccount returns the Account field value
func (o *AzureStorageCloudCredentialsSpecAllOf) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *AzureStorageCloudCredentialsSpecAllOf) GetAccountOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *AzureStorageCloudCredentialsSpecAllOf) SetAccount(v string) {
	o.Account = v
}

// GetSharedKey returns the SharedKey field value
func (o *AzureStorageCloudCredentialsSpecAllOf) GetSharedKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SharedKey
}

// GetSharedKeyOk returns a tuple with the SharedKey field value
// and a boolean to check if the value has been set.
func (o *AzureStorageCloudCredentialsSpecAllOf) GetSharedKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SharedKey, true
}

// SetSharedKey sets field value
func (o *AzureStorageCloudCredentialsSpecAllOf) SetSharedKey(v string) {
	o.SharedKey = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *AzureStorageCloudCredentialsSpecAllOf) GetTag() string {
	if o == nil || isNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureStorageCloudCredentialsSpecAllOf) GetTagOk() (*string, bool) {
	if o == nil || isNil(o.Tag) {
    return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *AzureStorageCloudCredentialsSpecAllOf) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *AzureStorageCloudCredentialsSpecAllOf) SetTag(v string) {
	o.Tag = &v
}

func (o AzureStorageCloudCredentialsSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["sharedKey"] = o.SharedKey
	}
	if !isNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableAzureStorageCloudCredentialsSpecAllOf struct {
	value *AzureStorageCloudCredentialsSpecAllOf
	isSet bool
}

func (v NullableAzureStorageCloudCredentialsSpecAllOf) Get() *AzureStorageCloudCredentialsSpecAllOf {
	return v.value
}

func (v *NullableAzureStorageCloudCredentialsSpecAllOf) Set(val *AzureStorageCloudCredentialsSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureStorageCloudCredentialsSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureStorageCloudCredentialsSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureStorageCloudCredentialsSpecAllOf(val *AzureStorageCloudCredentialsSpecAllOf) *NullableAzureStorageCloudCredentialsSpecAllOf {
	return &NullableAzureStorageCloudCredentialsSpecAllOf{value: val, isSet: true}
}

func (v NullableAzureStorageCloudCredentialsSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureStorageCloudCredentialsSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


