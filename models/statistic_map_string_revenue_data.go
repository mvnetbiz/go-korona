// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StatisticMapStringRevenueData statistic map string revenue data
//
// swagger:model StatisticMapStringRevenueData
type StatisticMapStringRevenueData struct {

	// current period
	CurrentPeriod *PeriodMapStringRevenueData `json:"currentPeriod,omitempty"`

	// organizational unit
	OrganizationalUnit *ModelReference `json:"organizationalUnit,omitempty"`

	// previous period
	PreviousPeriod *PeriodMapStringRevenueData `json:"previousPeriod,omitempty"`
}

// Validate validates this statistic map string revenue data
func (m *StatisticMapStringRevenueData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationalUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreviousPeriod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatisticMapStringRevenueData) validateCurrentPeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentPeriod) { // not required
		return nil
	}

	if m.CurrentPeriod != nil {
		if err := m.CurrentPeriod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currentPeriod")
			}
			return err
		}
	}

	return nil
}

func (m *StatisticMapStringRevenueData) validateOrganizationalUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganizationalUnit) { // not required
		return nil
	}

	if m.OrganizationalUnit != nil {
		if err := m.OrganizationalUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organizationalUnit")
			}
			return err
		}
	}

	return nil
}

func (m *StatisticMapStringRevenueData) validatePreviousPeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.PreviousPeriod) { // not required
		return nil
	}

	if m.PreviousPeriod != nil {
		if err := m.PreviousPeriod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("previousPeriod")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatisticMapStringRevenueData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatisticMapStringRevenueData) UnmarshalBinary(b []byte) error {
	var res StatisticMapStringRevenueData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
