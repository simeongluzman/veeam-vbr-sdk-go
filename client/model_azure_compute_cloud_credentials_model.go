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

// AzureComputeCloudCredentialsModel Microsoft Azure compute account.
type AzureComputeCloudCredentialsModel struct {
	CloudCredentialsModel
	// Name under which the cloud credentials record is shown in Veeam Backup & Replication.
	ConnectionName string `json:"connectionName"`
	Deployment AzureComputeCloudCredentialsDeploymentModel `json:"deployment"`
	Subscription AzureComputeCloudCredentialsSubscriptionModel `json:"subscription"`
	// Tag used to identify the cloud credentials record.
	Tag *string `json:"tag,omitempty"`
}

// NewAzureComputeCloudCredentialsModel instantiates a new AzureComputeCloudCredentialsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureComputeCloudCredentialsModel(connectionName string, deployment AzureComputeCloudCredentialsDeploymentModel, subscription AzureComputeCloudCredentialsSubscriptionModel, id string, type_ ECloudCredentialsType) *AzureComputeCloudCredentialsModel {
	this := AzureComputeCloudCredentialsModel{}
	this.Id = id
	this.Type = type_
	this.ConnectionName = connectionName
	this.Deployment = deployment
	this.Subscription = subscription
	return &this
}

// NewAzureComputeCloudCredentialsModelWithDefaults instantiates a new AzureComputeCloudCredentialsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureComputeCloudCredentialsModelWithDefaults() *AzureComputeCloudCredentialsModel {
	this := AzureComputeCloudCredentialsModel{}
	return &this
}

// GetConnectionName returns the ConnectionName field value
func (o *AzureComputeCloudCredentialsModel) GetConnectionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionName
}

// GetConnectionNameOk returns a tuple with the ConnectionName field value
// and a boolean to check if the value has been set.
func (o *AzureComputeCloudCredentialsModel) GetConnectionNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConnectionName, true
}

// SetConnectionName sets field value
func (o *AzureComputeCloudCredentialsModel) SetConnectionName(v string) {
	o.ConnectionName = v
}

// GetDeployment returns the Deployment field value
func (o *AzureComputeCloudCredentialsModel) GetDeployment() AzureComputeCloudCredentialsDeploymentModel {
	if o == nil {
		var ret AzureComputeCloudCredentialsDeploymentModel
		return ret
	}

	return o.Deployment
}

// GetDeploymentOk returns a tuple with the Deployment field value
// and a boolean to check if the value has been set.
func (o *AzureComputeCloudCredentialsModel) GetDeploymentOk() (*AzureComputeCloudCredentialsDeploymentModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Deployment, true
}

// SetDeployment sets field value
func (o *AzureComputeCloudCredentialsModel) SetDeployment(v AzureComputeCloudCredentialsDeploymentModel) {
	o.Deployment = v
}

// GetSubscription returns the Subscription field value
func (o *AzureComputeCloudCredentialsModel) GetSubscription() AzureComputeCloudCredentialsSubscriptionModel {
	if o == nil {
		var ret AzureComputeCloudCredentialsSubscriptionModel
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *AzureComputeCloudCredentialsModel) GetSubscriptionOk() (*AzureComputeCloudCredentialsSubscriptionModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *AzureComputeCloudCredentialsModel) SetSubscription(v AzureComputeCloudCredentialsSubscriptionModel) {
	o.Subscription = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *AzureComputeCloudCredentialsModel) GetTag() string {
	if o == nil || isNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureComputeCloudCredentialsModel) GetTagOk() (*string, bool) {
	if o == nil || isNil(o.Tag) {
    return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *AzureComputeCloudCredentialsModel) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *AzureComputeCloudCredentialsModel) SetTag(v string) {
	o.Tag = &v
}

func (o AzureComputeCloudCredentialsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudCredentialsModel, errCloudCredentialsModel := json.Marshal(o.CloudCredentialsModel)
	if errCloudCredentialsModel != nil {
		return []byte{}, errCloudCredentialsModel
	}
	errCloudCredentialsModel = json.Unmarshal([]byte(serializedCloudCredentialsModel), &toSerialize)
	if errCloudCredentialsModel != nil {
		return []byte{}, errCloudCredentialsModel
	}
	if true {
		toSerialize["connectionName"] = o.ConnectionName
	}
	if true {
		toSerialize["deployment"] = o.Deployment
	}
	if true {
		toSerialize["subscription"] = o.Subscription
	}
	if !isNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableAzureComputeCloudCredentialsModel struct {
	value *AzureComputeCloudCredentialsModel
	isSet bool
}

func (v NullableAzureComputeCloudCredentialsModel) Get() *AzureComputeCloudCredentialsModel {
	return v.value
}

func (v *NullableAzureComputeCloudCredentialsModel) Set(val *AzureComputeCloudCredentialsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureComputeCloudCredentialsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureComputeCloudCredentialsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureComputeCloudCredentialsModel(val *AzureComputeCloudCredentialsModel) *NullableAzureComputeCloudCredentialsModel {
	return &NullableAzureComputeCloudCredentialsModel{value: val, isSet: true}
}

func (v NullableAzureComputeCloudCredentialsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureComputeCloudCredentialsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


