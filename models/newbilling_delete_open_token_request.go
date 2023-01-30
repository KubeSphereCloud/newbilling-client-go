// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingDeleteOpenTokenRequest newbilling delete open token request
//
// swagger:model newbillingDeleteOpenTokenRequest
type NewbillingDeleteOpenTokenRequest struct {

	// 【接入系统编码】
	AccessSysID string `json:"access_sys_id,omitempty"`

	// action
	Action string `json:"action,omitempty"`

	// expires
	Expires string `json:"expires,omitempty"`

	// open name
	OpenName string `json:"open_name,omitempty"`

	// signature
	Signature string `json:"signature,omitempty"`

	// 【user_id】
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this newbilling delete open token request
func (m *NewbillingDeleteOpenTokenRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling delete open token request based on context it is used
func (m *NewbillingDeleteOpenTokenRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDeleteOpenTokenRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDeleteOpenTokenRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingDeleteOpenTokenRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}