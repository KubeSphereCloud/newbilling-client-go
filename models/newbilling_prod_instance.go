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

// NewbillingProdInstance 产品实例
//
// swagger:model newbillingProdInstance
type NewbillingProdInstance struct {

	// 应付金额
	AmountPayable float32 `json:"amount_payable,omitempty"`

	// 原价
	Cost float32 `json:"cost,omitempty"`

	// 产品实例创建时间
	// Format: date-time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// 是否自动续费
	IsAutoRenew int32 `json:"is_auto_renew,omitempty"`

	// 订单ID
	OrderID string `json:"order_id,omitempty"`

	// 计费方案ID
	PlanID string `json:"plan_id,omitempty"`

	// 产品ID
	ProdID string `json:"prod_id,omitempty"`

	// 产品实例ID
	ProdInstID string `json:"prod_inst_id,omitempty"`

	// 外部产品实例ID
	ProdInstIDExt string `json:"prod_inst_id_ext,omitempty"`

	// 产品名称
	ProdName string `json:"prod_name,omitempty"`

	// 到期自动续费时长
	RenewDuration string `json:"renew_duration,omitempty"`

	// 产品实例状态
	Status string `json:"status,omitempty"`

	// 产品实例更新时间
	// Format: date-time
	UpdateTime strfmt.DateTime `json:"update_time,omitempty"`
}

// Validate validates this newbilling prod instance
func (m *NewbillingProdInstance) Validate(formats strfmt.Registry) error {
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

func (m *NewbillingProdInstance) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingProdInstance) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling prod instance based on context it is used
func (m *NewbillingProdInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingProdInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingProdInstance) UnmarshalBinary(b []byte) error {
	var res NewbillingProdInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}