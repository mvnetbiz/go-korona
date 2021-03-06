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

// AddTicketDefinitionsReader is a Reader for the AddTicketDefinitions structure.
type AddTicketDefinitionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddTicketDefinitionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddTicketDefinitionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddTicketDefinitionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddTicketDefinitionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddTicketDefinitionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddTicketDefinitionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAddTicketDefinitionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddTicketDefinitionsOK creates a AddTicketDefinitionsOK with default headers values
func NewAddTicketDefinitionsOK() *AddTicketDefinitionsOK {
	return &AddTicketDefinitionsOK{}
}

/*AddTicketDefinitionsOK handles this case with default header values.

successful operation
*/
type AddTicketDefinitionsOK struct {
	Payload []*models.AddOrUpdateResult
}

func (o *AddTicketDefinitionsOK) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/ticketDefinitions][%d] addTicketDefinitionsOK  %+v", 200, o.Payload)
}

func (o *AddTicketDefinitionsOK) GetPayload() []*models.AddOrUpdateResult {
	return o.Payload
}

func (o *AddTicketDefinitionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTicketDefinitionsBadRequest creates a AddTicketDefinitionsBadRequest with default headers values
func NewAddTicketDefinitionsBadRequest() *AddTicketDefinitionsBadRequest {
	return &AddTicketDefinitionsBadRequest{}
}

/*AddTicketDefinitionsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type AddTicketDefinitionsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *AddTicketDefinitionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/ticketDefinitions][%d] addTicketDefinitionsBadRequest  %+v", 400, o.Payload)
}

func (o *AddTicketDefinitionsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *AddTicketDefinitionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTicketDefinitionsUnauthorized creates a AddTicketDefinitionsUnauthorized with default headers values
func NewAddTicketDefinitionsUnauthorized() *AddTicketDefinitionsUnauthorized {
	return &AddTicketDefinitionsUnauthorized{}
}

/*AddTicketDefinitionsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type AddTicketDefinitionsUnauthorized struct {
}

func (o *AddTicketDefinitionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/ticketDefinitions][%d] addTicketDefinitionsUnauthorized ", 401)
}

func (o *AddTicketDefinitionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddTicketDefinitionsForbidden creates a AddTicketDefinitionsForbidden with default headers values
func NewAddTicketDefinitionsForbidden() *AddTicketDefinitionsForbidden {
	return &AddTicketDefinitionsForbidden{}
}

/*AddTicketDefinitionsForbidden handles this case with default header values.

Requested action is not allowed
*/
type AddTicketDefinitionsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *AddTicketDefinitionsForbidden) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/ticketDefinitions][%d] addTicketDefinitionsForbidden  %+v", 403, o.Payload)
}

func (o *AddTicketDefinitionsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *AddTicketDefinitionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTicketDefinitionsNotFound creates a AddTicketDefinitionsNotFound with default headers values
func NewAddTicketDefinitionsNotFound() *AddTicketDefinitionsNotFound {
	return &AddTicketDefinitionsNotFound{}
}

/*AddTicketDefinitionsNotFound handles this case with default header values.

Object not found
*/
type AddTicketDefinitionsNotFound struct {
	Payload *models.NotFoundError
}

func (o *AddTicketDefinitionsNotFound) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/ticketDefinitions][%d] addTicketDefinitionsNotFound  %+v", 404, o.Payload)
}

func (o *AddTicketDefinitionsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *AddTicketDefinitionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTicketDefinitionsTooManyRequests creates a AddTicketDefinitionsTooManyRequests with default headers values
func NewAddTicketDefinitionsTooManyRequests() *AddTicketDefinitionsTooManyRequests {
	return &AddTicketDefinitionsTooManyRequests{}
}

/*AddTicketDefinitionsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type AddTicketDefinitionsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *AddTicketDefinitionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/ticketDefinitions][%d] addTicketDefinitionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AddTicketDefinitionsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *AddTicketDefinitionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
