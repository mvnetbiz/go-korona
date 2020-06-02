// Code generated by go-swagger; DO NOT EDIT.

package sectors

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

// NewGetSectorsParams creates a new GetSectorsParams object
// with the default values initialized.
func NewGetSectorsParams() *GetSectorsParams {
	var ()
	return &GetSectorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSectorsParamsWithTimeout creates a new GetSectorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSectorsParamsWithTimeout(timeout time.Duration) *GetSectorsParams {
	var ()
	return &GetSectorsParams{

		timeout: timeout,
	}
}

// NewGetSectorsParamsWithContext creates a new GetSectorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSectorsParamsWithContext(ctx context.Context) *GetSectorsParams {
	var ()
	return &GetSectorsParams{

		Context: ctx,
	}
}

// NewGetSectorsParamsWithHTTPClient creates a new GetSectorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSectorsParamsWithHTTPClient(client *http.Client) *GetSectorsParams {
	var ()
	return &GetSectorsParams{
		HTTPClient: client,
	}
}

/*GetSectorsParams contains all the parameters to send to the API endpoint
for the get sectors operation typically these are written to a http.Request
*/
type GetSectorsParams struct {

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

// WithTimeout adds the timeout to the get sectors params
func (o *GetSectorsParams) WithTimeout(timeout time.Duration) *GetSectorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sectors params
func (o *GetSectorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sectors params
func (o *GetSectorsParams) WithContext(ctx context.Context) *GetSectorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sectors params
func (o *GetSectorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sectors params
func (o *GetSectorsParams) WithHTTPClient(client *http.Client) *GetSectorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sectors params
func (o *GetSectorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeDeleted adds the includeDeleted to the get sectors params
func (o *GetSectorsParams) WithIncludeDeleted(includeDeleted *bool) *GetSectorsParams {
	o.SetIncludeDeleted(includeDeleted)
	return o
}

// SetIncludeDeleted adds the includeDeleted to the get sectors params
func (o *GetSectorsParams) SetIncludeDeleted(includeDeleted *bool) {
	o.IncludeDeleted = includeDeleted
}

// WithKoronaAccountID adds the koronaAccountID to the get sectors params
func (o *GetSectorsParams) WithKoronaAccountID(koronaAccountID string) *GetSectorsParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the get sectors params
func (o *GetSectorsParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WithNumber adds the number to the get sectors params
func (o *GetSectorsParams) WithNumber(number *string) *GetSectorsParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get sectors params
func (o *GetSectorsParams) SetNumber(number *string) {
	o.Number = number
}

// WithPage adds the page to the get sectors params
func (o *GetSectorsParams) WithPage(page *int32) *GetSectorsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get sectors params
func (o *GetSectorsParams) SetPage(page *int32) {
	o.Page = page
}

// WithRevision adds the revision to the get sectors params
func (o *GetSectorsParams) WithRevision(revision *int64) *GetSectorsParams {
	o.SetRevision(revision)
	return o
}

// SetRevision adds the revision to the get sectors params
func (o *GetSectorsParams) SetRevision(revision *int64) {
	o.Revision = revision
}

// WithSize adds the size to the get sectors params
func (o *GetSectorsParams) WithSize(size *int32) *GetSectorsParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get sectors params
func (o *GetSectorsParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the get sectors params
func (o *GetSectorsParams) WithSort(sort *string) *GetSectorsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get sectors params
func (o *GetSectorsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetSectorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
