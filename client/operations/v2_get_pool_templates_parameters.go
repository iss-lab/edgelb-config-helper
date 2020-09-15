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

// NewV2GetPoolTemplatesParams creates a new V2GetPoolTemplatesParams object
// with the default values initialized.
func NewV2GetPoolTemplatesParams() *V2GetPoolTemplatesParams {

	return &V2GetPoolTemplatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2GetPoolTemplatesParamsWithTimeout creates a new V2GetPoolTemplatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2GetPoolTemplatesParamsWithTimeout(timeout time.Duration) *V2GetPoolTemplatesParams {

	return &V2GetPoolTemplatesParams{

		timeout: timeout,
	}
}

// NewV2GetPoolTemplatesParamsWithContext creates a new V2GetPoolTemplatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2GetPoolTemplatesParamsWithContext(ctx context.Context) *V2GetPoolTemplatesParams {

	return &V2GetPoolTemplatesParams{

		Context: ctx,
	}
}

// NewV2GetPoolTemplatesParamsWithHTTPClient creates a new V2GetPoolTemplatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2GetPoolTemplatesParamsWithHTTPClient(client *http.Client) *V2GetPoolTemplatesParams {

	return &V2GetPoolTemplatesParams{
		HTTPClient: client,
	}
}

/*V2GetPoolTemplatesParams contains all the parameters to send to the API endpoint
for the v2 get pool templates operation typically these are written to a http.Request
*/
type V2GetPoolTemplatesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 get pool templates params
func (o *V2GetPoolTemplatesParams) WithTimeout(timeout time.Duration) *V2GetPoolTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 get pool templates params
func (o *V2GetPoolTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 get pool templates params
func (o *V2GetPoolTemplatesParams) WithContext(ctx context.Context) *V2GetPoolTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 get pool templates params
func (o *V2GetPoolTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 get pool templates params
func (o *V2GetPoolTemplatesParams) WithHTTPClient(client *http.Client) *V2GetPoolTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 get pool templates params
func (o *V2GetPoolTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V2GetPoolTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
