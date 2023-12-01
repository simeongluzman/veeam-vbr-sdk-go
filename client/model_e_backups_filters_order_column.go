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

// EBackupsFiltersOrderColumn the model 'EBackupsFiltersOrderColumn'
type EBackupsFiltersOrderColumn string

// List of EBackupsFiltersOrderColumn
const (
	EBACKUPSFILTERSORDERCOLUMN_NAME EBackupsFiltersOrderColumn = "Name"
	EBACKUPSFILTERSORDERCOLUMN_CREATION_TIME EBackupsFiltersOrderColumn = "CreationTime"
	EBACKUPSFILTERSORDERCOLUMN_PLATFORM_ID EBackupsFiltersOrderColumn = "PlatformId"
	EBACKUPSFILTERSORDERCOLUMN_JOB_ID EBackupsFiltersOrderColumn = "JobId"
	EBACKUPSFILTERSORDERCOLUMN_POLICY_TAG EBackupsFiltersOrderColumn = "PolicyTag"
)

// All allowed values of EBackupsFiltersOrderColumn enum
var AllowedEBackupsFiltersOrderColumnEnumValues = []EBackupsFiltersOrderColumn{
	"Name",
	"CreationTime",
	"PlatformId",
	"JobId",
	"PolicyTag",
}

func (v *EBackupsFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EBackupsFiltersOrderColumn(value)
	for _, existing := range AllowedEBackupsFiltersOrderColumnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EBackupsFiltersOrderColumn", value)
}

// NewEBackupsFiltersOrderColumnFromValue returns a pointer to a valid EBackupsFiltersOrderColumn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEBackupsFiltersOrderColumnFromValue(v string) (*EBackupsFiltersOrderColumn, error) {
	ev := EBackupsFiltersOrderColumn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EBackupsFiltersOrderColumn: valid values are %v", v, AllowedEBackupsFiltersOrderColumnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EBackupsFiltersOrderColumn) IsValid() bool {
	for _, existing := range AllowedEBackupsFiltersOrderColumnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EBackupsFiltersOrderColumn value
func (v EBackupsFiltersOrderColumn) Ptr() *EBackupsFiltersOrderColumn {
	return &v
}

type NullableEBackupsFiltersOrderColumn struct {
	value *EBackupsFiltersOrderColumn
	isSet bool
}

func (v NullableEBackupsFiltersOrderColumn) Get() *EBackupsFiltersOrderColumn {
	return v.value
}

func (v *NullableEBackupsFiltersOrderColumn) Set(val *EBackupsFiltersOrderColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableEBackupsFiltersOrderColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableEBackupsFiltersOrderColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEBackupsFiltersOrderColumn(val *EBackupsFiltersOrderColumn) *NullableEBackupsFiltersOrderColumn {
	return &NullableEBackupsFiltersOrderColumn{value: val, isSet: true}
}

func (v NullableEBackupsFiltersOrderColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEBackupsFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
