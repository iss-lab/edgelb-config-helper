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

// NewV2GetPoolMetadataParams creates a new V2GetPoolMetadataParams object
// with the default values initialized.
func NewV2GetPoolMetadataParams() *V2GetPoolMetadataParams {
	var ()
	return &V2GetPoolMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2GetPoolMetadataParamsWithTimeout creates a new V2GetPoolMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2GetPoolMetadataParamsWithTimeout(timeout time.Duration) *V2GetPoolMetadataParams {
	var ()
	return &V2GetPoolMetadataParams{

		timeout: timeout,
	}
}

// NewV2GetPoolMetadataParamsWithContext creates a new V2GetPoolMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2GetPoolMetadataParamsWithContext(ctx context.Context) *V2GetPoolMetadataParams {
	var ()
	return &V2GetPoolMetadataParams{

		Context: ctx,
	}
}

// NewV2GetPoolMetadataParamsWithHTTPClient creates a new V2GetPoolMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2GetPoolMetadataParamsWithHTTPClient(client *http.Client) *V2GetPoolMetadataParams {
	var ()
	return &V2GetPoolMetadataParams{
		HTTPClient: client,
	}
}

/*V2GetPoolMetadataParams contains all the parameters to send to the API endpoint
for the v2 get pool metadata operation typically these are written to a http.Request
*/
type V2GetPoolMetadataParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 get pool metadata params
func (o *V2GetPoolMetadataParams) WithTimeout(timeout time.Duration) *V2GetPoolMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 get pool metadata params
func (o *V2GetPoolMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 get pool metadata params
func (o *V2GetPoolMetadataParams) WithContext(ctx context.Context) *V2GetPoolMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 get pool metadata params
func (o *V2GetPoolMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 get pool metadata params
func (o *V2GetPoolMetadataParams) WithHTTPClient(client *http.Client) *V2GetPoolMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 get pool metadata params
func (o *V2GetPoolMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the v2 get pool metadata params
func (o *V2GetPoolMetadataParams) WithName(name string) *V2GetPoolMetadataParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the v2 get pool metadata params
func (o *V2GetPoolMetadataParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *V2GetPoolMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
