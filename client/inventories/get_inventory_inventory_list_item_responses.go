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

// GetInventoryInventoryListItemReader is a Reader for the GetInventoryInventoryListItem structure.
type GetInventoryInventoryListItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryInventoryListItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryInventoryListItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInventoryInventoryListItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInventoryInventoryListItemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInventoryInventoryListItemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInventoryInventoryListItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetInventoryInventoryListItemMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetInventoryInventoryListItemTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInventoryInventoryListItemOK creates a GetInventoryInventoryListItemOK with default headers values
func NewGetInventoryInventoryListItemOK() *GetInventoryInventoryListItemOK {
	return &GetInventoryInventoryListItemOK{}
}

/*GetInventoryInventoryListItemOK handles this case with default header values.

successful operation
*/
type GetInventoryInventoryListItemOK struct {
	Payload *models.InventoryListItem
}

func (o *GetInventoryInventoryListItemOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items/{productId}][%d] getInventoryInventoryListItemOK  %+v", 200, o.Payload)
}

func (o *GetInventoryInventoryListItemOK) GetPayload() *models.InventoryListItem {
	return o.Payload
}

func (o *GetInventoryInventoryListItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryListItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListItemBadRequest creates a GetInventoryInventoryListItemBadRequest with default headers values
func NewGetInventoryInventoryListItemBadRequest() *GetInventoryInventoryListItemBadRequest {
	return &GetInventoryInventoryListItemBadRequest{}
}

/*GetInventoryInventoryListItemBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetInventoryInventoryListItemBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetInventoryInventoryListItemBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items/{productId}][%d] getInventoryInventoryListItemBadRequest  %+v", 400, o.Payload)
}

func (o *GetInventoryInventoryListItemBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetInventoryInventoryListItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListItemUnauthorized creates a GetInventoryInventoryListItemUnauthorized with default headers values
func NewGetInventoryInventoryListItemUnauthorized() *GetInventoryInventoryListItemUnauthorized {
	return &GetInventoryInventoryListItemUnauthorized{}
}

/*GetInventoryInventoryListItemUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetInventoryInventoryListItemUnauthorized struct {
}

func (o *GetInventoryInventoryListItemUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items/{productId}][%d] getInventoryInventoryListItemUnauthorized ", 401)
}

func (o *GetInventoryInventoryListItemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInventoryInventoryListItemForbidden creates a GetInventoryInventoryListItemForbidden with default headers values
func NewGetInventoryInventoryListItemForbidden() *GetInventoryInventoryListItemForbidden {
	return &GetInventoryInventoryListItemForbidden{}
}

/*GetInventoryInventoryListItemForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetInventoryInventoryListItemForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetInventoryInventoryListItemForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items/{productId}][%d] getInventoryInventoryListItemForbidden  %+v", 403, o.Payload)
}

func (o *GetInventoryInventoryListItemForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetInventoryInventoryListItemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListItemNotFound creates a GetInventoryInventoryListItemNotFound with default headers values
func NewGetInventoryInventoryListItemNotFound() *GetInventoryInventoryListItemNotFound {
	return &GetInventoryInventoryListItemNotFound{}
}

/*GetInventoryInventoryListItemNotFound handles this case with default header values.

Object not found
*/
type GetInventoryInventoryListItemNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetInventoryInventoryListItemNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items/{productId}][%d] getInventoryInventoryListItemNotFound  %+v", 404, o.Payload)
}

func (o *GetInventoryInventoryListItemNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetInventoryInventoryListItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListItemMethodNotAllowed creates a GetInventoryInventoryListItemMethodNotAllowed with default headers values
func NewGetInventoryInventoryListItemMethodNotAllowed() *GetInventoryInventoryListItemMethodNotAllowed {
	return &GetInventoryInventoryListItemMethodNotAllowed{}
}

/*GetInventoryInventoryListItemMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type GetInventoryInventoryListItemMethodNotAllowed struct {
	Payload *models.MethodNotAllowedError
}

func (o *GetInventoryInventoryListItemMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items/{productId}][%d] getInventoryInventoryListItemMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetInventoryInventoryListItemMethodNotAllowed) GetPayload() *models.MethodNotAllowedError {
	return o.Payload
}

func (o *GetInventoryInventoryListItemMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MethodNotAllowedError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListItemTooManyRequests creates a GetInventoryInventoryListItemTooManyRequests with default headers values
func NewGetInventoryInventoryListItemTooManyRequests() *GetInventoryInventoryListItemTooManyRequests {
	return &GetInventoryInventoryListItemTooManyRequests{}
}

/*GetInventoryInventoryListItemTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetInventoryInventoryListItemTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetInventoryInventoryListItemTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items/{productId}][%d] getInventoryInventoryListItemTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetInventoryInventoryListItemTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetInventoryInventoryListItemTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
