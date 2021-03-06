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

// NewV1DeleteLoadBalancerPoolParams creates a new V1DeleteLoadBalancerPoolParams object
// with the default values initialized.
func NewV1DeleteLoadBalancerPoolParams() *V1DeleteLoadBalancerPoolParams {
	var ()
	return &V1DeleteLoadBalancerPoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1DeleteLoadBalancerPoolParamsWithTimeout creates a new V1DeleteLoadBalancerPoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1DeleteLoadBalancerPoolParamsWithTimeout(timeout time.Duration) *V1DeleteLoadBalancerPoolParams {
	var ()
	return &V1DeleteLoadBalancerPoolParams{

		timeout: timeout,
	}
}

// NewV1DeleteLoadBalancerPoolParamsWithContext creates a new V1DeleteLoadBalancerPoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1DeleteLoadBalancerPoolParamsWithContext(ctx context.Context) *V1DeleteLoadBalancerPoolParams {
	var ()
	return &V1DeleteLoadBalancerPoolParams{

		Context: ctx,
	}
}

// NewV1DeleteLoadBalancerPoolParamsWithHTTPClient creates a new V1DeleteLoadBalancerPoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1DeleteLoadBalancerPoolParamsWithHTTPClient(client *http.Client) *V1DeleteLoadBalancerPoolParams {
	var ()
	return &V1DeleteLoadBalancerPoolParams{
		HTTPClient: client,
	}
}

/*V1DeleteLoadBalancerPoolParams contains all the parameters to send to the API endpoint
for the v1 delete load balancer pool operation typically these are written to a http.Request
*/
type V1DeleteLoadBalancerPoolParams struct {

	/*Name*/
	Name string
	/*Token
	  DCOS Auth Token

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 delete load balancer pool params
func (o *V1DeleteLoadBalancerPoolParams) WithTimeout(timeout time.Duration) *V1DeleteLoadBalancerPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 delete load balancer pool params
func (o *V1DeleteLoadBalancerPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 delete load balancer pool params
func (o *V1DeleteLoadBalancerPoolParams) WithContext(ctx context.Context) *V1DeleteLoadBalancerPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 delete load balancer pool params
func (o *V1DeleteLoadBalancerPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 delete load balancer pool params
func (o *V1DeleteLoadBalancerPoolParams) WithHTTPClient(client *http.Client) *V1DeleteLoadBalancerPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 delete load balancer pool params
func (o *V1DeleteLoadBalancerPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the v1 delete load balancer pool params
func (o *V1DeleteLoadBalancerPoolParams) WithName(name string) *V1DeleteLoadBalancerPoolParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the v1 delete load balancer pool params
func (o *V1DeleteLoadBalancerPoolParams) SetName(name string) {
	o.Name = name
}

// WithToken adds the token to the v1 delete load balancer pool params
func (o *V1DeleteLoadBalancerPoolParams) WithToken(token *string) *V1DeleteLoadBalancerPoolParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the v1 delete load balancer pool params
func (o *V1DeleteLoadBalancerPoolParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *V1DeleteLoadBalancerPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Token != nil {

		// header param token
		if err := r.SetHeaderParam("token", *o.Token); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
