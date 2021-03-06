// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CurrencyDenomination currency denomination
//
// swagger:model CurrencyDenomination
type CurrencyDenomination struct {

	// type
	// Enum: [COIN BANKNOTE]
	Type string `json:"type,omitempty"`

	// value
	Value float64 `json:"value,omitempty"`
}

// Validate validates this currency denomination
func (m *CurrencyDenomination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var currencyDenominationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["COIN","BANKNOTE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		currencyDenominationTypeTypePropEnum = append(currencyDenominationTypeTypePropEnum, v)
	}
}

const (

	// CurrencyDenominationTypeCOIN captures enum value "COIN"
	CurrencyDenominationTypeCOIN string = "COIN"

	// CurrencyDenominationTypeBANKNOTE captures enum value "BANKNOTE"
	CurrencyDenominationTypeBANKNOTE string = "BANKNOTE"
)

// prop value enum
func (m *CurrencyDenomination) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, currencyDenominationTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CurrencyDenomination) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CurrencyDenomination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CurrencyDenomination) UnmarshalBinary(b []byte) error {
	var res CurrencyDenomination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
