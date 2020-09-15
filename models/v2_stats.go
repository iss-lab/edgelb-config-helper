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

// V2Stats v2 stats
//
// swagger:model V2Stats
type V2Stats struct {

	// bind address
	BindAddress string `json:"bindAddress,omitempty"`

	// bind port
	// Maximum: 65535
	// Minimum: 0
	BindPort *int32 `json:"bindPort,omitempty"`
}

func (m *V2Stats) UnmarshalJSON(b []byte) error {
	type V2StatsAlias V2Stats
	var t V2StatsAlias
	if err := json.Unmarshal([]byte("{\"bindAddress\":\"0.0.0.0\",\"bindPort\":9090}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = V2Stats(t)
	return nil
}

// Validate validates this v2 stats
func (m *V2Stats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBindPort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2Stats) validateBindPort(formats strfmt.Registry) error {

	if swag.IsZero(m.BindPort) { // not required
		return nil
	}

	if err := validate.MinimumInt("bindPort", "body", int64(*m.BindPort), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("bindPort", "body", int64(*m.BindPort), 65535, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2Stats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2Stats) UnmarshalBinary(b []byte) error {
	var res V2Stats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}