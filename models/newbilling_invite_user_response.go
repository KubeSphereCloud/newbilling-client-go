// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingInviteUserResponse 邀请用户响应信息
//
// swagger:model newbillingInviteUserResponse
type NewbillingInviteUserResponse struct {

	// 邀请用户CODE
	Code string `json:"code,omitempty"`

	// 是否发出邀请
	Ok bool `json:"ok,omitempty"`
}

// Validate validates this newbilling invite user response
func (m *NewbillingInviteUserResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling invite user response based on context it is used
func (m *NewbillingInviteUserResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingInviteUserResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingInviteUserResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingInviteUserResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
