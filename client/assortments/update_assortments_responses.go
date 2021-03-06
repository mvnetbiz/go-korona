// Code generated by go-swagger; DO NOT EDIT.

package assortments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// UpdateAssortmentsReader is a Reader for the UpdateAssortments structure.
type UpdateAssortmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAssortmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAssortmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAssortmentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateAssortmentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAssortmentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAssortmentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateAssortmentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAssortmentsOK creates a UpdateAssortmentsOK with default headers values
func NewUpdateAssortmentsOK() *UpdateAssortmentsOK {
	return &UpdateAssortmentsOK{}
}

/*UpdateAssortmentsOK handles this case with default header values.

successful operation
*/
type UpdateAssortmentsOK struct {
	Payload []*models.AddOrUpdateResult
}

func (o *UpdateAssortmentsOK) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/assortments][%d] updateAssortmentsOK  %+v", 200, o.Payload)
}

func (o *UpdateAssortmentsOK) GetPayload() []*models.AddOrUpdateResult {
	return o.Payload
}

func (o *UpdateAssortmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAssortmentsBadRequest creates a UpdateAssortmentsBadRequest with default headers values
func NewUpdateAssortmentsBadRequest() *UpdateAssortmentsBadRequest {
	return &UpdateAssortmentsBadRequest{}
}

/*UpdateAssortmentsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type UpdateAssortmentsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *UpdateAssortmentsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/assortments][%d] updateAssortmentsBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAssortmentsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *UpdateAssortmentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAssortmentsUnauthorized creates a UpdateAssortmentsUnauthorized with default headers values
func NewUpdateAssortmentsUnauthorized() *UpdateAssortmentsUnauthorized {
	return &UpdateAssortmentsUnauthorized{}
}

/*UpdateAssortmentsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type UpdateAssortmentsUnauthorized struct {
}

func (o *UpdateAssortmentsUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/assortments][%d] updateAssortmentsUnauthorized ", 401)
}

func (o *UpdateAssortmentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAssortmentsForbidden creates a UpdateAssortmentsForbidden with default headers values
func NewUpdateAssortmentsForbidden() *UpdateAssortmentsForbidden {
	return &UpdateAssortmentsForbidden{}
}

/*UpdateAssortmentsForbidden handles this case with default header values.

Requested action is not allowed
*/
type UpdateAssortmentsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *UpdateAssortmentsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/assortments][%d] updateAssortmentsForbidden  %+v", 403, o.Payload)
}

func (o *UpdateAssortmentsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *UpdateAssortmentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAssortmentsNotFound creates a UpdateAssortmentsNotFound with default headers values
func NewUpdateAssortmentsNotFound() *UpdateAssortmentsNotFound {
	return &UpdateAssortmentsNotFound{}
}

/*UpdateAssortmentsNotFound handles this case with default header values.

Object not found
*/
type UpdateAssortmentsNotFound struct {
	Payload *models.NotFoundError
}

func (o *UpdateAssortmentsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/assortments][%d] updateAssortmentsNotFound  %+v", 404, o.Payload)
}

func (o *UpdateAssortmentsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *UpdateAssortmentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAssortmentsTooManyRequests creates a UpdateAssortmentsTooManyRequests with default headers values
func NewUpdateAssortmentsTooManyRequests() *UpdateAssortmentsTooManyRequests {
	return &UpdateAssortmentsTooManyRequests{}
}

/*UpdateAssortmentsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type UpdateAssortmentsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *UpdateAssortmentsTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{koronaAccountId}/assortments][%d] updateAssortmentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateAssortmentsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *UpdateAssortmentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
