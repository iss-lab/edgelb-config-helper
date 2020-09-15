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

// NewV2GetPoolTemplateConfigTemplateParams creates a new V2GetPoolTemplateConfigTemplateParams object
// with the default values initialized.
func NewV2GetPoolTemplateConfigTemplateParams() *V2GetPoolTemplateConfigTemplateParams {
	var ()
	return &V2GetPoolTemplateConfigTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2GetPoolTemplateConfigTemplateParamsWithTimeout creates a new V2GetPoolTemplateConfigTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2GetPoolTemplateConfigTemplateParamsWithTimeout(timeout time.Duration) *V2GetPoolTemplateConfigTemplateParams {
	var ()
	return &V2GetPoolTemplateConfigTemplateParams{

		timeout: timeout,
	}
}

// NewV2GetPoolTemplateConfigTemplateParamsWithContext creates a new V2GetPoolTemplateConfigTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2GetPoolTemplateConfigTemplateParamsWithContext(ctx context.Context) *V2GetPoolTemplateConfigTemplateParams {
	var ()
	return &V2GetPoolTemplateConfigTemplateParams{

		Context: ctx,
	}
}

// NewV2GetPoolTemplateConfigTemplateParamsWithHTTPClient creates a new V2GetPoolTemplateConfigTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2GetPoolTemplateConfigTemplateParamsWithHTTPClient(client *http.Client) *V2GetPoolTemplateConfigTemplateParams {
	var ()
	return &V2GetPoolTemplateConfigTemplateParams{
		HTTPClient: client,
	}
}

/*V2GetPoolTemplateConfigTemplateParams contains all the parameters to send to the API endpoint
for the v2 get pool template config template operation typically these are written to a http.Request
*/
type V2GetPoolTemplateConfigTemplateParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 get pool template config template params
func (o *V2GetPoolTemplateConfigTemplateParams) WithTimeout(timeout time.Duration) *V2GetPoolTemplateConfigTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 get pool template config template params
func (o *V2GetPoolTemplateConfigTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 get pool template config template params
func (o *V2GetPoolTemplateConfigTemplateParams) WithContext(ctx context.Context) *V2GetPoolTemplateConfigTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 get pool template config template params
func (o *V2GetPoolTemplateConfigTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 get pool template config template params
func (o *V2GetPoolTemplateConfigTemplateParams) WithHTTPClient(client *http.Client) *V2GetPoolTemplateConfigTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 get pool template config template params
func (o *V2GetPoolTemplateConfigTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the v2 get pool template config template params
func (o *V2GetPoolTemplateConfigTemplateParams) WithName(name string) *V2GetPoolTemplateConfigTemplateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the v2 get pool template config template params
func (o *V2GetPoolTemplateConfigTemplateParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *V2GetPoolTemplateConfigTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
