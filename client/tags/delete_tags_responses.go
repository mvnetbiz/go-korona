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

// DeleteTagsReader is a Reader for the DeleteTags structure.
type DeleteTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteTagsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteTagsOK creates a DeleteTagsOK with default headers values
func NewDeleteTagsOK() *DeleteTagsOK {
	return &DeleteTagsOK{}
}

/*DeleteTagsOK handles this case with default header values.

successful operation
*/
type DeleteTagsOK struct {
	Payload []*models.AddOrUpdateResult
}

func (o *DeleteTagsOK) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/tags][%d] deleteTagsOK  %+v", 200, o.Payload)
}

func (o *DeleteTagsOK) GetPayload() []*models.AddOrUpdateResult {
	return o.Payload
}

func (o *DeleteTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTagsBadRequest creates a DeleteTagsBadRequest with default headers values
func NewDeleteTagsBadRequest() *DeleteTagsBadRequest {
	return &DeleteTagsBadRequest{}
}

/*DeleteTagsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type DeleteTagsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *DeleteTagsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/tags][%d] deleteTagsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTagsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *DeleteTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTagsUnauthorized creates a DeleteTagsUnauthorized with default headers values
func NewDeleteTagsUnauthorized() *DeleteTagsUnauthorized {
	return &DeleteTagsUnauthorized{}
}

/*DeleteTagsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type DeleteTagsUnauthorized struct {
}

func (o *DeleteTagsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/tags][%d] deleteTagsUnauthorized ", 401)
}

func (o *DeleteTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTagsForbidden creates a DeleteTagsForbidden with default headers values
func NewDeleteTagsForbidden() *DeleteTagsForbidden {
	return &DeleteTagsForbidden{}
}

/*DeleteTagsForbidden handles this case with default header values.

Requested action is not allowed
*/
type DeleteTagsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *DeleteTagsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/tags][%d] deleteTagsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTagsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *DeleteTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTagsNotFound creates a DeleteTagsNotFound with default headers values
func NewDeleteTagsNotFound() *DeleteTagsNotFound {
	return &DeleteTagsNotFound{}
}

/*DeleteTagsNotFound handles this case with default header values.

Object not found
*/
type DeleteTagsNotFound struct {
	Payload *models.NotFoundError
}

func (o *DeleteTagsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/tags][%d] deleteTagsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTagsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *DeleteTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTagsTooManyRequests creates a DeleteTagsTooManyRequests with default headers values
func NewDeleteTagsTooManyRequests() *DeleteTagsTooManyRequests {
	return &DeleteTagsTooManyRequests{}
}

/*DeleteTagsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type DeleteTagsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *DeleteTagsTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/tags][%d] deleteTagsTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTagsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *DeleteTagsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
