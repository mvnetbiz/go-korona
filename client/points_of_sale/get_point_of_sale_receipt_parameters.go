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

// NewGetPointOfSaleReceiptParams creates a new GetPointOfSaleReceiptParams object
// with the default values initialized.
func NewGetPointOfSaleReceiptParams() *GetPointOfSaleReceiptParams {
	var ()
	return &GetPointOfSaleReceiptParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPointOfSaleReceiptParamsWithTimeout creates a new GetPointOfSaleReceiptParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPointOfSaleReceiptParamsWithTimeout(timeout time.Duration) *GetPointOfSaleReceiptParams {
	var ()
	return &GetPointOfSaleReceiptParams{

		timeout: timeout,
	}
}

// NewGetPointOfSaleReceiptParamsWithContext creates a new GetPointOfSaleReceiptParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPointOfSaleReceiptParamsWithContext(ctx context.Context) *GetPointOfSaleReceiptParams {
	var ()
	return &GetPointOfSaleReceiptParams{

		Context: ctx,
	}
}

// NewGetPointOfSaleReceiptParamsWithHTTPClient creates a new GetPointOfSaleReceiptParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPointOfSaleReceiptParamsWithHTTPClient(client *http.Client) *GetPointOfSaleReceiptParams {
	var ()
	return &GetPointOfSaleReceiptParams{
		HTTPClient: client,
	}
}

/*GetPointOfSaleReceiptParams contains all the parameters to send to the API endpoint
for the get point of sale receipt operation typically these are written to a http.Request
*/
type GetPointOfSaleReceiptParams struct {

	/*CouplingID
	  the coupling-id of the device. It can be set to check whether your coupling-id is correct or not (works only, if point of sale is external).

	*/
	CouplingID *string
	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string
	/*PointOfSaleID
	  id of the related object (important: id should match the uuid-format)

	*/
	PointOfSaleID string
	/*ReceiptID
	  the id of the receipt

	*/
	ReceiptID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get point of sale receipt params
func (o *GetPointOfSaleReceiptParams) WithTimeout(timeout time.Duration) *GetPointOfSaleReceiptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get point of sale receipt params
func (o *GetPointOfSaleReceiptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get point of sale receipt params
func (o *GetPointOfSaleReceiptParams) WithContext(ctx context.Context) *GetPointOfSaleReceiptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get point of sale receipt params
func (o *GetPointOfSaleReceiptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get point of sale receipt params
func (o *GetPointOfSaleReceiptParams) WithHTTPClient(client *http.Client) *GetPointOfSaleReceiptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get point of sale receipt params
func (o *GetPointOfSaleReceiptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCouplingID adds the couplingID to the get point of sale receipt params
func (o *GetPointOfSaleReceiptParams) WithCouplingID(couplingID *string) *GetPointOfSaleReceiptParams {
	o.SetCouplingID(couplingID)
	return o
}

// SetCouplingID adds the couplingId to the get point of sale receipt params
func (o *GetPointOfSaleReceiptParams) SetCouplingID(couplingID *string) {
	o.CouplingID = couplingID
}

// WithKoronaAccountID adds the koronaAccountID to the get point of sale receipt params
func (o *GetPointOfSaleReceiptParams) WithKoronaAccountID(koronaAccountID string) *GetPointOfSaleReceiptParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the get point of sale receipt params
func (o *GetPointOfSaleReceiptParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WithPointOfSaleID adds the pointOfSaleID to the get point of sale receipt params
func (o *GetPointOfSaleReceiptParams) WithPointOfSaleID(pointOfSaleID string) *GetPointOfSaleReceiptParams {
	o.SetPointOfSaleID(pointOfSaleID)
	return o
}

// SetPointOfSaleID adds the pointOfSaleId to the get point of sale receipt params
func (o *GetPointOfSaleReceiptParams) SetPointOfSaleID(pointOfSaleID string) {
	o.PointOfSaleID = pointOfSaleID
}

// WithReceiptID adds the receiptID to the get point of sale receipt params
func (o *GetPointOfSaleReceiptParams) WithReceiptID(receiptID string) *GetPointOfSaleReceiptParams {
	o.SetReceiptID(receiptID)
	return o
}

// SetReceiptID adds the receiptId to the get point of sale receipt params
func (o *GetPointOfSaleReceiptParams) SetReceiptID(receiptID string) {
	o.ReceiptID = receiptID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPointOfSaleReceiptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param receiptId
	if err := r.SetPathParam("receiptId", o.ReceiptID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
