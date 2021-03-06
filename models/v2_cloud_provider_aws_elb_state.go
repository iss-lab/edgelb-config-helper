// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V2CloudProviderAwsElbState v2 cloud provider aws elb state
//
// swagger:model V2CloudProviderAwsElbState
type V2CloudProviderAwsElbState struct {

	// The reason explaining the current load balancer status.
	Reason *string `json:"reason,omitempty"`

	// The status of a load balancer.
	// Enum: [initializing pending active failed]
	Status string `json:"status,omitempty"`
}

func (m *V2CloudProviderAwsElbState) UnmarshalJSON(b []byte) error {
	type V2CloudProviderAwsElbStateAlias V2CloudProviderAwsElbState
	var t V2CloudProviderAwsElbStateAlias
	if err := json.Unmarshal([]byte("{\"status\":\"initializing\"}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = V2CloudProviderAwsElbState(t)
	return nil
}

// Validate validates this v2 cloud provider aws elb state
func (m *V2CloudProviderAwsElbState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v2CloudProviderAwsElbStateTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["initializing","pending","active","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v2CloudProviderAwsElbStateTypeStatusPropEnum = append(v2CloudProviderAwsElbStateTypeStatusPropEnum, v)
	}
}

const (

	// V2CloudProviderAwsElbStateStatusInitializing captures enum value "initializing"
	V2CloudProviderAwsElbStateStatusInitializing string = "initializing"

	// V2CloudProviderAwsElbStateStatusPending captures enum value "pending"
	V2CloudProviderAwsElbStateStatusPending string = "pending"

	// V2CloudProviderAwsElbStateStatusActive captures enum value "active"
	V2CloudProviderAwsElbStateStatusActive string = "active"

	// V2CloudProviderAwsElbStateStatusFailed captures enum value "failed"
	V2CloudProviderAwsElbStateStatusFailed string = "failed"
)

// prop value enum
func (m *V2CloudProviderAwsElbState) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v2CloudProviderAwsElbStateTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V2CloudProviderAwsElbState) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2CloudProviderAwsElbState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2CloudProviderAwsElbState) UnmarshalBinary(b []byte) error {
	var res V2CloudProviderAwsElbState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
