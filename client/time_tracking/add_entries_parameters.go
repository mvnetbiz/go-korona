// Code generated by go-swagger; DO NOT EDIT.

package time_tracking

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

// NewAddEntriesParams creates a new AddEntriesParams object
// with the default values initialized.
func NewAddEntriesParams() *AddEntriesParams {
	var ()
	return &AddEntriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddEntriesParamsWithTimeout creates a new AddEntriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddEntriesParamsWithTimeout(timeout time.Duration) *AddEntriesParams {
	var ()
	return &AddEntriesParams{

		timeout: timeout,
	}
}

// NewAddEntriesParamsWithContext creates a new AddEntriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddEntriesParamsWithContext(ctx context.Context) *AddEntriesParams {
	var ()
	return &AddEntriesParams{

		Context: ctx,
	}
}

// NewAddEntriesParamsWithHTTPClient creates a new AddEntriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddEntriesParamsWithHTTPClient(client *http.Client) *AddEntriesParams {
	var ()
	return &AddEntriesParams{
		HTTPClient: client,
	}
}

/*AddEntriesParams contains all the parameters to send to the API endpoint
for the add entries operation typically these are written to a http.Request
*/
type AddEntriesParams struct {

	/*Body
	  array of new time tracking entries

	*/
	Body []*models.TimeTrackingEntry
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

// WithTimeout adds the timeout to the add entries params
func (o *AddEntriesParams) WithTimeout(timeout time.Duration) *AddEntriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add entries params
func (o *AddEntriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add entries params
func (o *AddEntriesParams) WithContext(ctx context.Context) *AddEntriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add entries params
func (o *AddEntriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add entries params
func (o *AddEntriesParams) WithHTTPClient(client *http.Client) *AddEntriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add entries params
func (o *AddEntriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add entries params
func (o *AddEntriesParams) WithBody(body []*models.TimeTrackingEntry) *AddEntriesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add entries params
func (o *AddEntriesParams) SetBody(body []*models.TimeTrackingEntry) {
	o.Body = body
}

// WithKoronaAccountID adds the koronaAccountID to the add entries params
func (o *AddEntriesParams) WithKoronaAccountID(koronaAccountID string) *AddEntriesParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the add entries params
func (o *AddEntriesParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WithUpsert adds the upsert to the add entries params
func (o *AddEntriesParams) WithUpsert(upsert *bool) *AddEntriesParams {
	o.SetUpsert(upsert)
	return o
}

// SetUpsert adds the upsert to the add entries params
func (o *AddEntriesParams) SetUpsert(upsert *bool) {
	o.Upsert = upsert
}

// WriteToRequest writes these params to a swagger request
func (o *AddEntriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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