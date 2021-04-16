// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WindowsHostSpec windows host spec
//
// swagger:model WindowsHostSpec
type WindowsHostSpec struct {
	ManagedServerSpec

	// network settings
	NetworkSettings *WindowsHostPortsModel `json:"networkSettings,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WindowsHostSpec) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ManagedServerSpec
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ManagedServerSpec = aO0

	// AO1
	var dataAO1 struct {
		NetworkSettings *WindowsHostPortsModel `json:"networkSettings,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.NetworkSettings = dataAO1.NetworkSettings

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WindowsHostSpec) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ManagedServerSpec)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		NetworkSettings *WindowsHostPortsModel `json:"networkSettings,omitempty"`
	}

	dataAO1.NetworkSettings = m.NetworkSettings

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this windows host spec
func (m *WindowsHostSpec) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ManagedServerSpec
	if err := m.ManagedServerSpec.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WindowsHostSpec) validateNetworkSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkSettings) { // not required
		return nil
	}

	if m.NetworkSettings != nil {
		if err := m.NetworkSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkSettings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this windows host spec based on the context it is used
func (m *WindowsHostSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ManagedServerSpec
	if err := m.ManagedServerSpec.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WindowsHostSpec) contextValidateNetworkSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkSettings != nil {
		if err := m.NetworkSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WindowsHostSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WindowsHostSpec) UnmarshalBinary(b []byte) error {
	var res WindowsHostSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}