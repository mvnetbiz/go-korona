// Code generated by go-swagger; DO NOT EDIT.

package korona_account

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
)

// NewGetKoronaAccountParams creates a new GetKoronaAccountParams object
// with the default values initialized.
func NewGetKoronaAccountParams() *GetKoronaAccountParams {
	var ()
	return &GetKoronaAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKoronaAccountParamsWithTimeout creates a new GetKoronaAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKoronaAccountParamsWithTimeout(timeout time.Duration) *GetKoronaAccountParams {
	var ()
	return &GetKoronaAccountParams{

		timeout: timeout,
	}
}

// NewGetKoronaAccountParamsWithContext creates a new GetKoronaAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKoronaAccountParamsWithContext(ctx context.Context) *GetKoronaAccountParams {
	var ()
	return &GetKoronaAccountParams{

		Context: ctx,
	}
}

// NewGetKoronaAccountParamsWithHTTPClient creates a new GetKoronaAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKoronaAccountParamsWithHTTPClient(client *http.Client) *GetKoronaAccountParams {
	var ()
	return &GetKoronaAccountParams{
		HTTPClient: client,
	}
}

/*GetKoronaAccountParams contains all the parameters to send to the API endpoint
for the get korona account operation typically these are written to a http.Request
*/
type GetKoronaAccountParams struct {

	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get korona account params
func (o *GetKoronaAccountParams) WithTimeout(timeout time.Duration) *GetKoronaAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get korona account params
func (o *GetKoronaAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get korona account params
func (o *GetKoronaAccountParams) WithContext(ctx context.Context) *GetKoronaAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get korona account params
func (o *GetKoronaAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get korona account params
func (o *GetKoronaAccountParams) WithHTTPClient(client *http.Client) *GetKoronaAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get korona account params
func (o *GetKoronaAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKoronaAccountID adds the koronaAccountID to the get korona account params
func (o *GetKoronaAccountParams) WithKoronaAccountID(koronaAccountID string) *GetKoronaAccountParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the get korona account params
func (o *GetKoronaAccountParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetKoronaAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param koronaAccountId
	if err := r.SetPathParam("koronaAccountId", o.KoronaAccountID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}