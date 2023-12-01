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

// EVmwareDisksTypeToProcess Type of disk selection.
type EVmwareDisksTypeToProcess string

// List of EVmwareDisksTypeToProcess
const (
	EVMWAREDISKSTYPETOPROCESS_ALL_DISKS EVmwareDisksTypeToProcess = "AllDisks"
	EVMWAREDISKSTYPETOPROCESS_SYSTEM_ONLY EVmwareDisksTypeToProcess = "SystemOnly"
	EVMWAREDISKSTYPETOPROCESS_SELECTED_DISKS EVmwareDisksTypeToProcess = "SelectedDisks"
)

// All allowed values of EVmwareDisksTypeToProcess enum
var AllowedEVmwareDisksTypeToProcessEnumValues = []EVmwareDisksTypeToProcess{
	"AllDisks",
	"SystemOnly",
	"SelectedDisks",
}

func (v *EVmwareDisksTypeToProcess) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EVmwareDisksTypeToProcess(value)
	for _, existing := range AllowedEVmwareDisksTypeToProcessEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EVmwareDisksTypeToProcess", value)
}

// NewEVmwareDisksTypeToProcessFromValue returns a pointer to a valid EVmwareDisksTypeToProcess
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEVmwareDisksTypeToProcessFromValue(v string) (*EVmwareDisksTypeToProcess, error) {
	ev := EVmwareDisksTypeToProcess(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EVmwareDisksTypeToProcess: valid values are %v", v, AllowedEVmwareDisksTypeToProcessEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EVmwareDisksTypeToProcess) IsValid() bool {
	for _, existing := range AllowedEVmwareDisksTypeToProcessEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EVmwareDisksTypeToProcess value
func (v EVmwareDisksTypeToProcess) Ptr() *EVmwareDisksTypeToProcess {
	return &v
}

type NullableEVmwareDisksTypeToProcess struct {
	value *EVmwareDisksTypeToProcess
	isSet bool
}

func (v NullableEVmwareDisksTypeToProcess) Get() *EVmwareDisksTypeToProcess {
	return v.value
}

func (v *NullableEVmwareDisksTypeToProcess) Set(val *EVmwareDisksTypeToProcess) {
	v.value = val
	v.isSet = true
}

func (v NullableEVmwareDisksTypeToProcess) IsSet() bool {
	return v.isSet
}

func (v *NullableEVmwareDisksTypeToProcess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEVmwareDisksTypeToProcess(val *EVmwareDisksTypeToProcess) *NullableEVmwareDisksTypeToProcess {
	return &NullableEVmwareDisksTypeToProcess{value: val, isSet: true}
}

func (v NullableEVmwareDisksTypeToProcess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEVmwareDisksTypeToProcess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
