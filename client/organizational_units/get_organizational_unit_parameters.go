// Code generated by go-swagger; DO NOT EDIT.

package organizational_units

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

// NewGetOrganizationalUnitParams creates a new GetOrganizationalUnitParams object
// with the default values initialized.
func NewGetOrganizationalUnitParams() *GetOrganizationalUnitParams {
	var ()
	return &GetOrganizationalUnitParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationalUnitParamsWithTimeout creates a new GetOrganizationalUnitParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrganizationalUnitParamsWithTimeout(timeout time.Duration) *GetOrganizationalUnitParams {
	var ()
	return &GetOrganizationalUnitParams{

		timeout: timeout,
	}
}

// NewGetOrganizationalUnitParamsWithContext creates a new GetOrganizationalUnitParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrganizationalUnitParamsWithContext(ctx context.Context) *GetOrganizationalUnitParams {
	var ()
	return &GetOrganizationalUnitParams{

		Context: ctx,
	}
}

// NewGetOrganizationalUnitParamsWithHTTPClient creates a new GetOrganizationalUnitParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrganizationalUnitParamsWithHTTPClient(client *http.Client) *GetOrganizationalUnitParams {
	var ()
	return &GetOrganizationalUnitParams{
		HTTPClient: client,
	}
}

/*GetOrganizationalUnitParams contains all the parameters to send to the API endpoint
for the get organizational unit operation typically these are written to a http.Request
*/
type GetOrganizationalUnitParams struct {

	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string
	/*OrganizationalUnitID
	  id of the related object (important: id should match the uuid-format)

	*/
	OrganizationalUnitID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get organizational unit params
func (o *GetOrganizationalUnitParams) WithTimeout(timeout time.Duration) *GetOrganizationalUnitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organizational unit params
func (o *GetOrganizationalUnitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organizational unit params
func (o *GetOrganizationalUnitParams) WithContext(ctx context.Context) *GetOrganizationalUnitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organizational unit params
func (o *GetOrganizationalUnitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organizational unit params
func (o *GetOrganizationalUnitParams) WithHTTPClient(client *http.Client) *GetOrganizationalUnitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organizational unit params
func (o *GetOrganizationalUnitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKoronaAccountID adds the koronaAccountID to the get organizational unit params
func (o *GetOrganizationalUnitParams) WithKoronaAccountID(koronaAccountID string) *GetOrganizationalUnitParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the get organizational unit params
func (o *GetOrganizationalUnitParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WithOrganizationalUnitID adds the organizationalUnitID to the get organizational unit params
func (o *GetOrganizationalUnitParams) WithOrganizationalUnitID(organizationalUnitID string) *GetOrganizationalUnitParams {
	o.SetOrganizationalUnitID(organizationalUnitID)
	return o
}

// SetOrganizationalUnitID adds the organizationalUnitId to the get organizational unit params
func (o *GetOrganizationalUnitParams) SetOrganizationalUnitID(organizationalUnitID string) {
	o.OrganizationalUnitID = organizationalUnitID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationalUnitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param koronaAccountId
	if err := r.SetPathParam("koronaAccountId", o.KoronaAccountID); err != nil {
		return err
	}

	// path param organizationalUnitId
	if err := r.SetPathParam("organizationalUnitId", o.OrganizationalUnitID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}