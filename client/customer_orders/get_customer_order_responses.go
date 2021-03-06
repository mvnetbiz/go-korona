// Code generated by go-swagger; DO NOT EDIT.

package customer_orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// GetCustomerOrderReader is a Reader for the GetCustomerOrder structure.
type GetCustomerOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomerOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomerOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCustomerOrderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCustomerOrderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCustomerOrderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCustomerOrderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCustomerOrderTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCustomerOrderOK creates a GetCustomerOrderOK with default headers values
func NewGetCustomerOrderOK() *GetCustomerOrderOK {
	return &GetCustomerOrderOK{}
}

/*GetCustomerOrderOK handles this case with default header values.

successful operation
*/
type GetCustomerOrderOK struct {
	Payload *models.CustomerOrder
}

func (o *GetCustomerOrderOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/customerOrders/{customerOrderId}][%d] getCustomerOrderOK  %+v", 200, o.Payload)
}

func (o *GetCustomerOrderOK) GetPayload() *models.CustomerOrder {
	return o.Payload
}

func (o *GetCustomerOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerOrder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomerOrderBadRequest creates a GetCustomerOrderBadRequest with default headers values
func NewGetCustomerOrderBadRequest() *GetCustomerOrderBadRequest {
	return &GetCustomerOrderBadRequest{}
}

/*GetCustomerOrderBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetCustomerOrderBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetCustomerOrderBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/customerOrders/{customerOrderId}][%d] getCustomerOrderBadRequest  %+v", 400, o.Payload)
}

func (o *GetCustomerOrderBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetCustomerOrderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomerOrderUnauthorized creates a GetCustomerOrderUnauthorized with default headers values
func NewGetCustomerOrderUnauthorized() *GetCustomerOrderUnauthorized {
	return &GetCustomerOrderUnauthorized{}
}

/*GetCustomerOrderUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetCustomerOrderUnauthorized struct {
}

func (o *GetCustomerOrderUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/customerOrders/{customerOrderId}][%d] getCustomerOrderUnauthorized ", 401)
}

func (o *GetCustomerOrderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomerOrderForbidden creates a GetCustomerOrderForbidden with default headers values
func NewGetCustomerOrderForbidden() *GetCustomerOrderForbidden {
	return &GetCustomerOrderForbidden{}
}

/*GetCustomerOrderForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetCustomerOrderForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetCustomerOrderForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/customerOrders/{customerOrderId}][%d] getCustomerOrderForbidden  %+v", 403, o.Payload)
}

func (o *GetCustomerOrderForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetCustomerOrderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomerOrderNotFound creates a GetCustomerOrderNotFound with default headers values
func NewGetCustomerOrderNotFound() *GetCustomerOrderNotFound {
	return &GetCustomerOrderNotFound{}
}

/*GetCustomerOrderNotFound handles this case with default header values.

Object not found
*/
type GetCustomerOrderNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetCustomerOrderNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/customerOrders/{customerOrderId}][%d] getCustomerOrderNotFound  %+v", 404, o.Payload)
}

func (o *GetCustomerOrderNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetCustomerOrderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomerOrderTooManyRequests creates a GetCustomerOrderTooManyRequests with default headers values
func NewGetCustomerOrderTooManyRequests() *GetCustomerOrderTooManyRequests {
	return &GetCustomerOrderTooManyRequests{}
}

/*GetCustomerOrderTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetCustomerOrderTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetCustomerOrderTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/customerOrders/{customerOrderId}][%d] getCustomerOrderTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCustomerOrderTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetCustomerOrderTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
