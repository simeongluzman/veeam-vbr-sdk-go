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

// ERepositoryFiltersOrderColumn Sorts repositories by one of the repository parameters.
type ERepositoryFiltersOrderColumn string

// List of ERepositoryFiltersOrderColumn
const (
	EREPOSITORYFILTERSORDERCOLUMN_NAME ERepositoryFiltersOrderColumn = "Name"
	EREPOSITORYFILTERSORDERCOLUMN_DESCRIPTION ERepositoryFiltersOrderColumn = "Description"
	EREPOSITORYFILTERSORDERCOLUMN_TYPE ERepositoryFiltersOrderColumn = "Type"
	EREPOSITORYFILTERSORDERCOLUMN_HOST ERepositoryFiltersOrderColumn = "Host"
	EREPOSITORYFILTERSORDERCOLUMN_PATH ERepositoryFiltersOrderColumn = "Path"
)

// All allowed values of ERepositoryFiltersOrderColumn enum
var AllowedERepositoryFiltersOrderColumnEnumValues = []ERepositoryFiltersOrderColumn{
	"Name",
	"Description",
	"Type",
	"Host",
	"Path",
}

func (v *ERepositoryFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ERepositoryFiltersOrderColumn(value)
	for _, existing := range AllowedERepositoryFiltersOrderColumnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ERepositoryFiltersOrderColumn", value)
}

// NewERepositoryFiltersOrderColumnFromValue returns a pointer to a valid ERepositoryFiltersOrderColumn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewERepositoryFiltersOrderColumnFromValue(v string) (*ERepositoryFiltersOrderColumn, error) {
	ev := ERepositoryFiltersOrderColumn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ERepositoryFiltersOrderColumn: valid values are %v", v, AllowedERepositoryFiltersOrderColumnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ERepositoryFiltersOrderColumn) IsValid() bool {
	for _, existing := range AllowedERepositoryFiltersOrderColumnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ERepositoryFiltersOrderColumn value
func (v ERepositoryFiltersOrderColumn) Ptr() *ERepositoryFiltersOrderColumn {
	return &v
}

type NullableERepositoryFiltersOrderColumn struct {
	value *ERepositoryFiltersOrderColumn
	isSet bool
}

func (v NullableERepositoryFiltersOrderColumn) Get() *ERepositoryFiltersOrderColumn {
	return v.value
}

func (v *NullableERepositoryFiltersOrderColumn) Set(val *ERepositoryFiltersOrderColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableERepositoryFiltersOrderColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableERepositoryFiltersOrderColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableERepositoryFiltersOrderColumn(val *ERepositoryFiltersOrderColumn) *NullableERepositoryFiltersOrderColumn {
	return &NullableERepositoryFiltersOrderColumn{value: val, isSet: true}
}

func (v NullableERepositoryFiltersOrderColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableERepositoryFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

