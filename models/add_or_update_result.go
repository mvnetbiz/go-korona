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

// AddOrUpdateResult add or update result
//
// swagger:model AddOrUpdateResult
type AddOrUpdateResult struct {

	// action
	// Enum: [ADDED DELETED QUEUED UPDATED]
	Action string `json:"action,omitempty"`

	// href
	Href string `json:"href,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// number
	Number string `json:"number,omitempty"`

	// revision
	Revision int64 `json:"revision,omitempty"`

	// status
	// Enum: [ERROR OK]
	Status string `json:"status,omitempty"`
}

// Validate validates this add or update result
func (m *AddOrUpdateResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addOrUpdateResultTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ADDED","DELETED","QUEUED","UPDATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addOrUpdateResultTypeActionPropEnum = append(addOrUpdateResultTypeActionPropEnum, v)
	}
}

const (

	// AddOrUpdateResultActionADDED captures enum value "ADDED"
	AddOrUpdateResultActionADDED string = "ADDED"

	// AddOrUpdateResultActionDELETED captures enum value "DELETED"
	AddOrUpdateResultActionDELETED string = "DELETED"

	// AddOrUpdateResultActionQUEUED captures enum value "QUEUED"
	AddOrUpdateResultActionQUEUED string = "QUEUED"

	// AddOrUpdateResultActionUPDATED captures enum value "UPDATED"
	AddOrUpdateResultActionUPDATED string = "UPDATED"
)

// prop value enum
func (m *AddOrUpdateResult) validateActionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addOrUpdateResultTypeActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AddOrUpdateResult) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

var addOrUpdateResultTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ERROR","OK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addOrUpdateResultTypeStatusPropEnum = append(addOrUpdateResultTypeStatusPropEnum, v)
	}
}

const (

	// AddOrUpdateResultStatusERROR captures enum value "ERROR"
	AddOrUpdateResultStatusERROR string = "ERROR"

	// AddOrUpdateResultStatusOK captures enum value "OK"
	AddOrUpdateResultStatusOK string = "OK"
)

// prop value enum
func (m *AddOrUpdateResult) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addOrUpdateResultTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AddOrUpdateResult) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddOrUpdateResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddOrUpdateResult) UnmarshalBinary(b []byte) error {
	var res AddOrUpdateResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
