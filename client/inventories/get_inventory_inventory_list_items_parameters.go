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
	"github.com/go-openapi/swag"
)

// NewGetInventoryInventoryListItemsParams creates a new GetInventoryInventoryListItemsParams object
// with the default values initialized.
func NewGetInventoryInventoryListItemsParams() *GetInventoryInventoryListItemsParams {
	var ()
	return &GetInventoryInventoryListItemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInventoryInventoryListItemsParamsWithTimeout creates a new GetInventoryInventoryListItemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInventoryInventoryListItemsParamsWithTimeout(timeout time.Duration) *GetInventoryInventoryListItemsParams {
	var ()
	return &GetInventoryInventoryListItemsParams{

		timeout: timeout,
	}
}

// NewGetInventoryInventoryListItemsParamsWithContext creates a new GetInventoryInventoryListItemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInventoryInventoryListItemsParamsWithContext(ctx context.Context) *GetInventoryInventoryListItemsParams {
	var ()
	return &GetInventoryInventoryListItemsParams{

		Context: ctx,
	}
}

// NewGetInventoryInventoryListItemsParamsWithHTTPClient creates a new GetInventoryInventoryListItemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInventoryInventoryListItemsParamsWithHTTPClient(client *http.Client) *GetInventoryInventoryListItemsParams {
	var ()
	return &GetInventoryInventoryListItemsParams{
		HTTPClient: client,
	}
}

/*GetInventoryInventoryListItemsParams contains all the parameters to send to the API endpoint
for the get inventory inventory list items operation typically these are written to a http.Request
*/
type GetInventoryInventoryListItemsParams struct {

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
	/*Page
	  number of the page to fetch

	*/
	Page *int32
	/*Revision
	  last revision number, objects with a greater revision than this will be returned

	*/
	Revision *int64
	/*Size
	  amount of objects to return per page

	*/
	Size *int32
	/*Sort
	  attribute to sort by (multiple separated by comma; max. 5)

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) WithTimeout(timeout time.Duration) *GetInventoryInventoryListItemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) WithContext(ctx context.Context) *GetInventoryInventoryListItemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) WithHTTPClient(client *http.Client) *GetInventoryInventoryListItemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInventoryID adds the inventoryID to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) WithInventoryID(inventoryID string) *GetInventoryInventoryListItemsParams {
	o.SetInventoryID(inventoryID)
	return o
}

// SetInventoryID adds the inventoryId to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) SetInventoryID(inventoryID string) {
	o.InventoryID = inventoryID
}

// WithInventoryListID adds the inventoryListID to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) WithInventoryListID(inventoryListID string) *GetInventoryInventoryListItemsParams {
	o.SetInventoryListID(inventoryListID)
	return o
}

// SetInventoryListID adds the inventoryListId to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) SetInventoryListID(inventoryListID string) {
	o.InventoryListID = inventoryListID
}

// WithKoronaAccountID adds the koronaAccountID to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) WithKoronaAccountID(koronaAccountID string) *GetInventoryInventoryListItemsParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WithPage adds the page to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) WithPage(page *int32) *GetInventoryInventoryListItemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) SetPage(page *int32) {
	o.Page = page
}

// WithRevision adds the revision to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) WithRevision(revision *int64) *GetInventoryInventoryListItemsParams {
	o.SetRevision(revision)
	return o
}

// SetRevision adds the revision to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) SetRevision(revision *int64) {
	o.Revision = revision
}

// WithSize adds the size to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) WithSize(size *int32) *GetInventoryInventoryListItemsParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) WithSort(sort *string) *GetInventoryInventoryListItemsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get inventory inventory list items params
func (o *GetInventoryInventoryListItemsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetInventoryInventoryListItemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.Revision != nil {

		// query param revision
		var qrRevision int64
		if o.Revision != nil {
			qrRevision = *o.Revision
		}
		qRevision := swag.FormatInt64(qrRevision)
		if qRevision != "" {
			if err := r.SetQueryParam("revision", qRevision); err != nil {
				return err
			}
		}

	}

	if o.Size != nil {

		// query param size
		var qrSize int32
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
