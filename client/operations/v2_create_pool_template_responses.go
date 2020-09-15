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

// V2CreatePoolTemplateReader is a Reader for the V2CreatePoolTemplate structure.
type V2CreatePoolTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2CreatePoolTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2CreatePoolTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewV2CreatePoolTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV2CreatePoolTemplateOK creates a V2CreatePoolTemplateOK with default headers values
func NewV2CreatePoolTemplateOK() *V2CreatePoolTemplateOK {
	return &V2CreatePoolTemplateOK{}
}

/*V2CreatePoolTemplateOK handles this case with default header values.

Pool template response.
*/
type V2CreatePoolTemplateOK struct {
	Payload *models.PoolTemplate
}

func (o *V2CreatePoolTemplateOK) Error() string {
	return fmt.Sprintf("[POST /v2/pooltemplates][%d] v2CreatePoolTemplateOK  %+v", 200, o.Payload)
}

func (o *V2CreatePoolTemplateOK) GetPayload() *models.PoolTemplate {
	return o.Payload
}

func (o *V2CreatePoolTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PoolTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CreatePoolTemplateDefault creates a V2CreatePoolTemplateDefault with default headers values
func NewV2CreatePoolTemplateDefault(code int) *V2CreatePoolTemplateDefault {
	return &V2CreatePoolTemplateDefault{
		_statusCode: code,
	}
}

/*V2CreatePoolTemplateDefault handles this case with default header values.

Unexpected error.
*/
type V2CreatePoolTemplateDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the v2 create pool template default response
func (o *V2CreatePoolTemplateDefault) Code() int {
	return o._statusCode
}

func (o *V2CreatePoolTemplateDefault) Error() string {
	return fmt.Sprintf("[POST /v2/pooltemplates][%d] V2CreatePoolTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *V2CreatePoolTemplateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2CreatePoolTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}