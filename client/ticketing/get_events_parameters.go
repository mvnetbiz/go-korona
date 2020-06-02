// Code generated by go-swagger; DO NOT EDIT.

package ticketing

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

// NewGetEventsParams creates a new GetEventsParams object
// with the default values initialized.
func NewGetEventsParams() *GetEventsParams {
	var ()
	return &GetEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventsParamsWithTimeout creates a new GetEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEventsParamsWithTimeout(timeout time.Duration) *GetEventsParams {
	var ()
	return &GetEventsParams{

		timeout: timeout,
	}
}

// NewGetEventsParamsWithContext creates a new GetEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEventsParamsWithContext(ctx context.Context) *GetEventsParams {
	var ()
	return &GetEventsParams{

		Context: ctx,
	}
}

// NewGetEventsParamsWithHTTPClient creates a new GetEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEventsParamsWithHTTPClient(client *http.Client) *GetEventsParams {
	var ()
	return &GetEventsParams{
		HTTPClient: client,
	}
}

/*GetEventsParams contains all the parameters to send to the API endpoint
for the get events operation typically these are written to a http.Request
*/
type GetEventsParams struct {

	/*IncludeDeleted
	  indicates deleted objects should be loaded or not (default: false)

	*/
	IncludeDeleted *bool
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

// WithTimeout adds the timeout to the get events params
func (o *GetEventsParams) WithTimeout(timeout time.Duration) *GetEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get events params
func (o *GetEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get events params
func (o *GetEventsParams) WithContext(ctx context.Context) *GetEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get events params
func (o *GetEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get events params
func (o *GetEventsParams) WithHTTPClient(client *http.Client) *GetEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get events params
func (o *GetEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeDeleted adds the includeDeleted to the get events params
func (o *GetEventsParams) WithIncludeDeleted(includeDeleted *bool) *GetEventsParams {
	o.SetIncludeDeleted(includeDeleted)
	return o
}

// SetIncludeDeleted adds the includeDeleted to the get events params
func (o *GetEventsParams) SetIncludeDeleted(includeDeleted *bool) {
	o.IncludeDeleted = includeDeleted
}

// WithKoronaAccountID adds the koronaAccountID to the get events params
func (o *GetEventsParams) WithKoronaAccountID(koronaAccountID string) *GetEventsParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the get events params
func (o *GetEventsParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WithPage adds the page to the get events params
func (o *GetEventsParams) WithPage(page *int32) *GetEventsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get events params
func (o *GetEventsParams) SetPage(page *int32) {
	o.Page = page
}

// WithRevision adds the revision to the get events params
func (o *GetEventsParams) WithRevision(revision *int64) *GetEventsParams {
	o.SetRevision(revision)
	return o
}

// SetRevision adds the revision to the get events params
func (o *GetEventsParams) SetRevision(revision *int64) {
	o.Revision = revision
}

// WithSize adds the size to the get events params
func (o *GetEventsParams) WithSize(size *int32) *GetEventsParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get events params
func (o *GetEventsParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the get events params
func (o *GetEventsParams) WithSort(sort *string) *GetEventsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get events params
func (o *GetEventsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
