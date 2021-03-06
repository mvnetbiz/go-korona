// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PriceGroup price group
//
// swagger:model PriceGroup
type PriceGroup struct {

	// indicates whether the object is active for use or not
	// Read Only: true
	Active *bool `json:"active,omitempty"`

	// currency
	Currency *ModelReference `json:"currency,omitempty"`

	// global object uuid (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// net prices
	NetPrices bool `json:"netPrices,omitempty"`

	// number of the object, like it is set in backoffice; will be removed when active=false
	Number string `json:"number,omitempty"`

	// the revision number of the object. revision numbers are unique per object-type. there is is no object of the same type with identical revision numbers.
	// Read Only: true
	Revision int64 `json:"revision,omitempty"`
}

// Validate validates this price group
func (m *PriceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PriceGroup) validateCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	if m.Currency != nil {
		if err := m.Currency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currency")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PriceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PriceGroup) UnmarshalBinary(b []byte) error {
	var res PriceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
