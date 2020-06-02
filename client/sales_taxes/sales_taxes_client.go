// Code generated by go-swagger; DO NOT EDIT.

package sales_taxes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new sales taxes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sales taxes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetSalesTax(params *GetSalesTaxParams, authInfo runtime.ClientAuthInfoWriter) (*GetSalesTaxOK, error)

	GetSalesTaxes(params *GetSalesTaxesParams, authInfo runtime.ClientAuthInfoWriter) (*GetSalesTaxesOK, *GetSalesTaxesNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetSalesTax returns the single sales tax
*/
func (a *Client) GetSalesTax(params *GetSalesTaxParams, authInfo runtime.ClientAuthInfoWriter) (*GetSalesTaxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSalesTaxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSalesTax",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/salesTaxes/{salesTaxId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSalesTaxReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSalesTaxOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSalesTax: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSalesTaxes lists all sales taxes
*/
func (a *Client) GetSalesTaxes(params *GetSalesTaxesParams, authInfo runtime.ClientAuthInfoWriter) (*GetSalesTaxesOK, *GetSalesTaxesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSalesTaxesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSalesTaxes",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/salesTaxes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSalesTaxesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetSalesTaxesOK:
		return value, nil, nil
	case *GetSalesTaxesNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sales_taxes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
