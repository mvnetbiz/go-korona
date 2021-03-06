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

// GetPointOfSaleReader is a Reader for the GetPointOfSale structure.
type GetPointOfSaleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPointOfSaleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPointOfSaleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPointOfSaleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPointOfSaleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPointOfSaleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPointOfSaleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPointOfSaleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPointOfSaleOK creates a GetPointOfSaleOK with default headers values
func NewGetPointOfSaleOK() *GetPointOfSaleOK {
	return &GetPointOfSaleOK{}
}

/*GetPointOfSaleOK handles this case with default header values.

successful operation
*/
type GetPointOfSaleOK struct {
	Payload *models.Pos
}

func (o *GetPointOfSaleOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/pointsOfSale/{pointOfSaleId}][%d] getPointOfSaleOK  %+v", 200, o.Payload)
}

func (o *GetPointOfSaleOK) GetPayload() *models.Pos {
	return o.Payload
}

func (o *GetPointOfSaleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Pos)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPointOfSaleBadRequest creates a GetPointOfSaleBadRequest with default headers values
func NewGetPointOfSaleBadRequest() *GetPointOfSaleBadRequest {
	return &GetPointOfSaleBadRequest{}
}

/*GetPointOfSaleBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetPointOfSaleBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetPointOfSaleBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/pointsOfSale/{pointOfSaleId}][%d] getPointOfSaleBadRequest  %+v", 400, o.Payload)
}

func (o *GetPointOfSaleBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetPointOfSaleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPointOfSaleUnauthorized creates a GetPointOfSaleUnauthorized with default headers values
func NewGetPointOfSaleUnauthorized() *GetPointOfSaleUnauthorized {
	return &GetPointOfSaleUnauthorized{}
}

/*GetPointOfSaleUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetPointOfSaleUnauthorized struct {
}

func (o *GetPointOfSaleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/pointsOfSale/{pointOfSaleId}][%d] getPointOfSaleUnauthorized ", 401)
}

func (o *GetPointOfSaleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPointOfSaleForbidden creates a GetPointOfSaleForbidden with default headers values
func NewGetPointOfSaleForbidden() *GetPointOfSaleForbidden {
	return &GetPointOfSaleForbidden{}
}

/*GetPointOfSaleForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetPointOfSaleForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetPointOfSaleForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/pointsOfSale/{pointOfSaleId}][%d] getPointOfSaleForbidden  %+v", 403, o.Payload)
}

func (o *GetPointOfSaleForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetPointOfSaleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPointOfSaleNotFound creates a GetPointOfSaleNotFound with default headers values
func NewGetPointOfSaleNotFound() *GetPointOfSaleNotFound {
	return &GetPointOfSaleNotFound{}
}

/*GetPointOfSaleNotFound handles this case with default header values.

Object not found
*/
type GetPointOfSaleNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetPointOfSaleNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/pointsOfSale/{pointOfSaleId}][%d] getPointOfSaleNotFound  %+v", 404, o.Payload)
}

func (o *GetPointOfSaleNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetPointOfSaleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPointOfSaleTooManyRequests creates a GetPointOfSaleTooManyRequests with default headers values
func NewGetPointOfSaleTooManyRequests() *GetPointOfSaleTooManyRequests {
	return &GetPointOfSaleTooManyRequests{}
}

/*GetPointOfSaleTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetPointOfSaleTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetPointOfSaleTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/pointsOfSale/{pointOfSaleId}][%d] getPointOfSaleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPointOfSaleTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetPointOfSaleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
