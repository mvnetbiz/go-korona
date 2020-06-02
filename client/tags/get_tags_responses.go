// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// GetTagsReader is a Reader for the GetTags structure.
type GetTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetTagsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTagsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTagsOK creates a GetTagsOK with default headers values
func NewGetTagsOK() *GetTagsOK {
	return &GetTagsOK{}
}

/*GetTagsOK handles this case with default header values.

successful operation
*/
type GetTagsOK struct {
	Payload *models.ResultListTag
}

func (o *GetTagsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/tags][%d] getTagsOK  %+v", 200, o.Payload)
}

func (o *GetTagsOK) GetPayload() *models.ResultListTag {
	return o.Payload
}

func (o *GetTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResultListTag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagsNoContent creates a GetTagsNoContent with default headers values
func NewGetTagsNoContent() *GetTagsNoContent {
	return &GetTagsNoContent{}
}

/*GetTagsNoContent handles this case with default header values.

Request successful, but the list is empty. Either there is in general no object on the list or a set filter/restriction in querystring doesn't match any object.
*/
type GetTagsNoContent struct {
}

func (o *GetTagsNoContent) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/tags][%d] getTagsNoContent ", 204)
}

func (o *GetTagsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTagsBadRequest creates a GetTagsBadRequest with default headers values
func NewGetTagsBadRequest() *GetTagsBadRequest {
	return &GetTagsBadRequest{}
}

/*GetTagsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type GetTagsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/tags][%d] getTagsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTagsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagsUnauthorized creates a GetTagsUnauthorized with default headers values
func NewGetTagsUnauthorized() *GetTagsUnauthorized {
	return &GetTagsUnauthorized{}
}

/*GetTagsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type GetTagsUnauthorized struct {
}

func (o *GetTagsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/tags][%d] getTagsUnauthorized ", 401)
}

func (o *GetTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTagsForbidden creates a GetTagsForbidden with default headers values
func NewGetTagsForbidden() *GetTagsForbidden {
	return &GetTagsForbidden{}
}

/*GetTagsForbidden handles this case with default header values.

Requested action is not allowed
*/
type GetTagsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *GetTagsForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/tags][%d] getTagsForbidden  %+v", 403, o.Payload)
}

func (o *GetTagsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *GetTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagsNotFound creates a GetTagsNotFound with default headers values
func NewGetTagsNotFound() *GetTagsNotFound {
	return &GetTagsNotFound{}
}

/*GetTagsNotFound handles this case with default header values.

Object not found
*/
type GetTagsNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/tags][%d] getTagsNotFound  %+v", 404, o.Payload)
}

func (o *GetTagsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagsTooManyRequests creates a GetTagsTooManyRequests with default headers values
func NewGetTagsTooManyRequests() *GetTagsTooManyRequests {
	return &GetTagsTooManyRequests{}
}

/*GetTagsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type GetTagsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *GetTagsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{koronaAccountId}/tags][%d] getTagsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTagsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *GetTagsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}