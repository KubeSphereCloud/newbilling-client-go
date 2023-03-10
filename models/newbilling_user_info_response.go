// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingUserInfoResponse newbilling user info response
//
// swagger:model newbillingUserInfoResponse
type NewbillingUserInfoResponse struct {

	// console id
	ConsoleID string `json:"console_id,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// paid mode
	PaidMode string `json:"paid_mode,omitempty"`

	// root user id
	RootUserID string `json:"root_user_id,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`

	// user name
	UserName string `json:"user_name,omitempty"`
}

// Validate validates this newbilling user info response
func (m *NewbillingUserInfoResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling user info response based on context it is used
func (m *NewbillingUserInfoResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingUserInfoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingUserInfoResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingUserInfoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
