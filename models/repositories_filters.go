// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RepositoriesFilters repositories filters
//
// swagger:model RepositoriesFilters
type RepositoriesFilters struct {

	// host Id filter
	// Format: uuid
	HostIDFilter strfmt.UUID `json:"hostIdFilter,omitempty"`

	// limit
	Limit int32 `json:"limit,omitempty"`

	// name filter
	NameFilter string `json:"nameFilter,omitempty"`

	// order asc
	OrderAsc bool `json:"orderAsc,omitempty"`

	// order column
	OrderColumn ERepositoryFiltersOrderColumn `json:"orderColumn,omitempty"`

	// path filter
	PathFilter string `json:"pathFilter,omitempty"`

	// skip
	Skip int32 `json:"skip,omitempty"`

	// type filter
	TypeFilter ERepositoryType `json:"typeFilter,omitempty"`

	// VmbApiFilterModel as json serialized in base64 format (see VmbApiFilterModel)
	VmbAPIFilter string `json:"vmbApiFilter,omitempty"`
}

// Validate validates this repositories filters
func (m *RepositoriesFilters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostIDFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderColumn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeFilter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepositoriesFilters) validateHostIDFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.HostIDFilter) { // not required
		return nil
	}

	if err := validate.FormatOf("hostIdFilter", "body", "uuid", m.HostIDFilter.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RepositoriesFilters) validateOrderColumn(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderColumn) { // not required
		return nil
	}

	if err := m.OrderColumn.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("orderColumn")
		}
		return err
	}

	return nil
}

func (m *RepositoriesFilters) validateTypeFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeFilter) { // not required
		return nil
	}

	if err := m.TypeFilter.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("typeFilter")
		}
		return err
	}

	return nil
}

// ContextValidate validate this repositories filters based on the context it is used
func (m *RepositoriesFilters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrderColumn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepositoriesFilters) contextValidateOrderColumn(ctx context.Context, formats strfmt.Registry) error {

	if err := m.OrderColumn.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("orderColumn")
		}
		return err
	}

	return nil
}

func (m *RepositoriesFilters) contextValidateTypeFilter(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TypeFilter.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("typeFilter")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RepositoriesFilters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RepositoriesFilters) UnmarshalBinary(b []byte) error {
	var res RepositoriesFilters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}