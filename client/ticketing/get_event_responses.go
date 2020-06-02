// Code generated by go-swagger; DO NOT EDIT.

package ticketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// GetEventReader is a Reader for the GetEvent structure.
type GetEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEventBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetEventUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEventForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEventNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetEventTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEventOK creates a GetEventOK with default headers values
func NewGetEventOK() *GetEventOK {
	return &GetEventOK{}
}

/*GetEventOK handles this case with default header values.

successful operation
*/
type GetEventOK struct {
	Payload *models.Event
}

func (o *GetEventOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/events/{eventId}][%d] getEventOK  %+v", 200, o.Payload)
}

func (o *GetEventOK) GetPayload() *models.Event {
	return o.Payload
}

func (o *GetEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Event)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventBadRequest creates a GetEventBadRequest with default headers values
func NewGetEventBadRequest() *GetEventBadRequest {
	return &GetEventBadRequest{}
}

/*GetEventBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetEventBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetEventBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/events/{eventId}][%d] getEventBadRequest  %+v", 400, o.Payload)
}

func (o *GetEventBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetEventBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventUnauthorized creates a GetEventUnauthorized with default headers values
func NewGetEventUnauthorized() *GetEventUnauthorized {
	return &GetEventUnauthorized{}
}

/*GetEventUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetEventUnauthorized struct {
}

func (o *GetEventUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/events/{eventId}][%d] getEventUnauthorized ", 401)
}

func (o *GetEventUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEventForbidden creates a GetEventForbidden with default headers values
func NewGetEventForbidden() *GetEventForbidden {
	return &GetEventForbidden{}
}

/*GetEventForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetEventForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetEventForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/events/{eventId}][%d] getEventForbidden  %+v", 403, o.Payload)
}

func (o *GetEventForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetEventForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventNotFound creates a GetEventNotFound with default headers values
func NewGetEventNotFound() *GetEventNotFound {
	return &GetEventNotFound{}
}

/*GetEventNotFound handles this case with default header values.

Object not found
*/
type GetEventNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetEventNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/events/{eventId}][%d] getEventNotFound  %+v", 404, o.Payload)
}

func (o *GetEventNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetEventNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventTooManyRequests creates a GetEventTooManyRequests with default headers values
func NewGetEventTooManyRequests() *GetEventTooManyRequests {
	return &GetEventTooManyRequests{}
}

/*GetEventTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetEventTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetEventTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/events/{eventId}][%d] getEventTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetEventTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetEventTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}