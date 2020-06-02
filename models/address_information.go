// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddressInformation address information
//
// swagger:model AddressInformation
type AddressInformation struct {

	// address line1
	AddressLine1 string `json:"addressLine1,omitempty"`

	// address line2
	AddressLine2 string `json:"addressLine2,omitempty"`

	// city
	City string `json:"city,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// department
	Department string `json:"department,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// zip code
	ZipCode string `json:"zipCode,omitempty"`
}

// Validate validates this address information
func (m *AddressInformation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddressInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddressInformation) UnmarshalBinary(b []byte) error {
	var res AddressInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}