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

// GetInventoryInventoryListReader is a Reader for the GetInventoryInventoryList structure.
type GetInventoryInventoryListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryInventoryListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryInventoryListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInventoryInventoryListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInventoryInventoryListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInventoryInventoryListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInventoryInventoryListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetInventoryInventoryListMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetInventoryInventoryListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInventoryInventoryListOK creates a GetInventoryInventoryListOK with default headers values
func NewGetInventoryInventoryListOK() *GetInventoryInventoryListOK {
	return &GetInventoryInventoryListOK{}
}

/*GetInventoryInventoryListOK handles this case with default header values.

successful operation
*/
type GetInventoryInventoryListOK struct {
	Payload *models.InventoryList
}

func (o *GetInventoryInventoryListOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}][%d] getInventoryInventoryListOK  %+v", 200, o.Payload)
}

func (o *GetInventoryInventoryListOK) GetPayload() *models.InventoryList {
	return o.Payload
}

func (o *GetInventoryInventoryListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListBadRequest creates a GetInventoryInventoryListBadRequest with default headers values
func NewGetInventoryInventoryListBadRequest() *GetInventoryInventoryListBadRequest {
	return &GetInventoryInventoryListBadRequest{}
}

/*GetInventoryInventoryListBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetInventoryInventoryListBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetInventoryInventoryListBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}][%d] getInventoryInventoryListBadRequest  %+v", 400, o.Payload)
}

func (o *GetInventoryInventoryListBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetInventoryInventoryListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListUnauthorized creates a GetInventoryInventoryListUnauthorized with default headers values
func NewGetInventoryInventoryListUnauthorized() *GetInventoryInventoryListUnauthorized {
	return &GetInventoryInventoryListUnauthorized{}
}

/*GetInventoryInventoryListUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetInventoryInventoryListUnauthorized struct {
}

func (o *GetInventoryInventoryListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}][%d] getInventoryInventoryListUnauthorized ", 401)
}

func (o *GetInventoryInventoryListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInventoryInventoryListForbidden creates a GetInventoryInventoryListForbidden with default headers values
func NewGetInventoryInventoryListForbidden() *GetInventoryInventoryListForbidden {
	return &GetInventoryInventoryListForbidden{}
}

/*GetInventoryInventoryListForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetInventoryInventoryListForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetInventoryInventoryListForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}][%d] getInventoryInventoryListForbidden  %+v", 403, o.Payload)
}

func (o *GetInventoryInventoryListForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetInventoryInventoryListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListNotFound creates a GetInventoryInventoryListNotFound with default headers values
func NewGetInventoryInventoryListNotFound() *GetInventoryInventoryListNotFound {
	return &GetInventoryInventoryListNotFound{}
}

/*GetInventoryInventoryListNotFound handles this case with default header values.

Object not found
*/
type GetInventoryInventoryListNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetInventoryInventoryListNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}][%d] getInventoryInventoryListNotFound  %+v", 404, o.Payload)
}

func (o *GetInventoryInventoryListNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetInventoryInventoryListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListMethodNotAllowed creates a GetInventoryInventoryListMethodNotAllowed with default headers values
func NewGetInventoryInventoryListMethodNotAllowed() *GetInventoryInventoryListMethodNotAllowed {
	return &GetInventoryInventoryListMethodNotAllowed{}
}

/*GetInventoryInventoryListMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type GetInventoryInventoryListMethodNotAllowed struct {
	Payload *models.MethodNotAllowedError
}

func (o *GetInventoryInventoryListMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}][%d] getInventoryInventoryListMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetInventoryInventoryListMethodNotAllowed) GetPayload() *models.MethodNotAllowedError {
	return o.Payload
}

func (o *GetInventoryInventoryListMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MethodNotAllowedError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListTooManyRequests creates a GetInventoryInventoryListTooManyRequests with default headers values
func NewGetInventoryInventoryListTooManyRequests() *GetInventoryInventoryListTooManyRequests {
	return &GetInventoryInventoryListTooManyRequests{}
}

/*GetInventoryInventoryListTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetInventoryInventoryListTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetInventoryInventoryListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}][%d] getInventoryInventoryListTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetInventoryInventoryListTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetInventoryInventoryListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}