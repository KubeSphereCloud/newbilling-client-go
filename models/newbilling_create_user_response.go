// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingCreateUserResponse 创建用户帐号返回信息
//
// swagger:model newbillingCreateUserResponse
type NewbillingCreateUserResponse struct {

	// 创建用户生成的用户ID
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this newbilling create user response
func (m *NewbillingCreateUserResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling create user response based on context it is used
func (m *NewbillingCreateUserResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCreateUserResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCreateUserResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingCreateUserResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
