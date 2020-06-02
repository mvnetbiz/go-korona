// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AmountCountData amount count data
//
// swagger:model AmountCountData
type AmountCountData struct {

	// amount
	Amount float64 `json:"amount,omitempty"`

	// count
	Count float64 `json:"count,omitempty"`
}

// Validate validates this amount count data
func (m *AmountCountData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AmountCountData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AmountCountData) UnmarshalBinary(b []byte) error {
	var res AmountCountData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
