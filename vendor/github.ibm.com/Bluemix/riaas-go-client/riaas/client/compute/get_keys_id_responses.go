// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetKeysIDReader is a Reader for the GetKeysID structure.
type GetKeysIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKeysIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetKeysIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetKeysIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetKeysIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKeysIDOK creates a GetKeysIDOK with default headers values
func NewGetKeysIDOK() *GetKeysIDOK {
	return &GetKeysIDOK{}
}

/*GetKeysIDOK handles this case with default header values.

dummy
*/
type GetKeysIDOK struct {
	Payload *models.Key
}

func (o *GetKeysIDOK) Error() string {
	return fmt.Sprintf("[GET /keys/{id}][%d] getKeysIdOK  %+v", 200, o.Payload)
}

func (o *GetKeysIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Key)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeysIDNotFound creates a GetKeysIDNotFound with default headers values
func NewGetKeysIDNotFound() *GetKeysIDNotFound {
	return &GetKeysIDNotFound{}
}

/*GetKeysIDNotFound handles this case with default header values.

error
*/
type GetKeysIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetKeysIDNotFound) Error() string {
	return fmt.Sprintf("[GET /keys/{id}][%d] getKeysIdNotFound  %+v", 404, o.Payload)
}

func (o *GetKeysIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeysIDInternalServerError creates a GetKeysIDInternalServerError with default headers values
func NewGetKeysIDInternalServerError() *GetKeysIDInternalServerError {
	return &GetKeysIDInternalServerError{}
}

/*GetKeysIDInternalServerError handles this case with default header values.

error
*/
type GetKeysIDInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *GetKeysIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /keys/{id}][%d] getKeysIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKeysIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}