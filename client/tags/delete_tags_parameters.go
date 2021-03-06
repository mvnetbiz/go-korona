// Code generated by go-swagger; DO NOT EDIT.

package tags

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

// NewDeleteTagsParams creates a new DeleteTagsParams object
// with the default values initialized.
func NewDeleteTagsParams() *DeleteTagsParams {
	var ()
	return &DeleteTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTagsParamsWithTimeout creates a new DeleteTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTagsParamsWithTimeout(timeout time.Duration) *DeleteTagsParams {
	var ()
	return &DeleteTagsParams{

		timeout: timeout,
	}
}

// NewDeleteTagsParamsWithContext creates a new DeleteTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTagsParamsWithContext(ctx context.Context) *DeleteTagsParams {
	var ()
	return &DeleteTagsParams{

		Context: ctx,
	}
}

// NewDeleteTagsParamsWithHTTPClient creates a new DeleteTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTagsParamsWithHTTPClient(client *http.Client) *DeleteTagsParams {
	var ()
	return &DeleteTagsParams{
		HTTPClient: client,
	}
}

/*DeleteTagsParams contains all the parameters to send to the API endpoint
for the delete tags operation typically these are written to a http.Request
*/
type DeleteTagsParams struct {

	/*Body
	  array of existing tags (id or number required)

	*/
	Body []*models.Tag
	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete tags params
func (o *DeleteTagsParams) WithTimeout(timeout time.Duration) *DeleteTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete tags params
func (o *DeleteTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete tags params
func (o *DeleteTagsParams) WithContext(ctx context.Context) *DeleteTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete tags params
func (o *DeleteTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete tags params
func (o *DeleteTagsParams) WithHTTPClient(client *http.Client) *DeleteTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete tags params
func (o *DeleteTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete tags params
func (o *DeleteTagsParams) WithBody(body []*models.Tag) *DeleteTagsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete tags params
func (o *DeleteTagsParams) SetBody(body []*models.Tag) {
	o.Body = body
}

// WithKoronaAccountID adds the koronaAccountID to the delete tags params
func (o *DeleteTagsParams) WithKoronaAccountID(koronaAccountID string) *DeleteTagsParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the delete tags params
func (o *DeleteTagsParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
