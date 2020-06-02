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

// CreationTask creation task
//
// swagger:model CreationTask
type CreationTask struct {

	// created
	Created int64 `json:"created,omitempty"`

	// creation time
	// Format: date-time
	CreationTime strfmt.DateTime `json:"creationTime,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// entry Id
	EntryID int64 `json:"entryId,omitempty"`

	// personalization
	Personalization *TicketPersonalization `json:"personalization,omitempty"`

	// product number
	ProductNumber string `json:"productNumber,omitempty"`

	// quantity
	Quantity int64 `json:"quantity,omitempty"`

	// status
	// Enum: [PREPARED RUNNING FINISHED FAILED]
	Status string `json:"status,omitempty"`

	// ticket data
	TicketData *TicketData `json:"ticketData,omitempty"`

	// ticket number start
	TicketNumberStart string `json:"ticketNumberStart,omitempty"`

	// user
	User string `json:"user,omitempty"`
}

// Validate validates this creation task
func (m *CreationTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersonalization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTicketData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreationTask) validateCreationTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("creationTime", "body", "date-time", m.CreationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreationTask) validatePersonalization(formats strfmt.Registry) error {

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

var creationTaskTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PREPARED","RUNNING","FINISHED","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		creationTaskTypeStatusPropEnum = append(creationTaskTypeStatusPropEnum, v)
	}
}

const (

	// CreationTaskStatusPREPARED captures enum value "PREPARED"
	CreationTaskStatusPREPARED string = "PREPARED"

	// CreationTaskStatusRUNNING captures enum value "RUNNING"
	CreationTaskStatusRUNNING string = "RUNNING"

	// CreationTaskStatusFINISHED captures enum value "FINISHED"
	CreationTaskStatusFINISHED string = "FINISHED"

	// CreationTaskStatusFAILED captures enum value "FAILED"
	CreationTaskStatusFAILED string = "FAILED"
)

// prop value enum
func (m *CreationTask) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, creationTaskTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreationTask) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *CreationTask) validateTicketData(formats strfmt.Registry) error {

	if swag.IsZero(m.TicketData) { // not required
		return nil
	}

	if m.TicketData != nil {
		if err := m.TicketData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ticketData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreationTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreationTask) UnmarshalBinary(b []byte) error {
	var res CreationTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
