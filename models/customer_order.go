// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomerOrder customer order
//
// swagger:model CustomerOrder
type CustomerOrder struct {

	// indicates whether the object is active for use or not
	// Read Only: true
	Active *bool `json:"active,omitempty"`

	// booking time
	// Format: date-time
	BookingTime strfmt.DateTime `json:"bookingTime,omitempty"`

	// cashier
	Cashier *ModelReference `json:"cashier,omitempty"`

	// comment
	Comment string `json:"comment,omitempty"`

	// create time
	// Format: date-time
	CreateTime strfmt.DateTime `json:"createTime,omitempty"`

	// customer
	Customer *ModelReference `json:"customer,omitempty"`

	// customer data
	CustomerData *CustomerData `json:"customerData,omitempty"`

	// deposits
	Deposits []*Deposit `json:"deposits"`

	// finish time
	// Format: date-time
	FinishTime strfmt.DateTime `json:"finishTime,omitempty"`

	// global object uuid (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)
	ID string `json:"id,omitempty"`

	// info texts
	InfoTexts []string `json:"infoTexts"`

	// items
	Items []*ReceiptItem `json:"items"`

	// number of the object, like it is set in backoffice; will be removed when active=false
	Number string `json:"number,omitempty"`

	// organizational unit
	OrganizationalUnit *ModelReference `json:"organizationalUnit,omitempty"`

	// pick up time
	// Format: date-time
	PickUpTime strfmt.DateTime `json:"pickUpTime,omitempty"`

	// point of sale
	PointOfSale *ModelReference `json:"pointOfSale,omitempty"`

	// ready for pick up
	ReadyForPickUp bool `json:"readyForPickUp,omitempty"`

	// the revision number of the object. revision numbers are unique per object-type. there is is no object of the same type with identical revision numbers.
	// Read Only: true
	Revision int64 `json:"revision,omitempty"`

	// warehouse
	Warehouse *ModelReference `json:"warehouse,omitempty"`
}

// Validate validates this customer order
func (m *CustomerOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBookingTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCashier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeposits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinishTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationalUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePickUpTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePointOfSale(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWarehouse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerOrder) validateBookingTime(formats strfmt.Registry) error {

	if swag.IsZero(m.BookingTime) { // not required
		return nil
	}

	if err := validate.FormatOf("bookingTime", "body", "date-time", m.BookingTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomerOrder) validateCashier(formats strfmt.Registry) error {

	if swag.IsZero(m.Cashier) { // not required
		return nil
	}

	if m.Cashier != nil {
		if err := m.Cashier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cashier")
			}
			return err
		}
	}

	return nil
}

func (m *CustomerOrder) validateCreateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("createTime", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomerOrder) validateCustomer(formats strfmt.Registry) error {

	if swag.IsZero(m.Customer) { // not required
		return nil
	}

	if m.Customer != nil {
		if err := m.Customer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer")
			}
			return err
		}
	}

	return nil
}

func (m *CustomerOrder) validateCustomerData(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerData) { // not required
		return nil
	}

	if m.CustomerData != nil {
		if err := m.CustomerData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customerData")
			}
			return err
		}
	}

	return nil
}

func (m *CustomerOrder) validateDeposits(formats strfmt.Registry) error {

	if swag.IsZero(m.Deposits) { // not required
		return nil
	}

	for i := 0; i < len(m.Deposits); i++ {
		if swag.IsZero(m.Deposits[i]) { // not required
			continue
		}

		if m.Deposits[i] != nil {
			if err := m.Deposits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deposits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomerOrder) validateFinishTime(formats strfmt.Registry) error {

	if swag.IsZero(m.FinishTime) { // not required
		return nil
	}

	if err := validate.FormatOf("finishTime", "body", "date-time", m.FinishTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomerOrder) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomerOrder) validateOrganizationalUnit(formats strfmt.Registry) error {

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

func (m *CustomerOrder) validatePickUpTime(formats strfmt.Registry) error {

	if swag.IsZero(m.PickUpTime) { // not required
		return nil
	}

	if err := validate.FormatOf("pickUpTime", "body", "date-time", m.PickUpTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomerOrder) validatePointOfSale(formats strfmt.Registry) error {

	if swag.IsZero(m.PointOfSale) { // not required
		return nil
	}

	if m.PointOfSale != nil {
		if err := m.PointOfSale.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pointOfSale")
			}
			return err
		}
	}

	return nil
}

func (m *CustomerOrder) validateWarehouse(formats strfmt.Registry) error {

	if swag.IsZero(m.Warehouse) { // not required
		return nil
	}

	if m.Warehouse != nil {
		if err := m.Warehouse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("warehouse")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomerOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerOrder) UnmarshalBinary(b []byte) error {
	var res CustomerOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
