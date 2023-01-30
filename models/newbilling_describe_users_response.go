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

// NewbillingDescribeUsersResponse 查询用户返回信息
//
// swagger:model newbillingDescribeUsersResponse
type NewbillingDescribeUsersResponse struct {

	// 返回总数
	TotalCount int64 `json:"total_count,omitempty"`

	// 返回用户列表
	UserSet []*NewbillingUser `json:"user_set"`
}

// Validate validates this newbilling describe users response
func (m *NewbillingDescribeUsersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeUsersResponse) validateUserSet(formats strfmt.Registry) error {
	if swag.IsZero(m.UserSet) { // not required
		return nil
	}

	for i := 0; i < len(m.UserSet); i++ {
		if swag.IsZero(m.UserSet[i]) { // not required
			continue
		}

		if m.UserSet[i] != nil {
			if err := m.UserSet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("user_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("user_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling describe users response based on the context it is used
func (m *NewbillingDescribeUsersResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUserSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeUsersResponse) contextValidateUserSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UserSet); i++ {

		if m.UserSet[i] != nil {
			if err := m.UserSet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("user_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("user_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDescribeUsersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDescribeUsersResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingDescribeUsersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}