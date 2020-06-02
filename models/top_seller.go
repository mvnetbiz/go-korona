// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TopSeller top seller
//
// swagger:model TopSeller
type TopSeller struct {

	// gross amount
	GrossAmount float64 `json:"grossAmount,omitempty"`

	// global object uuid (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx). recommanded to use for linking purposes. will never change.
	ID string `json:"id,omitempty"`

	// name, like it is set in backoffice
	Name string `json:"name,omitempty"`

	// number, like it is set in backoffice
	Number string `json:"number,omitempty"`

	// quantity
	Quantity float64 `json:"quantity,omitempty"`
}

// Validate validates this top seller
func (m *TopSeller) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TopSeller) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopSeller) UnmarshalBinary(b []byte) error {
	var res TopSeller
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
