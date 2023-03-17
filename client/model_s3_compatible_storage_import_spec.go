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

// S3CompatibleStorageImportSpec struct for S3CompatibleStorageImportSpec
type S3CompatibleStorageImportSpec struct {
	// Name of the object storage repository.
	Name string `json:"name"`
	// Description of the object storage repository.
	Description string `json:"description"`
	// Tag that identifies the object storage repository.
	Tag string `json:"tag"`
	Type ERepositoryType `json:"type"`
	// If *true*, the maximum number of concurrent tasks is limited.
	EnableTaskLimit *bool `json:"enableTaskLimit,omitempty"`
	// Maximum number of concurrent tasks.
	MaxTaskCount *int32 `json:"maxTaskCount,omitempty"`
	Account S3CompatibleStorageAccountImportModel `json:"account"`
	Bucket S3CompatibleStorageBucketModel `json:"bucket"`
	MountServer MountServerSettingsImportSpec `json:"mountServer"`
	ProxyAppliance *S3CompatibleProxyModel `json:"proxyAppliance,omitempty"`
}

// NewS3CompatibleStorageImportSpec instantiates a new S3CompatibleStorageImportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3CompatibleStorageImportSpec(name string, description string, tag string, type_ ERepositoryType, account S3CompatibleStorageAccountImportModel, bucket S3CompatibleStorageBucketModel, mountServer MountServerSettingsImportSpec) *S3CompatibleStorageImportSpec {
	this := S3CompatibleStorageImportSpec{}
	this.Name = name
	this.Description = description
	this.Tag = tag
	this.Type = type_
	this.Account = account
	this.Bucket = bucket
	this.MountServer = mountServer
	return &this
}

// NewS3CompatibleStorageImportSpecWithDefaults instantiates a new S3CompatibleStorageImportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3CompatibleStorageImportSpecWithDefaults() *S3CompatibleStorageImportSpec {
	this := S3CompatibleStorageImportSpec{}
	return &this
}

// GetName returns the Name field value
func (o *S3CompatibleStorageImportSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageImportSpec) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *S3CompatibleStorageImportSpec) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *S3CompatibleStorageImportSpec) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageImportSpec) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *S3CompatibleStorageImportSpec) SetDescription(v string) {
	o.Description = v
}

// GetTag returns the Tag field value
func (o *S3CompatibleStorageImportSpec) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageImportSpec) GetTagOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *S3CompatibleStorageImportSpec) SetTag(v string) {
	o.Tag = v
}

// GetType returns the Type field value
func (o *S3CompatibleStorageImportSpec) GetType() ERepositoryType {
	if o == nil {
		var ret ERepositoryType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageImportSpec) GetTypeOk() (*ERepositoryType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *S3CompatibleStorageImportSpec) SetType(v ERepositoryType) {
	o.Type = v
}

// GetEnableTaskLimit returns the EnableTaskLimit field value if set, zero value otherwise.
func (o *S3CompatibleStorageImportSpec) GetEnableTaskLimit() bool {
	if o == nil || isNil(o.EnableTaskLimit) {
		var ret bool
		return ret
	}
	return *o.EnableTaskLimit
}

// GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageImportSpec) GetEnableTaskLimitOk() (*bool, bool) {
	if o == nil || isNil(o.EnableTaskLimit) {
    return nil, false
	}
	return o.EnableTaskLimit, true
}

// HasEnableTaskLimit returns a boolean if a field has been set.
func (o *S3CompatibleStorageImportSpec) HasEnableTaskLimit() bool {
	if o != nil && !isNil(o.EnableTaskLimit) {
		return true
	}

	return false
}

// SetEnableTaskLimit gets a reference to the given bool and assigns it to the EnableTaskLimit field.
func (o *S3CompatibleStorageImportSpec) SetEnableTaskLimit(v bool) {
	o.EnableTaskLimit = &v
}

// GetMaxTaskCount returns the MaxTaskCount field value if set, zero value otherwise.
func (o *S3CompatibleStorageImportSpec) GetMaxTaskCount() int32 {
	if o == nil || isNil(o.MaxTaskCount) {
		var ret int32
		return ret
	}
	return *o.MaxTaskCount
}

// GetMaxTaskCountOk returns a tuple with the MaxTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageImportSpec) GetMaxTaskCountOk() (*int32, bool) {
	if o == nil || isNil(o.MaxTaskCount) {
    return nil, false
	}
	return o.MaxTaskCount, true
}

