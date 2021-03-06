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

// NewV2DeletePoolTemplateParams creates a new V2DeletePoolTemplateParams object
// with the default values initialized.
func NewV2DeletePoolTemplateParams() *V2DeletePoolTemplateParams {
	var ()
	return &V2DeletePoolTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2DeletePoolTemplateParamsWithTimeout creates a new V2DeletePoolTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2DeletePoolTemplateParamsWithTimeout(timeout time.Duration) *V2DeletePoolTemplateParams {
	var ()
	return &V2DeletePoolTemplateParams{

		timeout: timeout,
	}
}

// NewV2DeletePoolTemplateParamsWithContext creates a new V2DeletePoolTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2DeletePoolTemplateParamsWithContext(ctx context.Context) *V2DeletePoolTemplateParams {
	var ()
	return &V2DeletePoolTemplateParams{

		Context: ctx,
	}
}

// NewV2DeletePoolTemplateParamsWithHTTPClient creates a new V2DeletePoolTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2DeletePoolTemplateParamsWithHTTPClient(client *http.Client) *V2DeletePoolTemplateParams {
	var ()
	return &V2DeletePoolTemplateParams{
		HTTPClient: client,
	}
}

/*V2DeletePoolTemplateParams contains all the parameters to send to the API endpoint
for the v2 delete pool template operation typically these are written to a http.Request
*/
type V2DeletePoolTemplateParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 delete pool template params
func (o *V2DeletePoolTemplateParams) WithTimeout(timeout time.Duration) *V2DeletePoolTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 delete pool template params
func (o *V2DeletePoolTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 delete pool template params
func (o *V2DeletePoolTemplateParams) WithContext(ctx context.Context) *V2DeletePoolTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 delete pool template params
func (o *V2DeletePoolTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 delete pool template params
func (o *V2DeletePoolTemplateParams) WithHTTPClient(client *http.Client) *V2DeletePoolTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 delete pool template params
func (o *V2DeletePoolTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the v2 delete pool template params
func (o *V2DeletePoolTemplateParams) WithName(name string) *V2DeletePoolTemplateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the v2 delete pool template params
func (o *V2DeletePoolTemplateParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *V2DeletePoolTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
