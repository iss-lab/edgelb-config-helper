// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1GetLoadBalancerArtifactReader is a Reader for the V1GetLoadBalancerArtifact structure.
type V1GetLoadBalancerArtifactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1GetLoadBalancerArtifactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1GetLoadBalancerArtifactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewV1GetLoadBalancerArtifactDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV1GetLoadBalancerArtifactOK creates a V1GetLoadBalancerArtifactOK with default headers values
func NewV1GetLoadBalancerArtifactOK() *V1GetLoadBalancerArtifactOK {
	return &V1GetLoadBalancerArtifactOK{}
}

/*V1GetLoadBalancerArtifactOK handles this case with default header values.

Configuration artifact for load balancer pool.
*/
type V1GetLoadBalancerArtifactOK struct {
	Payload string
}

func (o *V1GetLoadBalancerArtifactOK) Error() string {
	return fmt.Sprintf("[GET /v1/loadbalancers/{name}/artifacts/{artifactName}][%d] v1GetLoadBalancerArtifactOK  %+v", 200, o.Payload)
}

func (o *V1GetLoadBalancerArtifactOK) GetPayload() string {
	return o.Payload
}

func (o *V1GetLoadBalancerArtifactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1GetLoadBalancerArtifactDefault creates a V1GetLoadBalancerArtifactDefault with default headers values
func NewV1GetLoadBalancerArtifactDefault(code int) *V1GetLoadBalancerArtifactDefault {
	return &V1GetLoadBalancerArtifactDefault{
		_statusCode: code,
	}
}

/*V1GetLoadBalancerArtifactDefault handles this case with default header values.

Unexpected error.
*/
type V1GetLoadBalancerArtifactDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the v1 get load balancer artifact default response
func (o *V1GetLoadBalancerArtifactDefault) Code() int {
	return o._statusCode
}

func (o *V1GetLoadBalancerArtifactDefault) Error() string {
	return fmt.Sprintf("[GET /v1/loadbalancers/{name}/artifacts/{artifactName}][%d] V1GetLoadBalancerArtifact default  %+v", o._statusCode, o.Payload)
}

func (o *V1GetLoadBalancerArtifactDefault) GetPayload() string {
	return o.Payload
}

func (o *V1GetLoadBalancerArtifactDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}