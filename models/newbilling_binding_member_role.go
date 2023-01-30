// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingBindingMemberRole 用户角色绑定基本信息
//
// swagger:model newbillingBindingMemberRole
type NewbillingBindingMemberRole struct {

	// 角色ID
	RoleID []string `json:"role_id"`

	// 用户ID
	UserID []string `json:"user_id"`
}

// Validate validates this newbilling binding member role
func (m *NewbillingBindingMemberRole) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling binding member role based on context it is used
func (m *NewbillingBindingMemberRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingBindingMemberRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingBindingMemberRole) UnmarshalBinary(b []byte) error {
	var res NewbillingBindingMemberRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
