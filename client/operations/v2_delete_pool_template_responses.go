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

// V2DeletePoolTemplateReader is a Reader for the V2DeletePoolTemplate structure.
type V2DeletePoolTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2DeletePoolTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV2DeletePoolTemplateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewV2DeletePoolTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewV2DeletePoolTemplateNoContent creates a V2DeletePoolTemplateNoContent with default headers values
func NewV2DeletePoolTemplateNoContent() *V2DeletePoolTemplateNoContent {
	return &V2DeletePoolTemplateNoContent{}
}

/*V2DeletePoolTemplateNoContent handles this case with default header values.

Pool template deleted.
*/
type V2DeletePoolTemplateNoContent struct {
}

func (o *V2DeletePoolTemplateNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v2/pooltemplates/{name}][%d] v2DeletePoolTemplateNoContent ", 204)
}

func (o *V2DeletePoolTemplateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewV2DeletePoolTemplateDefault creates a V2DeletePoolTemplateDefault with default headers values
func NewV2DeletePoolTemplateDefault(code int) *V2DeletePoolTemplateDefault {
	return &V2DeletePoolTemplateDefault{
		_statusCode: code,
	}
}

/*V2DeletePoolTemplateDefault handles this case with default header values.

Unexpected error.
*/
type V2DeletePoolTemplateDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the v2 delete pool template default response
func (o *V2DeletePoolTemplateDefault) Code() int {
	return o._statusCode
}

func (o *V2DeletePoolTemplateDefault) Error() string {
	return fmt.Sprintf("[DELETE /v2/pooltemplates/{name}][%d] V2DeletePoolTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *V2DeletePoolTemplateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2DeletePoolTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
