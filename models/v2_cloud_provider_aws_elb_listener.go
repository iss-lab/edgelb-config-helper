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

// V2CloudProviderAwsElbListener An AWS ELB listener configuration.
//
// swagger:model V2CloudProviderAwsElbListener
type V2CloudProviderAwsElbListener struct {

	// An Edge-LB frontend to proxy requests to.
	// Required: true
	// Min Length: 1
	LinkFrontend *string `json:"linkFrontend"`

	// An AWS ELB port.
	// Required: true
	// Maximum: 65535
	// Minimum: 1
	Port *int32 `json:"port"`

	// A networking protocol.
	// Enum: [TCP TLS]
	Protocol string `json:"protocol,omitempty"`

	// tls
	TLS *V2CloudProviderAwsElbListenerTLS `json:"tls,omitempty"`
}

// Validate validates this v2 cloud provider aws elb listener
func (m *V2CloudProviderAwsElbListener) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinkFrontend(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2CloudProviderAwsElbListener) validateLinkFrontend(formats strfmt.Registry) error {

	if err := validate.Required("linkFrontend", "body", m.LinkFrontend); err != nil {
		return err
	}

	if err := validate.MinLength("linkFrontend", "body", string(*m.LinkFrontend), 1); err != nil {
		return err
	}

	return nil
}

func (m *V2CloudProviderAwsElbListener) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	if err := validate.MinimumInt("port", "body", int64(*m.Port), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(*m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

var v2CloudProviderAwsElbListenerTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TCP","TLS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v2CloudProviderAwsElbListenerTypeProtocolPropEnum = append(v2CloudProviderAwsElbListenerTypeProtocolPropEnum, v)
	}
}

const (

	// V2CloudProviderAwsElbListenerProtocolTCP captures enum value "TCP"
	V2CloudProviderAwsElbListenerProtocolTCP string = "TCP"

	// V2CloudProviderAwsElbListenerProtocolTLS captures enum value "TLS"
	V2CloudProviderAwsElbListenerProtocolTLS string = "TLS"
)

// prop value enum
func (m *V2CloudProviderAwsElbListener) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v2CloudProviderAwsElbListenerTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V2CloudProviderAwsElbListener) validateProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *V2CloudProviderAwsElbListener) validateTLS(formats strfmt.Registry) error {

	if swag.IsZero(m.TLS) { // not required
		return nil
	}

	if m.TLS != nil {
		if err := m.TLS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2CloudProviderAwsElbListener) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2CloudProviderAwsElbListener) UnmarshalBinary(b []byte) error {
	var res V2CloudProviderAwsElbListener
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
