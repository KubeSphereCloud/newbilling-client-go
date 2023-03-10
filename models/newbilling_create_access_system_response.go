// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingCreateAccessSystemResponse 创建接入系统返回信息
//
// swagger:model newbillingCreateAccessSystemResponse
type NewbillingCreateAccessSystemResponse struct {

	// 接入系统ID
	AccessSysID string `json:"access_sys_id,omitempty"`
}

// Validate validates this newbilling create access system response
func (m *NewbillingCreateAccessSystemResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling create access system response based on context it is used
func (m *NewbillingCreateAccessSystemResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCreateAccessSystemResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCreateAccessSystemResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingCreateAccessSystemResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
