// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewV2DeletePoolTemplateConfigTemplateParams creates a new V2DeletePoolTemplateConfigTemplateParams object
// with the default values initialized.
func NewV2DeletePoolTemplateConfigTemplateParams() *V2DeletePoolTemplateConfigTemplateParams {
	var ()
	return &V2DeletePoolTemplateConfigTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2DeletePoolTemplateConfigTemplateParamsWithTimeout creates a new V2DeletePoolTemplateConfigTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2DeletePoolTemplateConfigTemplateParamsWithTimeout(timeout time.Duration) *V2DeletePoolTemplateConfigTemplateParams {
	var ()
	return &V2DeletePoolTemplateConfigTemplateParams{

		timeout: timeout,
	}
}

// NewV2DeletePoolTemplateConfigTemplateParamsWithContext creates a new V2DeletePoolTemplateConfigTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2DeletePoolTemplateConfigTemplateParamsWithContext(ctx context.Context) *V2DeletePoolTemplateConfigTemplateParams {
	var ()
	return &V2DeletePoolTemplateConfigTemplateParams{

		Context: ctx,
	}
}

// NewV2DeletePoolTemplateConfigTemplateParamsWithHTTPClient creates a new V2DeletePoolTemplateConfigTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2DeletePoolTemplateConfigTemplateParamsWithHTTPClient(client *http.Client) *V2DeletePoolTemplateConfigTemplateParams {
	var ()
	return &V2DeletePoolTemplateConfigTemplateParams{
		HTTPClient: client,
	}
}

/*V2DeletePoolTemplateConfigTemplateParams contains all the parameters to send to the API endpoint
for the v2 delete pool template config template operation typically these are written to a http.Request
*/
type V2DeletePoolTemplateConfigTemplateParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 delete pool template config template params
func (o *V2DeletePoolTemplateConfigTemplateParams) WithTimeout(timeout time.Duration) *V2DeletePoolTemplateConfigTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 delete pool template config template params
func (o *V2DeletePoolTemplateConfigTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 delete pool template config template params
func (o *V2DeletePoolTemplateConfigTemplateParams) WithContext(ctx context.Context) *V2DeletePoolTemplateConfigTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 delete pool template config template params
func (o *V2DeletePoolTemplateConfigTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 delete pool template config template params
func (o *V2DeletePoolTemplateConfigTemplateParams) WithHTTPClient(client *http.Client) *V2DeletePoolTemplateConfigTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 delete pool template config template params
func (o *V2DeletePoolTemplateConfigTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the v2 delete pool template config template params
func (o *V2DeletePoolTemplateConfigTemplateParams) WithName(name string) *V2DeletePoolTemplateConfigTemplateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the v2 delete pool template config template params
func (o *V2DeletePoolTemplateConfigTemplateParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *V2DeletePoolTemplateConfigTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}