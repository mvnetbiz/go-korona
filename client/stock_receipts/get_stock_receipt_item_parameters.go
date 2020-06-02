// Code generated by go-swagger; DO NOT EDIT.

package stock_receipts

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

// NewGetStockReceiptItemParams creates a new GetStockReceiptItemParams object
// with the default values initialized.
func NewGetStockReceiptItemParams() *GetStockReceiptItemParams {
	var ()
	return &GetStockReceiptItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStockReceiptItemParamsWithTimeout creates a new GetStockReceiptItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStockReceiptItemParamsWithTimeout(timeout time.Duration) *GetStockReceiptItemParams {
	var ()
	return &GetStockReceiptItemParams{

		timeout: timeout,
	}
}

// NewGetStockReceiptItemParamsWithContext creates a new GetStockReceiptItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStockReceiptItemParamsWithContext(ctx context.Context) *GetStockReceiptItemParams {
	var ()
	return &GetStockReceiptItemParams{

		Context: ctx,
	}
}

// NewGetStockReceiptItemParamsWithHTTPClient creates a new GetStockReceiptItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStockReceiptItemParamsWithHTTPClient(client *http.Client) *GetStockReceiptItemParams {
	var ()
	return &GetStockReceiptItemParams{
		HTTPClient: client,
	}
}

/*GetStockReceiptItemParams contains all the parameters to send to the API endpoint
for the get stock receipt item operation typically these are written to a http.Request
*/
type GetStockReceiptItemParams struct {

	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string
	/*ProductID
	  id of the related object (important: id should match the uuid-format)

	*/
	ProductID string
	/*StockReceiptID
	  id of the related object (important: id should match the uuid-format)

	*/
	StockReceiptID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get stock receipt item params
func (o *GetStockReceiptItemParams) WithTimeout(timeout time.Duration) *GetStockReceiptItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get stock receipt item params
func (o *GetStockReceiptItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get stock receipt item params
func (o *GetStockReceiptItemParams) WithContext(ctx context.Context) *GetStockReceiptItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get stock receipt item params
func (o *GetStockReceiptItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get stock receipt item params
func (o *GetStockReceiptItemParams) WithHTTPClient(client *http.Client) *GetStockReceiptItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get stock receipt item params
func (o *GetStockReceiptItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKoronaAccountID adds the koronaAccountID to the get stock receipt item params
func (o *GetStockReceiptItemParams) WithKoronaAccountID(koronaAccountID string) *GetStockReceiptItemParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the get stock receipt item params
func (o *GetStockReceiptItemParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WithProductID adds the productID to the get stock receipt item params
func (o *GetStockReceiptItemParams) WithProductID(productID string) *GetStockReceiptItemParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the get stock receipt item params
func (o *GetStockReceiptItemParams) SetProductID(productID string) {
	o.ProductID = productID
}

// WithStockReceiptID adds the stockReceiptID to the get stock receipt item params
func (o *GetStockReceiptItemParams) WithStockReceiptID(stockReceiptID string) *GetStockReceiptItemParams {
	o.SetStockReceiptID(stockReceiptID)
	return o
}

// SetStockReceiptID adds the stockReceiptId to the get stock receipt item params
func (o *GetStockReceiptItemParams) SetStockReceiptID(stockReceiptID string) {
	o.StockReceiptID = stockReceiptID
}

// WriteToRequest writes these params to a swagger request
func (o *GetStockReceiptItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param koronaAccountId
	if err := r.SetPathParam("koronaAccountId", o.KoronaAccountID); err != nil {
		return err
	}

	// path param productId
	if err := r.SetPathParam("productId", o.ProductID); err != nil {
		return err
	}

	// path param stockReceiptId
	if err := r.SetPathParam("stockReceiptId", o.StockReceiptID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
