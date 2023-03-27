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

// EntireViVMCustomizedRestoreSpecAllOf struct for EntireViVMCustomizedRestoreSpecAllOf
type EntireViVMCustomizedRestoreSpecAllOf struct {
	DestinationHost *VmwareObjectModel `json:"destinationHost,omitempty"`
	ResourcePool *VmwareObjectModel `json:"resourcePool,omitempty"`
	Datastore *RestoreTargetDatastoreSpec `json:"datastore,omitempty"`
	Folder *RestoreTargetFolderSpec `json:"folder,omitempty"`
	Network *RestoreTargetNetworkSpec `json:"network,omitempty"`
}

// NewEntireViVMCustomizedRestoreSpecAllOf instantiates a new EntireViVMCustomizedRestoreSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntireViVMCustomizedRestoreSpecAllOf() *EntireViVMCustomizedRestoreSpecAllOf {
	this := EntireViVMCustomizedRestoreSpecAllOf{}
	return &this
}

// NewEntireViVMCustomizedRestoreSpecAllOfWithDefaults instantiates a new EntireViVMCustomizedRestoreSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntireViVMCustomizedRestoreSpecAllOfWithDefaults() *EntireViVMCustomizedRestoreSpecAllOf {
	this := EntireViVMCustomizedRestoreSpecAllOf{}
	return &this
}

// GetDestinationHost returns the DestinationHost field value if set, zero value otherwise.
func (o *EntireViVMCustomizedRestoreSpecAllOf) GetDestinationHost() VmwareObjectModel {
	if o == nil || isNil(o.DestinationHost) {
		var ret VmwareObjectModel
		return ret
	}
	return *o.DestinationHost
}

// GetDestinationHostOk returns a tuple with the DestinationHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntireViVMCustomizedRestoreSpecAllOf) GetDestinationHostOk() (*VmwareObjectModel, bool) {
	if o == nil || isNil(o.DestinationHost) {
    return nil, false
	}
	return o.DestinationHost, true
}

// HasDestinationHost returns a boolean if a field has been set.
func (o *EntireViVMCustomizedRestoreSpecAllOf) HasDestinationHost() bool {
	if o != nil && !isNil(o.DestinationHost) {
		return true
	}

	return false
}

// SetDestinationHost gets a reference to the given VmwareObjectModel and assigns it to the DestinationHost field.
func (o *EntireViVMCustomizedRestoreSpecAllOf) SetDestinationHost(v VmwareObjectModel) {
	o.DestinationHost = &v
}

// GetResourcePool returns the ResourcePool field value if set, zero value otherwise.
func (o *EntireViVMCustomizedRestoreSpecAllOf) GetResourcePool() VmwareObjectModel {
	if o == nil || isNil(o.ResourcePool) {
		var ret VmwareObjectModel
		return ret
	}
	return *o.ResourcePool
}

// GetResourcePoolOk returns a tuple with the ResourcePool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntireViVMCustomizedRestoreSpecAllOf) GetResourcePoolOk() (*VmwareObjectModel, bool) {
	if o == nil || isNil(o.ResourcePool) {
    return nil, false
	}
	return o.ResourcePool, true
}

// HasResourcePool returns a boolean if a field has been set.
func (o *EntireViVMCustomizedRestoreSpecAllOf) HasResourcePool() bool {
	if o != nil && !isNil(o.ResourcePool) {
		return true
	}

	return false
}

// SetResourcePool gets a reference to the given VmwareObjectModel and assigns it to the ResourcePool field.
func (o *EntireViVMCustomizedRestoreSpecAllOf) SetResourcePool(v VmwareObjectModel) {
	o.ResourcePool = &v
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *EntireViVMCustomizedRestoreSpecAllOf) GetDatastore() RestoreTargetDatastoreSpec {
	if o == nil || isNil(o.Datastore) {
		var ret RestoreTargetDatastoreSpec
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntireViVMCustomizedRestoreSpecAllOf) GetDatastoreOk() (*RestoreTargetDatastoreSpec, bool) {
	if o == nil || isNil(o.Datastore) {
    return nil, false
	}
	return o.Datastore, true
}

// HasDatastore returns a boolean if a field has been set.
func (o *EntireViVMCustomizedRestoreSpecAllOf) HasDatastore() bool {
	if o != nil && !isNil(o.Datastore) {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given RestoreTargetDatastoreSpec and assigns it to the Datastore field.
func (o *EntireViVMCustomizedRestoreSpecAllOf) SetDatastore(v RestoreTargetDatastoreSpec) {
	o.Datastore = &v
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *EntireViVMCustomizedRestoreSpecAllOf) GetFolder() RestoreTargetFolderSpec {
	if o == nil || isNil(o.Folder) {
		var ret RestoreTargetFolderSpec
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntireViVMCustomizedRestoreSpecAllOf) GetFolderOk() (*RestoreTargetFolderSpec, bool) {
	if o == nil || isNil(o.Folder) {
    return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *EntireViVMCustomizedRestoreSpecAllOf) HasFolder() bool {
	if o != nil && !isNil(o.Folder) {
		return true
	}

	return false
}

// SetFolder gets a reference to the given RestoreTargetFolderSpec and assigns it to the Folder field.
func (o *EntireViVMCustomizedRestoreSpecAllOf) SetFolder(v RestoreTargetFolderSpec) {
	o.Folder = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *EntireViVMCustomizedRestoreSpecAllOf) GetNetwork() RestoreTargetNetworkSpec {
	if o == nil || isNil(o.Network) {
		var ret RestoreTargetNetworkSpec
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntireViVMCustomizedRestoreSpecAllOf) GetNetworkOk() (*RestoreTargetNetworkSpec, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *EntireViVMCustomizedRestoreSpecAllOf) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given RestoreTargetNetworkSpec and assigns it to the Network field.
func (o *EntireViVMCustomizedRestoreSpecAllOf) SetNetwork(v RestoreTargetNetworkSpec) {
	o.Network = &v
}

func (o EntireViVMCustomizedRestoreSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DestinationHost) {
		toSerialize["destinationHost"] = o.DestinationHost
	}
	if !isNil(o.ResourcePool) {
		toSerialize["resourcePool"] = o.ResourcePool
	}
	if !isNil(o.Datastore) {
		toSerialize["datastore"] = o.Datastore
	}
	if !isNil(o.Folder) {
		toSerialize["folder"] = o.Folder
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableEntireViVMCustomizedRestoreSpecAllOf struct {
	value *EntireViVMCustomizedRestoreSpecAllOf
	isSet bool
}

func (v NullableEntireViVMCustomizedRestoreSpecAllOf) Get() *EntireViVMCustomizedRestoreSpecAllOf {
	return v.value
}

func (v *NullableEntireViVMCustomizedRestoreSpecAllOf) Set(val *EntireViVMCustomizedRestoreSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEntireViVMCustomizedRestoreSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEntireViVMCustomizedRestoreSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntireViVMCustomizedRestoreSpecAllOf(val *EntireViVMCustomizedRestoreSpecAllOf) *NullableEntireViVMCustomizedRestoreSpecAllOf {
	return &NullableEntireViVMCustomizedRestoreSpecAllOf{value: val, isSet: true}
}

func (v NullableEntireViVMCustomizedRestoreSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntireViVMCustomizedRestoreSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

