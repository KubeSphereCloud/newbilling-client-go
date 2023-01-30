// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingModifyAccessSystemUserResponse 修改接入平台用户返回信息
//
// swagger:model newbillingModifyAccessSystemUserResponse
type NewbillingModifyAccessSystemUserResponse struct {

	// 接入平台ID
	AccessSysID string `json:"access_sys_id,omitempty"`

	// 用户ID
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this newbilling modify access system user response
func (m *NewbillingModifyAccessSystemUserResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling modify access system user response based on context it is used
func (m *NewbillingModifyAccessSystemUserResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingModifyAccessSystemUserResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingModifyAccessSystemUserResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingModifyAccessSystemUserResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
