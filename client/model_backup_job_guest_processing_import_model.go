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

// BackupJobGuestProcessingImportModel Guest processing settings.
type BackupJobGuestProcessingImportModel struct {
	ApplicationAwareProcessing BackupApplicationAwareProcessingImportModel `json:"applicationAwareProcessing"`
	GuestFileSystemIndexing GuestFileSystemIndexingModel `json:"guestFileSystemIndexing"`
	GuestInteractionProxies *GuestInteractionProxiesSettingsImportModel `json:"guestInteractionProxies,omitempty"`
	GuestCredentials *GuestOsCredentialsImportModel `json:"guestCredentials,omitempty"`
}

// NewBackupJobGuestProcessingImportModel instantiates a new BackupJobGuestProcessingImportModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupJobGuestProcessingImportModel(applicationAwareProcessing BackupApplicationAwareProcessingImportModel, guestFileSystemIndexing GuestFileSystemIndexingModel) *BackupJobGuestProcessingImportModel {
	this := BackupJobGuestProcessingImportModel{}
	this.ApplicationAwareProcessing = applicationAwareProcessing
	this.GuestFileSystemIndexing = guestFileSystemIndexing
	return &this
}

// NewBackupJobGuestProcessingImportModelWithDefaults instantiates a new BackupJobGuestProcessingImportModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupJobGuestProcessingImportModelWithDefaults() *BackupJobGuestProcessingImportModel {
	this := BackupJobGuestProcessingImportModel{}
	return &this
}

// GetApplicationAwareProcessing returns the ApplicationAwareProcessing field value
func (o *BackupJobGuestProcessingImportModel) GetApplicationAwareProcessing() BackupApplicationAwareProcessingImportModel {
	if o == nil {
		var ret BackupApplicationAwareProcessingImportModel
		return ret
	}

	return o.ApplicationAwareProcessing
}

// GetApplicationAwareProcessingOk returns a tuple with the ApplicationAwareProcessing field value
// and a boolean to check if the value has been set.
func (o *BackupJobGuestProcessingImportModel) GetApplicationAwareProcessingOk() (*BackupApplicationAwareProcessingImportModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ApplicationAwareProcessing, true
}

// SetApplicationAwareProcessing sets field value
func (o *BackupJobGuestProcessingImportModel) SetApplicationAwareProcessing(v BackupApplicationAwareProcessingImportModel) {
	o.ApplicationAwareProcessing = v
}

// GetGuestFileSystemIndexing returns the GuestFileSystemIndexing field value
func (o *BackupJobGuestProcessingImportModel) GetGuestFileSystemIndexing() GuestFileSystemIndexingModel {
	if o == nil {
		var ret GuestFileSystemIndexingModel
		return ret
	}

	return o.GuestFileSystemIndexing
}

// GetGuestFileSystemIndexingOk returns a tuple with the GuestFileSystemIndexing field value
// and a boolean to check if the value has been set.
func (o *BackupJobGuestProcessingImportModel) GetGuestFileSystemIndexingOk() (*GuestFileSystemIndexingModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.GuestFileSystemIndexing, true
}

// SetGuestFileSystemIndexing sets field value
func (o *BackupJobGuestProcessingImportModel) SetGuestFileSystemIndexing(v GuestFileSystemIndexingModel) {
	o.GuestFileSystemIndexing = v
}

// GetGuestInteractionProxies returns the GuestInteractionProxies field value if set, zero value otherwise.
func (o *BackupJobGuestProcessingImportModel) GetGuestInteractionProxies() GuestInteractionProxiesSettingsImportModel {
	if o == nil || isNil(o.GuestInteractionProxies) {
		var ret GuestInteractionProxiesSettingsImportModel
		return ret
	}
	return *o.GuestInteractionProxies
}

// GetGuestInteractionProxiesOk returns a tuple with the GuestInteractionProxies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobGuestProcessingImportModel) GetGuestInteractionProxiesOk() (*GuestInteractionProxiesSettingsImportModel, bool) {
	if o == nil || isNil(o.GuestInteractionProxies) {
    return nil, false
	}
	return o.GuestInteractionProxies, true
}

// HasGuestInteractionProxies returns a boolean if a field has been set.
func (o *BackupJobGuestProcessingImportModel) HasGuestInteractionProxies() bool {
	if o != nil && !isNil(o.GuestInteractionProxies) {
		return true
	}

	return false
}

// SetGuestInteractionProxies gets a reference to the given GuestInteractionProxiesSettingsImportModel and assigns it to the GuestInteractionProxies field.
func (o *BackupJobGuestProcessingImportModel) SetGuestInteractionProxies(v GuestInteractionProxiesSettingsImportModel) {
	o.GuestInteractionProxies = &v
}

// GetGuestCredentials returns the GuestCredentials field value if set, zero value otherwise.
func (o *BackupJobGuestProcessingImportModel) GetGuestCredentials() GuestOsCredentialsImportModel {
	if o == nil || isNil(o.GuestCredentials) {
		var ret GuestOsCredentialsImportModel
		return ret
	}
	return *o.GuestCredentials
}

// GetGuestCredentialsOk returns a tuple with the GuestCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobGuestProcessingImportModel) GetGuestCredentialsOk() (*GuestOsCredentialsImportModel, bool) {
	if o == nil || isNil(o.GuestCredentials) {
    return nil, false
	}
	return o.GuestCredentials, true
}

// HasGuestCredentials returns a boolean if a field has been set.
func (o *BackupJobGuestProcessingImportModel) HasGuestCredentials() bool {
	if o != nil && !isNil(o.GuestCredentials) {
		return true
	}

	return false
}

// SetGuestCredentials gets a reference to the given GuestOsCredentialsImportModel and assigns it to the GuestCredentials field.
func (o *BackupJobGuestProcessingImportModel) SetGuestCredentials(v GuestOsCredentialsImportModel) {
	o.GuestCredentials = &v
}

func (o BackupJobGuestProcessingImportModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["applicationAwareProcessing"] = o.ApplicationAwareProcessing
	}
	if true {
		toSerialize["guestFileSystemIndexing"] = o.GuestFileSystemIndexing
	}
	if !isNil(o.GuestInteractionProxies) {
		toSerialize["guestInteractionProxies"] = o.GuestInteractionProxies
	}
	if !isNil(o.GuestCredentials) {
		toSerialize["guestCredentials"] = o.GuestCredentials
	}
	return json.Marshal(toSerialize)
}

type NullableBackupJobGuestProcessingImportModel struct {
	value *BackupJobGuestProcessingImportModel
	isSet bool
}

func (v NullableBackupJobGuestProcessingImportModel) Get() *BackupJobGuestProcessingImportModel {
	return v.value
}

func (v *NullableBackupJobGuestProcessingImportModel) Set(val *BackupJobGuestProcessingImportModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupJobGuestProcessingImportModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupJobGuestProcessingImportModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupJobGuestProcessingImportModel(val *BackupJobGuestProcessingImportModel) *NullableBackupJobGuestProcessingImportModel {
	return &NullableBackupJobGuestProcessingImportModel{value: val, isSet: true}
}

func (v NullableBackupJobGuestProcessingImportModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupJobGuestProcessingImportModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


