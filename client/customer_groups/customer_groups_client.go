// Code generated by go-swagger; DO NOT EDIT.

package customer_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new customer groups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for customer groups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetCustomerGroup(params *GetCustomerGroupParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomerGroupOK, error)

	GetCustomerGroups(params *GetCustomerGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomerGroupsOK, *GetCustomerGroupsNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCustomerGroup returns the single customer group
*/
func (a *Client) GetCustomerGroup(params *GetCustomerGroupParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomerGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomerGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCustomerGroup",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/customerGroups/{customerGroupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCustomerGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCustomerGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCustomerGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCustomerGroups lists all customer groups
*/
func (a *Client) GetCustomerGroups(params *GetCustomerGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomerGroupsOK, *GetCustomerGroupsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomerGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCustomerGroups",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/customerGroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCustomerGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetCustomerGroupsOK:
		return value, nil, nil
	case *GetCustomerGroupsNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for customer_groups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}