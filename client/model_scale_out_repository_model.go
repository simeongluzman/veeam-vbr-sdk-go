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

// ScaleOutRepositoryModel struct for ScaleOutRepositoryModel
type ScaleOutRepositoryModel struct {
	// ID of the scale-out backup repository.
	Id string `json:"id"`
	// Name of the scale-out backup repository.
	Name string `json:"name"`
	// Description of the scale-out backup repository.
	Description string `json:"description"`
	// Tag assigned to of the scale-out backup repository.
	Tag string `json:"tag"`
	PerformanceTier PerformanceTierModel `json:"performanceTier"`
	PlacementPolicy *PlacementPolicyModel `json:"placementPolicy,omitempty"`
	CapacityTier *CapacityTierModel `json:"capacityTier,omitempty"`
	ArchiveTier *ArchiveTierModel `json:"archiveTier,omitempty"`
}

// NewScaleOutRepositoryModel instantiates a new ScaleOutRepositoryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScaleOutRepositoryModel(id string, name string, description string, tag string, performanceTier PerformanceTierModel, ) *ScaleOutRepositoryModel {
	this := ScaleOutRepositoryModel{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Tag = tag
	this.PerformanceTier = performanceTier
	return &this
}

// NewScaleOutRepositoryModelWithDefaults instantiates a new ScaleOutRepositoryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScaleOutRepositoryModelWithDefaults() *ScaleOutRepositoryModel {
	this := ScaleOutRepositoryModel{}
	return &this
}

// GetId returns the Id field value
func (o *ScaleOutRepositoryModel) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScaleOutRepositoryModel) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScaleOutRepositoryModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ScaleOutRepositoryModel) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScaleOutRepositoryModel) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScaleOutRepositoryModel) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ScaleOutRepositoryModel) GetDescription() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ScaleOutRepositoryModel) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ScaleOutRepositoryModel) SetDescription(v string) {
	o.Description = v
}

// GetTag returns the Tag field value
func (o *ScaleOutRepositoryModel) GetTag() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *ScaleOutRepositoryModel) GetTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *ScaleOutRepositoryModel) SetTag(v string) {
	o.Tag = v
}

// GetPerformanceTier returns the PerformanceTier field value
func (o *ScaleOutRepositoryModel) GetPerformanceTier() PerformanceTierModel {
	if o == nil  {
		var ret PerformanceTierModel
		return ret
	}

	return o.PerformanceTier
}

// GetPerformanceTierOk returns a tuple with the PerformanceTier field value
// and a boolean to check if the value has been set.
func (o *ScaleOutRepositoryModel) GetPerformanceTierOk() (*PerformanceTierModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PerformanceTier, true
}

// SetPerformanceTier sets field value
func (o *ScaleOutRepositoryModel) SetPerformanceTier(v PerformanceTierModel) {
	o.PerformanceTier = v
}

// GetPlacementPolicy returns the PlacementPolicy field value if set, zero value otherwise.
func (o *ScaleOutRepositoryModel) GetPlacementPolicy() PlacementPolicyModel {
	if o == nil || o.PlacementPolicy == nil {
		var ret PlacementPolicyModel
		return ret
	}
	return *o.PlacementPolicy
}

// GetPlacementPolicyOk returns a tuple with the PlacementPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleOutRepositoryModel) GetPlacementPolicyOk() (*PlacementPolicyModel, bool) {
	if o == nil || o.PlacementPolicy == nil {
		return nil, false
	}
	return o.PlacementPolicy, true
}

// HasPlacementPolicy returns a boolean if a field has been set.
func (o *ScaleOutRepositoryModel) HasPlacementPolicy() bool {
	if o != nil && o.PlacementPolicy != nil {
		return true
	}

	return false
}

// SetPlacementPolicy gets a reference to the given PlacementPolicyModel and assigns it to the PlacementPolicy field.
func (o *ScaleOutRepositoryModel) SetPlacementPolicy(v PlacementPolicyModel) {
	o.PlacementPolicy = &v
}

// GetCapacityTier returns the CapacityTier field value if set, zero value otherwise.
func (o *ScaleOutRepositoryModel) GetCapacityTier() CapacityTierModel {
	if o == nil || o.CapacityTier == nil {
		var ret CapacityTierModel
		return ret
	}
	return *o.CapacityTier
}

// GetCapacityTierOk returns a tuple with the CapacityTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleOutRepositoryModel) GetCapacityTierOk() (*CapacityTierModel, bool) {
	if o == nil || o.CapacityTier == nil {
		return nil, false
	}
	return o.CapacityTier, true
}

// HasCapacityTier returns a boolean if a field has been set.
func (o *ScaleOutRepositoryModel) HasCapacityTier() bool {
	if o != nil && o.CapacityTier != nil {
		return true
	}

	return false
}

// SetCapacityTier gets a reference to the given CapacityTierModel and assigns it to the CapacityTier field.
func (o *ScaleOutRepositoryModel) SetCapacityTier(v CapacityTierModel) {
	o.CapacityTier = &v
}

// GetArchiveTier returns the ArchiveTier field value if set, zero value otherwise.
func (o *ScaleOutRepositoryModel) GetArchiveTier() ArchiveTierModel {
	if o == nil || o.ArchiveTier == nil {
		var ret ArchiveTierModel
		return ret
	}
	return *o.ArchiveTier
}

// GetArchiveTierOk returns a tuple with the ArchiveTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaleOutRepositoryModel) GetArchiveTierOk() (*ArchiveTierModel, bool) {
	if o == nil || o.ArchiveTier == nil {
		return nil, false
	}
	return o.ArchiveTier, true
}

// HasArchiveTier returns a boolean if a field has been set.
func (o *ScaleOutRepositoryModel) HasArchiveTier() bool {
	if o != nil && o.ArchiveTier != nil {
		return true
	}

	return false
}

// SetArchiveTier gets a reference to the given ArchiveTierModel and assigns it to the ArchiveTier field.
func (o *ScaleOutRepositoryModel) SetArchiveTier(v ArchiveTierModel) {
	o.ArchiveTier = &v
}

func (o ScaleOutRepositoryModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
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
		toSerialize["performanceTier"] = o.PerformanceTier
	}
	if o.PlacementPolicy != nil {
		toSerialize["placementPolicy"] = o.PlacementPolicy
	}
	if o.CapacityTier != nil {
		toSerialize["capacityTier"] = o.CapacityTier
	}
	if o.ArchiveTier != nil {
		toSerialize["archiveTier"] = o.ArchiveTier
	}
	return json.Marshal(toSerialize)
}

type NullableScaleOutRepositoryModel struct {
	value *ScaleOutRepositoryModel
	isSet bool
}

func (v NullableScaleOutRepositoryModel) Get() *ScaleOutRepositoryModel {
	return v.value
}

func (v *NullableScaleOutRepositoryModel) Set(val *ScaleOutRepositoryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableScaleOutRepositoryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableScaleOutRepositoryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScaleOutRepositoryModel(val *ScaleOutRepositoryModel) *NullableScaleOutRepositoryModel {
	return &NullableScaleOutRepositoryModel{value: val, isSet: true}
}

func (v NullableScaleOutRepositoryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScaleOutRepositoryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

