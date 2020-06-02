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

// NotFoundError not found error
//
// swagger:model NotFoundError
type NotFoundError struct {

	// the code for client-side evaluation of the error
	// Enum: [ACCOUNT_NOT_FOUND MODEL_ID_NOT_FOUND NUMBER_NOT_FOUND DATE_NOT_FOUND INDEX_NOT_FOUND CONDITION_MISMATCH]
	Code string `json:"code,omitempty"`

	// a short description of the error in english
	Message string `json:"message,omitempty"`
}

// Validate validates this not found error
func (m *NotFoundError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var notFoundErrorTypeCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACCOUNT_NOT_FOUND","MODEL_ID_NOT_FOUND","NUMBER_NOT_FOUND","DATE_NOT_FOUND","INDEX_NOT_FOUND","CONDITION_MISMATCH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		notFoundErrorTypeCodePropEnum = append(notFoundErrorTypeCodePropEnum, v)
	}
}

const (

	// NotFoundErrorCodeACCOUNTNOTFOUND captures enum value "ACCOUNT_NOT_FOUND"
	NotFoundErrorCodeACCOUNTNOTFOUND string = "ACCOUNT_NOT_FOUND"

	// NotFoundErrorCodeMODELIDNOTFOUND captures enum value "MODEL_ID_NOT_FOUND"
	NotFoundErrorCodeMODELIDNOTFOUND string = "MODEL_ID_NOT_FOUND"

	// NotFoundErrorCodeNUMBERNOTFOUND captures enum value "NUMBER_NOT_FOUND"
	NotFoundErrorCodeNUMBERNOTFOUND string = "NUMBER_NOT_FOUND"

	// NotFoundErrorCodeDATENOTFOUND captures enum value "DATE_NOT_FOUND"
	NotFoundErrorCodeDATENOTFOUND string = "DATE_NOT_FOUND"

	// NotFoundErrorCodeINDEXNOTFOUND captures enum value "INDEX_NOT_FOUND"
	NotFoundErrorCodeINDEXNOTFOUND string = "INDEX_NOT_FOUND"

	// NotFoundErrorCodeCONDITIONMISMATCH captures enum value "CONDITION_MISMATCH"
	NotFoundErrorCodeCONDITIONMISMATCH string = "CONDITION_MISMATCH"
)

// prop value enum
func (m *NotFoundError) validateCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, notFoundErrorTypeCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NotFoundError) validateCode(formats strfmt.Registry) error {

	if swag.IsZero(m.Code) { // not required
		return nil
	}

	// value enum
	if err := m.validateCodeEnum("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NotFoundError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotFoundError) UnmarshalBinary(b []byte) error {
	var res NotFoundError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}