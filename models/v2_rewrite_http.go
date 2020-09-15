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

// V2RewriteHTTP v2 rewrite Http
//
// swagger:model V2RewriteHttp
type V2RewriteHTTP struct {

	// Set the host header value
	Host string `json:"host,omitempty"`

	// path
	Path *V2RewriteHTTPPath `json:"path,omitempty"`

	// request
	Request *V2RewriteHTTPRequest `json:"request,omitempty"`

	// response
	Response *V2RewriteHTTPResponse `json:"response,omitempty"`

	// sticky
	Sticky *V2RewriteHTTPSticky `json:"sticky,omitempty"`
}

func (m *V2RewriteHTTP) UnmarshalJSON(b []byte) error {
	type V2RewriteHTTPAlias V2RewriteHTTP
	var t V2RewriteHTTPAlias
	if err := json.Unmarshal([]byte("{\"request\":{},\"response\":{}}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = V2RewriteHTTP(t)
	return nil
}

// Validate validates this v2 rewrite Http
func (m *V2RewriteHTTP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSticky(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2RewriteHTTP) validatePath(formats strfmt.Registry) error {

	if swag.IsZero(m.Path) { // not required
		return nil
	}

	if m.Path != nil {
		if err := m.Path.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("path")
			}
			return err
		}
	}

	return nil
}

func (m *V2RewriteHTTP) validateRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.Request) { // not required
		return nil
	}

	if m.Request != nil {
		if err := m.Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

func (m *V2RewriteHTTP) validateResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.Response) { // not required
		return nil
	}

	if m.Response != nil {
		if err := m.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

func (m *V2RewriteHTTP) validateSticky(formats strfmt.Registry) error {

	if swag.IsZero(m.Sticky) { // not required
		return nil
	}

	if m.Sticky != nil {
		if err := m.Sticky.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sticky")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2RewriteHTTP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2RewriteHTTP) UnmarshalBinary(b []byte) error {
	var res V2RewriteHTTP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V2RewriteHTTPPath Rewrite the HTTP URL path. Adding a slash to fromPath usually is not necessary.
//
// swagger:model V2RewriteHTTPPath
type V2RewriteHTTPPath struct {

	// from path
	FromPath *string `json:"fromPath,omitempty"`

	// to path
	ToPath *string `json:"toPath,omitempty"`
}

func (m *V2RewriteHTTPPath) UnmarshalJSON(b []byte) error {
	type V2RewriteHTTPPathAlias V2RewriteHTTPPath
	var t V2RewriteHTTPPathAlias
	if err := json.Unmarshal([]byte("{\"fromPath\":\"\",\"toPath\":\"\"}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = V2RewriteHTTPPath(t)
	return nil
}

// Validate validates this v2 rewrite HTTP path
func (m *V2RewriteHTTPPath) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2RewriteHTTPPath) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2RewriteHTTPPath) UnmarshalBinary(b []byte) error {
	var res V2RewriteHTTPPath
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V2RewriteHTTPSticky Sticky sessions via a cookie.
// To use the default values (recommended), set this field to the empty object.
//
// swagger:model V2RewriteHTTPSticky
type V2RewriteHTTPSticky struct {

	// custom str
	CustomStr string `json:"customStr,omitempty"`

	// enabled
	Enabled *bool `json:"enabled,omitempty"`
}

func (m *V2RewriteHTTPSticky) UnmarshalJSON(b []byte) error {
	type V2RewriteHTTPStickyAlias V2RewriteHTTPSticky
	var t V2RewriteHTTPStickyAlias
	if err := json.Unmarshal([]byte("{\"enabled\":true}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = V2RewriteHTTPSticky(t)
	return nil
}

// Validate validates this v2 rewrite HTTP sticky
func (m *V2RewriteHTTPSticky) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2RewriteHTTPSticky) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2RewriteHTTPSticky) UnmarshalBinary(b []byte) error {
	var res V2RewriteHTTPSticky
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}