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

// GetOrganizationalUnitsReader is a Reader for the GetOrganizationalUnits structure.
type GetOrganizationalUnitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationalUnitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationalUnitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetOrganizationalUnitsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrganizationalUnitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrganizationalUnitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationalUnitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationalUnitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrganizationalUnitsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganizationalUnitsOK creates a GetOrganizationalUnitsOK with default headers values
func NewGetOrganizationalUnitsOK() *GetOrganizationalUnitsOK {
	return &GetOrganizationalUnitsOK{}
}

/*GetOrganizationalUnitsOK handles this case with default header values.

successful operation
*/
type GetOrganizationalUnitsOK struct {
	Payload *models.ResultListOrganizationalUnit
}

func (o *GetOrganizationalUnitsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits][%d] getOrganizationalUnitsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationalUnitsOK) GetPayload() *models.ResultListOrganizationalUnit {
	return o.Payload
}

func (o *GetOrganizationalUnitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResultListOrganizationalUnit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationalUnitsNoContent creates a GetOrganizationalUnitsNoContent with default headers values
func NewGetOrganizationalUnitsNoContent() *GetOrganizationalUnitsNoContent {
	return &GetOrganizationalUnitsNoContent{}
}

/*GetOrganizationalUnitsNoContent handles this case with default header values.

Request successful, but the list is empty. Either there is in general no object on the list or a set filter/restriction in querystring doesn't match any object.
*/
type GetOrganizationalUnitsNoContent struct {
}

func (o *GetOrganizationalUnitsNoContent) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits][%d] getOrganizationalUnitsNoContent ", 204)
}

func (o *GetOrganizationalUnitsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationalUnitsBadRequest creates a GetOrganizationalUnitsBadRequest with default headers values
func NewGetOrganizationalUnitsBadRequest() *GetOrganizationalUnitsBadRequest {
	return &GetOrganizationalUnitsBadRequest{}
}

/*GetOrganizationalUnitsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetOrganizationalUnitsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetOrganizationalUnitsBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits][%d] getOrganizationalUnitsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationalUnitsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetOrganizationalUnitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationalUnitsUnauthorized creates a GetOrganizationalUnitsUnauthorized with default headers values
func NewGetOrganizationalUnitsUnauthorized() *GetOrganizationalUnitsUnauthorized {
	return &GetOrganizationalUnitsUnauthorized{}
}

/*GetOrganizationalUnitsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetOrganizationalUnitsUnauthorized struct {
}

func (o *GetOrganizationalUnitsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits][%d] getOrganizationalUnitsUnauthorized ", 401)
}

func (o *GetOrganizationalUnitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationalUnitsForbidden creates a GetOrganizationalUnitsForbidden with default headers values
func NewGetOrganizationalUnitsForbidden() *GetOrganizationalUnitsForbidden {
	return &GetOrganizationalUnitsForbidden{}
}

/*GetOrganizationalUnitsForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetOrganizationalUnitsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetOrganizationalUnitsForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits][%d] getOrganizationalUnitsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationalUnitsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetOrganizationalUnitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationalUnitsNotFound creates a GetOrganizationalUnitsNotFound with default headers values
func NewGetOrganizationalUnitsNotFound() *GetOrganizationalUnitsNotFound {
	return &GetOrganizationalUnitsNotFound{}
}

/*GetOrganizationalUnitsNotFound handles this case with default header values.

Object not found
*/
type GetOrganizationalUnitsNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetOrganizationalUnitsNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits][%d] getOrganizationalUnitsNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationalUnitsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetOrganizationalUnitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationalUnitsTooManyRequests creates a GetOrganizationalUnitsTooManyRequests with default headers values
func NewGetOrganizationalUnitsTooManyRequests() *GetOrganizationalUnitsTooManyRequests {
	return &GetOrganizationalUnitsTooManyRequests{}
}

/*GetOrganizationalUnitsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetOrganizationalUnitsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetOrganizationalUnitsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/organizationalUnits][%d] getOrganizationalUnitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrganizationalUnitsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetOrganizationalUnitsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}