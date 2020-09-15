// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Stats v1 stats
//
// swagger:model V1Stats
type V1Stats struct {

	// bind address
	BindAddress string `json:"bindAddress,omitempty"`

	// bind port
	BindPort int32 `json:"bindPort,omitempty"`
}

func (m *V1Stats) UnmarshalJSON(b []byte) error {
	type V1StatsAlias V1Stats
	var t V1StatsAlias
	if err := json.Unmarshal([]byte("{\"bindAddress\":\"0.0.0.0\",\"bindPort\":9090}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = V1Stats(t)
	return nil
}

// Validate validates this v1 stats
func (m *V1Stats) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Stats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Stats) UnmarshalBinary(b []byte) error {
	var res V1Stats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
