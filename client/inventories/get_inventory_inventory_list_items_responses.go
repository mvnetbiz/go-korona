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

// GetInventoryInventoryListItemsReader is a Reader for the GetInventoryInventoryListItems structure.
type GetInventoryInventoryListItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryInventoryListItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryInventoryListItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetInventoryInventoryListItemsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInventoryInventoryListItemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInventoryInventoryListItemsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInventoryInventoryListItemsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInventoryInventoryListItemsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetInventoryInventoryListItemsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetInventoryInventoryListItemsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInventoryInventoryListItemsOK creates a GetInventoryInventoryListItemsOK with default headers values
func NewGetInventoryInventoryListItemsOK() *GetInventoryInventoryListItemsOK {
	return &GetInventoryInventoryListItemsOK{}
}

/*GetInventoryInventoryListItemsOK handles this case with default header values.

successful operation
*/
type GetInventoryInventoryListItemsOK struct {
	Payload *models.ResultListInventoryListItem
}

func (o *GetInventoryInventoryListItemsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items][%d] getInventoryInventoryListItemsOK  %+v", 200, o.Payload)
}

func (o *GetInventoryInventoryListItemsOK) GetPayload() *models.ResultListInventoryListItem {
	return o.Payload
}

func (o *GetInventoryInventoryListItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResultListInventoryListItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListItemsNoContent creates a GetInventoryInventoryListItemsNoContent with default headers values
func NewGetInventoryInventoryListItemsNoContent() *GetInventoryInventoryListItemsNoContent {
	return &GetInventoryInventoryListItemsNoContent{}
}

/*GetInventoryInventoryListItemsNoContent handles this case with default header values.

Request successful, but the list is empty. Either there is in general no object on the list or a set filter/restriction in querystring doesn't match any object.
*/
type GetInventoryInventoryListItemsNoContent struct {
}

func (o *GetInventoryInventoryListItemsNoContent) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items][%d] getInventoryInventoryListItemsNoContent ", 204)
}

func (o *GetInventoryInventoryListItemsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInventoryInventoryListItemsBadRequest creates a GetInventoryInventoryListItemsBadRequest with default headers values
func NewGetInventoryInventoryListItemsBadRequest() *GetInventoryInventoryListItemsBadRequest {
	return &GetInventoryInventoryListItemsBadRequest{}
}

/*GetInventoryInventoryListItemsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetInventoryInventoryListItemsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetInventoryInventoryListItemsBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items][%d] getInventoryInventoryListItemsBadRequest  %+v", 400, o.Payload)
}

func (o *GetInventoryInventoryListItemsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetInventoryInventoryListItemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListItemsUnauthorized creates a GetInventoryInventoryListItemsUnauthorized with default headers values
func NewGetInventoryInventoryListItemsUnauthorized() *GetInventoryInventoryListItemsUnauthorized {
	return &GetInventoryInventoryListItemsUnauthorized{}
}

/*GetInventoryInventoryListItemsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetInventoryInventoryListItemsUnauthorized struct {
}

func (o *GetInventoryInventoryListItemsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items][%d] getInventoryInventoryListItemsUnauthorized ", 401)
}

func (o *GetInventoryInventoryListItemsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInventoryInventoryListItemsForbidden creates a GetInventoryInventoryListItemsForbidden with default headers values
func NewGetInventoryInventoryListItemsForbidden() *GetInventoryInventoryListItemsForbidden {
	return &GetInventoryInventoryListItemsForbidden{}
}

/*GetInventoryInventoryListItemsForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetInventoryInventoryListItemsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetInventoryInventoryListItemsForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items][%d] getInventoryInventoryListItemsForbidden  %+v", 403, o.Payload)
}

func (o *GetInventoryInventoryListItemsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetInventoryInventoryListItemsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListItemsNotFound creates a GetInventoryInventoryListItemsNotFound with default headers values
func NewGetInventoryInventoryListItemsNotFound() *GetInventoryInventoryListItemsNotFound {
	return &GetInventoryInventoryListItemsNotFound{}
}

/*GetInventoryInventoryListItemsNotFound handles this case with default header values.

Object not found
*/
type GetInventoryInventoryListItemsNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetInventoryInventoryListItemsNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items][%d] getInventoryInventoryListItemsNotFound  %+v", 404, o.Payload)
}

func (o *GetInventoryInventoryListItemsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetInventoryInventoryListItemsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListItemsMethodNotAllowed creates a GetInventoryInventoryListItemsMethodNotAllowed with default headers values
func NewGetInventoryInventoryListItemsMethodNotAllowed() *GetInventoryInventoryListItemsMethodNotAllowed {
	return &GetInventoryInventoryListItemsMethodNotAllowed{}
}

/*GetInventoryInventoryListItemsMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type GetInventoryInventoryListItemsMethodNotAllowed struct {
	Payload *models.MethodNotAllowedError
}

func (o *GetInventoryInventoryListItemsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items][%d] getInventoryInventoryListItemsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetInventoryInventoryListItemsMethodNotAllowed) GetPayload() *models.MethodNotAllowedError {
	return o.Payload
}

func (o *GetInventoryInventoryListItemsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MethodNotAllowedError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventoryInventoryListItemsTooManyRequests creates a GetInventoryInventoryListItemsTooManyRequests with default headers values
func NewGetInventoryInventoryListItemsTooManyRequests() *GetInventoryInventoryListItemsTooManyRequests {
	return &GetInventoryInventoryListItemsTooManyRequests{}
}

/*GetInventoryInventoryListItemsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetInventoryInventoryListItemsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetInventoryInventoryListItemsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items][%d] getInventoryInventoryListItemsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetInventoryInventoryListItemsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetInventoryInventoryListItemsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
