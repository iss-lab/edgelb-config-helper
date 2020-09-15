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

// V1GetConfigReader is a Reader for the V1GetConfig structure.
type V1GetConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1GetConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1GetConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewV1GetConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV1GetConfigOK creates a V1GetConfigOK with default headers values
func NewV1GetConfigOK() *V1GetConfigOK {
	return &V1GetConfigOK{}
}

/*V1GetConfigOK handles this case with default header values.

A configuration object containing all load balancer pools.
*/
type V1GetConfigOK struct {
	Payload *models.V1Config
}

func (o *V1GetConfigOK) Error() string {
	return fmt.Sprintf("[GET /v1/all][%d] v1GetConfigOK  %+v", 200, o.Payload)
}

func (o *V1GetConfigOK) GetPayload() *models.V1Config {
	return o.Payload
}

func (o *V1GetConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Config)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1GetConfigDefault creates a V1GetConfigDefault with default headers values
func NewV1GetConfigDefault(code int) *V1GetConfigDefault {
	return &V1GetConfigDefault{
		_statusCode: code,
	}
}

/*V1GetConfigDefault handles this case with default header values.

Unexpected error.
*/
type V1GetConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the v1 get config default response
func (o *V1GetConfigDefault) Code() int {
	return o._statusCode
}

func (o *V1GetConfigDefault) Error() string {
	return fmt.Sprintf("[GET /v1/all][%d] V1GetConfig default  %+v", o._statusCode, o.Payload)
}

func (o *V1GetConfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1GetConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
