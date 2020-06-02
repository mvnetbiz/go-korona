// Code generated by go-swagger; DO NOT EDIT.

package organizational_units

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// GetOrganizationalUnitProductStocksReader is a Reader for the GetOrganizationalUnitProductStocks structure.
type GetOrganizationalUnitProductStocksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationalUnitProductStocksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationalUnitProductStocksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetOrganizationalUnitProductStocksNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrganizationalUnitProductStocksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrganizationalUnitProductStocksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationalUnitProductStocksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationalUnitProductStocksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrganizationalUnitProductStocksTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganizationalUnitProductStocksOK creates a GetOrganizationalUnitProductStocksOK with default headers values
func NewGetOrganizationalUnitProductStocksOK() *GetOrganizationalUnitProductStocksOK {
	return &GetOrganizationalUnitProductStocksOK{}
}

/*GetOrganizationalUnitProductStocksOK handles this case with default header values.

successful operation
*/
type GetOrganizationalUnitProductStocksOK struct {
	Payload *models.ResultListProductStock
}

func (o *GetOrganizationalUnitProductStocksOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/productStocks][%d] getOrganizationalUnitProductStocksOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationalUnitProductStocksOK) GetPayload() *models.ResultListProductStock {
	return o.Payload
}

func (o *GetOrganizationalUnitProductStocksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResultListProductStock)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationalUnitProductStocksNoContent creates a GetOrganizationalUnitProductStocksNoContent with default headers values
func NewGetOrganizationalUnitProductStocksNoContent() *GetOrganizationalUnitProductStocksNoContent {
	return &GetOrganizationalUnitProductStocksNoContent{}
}

/*GetOrganizationalUnitProductStocksNoContent handles this case with default header values.

Request successful, but the list is empty. Either there is in general no object on the list or a set filter/restriction in querystring doesn't match any object.
*/
type GetOrganizationalUnitProductStocksNoContent struct {
}

func (o *GetOrganizationalUnitProductStocksNoContent) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/productStocks][%d] getOrganizationalUnitProductStocksNoContent ", 204)
}

func (o *GetOrganizationalUnitProductStocksNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationalUnitProductStocksBadRequest creates a GetOrganizationalUnitProductStocksBadRequest with default headers values
func NewGetOrganizationalUnitProductStocksBadRequest() *GetOrganizationalUnitProductStocksBadRequest {
	return &GetOrganizationalUnitProductStocksBadRequest{}
}

/*GetOrganizationalUnitProductStocksBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetOrganizationalUnitProductStocksBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetOrganizationalUnitProductStocksBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/productStocks][%d] getOrganizationalUnitProductStocksBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationalUnitProductStocksBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetOrganizationalUnitProductStocksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationalUnitProductStocksUnauthorized creates a GetOrganizationalUnitProductStocksUnauthorized with default headers values
func NewGetOrganizationalUnitProductStocksUnauthorized() *GetOrganizationalUnitProductStocksUnauthorized {
	return &GetOrganizationalUnitProductStocksUnauthorized{}
}

/*GetOrganizationalUnitProductStocksUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetOrganizationalUnitProductStocksUnauthorized struct {
}

func (o *GetOrganizationalUnitProductStocksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/productStocks][%d] getOrganizationalUnitProductStocksUnauthorized ", 401)
}

func (o *GetOrganizationalUnitProductStocksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationalUnitProductStocksForbidden creates a GetOrganizationalUnitProductStocksForbidden with default headers values
func NewGetOrganizationalUnitProductStocksForbidden() *GetOrganizationalUnitProductStocksForbidden {
	return &GetOrganizationalUnitProductStocksForbidden{}
}

/*GetOrganizationalUnitProductStocksForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetOrganizationalUnitProductStocksForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetOrganizationalUnitProductStocksForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/productStocks][%d] getOrganizationalUnitProductStocksForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationalUnitProductStocksForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetOrganizationalUnitProductStocksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationalUnitProductStocksNotFound creates a GetOrganizationalUnitProductStocksNotFound with default headers values
func NewGetOrganizationalUnitProductStocksNotFound() *GetOrganizationalUnitProductStocksNotFound {
	return &GetOrganizationalUnitProductStocksNotFound{}
}

/*GetOrganizationalUnitProductStocksNotFound handles this case with default header values.

Object not found
*/
type GetOrganizationalUnitProductStocksNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetOrganizationalUnitProductStocksNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/productStocks][%d] getOrganizationalUnitProductStocksNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationalUnitProductStocksNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetOrganizationalUnitProductStocksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationalUnitProductStocksTooManyRequests creates a GetOrganizationalUnitProductStocksTooManyRequests with default headers values
func NewGetOrganizationalUnitProductStocksTooManyRequests() *GetOrganizationalUnitProductStocksTooManyRequests {
	return &GetOrganizationalUnitProductStocksTooManyRequests{}
}

/*GetOrganizationalUnitProductStocksTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetOrganizationalUnitProductStocksTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetOrganizationalUnitProductStocksTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/productStocks][%d] getOrganizationalUnitProductStocksTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrganizationalUnitProductStocksTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetOrganizationalUnitProductStocksTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}