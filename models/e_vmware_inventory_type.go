// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// EVmwareInventoryType Type of the VMware vSphere object.
//
// swagger:model EVmwareInventoryType
type EVmwareInventoryType string

func NewEVmwareInventoryType(value EVmwareInventoryType) *EVmwareInventoryType {
	v := value
	return &v
}

const (

	// EVmwareInventoryTypeUnknown captures enum value "Unknown"
	EVmwareInventoryTypeUnknown EVmwareInventoryType = "Unknown"

	// EVmwareInventoryTypeVirtualMachine captures enum value "VirtualMachine"
	EVmwareInventoryTypeVirtualMachine EVmwareInventoryType = "VirtualMachine"

	// EVmwareInventoryTypeVCenterServer captures enum value "vCenterServer"
	EVmwareInventoryTypeVCenterServer EVmwareInventoryType = "vCenterServer"

	// EVmwareInventoryTypeDatacenter captures enum value "Datacenter"
	EVmwareInventoryTypeDatacenter EVmwareInventoryType = "Datacenter"

	// EVmwareInventoryTypeCluster captures enum value "Cluster"
	EVmwareInventoryTypeCluster EVmwareInventoryType = "Cluster"

	// EVmwareInventoryTypeHost captures enum value "Host"
	EVmwareInventoryTypeHost EVmwareInventoryType = "Host"

	// EVmwareInventoryTypeResourcePool captures enum value "ResourcePool"
	EVmwareInventoryTypeResourcePool EVmwareInventoryType = "ResourcePool"

	// EVmwareInventoryTypeFolder captures enum value "Folder"
	EVmwareInventoryTypeFolder EVmwareInventoryType = "Folder"

	// EVmwareInventoryTypeDatastore captures enum value "Datastore"
	EVmwareInventoryTypeDatastore EVmwareInventoryType = "Datastore"

	// EVmwareInventoryTypeDatastoreCluster captures enum value "DatastoreCluster"
	EVmwareInventoryTypeDatastoreCluster EVmwareInventoryType = "DatastoreCluster"

	// EVmwareInventoryTypeStoragePolicy captures enum value "StoragePolicy"
	EVmwareInventoryTypeStoragePolicy EVmwareInventoryType = "StoragePolicy"

	// EVmwareInventoryTypeTemplate captures enum value "Template"
	EVmwareInventoryTypeTemplate EVmwareInventoryType = "Template"

	// EVmwareInventoryTypeComputeResource captures enum value "ComputeResource"
	EVmwareInventoryTypeComputeResource EVmwareInventoryType = "ComputeResource"

	// EVmwareInventoryTypeVirtualApp captures enum value "VirtualApp"
	EVmwareInventoryTypeVirtualApp EVmwareInventoryType = "VirtualApp"

	// EVmwareInventoryTypeTag captures enum value "Tag"
	EVmwareInventoryTypeTag EVmwareInventoryType = "Tag"

	// EVmwareInventoryTypeCategory captures enum value "Category"
	EVmwareInventoryTypeCategory EVmwareInventoryType = "Category"

	// EVmwareInventoryTypeMultitag captures enum value "Multitag"
	EVmwareInventoryTypeMultitag EVmwareInventoryType = "Multitag"
)

// for schema
var eVmwareInventoryTypeEnum []interface{}

func init() {
	var res []EVmwareInventoryType
	if err := json.Unmarshal([]byte(`["Unknown","VirtualMachine","vCenterServer","Datacenter","Cluster","Host","ResourcePool","Folder","Datastore","DatastoreCluster","StoragePolicy","Template","ComputeResource","VirtualApp","Tag","Category","Multitag"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eVmwareInventoryTypeEnum = append(eVmwareInventoryTypeEnum, v)
	}
}

func (m EVmwareInventoryType) validateEVmwareInventoryTypeEnum(path, location string, value EVmwareInventoryType) error {
	if err := validate.EnumCase(path, location, value, eVmwareInventoryTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this e vmware inventory type
func (m EVmwareInventoryType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEVmwareInventoryTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this e vmware inventory type based on context it is used
func (m EVmwareInventoryType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}