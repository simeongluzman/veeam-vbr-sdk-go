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

// WasabiCloudStorageImportSpec struct for WasabiCloudStorageImportSpec
type WasabiCloudStorageImportSpec struct {
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
	Account WasabiCloudStorageAccountImportModel `json:"account"`
	Bucket WasabiCloudStorageBucketModel `json:"bucket"`
	MountServer MountServerSettingsImportSpec `json:"mountServer"`
	ProxyAppliance *S3CompatibleProxyModel `json:"proxyAppliance,omitempty"`
}

// NewWasabiCloudStorageImportSpec instantiates a new WasabiCloudStorageImportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWasabiCloudStorageImportSpec(name string, description string, tag string, type_ ERepositoryType, account WasabiCloudStorageAccountImportModel, bucket WasabiCloudStorageBucketModel, mountServer MountServerSettingsImportSpec) *WasabiCloudStorageImportSpec {
	this := WasabiCloudStorageImportSpec{}
	this.Name = name
	this.Description = description
	this.Tag = tag
	this.Type = type_
	this.Account = account
	this.Bucket = bucket
	this.MountServer = mountServer
	return &this
}

// NewWasabiCloudStorageImportSpecWithDefaults instantiates a new WasabiCloudStorageImportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWasabiCloudStorageImportSpecWithDefaults() *WasabiCloudStorageImportSpec {
	this := WasabiCloudStorageImportSpec{}
	return &this
}

// GetName returns the Name field value
func (o *WasabiCloudStorageImportSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WasabiCloudStorageImportSpec) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WasabiCloudStorageImportSpec) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *WasabiCloudStorageImportSpec) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *WasabiCloudStorageImportSpec) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *WasabiCloudStorageImportSpec) SetDescription(v string) {
	o.Description = v
}

// GetTag returns the Tag field value
func (o *WasabiCloudStorageImportSpec) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *WasabiCloudStorageImportSpec) GetTagOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *WasabiCloudStorageImportSpec) SetTag(v string) {
	o.Tag = v
}

// GetType returns the Type field value
func (o *WasabiCloudStorageImportSpec) GetType() ERepositoryType {
	if o == nil {
		var ret ERepositoryType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WasabiCloudStorageImportSpec) GetTypeOk() (*ERepositoryType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WasabiCloudStorageImportSpec) SetType(v ERepositoryType) {
	o.Type = v
}

// GetEnableTaskLimit returns the EnableTaskLimit field value if set, zero value otherwise.
func (o *WasabiCloudStorageImportSpec) GetEnableTaskLimit() bool {
	if o == nil || isNil(o.EnableTaskLimit) {
		var ret bool
		return ret
	}
	return *o.EnableTaskLimit
}

// GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WasabiCloudStorageImportSpec) GetEnableTaskLimitOk() (*bool, bool) {
	if o == nil || isNil(o.EnableTaskLimit) {
    return nil, false
	}
	return o.EnableTaskLimit, true
}

// HasEnableTaskLimit returns a boolean if a field has been set.
func (o *WasabiCloudStorageImportSpec) HasEnableTaskLimit() bool {
	if o != nil && !isNil(o.EnableTaskLimit) {
		return true
	}

	return false
}

// SetEnableTaskLimit gets a reference to the given bool and assigns it to the EnableTaskLimit field.
func (o *WasabiCloudStorageImportSpec) SetEnableTaskLimit(v bool) {
	o.EnableTaskLimit = &v
}

// GetMaxTaskCount returns the MaxTaskCount field value if set, zero value otherwise.
func (o *WasabiCloudStorageImportSpec) GetMaxTaskCount() int32 {
	if o == nil || isNil(o.MaxTaskCount) {
		var ret int32
		return ret
	}
	return *o.MaxTaskCount
}

// GetMaxTaskCountOk returns a tuple with the MaxTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WasabiCloudStorageImportSpec) GetMaxTaskCountOk() (*int32, bool) {
	if o == nil || isNil(o.MaxTaskCount) {
    return nil, false
	}
	return o.MaxTaskCount, true
}

