// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingChargeAlipayReturnRequest 支付宝回调请求信息
//
// swagger:model newbillingChargeAlipayReturnRequest
type NewbillingChargeAlipayReturnRequest struct {

	// access sys id
	AccessSysID string `json:"access_sys_id,omitempty"`

	// amount
	Amount string `json:"amount,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// order id
	OrderID string `json:"order_id,omitempty"`

	// out trade no
	OutTradeNo string `json:"out_trade_no,omitempty"`

	// passback
	Passback string `json:"passback,omitempty"`

	// pay channel
	PayChannel string `json:"pay_channel,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// trade no
	TradeNo string `json:"trade_no,omitempty"`
}

// Validate validates this newbilling charge alipay return request
func (m *NewbillingChargeAlipayReturnRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling charge alipay return request based on context it is used
func (m *NewbillingChargeAlipayReturnRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingChargeAlipayReturnRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingChargeAlipayReturnRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingChargeAlipayReturnRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
