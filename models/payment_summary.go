// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaymentSummary payment summary
//
// swagger:model PaymentSummary
type PaymentSummary struct {

	// actual amount
	ActualAmount float64 `json:"actualAmount,omitempty"`

	// expected amount
	ExpectedAmount float64 `json:"expectedAmount,omitempty"`

	// payment method
	PaymentMethod *ModelReference `json:"paymentMethod,omitempty"`
}

// Validate validates this payment summary
func (m *PaymentSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentSummary) validatePaymentMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentMethod) { // not required
		return nil
	}

	if m.PaymentMethod != nil {
		if err := m.PaymentMethod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("paymentMethod")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentSummary) UnmarshalBinary(b []byte) error {
	var res PaymentSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
