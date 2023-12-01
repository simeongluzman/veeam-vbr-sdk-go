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

// ScaleOutRepositoriesResult struct for ScaleOutRepositoriesResult
type ScaleOutRepositoriesResult struct {
	// Array of scale-out backup repositories.
	Data []ScaleOutRepositoryModel `json:"data"`
	Pagination PaginationResult `json:"pagination"`
}

// NewScaleOutRepositoriesResult instantiates a new ScaleOutRepositoriesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScaleOutRepositoriesResult(data []ScaleOutRepositoryModel, pagination PaginationResult) *ScaleOutRepositoriesResult {
	this := ScaleOutRepositoriesResult{}
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewScaleOutRepositoriesResultWithDefaults instantiates a new ScaleOutRepositoriesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScaleOutRepositoriesResultWithDefaults() *ScaleOutRepositoriesResult {
	this := ScaleOutRepositoriesResult{}
	return &this
}

// GetData returns the Data field value
func (o *ScaleOutRepositoriesResult) GetData() []ScaleOutRepositoryModel {
	if o == nil {
		var ret []ScaleOutRepositoryModel
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ScaleOutRepositoriesResult) GetDataOk() ([]ScaleOutRepositoryModel, bool) {
	if o == nil {
    return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ScaleOutRepositoriesResult) SetData(v []ScaleOutRepositoryModel) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *ScaleOutRepositoriesResult) GetPagination() PaginationResult {
	if o == nil {
		var ret PaginationResult
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ScaleOutRepositoriesResult) GetPaginationOk() (*PaginationResult, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ScaleOutRepositoriesResult) SetPagination(v PaginationResult) {
	o.Pagination = v
}

func (o ScaleOutRepositoriesResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableScaleOutRepositoriesResult struct {
	value *ScaleOutRepositoriesResult
	isSet bool
}

func (v NullableScaleOutRepositoriesResult) Get() *ScaleOutRepositoriesResult {
	return v.value
}

func (v *NullableScaleOutRepositoriesResult) Set(val *ScaleOutRepositoriesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableScaleOutRepositoriesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableScaleOutRepositoriesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScaleOutRepositoriesResult(val *ScaleOutRepositoriesResult) *NullableScaleOutRepositoriesResult {
	return &NullableScaleOutRepositoriesResult{value: val, isSet: true}
}

func (v NullableScaleOutRepositoriesResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScaleOutRepositoriesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

