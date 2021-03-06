// Code generated by go-swagger; DO NOT EDIT.

package ticketing

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
	"github.com/go-openapi/swag"

	"github.com/mvnetbiz/go-korona/models"
)

// NewAddEventsParams creates a new AddEventsParams object
// with the default values initialized.
func NewAddEventsParams() *AddEventsParams {
	var ()
	return &AddEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddEventsParamsWithTimeout creates a new AddEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddEventsParamsWithTimeout(timeout time.Duration) *AddEventsParams {
	var ()
	return &AddEventsParams{

		timeout: timeout,
	}
}

// NewAddEventsParamsWithContext creates a new AddEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddEventsParamsWithContext(ctx context.Context) *AddEventsParams {
	var ()
	return &AddEventsParams{

		Context: ctx,
	}
}

// NewAddEventsParamsWithHTTPClient creates a new AddEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddEventsParamsWithHTTPClient(client *http.Client) *AddEventsParams {
	var ()
	return &AddEventsParams{
		HTTPClient: client,
	}
}

/*AddEventsParams contains all the parameters to send to the API endpoint
for the add events operation typically these are written to a http.Request
*/
type AddEventsParams struct {

	/*Body
	  an array of new events

	*/
	Body []*models.Event
	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string
	/*Upsert
	  when set to true, updates the object instead of generating a already-exists-error

	*/
	Upsert *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add events params
func (o *AddEventsParams) WithTimeout(timeout time.Duration) *AddEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add events params
func (o *AddEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add events params
func (o *AddEventsParams) WithContext(ctx context.Context) *AddEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add events params
func (o *AddEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add events params
func (o *AddEventsParams) WithHTTPClient(client *http.Client) *AddEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add events params
func (o *AddEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add events params
func (o *AddEventsParams) WithBody(body []*models.Event) *AddEventsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add events params
func (o *AddEventsParams) SetBody(body []*models.Event) {
	o.Body = body
}

// WithKoronaAccountID adds the koronaAccountID to the add events params
func (o *AddEventsParams) WithKoronaAccountID(koronaAccountID string) *AddEventsParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the add events params
func (o *AddEventsParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WithUpsert adds the upsert to the add events params
func (o *AddEventsParams) WithUpsert(upsert *bool) *AddEventsParams {
	o.SetUpsert(upsert)
	return o
}

// SetUpsert adds the upsert to the add events params
func (o *AddEventsParams) SetUpsert(upsert *bool) {
	o.Upsert = upsert
}

// WriteToRequest writes these params to a swagger request
func (o *AddEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Upsert != nil {

		// query param upsert
		var qrUpsert bool
		if o.Upsert != nil {
			qrUpsert = *o.Upsert
		}
		qUpsert := swag.FormatBool(qrUpsert)
		if qUpsert != "" {
			if err := r.SetQueryParam("upsert", qUpsert); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
