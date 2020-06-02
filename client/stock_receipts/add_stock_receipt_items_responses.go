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

// AddStockReceiptItemsReader is a Reader for the AddStockReceiptItems structure.
type AddStockReceiptItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddStockReceiptItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddStockReceiptItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddStockReceiptItemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddStockReceiptItemsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddStockReceiptItemsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddStockReceiptItemsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewAddStockReceiptItemsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAddStockReceiptItemsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddStockReceiptItemsOK creates a AddStockReceiptItemsOK with default headers values
func NewAddStockReceiptItemsOK() *AddStockReceiptItemsOK {
	return &AddStockReceiptItemsOK{}
}

/*AddStockReceiptItemsOK handles this case with default header values.

successful operation
*/
type AddStockReceiptItemsOK struct {
	Payload []*models.AddOrUpdateResult
}

func (o *AddStockReceiptItemsOK) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/stockReceipts/{stockReceiptId}/items][%d] addStockReceiptItemsOK  %+v", 200, o.Payload)
}

func (o *AddStockReceiptItemsOK) GetPayload() []*models.AddOrUpdateResult {
	return o.Payload
}

func (o *AddStockReceiptItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddStockReceiptItemsBadRequest creates a AddStockReceiptItemsBadRequest with default headers values
func NewAddStockReceiptItemsBadRequest() *AddStockReceiptItemsBadRequest {
	return &AddStockReceiptItemsBadRequest{}
}

/*AddStockReceiptItemsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type AddStockReceiptItemsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *AddStockReceiptItemsBadRequest) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/stockReceipts/{stockReceiptId}/items][%d] addStockReceiptItemsBadRequest  %+v", 400, o.Payload)
}

func (o *AddStockReceiptItemsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *AddStockReceiptItemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddStockReceiptItemsUnauthorized creates a AddStockReceiptItemsUnauthorized with default headers values
func NewAddStockReceiptItemsUnauthorized() *AddStockReceiptItemsUnauthorized {
	return &AddStockReceiptItemsUnauthorized{}
}

/*AddStockReceiptItemsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type AddStockReceiptItemsUnauthorized struct {
}

func (o *AddStockReceiptItemsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/stockReceipts/{stockReceiptId}/items][%d] addStockReceiptItemsUnauthorized ", 401)
}

func (o *AddStockReceiptItemsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddStockReceiptItemsForbidden creates a AddStockReceiptItemsForbidden with default headers values
func NewAddStockReceiptItemsForbidden() *AddStockReceiptItemsForbidden {
	return &AddStockReceiptItemsForbidden{}
}

/*AddStockReceiptItemsForbidden handles this case with default header values.

Requested action is not allowed
*/
type AddStockReceiptItemsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *AddStockReceiptItemsForbidden) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/stockReceipts/{stockReceiptId}/items][%d] addStockReceiptItemsForbidden  %+v", 403, o.Payload)
}

func (o *AddStockReceiptItemsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *AddStockReceiptItemsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddStockReceiptItemsNotFound creates a AddStockReceiptItemsNotFound with default headers values
func NewAddStockReceiptItemsNotFound() *AddStockReceiptItemsNotFound {
	return &AddStockReceiptItemsNotFound{}
}

/*AddStockReceiptItemsNotFound handles this case with default header values.

Object not found
*/
type AddStockReceiptItemsNotFound struct {
	Payload *models.NotFoundError
}

func (o *AddStockReceiptItemsNotFound) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/stockReceipts/{stockReceiptId}/items][%d] addStockReceiptItemsNotFound  %+v", 404, o.Payload)
}

func (o *AddStockReceiptItemsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *AddStockReceiptItemsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddStockReceiptItemsMethodNotAllowed creates a AddStockReceiptItemsMethodNotAllowed with default headers values
func NewAddStockReceiptItemsMethodNotAllowed() *AddStockReceiptItemsMethodNotAllowed {
	return &AddStockReceiptItemsMethodNotAllowed{}
}

/*AddStockReceiptItemsMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type AddStockReceiptItemsMethodNotAllowed struct {
	Payload *models.MethodNotAllowedError
}

func (o *AddStockReceiptItemsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/stockReceipts/{stockReceiptId}/items][%d] addStockReceiptItemsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *AddStockReceiptItemsMethodNotAllowed) GetPayload() *models.MethodNotAllowedError {
	return o.Payload
}

func (o *AddStockReceiptItemsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MethodNotAllowedError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddStockReceiptItemsTooManyRequests creates a AddStockReceiptItemsTooManyRequests with default headers values
func NewAddStockReceiptItemsTooManyRequests() *AddStockReceiptItemsTooManyRequests {
	return &AddStockReceiptItemsTooManyRequests{}
}

/*AddStockReceiptItemsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type AddStockReceiptItemsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *AddStockReceiptItemsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/stockReceipts/{stockReceiptId}/items][%d] addStockReceiptItemsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AddStockReceiptItemsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *AddStockReceiptItemsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}