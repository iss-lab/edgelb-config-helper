// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// RenderStatus render status
//
// swagger:model RenderStatus
type RenderStatus string

const (

	// RenderStatusOK captures enum value "OK"
	RenderStatusOK RenderStatus = "OK"

	// RenderStatusRENDERERR captures enum value "RENDER_ERR"
	RenderStatusRENDERERR RenderStatus = "RENDER_ERR"

	// RenderStatusSUBMITERR captures enum value "SUBMIT_ERR"
	RenderStatusSUBMITERR RenderStatus = "SUBMIT_ERR"

	// RenderStatusNOTASKS captures enum value "NO_TASKS"
	RenderStatusNOTASKS RenderStatus = "NO_TASKS"

	// RenderStatusNEW captures enum value "NEW"
	RenderStatusNEW RenderStatus = "NEW"
)

// for schema
var renderStatusEnum []interface{}

func init() {
	var res []RenderStatus
	if err := json.Unmarshal([]byte(`["OK","RENDER_ERR","SUBMIT_ERR","NO_TASKS","NEW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		renderStatusEnum = append(renderStatusEnum, v)
	}
}

func (m RenderStatus) validateRenderStatusEnum(path, location string, value RenderStatus) error {
	if err := validate.EnumCase(path, location, value, renderStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this render status
func (m RenderStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRenderStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
