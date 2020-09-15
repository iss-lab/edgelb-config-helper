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

// V2GetPoolsReader is a Reader for the V2GetPools structure.
type V2GetPoolsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2GetPoolsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2GetPoolsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewV2GetPoolsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV2GetPoolsOK creates a V2GetPoolsOK with default headers values
func NewV2GetPoolsOK() *V2GetPoolsOK {
	return &V2GetPoolsOK{}
}

/*V2GetPoolsOK handles this case with default header values.

An array of load balancer pools.
*/
type V2GetPoolsOK struct {
	Payload []*models.V2Pool
}

func (o *V2GetPoolsOK) Error() string {
	return fmt.Sprintf("[GET /v2/pools][%d] v2GetPoolsOK  %+v", 200, o.Payload)
}

func (o *V2GetPoolsOK) GetPayload() []*models.V2Pool {
	return o.Payload
}

func (o *V2GetPoolsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetPoolsDefault creates a V2GetPoolsDefault with default headers values
func NewV2GetPoolsDefault(code int) *V2GetPoolsDefault {
	return &V2GetPoolsDefault{
		_statusCode: code,
	}
}

/*V2GetPoolsDefault handles this case with default header values.

Unexpected error.
*/
type V2GetPoolsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the v2 get pools default response
func (o *V2GetPoolsDefault) Code() int {
	return o._statusCode
}

func (o *V2GetPoolsDefault) Error() string {
	return fmt.Sprintf("[GET /v2/pools][%d] V2GetPools default  %+v", o._statusCode, o.Payload)
}

func (o *V2GetPoolsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2GetPoolsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}