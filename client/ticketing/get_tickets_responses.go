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

// GetTicketsReader is a Reader for the GetTickets structure.
type GetTicketsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTicketsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTicketsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetTicketsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTicketsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTicketsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTicketsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTicketsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTicketsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTicketsOK creates a GetTicketsOK with default headers values
func NewGetTicketsOK() *GetTicketsOK {
	return &GetTicketsOK{}
}

/*GetTicketsOK handles this case with default header values.

successful operation
*/
type GetTicketsOK struct {
	Payload *models.ResultListTicket
}

func (o *GetTicketsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/tickets][%d] getTicketsOK  %+v", 200, o.Payload)
}

func (o *GetTicketsOK) GetPayload() *models.ResultListTicket {
	return o.Payload
}

func (o *GetTicketsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResultListTicket)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTicketsNoContent creates a GetTicketsNoContent with default headers values
func NewGetTicketsNoContent() *GetTicketsNoContent {
	return &GetTicketsNoContent{}
}

/*GetTicketsNoContent handles this case with default header values.

Request successful, but the list is empty. Either there is in general no object on the list or a set filter/restriction in querystring doesn't match any object.
*/
type GetTicketsNoContent struct {
}

func (o *GetTicketsNoContent) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/tickets][%d] getTicketsNoContent ", 204)
}

func (o *GetTicketsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTicketsBadRequest creates a GetTicketsBadRequest with default headers values
func NewGetTicketsBadRequest() *GetTicketsBadRequest {
	return &GetTicketsBadRequest{}
}

/*GetTicketsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetTicketsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetTicketsBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/tickets][%d] getTicketsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTicketsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetTicketsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTicketsUnauthorized creates a GetTicketsUnauthorized with default headers values
func NewGetTicketsUnauthorized() *GetTicketsUnauthorized {
	return &GetTicketsUnauthorized{}
}

/*GetTicketsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetTicketsUnauthorized struct {
}

func (o *GetTicketsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/tickets][%d] getTicketsUnauthorized ", 401)
}

func (o *GetTicketsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTicketsForbidden creates a GetTicketsForbidden with default headers values
func NewGetTicketsForbidden() *GetTicketsForbidden {
	return &GetTicketsForbidden{}
}

/*GetTicketsForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetTicketsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetTicketsForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/tickets][%d] getTicketsForbidden  %+v", 403, o.Payload)
}

func (o *GetTicketsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetTicketsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTicketsNotFound creates a GetTicketsNotFound with default headers values
func NewGetTicketsNotFound() *GetTicketsNotFound {
	return &GetTicketsNotFound{}
}

/*GetTicketsNotFound handles this case with default header values.

Object not found
*/
type GetTicketsNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetTicketsNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/tickets][%d] getTicketsNotFound  %+v", 404, o.Payload)
}

func (o *GetTicketsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetTicketsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTicketsTooManyRequests creates a GetTicketsTooManyRequests with default headers values
func NewGetTicketsTooManyRequests() *GetTicketsTooManyRequests {
	return &GetTicketsTooManyRequests{}
}

/*GetTicketsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetTicketsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetTicketsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/tickets][%d] getTicketsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTicketsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetTicketsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
