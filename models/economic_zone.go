// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EconomicZone economic zone
//
// swagger:model EconomicZone
type EconomicZone struct {

	// indicates whether the object is active for use or not
	// Read Only: true
	Active *bool `json:"active,omitempty"`

	// global object uuid (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// number of the object, like it is set in backoffice; will be removed when active=false
	Number string `json:"number,omitempty"`

	// the revision number of the object. revision numbers are unique per object-type. there is is no object of the same type with identical revision numbers.
	// Read Only: true
	Revision int64 `json:"revision,omitempty"`
}

// Validate validates this economic zone
func (m *EconomicZone) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EconomicZone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EconomicZone) UnmarshalBinary(b []byte) error {
	var res EconomicZone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
