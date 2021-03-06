// Code generated by go-swagger; DO NOT EDIT.

package info_texts

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

// NewGetInfoTextParams creates a new GetInfoTextParams object
// with the default values initialized.
func NewGetInfoTextParams() *GetInfoTextParams {
	var ()
	return &GetInfoTextParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInfoTextParamsWithTimeout creates a new GetInfoTextParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInfoTextParamsWithTimeout(timeout time.Duration) *GetInfoTextParams {
	var ()
	return &GetInfoTextParams{

		timeout: timeout,
	}
}

// NewGetInfoTextParamsWithContext creates a new GetInfoTextParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInfoTextParamsWithContext(ctx context.Context) *GetInfoTextParams {
	var ()
	return &GetInfoTextParams{

		Context: ctx,
	}
}

// NewGetInfoTextParamsWithHTTPClient creates a new GetInfoTextParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInfoTextParamsWithHTTPClient(client *http.Client) *GetInfoTextParams {
	var ()
	return &GetInfoTextParams{
		HTTPClient: client,
	}
}

/*GetInfoTextParams contains all the parameters to send to the API endpoint
for the get info text operation typically these are written to a http.Request
*/
type GetInfoTextParams struct {

	/*InfoTextID
	  id of the related object (important: id should match the uuid-format)

	*/
	InfoTextID string
	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get info text params
func (o *GetInfoTextParams) WithTimeout(timeout time.Duration) *GetInfoTextParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get info text params
func (o *GetInfoTextParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get info text params
func (o *GetInfoTextParams) WithContext(ctx context.Context) *GetInfoTextParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get info text params
func (o *GetInfoTextParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get info text params
func (o *GetInfoTextParams) WithHTTPClient(client *http.Client) *GetInfoTextParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get info text params
func (o *GetInfoTextParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfoTextID adds the infoTextID to the get info text params
func (o *GetInfoTextParams) WithInfoTextID(infoTextID string) *GetInfoTextParams {
	o.SetInfoTextID(infoTextID)
	return o
}

// SetInfoTextID adds the infoTextId to the get info text params
func (o *GetInfoTextParams) SetInfoTextID(infoTextID string) {
	o.InfoTextID = infoTextID
}

// WithKoronaAccountID adds the koronaAccountID to the get info text params
func (o *GetInfoTextParams) WithKoronaAccountID(koronaAccountID string) *GetInfoTextParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the get info text params
func (o *GetInfoTextParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInfoTextParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param infoTextId
	if err := r.SetPathParam("infoTextId", o.InfoTextID); err != nil {
		return err
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
