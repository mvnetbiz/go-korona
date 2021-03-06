// Code generated by go-swagger; DO NOT EDIT.

package delivery_notes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// AddDeliveryNotesReader is a Reader for the AddDeliveryNotes structure.
type AddDeliveryNotesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDeliveryNotesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddDeliveryNotesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddDeliveryNotesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddDeliveryNotesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddDeliveryNotesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddDeliveryNotesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAddDeliveryNotesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddDeliveryNotesOK creates a AddDeliveryNotesOK with default headers values
func NewAddDeliveryNotesOK() *AddDeliveryNotesOK {
	return &AddDeliveryNotesOK{}
}

/*AddDeliveryNotesOK handles this case with default header values.

successful operation
*/
type AddDeliveryNotesOK struct {
	Payload []*models.AddOrUpdateResult
}

func (o *AddDeliveryNotesOK) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/deliveryNotes][%d] addDeliveryNotesOK  %+v", 200, o.Payload)
}

func (o *AddDeliveryNotesOK) GetPayload() []*models.AddOrUpdateResult {
	return o.Payload
}

func (o *AddDeliveryNotesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDeliveryNotesBadRequest creates a AddDeliveryNotesBadRequest with default headers values
func NewAddDeliveryNotesBadRequest() *AddDeliveryNotesBadRequest {
	return &AddDeliveryNotesBadRequest{}
}

/*AddDeliveryNotesBadRequest handles this case with default header values.

Malformed querystring or model
*/
type AddDeliveryNotesBadRequest struct {
	Payload *models.BadRequestError
}

func (o *AddDeliveryNotesBadRequest) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/deliveryNotes][%d] addDeliveryNotesBadRequest  %+v", 400, o.Payload)
}

func (o *AddDeliveryNotesBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *AddDeliveryNotesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDeliveryNotesUnauthorized creates a AddDeliveryNotesUnauthorized with default headers values
func NewAddDeliveryNotesUnauthorized() *AddDeliveryNotesUnauthorized {
	return &AddDeliveryNotesUnauthorized{}
}

/*AddDeliveryNotesUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type AddDeliveryNotesUnauthorized struct {
}

func (o *AddDeliveryNotesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/deliveryNotes][%d] addDeliveryNotesUnauthorized ", 401)
}

func (o *AddDeliveryNotesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddDeliveryNotesForbidden creates a AddDeliveryNotesForbidden with default headers values
func NewAddDeliveryNotesForbidden() *AddDeliveryNotesForbidden {
	return &AddDeliveryNotesForbidden{}
}

/*AddDeliveryNotesForbidden handles this case with default header values.

Requested action is not allowed
*/
type AddDeliveryNotesForbidden struct {
	Payload *models.ForbiddenError
}

func (o *AddDeliveryNotesForbidden) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/deliveryNotes][%d] addDeliveryNotesForbidden  %+v", 403, o.Payload)
}

func (o *AddDeliveryNotesForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *AddDeliveryNotesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDeliveryNotesNotFound creates a AddDeliveryNotesNotFound with default headers values
func NewAddDeliveryNotesNotFound() *AddDeliveryNotesNotFound {
	return &AddDeliveryNotesNotFound{}
}

/*AddDeliveryNotesNotFound handles this case with default header values.

Object not found
*/
type AddDeliveryNotesNotFound struct {
	Payload *models.NotFoundError
}

func (o *AddDeliveryNotesNotFound) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/deliveryNotes][%d] addDeliveryNotesNotFound  %+v", 404, o.Payload)
}

func (o *AddDeliveryNotesNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *AddDeliveryNotesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDeliveryNotesTooManyRequests creates a AddDeliveryNotesTooManyRequests with default headers values
func NewAddDeliveryNotesTooManyRequests() *AddDeliveryNotesTooManyRequests {
	return &AddDeliveryNotesTooManyRequests{}
}

/*AddDeliveryNotesTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type AddDeliveryNotesTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *AddDeliveryNotesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/deliveryNotes][%d] addDeliveryNotesTooManyRequests  %+v", 429, o.Payload)
}

func (o *AddDeliveryNotesTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *AddDeliveryNotesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
