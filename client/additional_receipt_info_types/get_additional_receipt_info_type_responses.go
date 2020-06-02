// Code generated by go-swagger; DO NOT EDIT.

package additional_receipt_info_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// GetAdditionalReceiptInfoTypeReader is a Reader for the GetAdditionalReceiptInfoType structure.
type GetAdditionalReceiptInfoTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdditionalReceiptInfoTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAdditionalReceiptInfoTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAdditionalReceiptInfoTypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAdditionalReceiptInfoTypeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAdditionalReceiptInfoTypeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAdditionalReceiptInfoTypeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAdditionalReceiptInfoTypeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAdditionalReceiptInfoTypeOK creates a GetAdditionalReceiptInfoTypeOK with default headers values
func NewGetAdditionalReceiptInfoTypeOK() *GetAdditionalReceiptInfoTypeOK {
	return &GetAdditionalReceiptInfoTypeOK{}
}

/*GetAdditionalReceiptInfoTypeOK handles this case with default header values.

successful operation
*/
type GetAdditionalReceiptInfoTypeOK struct {
	Payload *models.AdditionalReceiptInfoType
}

func (o *GetAdditionalReceiptInfoTypeOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/additionalReceiptInfoTypes/{additionalReceiptInfoTypeId}][%d] getAdditionalReceiptInfoTypeOK  %+v", 200, o.Payload)
}

func (o *GetAdditionalReceiptInfoTypeOK) GetPayload() *models.AdditionalReceiptInfoType {
	return o.Payload
}

func (o *GetAdditionalReceiptInfoTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdditionalReceiptInfoType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdditionalReceiptInfoTypeBadRequest creates a GetAdditionalReceiptInfoTypeBadRequest with default headers values
func NewGetAdditionalReceiptInfoTypeBadRequest() *GetAdditionalReceiptInfoTypeBadRequest {
	return &GetAdditionalReceiptInfoTypeBadRequest{}
}

/*GetAdditionalReceiptInfoTypeBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetAdditionalReceiptInfoTypeBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetAdditionalReceiptInfoTypeBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/additionalReceiptInfoTypes/{additionalReceiptInfoTypeId}][%d] getAdditionalReceiptInfoTypeBadRequest  %+v", 400, o.Payload)
}

func (o *GetAdditionalReceiptInfoTypeBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetAdditionalReceiptInfoTypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdditionalReceiptInfoTypeUnauthorized creates a GetAdditionalReceiptInfoTypeUnauthorized with default headers values
func NewGetAdditionalReceiptInfoTypeUnauthorized() *GetAdditionalReceiptInfoTypeUnauthorized {
	return &GetAdditionalReceiptInfoTypeUnauthorized{}
}

/*GetAdditionalReceiptInfoTypeUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetAdditionalReceiptInfoTypeUnauthorized struct {
}

func (o *GetAdditionalReceiptInfoTypeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/additionalReceiptInfoTypes/{additionalReceiptInfoTypeId}][%d] getAdditionalReceiptInfoTypeUnauthorized ", 401)
}

func (o *GetAdditionalReceiptInfoTypeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAdditionalReceiptInfoTypeForbidden creates a GetAdditionalReceiptInfoTypeForbidden with default headers values
func NewGetAdditionalReceiptInfoTypeForbidden() *GetAdditionalReceiptInfoTypeForbidden {
	return &GetAdditionalReceiptInfoTypeForbidden{}
}

/*GetAdditionalReceiptInfoTypeForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetAdditionalReceiptInfoTypeForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetAdditionalReceiptInfoTypeForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/additionalReceiptInfoTypes/{additionalReceiptInfoTypeId}][%d] getAdditionalReceiptInfoTypeForbidden  %+v", 403, o.Payload)
}

func (o *GetAdditionalReceiptInfoTypeForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetAdditionalReceiptInfoTypeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdditionalReceiptInfoTypeNotFound creates a GetAdditionalReceiptInfoTypeNotFound with default headers values
func NewGetAdditionalReceiptInfoTypeNotFound() *GetAdditionalReceiptInfoTypeNotFound {
	return &GetAdditionalReceiptInfoTypeNotFound{}
}

/*GetAdditionalReceiptInfoTypeNotFound handles this case with default header values.

Object not found
*/
type GetAdditionalReceiptInfoTypeNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetAdditionalReceiptInfoTypeNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/additionalReceiptInfoTypes/{additionalReceiptInfoTypeId}][%d] getAdditionalReceiptInfoTypeNotFound  %+v", 404, o.Payload)
}

func (o *GetAdditionalReceiptInfoTypeNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetAdditionalReceiptInfoTypeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdditionalReceiptInfoTypeTooManyRequests creates a GetAdditionalReceiptInfoTypeTooManyRequests with default headers values
func NewGetAdditionalReceiptInfoTypeTooManyRequests() *GetAdditionalReceiptInfoTypeTooManyRequests {
	return &GetAdditionalReceiptInfoTypeTooManyRequests{}
}

/*GetAdditionalReceiptInfoTypeTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetAdditionalReceiptInfoTypeTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetAdditionalReceiptInfoTypeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/additionalReceiptInfoTypes/{additionalReceiptInfoTypeId}][%d] getAdditionalReceiptInfoTypeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAdditionalReceiptInfoTypeTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetAdditionalReceiptInfoTypeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
