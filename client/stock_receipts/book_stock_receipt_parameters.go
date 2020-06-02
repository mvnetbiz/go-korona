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

// NewBookStockReceiptParams creates a new BookStockReceiptParams object
// with the default values initialized.
func NewBookStockReceiptParams() *BookStockReceiptParams {
	var ()
	return &BookStockReceiptParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBookStockReceiptParamsWithTimeout creates a new BookStockReceiptParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBookStockReceiptParamsWithTimeout(timeout time.Duration) *BookStockReceiptParams {
	var ()
	return &BookStockReceiptParams{

		timeout: timeout,
	}
}

// NewBookStockReceiptParamsWithContext creates a new BookStockReceiptParams object
// with the default values initialized, and the ability to set a context for a request
func NewBookStockReceiptParamsWithContext(ctx context.Context) *BookStockReceiptParams {
	var ()
	return &BookStockReceiptParams{

		Context: ctx,
	}
}

// NewBookStockReceiptParamsWithHTTPClient creates a new BookStockReceiptParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBookStockReceiptParamsWithHTTPClient(client *http.Client) *BookStockReceiptParams {
	var ()
	return &BookStockReceiptParams{
		HTTPClient: client,
	}
}

/*BookStockReceiptParams contains all the parameters to send to the API endpoint
for the book stock receipt operation typically these are written to a http.Request
*/
type BookStockReceiptParams struct {

	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string
	/*StockReceiptID
	  id of the related object (important: id should match the uuid-format)

	*/
	StockReceiptID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the book stock receipt params
func (o *BookStockReceiptParams) WithTimeout(timeout time.Duration) *BookStockReceiptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the book stock receipt params
func (o *BookStockReceiptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the book stock receipt params
func (o *BookStockReceiptParams) WithContext(ctx context.Context) *BookStockReceiptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the book stock receipt params
func (o *BookStockReceiptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the book stock receipt params
func (o *BookStockReceiptParams) WithHTTPClient(client *http.Client) *BookStockReceiptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the book stock receipt params
func (o *BookStockReceiptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKoronaAccountID adds the koronaAccountID to the book stock receipt params
func (o *BookStockReceiptParams) WithKoronaAccountID(koronaAccountID string) *BookStockReceiptParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the book stock receipt params
func (o *BookStockReceiptParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WithStockReceiptID adds the stockReceiptID to the book stock receipt params
func (o *BookStockReceiptParams) WithStockReceiptID(stockReceiptID string) *BookStockReceiptParams {
	o.SetStockReceiptID(stockReceiptID)
	return o
}

// SetStockReceiptID adds the stockReceiptId to the book stock receipt params
func (o *BookStockReceiptParams) SetStockReceiptID(stockReceiptID string) {
	o.StockReceiptID = stockReceiptID
}

// WriteToRequest writes these params to a swagger request
func (o *BookStockReceiptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param koronaAccountId
	if err := r.SetPathParam("koronaAccountId", o.KoronaAccountID); err != nil {
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
