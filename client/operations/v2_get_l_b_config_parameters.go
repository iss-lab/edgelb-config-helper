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

// NewV2GetLBConfigParams creates a new V2GetLBConfigParams object
// with the default values initialized.
func NewV2GetLBConfigParams() *V2GetLBConfigParams {
	var ()
	return &V2GetLBConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2GetLBConfigParamsWithTimeout creates a new V2GetLBConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2GetLBConfigParamsWithTimeout(timeout time.Duration) *V2GetLBConfigParams {
	var ()
	return &V2GetLBConfigParams{

		timeout: timeout,
	}
}

// NewV2GetLBConfigParamsWithContext creates a new V2GetLBConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2GetLBConfigParamsWithContext(ctx context.Context) *V2GetLBConfigParams {
	var ()
	return &V2GetLBConfigParams{

		Context: ctx,
	}
}

// NewV2GetLBConfigParamsWithHTTPClient creates a new V2GetLBConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2GetLBConfigParamsWithHTTPClient(client *http.Client) *V2GetLBConfigParams {
	var ()
	return &V2GetLBConfigParams{
		HTTPClient: client,
	}
}

/*V2GetLBConfigParams contains all the parameters to send to the API endpoint
for the v2 get l b config operation typically these are written to a http.Request
*/
type V2GetLBConfigParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 get l b config params
func (o *V2GetLBConfigParams) WithTimeout(timeout time.Duration) *V2GetLBConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 get l b config params
func (o *V2GetLBConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 get l b config params
func (o *V2GetLBConfigParams) WithContext(ctx context.Context) *V2GetLBConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 get l b config params
func (o *V2GetLBConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 get l b config params
func (o *V2GetLBConfigParams) WithHTTPClient(client *http.Client) *V2GetLBConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 get l b config params
func (o *V2GetLBConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the v2 get l b config params
func (o *V2GetLBConfigParams) WithName(name string) *V2GetLBConfigParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the v2 get l b config params
func (o *V2GetLBConfigParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *V2GetLBConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
