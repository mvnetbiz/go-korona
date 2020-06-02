// Code generated by go-swagger; DO NOT EDIT.

package delivery_notes

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

// NewBookDeliveryNoteParams creates a new BookDeliveryNoteParams object
// with the default values initialized.
func NewBookDeliveryNoteParams() *BookDeliveryNoteParams {
	var ()
	return &BookDeliveryNoteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBookDeliveryNoteParamsWithTimeout creates a new BookDeliveryNoteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBookDeliveryNoteParamsWithTimeout(timeout time.Duration) *BookDeliveryNoteParams {
	var ()
	return &BookDeliveryNoteParams{

		timeout: timeout,
	}
}

// NewBookDeliveryNoteParamsWithContext creates a new BookDeliveryNoteParams object
// with the default values initialized, and the ability to set a context for a request
func NewBookDeliveryNoteParamsWithContext(ctx context.Context) *BookDeliveryNoteParams {
	var ()
	return &BookDeliveryNoteParams{

		Context: ctx,
	}
}

// NewBookDeliveryNoteParamsWithHTTPClient creates a new BookDeliveryNoteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBookDeliveryNoteParamsWithHTTPClient(client *http.Client) *BookDeliveryNoteParams {
	var ()
	return &BookDeliveryNoteParams{
		HTTPClient: client,
	}
}

/*BookDeliveryNoteParams contains all the parameters to send to the API endpoint
for the book delivery note operation typically these are written to a http.Request
*/
type BookDeliveryNoteParams struct {

	/*DeliveryNoteID
	  id of the related object (important: id should match the uuid-format)

	*/
	DeliveryNoteID string
	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the book delivery note params
func (o *BookDeliveryNoteParams) WithTimeout(timeout time.Duration) *BookDeliveryNoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the book delivery note params
func (o *BookDeliveryNoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the book delivery note params
func (o *BookDeliveryNoteParams) WithContext(ctx context.Context) *BookDeliveryNoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the book delivery note params
func (o *BookDeliveryNoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the book delivery note params
func (o *BookDeliveryNoteParams) WithHTTPClient(client *http.Client) *BookDeliveryNoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the book delivery note params
func (o *BookDeliveryNoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeliveryNoteID adds the deliveryNoteID to the book delivery note params
func (o *BookDeliveryNoteParams) WithDeliveryNoteID(deliveryNoteID string) *BookDeliveryNoteParams {
	o.SetDeliveryNoteID(deliveryNoteID)
	return o
}

// SetDeliveryNoteID adds the deliveryNoteId to the book delivery note params
func (o *BookDeliveryNoteParams) SetDeliveryNoteID(deliveryNoteID string) {
	o.DeliveryNoteID = deliveryNoteID
}

// WithKoronaAccountID adds the koronaAccountID to the book delivery note params
func (o *BookDeliveryNoteParams) WithKoronaAccountID(koronaAccountID string) *BookDeliveryNoteParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the book delivery note params
func (o *BookDeliveryNoteParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *BookDeliveryNoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deliveryNoteId
	if err := r.SetPathParam("deliveryNoteId", o.DeliveryNoteID); err != nil {
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