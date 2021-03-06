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

// GetEventsReader is a Reader for the GetEvents structure.
type GetEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetEventsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEventsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEventsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetEventsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEventsOK creates a GetEventsOK with default headers values
func NewGetEventsOK() *GetEventsOK {
	return &GetEventsOK{}
}

/*GetEventsOK handles this case with default header values.

successful operation
*/
type GetEventsOK struct {
	Payload *models.ResultListEvent
}

func (o *GetEventsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/events][%d] getEventsOK  %+v", 200, o.Payload)
}

func (o *GetEventsOK) GetPayload() *models.ResultListEvent {
	return o.Payload
}

func (o *GetEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResultListEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsNoContent creates a GetEventsNoContent with default headers values
func NewGetEventsNoContent() *GetEventsNoContent {
	return &GetEventsNoContent{}
}

/*GetEventsNoContent handles this case with default header values.

Request successful, but the list is empty. Either there is in general no object on the list or a set filter/restriction in querystring doesn't match any object.
*/
type GetEventsNoContent struct {
}

func (o *GetEventsNoContent) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/events][%d] getEventsNoContent ", 204)
}

func (o *GetEventsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEventsBadRequest creates a GetEventsBadRequest with default headers values
func NewGetEventsBadRequest() *GetEventsBadRequest {
	return &GetEventsBadRequest{}
}

/*GetEventsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetEventsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetEventsBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/events][%d] getEventsBadRequest  %+v", 400, o.Payload)
}

func (o *GetEventsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetEventsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsUnauthorized creates a GetEventsUnauthorized with default headers values
func NewGetEventsUnauthorized() *GetEventsUnauthorized {
	return &GetEventsUnauthorized{}
}

/*GetEventsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetEventsUnauthorized struct {
}

func (o *GetEventsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/events][%d] getEventsUnauthorized ", 401)
}

func (o *GetEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEventsForbidden creates a GetEventsForbidden with default headers values
func NewGetEventsForbidden() *GetEventsForbidden {
	return &GetEventsForbidden{}
}

/*GetEventsForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetEventsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetEventsForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/events][%d] getEventsForbidden  %+v", 403, o.Payload)
}

func (o *GetEventsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsNotFound creates a GetEventsNotFound with default headers values
func NewGetEventsNotFound() *GetEventsNotFound {
	return &GetEventsNotFound{}
}

/*GetEventsNotFound handles this case with default header values.

Object not found
*/
type GetEventsNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetEventsNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/events][%d] getEventsNotFound  %+v", 404, o.Payload)
}

func (o *GetEventsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetEventsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsTooManyRequests creates a GetEventsTooManyRequests with default headers values
func NewGetEventsTooManyRequests() *GetEventsTooManyRequests {
	return &GetEventsTooManyRequests{}
}

/*GetEventsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetEventsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetEventsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/events][%d] getEventsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetEventsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetEventsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
