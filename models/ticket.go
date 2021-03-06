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

// Ticket ticket
//
// swagger:model Ticket
type Ticket struct {

	// creation date
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// duplicate
	Duplicate bool `json:"duplicate,omitempty"`

	// entry Id
	EntryID int64 `json:"entryId,omitempty"`

	// locked
	Locked bool `json:"locked,omitempty"`

	// locked to
	// Format: date-time
	LockedTo strfmt.DateTime `json:"lockedTo,omitempty"`

	// personalization
	Personalization *TicketPersonalization `json:"personalization,omitempty"`

	// ticket number
	TicketNumber string `json:"ticketNumber,omitempty"`
}

// Validate validates this ticket
func (m *Ticket) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLockedTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersonalization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Ticket) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Ticket) validateLockedTo(formats strfmt.Registry) error {

	if swag.IsZero(m.LockedTo) { // not required
		return nil
	}

	if err := validate.FormatOf("lockedTo", "body", "date-time", m.LockedTo.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Ticket) validatePersonalization(formats strfmt.Registry) error {

	if swag.IsZero(m.Personalization) { // not required
		return nil
	}

	if m.Personalization != nil {
		if err := m.Personalization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("personalization")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Ticket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Ticket) UnmarshalBinary(b []byte) error {
	var res Ticket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
