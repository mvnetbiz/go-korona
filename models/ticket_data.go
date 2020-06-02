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

// TicketData ticket data
//
// swagger:model TicketData
type TicketData struct {

	// activation delay
	ActivationDelay *TimePeriod `json:"activationDelay,omitempty"`

	// daily valid from
	DailyValidFrom string `json:"dailyValidFrom,omitempty"`

	// daily valid to
	DailyValidTo string `json:"dailyValidTo,omitempty"`

	// days blocked after use
	DaysBlockedAfterUse int32 `json:"daysBlockedAfterUse,omitempty"`

	// entry gate numbers
	EntryGateNumbers []string `json:"entryGateNumbers"`

	// max possible admissions
	MaxPossibleAdmissions int32 `json:"maxPossibleAdmissions,omitempty"`

	// ticket number prefix
	TicketNumberPrefix string `json:"ticketNumberPrefix,omitempty"`

	// valid from
	// Format: date-time
	ValidFrom strfmt.DateTime `json:"validFrom,omitempty"`

	// valid to
	// Format: date-time
	ValidTo strfmt.DateTime `json:"validTo,omitempty"`

	// validity period after entrance
	ValidityPeriodAfterEntrance *TimePeriod `json:"validityPeriodAfterEntrance,omitempty"`

	// validity period after purchase
	ValidityPeriodAfterPurchase *TimePeriod `json:"validityPeriodAfterPurchase,omitempty"`
}

// Validate validates this ticket data
func (m *TicketData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivationDelay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidityPeriodAfterEntrance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidityPeriodAfterPurchase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TicketData) validateActivationDelay(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivationDelay) { // not required
		return nil
	}

	if m.ActivationDelay != nil {
		if err := m.ActivationDelay.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activationDelay")
			}
			return err
		}
	}

	return nil
}

func (m *TicketData) validateValidFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidFrom) { // not required
		return nil
	}

	if err := validate.FormatOf("validFrom", "body", "date-time", m.ValidFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TicketData) validateValidTo(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidTo) { // not required
		return nil
	}

	if err := validate.FormatOf("validTo", "body", "date-time", m.ValidTo.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TicketData) validateValidityPeriodAfterEntrance(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidityPeriodAfterEntrance) { // not required
		return nil
	}

	if m.ValidityPeriodAfterEntrance != nil {
		if err := m.ValidityPeriodAfterEntrance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validityPeriodAfterEntrance")
			}
			return err
		}
	}

	return nil
}

func (m *TicketData) validateValidityPeriodAfterPurchase(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidityPeriodAfterPurchase) { // not required
		return nil
	}

	if m.ValidityPeriodAfterPurchase != nil {
		if err := m.ValidityPeriodAfterPurchase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validityPeriodAfterPurchase")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TicketData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TicketData) UnmarshalBinary(b []byte) error {
	var res TicketData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}