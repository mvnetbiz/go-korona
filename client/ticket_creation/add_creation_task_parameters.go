// Code generated by go-swagger; DO NOT EDIT.

package ticket_creation

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

// NewAddCreationTaskParams creates a new AddCreationTaskParams object
// with the default values initialized.
func NewAddCreationTaskParams() *AddCreationTaskParams {
	var ()
	return &AddCreationTaskParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddCreationTaskParamsWithTimeout creates a new AddCreationTaskParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddCreationTaskParamsWithTimeout(timeout time.Duration) *AddCreationTaskParams {
	var ()
	return &AddCreationTaskParams{

		timeout: timeout,
	}
}

// NewAddCreationTaskParamsWithContext creates a new AddCreationTaskParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddCreationTaskParamsWithContext(ctx context.Context) *AddCreationTaskParams {
	var ()
	return &AddCreationTaskParams{

		Context: ctx,
	}
}

// NewAddCreationTaskParamsWithHTTPClient creates a new AddCreationTaskParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddCreationTaskParamsWithHTTPClient(client *http.Client) *AddCreationTaskParams {
	var ()
	return &AddCreationTaskParams{
		HTTPClient: client,
	}
}

/*AddCreationTaskParams contains all the parameters to send to the API endpoint
for the add creation task operation typically these are written to a http.Request
*/
type AddCreationTaskParams struct {

	/*Body
	  a single ticket creation task

	*/
	Body *models.CreationTask
	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add creation task params
func (o *AddCreationTaskParams) WithTimeout(timeout time.Duration) *AddCreationTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add creation task params
func (o *AddCreationTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add creation task params
func (o *AddCreationTaskParams) WithContext(ctx context.Context) *AddCreationTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add creation task params
func (o *AddCreationTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add creation task params
func (o *AddCreationTaskParams) WithHTTPClient(client *http.Client) *AddCreationTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add creation task params
func (o *AddCreationTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add creation task params
func (o *AddCreationTaskParams) WithBody(body *models.CreationTask) *AddCreationTaskParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add creation task params
func (o *AddCreationTaskParams) SetBody(body *models.CreationTask) {
	o.Body = body
}

// WithKoronaAccountID adds the koronaAccountID to the add creation task params
func (o *AddCreationTaskParams) WithKoronaAccountID(koronaAccountID string) *AddCreationTaskParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the add creation task params
func (o *AddCreationTaskParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *AddCreationTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
