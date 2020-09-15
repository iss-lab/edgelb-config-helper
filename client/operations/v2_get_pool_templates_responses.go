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

// V2GetPoolTemplatesReader is a Reader for the V2GetPoolTemplates structure.
type V2GetPoolTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2GetPoolTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2GetPoolTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewV2GetPoolTemplatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV2GetPoolTemplatesOK creates a V2GetPoolTemplatesOK with default headers values
func NewV2GetPoolTemplatesOK() *V2GetPoolTemplatesOK {
	return &V2GetPoolTemplatesOK{}
}

/*V2GetPoolTemplatesOK handles this case with default header values.

An array of pool templates.
*/
type V2GetPoolTemplatesOK struct {
	Payload []*models.PoolTemplate
}

func (o *V2GetPoolTemplatesOK) Error() string {
	return fmt.Sprintf("[GET /v2/pooltemplates][%d] v2GetPoolTemplatesOK  %+v", 200, o.Payload)
}

func (o *V2GetPoolTemplatesOK) GetPayload() []*models.PoolTemplate {
	return o.Payload
}

func (o *V2GetPoolTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetPoolTemplatesDefault creates a V2GetPoolTemplatesDefault with default headers values
func NewV2GetPoolTemplatesDefault(code int) *V2GetPoolTemplatesDefault {
	return &V2GetPoolTemplatesDefault{
		_statusCode: code,
	}
}

/*V2GetPoolTemplatesDefault handles this case with default header values.

Unexpected error.
*/
type V2GetPoolTemplatesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the v2 get pool templates default response
func (o *V2GetPoolTemplatesDefault) Code() int {
	return o._statusCode
}

func (o *V2GetPoolTemplatesDefault) Error() string {
	return fmt.Sprintf("[GET /v2/pooltemplates][%d] V2GetPoolTemplates default  %+v", o._statusCode, o.Payload)
}

func (o *V2GetPoolTemplatesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2GetPoolTemplatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}