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

	"github.com/mvnetbiz/go-korona/models"
)

// NewDeleteAssortmentsParams creates a new DeleteAssortmentsParams object
// with the default values initialized.
func NewDeleteAssortmentsParams() *DeleteAssortmentsParams {
	var ()
	return &DeleteAssortmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAssortmentsParamsWithTimeout creates a new DeleteAssortmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAssortmentsParamsWithTimeout(timeout time.Duration) *DeleteAssortmentsParams {
	var ()
	return &DeleteAssortmentsParams{

		timeout: timeout,
	}
}

// NewDeleteAssortmentsParamsWithContext creates a new DeleteAssortmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAssortmentsParamsWithContext(ctx context.Context) *DeleteAssortmentsParams {
	var ()
	return &DeleteAssortmentsParams{

		Context: ctx,
	}
}

// NewDeleteAssortmentsParamsWithHTTPClient creates a new DeleteAssortmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAssortmentsParamsWithHTTPClient(client *http.Client) *DeleteAssortmentsParams {
	var ()
	return &DeleteAssortmentsParams{
		HTTPClient: client,
	}
}

/*DeleteAssortmentsParams contains all the parameters to send to the API endpoint
for the delete assortments operation typically these are written to a http.Request
*/
type DeleteAssortmentsParams struct {

	/*Body
	  array of existing assortments (id or number required)

	*/
	Body []*models.Assortment
	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete assortments params
func (o *DeleteAssortmentsParams) WithTimeout(timeout time.Duration) *DeleteAssortmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete assortments params
func (o *DeleteAssortmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete assortments params
func (o *DeleteAssortmentsParams) WithContext(ctx context.Context) *DeleteAssortmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete assortments params
func (o *DeleteAssortmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete assortments params
func (o *DeleteAssortmentsParams) WithHTTPClient(client *http.Client) *DeleteAssortmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete assortments params
func (o *DeleteAssortmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete assortments params
func (o *DeleteAssortmentsParams) WithBody(body []*models.Assortment) *DeleteAssortmentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete assortments params
func (o *DeleteAssortmentsParams) SetBody(body []*models.Assortment) {
	o.Body = body
}

// WithKoronaAccountID adds the koronaAccountID to the delete assortments params
func (o *DeleteAssortmentsParams) WithKoronaAccountID(koronaAccountID string) *DeleteAssortmentsParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the delete assortments params
func (o *DeleteAssortmentsParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAssortmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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