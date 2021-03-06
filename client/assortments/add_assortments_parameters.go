// Code generated by go-swagger; DO NOT EDIT.

package assortments

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

// NewAddAssortmentsParams creates a new AddAssortmentsParams object
// with the default values initialized.
func NewAddAssortmentsParams() *AddAssortmentsParams {
	var ()
	return &AddAssortmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddAssortmentsParamsWithTimeout creates a new AddAssortmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddAssortmentsParamsWithTimeout(timeout time.Duration) *AddAssortmentsParams {
	var ()
	return &AddAssortmentsParams{

		timeout: timeout,
	}
}

// NewAddAssortmentsParamsWithContext creates a new AddAssortmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddAssortmentsParamsWithContext(ctx context.Context) *AddAssortmentsParams {
	var ()
	return &AddAssortmentsParams{

		Context: ctx,
	}
}

// NewAddAssortmentsParamsWithHTTPClient creates a new AddAssortmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddAssortmentsParamsWithHTTPClient(client *http.Client) *AddAssortmentsParams {
	var ()
	return &AddAssortmentsParams{
		HTTPClient: client,
	}
}

/*AddAssortmentsParams contains all the parameters to send to the API endpoint
for the add assortments operation typically these are written to a http.Request
*/
type AddAssortmentsParams struct {

	/*Body
	  array of new assortments

	*/
	Body []*models.Assortment
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

// WithTimeout adds the timeout to the add assortments params
func (o *AddAssortmentsParams) WithTimeout(timeout time.Duration) *AddAssortmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add assortments params
func (o *AddAssortmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add assortments params
func (o *AddAssortmentsParams) WithContext(ctx context.Context) *AddAssortmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add assortments params
func (o *AddAssortmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add assortments params
func (o *AddAssortmentsParams) WithHTTPClient(client *http.Client) *AddAssortmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add assortments params
func (o *AddAssortmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add assortments params
func (o *AddAssortmentsParams) WithBody(body []*models.Assortment) *AddAssortmentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add assortments params
func (o *AddAssortmentsParams) SetBody(body []*models.Assortment) {
	o.Body = body
}

// WithKoronaAccountID adds the koronaAccountID to the add assortments params
func (o *AddAssortmentsParams) WithKoronaAccountID(koronaAccountID string) *AddAssortmentsParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the add assortments params
func (o *AddAssortmentsParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WithUpsert adds the upsert to the add assortments params
func (o *AddAssortmentsParams) WithUpsert(upsert *bool) *AddAssortmentsParams {
	o.SetUpsert(upsert)
	return o
}

// SetUpsert adds the upsert to the add assortments params
func (o *AddAssortmentsParams) SetUpsert(upsert *bool) {
	o.Upsert = upsert
}

// WriteToRequest writes these params to a swagger request
func (o *AddAssortmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
