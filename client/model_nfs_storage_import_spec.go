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

// NfsStorageImportSpec struct for NfsStorageImportSpec
type NfsStorageImportSpec struct {
	// Name of the backup repository.
	Name string `json:"name"`
	// Description of the backup repository.
	Description string `json:"description"`
	// Tag that identifies the backup repository.
	Tag string `json:"tag"`
	Type ERepositoryType `json:"type"`
	Share NfsRepositoryShareSettingsSpec `json:"share"`
	Repository NetworkRepositorySettingsModel `json:"repository"`
	MountServer MountServerSettingsImportSpec `json:"mountServer"`
}

// NewNfsStorageImportSpec instantiates a new NfsStorageImportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfsStorageImportSpec(name string, description string, tag string, type_ ERepositoryType, share NfsRepositoryShareSettingsSpec, repository NetworkRepositorySettingsModel, mountServer MountServerSettingsImportSpec) *NfsStorageImportSpec {
	this := NfsStorageImportSpec{}
	this.Name = name
	this.Description = description
	this.Tag = tag
	this.Type = type_
	this.Share = share
	this.Repository = repository
	this.MountServer = mountServer
	return &this
}

// NewNfsStorageImportSpecWithDefaults instantiates a new NfsStorageImportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfsStorageImportSpecWithDefaults() *NfsStorageImportSpec {
	this := NfsStorageImportSpec{}
	return &this
}

// GetName returns the Name field value
func (o *NfsStorageImportSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NfsStorageImportSpec) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NfsStorageImportSpec) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *NfsStorageImportSpec) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *NfsStorageImportSpec) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *NfsStorageImportSpec) SetDescription(v string) {
	o.Description = v
}

// GetTag returns the Tag field value
func (o *NfsStorageImportSpec) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *NfsStorageImportSpec) GetTagOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *NfsStorageImportSpec) SetTag(v string) {
	o.Tag = v
}

// GetType returns the Type field value
func (o *NfsStorageImportSpec) GetType() ERepositoryType {
	if o == nil {
		var ret ERepositoryType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NfsStorageImportSpec) GetTypeOk() (*ERepositoryType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NfsStorageImportSpec) SetType(v ERepositoryType) {
	o.Type = v
}

// GetShare returns the Share field value
func (o *NfsStorageImportSpec) GetShare() NfsRepositoryShareSettingsSpec {
	if o == nil {
		var ret NfsRepositoryShareSettingsSpec
		return ret
	}

	return o.Share
}

// GetShareOk returns a tuple with the Share field value
// and a boolean to check if the value has been set.
func (o *NfsStorageImportSpec) GetShareOk() (*NfsRepositoryShareSettingsSpec, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Share, true
}

// SetShare sets field value
func (o *NfsStorageImportSpec) SetShare(v NfsRepositoryShareSettingsSpec) {
	o.Share = v
}

// GetRepository returns the Repository field value
func (o *NfsStorageImportSpec) GetRepository() NetworkRepositorySettingsModel {
	if o == nil {
		var ret NetworkRepositorySettingsModel
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *NfsStorageImportSpec) GetRepositoryOk() (*NetworkRepositorySettingsModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *NfsStorageImportSpec) SetRepository(v NetworkRepositorySettingsModel) {
	o.Repository = v
}

// GetMountServer returns the MountServer field value
func (o *NfsStorageImportSpec) GetMountServer() MountServerSettingsImportSpec {
	if o == nil {
		var ret MountServerSettingsImportSpec
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *NfsStorageImportSpec) GetMountServerOk() (*MountServerSettingsImportSpec, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *NfsStorageImportSpec) SetMountServer(v MountServerSettingsImportSpec) {
	o.MountServer = v
}

func (o NfsStorageImportSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["tag"] = o.Tag
	}
	if true {
		toSerialize["type"] = o.Type
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

type NullableNfsStorageImportSpec struct {
	value *NfsStorageImportSpec
	isSet bool
}

func (v NullableNfsStorageImportSpec) Get() *NfsStorageImportSpec {
	return v.value
}

func (v *NullableNfsStorageImportSpec) Set(val *NfsStorageImportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableNfsStorageImportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableNfsStorageImportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfsStorageImportSpec(val *NfsStorageImportSpec) *NullableNfsStorageImportSpec {
	return &NullableNfsStorageImportSpec{value: val, isSet: true}
}

func (v NullableNfsStorageImportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfsStorageImportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


