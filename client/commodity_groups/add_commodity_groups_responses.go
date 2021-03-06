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

// AddCommodityGroupsReader is a Reader for the AddCommodityGroups structure.
type AddCommodityGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddCommodityGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddCommodityGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddCommodityGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddCommodityGroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddCommodityGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddCommodityGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAddCommodityGroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddCommodityGroupsOK creates a AddCommodityGroupsOK with default headers values
func NewAddCommodityGroupsOK() *AddCommodityGroupsOK {
	return &AddCommodityGroupsOK{}
}

/*AddCommodityGroupsOK handles this case with default header values.

successful operation
*/
type AddCommodityGroupsOK struct {
	Payload []*models.AddOrUpdateResult
}

func (o *AddCommodityGroupsOK) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/commodityGroups][%d] addCommodityGroupsOK  %+v", 200, o.Payload)
}

func (o *AddCommodityGroupsOK) GetPayload() []*models.AddOrUpdateResult {
	return o.Payload
}

func (o *AddCommodityGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddCommodityGroupsBadRequest creates a AddCommodityGroupsBadRequest with default headers values
func NewAddCommodityGroupsBadRequest() *AddCommodityGroupsBadRequest {
	return &AddCommodityGroupsBadRequest{}
}

/*AddCommodityGroupsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type AddCommodityGroupsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *AddCommodityGroupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/commodityGroups][%d] addCommodityGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *AddCommodityGroupsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *AddCommodityGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddCommodityGroupsUnauthorized creates a AddCommodityGroupsUnauthorized with default headers values
func NewAddCommodityGroupsUnauthorized() *AddCommodityGroupsUnauthorized {
	return &AddCommodityGroupsUnauthorized{}
}

/*AddCommodityGroupsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type AddCommodityGroupsUnauthorized struct {
}

func (o *AddCommodityGroupsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/commodityGroups][%d] addCommodityGroupsUnauthorized ", 401)
}

func (o *AddCommodityGroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddCommodityGroupsForbidden creates a AddCommodityGroupsForbidden with default headers values
func NewAddCommodityGroupsForbidden() *AddCommodityGroupsForbidden {
	return &AddCommodityGroupsForbidden{}
}

/*AddCommodityGroupsForbidden handles this case with default header values.

Requested action is not allowed
*/
type AddCommodityGroupsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *AddCommodityGroupsForbidden) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/commodityGroups][%d] addCommodityGroupsForbidden  %+v", 403, o.Payload)
}

func (o *AddCommodityGroupsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *AddCommodityGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddCommodityGroupsNotFound creates a AddCommodityGroupsNotFound with default headers values
func NewAddCommodityGroupsNotFound() *AddCommodityGroupsNotFound {
	return &AddCommodityGroupsNotFound{}
}

/*AddCommodityGroupsNotFound handles this case with default header values.

Object not found
*/
type AddCommodityGroupsNotFound struct {
	Payload *models.NotFoundError
}

func (o *AddCommodityGroupsNotFound) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/commodityGroups][%d] addCommodityGroupsNotFound  %+v", 404, o.Payload)
}

func (o *AddCommodityGroupsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *AddCommodityGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddCommodityGroupsTooManyRequests creates a AddCommodityGroupsTooManyRequests with default headers values
func NewAddCommodityGroupsTooManyRequests() *AddCommodityGroupsTooManyRequests {
	return &AddCommodityGroupsTooManyRequests{}
}

/*AddCommodityGroupsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type AddCommodityGroupsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *AddCommodityGroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/commodityGroups][%d] addCommodityGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AddCommodityGroupsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *AddCommodityGroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
