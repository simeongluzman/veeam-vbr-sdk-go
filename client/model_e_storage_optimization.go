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
	"fmt"
)

// EStorageOptimization Storage optimization that depends on the target type.
type EStorageOptimization string

// List of EStorageOptimization
const (
	ESTORAGEOPTIMIZATION_AUTO EStorageOptimization = "Auto"
	ESTORAGEOPTIMIZATION_LOCAL_TARGET_LARGE EStorageOptimization = "LocalTargetLarge"
	ESTORAGEOPTIMIZATION_LOCAL_TARGET EStorageOptimization = "LocalTarget"
	ESTORAGEOPTIMIZATION_LAN_TARGET EStorageOptimization = "LANTarget"
	ESTORAGEOPTIMIZATION_WAN_TARGET EStorageOptimization = "WANTarget"
	ESTORAGEOPTIMIZATION_LOCAL_TARGET_LARGE8192 EStorageOptimization = "LocalTargetLarge8192"
	ESTORAGEOPTIMIZATION_LOCAL_TARGET_LARGE4096 EStorageOptimization = "LocalTargetLarge4096"
)

func (v *EStorageOptimization) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EStorageOptimization(value)
	for _, existing := range []EStorageOptimization{ "Auto", "LocalTargetLarge", "LocalTarget", "LANTarget", "WANTarget", "LocalTargetLarge8192", "LocalTargetLarge4096",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EStorageOptimization", value)
}

// Ptr returns reference to EStorageOptimization value
func (v EStorageOptimization) Ptr() *EStorageOptimization {
	return &v
}

type NullableEStorageOptimization struct {
	value *EStorageOptimization
	isSet bool
}

func (v NullableEStorageOptimization) Get() *EStorageOptimization {
	return v.value
}

func (v *NullableEStorageOptimization) Set(val *EStorageOptimization) {
	v.value = val
	v.isSet = true
}

func (v NullableEStorageOptimization) IsSet() bool {
	return v.isSet
}

func (v *NullableEStorageOptimization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEStorageOptimization(val *EStorageOptimization) *NullableEStorageOptimization {
	return &NullableEStorageOptimization{value: val, isSet: true}
}

func (v NullableEStorageOptimization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEStorageOptimization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
