// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingRefreshOpenTokenRequest newbilling refresh open token request
//
// swagger:model newbillingRefreshOpenTokenRequest
type NewbillingRefreshOpenTokenRequest struct {

	// 接入系统编码
	AccessSysID string `json:"access_sys_id,omitempty"`

	// action
	Action string `json:"action,omitempty"`

	// expires
	Expires string `json:"expires,omitempty"`

	// open name
	OpenName string `json:"open_name,omitempty"`

	// signature
	Signature string `json:"signature,omitempty"`

	// 用户ID
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this newbilling refresh open token request
func (m *NewbillingRefreshOpenTokenRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling refresh open token request based on context it is used
func (m *NewbillingRefreshOpenTokenRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingRefreshOpenTokenRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingRefreshOpenTokenRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingRefreshOpenTokenRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
