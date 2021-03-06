// Code generated by go-swagger; DO NOT EDIT.

package customer_orders

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

// NewUpdateCustomerOrdersParams creates a new UpdateCustomerOrdersParams object
// with the default values initialized.
func NewUpdateCustomerOrdersParams() *UpdateCustomerOrdersParams {
	var ()
	return &UpdateCustomerOrdersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCustomerOrdersParamsWithTimeout creates a new UpdateCustomerOrdersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCustomerOrdersParamsWithTimeout(timeout time.Duration) *UpdateCustomerOrdersParams {
	var ()
	return &UpdateCustomerOrdersParams{

		timeout: timeout,
	}
}

// NewUpdateCustomerOrdersParamsWithContext creates a new UpdateCustomerOrdersParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCustomerOrdersParamsWithContext(ctx context.Context) *UpdateCustomerOrdersParams {
	var ()
	return &UpdateCustomerOrdersParams{

		Context: ctx,
	}
}

// NewUpdateCustomerOrdersParamsWithHTTPClient creates a new UpdateCustomerOrdersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCustomerOrdersParamsWithHTTPClient(client *http.Client) *UpdateCustomerOrdersParams {
	var ()
	return &UpdateCustomerOrdersParams{
		HTTPClient: client,
	}
}

/*UpdateCustomerOrdersParams contains all the parameters to send to the API endpoint
for the update customer orders operation typically these are written to a http.Request
*/
type UpdateCustomerOrdersParams struct {

	/*Body
	  an array of existing customer orders

	*/
	Body []*models.CustomerOrder
	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update customer orders params
func (o *UpdateCustomerOrdersParams) WithTimeout(timeout time.Duration) *UpdateCustomerOrdersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update customer orders params
func (o *UpdateCustomerOrdersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update customer orders params
func (o *UpdateCustomerOrdersParams) WithContext(ctx context.Context) *UpdateCustomerOrdersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update customer orders params
func (o *UpdateCustomerOrdersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update customer orders params
func (o *UpdateCustomerOrdersParams) WithHTTPClient(client *http.Client) *UpdateCustomerOrdersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update customer orders params
func (o *UpdateCustomerOrdersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update customer orders params
func (o *UpdateCustomerOrdersParams) WithBody(body []*models.CustomerOrder) *UpdateCustomerOrdersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update customer orders params
func (o *UpdateCustomerOrdersParams) SetBody(body []*models.CustomerOrder) {
	o.Body = body
}

// WithKoronaAccountID adds the koronaAccountID to the update customer orders params
func (o *UpdateCustomerOrdersParams) WithKoronaAccountID(koronaAccountID string) *UpdateCustomerOrdersParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the update customer orders params
func (o *UpdateCustomerOrdersParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCustomerOrdersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
