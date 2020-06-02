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

// UpdateInventoryInventoryListItemsReader is a Reader for the UpdateInventoryInventoryListItems structure.
type UpdateInventoryInventoryListItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInventoryInventoryListItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateInventoryInventoryListItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateInventoryInventoryListItemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateInventoryInventoryListItemsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateInventoryInventoryListItemsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateInventoryInventoryListItemsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewUpdateInventoryInventoryListItemsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateInventoryInventoryListItemsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateInventoryInventoryListItemsOK creates a UpdateInventoryInventoryListItemsOK with default headers values
func NewUpdateInventoryInventoryListItemsOK() *UpdateInventoryInventoryListItemsOK {
	return &UpdateInventoryInventoryListItemsOK{}
}

/*UpdateInventoryInventoryListItemsOK handles this case with default header values.

successful operation
*/
type UpdateInventoryInventoryListItemsOK struct {
	Payload []*models.AddOrUpdateResult
}

func (o *UpdateInventoryInventoryListItemsOK) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items][%d] updateInventoryInventoryListItemsOK  %+v", 200, o.Payload)
}

func (o *UpdateInventoryInventoryListItemsOK) GetPayload() []*models.AddOrUpdateResult {
	return o.Payload
}

func (o *UpdateInventoryInventoryListItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInventoryInventoryListItemsBadRequest creates a UpdateInventoryInventoryListItemsBadRequest with default headers values
func NewUpdateInventoryInventoryListItemsBadRequest() *UpdateInventoryInventoryListItemsBadRequest {
	return &UpdateInventoryInventoryListItemsBadRequest{}
}

/*UpdateInventoryInventoryListItemsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type UpdateInventoryInventoryListItemsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *UpdateInventoryInventoryListItemsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items][%d] updateInventoryInventoryListItemsBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateInventoryInventoryListItemsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *UpdateInventoryInventoryListItemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInventoryInventoryListItemsUnauthorized creates a UpdateInventoryInventoryListItemsUnauthorized with default headers values
func NewUpdateInventoryInventoryListItemsUnauthorized() *UpdateInventoryInventoryListItemsUnauthorized {
	return &UpdateInventoryInventoryListItemsUnauthorized{}
}

/*UpdateInventoryInventoryListItemsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type UpdateInventoryInventoryListItemsUnauthorized struct {
}

func (o *UpdateInventoryInventoryListItemsUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items][%d] updateInventoryInventoryListItemsUnauthorized ", 401)
}

func (o *UpdateInventoryInventoryListItemsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInventoryInventoryListItemsForbidden creates a UpdateInventoryInventoryListItemsForbidden with default headers values
func NewUpdateInventoryInventoryListItemsForbidden() *UpdateInventoryInventoryListItemsForbidden {
	return &UpdateInventoryInventoryListItemsForbidden{}
}

/*UpdateInventoryInventoryListItemsForbidden handles this case with default header values.

Requested action is not allowed
*/
type UpdateInventoryInventoryListItemsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *UpdateInventoryInventoryListItemsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items][%d] updateInventoryInventoryListItemsForbidden  %+v", 403, o.Payload)
}

func (o *UpdateInventoryInventoryListItemsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *UpdateInventoryInventoryListItemsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInventoryInventoryListItemsNotFound creates a UpdateInventoryInventoryListItemsNotFound with default headers values
func NewUpdateInventoryInventoryListItemsNotFound() *UpdateInventoryInventoryListItemsNotFound {
	return &UpdateInventoryInventoryListItemsNotFound{}
}

/*UpdateInventoryInventoryListItemsNotFound handles this case with default header values.

Object not found
*/
type UpdateInventoryInventoryListItemsNotFound struct {
	Payload *models.NotFoundError
}

func (o *UpdateInventoryInventoryListItemsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items][%d] updateInventoryInventoryListItemsNotFound  %+v", 404, o.Payload)
}

func (o *UpdateInventoryInventoryListItemsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *UpdateInventoryInventoryListItemsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInventoryInventoryListItemsMethodNotAllowed creates a UpdateInventoryInventoryListItemsMethodNotAllowed with default headers values
func NewUpdateInventoryInventoryListItemsMethodNotAllowed() *UpdateInventoryInventoryListItemsMethodNotAllowed {
	return &UpdateInventoryInventoryListItemsMethodNotAllowed{}
}

/*UpdateInventoryInventoryListItemsMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type UpdateInventoryInventoryListItemsMethodNotAllowed struct {
	Payload *models.MethodNotAllowedError
}

func (o *UpdateInventoryInventoryListItemsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items][%d] updateInventoryInventoryListItemsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *UpdateInventoryInventoryListItemsMethodNotAllowed) GetPayload() *models.MethodNotAllowedError {
	return o.Payload
}

func (o *UpdateInventoryInventoryListItemsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MethodNotAllowedError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInventoryInventoryListItemsTooManyRequests creates a UpdateInventoryInventoryListItemsTooManyRequests with default headers values
func NewUpdateInventoryInventoryListItemsTooManyRequests() *UpdateInventoryInventoryListItemsTooManyRequests {
	return &UpdateInventoryInventoryListItemsTooManyRequests{}
}

/*UpdateInventoryInventoryListItemsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type UpdateInventoryInventoryListItemsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *UpdateInventoryInventoryListItemsTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/inventories/{inventoryId}/inventoryLists/{inventoryListId}/items][%d] updateInventoryInventoryListItemsTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateInventoryInventoryListItemsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *UpdateInventoryInventoryListItemsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}