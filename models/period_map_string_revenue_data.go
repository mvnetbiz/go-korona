// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PeriodMapStringRevenueData period map string revenue data
//
// swagger:model PeriodMapStringRevenueData
type PeriodMapStringRevenueData struct {

	// end
	// Format: date-time
	End strfmt.DateTime `json:"end,omitempty"`

	// results
	Results map[string]RevenueData `json:"results,omitempty"`

	// start
	// Format: date-time
	Start strfmt.DateTime `json:"start,omitempty"`
}

// Validate validates this period map string revenue data
func (m *PeriodMapStringRevenueData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeriodMapStringRevenueData) validateEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PeriodMapStringRevenueData) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for k := range m.Results {

		if err := validate.Required("results"+"."+k, "body", m.Results[k]); err != nil {
			return err
		}
		if val, ok := m.Results[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *PeriodMapStringRevenueData) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PeriodMapStringRevenueData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeriodMapStringRevenueData) UnmarshalBinary(b []byte) error {
	var res PeriodMapStringRevenueData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
