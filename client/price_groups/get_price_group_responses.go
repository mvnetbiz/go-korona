// Code generated by go-swagger; DO NOT EDIT.

package price_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// GetPriceGroupReader is a Reader for the GetPriceGroup structure.
type GetPriceGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPriceGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPriceGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPriceGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPriceGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPriceGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPriceGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPriceGroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPriceGroupOK creates a GetPriceGroupOK with default headers values
func NewGetPriceGroupOK() *GetPriceGroupOK {
	return &GetPriceGroupOK{}
}

/*GetPriceGroupOK handles this case with default header values.

successful operation
*/
type GetPriceGroupOK struct {
	Payload *models.PriceGroup
}

func (o *GetPriceGroupOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/priceGroups/{priceGroupId}][%d] getPriceGroupOK  %+v", 200, o.Payload)
}

func (o *GetPriceGroupOK) GetPayload() *models.PriceGroup {
	return o.Payload
}

func (o *GetPriceGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PriceGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPriceGroupBadRequest creates a GetPriceGroupBadRequest with default headers values
func NewGetPriceGroupBadRequest() *GetPriceGroupBadRequest {
	return &GetPriceGroupBadRequest{}
}

/*GetPriceGroupBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetPriceGroupBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetPriceGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/priceGroups/{priceGroupId}][%d] getPriceGroupBadRequest  %+v", 400, o.Payload)
}

func (o *GetPriceGroupBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetPriceGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPriceGroupUnauthorized creates a GetPriceGroupUnauthorized with default headers values
func NewGetPriceGroupUnauthorized() *GetPriceGroupUnauthorized {
	return &GetPriceGroupUnauthorized{}
}

/*GetPriceGroupUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetPriceGroupUnauthorized struct {
}

func (o *GetPriceGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/priceGroups/{priceGroupId}][%d] getPriceGroupUnauthorized ", 401)
}

func (o *GetPriceGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPriceGroupForbidden creates a GetPriceGroupForbidden with default headers values
func NewGetPriceGroupForbidden() *GetPriceGroupForbidden {
	return &GetPriceGroupForbidden{}
}

/*GetPriceGroupForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetPriceGroupForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetPriceGroupForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/priceGroups/{priceGroupId}][%d] getPriceGroupForbidden  %+v", 403, o.Payload)
}

func (o *GetPriceGroupForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetPriceGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPriceGroupNotFound creates a GetPriceGroupNotFound with default headers values
func NewGetPriceGroupNotFound() *GetPriceGroupNotFound {
	return &GetPriceGroupNotFound{}
}

/*GetPriceGroupNotFound handles this case with default header values.

Object not found
*/
type GetPriceGroupNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetPriceGroupNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/priceGroups/{priceGroupId}][%d] getPriceGroupNotFound  %+v", 404, o.Payload)
}

func (o *GetPriceGroupNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetPriceGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPriceGroupTooManyRequests creates a GetPriceGroupTooManyRequests with default headers values
func NewGetPriceGroupTooManyRequests() *GetPriceGroupTooManyRequests {
	return &GetPriceGroupTooManyRequests{}
}

/*GetPriceGroupTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetPriceGroupTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetPriceGroupTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/priceGroups/{priceGroupId}][%d] getPriceGroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPriceGroupTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetPriceGroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
