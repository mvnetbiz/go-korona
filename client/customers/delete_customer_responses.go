// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// DeleteCustomerReader is a Reader for the DeleteCustomer structure.
type DeleteCustomerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCustomerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCustomerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteCustomerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCustomerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCustomerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteCustomerTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCustomerNoContent creates a DeleteCustomerNoContent with default headers values
func NewDeleteCustomerNoContent() *DeleteCustomerNoContent {
	return &DeleteCustomerNoContent{}
}

/*DeleteCustomerNoContent handles this case with default header values.

Request successful, delete on single resource does not return any content
*/
type DeleteCustomerNoContent struct {
}

func (o *DeleteCustomerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/customers/{customerId}][%d] deleteCustomerNoContent ", 204)
}

func (o *DeleteCustomerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomerBadRequest creates a DeleteCustomerBadRequest with default headers values
func NewDeleteCustomerBadRequest() *DeleteCustomerBadRequest {
	return &DeleteCustomerBadRequest{}
}

/*DeleteCustomerBadRequest handles this case with default header values.

Malformed querystring or model
*/
type DeleteCustomerBadRequest struct {
	Payload *models.BadRequestError
}

func (o *DeleteCustomerBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/customers/{customerId}][%d] deleteCustomerBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCustomerBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *DeleteCustomerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomerUnauthorized creates a DeleteCustomerUnauthorized with default headers values
func NewDeleteCustomerUnauthorized() *DeleteCustomerUnauthorized {
	return &DeleteCustomerUnauthorized{}
}

/*DeleteCustomerUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type DeleteCustomerUnauthorized struct {
}

func (o *DeleteCustomerUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/customers/{customerId}][%d] deleteCustomerUnauthorized ", 401)
}

func (o *DeleteCustomerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomerForbidden creates a DeleteCustomerForbidden with default headers values
func NewDeleteCustomerForbidden() *DeleteCustomerForbidden {
	return &DeleteCustomerForbidden{}
}

/*DeleteCustomerForbidden handles this case with default header values.

Requested action is not allowed
*/
type DeleteCustomerForbidden struct {
	Payload *models.ForbiddenError
}

func (o *DeleteCustomerForbidden) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/customers/{customerId}][%d] deleteCustomerForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCustomerForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *DeleteCustomerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomerNotFound creates a DeleteCustomerNotFound with default headers values
func NewDeleteCustomerNotFound() *DeleteCustomerNotFound {
	return &DeleteCustomerNotFound{}
}

/*DeleteCustomerNotFound handles this case with default header values.

Object not found
*/
type DeleteCustomerNotFound struct {
	Payload *models.NotFoundError
}

func (o *DeleteCustomerNotFound) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/customers/{customerId}][%d] deleteCustomerNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCustomerNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *DeleteCustomerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomerTooManyRequests creates a DeleteCustomerTooManyRequests with default headers values
func NewDeleteCustomerTooManyRequests() *DeleteCustomerTooManyRequests {
	return &DeleteCustomerTooManyRequests{}
}

/*DeleteCustomerTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type DeleteCustomerTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *DeleteCustomerTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/customers/{customerId}][%d] deleteCustomerTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteCustomerTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *DeleteCustomerTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
