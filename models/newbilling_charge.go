// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewbillingCharge 支付记录
//
// swagger:model newbillingCharge
type NewbillingCharge struct {

	// 接入系统ID
	AccessSysID string `json:"access_sys_id,omitempty"`

	// 账单ID
	BillID string `json:"bill_id,omitempty"`

	// 支付金额
	ChargeAmount float32 `json:"charge_amount,omitempty"`

	// 支付渠道- 余额、直付、外部
	ChargeChannel string `json:"charge_channel,omitempty"`

	// 支付记录ID
	ChargeID string `json:"charge_id,omitempty"`

	// 支付状态
	ChargeStatus string `json:"charge_status,omitempty"`

	// 支付记录创建时间
	// Format: date-time
	CreateTime *strfmt.DateTime `json:"create_time,omitempty"`

	// 支付记录创建人
	CreatedBy string `json:"created_by,omitempty"`

	// 支付货币单位
	CurrencyCode string `json:"currency_code,omitempty"`

	// 额外信息
	ExtraInfo string `json:"extra_info,omitempty"`

	// 付款类型- charging 支付，recharge 充值
	PayType string `json:"pay_type,omitempty"`

	// 备注
	Remarks string `json:"remarks,omitempty"`

	// 第三方支付交易号
	TradeNo string `json:"trade_no,omitempty"`

	// 支付记录更新时间
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"update_time,omitempty"`

	// 用户ID
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this newbilling charge
func (m *NewbillingCharge) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingCharge) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingCharge) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling charge based on context it is used
func (m *NewbillingCharge) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCharge) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCharge) UnmarshalBinary(b []byte) error {
	var res NewbillingCharge
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
