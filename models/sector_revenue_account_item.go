// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SectorRevenueAccountItem sector revenue account item
//
// swagger:model SectorRevenueAccountItem
type SectorRevenueAccountItem struct {

	// economic zone
	EconomicZone *ModelReference `json:"economicZone,omitempty"`

	// revenue account
	RevenueAccount *ModelReference `json:"revenueAccount,omitempty"`
}

// Validate validates this sector revenue account item
func (m *SectorRevenueAccountItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEconomicZone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevenueAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SectorRevenueAccountItem) validateEconomicZone(formats strfmt.Registry) error {

	if swag.IsZero(m.EconomicZone) { // not required
		return nil
	}

	if m.EconomicZone != nil {
		if err := m.EconomicZone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("economicZone")
			}
			return err
		}
	}

	return nil
}

func (m *SectorRevenueAccountItem) validateRevenueAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.RevenueAccount) { // not required
		return nil
	}

	if m.RevenueAccount != nil {
		if err := m.RevenueAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revenueAccount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SectorRevenueAccountItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SectorRevenueAccountItem) UnmarshalBinary(b []byte) error {
	var res SectorRevenueAccountItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}