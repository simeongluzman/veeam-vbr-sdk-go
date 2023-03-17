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

// NfsRepositoryShareSettingsModel Share settings.
type NfsRepositoryShareSettingsModel struct {
	// Path to the shared folder that is used as a backup repository.
	SharePath string `json:"sharePath"`
	GatewayServer *RepositoryShareGatewayModel `json:"gatewayServer,omitempty"`
}

// NewNfsRepositoryShareSettingsModel instantiates a new NfsRepositoryShareSettingsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfsRepositoryShareSettingsModel(sharePath string) *NfsRepositoryShareSettingsModel {
	this := NfsRepositoryShareSettingsModel{}
	this.SharePath = sharePath
	return &this
}

// NewNfsRepositoryShareSettingsModelWithDefaults instantiates a new NfsRepositoryShareSettingsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfsRepositoryShareSettingsModelWithDefaults() *NfsRepositoryShareSettingsModel {
	this := NfsRepositoryShareSettingsModel{}
	return &this
}

// GetSharePath returns the SharePath field value
func (o *NfsRepositoryShareSettingsModel) GetSharePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SharePath
}

// GetSharePathOk returns a tuple with the SharePath field value
// and a boolean to check if the value has been set.
func (o *NfsRepositoryShareSettingsModel) GetSharePathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SharePath, true
}

// SetSharePath sets field value
func (o *NfsRepositoryShareSettingsModel) SetSharePath(v string) {
	o.SharePath = v
}

// GetGatewayServer returns the GatewayServer field value if set, zero value otherwise.
func (o *NfsRepositoryShareSettingsModel) GetGatewayServer() RepositoryShareGatewayModel {
	if o == nil || isNil(o.GatewayServer) {
		var ret RepositoryShareGatewayModel
		return ret
	}
	return *o.GatewayServer
}

// GetGatewayServerOk returns a tuple with the GatewayServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsRepositoryShareSettingsModel) GetGatewayServerOk() (*RepositoryShareGatewayModel, bool) {
	if o == nil || isNil(o.GatewayServer) {
    return nil, false
	}
	return o.GatewayServer, true
}

// HasGatewayServer returns a boolean if a field has been set.
func (o *NfsRepositoryShareSettingsModel) HasGatewayServer() bool {
	if o != nil && !isNil(o.GatewayServer) {
		return true
	}

	return false
}

// SetGatewayServer gets a reference to the given RepositoryShareGatewayModel and assigns it to the GatewayServer field.
func (o *NfsRepositoryShareSettingsModel) SetGatewayServer(v RepositoryShareGatewayModel) {
	o.GatewayServer = &v
}

func (o NfsRepositoryShareSettingsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sharePath"] = o.SharePath
	}
	if !isNil(o.GatewayServer) {
		toSerialize["gatewayServer"] = o.GatewayServer
	}
	return json.Marshal(toSerialize)
}

type NullableNfsRepositoryShareSettingsModel struct {
	value *NfsRepositoryShareSettingsModel
	isSet bool
}

func (v NullableNfsRepositoryShareSettingsModel) Get() *NfsRepositoryShareSettingsModel {
	return v.value
}

func (v *NullableNfsRepositoryShareSettingsModel) Set(val *NfsRepositoryShareSettingsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableNfsRepositoryShareSettingsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableNfsRepositoryShareSettingsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfsRepositoryShareSettingsModel(val *NfsRepositoryShareSettingsModel) *NullableNfsRepositoryShareSettingsModel {
	return &NullableNfsRepositoryShareSettingsModel{value: val, isSet: true}
}

func (v NullableNfsRepositoryShareSettingsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfsRepositoryShareSettingsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


