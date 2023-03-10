// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingAccessSystemInitResponse 用户绑定接入系统请求信息
//
// swagger:model newbillingAccessSystemInitResponse
type NewbillingAccessSystemInitResponse struct {

	// 接入系统ID
	AccessSysID string `json:"access_sys_id,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this newbilling access system init response
func (m *NewbillingAccessSystemInitResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling access system init response based on context it is used
func (m *NewbillingAccessSystemInitResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingAccessSystemInitResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingAccessSystemInitResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingAccessSystemInitResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
