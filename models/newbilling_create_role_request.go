// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingCreateRoleRequest 创建角色请求信息
//
// swagger:model newbillingCreateRoleRequest
type NewbillingCreateRoleRequest struct {

	// 角色描述
	Description string `json:"description,omitempty"`

	// 角色名称
	RoleName string `json:"role_name,omitempty"`

	// role type
	RoleType int64 `json:"role_type,omitempty"`
}

// Validate validates this newbilling create role request
func (m *NewbillingCreateRoleRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling create role request based on context it is used
func (m *NewbillingCreateRoleRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCreateRoleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCreateRoleRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingCreateRoleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
