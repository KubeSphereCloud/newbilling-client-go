// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingDescribeRolesResponse 查询角色返回信息
//
// swagger:model newbillingDescribeRolesResponse
type NewbillingDescribeRolesResponse struct {

	// 返回角色列表
	RoleSet []*NewbillingRole `json:"role_set"`

	// 返回总数
	TotalCount int64 `json:"total_count,omitempty"`
}

// Validate validates this newbilling describe roles response
func (m *NewbillingDescribeRolesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoleSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeRolesResponse) validateRoleSet(formats strfmt.Registry) error {
	if swag.IsZero(m.RoleSet) { // not required
		return nil
	}

	for i := 0; i < len(m.RoleSet); i++ {
		if swag.IsZero(m.RoleSet[i]) { // not required
			continue
		}

		if m.RoleSet[i] != nil {
			if err := m.RoleSet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("role_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("role_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling describe roles response based on the context it is used
func (m *NewbillingDescribeRolesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRoleSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeRolesResponse) contextValidateRoleSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RoleSet); i++ {

		if m.RoleSet[i] != nil {
			if err := m.RoleSet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("role_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("role_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDescribeRolesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDescribeRolesResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingDescribeRolesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
