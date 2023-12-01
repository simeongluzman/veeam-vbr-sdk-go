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

// ProxyServerSettingsImportSpec Settings of the server that is used as a backup proxy.
type ProxyServerSettingsImportSpec struct {
	// Name of the server.
	HostName string `json:"hostName"`
	// Tag assigned to the server.
	HostTag *string `json:"hostTag,omitempty"`
	TransportMode *EBackupProxyTransportMode `json:"transportMode,omitempty"`
	// (For the Direct storage access and Virtual appliance transport modes) If *true*, Veeam Backup & Replication failovers to the network transport mode in case the primary mode fails or is unavailable.
	FailoverToNetwork *bool `json:"failoverToNetwork,omitempty"`
	// (For the Network mode) If *true*, VM data is transferred over an encrypted TLS connection.
	HostToProxyEncryption *bool `json:"hostToProxyEncryption,omitempty"`
	ConnectedDatastores *ProxyDatastoreSettingsModel `json:"connectedDatastores,omitempty"`
	// Maximum number of concurrent tasks.
	MaxTaskCount *int32 `json:"maxTaskCount,omitempty"`
}

// NewProxyServerSettingsImportSpec instantiates a new ProxyServerSettingsImportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyServerSettingsImportSpec(hostName string) *ProxyServerSettingsImportSpec {
	this := ProxyServerSettingsImportSpec{}
	this.HostName = hostName
	return &this
}

// NewProxyServerSettingsImportSpecWithDefaults instantiates a new ProxyServerSettingsImportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyServerSettingsImportSpecWithDefaults() *ProxyServerSettingsImportSpec {
	this := ProxyServerSettingsImportSpec{}
	return &this
}

// GetHostName returns the HostName field value
func (o *ProxyServerSettingsImportSpec) GetHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value
// and a boolean to check if the value has been set.
func (o *ProxyServerSettingsImportSpec) GetHostNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HostName, true
}

// SetHostName sets field value
func (o *ProxyServerSettingsImportSpec) SetHostName(v string) {
	o.HostName = v
}

// GetHostTag returns the HostTag field value if set, zero value otherwise.
func (o *ProxyServerSettingsImportSpec) GetHostTag() string {
	if o == nil || isNil(o.HostTag) {
		var ret string
		return ret
	}
	return *o.HostTag
}

// GetHostTagOk returns a tuple with the HostTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyServerSettingsImportSpec) GetHostTagOk() (*string, bool) {
	if o == nil || isNil(o.HostTag) {
    return nil, false
	}
	return o.HostTag, true
}

// HasHostTag returns a boolean if a field has been set.
func (o *ProxyServerSettingsImportSpec) HasHostTag() bool {
	if o != nil && !isNil(o.HostTag) {
		return true
	}

	return false
}

// SetHostTag gets a reference to the given string and assigns it to the HostTag field.
func (o *ProxyServerSettingsImportSpec) SetHostTag(v string) {
	o.HostTag = &v
}

// GetTransportMode returns the TransportMode field value if set, zero value otherwise.
func (o *ProxyServerSettingsImportSpec) GetTransportMode() EBackupProxyTransportMode {
	if o == nil || isNil(o.TransportMode) {
		var ret EBackupProxyTransportMode
		return ret
	}
	return *o.TransportMode
}

// GetTransportModeOk returns a tuple with the TransportMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyServerSettingsImportSpec) GetTransportModeOk() (*EBackupProxyTransportMode, bool) {
	if o == nil || isNil(o.TransportMode) {
    return nil, false
	}
	return o.TransportMode, true
}

// HasTransportMode returns a boolean if a field has been set.
func (o *ProxyServerSettingsImportSpec) HasTransportMode() bool {
	if o != nil && !isNil(o.TransportMode) {
		return true
	}

	return false
}

// SetTransportMode gets a reference to the given EBackupProxyTransportMode and assigns it to the TransportMode field.
func (o *ProxyServerSettingsImportSpec) SetTransportMode(v EBackupProxyTransportMode) {
	o.TransportMode = &v
}

// GetFailoverToNetwork returns the FailoverToNetwork field value if set, zero value otherwise.
func (o *ProxyServerSettingsImportSpec) GetFailoverToNetwork() bool {
	if o == nil || isNil(o.FailoverToNetwork) {
		var ret bool
		return ret
	}
	return *o.FailoverToNetwork
}

// GetFailoverToNetworkOk returns a tuple with the FailoverToNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyServerSettingsImportSpec) GetFailoverToNetworkOk() (*bool, bool) {
	if o == nil || isNil(o.FailoverToNetwork) {
    return nil, false
	}
	return o.FailoverToNetwork, true
}

// HasFailoverToNetwork returns a boolean if a field has been set.
func (o *ProxyServerSettingsImportSpec) HasFailoverToNetwork() bool {
	if o != nil && !isNil(o.FailoverToNetwork) {
		return true
	}

	return false
}

// SetFailoverToNetwork gets a reference to the given bool and assigns it to the FailoverToNetwork field.
func (o *ProxyServerSettingsImportSpec) SetFailoverToNetwork(v bool) {
	o.FailoverToNetwork = &v
}

