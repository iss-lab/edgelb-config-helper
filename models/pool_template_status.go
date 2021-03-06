// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PoolTemplateStatus Describes the status of the given pool template, e.g. rendering information
//
// swagger:model PoolTemplateStatus
type PoolTemplateStatus map[string]PoolTemplateStatusAnon

// Validate validates this pool template status
func (m PoolTemplateStatus) Validate(formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if swag.IsZero(m[k]) { // not required
			continue
		}
		if val, ok := m[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// PoolTemplateStatusAnon pool template status anon
//
// swagger:model PoolTemplateStatusAnon
type PoolTemplateStatusAnon struct {

	// message
	Message string `json:"message,omitempty"`

	// Describes the status of the rendering of the given pool template into a given pool
	Status RenderStatus `json:"status,omitempty"`

	// Describes the per task status of the rendering of the given pool template
	Statuses map[string]TaskStatus `json:"statuses,omitempty"`
}

// Validate validates this pool template status anon
func (m *PoolTemplateStatusAnon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatuses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolTemplateStatusAnon) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *PoolTemplateStatusAnon) validateStatuses(formats strfmt.Registry) error {

	if swag.IsZero(m.Statuses) { // not required
		return nil
	}

	for k := range m.Statuses {

		if val, ok := m.Statuses[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoolTemplateStatusAnon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoolTemplateStatusAnon) UnmarshalBinary(b []byte) error {
	var res PoolTemplateStatusAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
