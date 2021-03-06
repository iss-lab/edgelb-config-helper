// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PoolContainer Object used internally as an interface to handle multple model versions.
//
// swagger:model PoolContainer
type PoolContainer struct {

	// The api / schema version of this pool object
	APIVersion APIVersion `json:"apiVersion,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace *string `json:"namespace,omitempty"`

	// v1
	V1 *V1Pool `json:"v1,omitempty"`

	// v2
	V2 *V2Pool `json:"v2,omitempty"`
}

func (m *PoolContainer) UnmarshalJSON(b []byte) error {
	type PoolContainerAlias PoolContainer
	var t PoolContainerAlias
	if err := json.Unmarshal([]byte("{\"apiVersion\":\"V1\"}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = PoolContainer(t)
	return nil
}

// Validate validates this pool container
func (m *PoolContainer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateV1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateV2(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolContainer) validateAPIVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.APIVersion) { // not required
		return nil
	}

	if err := m.APIVersion.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("apiVersion")
		}
		return err
	}

	return nil
}

func (m *PoolContainer) validateV1(formats strfmt.Registry) error {

	if swag.IsZero(m.V1) { // not required
		return nil
	}

	if m.V1 != nil {
		if err := m.V1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("v1")
			}
			return err
		}
	}

	return nil
}

func (m *PoolContainer) validateV2(formats strfmt.Registry) error {

	if swag.IsZero(m.V2) { // not required
		return nil
	}

	if m.V2 != nil {
		if err := m.V2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("v2")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoolContainer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoolContainer) UnmarshalBinary(b []byte) error {
	var res PoolContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