// GetHostToProxyEncryption returns the HostToProxyEncryption field value if set, zero value otherwise.
func (o *ProxyServerSettingsImportSpec) GetHostToProxyEncryption() bool {
	if o == nil || isNil(o.HostToProxyEncryption) {
		var ret bool
		return ret
	}
	return *o.HostToProxyEncryption
}

// GetHostToProxyEncryptionOk returns a tuple with the HostToProxyEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyServerSettingsImportSpec) GetHostToProxyEncryptionOk() (*bool, bool) {
	if o == nil || isNil(o.HostToProxyEncryption) {
    return nil, false
	}
	return o.HostToProxyEncryption, true
}

// HasHostToProxyEncryption returns a boolean if a field has been set.
func (o *ProxyServerSettingsImportSpec) HasHostToProxyEncryption() bool {
	if o != nil && !isNil(o.HostToProxyEncryption) {
		return true
	}

	return false
}

// SetHostToProxyEncryption gets a reference to the given bool and assigns it to the HostToProxyEncryption field.
func (o *ProxyServerSettingsImportSpec) SetHostToProxyEncryption(v bool) {
	o.HostToProxyEncryption = &v
}

// GetConnectedDatastores returns the ConnectedDatastores field value if set, zero value otherwise.
func (o *ProxyServerSettingsImportSpec) GetConnectedDatastores() ProxyDatastoreSettingsModel {
	if o == nil || isNil(o.ConnectedDatastores) {
		var ret ProxyDatastoreSettingsModel
		return ret
	}
	return *o.ConnectedDatastores
}

// GetConnectedDatastoresOk returns a tuple with the ConnectedDatastores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyServerSettingsImportSpec) GetConnectedDatastoresOk() (*ProxyDatastoreSettingsModel, bool) {
	if o == nil || isNil(o.ConnectedDatastores) {
    return nil, false
	}
	return o.ConnectedDatastores, true
}

// HasConnectedDatastores returns a boolean if a field has been set.
func (o *ProxyServerSettingsImportSpec) HasConnectedDatastores() bool {
	if o != nil && !isNil(o.ConnectedDatastores) {
		return true
	}

	return false
}

// SetConnectedDatastores gets a reference to the given ProxyDatastoreSettingsModel and assigns it to the ConnectedDatastores field.
func (o *ProxyServerSettingsImportSpec) SetConnectedDatastores(v ProxyDatastoreSettingsModel) {
	o.ConnectedDatastores = &v
}

// GetMaxTaskCount returns the MaxTaskCount field value if set, zero value otherwise.
func (o *ProxyServerSettingsImportSpec) GetMaxTaskCount() int32 {
	if o == nil || isNil(o.MaxTaskCount) {
		var ret int32
		return ret
	}
	return *o.MaxTaskCount
}

// GetMaxTaskCountOk returns a tuple with the MaxTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyServerSettingsImportSpec) GetMaxTaskCountOk() (*int32, bool) {
	if o == nil || isNil(o.MaxTaskCount) {
    return nil, false
	}
	return o.MaxTaskCount, true
}

// HasMaxTaskCount returns a boolean if a field has been set.
func (o *ProxyServerSettingsImportSpec) HasMaxTaskCount() bool {
	if o != nil && !isNil(o.MaxTaskCount) {
		return true
	}

	return false
}

// SetMaxTaskCount gets a reference to the given int32 and assigns it to the MaxTaskCount field.
func (o *ProxyServerSettingsImportSpec) SetMaxTaskCount(v int32) {
	o.MaxTaskCount = &v
}

func (o ProxyServerSettingsImportSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hostName"] = o.HostName
	}
	if !isNil(o.HostTag) {
		toSerialize["hostTag"] = o.HostTag
	}
	if !isNil(o.TransportMode) {
		toSerialize["transportMode"] = o.TransportMode
	}
	if !isNil(o.FailoverToNetwork) {
		toSerialize["failoverToNetwork"] = o.FailoverToNetwork
	}
	if !isNil(o.HostToProxyEncryption) {
		toSerialize["hostToProxyEncryption"] = o.HostToProxyEncryption
	}
	if !isNil(o.ConnectedDatastores) {
		toSerialize["connectedDatastores"] = o.ConnectedDatastores
	}
	if !isNil(o.MaxTaskCount) {
		toSerialize["maxTaskCount"] = o.MaxTaskCount
	}
	return json.Marshal(toSerialize)
}

type NullableProxyServerSettingsImportSpec struct {
	value *ProxyServerSettingsImportSpec
	isSet bool
}

func (v NullableProxyServerSettingsImportSpec) Get() *ProxyServerSettingsImportSpec {
	return v.value
}

func (v *NullableProxyServerSettingsImportSpec) Set(val *ProxyServerSettingsImportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyServerSettingsImportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyServerSettingsImportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyServerSettingsImportSpec(val *ProxyServerSettingsImportSpec) *NullableProxyServerSettingsImportSpec {
	return &NullableProxyServerSettingsImportSpec{value: val, isSet: true}
}

func (v NullableProxyServerSettingsImportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyServerSettingsImportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

