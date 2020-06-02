// Code generated by go-swagger; DO NOT EDIT.

package stock_receipts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// UpdateStockReceiptItemsReader is a Reader for the UpdateStockReceiptItems structure.
type UpdateStockReceiptItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateStockReceiptItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateStockReceiptItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateStockReceiptItemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateStockReceiptItemsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateStockReceiptItemsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateStockReceiptItemsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewUpdateStockReceiptItemsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateStockReceiptItemsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateStockReceiptItemsOK creates a UpdateStockReceiptItemsOK with default headers values
func NewUpdateStockReceiptItemsOK() *UpdateStockReceiptItemsOK {
	return &UpdateStockReceiptItemsOK{}
}

/*UpdateStockReceiptItemsOK handles this case with default header values.

successful operation
*/
type UpdateStockReceiptItemsOK struct {
	Payload []*models.AddOrUpdateResult
}

func (o *UpdateStockReceiptItemsOK) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/stockReceipts/{stockReceiptId}/items][%d] updateStockReceiptItemsOK  %+v", 200, o.Payload)
}

func (o *UpdateStockReceiptItemsOK) GetPayload() []*models.AddOrUpdateResult {
	return o.Payload
}

func (o *UpdateStockReceiptItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateStockReceiptItemsBadRequest creates a UpdateStockReceiptItemsBadRequest with default headers values
func NewUpdateStockReceiptItemsBadRequest() *UpdateStockReceiptItemsBadRequest {
	return &UpdateStockReceiptItemsBadRequest{}
}

/*UpdateStockReceiptItemsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type UpdateStockReceiptItemsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *UpdateStockReceiptItemsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/stockReceipts/{stockReceiptId}/items][%d] updateStockReceiptItemsBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateStockReceiptItemsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *UpdateStockReceiptItemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateStockReceiptItemsUnauthorized creates a UpdateStockReceiptItemsUnauthorized with default headers values
func NewUpdateStockReceiptItemsUnauthorized() *UpdateStockReceiptItemsUnauthorized {
	return &UpdateStockReceiptItemsUnauthorized{}
}

/*UpdateStockReceiptItemsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type UpdateStockReceiptItemsUnauthorized struct {
}

func (o *UpdateStockReceiptItemsUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/stockReceipts/{stockReceiptId}/items][%d] updateStockReceiptItemsUnauthorized ", 401)
}

func (o *UpdateStockReceiptItemsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateStockReceiptItemsForbidden creates a UpdateStockReceiptItemsForbidden with default headers values
func NewUpdateStockReceiptItemsForbidden() *UpdateStockReceiptItemsForbidden {
	return &UpdateStockReceiptItemsForbidden{}
}

/*UpdateStockReceiptItemsForbidden handles this case with default header values.

Requested action is not allowed
*/
type UpdateStockReceiptItemsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *UpdateStockReceiptItemsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/stockReceipts/{stockReceiptId}/items][%d] updateStockReceiptItemsForbidden  %+v", 403, o.Payload)
}

func (o *UpdateStockReceiptItemsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *UpdateStockReceiptItemsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateStockReceiptItemsNotFound creates a UpdateStockReceiptItemsNotFound with default headers values
func NewUpdateStockReceiptItemsNotFound() *UpdateStockReceiptItemsNotFound {
	return &UpdateStockReceiptItemsNotFound{}
}

/*UpdateStockReceiptItemsNotFound handles this case with default header values.

Object not found
*/
type UpdateStockReceiptItemsNotFound struct {
	Payload *models.NotFoundError
}

func (o *UpdateStockReceiptItemsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/stockReceipts/{stockReceiptId}/items][%d] updateStockReceiptItemsNotFound  %+v", 404, o.Payload)
}

func (o *UpdateStockReceiptItemsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *UpdateStockReceiptItemsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateStockReceiptItemsMethodNotAllowed creates a UpdateStockReceiptItemsMethodNotAllowed with default headers values
func NewUpdateStockReceiptItemsMethodNotAllowed() *UpdateStockReceiptItemsMethodNotAllowed {
	return &UpdateStockReceiptItemsMethodNotAllowed{}
}

/*UpdateStockReceiptItemsMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type UpdateStockReceiptItemsMethodNotAllowed struct {
	Payload *models.MethodNotAllowedError
}

func (o *UpdateStockReceiptItemsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/stockReceipts/{stockReceiptId}/items][%d] updateStockReceiptItemsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *UpdateStockReceiptItemsMethodNotAllowed) GetPayload() *models.MethodNotAllowedError {
	return o.Payload
}

func (o *UpdateStockReceiptItemsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MethodNotAllowedError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateStockReceiptItemsTooManyRequests creates a UpdateStockReceiptItemsTooManyRequests with default headers values
func NewUpdateStockReceiptItemsTooManyRequests() *UpdateStockReceiptItemsTooManyRequests {
	return &UpdateStockReceiptItemsTooManyRequests{}
}

/*UpdateStockReceiptItemsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type UpdateStockReceiptItemsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *UpdateStockReceiptItemsTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/stockReceipts/{stockReceiptId}/items][%d] updateStockReceiptItemsTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateStockReceiptItemsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *UpdateStockReceiptItemsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
