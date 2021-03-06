// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/iss-lab/edgelb-config-helper/models"
)

// V1DeleteLoadBalancerPoolReader is a Reader for the V1DeleteLoadBalancerPool structure.
type V1DeleteLoadBalancerPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1DeleteLoadBalancerPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1DeleteLoadBalancerPoolNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewV1DeleteLoadBalancerPoolDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV1DeleteLoadBalancerPoolNoContent creates a V1DeleteLoadBalancerPoolNoContent with default headers values
func NewV1DeleteLoadBalancerPoolNoContent() *V1DeleteLoadBalancerPoolNoContent {
	return &V1DeleteLoadBalancerPoolNoContent{}
}

/*V1DeleteLoadBalancerPoolNoContent handles this case with default header values.

Load balancer deleted.
*/
type V1DeleteLoadBalancerPoolNoContent struct {
}

func (o *V1DeleteLoadBalancerPoolNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/loadbalancers/{name}][%d] v1DeleteLoadBalancerPoolNoContent ", 204)
}

func (o *V1DeleteLoadBalancerPoolNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewV1DeleteLoadBalancerPoolDefault creates a V1DeleteLoadBalancerPoolDefault with default headers values
func NewV1DeleteLoadBalancerPoolDefault(code int) *V1DeleteLoadBalancerPoolDefault {
	return &V1DeleteLoadBalancerPoolDefault{
		_statusCode: code,
	}
}

/*V1DeleteLoadBalancerPoolDefault handles this case with default header values.

Unexpected error.
*/
type V1DeleteLoadBalancerPoolDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the v1 delete load balancer pool default response
func (o *V1DeleteLoadBalancerPoolDefault) Code() int {
	return o._statusCode
}

func (o *V1DeleteLoadBalancerPoolDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/loadbalancers/{name}][%d] V1DeleteLoadBalancerPool default  %+v", o._statusCode, o.Payload)
}

func (o *V1DeleteLoadBalancerPoolDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1DeleteLoadBalancerPoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
