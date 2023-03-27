/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Some request and response bodies refer to reusable schema objects that are defined in the `schemas` section of the REST API specification. Schemas, in turn, may inherit a part of their properties by referring to other schemas, and be polymorphic by referring to multiple alternate schemas.

API version: 1.1-rev0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// ECredentialsStorageType Credentials type used to connect to the Linux server.
type ECredentialsStorageType string

// List of ECredentialsStorageType
const (
	ECREDENTIALSSTORAGETYPE_PERMANENT ECredentialsStorageType = "Permanent"
	ECREDENTIALSSTORAGETYPE_SINGLE_USE ECredentialsStorageType = "SingleUse"
)

// All allowed values of ECredentialsStorageType enum
var AllowedECredentialsStorageTypeEnumValues = []ECredentialsStorageType{
	"Permanent",
	"SingleUse",
}

func (v *ECredentialsStorageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ECredentialsStorageType(value)
	for _, existing := range AllowedECredentialsStorageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ECredentialsStorageType", value)
}

// NewECredentialsStorageTypeFromValue returns a pointer to a valid ECredentialsStorageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewECredentialsStorageTypeFromValue(v string) (*ECredentialsStorageType, error) {
	ev := ECredentialsStorageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ECredentialsStorageType: valid values are %v", v, AllowedECredentialsStorageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ECredentialsStorageType) IsValid() bool {
	for _, existing := range AllowedECredentialsStorageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ECredentialsStorageType value
func (v ECredentialsStorageType) Ptr() *ECredentialsStorageType {
	return &v
}

type NullableECredentialsStorageType struct {
	value *ECredentialsStorageType
	isSet bool
}

func (v NullableECredentialsStorageType) Get() *ECredentialsStorageType {
	return v.value
}

func (v *NullableECredentialsStorageType) Set(val *ECredentialsStorageType) {
	v.value = val
	v.isSet = true
}

func (v NullableECredentialsStorageType) IsSet() bool {
	return v.isSet
}

func (v *NullableECredentialsStorageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECredentialsStorageType(val *ECredentialsStorageType) *NullableECredentialsStorageType {
	return &NullableECredentialsStorageType{value: val, isSet: true}
}

func (v NullableECredentialsStorageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECredentialsStorageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
