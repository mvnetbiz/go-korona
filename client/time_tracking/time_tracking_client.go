// Code generated by go-swagger; DO NOT EDIT.

package time_tracking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new time tracking API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for time tracking API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddEntities(params *AddEntitiesParams, authInfo runtime.ClientAuthInfoWriter) (*AddEntitiesOK, error)

	AddEntries(params *AddEntriesParams, authInfo runtime.ClientAuthInfoWriter) (*AddEntriesOK, error)

	GetEntities(params *GetEntitiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetEntitiesOK, *GetEntitiesNoContent, error)

	GetEntries(params *GetEntriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetEntriesOK, *GetEntriesNoContent, error)

	GetTimeTrackingEntity(params *GetTimeTrackingEntityParams, authInfo runtime.ClientAuthInfoWriter) (*GetTimeTrackingEntityOK, error)

	GetTimeTrackingEntry(params *GetTimeTrackingEntryParams, authInfo runtime.ClientAuthInfoWriter) (*GetTimeTrackingEntryOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddEntities adds a batch of new time tracking entities
*/
func (a *Client) AddEntities(params *AddEntitiesParams, authInfo runtime.ClientAuthInfoWriter) (*AddEntitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddEntitiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addEntities",
		Method:             "POST",
		PathPattern:        "/accounts/{koronaAccountId}/timeTrackingEntities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddEntitiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddEntitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addEntities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AddEntries adds a batch of new time tracking entries
*/
func (a *Client) AddEntries(params *AddEntriesParams, authInfo runtime.ClientAuthInfoWriter) (*AddEntriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddEntriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addEntries",
		Method:             "POST",
		PathPattern:        "/accounts/{koronaAccountId}/timeTrackingEntries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddEntriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddEntriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addEntries: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEntities lists time tracking entities
*/
func (a *Client) GetEntities(params *GetEntitiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetEntitiesOK, *GetEntitiesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEntitiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEntities",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/timeTrackingEntities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEntitiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetEntitiesOK:
		return value, nil, nil
	case *GetEntitiesNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for time_tracking: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEntries lists time tracking entries
*/
func (a *Client) GetEntries(params *GetEntriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetEntriesOK, *GetEntriesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEntriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEntries",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/timeTrackingEntries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEntriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetEntriesOK:
		return value, nil, nil
	case *GetEntriesNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for time_tracking: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTimeTrackingEntity returns a single time tracking entity
*/
func (a *Client) GetTimeTrackingEntity(params *GetTimeTrackingEntityParams, authInfo runtime.ClientAuthInfoWriter) (*GetTimeTrackingEntityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTimeTrackingEntityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTimeTrackingEntity",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/timeTrackingEntities/{timeTrackingEntityId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTimeTrackingEntityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTimeTrackingEntityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTimeTrackingEntity: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTimeTrackingEntry returns a single time tracking entry
*/
func (a *Client) GetTimeTrackingEntry(params *GetTimeTrackingEntryParams, authInfo runtime.ClientAuthInfoWriter) (*GetTimeTrackingEntryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTimeTrackingEntryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTimeTrackingEntry",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/timeTrackingEntries/{timeTrackingEntryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTimeTrackingEntryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTimeTrackingEntryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTimeTrackingEntry: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}