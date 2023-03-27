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

// NfsStorageModel NFS backup repository.
type NfsStorageModel struct {
	RepositoryModel
	Share NfsRepositoryShareSettingsModel `json:"share"`
	Repository NetworkRepositorySettingsModel `json:"repository"`
	MountServer MountServerSettingsModel `json:"mountServer"`
}

// NewNfsStorageModel instantiates a new NfsStorageModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfsStorageModel(share NfsRepositoryShareSettingsModel, repository NetworkRepositorySettingsModel, mountServer MountServerSettingsModel, id string, name string, description string, type_ ERepositoryType) *NfsStorageModel {
	this := NfsStorageModel{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Type = type_
	this.Share = share
	this.Repository = repository
	this.MountServer = mountServer
	return &this
}

// NewNfsStorageModelWithDefaults instantiates a new NfsStorageModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfsStorageModelWithDefaults() *NfsStorageModel {
	this := NfsStorageModel{}
	return &this
}

// GetShare returns the Share field value
func (o *NfsStorageModel) GetShare() NfsRepositoryShareSettingsModel {
	if o == nil {
		var ret NfsRepositoryShareSettingsModel
		return ret
	}

	return o.Share
}

// GetShareOk returns a tuple with the Share field value
// and a boolean to check if the value has been set.
func (o *NfsStorageModel) GetShareOk() (*NfsRepositoryShareSettingsModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Share, true
}

// SetShare sets field value
func (o *NfsStorageModel) SetShare(v NfsRepositoryShareSettingsModel) {
	o.Share = v
}

// GetRepository returns the Repository field value
func (o *NfsStorageModel) GetRepository() NetworkRepositorySettingsModel {
	if o == nil {
		var ret NetworkRepositorySettingsModel
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *NfsStorageModel) GetRepositoryOk() (*NetworkRepositorySettingsModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *NfsStorageModel) SetRepository(v NetworkRepositorySettingsModel) {
	o.Repository = v
}

// GetMountServer returns the MountServer field value
func (o *NfsStorageModel) GetMountServer() MountServerSettingsModel {
	if o == nil {
		var ret MountServerSettingsModel
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *NfsStorageModel) GetMountServerOk() (*MountServerSettingsModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *NfsStorageModel) SetMountServer(v MountServerSettingsModel) {
	o.MountServer = v
}

func (o NfsStorageModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRepositoryModel, errRepositoryModel := json.Marshal(o.RepositoryModel)
	if errRepositoryModel != nil {
		return []byte{}, errRepositoryModel
	}
	errRepositoryModel = json.Unmarshal([]byte(serializedRepositoryModel), &toSerialize)
	if errRepositoryModel != nil {
		return []byte{}, errRepositoryModel
	}
	if true {
		toSerialize["share"] = o.Share
	}
	if true {
		toSerialize["repository"] = o.Repository
	}
	if true {
		toSerialize["mountServer"] = o.MountServer
	}
	return json.Marshal(toSerialize)
}

type NullableNfsStorageModel struct {
	value *NfsStorageModel
	isSet bool
}

func (v NullableNfsStorageModel) Get() *NfsStorageModel {
	return v.value
}

func (v *NullableNfsStorageModel) Set(val *NfsStorageModel) {
	v.value = val
	v.isSet = true
}

func (v NullableNfsStorageModel) IsSet() bool {
	return v.isSet
}

func (v *NullableNfsStorageModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfsStorageModel(val *NfsStorageModel) *NullableNfsStorageModel {
	return &NullableNfsStorageModel{value: val, isSet: true}
}

func (v NullableNfsStorageModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfsStorageModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


