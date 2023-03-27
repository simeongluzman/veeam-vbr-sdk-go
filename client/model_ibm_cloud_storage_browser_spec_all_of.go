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

// IBMCloudStorageBrowserSpecAllOf struct for IBMCloudStorageBrowserSpecAllOf
type IBMCloudStorageBrowserSpecAllOf struct {
	// Region of the IBM Cloud object storage.
	RegionId string `json:"regionId"`
	// Endpoint address and port number of the IBM Cloud object storage.
	ServicePoint string `json:"servicePoint"`
	// ID of a gateway server you want to use to connect to the object storage. Specify this parameter to check internet connection of the server. As a gateway server you can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used.
	GatewayServerId *string `json:"gatewayServerId,omitempty"`
}

// NewIBMCloudStorageBrowserSpecAllOf instantiates a new IBMCloudStorageBrowserSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIBMCloudStorageBrowserSpecAllOf(regionId string, servicePoint string) *IBMCloudStorageBrowserSpecAllOf {
	this := IBMCloudStorageBrowserSpecAllOf{}
	this.RegionId = regionId
	this.ServicePoint = servicePoint
	return &this
}

// NewIBMCloudStorageBrowserSpecAllOfWithDefaults instantiates a new IBMCloudStorageBrowserSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIBMCloudStorageBrowserSpecAllOfWithDefaults() *IBMCloudStorageBrowserSpecAllOf {
	this := IBMCloudStorageBrowserSpecAllOf{}
	return &this
}

// GetRegionId returns the RegionId field value
func (o *IBMCloudStorageBrowserSpecAllOf) GetRegionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageBrowserSpecAllOf) GetRegionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RegionId, true
}

// SetRegionId sets field value
func (o *IBMCloudStorageBrowserSpecAllOf) SetRegionId(v string) {
	o.RegionId = v
}

// GetServicePoint returns the ServicePoint field value
func (o *IBMCloudStorageBrowserSpecAllOf) GetServicePoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServicePoint
}

// GetServicePointOk returns a tuple with the ServicePoint field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageBrowserSpecAllOf) GetServicePointOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServicePoint, true
}

// SetServicePoint sets field value
func (o *IBMCloudStorageBrowserSpecAllOf) SetServicePoint(v string) {
	o.ServicePoint = v
}

// GetGatewayServerId returns the GatewayServerId field value if set, zero value otherwise.
func (o *IBMCloudStorageBrowserSpecAllOf) GetGatewayServerId() string {
	if o == nil || isNil(o.GatewayServerId) {
		var ret string
		return ret
	}
	return *o.GatewayServerId
}

// GetGatewayServerIdOk returns a tuple with the GatewayServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageBrowserSpecAllOf) GetGatewayServerIdOk() (*string, bool) {
	if o == nil || isNil(o.GatewayServerId) {
    return nil, false
	}
	return o.GatewayServerId, true
}

// HasGatewayServerId returns a boolean if a field has been set.
func (o *IBMCloudStorageBrowserSpecAllOf) HasGatewayServerId() bool {
	if o != nil && !isNil(o.GatewayServerId) {
		return true
	}

	return false
}

// SetGatewayServerId gets a reference to the given string and assigns it to the GatewayServerId field.
func (o *IBMCloudStorageBrowserSpecAllOf) SetGatewayServerId(v string) {
	o.GatewayServerId = &v
}

func (o IBMCloudStorageBrowserSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["regionId"] = o.RegionId
	}
	if true {
		toSerialize["servicePoint"] = o.ServicePoint
	}
	if !isNil(o.GatewayServerId) {
		toSerialize["gatewayServerId"] = o.GatewayServerId
	}
	return json.Marshal(toSerialize)
}

type NullableIBMCloudStorageBrowserSpecAllOf struct {
	value *IBMCloudStorageBrowserSpecAllOf
	isSet bool
}

func (v NullableIBMCloudStorageBrowserSpecAllOf) Get() *IBMCloudStorageBrowserSpecAllOf {
	return v.value
}

func (v *NullableIBMCloudStorageBrowserSpecAllOf) Set(val *IBMCloudStorageBrowserSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIBMCloudStorageBrowserSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIBMCloudStorageBrowserSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIBMCloudStorageBrowserSpecAllOf(val *IBMCloudStorageBrowserSpecAllOf) *NullableIBMCloudStorageBrowserSpecAllOf {
	return &NullableIBMCloudStorageBrowserSpecAllOf{value: val, isSet: true}
}

func (v NullableIBMCloudStorageBrowserSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIBMCloudStorageBrowserSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


