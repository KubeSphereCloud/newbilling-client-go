// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewbillingRole 角色基本信息
//
// swagger:model newbillingRole
type NewbillingRole struct {

	// 接入平台ID
	AccessSysID string `json:"access_sys_id,omitempty"`

	// 创建时间
	// Format: date-time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// 创建人ID
	CreatedBy string `json:"created_by,omitempty"`

	// 角色描述
	Description string `json:"description,omitempty"`

	// 角色ID
	RoleID string `json:"role_id,omitempty"`

	// 角色名秒
	RoleName string `json:"role_name,omitempty"`

	// 角色类型- value = 1 owner, value = 2 普通角色
	RoleType int64 `json:"role_type,omitempty"`

	// 角色状态- value = 1 可用, value = 2 禁用
	Status int64 `json:"status,omitempty"`

	// 更新时间
	// Format: date-time
	UpdateTime strfmt.DateTime `json:"update_time,omitempty"`

	// 角色的成员数量
	Users int64 `json:"users,omitempty"`
}

// Validate validates this newbilling role
func (m *NewbillingRole) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingRole) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingRole) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling role based on context it is used
func (m *NewbillingRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingRole) UnmarshalBinary(b []byte) error {
	var res NewbillingRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
