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

	"github.com/mvnetbiz/go-korona/models"
)

// NewAddSectorsParams creates a new AddSectorsParams object
// with the default values initialized.
func NewAddSectorsParams() *AddSectorsParams {
	var ()
	return &AddSectorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddSectorsParamsWithTimeout creates a new AddSectorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddSectorsParamsWithTimeout(timeout time.Duration) *AddSectorsParams {
	var ()
	return &AddSectorsParams{

		timeout: timeout,
	}
}

// NewAddSectorsParamsWithContext creates a new AddSectorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddSectorsParamsWithContext(ctx context.Context) *AddSectorsParams {
	var ()
	return &AddSectorsParams{

		Context: ctx,
	}
}

// NewAddSectorsParamsWithHTTPClient creates a new AddSectorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddSectorsParamsWithHTTPClient(client *http.Client) *AddSectorsParams {
	var ()
	return &AddSectorsParams{
		HTTPClient: client,
	}
}

/*AddSectorsParams contains all the parameters to send to the API endpoint
for the add sectors operation typically these are written to a http.Request
*/
type AddSectorsParams struct {

	/*Body
	  array of new sectors

	*/
	Body []*models.Sector
	/*KoronaAccountID
	  account id of the KORONA.cloud account

	*/
	KoronaAccountID string
	/*Upsert
	  when set to true, updates the object instead of generating a already-exists-error

	*/
	Upsert *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add sectors params
func (o *AddSectorsParams) WithTimeout(timeout time.Duration) *AddSectorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add sectors params
func (o *AddSectorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add sectors params
func (o *AddSectorsParams) WithContext(ctx context.Context) *AddSectorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add sectors params
func (o *AddSectorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add sectors params
func (o *AddSectorsParams) WithHTTPClient(client *http.Client) *AddSectorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add sectors params
func (o *AddSectorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add sectors params
func (o *AddSectorsParams) WithBody(body []*models.Sector) *AddSectorsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add sectors params
func (o *AddSectorsParams) SetBody(body []*models.Sector) {
	o.Body = body
}

// WithKoronaAccountID adds the koronaAccountID to the add sectors params
func (o *AddSectorsParams) WithKoronaAccountID(koronaAccountID string) *AddSectorsParams {
	o.SetKoronaAccountID(koronaAccountID)
	return o
}

// SetKoronaAccountID adds the koronaAccountId to the add sectors params
func (o *AddSectorsParams) SetKoronaAccountID(koronaAccountID string) {
	o.KoronaAccountID = koronaAccountID
}

// WithUpsert adds the upsert to the add sectors params
func (o *AddSectorsParams) WithUpsert(upsert *bool) *AddSectorsParams {
	o.SetUpsert(upsert)
	return o
}

// SetUpsert adds the upsert to the add sectors params
func (o *AddSectorsParams) SetUpsert(upsert *bool) {
	o.Upsert = upsert
}

// WriteToRequest writes these params to a swagger request
func (o *AddSectorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Upsert != nil {

		// query param upsert
		var qrUpsert bool
		if o.Upsert != nil {
			qrUpsert = *o.Upsert
		}
		qUpsert := swag.FormatBool(qrUpsert)
		if qUpsert != "" {
			if err := r.SetQueryParam("upsert", qUpsert); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
