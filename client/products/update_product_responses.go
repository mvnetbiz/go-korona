// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// UpdateProductReader is a Reader for the UpdateProduct structure.
type UpdateProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateProductNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateProductBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateProductUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateProductForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateProductNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateProductTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateProductNoContent creates a UpdateProductNoContent with default headers values
func NewUpdateProductNoContent() *UpdateProductNoContent {
	return &UpdateProductNoContent{}
}

/*UpdateProductNoContent handles this case with default header values.

Request successful, patch on single resource does not return any content
*/
type UpdateProductNoContent struct {
}

func (o *UpdateProductNoContent) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/products/{productId}][%d] updateProductNoContent ", 204)
}

func (o *UpdateProductNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProductBadRequest creates a UpdateProductBadRequest with default headers values
func NewUpdateProductBadRequest() *UpdateProductBadRequest {
	return &UpdateProductBadRequest{}
}

/*UpdateProductBadRequest handles this case with default header values.

Malformed querystring or model
*/
type UpdateProductBadRequest struct {
	Payload *models.BadRequestError
}

func (o *UpdateProductBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/products/{productId}][%d] updateProductBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateProductBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *UpdateProductBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProductUnauthorized creates a UpdateProductUnauthorized with default headers values
func NewUpdateProductUnauthorized() *UpdateProductUnauthorized {
	return &UpdateProductUnauthorized{}
}

/*UpdateProductUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type UpdateProductUnauthorized struct {
}

func (o *UpdateProductUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/products/{productId}][%d] updateProductUnauthorized ", 401)
}

func (o *UpdateProductUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProductForbidden creates a UpdateProductForbidden with default headers values
func NewUpdateProductForbidden() *UpdateProductForbidden {
	return &UpdateProductForbidden{}
}

/*UpdateProductForbidden handles this case with default header values.

Requested action is not allowed
*/
type UpdateProductForbidden struct {
	Payload *models.ForbiddenError
}

func (o *UpdateProductForbidden) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/products/{productId}][%d] updateProductForbidden  %+v", 403, o.Payload)
}

func (o *UpdateProductForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *UpdateProductForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProductNotFound creates a UpdateProductNotFound with default headers values
func NewUpdateProductNotFound() *UpdateProductNotFound {
	return &UpdateProductNotFound{}
}

/*UpdateProductNotFound handles this case with default header values.

Object not found
*/
type UpdateProductNotFound struct {
	Payload *models.NotFoundError
}

func (o *UpdateProductNotFound) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/products/{productId}][%d] updateProductNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProductNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *UpdateProductNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProductTooManyRequests creates a UpdateProductTooManyRequests with default headers values
func NewUpdateProductTooManyRequests() *UpdateProductTooManyRequests {
	return &UpdateProductTooManyRequests{}
}

/*UpdateProductTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type UpdateProductTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *UpdateProductTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/products/{productId}][%d] updateProductTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateProductTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *UpdateProductTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
