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

// ERepositoryType Repository type.
type ERepositoryType string

// List of ERepositoryType
const (
	EREPOSITORYTYPE_WIN_LOCAL ERepositoryType = "WinLocal"
	EREPOSITORYTYPE_LINUX_LOCAL ERepositoryType = "LinuxLocal"
	EREPOSITORYTYPE_SMB ERepositoryType = "Smb"
	EREPOSITORYTYPE_NFS ERepositoryType = "Nfs"
	EREPOSITORYTYPE_AZURE_BLOB ERepositoryType = "AzureBlob"
	EREPOSITORYTYPE_AZURE_DATA_BOX ERepositoryType = "AzureDataBox"
	EREPOSITORYTYPE_AZURE_ARCHIVE ERepositoryType = "AzureArchive"
	EREPOSITORYTYPE_AMAZON_S3 ERepositoryType = "AmazonS3"
	EREPOSITORYTYPE_AMAZON_SNOWBALL_EDGE ERepositoryType = "AmazonSnowballEdge"
	EREPOSITORYTYPE_AMAZON_S3_GLACIER ERepositoryType = "AmazonS3Glacier"
	EREPOSITORYTYPE_S3_COMPATIBLE ERepositoryType = "S3Compatible"
	EREPOSITORYTYPE_GOOGLE_CLOUD ERepositoryType = "GoogleCloud"
	EREPOSITORYTYPE_IBM_CLOUD ERepositoryType = "IBMCloud"
	EREPOSITORYTYPE_EXTENDABLE_REPOSITORY ERepositoryType = "ExtendableRepository"
	EREPOSITORYTYPE_DD_BOOST ERepositoryType = "DDBoost"
	EREPOSITORYTYPE_EXA_GRID ERepositoryType = "ExaGrid"
	EREPOSITORYTYPE_HP_STORE_ONCE_INTEGRATION ERepositoryType = "HPStoreOnceIntegration"
	EREPOSITORYTYPE_QUANTUM ERepositoryType = "Quantum"
	EREPOSITORYTYPE_WASABI_CLOUD ERepositoryType = "WasabiCloud"
	EREPOSITORYTYPE_LINUX_HARDENED ERepositoryType = "LinuxHardened"
	EREPOSITORYTYPE_INFINIDAT ERepositoryType = "Infinidat"
	EREPOSITORYTYPE_FUJITSU ERepositoryType = "Fujitsu"
)

// All allowed values of ERepositoryType enum
var AllowedERepositoryTypeEnumValues = []ERepositoryType{
	"WinLocal",
	"LinuxLocal",
	"Smb",
	"Nfs",
	"AzureBlob",
	"AzureDataBox",
	"AzureArchive",
	"AmazonS3",
	"AmazonSnowballEdge",
	"AmazonS3Glacier",
	"S3Compatible",
	"GoogleCloud",
	"IBMCloud",
	"ExtendableRepository",
	"DDBoost",
	"ExaGrid",
	"HPStoreOnceIntegration",
	"Quantum",
	"WasabiCloud",
	"LinuxHardened",
	"Infinidat",
	"Fujitsu",
}

func (v *ERepositoryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ERepositoryType(value)
	for _, existing := range AllowedERepositoryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ERepositoryType", value)
}

// NewERepositoryTypeFromValue returns a pointer to a valid ERepositoryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewERepositoryTypeFromValue(v string) (*ERepositoryType, error) {
	ev := ERepositoryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ERepositoryType: valid values are %v", v, AllowedERepositoryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ERepositoryType) IsValid() bool {
	for _, existing := range AllowedERepositoryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ERepositoryType value
func (v ERepositoryType) Ptr() *ERepositoryType {
	return &v
}

type NullableERepositoryType struct {
	value *ERepositoryType
	isSet bool
}

func (v NullableERepositoryType) Get() *ERepositoryType {
	return v.value
}

func (v *NullableERepositoryType) Set(val *ERepositoryType) {
	v.value = val
	v.isSet = true
}

func (v NullableERepositoryType) IsSet() bool {
	return v.isSet
}

func (v *NullableERepositoryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableERepositoryType(val *ERepositoryType) *NullableERepositoryType {
	return &NullableERepositoryType{value: val, isSet: true}
}

func (v NullableERepositoryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableERepositoryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
