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

// GetOrganizationalUnitDayRatingsReader is a Reader for the GetOrganizationalUnitDayRatings structure.
type GetOrganizationalUnitDayRatingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationalUnitDayRatingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationalUnitDayRatingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrganizationalUnitDayRatingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrganizationalUnitDayRatingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationalUnitDayRatingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationalUnitDayRatingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrganizationalUnitDayRatingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganizationalUnitDayRatingsOK creates a GetOrganizationalUnitDayRatingsOK with default headers values
func NewGetOrganizationalUnitDayRatingsOK() *GetOrganizationalUnitDayRatingsOK {
	return &GetOrganizationalUnitDayRatingsOK{}
}

/*GetOrganizationalUnitDayRatingsOK handles this case with default header values.

successful operation
*/
type GetOrganizationalUnitDayRatingsOK struct {
	Payload *models.ResultListDayRating
}

func (o *GetOrganizationalUnitDayRatingsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/dayRatings][%d] getOrganizationalUnitDayRatingsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationalUnitDayRatingsOK) GetPayload() *models.ResultListDayRating {
	return o.Payload
}

func (o *GetOrganizationalUnitDayRatingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResultListDayRating)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationalUnitDayRatingsBadRequest creates a GetOrganizationalUnitDayRatingsBadRequest with default headers values
func NewGetOrganizationalUnitDayRatingsBadRequest() *GetOrganizationalUnitDayRatingsBadRequest {
	return &GetOrganizationalUnitDayRatingsBadRequest{}
}

/*GetOrganizationalUnitDayRatingsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetOrganizationalUnitDayRatingsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetOrganizationalUnitDayRatingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/dayRatings][%d] getOrganizationalUnitDayRatingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationalUnitDayRatingsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetOrganizationalUnitDayRatingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationalUnitDayRatingsUnauthorized creates a GetOrganizationalUnitDayRatingsUnauthorized with default headers values
func NewGetOrganizationalUnitDayRatingsUnauthorized() *GetOrganizationalUnitDayRatingsUnauthorized {
	return &GetOrganizationalUnitDayRatingsUnauthorized{}
}

/*GetOrganizationalUnitDayRatingsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetOrganizationalUnitDayRatingsUnauthorized struct {
}

func (o *GetOrganizationalUnitDayRatingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/dayRatings][%d] getOrganizationalUnitDayRatingsUnauthorized ", 401)
}

func (o *GetOrganizationalUnitDayRatingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationalUnitDayRatingsForbidden creates a GetOrganizationalUnitDayRatingsForbidden with default headers values
func NewGetOrganizationalUnitDayRatingsForbidden() *GetOrganizationalUnitDayRatingsForbidden {
	return &GetOrganizationalUnitDayRatingsForbidden{}
}

/*GetOrganizationalUnitDayRatingsForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetOrganizationalUnitDayRatingsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetOrganizationalUnitDayRatingsForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/dayRatings][%d] getOrganizationalUnitDayRatingsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationalUnitDayRatingsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetOrganizationalUnitDayRatingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationalUnitDayRatingsNotFound creates a GetOrganizationalUnitDayRatingsNotFound with default headers values
func NewGetOrganizationalUnitDayRatingsNotFound() *GetOrganizationalUnitDayRatingsNotFound {
	return &GetOrganizationalUnitDayRatingsNotFound{}
}

/*GetOrganizationalUnitDayRatingsNotFound handles this case with default header values.

Object not found
*/
type GetOrganizationalUnitDayRatingsNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetOrganizationalUnitDayRatingsNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/dayRatings][%d] getOrganizationalUnitDayRatingsNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationalUnitDayRatingsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetOrganizationalUnitDayRatingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationalUnitDayRatingsTooManyRequests creates a GetOrganizationalUnitDayRatingsTooManyRequests with default headers values
func NewGetOrganizationalUnitDayRatingsTooManyRequests() *GetOrganizationalUnitDayRatingsTooManyRequests {
	return &GetOrganizationalUnitDayRatingsTooManyRequests{}
}

/*GetOrganizationalUnitDayRatingsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetOrganizationalUnitDayRatingsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetOrganizationalUnitDayRatingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/dayRatings][%d] getOrganizationalUnitDayRatingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrganizationalUnitDayRatingsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetOrganizationalUnitDayRatingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
