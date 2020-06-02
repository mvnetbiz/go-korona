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

// GetOrganizationalUnitDayRatingReader is a Reader for the GetOrganizationalUnitDayRating structure.
type GetOrganizationalUnitDayRatingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationalUnitDayRatingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationalUnitDayRatingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrganizationalUnitDayRatingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrganizationalUnitDayRatingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationalUnitDayRatingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationalUnitDayRatingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrganizationalUnitDayRatingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganizationalUnitDayRatingOK creates a GetOrganizationalUnitDayRatingOK with default headers values
func NewGetOrganizationalUnitDayRatingOK() *GetOrganizationalUnitDayRatingOK {
	return &GetOrganizationalUnitDayRatingOK{}
}

/*GetOrganizationalUnitDayRatingOK handles this case with default header values.

successful operation
*/
type GetOrganizationalUnitDayRatingOK struct {
	Payload *models.DayRating
}

func (o *GetOrganizationalUnitDayRatingOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/dayRatings/{dayRatingIdOrDate}][%d] getOrganizationalUnitDayRatingOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationalUnitDayRatingOK) GetPayload() *models.DayRating {
	return o.Payload
}

func (o *GetOrganizationalUnitDayRatingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DayRating)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationalUnitDayRatingBadRequest creates a GetOrganizationalUnitDayRatingBadRequest with default headers values
func NewGetOrganizationalUnitDayRatingBadRequest() *GetOrganizationalUnitDayRatingBadRequest {
	return &GetOrganizationalUnitDayRatingBadRequest{}
}

/*GetOrganizationalUnitDayRatingBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetOrganizationalUnitDayRatingBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetOrganizationalUnitDayRatingBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/dayRatings/{dayRatingIdOrDate}][%d] getOrganizationalUnitDayRatingBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationalUnitDayRatingBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetOrganizationalUnitDayRatingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationalUnitDayRatingUnauthorized creates a GetOrganizationalUnitDayRatingUnauthorized with default headers values
func NewGetOrganizationalUnitDayRatingUnauthorized() *GetOrganizationalUnitDayRatingUnauthorized {
	return &GetOrganizationalUnitDayRatingUnauthorized{}
}

/*GetOrganizationalUnitDayRatingUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetOrganizationalUnitDayRatingUnauthorized struct {
}

func (o *GetOrganizationalUnitDayRatingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/dayRatings/{dayRatingIdOrDate}][%d] getOrganizationalUnitDayRatingUnauthorized ", 401)
}

func (o *GetOrganizationalUnitDayRatingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationalUnitDayRatingForbidden creates a GetOrganizationalUnitDayRatingForbidden with default headers values
func NewGetOrganizationalUnitDayRatingForbidden() *GetOrganizationalUnitDayRatingForbidden {
	return &GetOrganizationalUnitDayRatingForbidden{}
}

/*GetOrganizationalUnitDayRatingForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetOrganizationalUnitDayRatingForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetOrganizationalUnitDayRatingForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/dayRatings/{dayRatingIdOrDate}][%d] getOrganizationalUnitDayRatingForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationalUnitDayRatingForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetOrganizationalUnitDayRatingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationalUnitDayRatingNotFound creates a GetOrganizationalUnitDayRatingNotFound with default headers values
func NewGetOrganizationalUnitDayRatingNotFound() *GetOrganizationalUnitDayRatingNotFound {
	return &GetOrganizationalUnitDayRatingNotFound{}
}

/*GetOrganizationalUnitDayRatingNotFound handles this case with default header values.

Object not found
*/
type GetOrganizationalUnitDayRatingNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetOrganizationalUnitDayRatingNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/dayRatings/{dayRatingIdOrDate}][%d] getOrganizationalUnitDayRatingNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationalUnitDayRatingNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetOrganizationalUnitDayRatingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationalUnitDayRatingTooManyRequests creates a GetOrganizationalUnitDayRatingTooManyRequests with default headers values
func NewGetOrganizationalUnitDayRatingTooManyRequests() *GetOrganizationalUnitDayRatingTooManyRequests {
	return &GetOrganizationalUnitDayRatingTooManyRequests{}
}

/*GetOrganizationalUnitDayRatingTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetOrganizationalUnitDayRatingTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetOrganizationalUnitDayRatingTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/dayRatings/{dayRatingIdOrDate}][%d] getOrganizationalUnitDayRatingTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrganizationalUnitDayRatingTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetOrganizationalUnitDayRatingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
