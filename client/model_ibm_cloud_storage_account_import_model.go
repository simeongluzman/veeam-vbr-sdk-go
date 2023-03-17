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

// IBMCloudStorageAccountImportModel Account used to access the IBM Cloud storage.
type IBMCloudStorageAccountImportModel struct {
	// Endpoint address of the storage.
	ServicePoint string `json:"servicePoint"`
	// ID of a region where the storage is located.
	RegionId string `json:"regionId"`
	Credentials CloudCredentialsImportModel `json:"credentials"`
	ConnectionSettings ObjectStorageConnectionImportSpec `json:"connectionSettings"`
}

// NewIBMCloudStorageAccountImportModel instantiates a new IBMCloudStorageAccountImportModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIBMCloudStorageAccountImportModel(servicePoint string, regionId string, credentials CloudCredentialsImportModel, connectionSettings ObjectStorageConnectionImportSpec) *IBMCloudStorageAccountImportModel {
	this := IBMCloudStorageAccountImportModel{}
	this.ServicePoint = servicePoint
	this.RegionId = regionId
	this.Credentials = credentials
	this.ConnectionSettings = connectionSettings
	return &this
}

// NewIBMCloudStorageAccountImportModelWithDefaults instantiates a new IBMCloudStorageAccountImportModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIBMCloudStorageAccountImportModelWithDefaults() *IBMCloudStorageAccountImportModel {
	this := IBMCloudStorageAccountImportModel{}
	return &this
}

// GetServicePoint returns the ServicePoint field value
func (o *IBMCloudStorageAccountImportModel) GetServicePoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServicePoint
}

// GetServicePointOk returns a tuple with the ServicePoint field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageAccountImportModel) GetServicePointOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServicePoint, true
}

// SetServicePoint sets field value
func (o *IBMCloudStorageAccountImportModel) SetServicePoint(v string) {
	o.ServicePoint = v
}

// GetRegionId returns the RegionId field value
func (o *IBMCloudStorageAccountImportModel) GetRegionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageAccountImportModel) GetRegionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RegionId, true
}

// SetRegionId sets field value
func (o *IBMCloudStorageAccountImportModel) SetRegionId(v string) {
	o.RegionId = v
}

// GetCredentials returns the Credentials field value
func (o *IBMCloudStorageAccountImportModel) GetCredentials() CloudCredentialsImportModel {
	if o == nil {
		var ret CloudCredentialsImportModel
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageAccountImportModel) GetCredentialsOk() (*CloudCredentialsImportModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *IBMCloudStorageAccountImportModel) SetCredentials(v CloudCredentialsImportModel) {
	o.Credentials = v
}

// GetConnectionSettings returns the ConnectionSettings field value
func (o *IBMCloudStorageAccountImportModel) GetConnectionSettings() ObjectStorageConnectionImportSpec {
	if o == nil {
		var ret ObjectStorageConnectionImportSpec
		return ret
	}

	return o.ConnectionSettings
}

// GetConnectionSettingsOk returns a tuple with the ConnectionSettings field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageAccountImportModel) GetConnectionSettingsOk() (*ObjectStorageConnectionImportSpec, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConnectionSettings, true
}

// SetConnectionSettings sets field value
func (o *IBMCloudStorageAccountImportModel) SetConnectionSettings(v ObjectStorageConnectionImportSpec) {
	o.ConnectionSettings = v
}

func (o IBMCloudStorageAccountImportModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["servicePoint"] = o.ServicePoint
	}
	if true {
		toSerialize["regionId"] = o.RegionId
	}
	if true {
		toSerialize["credentials"] = o.Credentials
	}
	if true {
		toSerialize["connectionSettings"] = o.ConnectionSettings
	}
	return json.Marshal(toSerialize)
}

type NullableIBMCloudStorageAccountImportModel struct {
	value *IBMCloudStorageAccountImportModel
	isSet bool
}

func (v NullableIBMCloudStorageAccountImportModel) Get() *IBMCloudStorageAccountImportModel {
	return v.value
}

func (v *NullableIBMCloudStorageAccountImportModel) Set(val *IBMCloudStorageAccountImportModel) {
	v.value = val
	v.isSet = true
}

func (v NullableIBMCloudStorageAccountImportModel) IsSet() bool {
	return v.isSet
}

func (v *NullableIBMCloudStorageAccountImportModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIBMCloudStorageAccountImportModel(val *IBMCloudStorageAccountImportModel) *NullableIBMCloudStorageAccountImportModel {
	return &NullableIBMCloudStorageAccountImportModel{value: val, isSet: true}
}

func (v NullableIBMCloudStorageAccountImportModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIBMCloudStorageAccountImportModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


