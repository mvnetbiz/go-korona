// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EntryGate entry gate
//
// swagger:model EntryGate
type EntryGate struct {

	// indicates whether the object is active for use or not
	// Read Only: true
	Active *bool `json:"active,omitempty"`

	// has children
	HasChildren bool `json:"hasChildren,omitempty"`

	// global object uuid (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// number of the object, like it is set in backoffice; will be removed when active=false
	Number string `json:"number,omitempty"`

	// parent entry gate
	ParentEntryGate *ModelReference `json:"parentEntryGate,omitempty"`

	// the revision number of the object. revision numbers are unique per object-type. there is is no object of the same type with identical revision numbers.
	// Read Only: true
	Revision int64 `json:"revision,omitempty"`
}

// Validate validates this entry gate
func (m *EntryGate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParentEntryGate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntryGate) validateParentEntryGate(formats strfmt.Registry) error {

	if swag.IsZero(m.ParentEntryGate) { // not required
		return nil
	}

	if m.ParentEntryGate != nil {
		if err := m.ParentEntryGate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentEntryGate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EntryGate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntryGate) UnmarshalBinary(b []byte) error {
	var res EntryGate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
