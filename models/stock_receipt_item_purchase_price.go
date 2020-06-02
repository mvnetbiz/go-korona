// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StockReceiptItemPurchasePrice stock receipt item purchase price
//
// swagger:model StockReceiptItemPurchasePrice
type StockReceiptItemPurchasePrice struct {

	// actual
	Actual float64 `json:"actual,omitempty"`

	// old
	Old float64 `json:"old,omitempty"`
}

// Validate validates this stock receipt item purchase price
func (m *StockReceiptItemPurchasePrice) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StockReceiptItemPurchasePrice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StockReceiptItemPurchasePrice) UnmarshalBinary(b []byte) error {
	var res StockReceiptItemPurchasePrice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}