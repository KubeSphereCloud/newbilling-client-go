// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingBindingRoleActionResponse 角色操作绑定返回信息
//
// swagger:model newbillingBindingRoleActionResponse
type NewbillingBindingRoleActionResponse struct {

	// 是否绑定成功
	Ok bool `json:"ok,omitempty"`
}

// Validate validates this newbilling binding role action response
func (m *NewbillingBindingRoleActionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling binding role action response based on context it is used
func (m *NewbillingBindingRoleActionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingBindingRoleActionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingBindingRoleActionResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingBindingRoleActionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
