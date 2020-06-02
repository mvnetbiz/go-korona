// Code generated by go-swagger; DO NOT EDIT.

package organizational_units

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new organizational units API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organizational units API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddOrganizationalUnitDayRatings(params *AddOrganizationalUnitDayRatingsParams, authInfo runtime.ClientAuthInfoWriter) (*AddOrganizationalUnitDayRatingsOK, error)

	DeleteOrganizationalUnitDayRating(params *DeleteOrganizationalUnitDayRatingParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOrganizationalUnitDayRatingNoContent, error)

	GetOrganizationalUnit(params *GetOrganizationalUnitParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationalUnitOK, error)

	GetOrganizationalUnitDayRating(params *GetOrganizationalUnitDayRatingParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationalUnitDayRatingOK, error)

	GetOrganizationalUnitDayRatings(params *GetOrganizationalUnitDayRatingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationalUnitDayRatingsOK, error)

	GetOrganizationalUnitInventoryLists(params *GetOrganizationalUnitInventoryListsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationalUnitInventoryListsOK, *GetOrganizationalUnitInventoryListsNoContent, error)

	GetOrganizationalUnitProductStocks(params *GetOrganizationalUnitProductStocksParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationalUnitProductStocksOK, *GetOrganizationalUnitProductStocksNoContent, error)

	GetOrganizationalUnitStockReceipts(params *GetOrganizationalUnitStockReceiptsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationalUnitStockReceiptsOK, *GetOrganizationalUnitStockReceiptsNoContent, error)

	GetOrganizationalUnits(params *GetOrganizationalUnitsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationalUnitsOK, *GetOrganizationalUnitsNoContent, error)

	UpdateOrganizationalUnitDayRating(params *UpdateOrganizationalUnitDayRatingParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationalUnitDayRatingNoContent, error)

	UpdateOrganizationalUnitDayRatings(params *UpdateOrganizationalUnitDayRatingsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationalUnitDayRatingsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddOrganizationalUnitDayRatings adds a batch of new day ratings
*/
func (a *Client) AddOrganizationalUnitDayRatings(params *AddOrganizationalUnitDayRatingsParams, authInfo runtime.ClientAuthInfoWriter) (*AddOrganizationalUnitDayRatingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddOrganizationalUnitDayRatingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addOrganizationalUnitDayRatings",
		Method:             "POST",
		PathPattern:        "/accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/dayRatings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddOrganizationalUnitDayRatingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddOrganizationalUnitDayRatingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addOrganizationalUnitDayRatings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteOrganizationalUnitDayRating deletes the single day rating by its id or date
*/
func (a *Client) DeleteOrganizationalUnitDayRating(params *DeleteOrganizationalUnitDayRatingParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOrganizationalUnitDayRatingNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationalUnitDayRatingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOrganizationalUnitDayRating",
		Method:             "DELETE",
		PathPattern:        "/accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/dayRatings/{dayRatingIdOrDate}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOrganizationalUnitDayRatingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteOrganizationalUnitDayRatingNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteOrganizationalUnitDayRating: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationalUnit returns the single organizational unit
*/
func (a *Client) GetOrganizationalUnit(params *GetOrganizationalUnitParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationalUnitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationalUnitParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationalUnit",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationalUnitReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationalUnitOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationalUnit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationalUnitDayRating returns the single day rating by its id or date
*/
func (a *Client) GetOrganizationalUnitDayRating(params *GetOrganizationalUnitDayRatingParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationalUnitDayRatingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationalUnitDayRatingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationalUnitDayRating",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/dayRatings/{dayRatingIdOrDate}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationalUnitDayRatingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationalUnitDayRatingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationalUnitDayRating: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationalUnitDayRatings lists all organizational unit related day ratings
*/
func (a *Client) GetOrganizationalUnitDayRatings(params *GetOrganizationalUnitDayRatingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationalUnitDayRatingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationalUnitDayRatingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationalUnitDayRatings",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/dayRatings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationalUnitDayRatingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationalUnitDayRatingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationalUnitDayRatings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationalUnitInventoryLists lists the inventory lists belonging to the organizational unit k o r o n a retail required
*/
func (a *Client) GetOrganizationalUnitInventoryLists(params *GetOrganizationalUnitInventoryListsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationalUnitInventoryListsOK, *GetOrganizationalUnitInventoryListsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationalUnitInventoryListsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationalUnitInventoryLists",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/inventoryLists",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationalUnitInventoryListsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetOrganizationalUnitInventoryListsOK:
		return value, nil, nil
	case *GetOrganizationalUnitInventoryListsNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for organizational_units: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationalUnitProductStocks lists the product stocks of the organizational unit in case it contains a warehouse k o r o n a retail required
*/
func (a *Client) GetOrganizationalUnitProductStocks(params *GetOrganizationalUnitProductStocksParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationalUnitProductStocksOK, *GetOrganizationalUnitProductStocksNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationalUnitProductStocksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationalUnitProductStocks",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/productStocks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationalUnitProductStocksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetOrganizationalUnitProductStocksOK:
		return value, nil, nil
	case *GetOrganizationalUnitProductStocksNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for organizational_units: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationalUnitStockReceipts lists the stock receipts belonging to the organizational unit k o r o n a retail required
*/
func (a *Client) GetOrganizationalUnitStockReceipts(params *GetOrganizationalUnitStockReceiptsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationalUnitStockReceiptsOK, *GetOrganizationalUnitStockReceiptsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationalUnitStockReceiptsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationalUnitStockReceipts",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/stockReceipts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationalUnitStockReceiptsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetOrganizationalUnitStockReceiptsOK:
		return value, nil, nil
	case *GetOrganizationalUnitStockReceiptsNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for organizational_units: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationalUnits lists all organizational units
*/
func (a *Client) GetOrganizationalUnits(params *GetOrganizationalUnitsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationalUnitsOK, *GetOrganizationalUnitsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationalUnitsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationalUnits",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/organizationalUnits",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationalUnitsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetOrganizationalUnitsOK:
		return value, nil, nil
	case *GetOrganizationalUnitsNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for organizational_units: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateOrganizationalUnitDayRating updates the day rating by its id or date
*/
func (a *Client) UpdateOrganizationalUnitDayRating(params *UpdateOrganizationalUnitDayRatingParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationalUnitDayRatingNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrganizationalUnitDayRatingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateOrganizationalUnitDayRating",
		Method:             "PATCH",
		PathPattern:        "/accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/dayRatings/{dayRatingIdOrDate}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateOrganizationalUnitDayRatingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateOrganizationalUnitDayRatingNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateOrganizationalUnitDayRating: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateOrganizationalUnitDayRatings updates a batch of day ratings
*/
func (a *Client) UpdateOrganizationalUnitDayRatings(params *UpdateOrganizationalUnitDayRatingsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationalUnitDayRatingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrganizationalUnitDayRatingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateOrganizationalUnitDayRatings",
		Method:             "PATCH",
		PathPattern:        "/accounts/{koronaAccountId}/organizationalUnits/{organizationalUnitId}/dayRatings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateOrganizationalUnitDayRatingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateOrganizationalUnitDayRatingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateOrganizationalUnitDayRatings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}