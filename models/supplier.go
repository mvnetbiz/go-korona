// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Supplier supplier
//
// swagger:model Supplier
type Supplier struct {

	// indicates whether the object is active for use or not
	// Read Only: true
	Active *bool `json:"active,omitempty"`

	// additional information
	AdditionalInformation string `json:"additionalInformation,omitempty"`

	// address
	Address *ProductTransferInvolvedPartyInformationData `json:"address,omitempty"`

	// contact
	Contact *SupplierContact `json:"contact,omitempty"`

	// contact person
	ContactPerson *SupplierContactPerson `json:"contactPerson,omitempty"`

	// customer number
	CustomerNumber string `json:"customerNumber,omitempty"`

	// global object uuid (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// number of the object, like it is set in backoffice; will be removed when active=false
	Number string `json:"number,omitempty"`

	// payment
	Payment *SupplierPaymentInformation `json:"payment,omitempty"`

	// return address
	ReturnAddress *ProductTransferInvolvedPartyInformationData `json:"returnAddress,omitempty"`

	// the revision number of the object. revision numbers are unique per object-type. there is is no object of the same type with identical revision numbers.
	// Read Only: true
	Revision int64 `json:"revision,omitempty"`
}

// Validate validates this supplier
func (m *Supplier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactPerson(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Supplier) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *Supplier) validateContact(formats strfmt.Registry) error {

	if swag.IsZero(m.Contact) { // not required
		return nil
	}

	if m.Contact != nil {
		if err := m.Contact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contact")
			}
			return err
		}
	}

	return nil
}

func (m *Supplier) validateContactPerson(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactPerson) { // not required
		return nil
	}

	if m.ContactPerson != nil {
		if err := m.ContactPerson.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contactPerson")
			}
			return err
		}
	}

	return nil
}

func (m *Supplier) validatePayment(formats strfmt.Registry) error {

	if swag.IsZero(m.Payment) { // not required
		return nil
	}

	if m.Payment != nil {
		if err := m.Payment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment")
			}
			return err
		}
	}

	return nil
}

func (m *Supplier) validateReturnAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.ReturnAddress) { // not required
		return nil
	}

	if m.ReturnAddress != nil {
		if err := m.ReturnAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("returnAddress")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Supplier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Supplier) UnmarshalBinary(b []byte) error {
	var res Supplier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}