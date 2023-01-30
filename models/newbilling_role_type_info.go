// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingRoleTypeInfo newbilling role type info
//
// swagger:model newbillingRoleTypeInfo
type NewbillingRoleTypeInfo struct {

	// role id
	RoleID string `json:"role_id,omitempty"`

	// 角色类型- value = 1 owner, value = 2 普通角色
	RoleType int64 `json:"role_type,omitempty"`
}

// Validate validates this newbilling role type info
func (m *NewbillingRoleTypeInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling role type info based on context it is used
func (m *NewbillingRoleTypeInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingRoleTypeInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingRoleTypeInfo) UnmarshalBinary(b []byte) error {
	var res NewbillingRoleTypeInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}