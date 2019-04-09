// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TransactionInput transaction input
// swagger:model TransactionInput
type TransactionInput struct {

	// hash
	// Required: true
	Hash *string `json:"hash"`

	// index
	// Required: true
	Index *int64 `json:"index"`
}

// Validate validates this transaction input
func (m *TransactionInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionInput) validateHash(formats strfmt.Registry) error {

	if err := validate.Required("hash", "body", m.Hash); err != nil {
		return err
	}

	return nil
}

func (m *TransactionInput) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionInput) UnmarshalBinary(b []byte) error {
	var res TransactionInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
