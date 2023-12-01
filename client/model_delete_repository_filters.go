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

// DeleteRepositoryFilters struct for DeleteRepositoryFilters
type DeleteRepositoryFilters struct {
	DeleteBackups *bool `json:"deleteBackups,omitempty"`
}

// NewDeleteRepositoryFilters instantiates a new DeleteRepositoryFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRepositoryFilters() *DeleteRepositoryFilters {
	this := DeleteRepositoryFilters{}
	return &this
}

// NewDeleteRepositoryFiltersWithDefaults instantiates a new DeleteRepositoryFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRepositoryFiltersWithDefaults() *DeleteRepositoryFilters {
	this := DeleteRepositoryFilters{}
	return &this
}

// GetDeleteBackups returns the DeleteBackups field value if set, zero value otherwise.
func (o *DeleteRepositoryFilters) GetDeleteBackups() bool {
	if o == nil || isNil(o.DeleteBackups) {
		var ret bool
		return ret
	}
	return *o.DeleteBackups
}

// GetDeleteBackupsOk returns a tuple with the DeleteBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRepositoryFilters) GetDeleteBackupsOk() (*bool, bool) {
	if o == nil || isNil(o.DeleteBackups) {
    return nil, false
	}
	return o.DeleteBackups, true
}

// HasDeleteBackups returns a boolean if a field has been set.
func (o *DeleteRepositoryFilters) HasDeleteBackups() bool {
	if o != nil && !isNil(o.DeleteBackups) {
		return true
	}

	return false
}

// SetDeleteBackups gets a reference to the given bool and assigns it to the DeleteBackups field.
func (o *DeleteRepositoryFilters) SetDeleteBackups(v bool) {
	o.DeleteBackups = &v
}

func (o DeleteRepositoryFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DeleteBackups) {
		toSerialize["deleteBackups"] = o.DeleteBackups
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteRepositoryFilters struct {
	value *DeleteRepositoryFilters
	isSet bool
}

func (v NullableDeleteRepositoryFilters) Get() *DeleteRepositoryFilters {
	return v.value
}

func (v *NullableDeleteRepositoryFilters) Set(val *DeleteRepositoryFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRepositoryFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRepositoryFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRepositoryFilters(val *DeleteRepositoryFilters) *NullableDeleteRepositoryFilters {
	return &NullableDeleteRepositoryFilters{value: val, isSet: true}
}

func (v NullableDeleteRepositoryFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRepositoryFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

