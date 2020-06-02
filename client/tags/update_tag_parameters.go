// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/mvnetbiz/go-korona/models"
)

// NewUpdateTagParams creates a new UpdateTagParams object
// with the default values initialized.
func NewUpdateTagParams() *UpdateTagParams {
	var ()
	return &UpdateTagParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTagParamsWithTimeout creates a new UpdateTagParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateTagParamsWithTimeout(timeout time.Duration) *UpdateTagParams {
	var ()
	return &UpdateTagParams{

		timeout: timeout,
	}
}

// NewUpdateTagParamsWithContext creates a new UpdateTagParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateTagParamsWithContext(ctx context.Context) *UpdateTagParams {
	var ()
	return &UpdateTagParams{

		Context: ctx,
	}
}

// NewUpdateTagParamsWithHTTPClient creates a new UpdateTagParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateTagParamsWithHTTPClient(client *http.Client) *UpdateTagParams {
	var ()
	return &UpdateTagParams{
		HTTPClient: client,
	}
}

/*UpdateTagParams contains all the parameters to send to the API endpoint
for the update tag operation typically these are written to a http.Request
*/
type UpdateTagParams struct {

	/*Body
	  the properties to update of the tag

	*/
	Body *models.Tag
	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string
	/*TagID
	  id of the related object (important: id should match the uuid-format)

	*/
	TagID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update tag params
func (o *UpdateTagParams) WithTimeout(timeout time.Duration) *UpdateTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update tag params
func (o *UpdateTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update tag params
func (o *UpdateTagParams) WithContext(ctx context.Context) *UpdateTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update tag params
func (o *UpdateTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update tag params
func (o *UpdateTagParams) WithHTTPClient(client *http.Client) *UpdateTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update tag params
func (o *UpdateTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update tag params
func (o *UpdateTagParams) WithBody(body *models.Tag) *UpdateTagParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update tag params
func (o *UpdateTagParams) SetBody(body *models.Tag) {
	o.Body = body
}

// WithKoronaAccountID adds the koronaAccountID to the update tag params
func (o *UpdateTagParams) WithKoronaAccountID(koronaAccountID string) *UpdateTagParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the update tag params
func (o *UpdateTagParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WithTagID adds the tagID to the update tag params
func (o *UpdateTagParams) WithTagID(tagID string) *UpdateTagParams {
	o.SetTagID(tagID)
	return o
}

// SetTagID adds the tagId to the update tag params
func (o *UpdateTagParams) SetTagID(tagID string) {
	o.TagID = tagID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param koronaAccountId
	if err := r.SetPathParam("koronaAccountId", o.KoronaAccountID); err != nil {
		return err
	}

	// path param tagId
	if err := r.SetPathParam("tagId", o.TagID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
