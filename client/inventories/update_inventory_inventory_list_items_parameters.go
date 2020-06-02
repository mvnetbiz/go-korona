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

	"github.com/mvnetbiz/go-korona/models"
)

// NewUpdateInventoryInventoryListItemsParams creates a new UpdateInventoryInventoryListItemsParams object
// with the default values initialized.
func NewUpdateInventoryInventoryListItemsParams() *UpdateInventoryInventoryListItemsParams {
	var ()
	return &UpdateInventoryInventoryListItemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateInventoryInventoryListItemsParamsWithTimeout creates a new UpdateInventoryInventoryListItemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateInventoryInventoryListItemsParamsWithTimeout(timeout time.Duration) *UpdateInventoryInventoryListItemsParams {
	var ()
	return &UpdateInventoryInventoryListItemsParams{

		timeout: timeout,
	}
}

// NewUpdateInventoryInventoryListItemsParamsWithContext creates a new UpdateInventoryInventoryListItemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateInventoryInventoryListItemsParamsWithContext(ctx context.Context) *UpdateInventoryInventoryListItemsParams {
	var ()
	return &UpdateInventoryInventoryListItemsParams{

		Context: ctx,
	}
}

// NewUpdateInventoryInventoryListItemsParamsWithHTTPClient creates a new UpdateInventoryInventoryListItemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateInventoryInventoryListItemsParamsWithHTTPClient(client *http.Client) *UpdateInventoryInventoryListItemsParams {
	var ()
	return &UpdateInventoryInventoryListItemsParams{
		HTTPClient: client,
	}
}

/*UpdateInventoryInventoryListItemsParams contains all the parameters to send to the API endpoint
for the update inventory inventory list items operation typically these are written to a http.Request
*/
type UpdateInventoryInventoryListItemsParams struct {

	/*Body
	  an array of inventory list items to update

	*/
	Body []*models.InventoryListItem
	/*InventoryID
	  id of the related object (important: id should match the uuid-format)

	*/
	InventoryID string
	/*InventoryListID
	  id of the related object (important: id should match the uuid-format)

	*/
	InventoryListID string
	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update inventory inventory list items params
func (o *UpdateInventoryInventoryListItemsParams) WithTimeout(timeout time.Duration) *UpdateInventoryInventoryListItemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update inventory inventory list items params
func (o *UpdateInventoryInventoryListItemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update inventory inventory list items params
func (o *UpdateInventoryInventoryListItemsParams) WithContext(ctx context.Context) *UpdateInventoryInventoryListItemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update inventory inventory list items params
func (o *UpdateInventoryInventoryListItemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update inventory inventory list items params
func (o *UpdateInventoryInventoryListItemsParams) WithHTTPClient(client *http.Client) *UpdateInventoryInventoryListItemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update inventory inventory list items params
func (o *UpdateInventoryInventoryListItemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update inventory inventory list items params
func (o *UpdateInventoryInventoryListItemsParams) WithBody(body []*models.InventoryListItem) *UpdateInventoryInventoryListItemsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update inventory inventory list items params
func (o *UpdateInventoryInventoryListItemsParams) SetBody(body []*models.InventoryListItem) {
	o.Body = body
}

// WithInventoryID adds the inventoryID to the update inventory inventory list items params
func (o *UpdateInventoryInventoryListItemsParams) WithInventoryID(inventoryID string) *UpdateInventoryInventoryListItemsParams {
	o.SetInventoryID(inventoryID)
	return o
}

// SetInventoryID adds the inventoryId to the update inventory inventory list items params
func (o *UpdateInventoryInventoryListItemsParams) SetInventoryID(inventoryID string) {
	o.InventoryID = inventoryID
}

// WithInventoryListID adds the inventoryListID to the update inventory inventory list items params
func (o *UpdateInventoryInventoryListItemsParams) WithInventoryListID(inventoryListID string) *UpdateInventoryInventoryListItemsParams {
	o.SetInventoryListID(inventoryListID)
	return o
}

// SetInventoryListID adds the inventoryListId to the update inventory inventory list items params
func (o *UpdateInventoryInventoryListItemsParams) SetInventoryListID(inventoryListID string) {
	o.InventoryListID = inventoryListID
}

// WithKoronaAccountID adds the koronaAccountID to the update inventory inventory list items params
func (o *UpdateInventoryInventoryListItemsParams) WithKoronaAccountID(koronaAccountID string) *UpdateInventoryInventoryListItemsParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the update inventory inventory list items params
func (o *UpdateInventoryInventoryListItemsParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateInventoryInventoryListItemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param inventoryId
	if err := r.SetPathParam("inventoryId", o.InventoryID); err != nil {
		return err
	}

	// path param inventoryListId
	if err := r.SetPathParam("inventoryListId", o.InventoryListID); err != nil {
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
