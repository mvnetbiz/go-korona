// Code generated by go-swagger; DO NOT EDIT.

package ticketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// GetEntryGatesReader is a Reader for the GetEntryGates structure.
type GetEntryGatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEntryGatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEntryGatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetEntryGatesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEntryGatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetEntryGatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEntryGatesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEntryGatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetEntryGatesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEntryGatesOK creates a GetEntryGatesOK with default headers values
func NewGetEntryGatesOK() *GetEntryGatesOK {
	return &GetEntryGatesOK{}
}

/*GetEntryGatesOK handles this case with default header values.

successful operation
*/
type GetEntryGatesOK struct {
	Payload *models.ResultListEntryGate
}

func (o *GetEntryGatesOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/entryGates][%d] getEntryGatesOK  %+v", 200, o.Payload)
}

func (o *GetEntryGatesOK) GetPayload() *models.ResultListEntryGate {
	return o.Payload
}

func (o *GetEntryGatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResultListEntryGate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEntryGatesNoContent creates a GetEntryGatesNoContent with default headers values
func NewGetEntryGatesNoContent() *GetEntryGatesNoContent {
	return &GetEntryGatesNoContent{}
}

/*GetEntryGatesNoContent handles this case with default header values.

Request successful, but the list is empty. Either there is in general no object on the list or a set filter/restriction in querystring doesn't match any object.
*/
type GetEntryGatesNoContent struct {
}

func (o *GetEntryGatesNoContent) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/entryGates][%d] getEntryGatesNoContent ", 204)
}

func (o *GetEntryGatesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEntryGatesBadRequest creates a GetEntryGatesBadRequest with default headers values
func NewGetEntryGatesBadRequest() *GetEntryGatesBadRequest {
	return &GetEntryGatesBadRequest{}
}

/*GetEntryGatesBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetEntryGatesBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetEntryGatesBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/entryGates][%d] getEntryGatesBadRequest  %+v", 400, o.Payload)
}

func (o *GetEntryGatesBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetEntryGatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEntryGatesUnauthorized creates a GetEntryGatesUnauthorized with default headers values
func NewGetEntryGatesUnauthorized() *GetEntryGatesUnauthorized {
	return &GetEntryGatesUnauthorized{}
}

/*GetEntryGatesUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetEntryGatesUnauthorized struct {
}

func (o *GetEntryGatesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/entryGates][%d] getEntryGatesUnauthorized ", 401)
}

func (o *GetEntryGatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEntryGatesForbidden creates a GetEntryGatesForbidden with default headers values
func NewGetEntryGatesForbidden() *GetEntryGatesForbidden {
	return &GetEntryGatesForbidden{}
}

/*GetEntryGatesForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetEntryGatesForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetEntryGatesForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/entryGates][%d] getEntryGatesForbidden  %+v", 403, o.Payload)
}

func (o *GetEntryGatesForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetEntryGatesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEntryGatesNotFound creates a GetEntryGatesNotFound with default headers values
func NewGetEntryGatesNotFound() *GetEntryGatesNotFound {
	return &GetEntryGatesNotFound{}
}

/*GetEntryGatesNotFound handles this case with default header values.

Object not found
*/
type GetEntryGatesNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetEntryGatesNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/entryGates][%d] getEntryGatesNotFound  %+v", 404, o.Payload)
}

func (o *GetEntryGatesNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetEntryGatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEntryGatesTooManyRequests creates a GetEntryGatesTooManyRequests with default headers values
func NewGetEntryGatesTooManyRequests() *GetEntryGatesTooManyRequests {
	return &GetEntryGatesTooManyRequests{}
}

/*GetEntryGatesTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetEntryGatesTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetEntryGatesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/entryGates][%d] getEntryGatesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetEntryGatesTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetEntryGatesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
