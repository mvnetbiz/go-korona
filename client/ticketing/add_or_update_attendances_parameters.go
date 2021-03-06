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

	"github.com/mvnetbiz/go-korona/models"
)

// NewAddOrUpdateAttendancesParams creates a new AddOrUpdateAttendancesParams object
// with the default values initialized.
func NewAddOrUpdateAttendancesParams() *AddOrUpdateAttendancesParams {
	var ()
	return &AddOrUpdateAttendancesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddOrUpdateAttendancesParamsWithTimeout creates a new AddOrUpdateAttendancesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddOrUpdateAttendancesParamsWithTimeout(timeout time.Duration) *AddOrUpdateAttendancesParams {
	var ()
	return &AddOrUpdateAttendancesParams{

		timeout: timeout,
	}
}

// NewAddOrUpdateAttendancesParamsWithContext creates a new AddOrUpdateAttendancesParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddOrUpdateAttendancesParamsWithContext(ctx context.Context) *AddOrUpdateAttendancesParams {
	var ()
	return &AddOrUpdateAttendancesParams{

		Context: ctx,
	}
}

// NewAddOrUpdateAttendancesParamsWithHTTPClient creates a new AddOrUpdateAttendancesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddOrUpdateAttendancesParamsWithHTTPClient(client *http.Client) *AddOrUpdateAttendancesParams {
	var ()
	return &AddOrUpdateAttendancesParams{
		HTTPClient: client,
	}
}

/*AddOrUpdateAttendancesParams contains all the parameters to send to the API endpoint
for the add or update attendances operation typically these are written to a http.Request
*/
type AddOrUpdateAttendancesParams struct {

	/*Body
	  an array of attendances

	*/
	Body []*models.Attendance
	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add or update attendances params
func (o *AddOrUpdateAttendancesParams) WithTimeout(timeout time.Duration) *AddOrUpdateAttendancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add or update attendances params
func (o *AddOrUpdateAttendancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add or update attendances params
func (o *AddOrUpdateAttendancesParams) WithContext(ctx context.Context) *AddOrUpdateAttendancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add or update attendances params
func (o *AddOrUpdateAttendancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add or update attendances params
func (o *AddOrUpdateAttendancesParams) WithHTTPClient(client *http.Client) *AddOrUpdateAttendancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add or update attendances params
func (o *AddOrUpdateAttendancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add or update attendances params
func (o *AddOrUpdateAttendancesParams) WithBody(body []*models.Attendance) *AddOrUpdateAttendancesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add or update attendances params
func (o *AddOrUpdateAttendancesParams) SetBody(body []*models.Attendance) {
	o.Body = body
}

// WithKoronaAccountID adds the koronaAccountID to the add or update attendances params
func (o *AddOrUpdateAttendancesParams) WithKoronaAccountID(koronaAccountID string) *AddOrUpdateAttendancesParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the add or update attendances params
func (o *AddOrUpdateAttendancesParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *AddOrUpdateAttendancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
