// Code generated by go-swagger; DO NOT EDIT.

package economic_zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// GetEconomicZoneReader is a Reader for the GetEconomicZone structure.
type GetEconomicZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEconomicZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEconomicZoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEconomicZoneBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetEconomicZoneUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEconomicZoneForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEconomicZoneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetEconomicZoneTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEconomicZoneOK creates a GetEconomicZoneOK with default headers values
func NewGetEconomicZoneOK() *GetEconomicZoneOK {
	return &GetEconomicZoneOK{}
}

/*GetEconomicZoneOK handles this case with default header values.

successful operation
*/
type GetEconomicZoneOK struct {
	Payload *models.EconomicZone
}

func (o *GetEconomicZoneOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/economicZones/{economicZoneId}][%d] getEconomicZoneOK  %+v", 200, o.Payload)
}

func (o *GetEconomicZoneOK) GetPayload() *models.EconomicZone {
	return o.Payload
}

func (o *GetEconomicZoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EconomicZone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEconomicZoneBadRequest creates a GetEconomicZoneBadRequest with default headers values
func NewGetEconomicZoneBadRequest() *GetEconomicZoneBadRequest {
	return &GetEconomicZoneBadRequest{}
}

/*GetEconomicZoneBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetEconomicZoneBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetEconomicZoneBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/economicZones/{economicZoneId}][%d] getEconomicZoneBadRequest  %+v", 400, o.Payload)
}

func (o *GetEconomicZoneBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetEconomicZoneBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEconomicZoneUnauthorized creates a GetEconomicZoneUnauthorized with default headers values
func NewGetEconomicZoneUnauthorized() *GetEconomicZoneUnauthorized {
	return &GetEconomicZoneUnauthorized{}
}

/*GetEconomicZoneUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetEconomicZoneUnauthorized struct {
}

func (o *GetEconomicZoneUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/economicZones/{economicZoneId}][%d] getEconomicZoneUnauthorized ", 401)
}

func (o *GetEconomicZoneUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEconomicZoneForbidden creates a GetEconomicZoneForbidden with default headers values
func NewGetEconomicZoneForbidden() *GetEconomicZoneForbidden {
	return &GetEconomicZoneForbidden{}
}

/*GetEconomicZoneForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetEconomicZoneForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetEconomicZoneForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/economicZones/{economicZoneId}][%d] getEconomicZoneForbidden  %+v", 403, o.Payload)
}

func (o *GetEconomicZoneForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetEconomicZoneForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEconomicZoneNotFound creates a GetEconomicZoneNotFound with default headers values
func NewGetEconomicZoneNotFound() *GetEconomicZoneNotFound {
	return &GetEconomicZoneNotFound{}
}

/*GetEconomicZoneNotFound handles this case with default header values.

Object not found
*/
type GetEconomicZoneNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetEconomicZoneNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/economicZones/{economicZoneId}][%d] getEconomicZoneNotFound  %+v", 404, o.Payload)
}

func (o *GetEconomicZoneNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetEconomicZoneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEconomicZoneTooManyRequests creates a GetEconomicZoneTooManyRequests with default headers values
func NewGetEconomicZoneTooManyRequests() *GetEconomicZoneTooManyRequests {
	return &GetEconomicZoneTooManyRequests{}
}

/*GetEconomicZoneTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetEconomicZoneTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetEconomicZoneTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/economicZones/{economicZoneId}][%d] getEconomicZoneTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetEconomicZoneTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetEconomicZoneTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
