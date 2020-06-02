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

// ResultListProduct result list product
//
// swagger:model ResultListProduct
type ResultListProduct struct {

	// number of the current page, starting from 1
	// Required: true
	CurrentPage *int32 `json:"currentPage"`

	// navigation links (previous, self, next)
	// Required: true
	Links map[string]string `json:"links"`

	// pages count total
	// Required: true
	PagesTotal *int32 `json:"pagesTotal"`

	// result list
	// Required: true
	Results []*Product `json:"results"`

	// result count of the current page
	// Required: true
	ResultsOfPage *int32 `json:"resultsOfPage"`

	// result count total
	// Required: true
	ResultsTotal *int32 `json:"resultsTotal"`
}

// Validate validates this result list product
func (m *ResultListProduct) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentPage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagesTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultsOfPage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultsTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResultListProduct) validateCurrentPage(formats strfmt.Registry) error {

	if err := validate.Required("currentPage", "body", m.CurrentPage); err != nil {
		return err
	}

	return nil
}

func (m *ResultListProduct) validateLinks(formats strfmt.Registry) error {

	return nil
}

func (m *ResultListProduct) validatePagesTotal(formats strfmt.Registry) error {

	if err := validate.Required("pagesTotal", "body", m.PagesTotal); err != nil {
		return err
	}

	return nil
}

func (m *ResultListProduct) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("results", "body", m.Results); err != nil {
		return err
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResultListProduct) validateResultsOfPage(formats strfmt.Registry) error {

	if err := validate.Required("resultsOfPage", "body", m.ResultsOfPage); err != nil {
		return err
	}

	return nil
}

func (m *ResultListProduct) validateResultsTotal(formats strfmt.Registry) error {

	if err := validate.Required("resultsTotal", "body", m.ResultsTotal); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResultListProduct) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResultListProduct) UnmarshalBinary(b []byte) error {
	var res ResultListProduct
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
