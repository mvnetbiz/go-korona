// Code generated by go-swagger; DO NOT EDIT.

package cashiers

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

// NewGetCashiersParams creates a new GetCashiersParams object
// with the default values initialized.
func NewGetCashiersParams() *GetCashiersParams {
	var ()
	return &GetCashiersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCashiersParamsWithTimeout creates a new GetCashiersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCashiersParamsWithTimeout(timeout time.Duration) *GetCashiersParams {
	var ()
	return &GetCashiersParams{

		timeout: timeout,
	}
}

// NewGetCashiersParamsWithContext creates a new GetCashiersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCashiersParamsWithContext(ctx context.Context) *GetCashiersParams {
	var ()
	return &GetCashiersParams{

		Context: ctx,
	}
}

// NewGetCashiersParamsWithHTTPClient creates a new GetCashiersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCashiersParamsWithHTTPClient(client *http.Client) *GetCashiersParams {
	var ()
	return &GetCashiersParams{
		HTTPClient: client,
	}
}

/*GetCashiersParams contains all the parameters to send to the API endpoint
for the get cashiers operation typically these are written to a http.Request
*/
type GetCashiersParams struct {

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

// WithTimeout adds the timeout to the get cashiers params
func (o *GetCashiersParams) WithTimeout(timeout time.Duration) *GetCashiersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cashiers params
func (o *GetCashiersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cashiers params
func (o *GetCashiersParams) WithContext(ctx context.Context) *GetCashiersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cashiers params
func (o *GetCashiersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cashiers params
func (o *GetCashiersParams) WithHTTPClient(client *http.Client) *GetCashiersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cashiers params
func (o *GetCashiersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeDeleted adds the includeDeleted to the get cashiers params
func (o *GetCashiersParams) WithIncludeDeleted(includeDeleted *bool) *GetCashiersParams {
	o.SetIncludeDeleted(includeDeleted)
	return o
}

// SetIncludeDeleted adds the includeDeleted to the get cashiers params
func (o *GetCashiersParams) SetIncludeDeleted(includeDeleted *bool) {
	o.IncludeDeleted = includeDeleted
}

// WithKoronaAccountID adds the koronaAccountID to the get cashiers params
func (o *GetCashiersParams) WithKoronaAccountID(koronaAccountID string) *GetCashiersParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the get cashiers params
func (o *GetCashiersParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WithNumber adds the number to the get cashiers params
func (o *GetCashiersParams) WithNumber(number *string) *GetCashiersParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get cashiers params
func (o *GetCashiersParams) SetNumber(number *string) {
	o.Number = number
}

// WithPage adds the page to the get cashiers params
func (o *GetCashiersParams) WithPage(page *int32) *GetCashiersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get cashiers params
func (o *GetCashiersParams) SetPage(page *int32) {
	o.Page = page
}

// WithRevision adds the revision to the get cashiers params
func (o *GetCashiersParams) WithRevision(revision *int64) *GetCashiersParams {
	o.SetRevision(revision)
	return o
}

// SetRevision adds the revision to the get cashiers params
func (o *GetCashiersParams) SetRevision(revision *int64) {
	o.Revision = revision
}

// WithSize adds the size to the get cashiers params
func (o *GetCashiersParams) WithSize(size *int32) *GetCashiersParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get cashiers params
func (o *GetCashiersParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the get cashiers params
func (o *GetCashiersParams) WithSort(sort *string) *GetCashiersParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get cashiers params
func (o *GetCashiersParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetCashiersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
