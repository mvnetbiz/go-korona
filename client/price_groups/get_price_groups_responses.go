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

// GetPriceGroupsReader is a Reader for the GetPriceGroups structure.
type GetPriceGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPriceGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPriceGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetPriceGroupsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPriceGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPriceGroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPriceGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPriceGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPriceGroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPriceGroupsOK creates a GetPriceGroupsOK with default headers values
func NewGetPriceGroupsOK() *GetPriceGroupsOK {
	return &GetPriceGroupsOK{}
}

/*GetPriceGroupsOK handles this case with default header values.

successful operation
*/
type GetPriceGroupsOK struct {
	Payload *models.ResultListPriceGroup
}

func (o *GetPriceGroupsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/priceGroups][%d] getPriceGroupsOK  %+v", 200, o.Payload)
}

func (o *GetPriceGroupsOK) GetPayload() *models.ResultListPriceGroup {
	return o.Payload
}

func (o *GetPriceGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResultListPriceGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPriceGroupsNoContent creates a GetPriceGroupsNoContent with default headers values
func NewGetPriceGroupsNoContent() *GetPriceGroupsNoContent {
	return &GetPriceGroupsNoContent{}
}

/*GetPriceGroupsNoContent handles this case with default header values.

Request successful, but the list is empty. Either there is in general no object on the list or a set filter/restriction in querystring doesn't match any object.
*/
type GetPriceGroupsNoContent struct {
}

func (o *GetPriceGroupsNoContent) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/priceGroups][%d] getPriceGroupsNoContent ", 204)
}

func (o *GetPriceGroupsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPriceGroupsBadRequest creates a GetPriceGroupsBadRequest with default headers values
func NewGetPriceGroupsBadRequest() *GetPriceGroupsBadRequest {
	return &GetPriceGroupsBadRequest{}
}

/*GetPriceGroupsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetPriceGroupsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetPriceGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/priceGroups][%d] getPriceGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetPriceGroupsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetPriceGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPriceGroupsUnauthorized creates a GetPriceGroupsUnauthorized with default headers values
func NewGetPriceGroupsUnauthorized() *GetPriceGroupsUnauthorized {
	return &GetPriceGroupsUnauthorized{}
}

/*GetPriceGroupsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetPriceGroupsUnauthorized struct {
}

func (o *GetPriceGroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/priceGroups][%d] getPriceGroupsUnauthorized ", 401)
}

func (o *GetPriceGroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPriceGroupsForbidden creates a GetPriceGroupsForbidden with default headers values
func NewGetPriceGroupsForbidden() *GetPriceGroupsForbidden {
	return &GetPriceGroupsForbidden{}
}

/*GetPriceGroupsForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetPriceGroupsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetPriceGroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/priceGroups][%d] getPriceGroupsForbidden  %+v", 403, o.Payload)
}

func (o *GetPriceGroupsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetPriceGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPriceGroupsNotFound creates a GetPriceGroupsNotFound with default headers values
func NewGetPriceGroupsNotFound() *GetPriceGroupsNotFound {
	return &GetPriceGroupsNotFound{}
}

/*GetPriceGroupsNotFound handles this case with default header values.

Object not found
*/
type GetPriceGroupsNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetPriceGroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/priceGroups][%d] getPriceGroupsNotFound  %+v", 404, o.Payload)
}

func (o *GetPriceGroupsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetPriceGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPriceGroupsTooManyRequests creates a GetPriceGroupsTooManyRequests with default headers values
func NewGetPriceGroupsTooManyRequests() *GetPriceGroupsTooManyRequests {
	return &GetPriceGroupsTooManyRequests{}
}

/*GetPriceGroupsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetPriceGroupsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetPriceGroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/priceGroups][%d] getPriceGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPriceGroupsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetPriceGroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}