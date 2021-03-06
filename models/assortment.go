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

// Assortment assortment
//
// swagger:model Assortment
type Assortment struct {

	// indicates whether the object is active for use or not
	// Read Only: true
	Active *bool `json:"active,omitempty"`

	// cost center
	// Read Only: true
	CostCenter *ModelReference `json:"costCenter,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// global object uuid (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)
	ID string `json:"id,omitempty"`

	// last clean up
	// Read Only: true
	// Format: date-time
	LastCleanUp strfmt.DateTime `json:"lastCleanUp,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// number of the object, like it is set in backoffice; will be removed when active=false
	Number string `json:"number,omitempty"`

	// the revision number of the object. revision numbers are unique per object-type. there is is no object of the same type with identical revision numbers.
	// Read Only: true
	Revision int64 `json:"revision,omitempty"`
}

// Validate validates this assortment
func (m *Assortment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCostCenter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastCleanUp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Assortment) validateCostCenter(formats strfmt.Registry) error {

	if swag.IsZero(m.CostCenter) { // not required
		return nil
	}

	if m.CostCenter != nil {
		if err := m.CostCenter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("costCenter")
			}
			return err
		}
	}

	return nil
}

func (m *Assortment) validateLastCleanUp(formats strfmt.Registry) error {

	if swag.IsZero(m.LastCleanUp) { // not required
		return nil
	}

	if err := validate.FormatOf("lastCleanUp", "body", "date-time", m.LastCleanUp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Assortment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Assortment) UnmarshalBinary(b []byte) error {
	var res Assortment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
