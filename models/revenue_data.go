// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RevenueData revenue data
//
// swagger:model RevenueData
type RevenueData struct {

	// average gross value
	AverageGrossValue float64 `json:"averageGrossValue,omitempty"`

	// average net value
	AverageNetValue float64 `json:"averageNetValue,omitempty"`

	// customers total
	CustomersTotal int64 `json:"customersTotal,omitempty"`

	// gross value
	GrossValue float64 `json:"grossValue,omitempty"`

	// net value
	NetValue float64 `json:"netValue,omitempty"`
}

// Validate validates this revenue data
func (m *RevenueData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RevenueData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RevenueData) UnmarshalBinary(b []byte) error {
	var res RevenueData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
