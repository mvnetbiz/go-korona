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

// ReceiptItem receipt item
//
// swagger:model ReceiptItem
type ReceiptItem struct {

	// booking time
	// Format: date-time
	BookingTime strfmt.DateTime `json:"bookingTime,omitempty"`

	// commodity group
	CommodityGroup *ModelReference `json:"commodityGroup,omitempty"`

	// delivery date
	// Format: date-time
	DeliveryDate strfmt.DateTime `json:"deliveryDate,omitempty"`

	// discounts
	Discounts []*Discount `json:"discounts"`

	// indent
	Indent int32 `json:"indent,omitempty"`

	// info texts
	InfoTexts []string `json:"infoTexts"`

	// manual price
	ManualPrice bool `json:"manualPrice,omitempty"`

	// product
	Product *ModelReference `json:"product,omitempty"`

	// quantity
	Quantity float64 `json:"quantity,omitempty"`

	// recognition number
	RecognitionNumber string `json:"recognitionNumber,omitempty"`

	// sector
	Sector *ModelReference `json:"sector,omitempty"`

	// serial numbers
	// Unique: true
	SerialNumbers []string `json:"serialNumbers"`

	// total
	Total *TotalPrice `json:"total,omitempty"`
}

// Validate validates this receipt item
func (m *ReceiptItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBookingTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommodityGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeliveryDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerialNumbers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReceiptItem) validateBookingTime(formats strfmt.Registry) error {

	if swag.IsZero(m.BookingTime) { // not required
		return nil
	}

	if err := validate.FormatOf("bookingTime", "body", "date-time", m.BookingTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReceiptItem) validateCommodityGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.CommodityGroup) { // not required
		return nil
	}

	if m.CommodityGroup != nil {
		if err := m.CommodityGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commodityGroup")
			}
			return err
		}
	}

	return nil
}

func (m *ReceiptItem) validateDeliveryDate(formats strfmt.Registry) error {

	if swag.IsZero(m.DeliveryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("deliveryDate", "body", "date-time", m.DeliveryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReceiptItem) validateDiscounts(formats strfmt.Registry) error {

	if swag.IsZero(m.Discounts) { // not required
		return nil
	}

	for i := 0; i < len(m.Discounts); i++ {
		if swag.IsZero(m.Discounts[i]) { // not required
			continue
		}

		if m.Discounts[i] != nil {
			if err := m.Discounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("discounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReceiptItem) validateProduct(formats strfmt.Registry) error {

	if swag.IsZero(m.Product) { // not required
		return nil
	}

	if m.Product != nil {
		if err := m.Product.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("product")
			}
			return err
		}
	}

	return nil
}

func (m *ReceiptItem) validateSector(formats strfmt.Registry) error {

	if swag.IsZero(m.Sector) { // not required
		return nil
	}

	if m.Sector != nil {
		if err := m.Sector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sector")
			}
			return err
		}
	}

	return nil
}

func (m *ReceiptItem) validateSerialNumbers(formats strfmt.Registry) error {

	if swag.IsZero(m.SerialNumbers) { // not required
		return nil
	}

	if err := validate.UniqueItems("serialNumbers", "body", m.SerialNumbers); err != nil {
		return err
	}

	return nil
}

func (m *ReceiptItem) validateTotal(formats strfmt.Registry) error {

	if swag.IsZero(m.Total) { // not required
		return nil
	}

	if m.Total != nil {
		if err := m.Total.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReceiptItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReceiptItem) UnmarshalBinary(b []byte) error {
	var res ReceiptItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
