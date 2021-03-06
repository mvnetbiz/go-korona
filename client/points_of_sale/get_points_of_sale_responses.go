// Code generated by go-swagger; DO NOT EDIT.

package points_of_sale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// GetPointsOfSaleReader is a Reader for the GetPointsOfSale structure.
type GetPointsOfSaleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPointsOfSaleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPointsOfSaleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetPointsOfSaleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPointsOfSaleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPointsOfSaleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPointsOfSaleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPointsOfSaleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPointsOfSaleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPointsOfSaleOK creates a GetPointsOfSaleOK with default headers values
func NewGetPointsOfSaleOK() *GetPointsOfSaleOK {
	return &GetPointsOfSaleOK{}
}

/*GetPointsOfSaleOK handles this case with default header values.

successful operation
*/
type GetPointsOfSaleOK struct {
	Payload *models.ResultListPos
}

func (o *GetPointsOfSaleOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/pointsOfSale][%d] getPointsOfSaleOK  %+v", 200, o.Payload)
}

func (o *GetPointsOfSaleOK) GetPayload() *models.ResultListPos {
	return o.Payload
}

func (o *GetPointsOfSaleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResultListPos)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPointsOfSaleNoContent creates a GetPointsOfSaleNoContent with default headers values
func NewGetPointsOfSaleNoContent() *GetPointsOfSaleNoContent {
	return &GetPointsOfSaleNoContent{}
}

/*GetPointsOfSaleNoContent handles this case with default header values.

Request successful, but the list is empty. Either there is in general no object on the list or a set filter/restriction in querystring doesn't match any object.
*/
type GetPointsOfSaleNoContent struct {
}

func (o *GetPointsOfSaleNoContent) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/pointsOfSale][%d] getPointsOfSaleNoContent ", 204)
}

func (o *GetPointsOfSaleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPointsOfSaleBadRequest creates a GetPointsOfSaleBadRequest with default headers values
func NewGetPointsOfSaleBadRequest() *GetPointsOfSaleBadRequest {
	return &GetPointsOfSaleBadRequest{}
}

/*GetPointsOfSaleBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetPointsOfSaleBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetPointsOfSaleBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/pointsOfSale][%d] getPointsOfSaleBadRequest  %+v", 400, o.Payload)
}

func (o *GetPointsOfSaleBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetPointsOfSaleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPointsOfSaleUnauthorized creates a GetPointsOfSaleUnauthorized with default headers values
func NewGetPointsOfSaleUnauthorized() *GetPointsOfSaleUnauthorized {
	return &GetPointsOfSaleUnauthorized{}
}

/*GetPointsOfSaleUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetPointsOfSaleUnauthorized struct {
}

func (o *GetPointsOfSaleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/pointsOfSale][%d] getPointsOfSaleUnauthorized ", 401)
}

func (o *GetPointsOfSaleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPointsOfSaleForbidden creates a GetPointsOfSaleForbidden with default headers values
func NewGetPointsOfSaleForbidden() *GetPointsOfSaleForbidden {
	return &GetPointsOfSaleForbidden{}
}

/*GetPointsOfSaleForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetPointsOfSaleForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetPointsOfSaleForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/pointsOfSale][%d] getPointsOfSaleForbidden  %+v", 403, o.Payload)
}

func (o *GetPointsOfSaleForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetPointsOfSaleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPointsOfSaleNotFound creates a GetPointsOfSaleNotFound with default headers values
func NewGetPointsOfSaleNotFound() *GetPointsOfSaleNotFound {
	return &GetPointsOfSaleNotFound{}
}

/*GetPointsOfSaleNotFound handles this case with default header values.

Object not found
*/
type GetPointsOfSaleNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetPointsOfSaleNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/pointsOfSale][%d] getPointsOfSaleNotFound  %+v", 404, o.Payload)
}

func (o *GetPointsOfSaleNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetPointsOfSaleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPointsOfSaleTooManyRequests creates a GetPointsOfSaleTooManyRequests with default headers values
func NewGetPointsOfSaleTooManyRequests() *GetPointsOfSaleTooManyRequests {
	return &GetPointsOfSaleTooManyRequests{}
}

/*GetPointsOfSaleTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetPointsOfSaleTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetPointsOfSaleTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/pointsOfSale][%d] getPointsOfSaleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPointsOfSaleTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetPointsOfSaleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