// HasMaxTaskCount returns a boolean if a field has been set.
func (o *WasabiCloudStorageImportSpec) HasMaxTaskCount() bool {
	if o != nil && !isNil(o.MaxTaskCount) {
		return true
	}

	return false
}

// SetMaxTaskCount gets a reference to the given int32 and assigns it to the MaxTaskCount field.
func (o *WasabiCloudStorageImportSpec) SetMaxTaskCount(v int32) {
	o.MaxTaskCount = &v
}

// GetAccount returns the Account field value
func (o *WasabiCloudStorageImportSpec) GetAccount() WasabiCloudStorageAccountImportModel {
	if o == nil {
		var ret WasabiCloudStorageAccountImportModel
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *WasabiCloudStorageImportSpec) GetAccountOk() (*WasabiCloudStorageAccountImportModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *WasabiCloudStorageImportSpec) SetAccount(v WasabiCloudStorageAccountImportModel) {
	o.Account = v
}

// GetBucket returns the Bucket field value
func (o *WasabiCloudStorageImportSpec) GetBucket() WasabiCloudStorageBucketModel {
	if o == nil {
		var ret WasabiCloudStorageBucketModel
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *WasabiCloudStorageImportSpec) GetBucketOk() (*WasabiCloudStorageBucketModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *WasabiCloudStorageImportSpec) SetBucket(v WasabiCloudStorageBucketModel) {
	o.Bucket = v
}

// GetMountServer returns the MountServer field value
func (o *WasabiCloudStorageImportSpec) GetMountServer() MountServerSettingsImportSpec {
	if o == nil {
		var ret MountServerSettingsImportSpec
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *WasabiCloudStorageImportSpec) GetMountServerOk() (*MountServerSettingsImportSpec, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *WasabiCloudStorageImportSpec) SetMountServer(v MountServerSettingsImportSpec) {
	o.MountServer = v
}

// GetProxyAppliance returns the ProxyAppliance field value if set, zero value otherwise.
func (o *WasabiCloudStorageImportSpec) GetProxyAppliance() S3CompatibleProxyModel {
	if o == nil || isNil(o.ProxyAppliance) {
		var ret S3CompatibleProxyModel
		return ret
	}
	return *o.ProxyAppliance
}

// GetProxyApplianceOk returns a tuple with the ProxyAppliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WasabiCloudStorageImportSpec) GetProxyApplianceOk() (*S3CompatibleProxyModel, bool) {
	if o == nil || isNil(o.ProxyAppliance) {
    return nil, false
	}
	return o.ProxyAppliance, true
}

// HasProxyAppliance returns a boolean if a field has been set.
func (o *WasabiCloudStorageImportSpec) HasProxyAppliance() bool {
	if o != nil && !isNil(o.ProxyAppliance) {
		return true
	}

	return false
}

// SetProxyAppliance gets a reference to the given S3CompatibleProxyModel and assigns it to the ProxyAppliance field.
func (o *WasabiCloudStorageImportSpec) SetProxyAppliance(v S3CompatibleProxyModel) {
	o.ProxyAppliance = &v
}

func (o WasabiCloudStorageImportSpec) MarshalJSON() ([]byte, error) {
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

type NullableWasabiCloudStorageImportSpec struct {
	value *WasabiCloudStorageImportSpec
	isSet bool
}

func (v NullableWasabiCloudStorageImportSpec) Get() *WasabiCloudStorageImportSpec {
	return v.value
}

func (v *NullableWasabiCloudStorageImportSpec) Set(val *WasabiCloudStorageImportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableWasabiCloudStorageImportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableWasabiCloudStorageImportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWasabiCloudStorageImportSpec(val *WasabiCloudStorageImportSpec) *NullableWasabiCloudStorageImportSpec {
	return &NullableWasabiCloudStorageImportSpec{value: val, isSet: true}
}

func (v NullableWasabiCloudStorageImportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWasabiCloudStorageImportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


