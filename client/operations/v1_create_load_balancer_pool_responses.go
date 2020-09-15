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

// V1CreateLoadBalancerPoolReader is a Reader for the V1CreateLoadBalancerPool structure.
type V1CreateLoadBalancerPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CreateLoadBalancerPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CreateLoadBalancerPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewV1CreateLoadBalancerPoolDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV1CreateLoadBalancerPoolOK creates a V1CreateLoadBalancerPoolOK with default headers values
func NewV1CreateLoadBalancerPoolOK() *V1CreateLoadBalancerPoolOK {
	return &V1CreateLoadBalancerPoolOK{}
}

/*V1CreateLoadBalancerPoolOK handles this case with default header values.

Load Balancer Pool response.
*/
type V1CreateLoadBalancerPoolOK struct {
	Payload *models.V1Pool
}

func (o *V1CreateLoadBalancerPoolOK) Error() string {
	return fmt.Sprintf("[POST /v1/loadbalancers][%d] v1CreateLoadBalancerPoolOK  %+v", 200, o.Payload)
}

func (o *V1CreateLoadBalancerPoolOK) GetPayload() *models.V1Pool {
	return o.Payload
}

func (o *V1CreateLoadBalancerPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Pool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1CreateLoadBalancerPoolDefault creates a V1CreateLoadBalancerPoolDefault with default headers values
func NewV1CreateLoadBalancerPoolDefault(code int) *V1CreateLoadBalancerPoolDefault {
	return &V1CreateLoadBalancerPoolDefault{
		_statusCode: code,
	}
}

/*V1CreateLoadBalancerPoolDefault handles this case with default header values.

Unexpected error.
*/
type V1CreateLoadBalancerPoolDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the v1 create load balancer pool default response
func (o *V1CreateLoadBalancerPoolDefault) Code() int {
	return o._statusCode
}

func (o *V1CreateLoadBalancerPoolDefault) Error() string {
	return fmt.Sprintf("[POST /v1/loadbalancers][%d] V1CreateLoadBalancerPool default  %+v", o._statusCode, o.Payload)
}

func (o *V1CreateLoadBalancerPoolDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1CreateLoadBalancerPoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
