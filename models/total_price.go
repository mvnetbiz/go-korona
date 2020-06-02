// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TotalPrice Contains all price-related information
//
// swagger:model TotalPrice
type TotalPrice struct {

	// base price: price of the receipt or itemrow, excluded not included taxes, included discount. example: 19% tax included, 2€ discount included
	// Required: true
	Base *float64 `json:"base"`

	// discount: discount amount of the receipt or item row. base + discount = undiscounted base price.
	Discount float64 `json:"discount,omitempty"`

	// gross price: base + not includes taxes. example: 19% tax included, 7% tax excluded
	Gross float64 `json:"gross,omitempty"`

	// base price: price of the receipt or itemrow, excluded all taxes. example: every taxes excluded
	Net float64 `json:"net,omitempty"`

	// taxPayments: detailed tax information. could also be used to calculate the gross and net price depending on the base price
	// Required: true
	TaxPayments []*TaxPayment `json:"taxPayments"`
}

// Validate validates this total price
func (m *TotalPrice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxPayments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TotalPrice) validateBase(formats strfmt.Registry) error {

	if err := validate.Required("base", "body", m.Base); err != nil {
		return err
	}

	return nil
}

func (m *TotalPrice) validateTaxPayments(formats strfmt.Registry) error {

	if err := validate.Required("taxPayments", "body", m.TaxPayments); err != nil {
		return err
	}

	for i := 0; i < len(m.TaxPayments); i++ {
		if swag.IsZero(m.TaxPayments[i]) { // not required
			continue
		}

		if m.TaxPayments[i] != nil {
			if err := m.TaxPayments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taxPayments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TotalPrice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TotalPrice) UnmarshalBinary(b []byte) error {
	var res TotalPrice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}