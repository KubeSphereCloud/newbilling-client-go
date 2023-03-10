// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingCreateRechargeRequest 创建充值记录请求信息
//
// swagger:model newbillingCreateRechargeRequest
type NewbillingCreateRechargeRequest struct {

	// 充值渠道
	Channel string `json:"channel,omitempty"`

	// Console ID
	ConsoleID string `json:"console_id,omitempty"`

	// 充值币种
	Currency string `json:"currency,omitempty"`

	// 用户ID
	CustomerID string `json:"customer_id,omitempty"`

	// 外部交易号
	OutTradeNo string `json:"out_trade_no,omitempty"`

	// 充值金额
	RechargeAmount float32 `json:"recharge_amount,omitempty"`

	// 充值类型
	RechargeType string `json:"recharge_type,omitempty"`

	// 备注
	Remarks string `json:"remarks,omitempty"`
}

// Validate validates this newbilling create recharge request
func (m *NewbillingCreateRechargeRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling create recharge request based on context it is used
func (m *NewbillingCreateRechargeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCreateRechargeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCreateRechargeRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingCreateRechargeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
