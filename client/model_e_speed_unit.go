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

// ESpeedUnit Traffic speed unit.
type ESpeedUnit string

// List of ESpeedUnit
const (
	ESPEEDUNIT_KBYTE_PER_SEC ESpeedUnit = "KbytePerSec"
	ESPEEDUNIT_MBIT_PER_SPEC ESpeedUnit = "MbitPerSpec"
	ESPEEDUNIT_MBYTE_PER_SEC ESpeedUnit = "MbytePerSec"
)

// All allowed values of ESpeedUnit enum
var AllowedESpeedUnitEnumValues = []ESpeedUnit{
	"KbytePerSec",
	"MbitPerSpec",
	"MbytePerSec",
}

func (v *ESpeedUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ESpeedUnit(value)
	for _, existing := range AllowedESpeedUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ESpeedUnit", value)
}

// NewESpeedUnitFromValue returns a pointer to a valid ESpeedUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewESpeedUnitFromValue(v string) (*ESpeedUnit, error) {
	ev := ESpeedUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ESpeedUnit: valid values are %v", v, AllowedESpeedUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ESpeedUnit) IsValid() bool {
	for _, existing := range AllowedESpeedUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ESpeedUnit value
func (v ESpeedUnit) Ptr() *ESpeedUnit {
	return &v
}

type NullableESpeedUnit struct {
	value *ESpeedUnit
	isSet bool
}

func (v NullableESpeedUnit) Get() *ESpeedUnit {
	return v.value
}

func (v *NullableESpeedUnit) Set(val *ESpeedUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableESpeedUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableESpeedUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableESpeedUnit(val *ESpeedUnit) *NullableESpeedUnit {
	return &NullableESpeedUnit{value: val, isSet: true}
}

func (v NullableESpeedUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableESpeedUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
