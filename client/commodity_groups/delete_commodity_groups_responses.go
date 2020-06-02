// Code generated by go-swagger; DO NOT EDIT.

package commodity_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// DeleteCommodityGroupsReader is a Reader for the DeleteCommodityGroups structure.
type DeleteCommodityGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCommodityGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCommodityGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCommodityGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteCommodityGroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCommodityGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCommodityGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteCommodityGroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCommodityGroupsOK creates a DeleteCommodityGroupsOK with default headers values
func NewDeleteCommodityGroupsOK() *DeleteCommodityGroupsOK {
	return &DeleteCommodityGroupsOK{}
}

/*DeleteCommodityGroupsOK handles this case with default header values.

successful operation
*/
type DeleteCommodityGroupsOK struct {
	Payload []*models.AddOrUpdateResult
}

func (o *DeleteCommodityGroupsOK) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/commodityGroups][%d] deleteCommodityGroupsOK  %+v", 200, o.Payload)
}

func (o *DeleteCommodityGroupsOK) GetPayload() []*models.AddOrUpdateResult {
	return o.Payload
}

func (o *DeleteCommodityGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCommodityGroupsBadRequest creates a DeleteCommodityGroupsBadRequest with default headers values
func NewDeleteCommodityGroupsBadRequest() *DeleteCommodityGroupsBadRequest {
	return &DeleteCommodityGroupsBadRequest{}
}

/*DeleteCommodityGroupsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type DeleteCommodityGroupsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *DeleteCommodityGroupsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/commodityGroups][%d] deleteCommodityGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCommodityGroupsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *DeleteCommodityGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCommodityGroupsUnauthorized creates a DeleteCommodityGroupsUnauthorized with default headers values
func NewDeleteCommodityGroupsUnauthorized() *DeleteCommodityGroupsUnauthorized {
	return &DeleteCommodityGroupsUnauthorized{}
}

/*DeleteCommodityGroupsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type DeleteCommodityGroupsUnauthorized struct {
}

func (o *DeleteCommodityGroupsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/commodityGroups][%d] deleteCommodityGroupsUnauthorized ", 401)
}

func (o *DeleteCommodityGroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCommodityGroupsForbidden creates a DeleteCommodityGroupsForbidden with default headers values
func NewDeleteCommodityGroupsForbidden() *DeleteCommodityGroupsForbidden {
	return &DeleteCommodityGroupsForbidden{}
}

/*DeleteCommodityGroupsForbidden handles this case with default header values.

Requested action is not allowed
*/
type DeleteCommodityGroupsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *DeleteCommodityGroupsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/commodityGroups][%d] deleteCommodityGroupsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCommodityGroupsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *DeleteCommodityGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCommodityGroupsNotFound creates a DeleteCommodityGroupsNotFound with default headers values
func NewDeleteCommodityGroupsNotFound() *DeleteCommodityGroupsNotFound {
	return &DeleteCommodityGroupsNotFound{}
}

/*DeleteCommodityGroupsNotFound handles this case with default header values.

Object not found
*/
type DeleteCommodityGroupsNotFound struct {
	Payload *models.NotFoundError
}

func (o *DeleteCommodityGroupsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/commodityGroups][%d] deleteCommodityGroupsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCommodityGroupsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *DeleteCommodityGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCommodityGroupsTooManyRequests creates a DeleteCommodityGroupsTooManyRequests with default headers values
func NewDeleteCommodityGroupsTooManyRequests() *DeleteCommodityGroupsTooManyRequests {
	return &DeleteCommodityGroupsTooManyRequests{}
}

/*DeleteCommodityGroupsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type DeleteCommodityGroupsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *DeleteCommodityGroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{koronaAccountId}/commodityGroups][%d] deleteCommodityGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteCommodityGroupsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *DeleteCommodityGroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}