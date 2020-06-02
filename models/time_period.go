// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TimePeriod time period
//
// swagger:model TimePeriod
type TimePeriod struct {

	// days
	Days int32 `json:"days,omitempty"`

	// hours
	Hours int32 `json:"hours,omitempty"`

	// minutes
	Minutes int32 `json:"minutes,omitempty"`

	// months
	Months int32 `json:"months,omitempty"`

	// seconds
	Seconds int32 `json:"seconds,omitempty"`

	// weeks
	Weeks int32 `json:"weeks,omitempty"`

	// years
	Years int32 `json:"years,omitempty"`
}

// Validate validates this time period
func (m *TimePeriod) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TimePeriod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimePeriod) UnmarshalBinary(b []byte) error {
	var res TimePeriod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
