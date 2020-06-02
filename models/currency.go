// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Currency currency
//
// swagger:model Currency
type Currency struct {

	// indicates whether the object is active for use or not
	// Read Only: true
	Active *bool `json:"active,omitempty"`

	// cent name
	CentName string `json:"centName,omitempty"`

	// decimal places
	DecimalPlaces int32 `json:"decimalPlaces,omitempty"`

	// denominations
	Denominations []*CurrencyDenomination `json:"denominations"`

	// exchange rates
	ExchangeRates []*CurrencyExchangeRate `json:"exchangeRates"`

	// global object uuid (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)
	ID string `json:"id,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// number of the object, like it is set in backoffice; will be removed when active=false
	Number string `json:"number,omitempty"`

	// the revision number of the object. revision numbers are unique per object-type. there is is no object of the same type with identical revision numbers.
	// Read Only: true
	Revision int64 `json:"revision,omitempty"`

	// symbol
	Symbol string `json:"symbol,omitempty"`

	// system currency
	SystemCurrency bool `json:"systemCurrency,omitempty"`
}

// Validate validates this currency
func (m *Currency) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDenominations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExchangeRates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Currency) validateDenominations(formats strfmt.Registry) error {

	if swag.IsZero(m.Denominations) { // not required
		return nil
	}

	for i := 0; i < len(m.Denominations); i++ {
		if swag.IsZero(m.Denominations[i]) { // not required
			continue
		}

		if m.Denominations[i] != nil {
			if err := m.Denominations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("denominations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Currency) validateExchangeRates(formats strfmt.Registry) error {

	if swag.IsZero(m.ExchangeRates) { // not required
		return nil
	}

	for i := 0; i < len(m.ExchangeRates); i++ {
		if swag.IsZero(m.ExchangeRates[i]) { // not required
			continue
		}

		if m.ExchangeRates[i] != nil {
			if err := m.ExchangeRates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("exchangeRates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Currency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Currency) UnmarshalBinary(b []byte) error {
	var res Currency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
