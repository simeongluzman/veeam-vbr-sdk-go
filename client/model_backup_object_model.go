/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev2
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// BackupObjectModel struct for BackupObjectModel
type BackupObjectModel struct {
	// ID of the object.
	Id string `json:"id"`
	// Name of the object.
	Name *string `json:"name,omitempty"`
	// Type of the object.
	Type *string `json:"type,omitempty"`
	PlatformName EPlatformType `json:"platformName"`
	// Id of the platform where the object was created.
	PlatformId *string `json:"platformId,omitempty"`
	// Number of restore points.
	RestorePointsCount *int32 `json:"restorePointsCount,omitempty"`
}

// NewBackupObjectModel instantiates a new BackupObjectModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupObjectModel(id string, platformName EPlatformType, ) *BackupObjectModel {
	this := BackupObjectModel{}
	this.Id = id
	this.PlatformName = platformName
	return &this
}

// NewBackupObjectModelWithDefaults instantiates a new BackupObjectModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupObjectModelWithDefaults() *BackupObjectModel {
	this := BackupObjectModel{}
	return &this
}

// GetId returns the Id field value
func (o *BackupObjectModel) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BackupObjectModel) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BackupObjectModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BackupObjectModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupObjectModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BackupObjectModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BackupObjectModel) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BackupObjectModel) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupObjectModel) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BackupObjectModel) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BackupObjectModel) SetType(v string) {
	o.Type = &v
}

// GetPlatformName returns the PlatformName field value
func (o *BackupObjectModel) GetPlatformName() EPlatformType {
	if o == nil  {
		var ret EPlatformType
		return ret
	}

	return o.PlatformName
}

// GetPlatformNameOk returns a tuple with the PlatformName field value
// and a boolean to check if the value has been set.
func (o *BackupObjectModel) GetPlatformNameOk() (*EPlatformType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PlatformName, true
}

// SetPlatformName sets field value
func (o *BackupObjectModel) SetPlatformName(v EPlatformType) {
	o.PlatformName = v
}

// GetPlatformId returns the PlatformId field value if set, zero value otherwise.
func (o *BackupObjectModel) GetPlatformId() string {
	if o == nil || o.PlatformId == nil {
		var ret string
		return ret
	}
	return *o.PlatformId
}

// GetPlatformIdOk returns a tuple with the PlatformId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupObjectModel) GetPlatformIdOk() (*string, bool) {
	if o == nil || o.PlatformId == nil {
		return nil, false
	}
	return o.PlatformId, true
}

// HasPlatformId returns a boolean if a field has been set.
func (o *BackupObjectModel) HasPlatformId() bool {
	if o != nil && o.PlatformId != nil {
		return true
	}

	return false
}

// SetPlatformId gets a reference to the given string and assigns it to the PlatformId field.
func (o *BackupObjectModel) SetPlatformId(v string) {
	o.PlatformId = &v
}

// GetRestorePointsCount returns the RestorePointsCount field value if set, zero value otherwise.
func (o *BackupObjectModel) GetRestorePointsCount() int32 {
	if o == nil || o.RestorePointsCount == nil {
		var ret int32
		return ret
	}
	return *o.RestorePointsCount
}

// GetRestorePointsCountOk returns a tuple with the RestorePointsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupObjectModel) GetRestorePointsCountOk() (*int32, bool) {
	if o == nil || o.RestorePointsCount == nil {
		return nil, false
	}
	return o.RestorePointsCount, true
}

// HasRestorePointsCount returns a boolean if a field has been set.
func (o *BackupObjectModel) HasRestorePointsCount() bool {
	if o != nil && o.RestorePointsCount != nil {
		return true
	}

	return false
}

// SetRestorePointsCount gets a reference to the given int32 and assigns it to the RestorePointsCount field.
func (o *BackupObjectModel) SetRestorePointsCount(v int32) {
	o.RestorePointsCount = &v
}

func (o BackupObjectModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["platformName"] = o.PlatformName
	}
	if o.PlatformId != nil {
		toSerialize["platformId"] = o.PlatformId
	}
	if o.RestorePointsCount != nil {
		toSerialize["restorePointsCount"] = o.RestorePointsCount
	}
	return json.Marshal(toSerialize)
}

type NullableBackupObjectModel struct {
	value *BackupObjectModel
	isSet bool
}

func (v NullableBackupObjectModel) Get() *BackupObjectModel {
	return v.value
}

func (v *NullableBackupObjectModel) Set(val *BackupObjectModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupObjectModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupObjectModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupObjectModel(val *BackupObjectModel) *NullableBackupObjectModel {
	return &NullableBackupObjectModel{value: val, isSet: true}
}

func (v NullableBackupObjectModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupObjectModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

