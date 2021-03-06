// Code generated by go-swagger; DO NOT EDIT.

package sectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new sectors API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sectors API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddSectors(params *AddSectorsParams, authInfo runtime.ClientAuthInfoWriter) (*AddSectorsOK, error)

	GetSector(params *GetSectorParams, authInfo runtime.ClientAuthInfoWriter) (*GetSectorOK, error)

	GetSectors(params *GetSectorsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSectorsOK, *GetSectorsNoContent, error)

	UpdateSectors(params *UpdateSectorsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSectorsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddSectors adds a batch of new sectors
*/
func (a *Client) AddSectors(params *AddSectorsParams, authInfo runtime.ClientAuthInfoWriter) (*AddSectorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddSectorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addSectors",
		Method:             "POST",
		PathPattern:        "/accounts/{koronaAccountId}/sectors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddSectorsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddSectorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addSectors: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSector returns the single sector
*/
func (a *Client) GetSector(params *GetSectorParams, authInfo runtime.ClientAuthInfoWriter) (*GetSectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSectorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSector",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/sectors/{sectorId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSectorReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSectors lists all sectors
*/
func (a *Client) GetSectors(params *GetSectorsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSectorsOK, *GetSectorsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSectorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSectors",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/sectors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSectorsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetSectorsOK:
		return value, nil, nil
	case *GetSectorsNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sectors: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateSectors updates a batch of sectors
*/
func (a *Client) UpdateSectors(params *UpdateSectorsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSectorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSectorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSectors",
		Method:             "PATCH",
		PathPattern:        "/accounts/{koronaAccountId}/sectors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSectorsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateSectorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateSectors: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
