// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PlacementPolicyModel Backup file placement policy.
//
// swagger:model PlacementPolicyModel
type PlacementPolicyModel struct {

	// settings
	// Required: true
	Settings []*BackupPlacementSettingsModel `json:"settings"`

	// type
	// Required: true
	Type *EPlacementPolicyType `json:"type"`
}

// Validate validates this placement policy model
func (m *PlacementPolicyModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlacementPolicyModel) validateSettings(formats strfmt.Registry) error {

	if err := validate.Required("settings", "body", m.Settings); err != nil {
		return err
	}

	for i := 0; i < len(m.Settings); i++ {
		if swag.IsZero(m.Settings[i]) { // not required
			continue
		}

		if m.Settings[i] != nil {
			if err := m.Settings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("settings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlacementPolicyModel) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this placement policy model based on the context it is used
func (m *PlacementPolicyModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlacementPolicyModel) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Settings); i++ {

		if m.Settings[i] != nil {
			if err := m.Settings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("settings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlacementPolicyModel) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlacementPolicyModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlacementPolicyModel) UnmarshalBinary(b []byte) error {
	var res PlacementPolicyModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}