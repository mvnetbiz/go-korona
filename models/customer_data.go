// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomerData customer data
//
// swagger:model CustomerData
type CustomerData struct {

	// name
	Name string `json:"name,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`
}

// Validate validates this customer data
func (m *CustomerData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomerData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerData) UnmarshalBinary(b []byte) error {
	var res CustomerData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}