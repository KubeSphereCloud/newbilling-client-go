// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingCreateAccessSystemRequest 创建接入系统请求信息
//
// swagger:model newbillingCreateAccessSystemRequest
type NewbillingCreateAccessSystemRequest struct {

	// 接入系统编码
	AccessSysCode string `json:"access_sys_code,omitempty"`

	// 邮箱 @gotags: valid:"Required"
	Email string `json:"email,omitempty"`

	// 接入系统名称 @gotags: valid:"Required"
	Name string `json:"name,omitempty"`
}

// Validate validates this newbilling create access system request
func (m *NewbillingCreateAccessSystemRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling create access system request based on context it is used
func (m *NewbillingCreateAccessSystemRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCreateAccessSystemRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCreateAccessSystemRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingCreateAccessSystemRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}