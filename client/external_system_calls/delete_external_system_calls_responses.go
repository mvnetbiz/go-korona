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

// DeleteExternalSystemCallsReader is a Reader for the DeleteExternalSystemCalls structure.
type DeleteExternalSystemCallsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteExternalSystemCallsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteExternalSystemCallsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteExternalSystemCallsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteExternalSystemCallsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteExternalSystemCallsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteExternalSystemCallsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteExternalSystemCallsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteExternalSystemCallsOK creates a DeleteExternalSystemCallsOK with default headers values
func NewDeleteExternalSystemCallsOK() *DeleteExternalSystemCallsOK {
	return &DeleteExternalSystemCallsOK{}
}

/*DeleteExternalSystemCallsOK handles this case with default header values.

successful operation
*/
type DeleteExternalSystemCallsOK struct {
	Payload []*models.AddOrUpdateResult
}

func (o *DeleteExternalSystemCallsOK) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/externalSystemCalls][%d] deleteExternalSystemCallsOK  %+v", 200, o.Payload)
}

func (o *DeleteExternalSystemCallsOK) GetPayload() []*models.AddOrUpdateResult {
	return o.Payload
}

func (o *DeleteExternalSystemCallsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalSystemCallsBadRequest creates a DeleteExternalSystemCallsBadRequest with default headers values
func NewDeleteExternalSystemCallsBadRequest() *DeleteExternalSystemCallsBadRequest {
	return &DeleteExternalSystemCallsBadRequest{}
}

/*DeleteExternalSystemCallsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type DeleteExternalSystemCallsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *DeleteExternalSystemCallsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/externalSystemCalls][%d] deleteExternalSystemCallsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteExternalSystemCallsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *DeleteExternalSystemCallsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalSystemCallsUnauthorized creates a DeleteExternalSystemCallsUnauthorized with default headers values
func NewDeleteExternalSystemCallsUnauthorized() *DeleteExternalSystemCallsUnauthorized {
	return &DeleteExternalSystemCallsUnauthorized{}
}

/*DeleteExternalSystemCallsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type DeleteExternalSystemCallsUnauthorized struct {
}

func (o *DeleteExternalSystemCallsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/externalSystemCalls][%d] deleteExternalSystemCallsUnauthorized ", 401)
}

func (o *DeleteExternalSystemCallsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteExternalSystemCallsForbidden creates a DeleteExternalSystemCallsForbidden with default headers values
func NewDeleteExternalSystemCallsForbidden() *DeleteExternalSystemCallsForbidden {
	return &DeleteExternalSystemCallsForbidden{}
}

/*DeleteExternalSystemCallsForbidden handles this case with default header values.

Requested action is not allowed
*/
type DeleteExternalSystemCallsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *DeleteExternalSystemCallsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/externalSystemCalls][%d] deleteExternalSystemCallsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteExternalSystemCallsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *DeleteExternalSystemCallsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalSystemCallsNotFound creates a DeleteExternalSystemCallsNotFound with default headers values
func NewDeleteExternalSystemCallsNotFound() *DeleteExternalSystemCallsNotFound {
	return &DeleteExternalSystemCallsNotFound{}
}

/*DeleteExternalSystemCallsNotFound handles this case with default header values.

Object not found
*/
type DeleteExternalSystemCallsNotFound struct {
	Payload *models.NotFoundError
}

func (o *DeleteExternalSystemCallsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/externalSystemCalls][%d] deleteExternalSystemCallsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteExternalSystemCallsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *DeleteExternalSystemCallsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalSystemCallsTooManyRequests creates a DeleteExternalSystemCallsTooManyRequests with default headers values
func NewDeleteExternalSystemCallsTooManyRequests() *DeleteExternalSystemCallsTooManyRequests {
	return &DeleteExternalSystemCallsTooManyRequests{}
}

/*DeleteExternalSystemCallsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type DeleteExternalSystemCallsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *DeleteExternalSystemCallsTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/externalSystemCalls][%d] deleteExternalSystemCallsTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteExternalSystemCallsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *DeleteExternalSystemCallsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
