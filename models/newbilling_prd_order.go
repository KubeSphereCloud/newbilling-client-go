// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewbillingPrdOrder 主订单
//
// swagger:model newbillingPrdOrder
type NewbillingPrdOrder struct {

	// 接入系统ID
	AccessSysID string `json:"access_sys_id,omitempty"`

	// amount payable
	AmountPayable float32 `json:"amount_payable,omitempty"`

	// 支付时间
	// Format: date-time
	ChargeTime *strfmt.DateTime `json:"charge_time,omitempty"`

	// cost
	Cost float32 `json:"cost,omitempty"`

	// 主订单创建时间
	// Format: date-time
	CreateTime *strfmt.DateTime `json:"create_time,omitempty"`

	// 自定义信息
	ExtraInfo string `json:"extra_info,omitempty"`

	// order amount
	OrderAmount float32 `json:"order_amount,omitempty"`

	// order cost
	OrderCost float32 `json:"order_cost,omitempty"`

	// 主订单ID
	OrderID string `json:"order_id,omitempty"`

	// 主订单状态- 待支付、已支付、作废
	OrderStatus string `json:"order_status,omitempty"`

	// order type
	OrderType string `json:"order_type,omitempty"`

	// 主订单用户ID
	OrderUserID string `json:"order_user_id,omitempty"`

	// 主订单用户名称
	OrderUserName string `json:"order_user_name,omitempty"`

	// prod insts
	ProdInsts []*NewbillingProdInst `json:"prod_insts"`

	// prods
	Prods []*Newbillingprod `json:"prods"`

	// 主订单更新时间
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"update_time,omitempty"`
}

// Validate validates this newbilling prd order
func (m *NewbillingPrdOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChargeTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProdInsts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProds(formats); err != nil {
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

func (m *NewbillingPrdOrder) validateChargeTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ChargeTime) { // not required
		return nil
	}

	if err := validate.FormatOf("charge_time", "body", "date-time", m.ChargeTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingPrdOrder) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingPrdOrder) validateProdInsts(formats strfmt.Registry) error {
	if swag.IsZero(m.ProdInsts) { // not required
		return nil
	}

	for i := 0; i < len(m.ProdInsts); i++ {
		if swag.IsZero(m.ProdInsts[i]) { // not required
			continue
		}

		if m.ProdInsts[i] != nil {
			if err := m.ProdInsts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prod_insts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("prod_insts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NewbillingPrdOrder) validateProds(formats strfmt.Registry) error {
	if swag.IsZero(m.Prods) { // not required
		return nil
	}

	for i := 0; i < len(m.Prods); i++ {
		if swag.IsZero(m.Prods[i]) { // not required
			continue
		}

		if m.Prods[i] != nil {
			if err := m.Prods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("prods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NewbillingPrdOrder) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this newbilling prd order based on the context it is used
func (m *NewbillingPrdOrder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProdInsts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingPrdOrder) contextValidateProdInsts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProdInsts); i++ {

		if m.ProdInsts[i] != nil {
			if err := m.ProdInsts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prod_insts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("prod_insts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NewbillingPrdOrder) contextValidateProds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Prods); i++ {

		if m.Prods[i] != nil {
			if err := m.Prods[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("prods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingPrdOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingPrdOrder) UnmarshalBinary(b []byte) error {
	var res NewbillingPrdOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
