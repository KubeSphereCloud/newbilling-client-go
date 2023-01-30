// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingCreateOrderRequest newbilling create order request
//
// swagger:model newbillingCreateOrderRequest
type NewbillingCreateOrderRequest struct {

	// access sys id
	AccessSysID string `json:"access_sys_id,omitempty"`

	// amount
	Amount float32 `json:"amount,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// expire time
	ExpireTime string `json:"expire_time,omitempty"`

	// notify url
	NotifyURL string `json:"notify_url,omitempty"`

	// out trade no
	OutTradeNo string `json:"out_trade_no,omitempty"`

	// passback
	Passback string `json:"passback,omitempty"`

	// remark
	Remark string `json:"remark,omitempty"`

	// return url
	ReturnURL string `json:"return_url,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`
}

// Validate validates this newbilling create order request
func (m *NewbillingCreateOrderRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling create order request based on context it is used
func (m *NewbillingCreateOrderRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCreateOrderRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCreateOrderRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingCreateOrderRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
