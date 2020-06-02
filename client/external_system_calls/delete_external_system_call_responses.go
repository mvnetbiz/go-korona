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

// DeleteExternalSystemCallReader is a Reader for the DeleteExternalSystemCall structure.
type DeleteExternalSystemCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteExternalSystemCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteExternalSystemCallNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteExternalSystemCallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteExternalSystemCallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteExternalSystemCallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteExternalSystemCallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteExternalSystemCallTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteExternalSystemCallNoContent creates a DeleteExternalSystemCallNoContent with default headers values
func NewDeleteExternalSystemCallNoContent() *DeleteExternalSystemCallNoContent {
	return &DeleteExternalSystemCallNoContent{}
}

/*DeleteExternalSystemCallNoContent handles this case with default header values.

Request successful, delete on single resource does not return any content
*/
type DeleteExternalSystemCallNoContent struct {
}

func (o *DeleteExternalSystemCallNoContent) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/externalSystemCalls/{externalSystemCallId}][%d] deleteExternalSystemCallNoContent ", 204)
}

func (o *DeleteExternalSystemCallNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteExternalSystemCallBadRequest creates a DeleteExternalSystemCallBadRequest with default headers values
func NewDeleteExternalSystemCallBadRequest() *DeleteExternalSystemCallBadRequest {
	return &DeleteExternalSystemCallBadRequest{}
}

/*DeleteExternalSystemCallBadRequest handles this case with default header values.

Malformed querystring or model
*/
type DeleteExternalSystemCallBadRequest struct {
	Payload *models.BadRequestError
}

func (o *DeleteExternalSystemCallBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/externalSystemCalls/{externalSystemCallId}][%d] deleteExternalSystemCallBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteExternalSystemCallBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *DeleteExternalSystemCallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalSystemCallUnauthorized creates a DeleteExternalSystemCallUnauthorized with default headers values
func NewDeleteExternalSystemCallUnauthorized() *DeleteExternalSystemCallUnauthorized {
	return &DeleteExternalSystemCallUnauthorized{}
}

/*DeleteExternalSystemCallUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type DeleteExternalSystemCallUnauthorized struct {
}

func (o *DeleteExternalSystemCallUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/externalSystemCalls/{externalSystemCallId}][%d] deleteExternalSystemCallUnauthorized ", 401)
}

func (o *DeleteExternalSystemCallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteExternalSystemCallForbidden creates a DeleteExternalSystemCallForbidden with default headers values
func NewDeleteExternalSystemCallForbidden() *DeleteExternalSystemCallForbidden {
	return &DeleteExternalSystemCallForbidden{}
}

/*DeleteExternalSystemCallForbidden handles this case with default header values.

Requested action is not allowed
*/
type DeleteExternalSystemCallForbidden struct {
	Payload *models.ForbiddenError
}

func (o *DeleteExternalSystemCallForbidden) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/externalSystemCalls/{externalSystemCallId}][%d] deleteExternalSystemCallForbidden  %+v", 403, o.Payload)
}

func (o *DeleteExternalSystemCallForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *DeleteExternalSystemCallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalSystemCallNotFound creates a DeleteExternalSystemCallNotFound with default headers values
func NewDeleteExternalSystemCallNotFound() *DeleteExternalSystemCallNotFound {
	return &DeleteExternalSystemCallNotFound{}
}

/*DeleteExternalSystemCallNotFound handles this case with default header values.

Object not found
*/
type DeleteExternalSystemCallNotFound struct {
	Payload *models.NotFoundError
}

func (o *DeleteExternalSystemCallNotFound) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/externalSystemCalls/{externalSystemCallId}][%d] deleteExternalSystemCallNotFound  %+v", 404, o.Payload)
}

func (o *DeleteExternalSystemCallNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *DeleteExternalSystemCallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalSystemCallTooManyRequests creates a DeleteExternalSystemCallTooManyRequests with default headers values
func NewDeleteExternalSystemCallTooManyRequests() *DeleteExternalSystemCallTooManyRequests {
	return &DeleteExternalSystemCallTooManyRequests{}
}

/*DeleteExternalSystemCallTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type DeleteExternalSystemCallTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *DeleteExternalSystemCallTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/externalSystemCalls/{externalSystemCallId}][%d] deleteExternalSystemCallTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteExternalSystemCallTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *DeleteExternalSystemCallTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
