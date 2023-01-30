// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingAccessSystemUserResponse newbilling access system user response
//
// swagger:model newbillingAccessSystemUserResponse
type NewbillingAccessSystemUserResponse struct {

	// ok
	Ok bool `json:"ok,omitempty"`
}

// Validate validates this newbilling access system user response
func (m *NewbillingAccessSystemUserResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling access system user response based on context it is used
func (m *NewbillingAccessSystemUserResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingAccessSystemUserResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingAccessSystemUserResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingAccessSystemUserResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}