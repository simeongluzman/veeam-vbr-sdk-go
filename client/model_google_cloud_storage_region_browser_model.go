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

// GoogleCloudStorageRegionBrowserModel struct for GoogleCloudStorageRegionBrowserModel
type GoogleCloudStorageRegionBrowserModel struct {
	// Region ID.
	Id *string `json:"id,omitempty"`
	// Region name.
	Name *string `json:"name,omitempty"`
	// Array of buckets located in the region.
	Buckets []GoogleCloudStorageBucketBrowserModel `json:"buckets,omitempty"`
}

// NewGoogleCloudStorageRegionBrowserModel instantiates a new GoogleCloudStorageRegionBrowserModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleCloudStorageRegionBrowserModel() *GoogleCloudStorageRegionBrowserModel {
	this := GoogleCloudStorageRegionBrowserModel{}
	return &this
}

// NewGoogleCloudStorageRegionBrowserModelWithDefaults instantiates a new GoogleCloudStorageRegionBrowserModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudStorageRegionBrowserModelWithDefaults() *GoogleCloudStorageRegionBrowserModel {
	this := GoogleCloudStorageRegionBrowserModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GoogleCloudStorageRegionBrowserModel) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudStorageRegionBrowserModel) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GoogleCloudStorageRegionBrowserModel) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GoogleCloudStorageRegionBrowserModel) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GoogleCloudStorageRegionBrowserModel) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudStorageRegionBrowserModel) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GoogleCloudStorageRegionBrowserModel) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GoogleCloudStorageRegionBrowserModel) SetName(v string) {
	o.Name = &v
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *GoogleCloudStorageRegionBrowserModel) GetBuckets() []GoogleCloudStorageBucketBrowserModel {
	if o == nil || isNil(o.Buckets) {
		var ret []GoogleCloudStorageBucketBrowserModel
		return ret
	}
	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudStorageRegionBrowserModel) GetBucketsOk() ([]GoogleCloudStorageBucketBrowserModel, bool) {
	if o == nil || isNil(o.Buckets) {
    return nil, false
	}
	return o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *GoogleCloudStorageRegionBrowserModel) HasBuckets() bool {
	if o != nil && !isNil(o.Buckets) {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given []GoogleCloudStorageBucketBrowserModel and assigns it to the Buckets field.
func (o *GoogleCloudStorageRegionBrowserModel) SetBuckets(v []GoogleCloudStorageBucketBrowserModel) {
	o.Buckets = v
}

func (o GoogleCloudStorageRegionBrowserModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Buckets) {
		toSerialize["buckets"] = o.Buckets
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleCloudStorageRegionBrowserModel struct {
	value *GoogleCloudStorageRegionBrowserModel
	isSet bool
}

func (v NullableGoogleCloudStorageRegionBrowserModel) Get() *GoogleCloudStorageRegionBrowserModel {
	return v.value
}

func (v *NullableGoogleCloudStorageRegionBrowserModel) Set(val *GoogleCloudStorageRegionBrowserModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleCloudStorageRegionBrowserModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleCloudStorageRegionBrowserModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleCloudStorageRegionBrowserModel(val *GoogleCloudStorageRegionBrowserModel) *NullableGoogleCloudStorageRegionBrowserModel {
	return &NullableGoogleCloudStorageRegionBrowserModel{value: val, isSet: true}
}

func (v NullableGoogleCloudStorageRegionBrowserModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleCloudStorageRegionBrowserModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


