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
)

// NewGetCustomerParams creates a new GetCustomerParams object
// with the default values initialized.
func NewGetCustomerParams() *GetCustomerParams {
	var ()
	return &GetCustomerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomerParamsWithTimeout creates a new GetCustomerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomerParamsWithTimeout(timeout time.Duration) *GetCustomerParams {
	var ()
	return &GetCustomerParams{

		timeout: timeout,
	}
}

// NewGetCustomerParamsWithContext creates a new GetCustomerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomerParamsWithContext(ctx context.Context) *GetCustomerParams {
	var ()
	return &GetCustomerParams{

		Context: ctx,
	}
}

// NewGetCustomerParamsWithHTTPClient creates a new GetCustomerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomerParamsWithHTTPClient(client *http.Client) *GetCustomerParams {
	var ()
	return &GetCustomerParams{
		HTTPClient: client,
	}
}

/*GetCustomerParams contains all the parameters to send to the API endpoint
for the get customer operation typically these are written to a http.Request
*/
type GetCustomerParams struct {

	/*CustomerID
	  id of the related object (important: id should match the uuid-format)

	*/
	CustomerID string
	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get customer params
func (o *GetCustomerParams) WithTimeout(timeout time.Duration) *GetCustomerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customer params
func (o *GetCustomerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customer params
func (o *GetCustomerParams) WithContext(ctx context.Context) *GetCustomerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customer params
func (o *GetCustomerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customer params
func (o *GetCustomerParams) WithHTTPClient(client *http.Client) *GetCustomerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customer params
func (o *GetCustomerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get customer params
func (o *GetCustomerParams) WithCustomerID(customerID string) *GetCustomerParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get customer params
func (o *GetCustomerParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithKoronaAccountID adds the koronaAccountID to the get customer params
func (o *GetCustomerParams) WithKoronaAccountID(koronaAccountID string) *GetCustomerParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the get customer params
func (o *GetCustomerParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", o.CustomerID); err != nil {
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
