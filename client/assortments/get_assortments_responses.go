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

// GetAssortmentsReader is a Reader for the GetAssortments structure.
type GetAssortmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssortmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAssortmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetAssortmentsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAssortmentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAssortmentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAssortmentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAssortmentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAssortmentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAssortmentsOK creates a GetAssortmentsOK with default headers values
func NewGetAssortmentsOK() *GetAssortmentsOK {
	return &GetAssortmentsOK{}
}

/*GetAssortmentsOK handles this case with default header values.

successful operation
*/
type GetAssortmentsOK struct {
	Payload *models.ResultListAssortment
}

func (o *GetAssortmentsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/assortments][%d] getAssortmentsOK  %+v", 200, o.Payload)
}

func (o *GetAssortmentsOK) GetPayload() *models.ResultListAssortment {
	return o.Payload
}

func (o *GetAssortmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResultListAssortment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssortmentsNoContent creates a GetAssortmentsNoContent with default headers values
func NewGetAssortmentsNoContent() *GetAssortmentsNoContent {
	return &GetAssortmentsNoContent{}
}

/*GetAssortmentsNoContent handles this case with default header values.

Request successful, but the list is empty. Either there is in general no object on the list or a set filter/restriction in querystring doesn't match any object.
*/
type GetAssortmentsNoContent struct {
}

func (o *GetAssortmentsNoContent) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/assortments][%d] getAssortmentsNoContent ", 204)
}

func (o *GetAssortmentsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAssortmentsBadRequest creates a GetAssortmentsBadRequest with default headers values
func NewGetAssortmentsBadRequest() *GetAssortmentsBadRequest {
	return &GetAssortmentsBadRequest{}
}

/*GetAssortmentsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetAssortmentsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetAssortmentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/assortments][%d] getAssortmentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAssortmentsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetAssortmentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssortmentsUnauthorized creates a GetAssortmentsUnauthorized with default headers values
func NewGetAssortmentsUnauthorized() *GetAssortmentsUnauthorized {
	return &GetAssortmentsUnauthorized{}
}

/*GetAssortmentsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetAssortmentsUnauthorized struct {
}

func (o *GetAssortmentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/assortments][%d] getAssortmentsUnauthorized ", 401)
}

func (o *GetAssortmentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAssortmentsForbidden creates a GetAssortmentsForbidden with default headers values
func NewGetAssortmentsForbidden() *GetAssortmentsForbidden {
	return &GetAssortmentsForbidden{}
}

/*GetAssortmentsForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetAssortmentsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetAssortmentsForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/assortments][%d] getAssortmentsForbidden  %+v", 403, o.Payload)
}

func (o *GetAssortmentsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetAssortmentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssortmentsNotFound creates a GetAssortmentsNotFound with default headers values
func NewGetAssortmentsNotFound() *GetAssortmentsNotFound {
	return &GetAssortmentsNotFound{}
}

/*GetAssortmentsNotFound handles this case with default header values.

Object not found
*/
type GetAssortmentsNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetAssortmentsNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/assortments][%d] getAssortmentsNotFound  %+v", 404, o.Payload)
}

func (o *GetAssortmentsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetAssortmentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssortmentsTooManyRequests creates a GetAssortmentsTooManyRequests with default headers values
func NewGetAssortmentsTooManyRequests() *GetAssortmentsTooManyRequests {
	return &GetAssortmentsTooManyRequests{}
}

/*GetAssortmentsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetAssortmentsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetAssortmentsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/assortments][%d] getAssortmentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAssortmentsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetAssortmentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
