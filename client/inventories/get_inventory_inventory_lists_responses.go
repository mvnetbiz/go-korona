// Code generated by go-swagger; DO NOT EDIT.

package inventories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// GetInventoryInventoryListsReader is a Reader for the GetInventoryInventoryLists structure.
type GetInventoryInventoryListsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryInventoryListsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryInventoryListsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetInventoryInventoryListsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInventoryInventoryListsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInventoryInventoryListsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInventoryInventoryListsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInventoryInventoryListsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetInventoryInventoryListsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetInventoryInventoryListsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInventoryInventoryListsOK creates a GetInventoryInventoryListsOK with default headers values
func NewGetInventoryInventoryListsOK() *GetInventoryInventoryListsOK {
	return &GetInventoryInventoryListsOK{}
}

/*GetInventoryInventoryListsOK handles this case with default header values.

successful operation
*/
type GetInventoryInventoryListsOK struct {
	Payload *models.ResultListInventoryList
}

func (o *GetInventoryInventoryListsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists][%d] getInventoryInventoryListsOK  %+v", 200, o.Payload)
}

func (o *GetInventoryInventoryListsOK) GetPayload() *models.ResultListInventoryList {
	return o.Payload
}

func (o *GetInventoryInventoryListsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResultListInventoryList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListsNoContent creates a GetInventoryInventoryListsNoContent with default headers values
func NewGetInventoryInventoryListsNoContent() *GetInventoryInventoryListsNoContent {
	return &GetInventoryInventoryListsNoContent{}
}

/*GetInventoryInventoryListsNoContent handles this case with default header values.

Request successful, but the list is empty. Either there is in general no object on the list or a set filter/restriction in querystring doesn't match any object.
*/
type GetInventoryInventoryListsNoContent struct {
}

func (o *GetInventoryInventoryListsNoContent) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists][%d] getInventoryInventoryListsNoContent ", 204)
}

func (o *GetInventoryInventoryListsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInventoryInventoryListsBadRequest creates a GetInventoryInventoryListsBadRequest with default headers values
func NewGetInventoryInventoryListsBadRequest() *GetInventoryInventoryListsBadRequest {
	return &GetInventoryInventoryListsBadRequest{}
}

/*GetInventoryInventoryListsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetInventoryInventoryListsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetInventoryInventoryListsBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists][%d] getInventoryInventoryListsBadRequest  %+v", 400, o.Payload)
}

func (o *GetInventoryInventoryListsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetInventoryInventoryListsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListsUnauthorized creates a GetInventoryInventoryListsUnauthorized with default headers values
func NewGetInventoryInventoryListsUnauthorized() *GetInventoryInventoryListsUnauthorized {
	return &GetInventoryInventoryListsUnauthorized{}
}

/*GetInventoryInventoryListsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetInventoryInventoryListsUnauthorized struct {
}

func (o *GetInventoryInventoryListsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists][%d] getInventoryInventoryListsUnauthorized ", 401)
}

func (o *GetInventoryInventoryListsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInventoryInventoryListsForbidden creates a GetInventoryInventoryListsForbidden with default headers values
func NewGetInventoryInventoryListsForbidden() *GetInventoryInventoryListsForbidden {
	return &GetInventoryInventoryListsForbidden{}
}

/*GetInventoryInventoryListsForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetInventoryInventoryListsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetInventoryInventoryListsForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists][%d] getInventoryInventoryListsForbidden  %+v", 403, o.Payload)
}

func (o *GetInventoryInventoryListsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetInventoryInventoryListsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListsNotFound creates a GetInventoryInventoryListsNotFound with default headers values
func NewGetInventoryInventoryListsNotFound() *GetInventoryInventoryListsNotFound {
	return &GetInventoryInventoryListsNotFound{}
}

/*GetInventoryInventoryListsNotFound handles this case with default header values.

Object not found
*/
type GetInventoryInventoryListsNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetInventoryInventoryListsNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists][%d] getInventoryInventoryListsNotFound  %+v", 404, o.Payload)
}

func (o *GetInventoryInventoryListsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetInventoryInventoryListsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListsMethodNotAllowed creates a GetInventoryInventoryListsMethodNotAllowed with default headers values
func NewGetInventoryInventoryListsMethodNotAllowed() *GetInventoryInventoryListsMethodNotAllowed {
	return &GetInventoryInventoryListsMethodNotAllowed{}
}

/*GetInventoryInventoryListsMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type GetInventoryInventoryListsMethodNotAllowed struct {
	Payload *models.MethodNotAllowedError
}

func (o *GetInventoryInventoryListsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists][%d] getInventoryInventoryListsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetInventoryInventoryListsMethodNotAllowed) GetPayload() *models.MethodNotAllowedError {
	return o.Payload
}

func (o *GetInventoryInventoryListsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MethodNotAllowedError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListsTooManyRequests creates a GetInventoryInventoryListsTooManyRequests with default headers values
func NewGetInventoryInventoryListsTooManyRequests() *GetInventoryInventoryListsTooManyRequests {
	return &GetInventoryInventoryListsTooManyRequests{}
}

/*GetInventoryInventoryListsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetInventoryInventoryListsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetInventoryInventoryListsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists][%d] getInventoryInventoryListsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetInventoryInventoryListsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetInventoryInventoryListsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
