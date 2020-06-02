// Code generated by go-swagger; DO NOT EDIT.

package korona_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// GetKoronaAccountReader is a Reader for the GetKoronaAccount structure.
type GetKoronaAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKoronaAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKoronaAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKoronaAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKoronaAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKoronaAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKoronaAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetKoronaAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKoronaAccountOK creates a GetKoronaAccountOK with default headers values
func NewGetKoronaAccountOK() *GetKoronaAccountOK {
	return &GetKoronaAccountOK{}
}

/*GetKoronaAccountOK handles this case with default header values.

successful operation
*/
type GetKoronaAccountOK struct {
	Payload *models.KoronaAccount
}

func (o *GetKoronaAccountOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}][%d] getKoronaAccountOK  %+v", 200, o.Payload)
}

func (o *GetKoronaAccountOK) GetPayload() *models.KoronaAccount {
	return o.Payload
}

func (o *GetKoronaAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KoronaAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKoronaAccountBadRequest creates a GetKoronaAccountBadRequest with default headers values
func NewGetKoronaAccountBadRequest() *GetKoronaAccountBadRequest {
	return &GetKoronaAccountBadRequest{}
}

/*GetKoronaAccountBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetKoronaAccountBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetKoronaAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}][%d] getKoronaAccountBadRequest  %+v", 400, o.Payload)
}

func (o *GetKoronaAccountBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetKoronaAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKoronaAccountUnauthorized creates a GetKoronaAccountUnauthorized with default headers values
func NewGetKoronaAccountUnauthorized() *GetKoronaAccountUnauthorized {
	return &GetKoronaAccountUnauthorized{}
}

/*GetKoronaAccountUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetKoronaAccountUnauthorized struct {
}

func (o *GetKoronaAccountUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}][%d] getKoronaAccountUnauthorized ", 401)
}

func (o *GetKoronaAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKoronaAccountForbidden creates a GetKoronaAccountForbidden with default headers values
func NewGetKoronaAccountForbidden() *GetKoronaAccountForbidden {
	return &GetKoronaAccountForbidden{}
}

/*GetKoronaAccountForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetKoronaAccountForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetKoronaAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}][%d] getKoronaAccountForbidden  %+v", 403, o.Payload)
}

func (o *GetKoronaAccountForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetKoronaAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKoronaAccountNotFound creates a GetKoronaAccountNotFound with default headers values
func NewGetKoronaAccountNotFound() *GetKoronaAccountNotFound {
	return &GetKoronaAccountNotFound{}
}

/*GetKoronaAccountNotFound handles this case with default header values.

Object not found
*/
type GetKoronaAccountNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetKoronaAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}][%d] getKoronaAccountNotFound  %+v", 404, o.Payload)
}

func (o *GetKoronaAccountNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetKoronaAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKoronaAccountTooManyRequests creates a GetKoronaAccountTooManyRequests with default headers values
func NewGetKoronaAccountTooManyRequests() *GetKoronaAccountTooManyRequests {
	return &GetKoronaAccountTooManyRequests{}
}

/*GetKoronaAccountTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetKoronaAccountTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetKoronaAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}][%d] getKoronaAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKoronaAccountTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetKoronaAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}