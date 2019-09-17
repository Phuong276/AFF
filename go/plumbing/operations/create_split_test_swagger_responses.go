// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/netlify/open-api/go/models"
)

// CreateSplitTestReader is a Reader for the CreateSplitTest structure.
type CreateSplitTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSplitTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateSplitTestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateSplitTestDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSplitTestCreated creates a CreateSplitTestCreated with default headers values
func NewCreateSplitTestCreated() *CreateSplitTestCreated {
	return &CreateSplitTestCreated{}
}

/*CreateSplitTestCreated handles this case with default header values.

Created
*/
type CreateSplitTestCreated struct {
	Payload *models.SplitTest
}

func (o *CreateSplitTestCreated) Error() string {
	return fmt.Sprintf("[POST /site/{site_id}/traffic_splits][%d] createSplitTestCreated  %+v", 201, o.Payload)
}

func (o *CreateSplitTestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SplitTest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSplitTestDefault creates a CreateSplitTestDefault with default headers values
func NewCreateSplitTestDefault(code int) *CreateSplitTestDefault {
	return &CreateSplitTestDefault{
		_statusCode: code,
	}
}

/*CreateSplitTestDefault handles this case with default header values.

error
*/
type CreateSplitTestDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create split test default response
func (o *CreateSplitTestDefault) Code() int {
	return o._statusCode
}

func (o *CreateSplitTestDefault) Error() string {
	return fmt.Sprintf("[POST /site/{site_id}/traffic_splits][%d] createSplitTest default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSplitTestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}