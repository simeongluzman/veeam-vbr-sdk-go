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

// EDiskInfoType Type of the disk.
type EDiskInfoType string

// List of EDiskInfoType
const (
	EDISKINFOTYPE_UNKNOWN EDiskInfoType = "Unknown"
	EDISKINFOTYPE_SIMPLE EDiskInfoType = "Simple"
	EDISKINFOTYPE_VI EDiskInfoType = "Vi"
	EDISKINFOTYPE_HV EDiskInfoType = "Hv"
	EDISKINFOTYPE_HV_RAW_DISK_FILE EDiskInfoType = "HvRawDiskFile"
	EDISKINFOTYPE_END_POINT EDiskInfoType = "EndPoint"
)

func (v *EDiskInfoType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EDiskInfoType(value)
	for _, existing := range []EDiskInfoType{ "Unknown", "Simple", "Vi", "Hv", "HvRawDiskFile", "EndPoint",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EDiskInfoType", value)
}

// Ptr returns reference to EDiskInfoType value
func (v EDiskInfoType) Ptr() *EDiskInfoType {
	return &v
}

type NullableEDiskInfoType struct {
	value *EDiskInfoType
	isSet bool
}

func (v NullableEDiskInfoType) Get() *EDiskInfoType {
	return v.value
}

func (v *NullableEDiskInfoType) Set(val *EDiskInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableEDiskInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableEDiskInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEDiskInfoType(val *EDiskInfoType) *NullableEDiskInfoType {
	return &NullableEDiskInfoType{value: val, isSet: true}
}

func (v NullableEDiskInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEDiskInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
