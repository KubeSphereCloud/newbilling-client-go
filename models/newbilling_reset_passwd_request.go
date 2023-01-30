// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingResetPasswdRequest 重置密码请求信息
//
// swagger:model newbillingResetPasswdRequest
type NewbillingResetPasswdRequest struct {

	// 重置密码CODE
	Code string `json:"code,omitempty"`

	// 用户邮件-帐号
	Email string `json:"email,omitempty"`

	// 用户新密码
	PasswdNew string `json:"passwd_new,omitempty"`
}

// Validate validates this newbilling reset passwd request
func (m *NewbillingResetPasswdRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling reset passwd request based on context it is used
func (m *NewbillingResetPasswdRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingResetPasswdRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingResetPasswdRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingResetPasswdRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
