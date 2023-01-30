// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingModifyUserResponse 修改用户返回信息
//
// swagger:model newbillingModifyUserResponse
type NewbillingModifyUserResponse struct {

	// 修改的用户ID
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this newbilling modify user response
func (m *NewbillingModifyUserResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling modify user response based on context it is used
func (m *NewbillingModifyUserResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingModifyUserResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingModifyUserResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingModifyUserResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
