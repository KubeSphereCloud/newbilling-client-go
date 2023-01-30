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

// NewbillingCustomer 客户管理
//
// swagger:model newbillingCustomer
type NewbillingCustomer struct {

	// 接入系统的owner的user id
	AccessSysID string `json:"access_sys_id,omitempty"`

	// 用户余额, 单位元，保留两位小数
	Balance float32 `json:"balance,omitempty"`

	// c id
	CID string `json:"c_id,omitempty"`

	// 公司名称
	CompanyName string `json:"company_name,omitempty"`

	// 货币种类
	Currency string `json:"currency,omitempty"`

	// 客户ID
	CustomerID string `json:"customer_id,omitempty"`

	// 客户类型下的客户来源
	CustomerSource string `json:"customer_source,omitempty"`

	// 客户email
	Email string `json:"email,omitempty"`

	// 是否独立计费
	IsOwnBilling int64 `json:"is_own_billing,omitempty"`

	// 客户名称
	Name string `json:"name,omitempty"`

	// 透支开始时间
	// Format: date-time
	OverdraftBeginTime strfmt.DateTime `json:"overdraft_begin_time,omitempty"`

	// 透支类型
	OverdraftType string `json:"overdraft_type,omitempty"`

	// 透支能力值
	OverdraftValue string `json:"overdraft_value,omitempty"`

	// 父客户Id
	ParentCustomerID string `json:"parent_customer_id,omitempty"`

	// 付费模式 1.先付费 2.后付费
	PayMode int64 `json:"pay_mode,omitempty"`

	// 客户电话
	Phone string `json:"phone,omitempty"`
}

// Validate validates this newbilling customer
func (m *NewbillingCustomer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOverdraftBeginTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingCustomer) validateOverdraftBeginTime(formats strfmt.Registry) error {
	if swag.IsZero(m.OverdraftBeginTime) { // not required
		return nil
	}

	if err := validate.FormatOf("overdraft_begin_time", "body", "date-time", m.OverdraftBeginTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling customer based on context it is used
func (m *NewbillingCustomer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCustomer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCustomer) UnmarshalBinary(b []byte) error {
	var res NewbillingCustomer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}