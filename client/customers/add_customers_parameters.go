// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// NewAddCustomersParams creates a new AddCustomersParams object
// with the default values initialized.
func NewAddCustomersParams() *AddCustomersParams {
	var ()
	return &AddCustomersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddCustomersParamsWithTimeout creates a new AddCustomersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddCustomersParamsWithTimeout(timeout time.Duration) *AddCustomersParams {
	var ()
	return &AddCustomersParams{

		timeout: timeout,
	}
}

// NewAddCustomersParamsWithContext creates a new AddCustomersParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddCustomersParamsWithContext(ctx context.Context) *AddCustomersParams {
	var ()
	return &AddCustomersParams{

		Context: ctx,
	}
}

// NewAddCustomersParamsWithHTTPClient creates a new AddCustomersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddCustomersParamsWithHTTPClient(client *http.Client) *AddCustomersParams {
	var ()
	return &AddCustomersParams{
		HTTPClient: client,
	}
}

/*AddCustomersParams contains all the parameters to send to the API endpoint
for the add customers operation typically these are written to a http.Request
*/
type AddCustomersParams struct {

	/*Body
	  array of new customers

	*/
	Body []*models.Customer
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

// WithTimeout adds the timeout to the add customers params
func (o *AddCustomersParams) WithTimeout(timeout time.Duration) *AddCustomersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add customers params
func (o *AddCustomersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add customers params
func (o *AddCustomersParams) WithContext(ctx context.Context) *AddCustomersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add customers params
func (o *AddCustomersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add customers params
func (o *AddCustomersParams) WithHTTPClient(client *http.Client) *AddCustomersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add customers params
func (o *AddCustomersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add customers params
func (o *AddCustomersParams) WithBody(body []*models.Customer) *AddCustomersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add customers params
func (o *AddCustomersParams) SetBody(body []*models.Customer) {
	o.Body = body
}

// WithKoronaAccountID adds the koronaAccountID to the add customers params
func (o *AddCustomersParams) WithKoronaAccountID(koronaAccountID string) *AddCustomersParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the add customers params
func (o *AddCustomersParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WithUpsert adds the upsert to the add customers params
func (o *AddCustomersParams) WithUpsert(upsert *bool) *AddCustomersParams {
	o.SetUpsert(upsert)
	return o
}

// SetUpsert adds the upsert to the add customers params
func (o *AddCustomersParams) SetUpsert(upsert *bool) {
	o.Upsert = upsert
}

// WriteToRequest writes these params to a swagger request
func (o *AddCustomersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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