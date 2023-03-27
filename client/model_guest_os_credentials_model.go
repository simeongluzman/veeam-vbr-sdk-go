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

// GuestOsCredentialsModel VM custom credentials.
type GuestOsCredentialsModel struct {
	// Credentials ID for Microsoft Windows VMs.
	CredsId string `json:"credsId"`
	CredsType ECredentialsType `json:"credsType"`
	// Individual credentials for VMs.
	CredentialsPerMachine []GuestOsCredentialsPerMachineModel `json:"credentialsPerMachine,omitempty"`
}

// NewGuestOsCredentialsModel instantiates a new GuestOsCredentialsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuestOsCredentialsModel(credsId string, credsType ECredentialsType) *GuestOsCredentialsModel {
	this := GuestOsCredentialsModel{}
	this.CredsId = credsId
	this.CredsType = credsType
	return &this
}

// NewGuestOsCredentialsModelWithDefaults instantiates a new GuestOsCredentialsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuestOsCredentialsModelWithDefaults() *GuestOsCredentialsModel {
	this := GuestOsCredentialsModel{}
	return &this
}

// GetCredsId returns the CredsId field value
func (o *GuestOsCredentialsModel) GetCredsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CredsId
}

// GetCredsIdOk returns a tuple with the CredsId field value
// and a boolean to check if the value has been set.
func (o *GuestOsCredentialsModel) GetCredsIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CredsId, true
}

// SetCredsId sets field value
func (o *GuestOsCredentialsModel) SetCredsId(v string) {
	o.CredsId = v
}

// GetCredsType returns the CredsType field value
func (o *GuestOsCredentialsModel) GetCredsType() ECredentialsType {
	if o == nil {
		var ret ECredentialsType
		return ret
	}

	return o.CredsType
}

// GetCredsTypeOk returns a tuple with the CredsType field value
// and a boolean to check if the value has been set.
func (o *GuestOsCredentialsModel) GetCredsTypeOk() (*ECredentialsType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CredsType, true
}

// SetCredsType sets field value
func (o *GuestOsCredentialsModel) SetCredsType(v ECredentialsType) {
	o.CredsType = v
}

// GetCredentialsPerMachine returns the CredentialsPerMachine field value if set, zero value otherwise.
func (o *GuestOsCredentialsModel) GetCredentialsPerMachine() []GuestOsCredentialsPerMachineModel {
	if o == nil || isNil(o.CredentialsPerMachine) {
		var ret []GuestOsCredentialsPerMachineModel
		return ret
	}
	return o.CredentialsPerMachine
}

// GetCredentialsPerMachineOk returns a tuple with the CredentialsPerMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestOsCredentialsModel) GetCredentialsPerMachineOk() ([]GuestOsCredentialsPerMachineModel, bool) {
	if o == nil || isNil(o.CredentialsPerMachine) {
    return nil, false
	}
	return o.CredentialsPerMachine, true
}

// HasCredentialsPerMachine returns a boolean if a field has been set.
func (o *GuestOsCredentialsModel) HasCredentialsPerMachine() bool {
	if o != nil && !isNil(o.CredentialsPerMachine) {
		return true
	}

	return false
}

// SetCredentialsPerMachine gets a reference to the given []GuestOsCredentialsPerMachineModel and assigns it to the CredentialsPerMachine field.
func (o *GuestOsCredentialsModel) SetCredentialsPerMachine(v []GuestOsCredentialsPerMachineModel) {
	o.CredentialsPerMachine = v
}

func (o GuestOsCredentialsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["credsId"] = o.CredsId
	}
	if true {
		toSerialize["credsType"] = o.CredsType
	}
	if !isNil(o.CredentialsPerMachine) {
		toSerialize["credentialsPerMachine"] = o.CredentialsPerMachine
	}
	return json.Marshal(toSerialize)
}

type NullableGuestOsCredentialsModel struct {
	value *GuestOsCredentialsModel
	isSet bool
}

func (v NullableGuestOsCredentialsModel) Get() *GuestOsCredentialsModel {
	return v.value
}

func (v *NullableGuestOsCredentialsModel) Set(val *GuestOsCredentialsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGuestOsCredentialsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGuestOsCredentialsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuestOsCredentialsModel(val *GuestOsCredentialsModel) *NullableGuestOsCredentialsModel {
	return &NullableGuestOsCredentialsModel{value: val, isSet: true}
}

func (v NullableGuestOsCredentialsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuestOsCredentialsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


