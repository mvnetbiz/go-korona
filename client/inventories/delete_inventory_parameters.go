// Code generated by go-swagger; DO NOT EDIT.

package inventories

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

// NewDeleteInventoryParams creates a new DeleteInventoryParams object
// with the default values initialized.
func NewDeleteInventoryParams() *DeleteInventoryParams {
	var ()
	return &DeleteInventoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInventoryParamsWithTimeout creates a new DeleteInventoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInventoryParamsWithTimeout(timeout time.Duration) *DeleteInventoryParams {
	var ()
	return &DeleteInventoryParams{

		timeout: timeout,
	}
}

// NewDeleteInventoryParamsWithContext creates a new DeleteInventoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInventoryParamsWithContext(ctx context.Context) *DeleteInventoryParams {
	var ()
	return &DeleteInventoryParams{

		Context: ctx,
	}
}

// NewDeleteInventoryParamsWithHTTPClient creates a new DeleteInventoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInventoryParamsWithHTTPClient(client *http.Client) *DeleteInventoryParams {
	var ()
	return &DeleteInventoryParams{
		HTTPClient: client,
	}
}

/*DeleteInventoryParams contains all the parameters to send to the API endpoint
for the delete inventory operation typically these are written to a http.Request
*/
type DeleteInventoryParams struct {

	/*InventoryID
	  id of the related object (important: id should match the uuid-format)

	*/
	InventoryID string
	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete inventory params
func (o *DeleteInventoryParams) WithTimeout(timeout time.Duration) *DeleteInventoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete inventory params
func (o *DeleteInventoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete inventory params
func (o *DeleteInventoryParams) WithContext(ctx context.Context) *DeleteInventoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete inventory params
func (o *DeleteInventoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete inventory params
func (o *DeleteInventoryParams) WithHTTPClient(client *http.Client) *DeleteInventoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete inventory params
func (o *DeleteInventoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInventoryID adds the inventoryID to the delete inventory params
func (o *DeleteInventoryParams) WithInventoryID(inventoryID string) *DeleteInventoryParams {
	o.SetInventoryID(inventoryID)
	return o
}

// SetInventoryID adds the inventoryId to the delete inventory params
func (o *DeleteInventoryParams) SetInventoryID(inventoryID string) {
	o.InventoryID = inventoryID
}

// WithKoronaAccountID adds the koronaAccountID to the delete inventory params
func (o *DeleteInventoryParams) WithKoronaAccountID(koronaAccountID string) *DeleteInventoryParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the delete inventory params
func (o *DeleteInventoryParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInventoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param inventoryId
	if err := r.SetPathParam("inventoryId", o.InventoryID); err != nil {
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