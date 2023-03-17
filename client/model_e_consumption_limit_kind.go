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

// EConsumptionLimitKind Measurement unit.
type EConsumptionLimitKind string

// List of EConsumptionLimitKind
const (
	ECONSUMPTIONLIMITKIND_TB EConsumptionLimitKind = "TB"
	ECONSUMPTIONLIMITKIND_PB EConsumptionLimitKind = "PB"
)

// All allowed values of EConsumptionLimitKind enum
var AllowedEConsumptionLimitKindEnumValues = []EConsumptionLimitKind{
	"TB",
	"PB",
}

func (v *EConsumptionLimitKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EConsumptionLimitKind(value)
	for _, existing := range AllowedEConsumptionLimitKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EConsumptionLimitKind", value)
}

// NewEConsumptionLimitKindFromValue returns a pointer to a valid EConsumptionLimitKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEConsumptionLimitKindFromValue(v string) (*EConsumptionLimitKind, error) {
	ev := EConsumptionLimitKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EConsumptionLimitKind: valid values are %v", v, AllowedEConsumptionLimitKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EConsumptionLimitKind) IsValid() bool {
	for _, existing := range AllowedEConsumptionLimitKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EConsumptionLimitKind value
func (v EConsumptionLimitKind) Ptr() *EConsumptionLimitKind {
	return &v
}

type NullableEConsumptionLimitKind struct {
	value *EConsumptionLimitKind
	isSet bool
}

func (v NullableEConsumptionLimitKind) Get() *EConsumptionLimitKind {
	return v.value
}

func (v *NullableEConsumptionLimitKind) Set(val *EConsumptionLimitKind) {
	v.value = val
	v.isSet = true
}

func (v NullableEConsumptionLimitKind) IsSet() bool {
	return v.isSet
}

func (v *NullableEConsumptionLimitKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEConsumptionLimitKind(val *EConsumptionLimitKind) *NullableEConsumptionLimitKind {
	return &NullableEConsumptionLimitKind{value: val, isSet: true}
}

func (v NullableEConsumptionLimitKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEConsumptionLimitKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

