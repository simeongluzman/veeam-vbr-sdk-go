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

// PreferredNetworkModel struct for PreferredNetworkModel
type PreferredNetworkModel struct {
	// IP address.
	IpAddress *string `json:"ipAddress,omitempty"`
	// Subnet mask.
	SubnetMask *string `json:"subnetMask,omitempty"`
	// CIDR notation.
	CidrNotation *string `json:"cidrNotation,omitempty"`
}

// NewPreferredNetworkModel instantiates a new PreferredNetworkModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreferredNetworkModel() *PreferredNetworkModel {
	this := PreferredNetworkModel{}
	return &this
}

// NewPreferredNetworkModelWithDefaults instantiates a new PreferredNetworkModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferredNetworkModelWithDefaults() *PreferredNetworkModel {
	this := PreferredNetworkModel{}
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *PreferredNetworkModel) GetIpAddress() string {
	if o == nil || isNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredNetworkModel) GetIpAddressOk() (*string, bool) {
	if o == nil || isNil(o.IpAddress) {
    return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *PreferredNetworkModel) HasIpAddress() bool {
	if o != nil && !isNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *PreferredNetworkModel) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetSubnetMask returns the SubnetMask field value if set, zero value otherwise.
func (o *PreferredNetworkModel) GetSubnetMask() string {
	if o == nil || isNil(o.SubnetMask) {
		var ret string
		return ret
	}
	return *o.SubnetMask
}

// GetSubnetMaskOk returns a tuple with the SubnetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredNetworkModel) GetSubnetMaskOk() (*string, bool) {
	if o == nil || isNil(o.SubnetMask) {
    return nil, false
	}
	return o.SubnetMask, true
}

// HasSubnetMask returns a boolean if a field has been set.
func (o *PreferredNetworkModel) HasSubnetMask() bool {
	if o != nil && !isNil(o.SubnetMask) {
		return true
	}

	return false
}

// SetSubnetMask gets a reference to the given string and assigns it to the SubnetMask field.
func (o *PreferredNetworkModel) SetSubnetMask(v string) {
	o.SubnetMask = &v
}

// GetCidrNotation returns the CidrNotation field value if set, zero value otherwise.
func (o *PreferredNetworkModel) GetCidrNotation() string {
	if o == nil || isNil(o.CidrNotation) {
		var ret string
		return ret
	}
	return *o.CidrNotation
}

// GetCidrNotationOk returns a tuple with the CidrNotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredNetworkModel) GetCidrNotationOk() (*string, bool) {
	if o == nil || isNil(o.CidrNotation) {
    return nil, false
	}
	return o.CidrNotation, true
}

// HasCidrNotation returns a boolean if a field has been set.
func (o *PreferredNetworkModel) HasCidrNotation() bool {
	if o != nil && !isNil(o.CidrNotation) {
		return true
	}

	return false
}

// SetCidrNotation gets a reference to the given string and assigns it to the CidrNotation field.
func (o *PreferredNetworkModel) SetCidrNotation(v string) {
	o.CidrNotation = &v
}

func (o PreferredNetworkModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !isNil(o.SubnetMask) {
		toSerialize["subnetMask"] = o.SubnetMask
	}
	if !isNil(o.CidrNotation) {
		toSerialize["cidrNotation"] = o.CidrNotation
	}
	return json.Marshal(toSerialize)
}

type NullablePreferredNetworkModel struct {
	value *PreferredNetworkModel
	isSet bool
}

func (v NullablePreferredNetworkModel) Get() *PreferredNetworkModel {
	return v.value
}

func (v *NullablePreferredNetworkModel) Set(val *PreferredNetworkModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePreferredNetworkModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePreferredNetworkModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreferredNetworkModel(val *PreferredNetworkModel) *NullablePreferredNetworkModel {
	return &NullablePreferredNetworkModel{value: val, isSet: true}
}

func (v NullablePreferredNetworkModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreferredNetworkModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


