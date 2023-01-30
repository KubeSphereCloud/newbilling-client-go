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
)

// NewbillingCreatePrdOrderV2Request 创建主订单请求信息
//
// swagger:model newbillingCreatePrdOrderV2Request
type NewbillingCreatePrdOrderV2Request struct {

	// 接入系统ID
	AccessSysID string `json:"access_sys_id,omitempty"`

	// 描述信息
	Description string `json:"description,omitempty"`

	// 扩展信息
	ExtraInfo string `json:"extra_info,omitempty"`

	// 主订单产品信息列表
	OrderProductSet []*NewbillingOrderProductV2 `json:"order_product_set"`

	// 主订单用户ID,对应【customer_id】
	OrderUserID string `json:"order_user_id,omitempty"`

	// 主订单用户名称
	OrderUserName string `json:"order_user_name,omitempty"`
}

// Validate validates this newbilling create prd order v2 request
func (m *NewbillingCreatePrdOrderV2Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderProductSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingCreatePrdOrderV2Request) validateOrderProductSet(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderProductSet) { // not required
		return nil
	}

	for i := 0; i < len(m.OrderProductSet); i++ {
		if swag.IsZero(m.OrderProductSet[i]) { // not required
			continue
		}

		if m.OrderProductSet[i] != nil {
			if err := m.OrderProductSet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("order_product_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("order_product_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling create prd order v2 request based on the context it is used
func (m *NewbillingCreatePrdOrderV2Request) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrderProductSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingCreatePrdOrderV2Request) contextValidateOrderProductSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OrderProductSet); i++ {

		if m.OrderProductSet[i] != nil {
			if err := m.OrderProductSet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("order_product_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("order_product_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCreatePrdOrderV2Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCreatePrdOrderV2Request) UnmarshalBinary(b []byte) error {
	var res NewbillingCreatePrdOrderV2Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
