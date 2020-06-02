// Code generated by go-swagger; DO NOT EDIT.

package points_of_sale

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

// NewGetPointOfSaleParams creates a new GetPointOfSaleParams object
// with the default values initialized.
func NewGetPointOfSaleParams() *GetPointOfSaleParams {
	var ()
	return &GetPointOfSaleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPointOfSaleParamsWithTimeout creates a new GetPointOfSaleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPointOfSaleParamsWithTimeout(timeout time.Duration) *GetPointOfSaleParams {
	var ()
	return &GetPointOfSaleParams{

		timeout: timeout,
	}
}

// NewGetPointOfSaleParamsWithContext creates a new GetPointOfSaleParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPointOfSaleParamsWithContext(ctx context.Context) *GetPointOfSaleParams {
	var ()
	return &GetPointOfSaleParams{

		Context: ctx,
	}
}

// NewGetPointOfSaleParamsWithHTTPClient creates a new GetPointOfSaleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPointOfSaleParamsWithHTTPClient(client *http.Client) *GetPointOfSaleParams {
	var ()
	return &GetPointOfSaleParams{
		HTTPClient: client,
	}
}

/*GetPointOfSaleParams contains all the parameters to send to the API endpoint
for the get point of sale operation typically these are written to a http.Request
*/
type GetPointOfSaleParams struct {

	/*CouplingID
	  the coupling-id of the device. It can be set to check whether your coupling-id is correct or not (works only, if point of sale is external).

	*/
	CouplingID *string
	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string
	/*PointOfSaleID
	  the number of the point of sale

	*/
	PointOfSaleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get point of sale params
func (o *GetPointOfSaleParams) WithTimeout(timeout time.Duration) *GetPointOfSaleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get point of sale params
func (o *GetPointOfSaleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get point of sale params
func (o *GetPointOfSaleParams) WithContext(ctx context.Context) *GetPointOfSaleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get point of sale params
func (o *GetPointOfSaleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get point of sale params
func (o *GetPointOfSaleParams) WithHTTPClient(client *http.Client) *GetPointOfSaleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get point of sale params
func (o *GetPointOfSaleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCouplingID adds the couplingID to the get point of sale params
func (o *GetPointOfSaleParams) WithCouplingID(couplingID *string) *GetPointOfSaleParams {
	o.SetCouplingID(couplingID)
	return o
}

// SetCouplingID adds the couplingId to the get point of sale params
func (o *GetPointOfSaleParams) SetCouplingID(couplingID *string) {
	o.CouplingID = couplingID
}

// WithKoronaAccountID adds the koronaAccountID to the get point of sale params
func (o *GetPointOfSaleParams) WithKoronaAccountID(koronaAccountID string) *GetPointOfSaleParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the get point of sale params
func (o *GetPointOfSaleParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WithPointOfSaleID adds the pointOfSaleID to the get point of sale params
func (o *GetPointOfSaleParams) WithPointOfSaleID(pointOfSaleID string) *GetPointOfSaleParams {
	o.SetPointOfSaleID(pointOfSaleID)
	return o
}

// SetPointOfSaleID adds the pointOfSaleId to the get point of sale params
func (o *GetPointOfSaleParams) SetPointOfSaleID(pointOfSaleID string) {
	o.PointOfSaleID = pointOfSaleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPointOfSaleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CouplingID != nil {

		// query param couplingId
		var qrCouplingID string
		if o.CouplingID != nil {
			qrCouplingID = *o.CouplingID
		}
		qCouplingID := qrCouplingID
		if qCouplingID != "" {
			if err := r.SetQueryParam("couplingId", qCouplingID); err != nil {
				return err
			}
		}

	}

	// path param koronaAccountId
	if err := r.SetPathParam("koronaAccountId", o.KoronaAccountID); err != nil {
		return err
	}

	// path param pointOfSaleId
	if err := r.SetPathParam("pointOfSaleId", o.PointOfSaleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