// HasMaxTaskCount returns a boolean if a field has been set.
func (o *S3CompatibleStorageImportSpec) HasMaxTaskCount() bool {
	if o != nil && !isNil(o.MaxTaskCount) {
		return true
	}

	return false
}

// SetMaxTaskCount gets a reference to the given int32 and assigns it to the MaxTaskCount field.
func (o *S3CompatibleStorageImportSpec) SetMaxTaskCount(v int32) {
	o.MaxTaskCount = &v
}

// GetAccount returns the Account field value
func (o *S3CompatibleStorageImportSpec) GetAccount() S3CompatibleStorageAccountImportModel {
	if o == nil {
		var ret S3CompatibleStorageAccountImportModel
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageImportSpec) GetAccountOk() (*S3CompatibleStorageAccountImportModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *S3CompatibleStorageImportSpec) SetAccount(v S3CompatibleStorageAccountImportModel) {
	o.Account = v
}

// GetBucket returns the Bucket field value
func (o *S3CompatibleStorageImportSpec) GetBucket() S3CompatibleStorageBucketModel {
	if o == nil {
		var ret S3CompatibleStorageBucketModel
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageImportSpec) GetBucketOk() (*S3CompatibleStorageBucketModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *S3CompatibleStorageImportSpec) SetBucket(v S3CompatibleStorageBucketModel) {
	o.Bucket = v
}

// GetMountServer returns the MountServer field value
func (o *S3CompatibleStorageImportSpec) GetMountServer() MountServerSettingsImportSpec {
	if o == nil {
		var ret MountServerSettingsImportSpec
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageImportSpec) GetMountServerOk() (*MountServerSettingsImportSpec, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *S3CompatibleStorageImportSpec) SetMountServer(v MountServerSettingsImportSpec) {
	o.MountServer = v
}

// GetProxyAppliance returns the ProxyAppliance field value if set, zero value otherwise.
func (o *S3CompatibleStorageImportSpec) GetProxyAppliance() S3CompatibleProxyModel {
	if o == nil || isNil(o.ProxyAppliance) {
		var ret S3CompatibleProxyModel
		return ret
	}
	return *o.ProxyAppliance
}

// GetProxyApplianceOk returns a tuple with the ProxyAppliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageImportSpec) GetProxyApplianceOk() (*S3CompatibleProxyModel, bool) {
	if o == nil || isNil(o.ProxyAppliance) {
    return nil, false
	}
	return o.ProxyAppliance, true
}

// HasProxyAppliance returns a boolean if a field has been set.
func (o *S3CompatibleStorageImportSpec) HasProxyAppliance() bool {
	if o != nil && !isNil(o.ProxyAppliance) {
		return true
	}

	return false
}

// SetProxyAppliance gets a reference to the given S3CompatibleProxyModel and assigns it to the ProxyAppliance field.
func (o *S3CompatibleStorageImportSpec) SetProxyAppliance(v S3CompatibleProxyModel) {
	o.ProxyAppliance = &v
}

func (o S3CompatibleStorageImportSpec) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.EnableTaskLimit) {
		toSerialize["enableTaskLimit"] = o.EnableTaskLimit
	}
	if !isNil(o.MaxTaskCount) {
		toSerialize["maxTaskCount"] = o.MaxTaskCount
	}
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["bucket"] = o.Bucket
	}
	if true {
		toSerialize["mountServer"] = o.MountServer
	}
	if !isNil(o.ProxyAppliance) {
		toSerialize["proxyAppliance"] = o.ProxyAppliance
	}
	return json.Marshal(toSerialize)
}

type NullableS3CompatibleStorageImportSpec struct {
	value *S3CompatibleStorageImportSpec
	isSet bool
}

func (v NullableS3CompatibleStorageImportSpec) Get() *S3CompatibleStorageImportSpec {
	return v.value
}

func (v *NullableS3CompatibleStorageImportSpec) Set(val *S3CompatibleStorageImportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableS3CompatibleStorageImportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableS3CompatibleStorageImportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3CompatibleStorageImportSpec(val *S3CompatibleStorageImportSpec) *NullableS3CompatibleStorageImportSpec {
	return &NullableS3CompatibleStorageImportSpec{value: val, isSet: true}
}

func (v NullableS3CompatibleStorageImportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3CompatibleStorageImportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


