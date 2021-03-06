// Code generated by go-swagger; DO NOT EDIT.

package sectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// AddSectorsReader is a Reader for the AddSectors structure.
type AddSectorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddSectorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddSectorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddSectorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddSectorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddSectorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddSectorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAddSectorsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddSectorsOK creates a AddSectorsOK with default headers values
func NewAddSectorsOK() *AddSectorsOK {
	return &AddSectorsOK{}
}

/*AddSectorsOK handles this case with default header values.

successful operation
*/
type AddSectorsOK struct {
	Payload []*models.AddOrUpdateResult
}

func (o *AddSectorsOK) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/sectors][%d] addSectorsOK  %+v", 200, o.Payload)
}

func (o *AddSectorsOK) GetPayload() []*models.AddOrUpdateResult {
	return o.Payload
}

func (o *AddSectorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSectorsBadRequest creates a AddSectorsBadRequest with default headers values
func NewAddSectorsBadRequest() *AddSectorsBadRequest {
	return &AddSectorsBadRequest{}
}

/*AddSectorsBadRequest handles this case with default header values.

Malformed querystring or model
*/
type AddSectorsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *AddSectorsBadRequest) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/sectors][%d] addSectorsBadRequest  %+v", 400, o.Payload)
}

func (o *AddSectorsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *AddSectorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSectorsUnauthorized creates a AddSectorsUnauthorized with default headers values
func NewAddSectorsUnauthorized() *AddSectorsUnauthorized {
	return &AddSectorsUnauthorized{}
}

/*AddSectorsUnauthorized handles this case with default header values.

Missing or invalid http-authentication
*/
type AddSectorsUnauthorized struct {
}

func (o *AddSectorsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/sectors][%d] addSectorsUnauthorized ", 401)
}

func (o *AddSectorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddSectorsForbidden creates a AddSectorsForbidden with default headers values
func NewAddSectorsForbidden() *AddSectorsForbidden {
	return &AddSectorsForbidden{}
}

/*AddSectorsForbidden handles this case with default header values.

Requested action is not allowed
*/
type AddSectorsForbidden struct {
	Payload *models.ForbiddenError
}

func (o *AddSectorsForbidden) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/sectors][%d] addSectorsForbidden  %+v", 403, o.Payload)
}

func (o *AddSectorsForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *AddSectorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSectorsNotFound creates a AddSectorsNotFound with default headers values
func NewAddSectorsNotFound() *AddSectorsNotFound {
	return &AddSectorsNotFound{}
}

/*AddSectorsNotFound handles this case with default header values.

Object not found
*/
type AddSectorsNotFound struct {
	Payload *models.NotFoundError
}

func (o *AddSectorsNotFound) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/sectors][%d] addSectorsNotFound  %+v", 404, o.Payload)
}

func (o *AddSectorsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *AddSectorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSectorsTooManyRequests creates a AddSectorsTooManyRequests with default headers values
func NewAddSectorsTooManyRequests() *AddSectorsTooManyRequests {
	return &AddSectorsTooManyRequests{}
}

/*AddSectorsTooManyRequests handles this case with default header values.

Too many requests in a specified period
*/
type AddSectorsTooManyRequests struct {
	Payload *models.TooManyRequestsError
}

func (o *AddSectorsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /accounts/{koronaAccountId}/sectors][%d] addSectorsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AddSectorsTooManyRequests) GetPayload() *models.TooManyRequestsError {
	return o.Payload
}

func (o *AddSectorsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TooManyRequestsError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
