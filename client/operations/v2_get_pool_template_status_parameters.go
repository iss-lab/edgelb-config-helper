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

// NewV2GetPoolTemplateStatusParams creates a new V2GetPoolTemplateStatusParams object
// with the default values initialized.
func NewV2GetPoolTemplateStatusParams() *V2GetPoolTemplateStatusParams {
	var ()
	return &V2GetPoolTemplateStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2GetPoolTemplateStatusParamsWithTimeout creates a new V2GetPoolTemplateStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2GetPoolTemplateStatusParamsWithTimeout(timeout time.Duration) *V2GetPoolTemplateStatusParams {
	var ()
	return &V2GetPoolTemplateStatusParams{

		timeout: timeout,
	}
}

// NewV2GetPoolTemplateStatusParamsWithContext creates a new V2GetPoolTemplateStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2GetPoolTemplateStatusParamsWithContext(ctx context.Context) *V2GetPoolTemplateStatusParams {
	var ()
	return &V2GetPoolTemplateStatusParams{

		Context: ctx,
	}
}

// NewV2GetPoolTemplateStatusParamsWithHTTPClient creates a new V2GetPoolTemplateStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2GetPoolTemplateStatusParamsWithHTTPClient(client *http.Client) *V2GetPoolTemplateStatusParams {
	var ()
	return &V2GetPoolTemplateStatusParams{
		HTTPClient: client,
	}
}

/*V2GetPoolTemplateStatusParams contains all the parameters to send to the API endpoint
for the v2 get pool template status operation typically these are written to a http.Request
*/
type V2GetPoolTemplateStatusParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 get pool template status params
func (o *V2GetPoolTemplateStatusParams) WithTimeout(timeout time.Duration) *V2GetPoolTemplateStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 get pool template status params
func (o *V2GetPoolTemplateStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 get pool template status params
func (o *V2GetPoolTemplateStatusParams) WithContext(ctx context.Context) *V2GetPoolTemplateStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 get pool template status params
func (o *V2GetPoolTemplateStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 get pool template status params
func (o *V2GetPoolTemplateStatusParams) WithHTTPClient(client *http.Client) *V2GetPoolTemplateStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 get pool template status params
func (o *V2GetPoolTemplateStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the v2 get pool template status params
func (o *V2GetPoolTemplateStatusParams) WithName(name string) *V2GetPoolTemplateStatusParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the v2 get pool template status params
func (o *V2GetPoolTemplateStatusParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *V2GetPoolTemplateStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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