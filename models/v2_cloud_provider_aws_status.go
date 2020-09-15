// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2CloudProviderAwsStatus An array of AWS ELB statuses.
//
// swagger:model V2CloudProviderAwsStatus
type V2CloudProviderAwsStatus struct {

	// elbs
	Elbs []*V2CloudProviderAwsElbStatus `json:"elbs"`
}

// Validate validates this v2 cloud provider aws status
func (m *V2CloudProviderAwsStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElbs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2CloudProviderAwsStatus) validateElbs(formats strfmt.Registry) error {

	if swag.IsZero(m.Elbs) { // not required
		return nil
	}

	for i := 0; i < len(m.Elbs); i++ {
		if swag.IsZero(m.Elbs[i]) { // not required
			continue
		}

		if m.Elbs[i] != nil {
			if err := m.Elbs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elbs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2CloudProviderAwsStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2CloudProviderAwsStatus) UnmarshalBinary(b []byte) error {
	var res V2CloudProviderAwsStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
