// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SupplierContact supplier contact
//
// swagger:model SupplierContact
type SupplierContact struct {

	// email
	Email string `json:"email,omitempty"`

	// fax
	Fax string `json:"fax,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`

	// website
	Website string `json:"website,omitempty"`
}

// Validate validates this supplier contact
func (m *SupplierContact) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SupplierContact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupplierContact) UnmarshalBinary(b []byte) error {
	var res SupplierContact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}