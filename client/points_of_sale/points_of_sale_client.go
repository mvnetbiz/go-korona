// Code generated by go-swagger; DO NOT EDIT.

package points_of_sale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new points of sale API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for points of sale API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddPointOfSaleEndOfDayStatements(params *AddPointOfSaleEndOfDayStatementsParams, authInfo runtime.ClientAuthInfoWriter) (*AddPointOfSaleEndOfDayStatementsOK, error)

	AddPointOfSaleReceipts(params *AddPointOfSaleReceiptsParams, authInfo runtime.ClientAuthInfoWriter) (*AddPointOfSaleReceiptsOK, error)

	GetPointOfSale(params *GetPointOfSaleParams, authInfo runtime.ClientAuthInfoWriter) (*GetPointOfSaleOK, error)

	GetPointOfSaleEndOfDayStatements(params *GetPointOfSaleEndOfDayStatementsParams, authInfo runtime.ClientAuthInfoWriter) (*GetPointOfSaleEndOfDayStatementsOK, *GetPointOfSaleEndOfDayStatementsNoContent, error)

	GetPointOfSaleReceipt(params *GetPointOfSaleReceiptParams, authInfo runtime.ClientAuthInfoWriter) (*GetPointOfSaleReceiptOK, error)

	GetPointOfSaleReceipts(params *GetPointOfSaleReceiptsParams, authInfo runtime.ClientAuthInfoWriter) (*GetPointOfSaleReceiptsOK, *GetPointOfSaleReceiptsNoContent, error)

	GetPointsOfSale(params *GetPointsOfSaleParams, authInfo runtime.ClientAuthInfoWriter) (*GetPointsOfSaleOK, *GetPointsOfSaleNoContent, error)

	UpdatePointOfSale(params *UpdatePointOfSaleParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePointOfSaleNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddPointOfSaleEndOfDayStatements adds a batch of point of sale related end of day statements
*/
func (a *Client) AddPointOfSaleEndOfDayStatements(params *AddPointOfSaleEndOfDayStatementsParams, authInfo runtime.ClientAuthInfoWriter) (*AddPointOfSaleEndOfDayStatementsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddPointOfSaleEndOfDayStatementsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addPointOfSaleEndOfDayStatements",
		Method:             "POST",
		PathPattern:        "/accounts/{koronaAccountId}/pointsOfSale/{pointOfSaleId}/endOfDayStatements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddPointOfSaleEndOfDayStatementsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddPointOfSaleEndOfDayStatementsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addPointOfSaleEndOfDayStatements: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AddPointOfSaleReceipts adds a batch of point of sale related receipts
*/
func (a *Client) AddPointOfSaleReceipts(params *AddPointOfSaleReceiptsParams, authInfo runtime.ClientAuthInfoWriter) (*AddPointOfSaleReceiptsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddPointOfSaleReceiptsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addPointOfSaleReceipts",
		Method:             "POST",
		PathPattern:        "/accounts/{koronaAccountId}/pointsOfSale/{pointOfSaleId}/receipts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddPointOfSaleReceiptsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddPointOfSaleReceiptsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addPointOfSaleReceipts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPointOfSale returns the single point of sale
*/
func (a *Client) GetPointOfSale(params *GetPointOfSaleParams, authInfo runtime.ClientAuthInfoWriter) (*GetPointOfSaleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPointOfSaleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPointOfSale",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/pointsOfSale/{pointOfSaleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPointOfSaleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPointOfSaleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPointOfSale: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPointOfSaleEndOfDayStatements lists all point of sale related end of day statements
*/
func (a *Client) GetPointOfSaleEndOfDayStatements(params *GetPointOfSaleEndOfDayStatementsParams, authInfo runtime.ClientAuthInfoWriter) (*GetPointOfSaleEndOfDayStatementsOK, *GetPointOfSaleEndOfDayStatementsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPointOfSaleEndOfDayStatementsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPointOfSaleEndOfDayStatements",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/pointsOfSale/{pointOfSaleId}/endOfDayStatements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPointOfSaleEndOfDayStatementsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetPointOfSaleEndOfDayStatementsOK:
		return value, nil, nil
	case *GetPointOfSaleEndOfDayStatementsNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for points_of_sale: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPointOfSaleReceipt returns the single point of sale related receipt
*/
func (a *Client) GetPointOfSaleReceipt(params *GetPointOfSaleReceiptParams, authInfo runtime.ClientAuthInfoWriter) (*GetPointOfSaleReceiptOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPointOfSaleReceiptParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPointOfSaleReceipt",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/pointsOfSale/{pointOfSaleId}/receipts/{receiptId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPointOfSaleReceiptReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPointOfSaleReceiptOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPointOfSaleReceipt: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPointOfSaleReceipts lists all point of sale related receipts
*/
func (a *Client) GetPointOfSaleReceipts(params *GetPointOfSaleReceiptsParams, authInfo runtime.ClientAuthInfoWriter) (*GetPointOfSaleReceiptsOK, *GetPointOfSaleReceiptsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPointOfSaleReceiptsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPointOfSaleReceipts",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/pointsOfSale/{pointOfSaleId}/receipts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPointOfSaleReceiptsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetPointOfSaleReceiptsOK:
		return value, nil, nil
	case *GetPointOfSaleReceiptsNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for points_of_sale: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPointsOfSale lists all points of sale
*/
func (a *Client) GetPointsOfSale(params *GetPointsOfSaleParams, authInfo runtime.ClientAuthInfoWriter) (*GetPointsOfSaleOK, *GetPointsOfSaleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPointsOfSaleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPointsOfSale",
		Method:             "GET",
		PathPattern:        "/accounts/{koronaAccountId}/pointsOfSale",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPointsOfSaleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetPointsOfSaleOK:
		return value, nil, nil
	case *GetPointsOfSaleNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for points_of_sale: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdatePointOfSale updates a point of sale works only for coupling attribute coupling Id or updating device information attribute device information
*/
func (a *Client) UpdatePointOfSale(params *UpdatePointOfSaleParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePointOfSaleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePointOfSaleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updatePointOfSale",
		Method:             "PATCH",
		PathPattern:        "/accounts/{koronaAccountId}/pointsOfSale/{pointOfSaleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePointOfSaleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdatePointOfSaleNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updatePointOfSale: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
