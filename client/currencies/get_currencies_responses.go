// Code generated by go-swagger; DO NOT EDIT.

package currencies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// GetCurrenciesReader is a Reader for the GetCurrencies structure.
type GetCurrenciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrenciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrenciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetCurrenciesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCurrenciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCurrenciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCurrenciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCurrenciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCurrenciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCurrenciesOK creates a GetCurrenciesOK with default headers values
func NewGetCurrenciesOK() *GetCurrenciesOK {
	return &GetCurrenciesOK{}
}

/*GetCurrenciesOK handles this case with default header values.

successful operation
*/
type GetCurrenciesOK struct {
	Payload *models.ResultListCurrency
}

func (o *GetCurrenciesOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/currencies][%d] getCurrenciesOK  %+v", 200, o.Payload)
}

func (o *GetCurrenciesOK) GetPayload() *models.ResultListCurrency {
	return o.Payload
}

func (o *GetCurrenciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResultListCurrency)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrenciesNoContent creates a GetCurrenciesNoContent with default headers values
func NewGetCurrenciesNoContent() *GetCurrenciesNoContent {
	return &GetCurrenciesNoContent{}
}

/*GetCurrenciesNoContent handles this case with default header values.

Request successful, but the list is empty. Either there is in general no object on the list or a set filter/restriction in querystring doesn't match any object.
*/
type GetCurrenciesNoContent struct {
}

func (o *GetCurrenciesNoContent) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/currencies][%d] getCurrenciesNoContent ", 204)
}

func (o *GetCurrenciesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCurrenciesBadRequest creates a GetCurrenciesBadRequest with default headers values
func NewGetCurrenciesBadRequest() *GetCurrenciesBadRequest {
	return &GetCurrenciesBadRequest{}
}

/*GetCurrenciesBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetCurrenciesBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetCurrenciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/currencies][%d] getCurrenciesBadRequest  %+v", 400, o.Payload)
}

func (o *GetCurrenciesBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetCurrenciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrenciesUnauthorized creates a GetCurrenciesUnauthorized with default headers values
func NewGetCurrenciesUnauthorized() *GetCurrenciesUnauthorized {
	return &GetCurrenciesUnauthorized{}
}

/*GetCurrenciesUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetCurrenciesUnauthorized struct {
}

func (o *GetCurrenciesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/currencies][%d] getCurrenciesUnauthorized ", 401)
}

func (o *GetCurrenciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCurrenciesForbidden creates a GetCurrenciesForbidden with default headers values
func NewGetCurrenciesForbidden() *GetCurrenciesForbidden {
	return &GetCurrenciesForbidden{}
}

/*GetCurrenciesForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetCurrenciesForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetCurrenciesForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/currencies][%d] getCurrenciesForbidden  %+v", 403, o.Payload)
}

func (o *GetCurrenciesForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetCurrenciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrenciesNotFound creates a GetCurrenciesNotFound with default headers values
func NewGetCurrenciesNotFound() *GetCurrenciesNotFound {
	return &GetCurrenciesNotFound{}
}

/*GetCurrenciesNotFound handles this case with default header values.

Object not found
*/
type GetCurrenciesNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetCurrenciesNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/currencies][%d] getCurrenciesNotFound  %+v", 404, o.Payload)
}

func (o *GetCurrenciesNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetCurrenciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrenciesTooManyRequests creates a GetCurrenciesTooManyRequests with default headers values
func NewGetCurrenciesTooManyRequests() *GetCurrenciesTooManyRequests {
	return &GetCurrenciesTooManyRequests{}
}

/*GetCurrenciesTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetCurrenciesTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetCurrenciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/currencies][%d] getCurrenciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCurrenciesTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetCurrenciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
