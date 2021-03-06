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

// AddProductsReader is a Reader for the AddProducts structure.
type AddProductsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddProductsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddProductsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddProductsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddProductsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddProductsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddProductsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAddProductsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddProductsOK creates a AddProductsOK with default headers values
func NewAddProductsOK() *AddProductsOK {
	return &AddProductsOK{}
}

/*AddProductsOK handles this case with default header values.

successful operation
*/
type AddProductsOK struct {
	Payload []*models.AddOrUpdateResult
}

func (o *AddProductsOK) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/products][%d] addProductsOK  %+v", 200, o.Payload)
}

func (o *AddProductsOK) GetPayload() []*models.AddOrUpdateResult {
	return o.Payload
}

func (o *AddProductsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProductsBadRequest creates a AddProductsBadRequest with default headers values
func NewAddProductsBadRequest() *AddProductsBadRequest {
	return &AddProductsBadRequest{}
}

/*AddProductsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type AddProductsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *AddProductsBadRequest) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/products][%d] addProductsBadRequest  %+v", 400, o.Payload)
}

func (o *AddProductsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *AddProductsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProductsUnauthorized creates a AddProductsUnauthorized with default headers values
func NewAddProductsUnauthorized() *AddProductsUnauthorized {
	return &AddProductsUnauthorized{}
}

/*AddProductsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type AddProductsUnauthorized struct {
}

func (o *AddProductsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/products][%d] addProductsUnauthorized ", 401)
}

func (o *AddProductsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddProductsForbidden creates a AddProductsForbidden with default headers values
func NewAddProductsForbidden() *AddProductsForbidden {
	return &AddProductsForbidden{}
}

/*AddProductsForbidden handles this case with default header values.

Requested action is not allowed
*/
type AddProductsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *AddProductsForbidden) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/products][%d] addProductsForbidden  %+v", 403, o.Payload)
}

func (o *AddProductsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *AddProductsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProductsNotFound creates a AddProductsNotFound with default headers values
func NewAddProductsNotFound() *AddProductsNotFound {
	return &AddProductsNotFound{}
}

/*AddProductsNotFound handles this case with default header values.

Object not found
*/
type AddProductsNotFound struct {
	Payload *models.NotFoundError
}

func (o *AddProductsNotFound) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/products][%d] addProductsNotFound  %+v", 404, o.Payload)
}

func (o *AddProductsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *AddProductsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProductsTooManyRequests creates a AddProductsTooManyRequests with default headers values
func NewAddProductsTooManyRequests() *AddProductsTooManyRequests {
	return &AddProductsTooManyRequests{}
}

/*AddProductsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type AddProductsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *AddProductsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/products][%d] addProductsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AddProductsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *AddProductsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
