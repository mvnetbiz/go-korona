// Code generated by go-swagger; DO NOT EDIT.

package external_system_calls

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// UpdateExternalSystemCallReader is a Reader for the UpdateExternalSystemCall structure.
type UpdateExternalSystemCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateExternalSystemCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateExternalSystemCallNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateExternalSystemCallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateExternalSystemCallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateExternalSystemCallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateExternalSystemCallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateExternalSystemCallTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateExternalSystemCallNoContent creates a UpdateExternalSystemCallNoContent with default headers values
func NewUpdateExternalSystemCallNoContent() *UpdateExternalSystemCallNoContent {
	return &UpdateExternalSystemCallNoContent{}
}

/*UpdateExternalSystemCallNoContent handles this case with default header values.

Request successful, patch on single resource does not return any content
*/
type UpdateExternalSystemCallNoContent struct {
}

func (o *UpdateExternalSystemCallNoContent) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/externalSystemCalls/{externalSystemCallId}][%d] updateExternalSystemCallNoContent ", 204)
}

func (o *UpdateExternalSystemCallNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateExternalSystemCallBadRequest creates a UpdateExternalSystemCallBadRequest with default headers values
func NewUpdateExternalSystemCallBadRequest() *UpdateExternalSystemCallBadRequest {
	return &UpdateExternalSystemCallBadRequest{}
}

/*UpdateExternalSystemCallBadRequest handles this case with default header values.

Malformed querystring or model
*/
type UpdateExternalSystemCallBadRequest struct {
	Payload *models.BadRequestError
}

func (o *UpdateExternalSystemCallBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/externalSystemCalls/{externalSystemCallId}][%d] updateExternalSystemCallBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateExternalSystemCallBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *UpdateExternalSystemCallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateExternalSystemCallUnauthorized creates a UpdateExternalSystemCallUnauthorized with default headers values
func NewUpdateExternalSystemCallUnauthorized() *UpdateExternalSystemCallUnauthorized {
	return &UpdateExternalSystemCallUnauthorized{}
}

/*UpdateExternalSystemCallUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type UpdateExternalSystemCallUnauthorized struct {
}

func (o *UpdateExternalSystemCallUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/externalSystemCalls/{externalSystemCallId}][%d] updateExternalSystemCallUnauthorized ", 401)
}

func (o *UpdateExternalSystemCallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateExternalSystemCallForbidden creates a UpdateExternalSystemCallForbidden with default headers values
func NewUpdateExternalSystemCallForbidden() *UpdateExternalSystemCallForbidden {
	return &UpdateExternalSystemCallForbidden{}
}

/*UpdateExternalSystemCallForbidden handles this case with default header values.

Requested action is not allowed
*/
type UpdateExternalSystemCallForbidden struct {
	Payload *models.ForbiddenError
}

func (o *UpdateExternalSystemCallForbidden) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/externalSystemCalls/{externalSystemCallId}][%d] updateExternalSystemCallForbidden  %+v", 403, o.Payload)
}

func (o *UpdateExternalSystemCallForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *UpdateExternalSystemCallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateExternalSystemCallNotFound creates a UpdateExternalSystemCallNotFound with default headers values
func NewUpdateExternalSystemCallNotFound() *UpdateExternalSystemCallNotFound {
	return &UpdateExternalSystemCallNotFound{}
}

/*UpdateExternalSystemCallNotFound handles this case with default header values.

Object not found
*/
type UpdateExternalSystemCallNotFound struct {
	Payload *models.NotFoundError
}

func (o *UpdateExternalSystemCallNotFound) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/externalSystemCalls/{externalSystemCallId}][%d] updateExternalSystemCallNotFound  %+v", 404, o.Payload)
}

func (o *UpdateExternalSystemCallNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *UpdateExternalSystemCallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateExternalSystemCallTooManyRequests creates a UpdateExternalSystemCallTooManyRequests with default headers values
func NewUpdateExternalSystemCallTooManyRequests() *UpdateExternalSystemCallTooManyRequests {
	return &UpdateExternalSystemCallTooManyRequests{}
}

/*UpdateExternalSystemCallTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type UpdateExternalSystemCallTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *UpdateExternalSystemCallTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/externalSystemCalls/{externalSystemCallId}][%d] updateExternalSystemCallTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateExternalSystemCallTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *UpdateExternalSystemCallTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
