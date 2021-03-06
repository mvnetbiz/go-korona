// Code generated by go-swagger; DO NOT EDIT.

package sales_taxes

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

// NewGetSalesTaxesParams creates a new GetSalesTaxesParams object
// with the default values initialized.
func NewGetSalesTaxesParams() *GetSalesTaxesParams {
	var ()
	return &GetSalesTaxesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSalesTaxesParamsWithTimeout creates a new GetSalesTaxesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSalesTaxesParamsWithTimeout(timeout time.Duration) *GetSalesTaxesParams {
	var ()
	return &GetSalesTaxesParams{

		timeout: timeout,
	}
}

// NewGetSalesTaxesParamsWithContext creates a new GetSalesTaxesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSalesTaxesParamsWithContext(ctx context.Context) *GetSalesTaxesParams {
	var ()
	return &GetSalesTaxesParams{

		Context: ctx,
	}
}

// NewGetSalesTaxesParamsWithHTTPClient creates a new GetSalesTaxesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSalesTaxesParamsWithHTTPClient(client *http.Client) *GetSalesTaxesParams {
	var ()
	return &GetSalesTaxesParams{
		HTTPClient: client,
	}
}

/*GetSalesTaxesParams contains all the parameters to send to the API endpoint
for the get sales taxes operation typically these are written to a http.Request
*/
type GetSalesTaxesParams struct {

	/*IncludeDeleted
	  indicates deleted objects should be loaded or not (default: false)

	*/
	IncludeDeleted *bool
	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string
	/*Number
	  number of the related object

	*/
	Number *string
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

// WithTimeout adds the timeout to the get sales taxes params
func (o *GetSalesTaxesParams) WithTimeout(timeout time.Duration) *GetSalesTaxesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sales taxes params
func (o *GetSalesTaxesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sales taxes params
func (o *GetSalesTaxesParams) WithContext(ctx context.Context) *GetSalesTaxesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sales taxes params
func (o *GetSalesTaxesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sales taxes params
func (o *GetSalesTaxesParams) WithHTTPClient(client *http.Client) *GetSalesTaxesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sales taxes params
func (o *GetSalesTaxesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeDeleted adds the includeDeleted to the get sales taxes params
func (o *GetSalesTaxesParams) WithIncludeDeleted(includeDeleted *bool) *GetSalesTaxesParams {
	o.SetIncludeDeleted(includeDeleted)
	return o
}

// SetIncludeDeleted adds the includeDeleted to the get sales taxes params
func (o *GetSalesTaxesParams) SetIncludeDeleted(includeDeleted *bool) {
	o.IncludeDeleted = includeDeleted
}

// WithKoronaAccountID adds the koronaAccountID to the get sales taxes params
func (o *GetSalesTaxesParams) WithKoronaAccountID(koronaAccountID string) *GetSalesTaxesParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the get sales taxes params
func (o *GetSalesTaxesParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WithNumber adds the number to the get sales taxes params
func (o *GetSalesTaxesParams) WithNumber(number *string) *GetSalesTaxesParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get sales taxes params
func (o *GetSalesTaxesParams) SetNumber(number *string) {
	o.Number = number
}

// WithPage adds the page to the get sales taxes params
func (o *GetSalesTaxesParams) WithPage(page *int32) *GetSalesTaxesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get sales taxes params
func (o *GetSalesTaxesParams) SetPage(page *int32) {
	o.Page = page
}

// WithRevision adds the revision to the get sales taxes params
func (o *GetSalesTaxesParams) WithRevision(revision *int64) *GetSalesTaxesParams {
	o.SetRevision(revision)
	return o
}

// SetRevision adds the revision to the get sales taxes params
func (o *GetSalesTaxesParams) SetRevision(revision *int64) {
	o.Revision = revision
}

// WithSize adds the size to the get sales taxes params
func (o *GetSalesTaxesParams) WithSize(size *int32) *GetSalesTaxesParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get sales taxes params
func (o *GetSalesTaxesParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the get sales taxes params
func (o *GetSalesTaxesParams) WithSort(sort *string) *GetSalesTaxesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get sales taxes params
func (o *GetSalesTaxesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetSalesTaxesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IncludeDeleted != nil {

		// query param includeDeleted
		var qrIncludeDeleted bool
		if o.IncludeDeleted != nil {
			qrIncludeDeleted = *o.IncludeDeleted
		}
		qIncludeDeleted := swag.FormatBool(qrIncludeDeleted)
		if qIncludeDeleted != "" {
			if err := r.SetQueryParam("includeDeleted", qIncludeDeleted); err != nil {
				return err
			}
		}

	}

	// path param koronaAccountId
	if err := r.SetPathParam("koronaAccountId", o.KoronaAccountID); err != nil {
		return err
	}

	if o.Number != nil {

		// query param number
		var qrNumber string
		if o.Number != nil {
			qrNumber = *o.Number
		}
		qNumber := qrNumber
		if qNumber != "" {
			if err := r.SetQueryParam("number", qNumber); err != nil {
				return err
			}
		}

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
