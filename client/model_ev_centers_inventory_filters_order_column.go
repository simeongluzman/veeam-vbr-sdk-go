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

// EvCentersInventoryFiltersOrderColumn Sorts vCenter Servers by one of the job parameters.
type EvCentersInventoryFiltersOrderColumn string

// List of EvCentersInventoryFiltersOrderColumn
const (
	EVCENTERSINVENTORYFILTERSORDERCOLUMN_NAME EvCentersInventoryFiltersOrderColumn = "Name"
	EVCENTERSINVENTORYFILTERSORDERCOLUMN_OBJECT_ID EvCentersInventoryFiltersOrderColumn = "ObjectId"
	EVCENTERSINVENTORYFILTERSORDERCOLUMN_TYPE EvCentersInventoryFiltersOrderColumn = "Type"
)

// All allowed values of EvCentersInventoryFiltersOrderColumn enum
var AllowedEvCentersInventoryFiltersOrderColumnEnumValues = []EvCentersInventoryFiltersOrderColumn{
	"Name",
	"ObjectId",
	"Type",
}

func (v *EvCentersInventoryFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EvCentersInventoryFiltersOrderColumn(value)
	for _, existing := range AllowedEvCentersInventoryFiltersOrderColumnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EvCentersInventoryFiltersOrderColumn", value)
}

// NewEvCentersInventoryFiltersOrderColumnFromValue returns a pointer to a valid EvCentersInventoryFiltersOrderColumn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEvCentersInventoryFiltersOrderColumnFromValue(v string) (*EvCentersInventoryFiltersOrderColumn, error) {
	ev := EvCentersInventoryFiltersOrderColumn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EvCentersInventoryFiltersOrderColumn: valid values are %v", v, AllowedEvCentersInventoryFiltersOrderColumnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EvCentersInventoryFiltersOrderColumn) IsValid() bool {
	for _, existing := range AllowedEvCentersInventoryFiltersOrderColumnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EvCentersInventoryFiltersOrderColumn value
func (v EvCentersInventoryFiltersOrderColumn) Ptr() *EvCentersInventoryFiltersOrderColumn {
	return &v
}

type NullableEvCentersInventoryFiltersOrderColumn struct {
	value *EvCentersInventoryFiltersOrderColumn
	isSet bool
}

func (v NullableEvCentersInventoryFiltersOrderColumn) Get() *EvCentersInventoryFiltersOrderColumn {
	return v.value
}

func (v *NullableEvCentersInventoryFiltersOrderColumn) Set(val *EvCentersInventoryFiltersOrderColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableEvCentersInventoryFiltersOrderColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableEvCentersInventoryFiltersOrderColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvCentersInventoryFiltersOrderColumn(val *EvCentersInventoryFiltersOrderColumn) *NullableEvCentersInventoryFiltersOrderColumn {
	return &NullableEvCentersInventoryFiltersOrderColumn{value: val, isSet: true}
}

func (v NullableEvCentersInventoryFiltersOrderColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvCentersInventoryFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

