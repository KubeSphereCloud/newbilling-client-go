// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingCompIDAndCode newbilling comp Id and code
//
// swagger:model newbillingCompIdAndCode
type NewbillingCompIDAndCode struct {

	// code
	Code string `json:"code,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this newbilling comp Id and code
func (m *NewbillingCompIDAndCode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling comp Id and code based on context it is used
func (m *NewbillingCompIDAndCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCompIDAndCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCompIDAndCode) UnmarshalBinary(b []byte) error {
	var res NewbillingCompIDAndCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
