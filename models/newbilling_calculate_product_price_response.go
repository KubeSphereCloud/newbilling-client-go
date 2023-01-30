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

// NewbillingCalculateProductPriceResponse 返回-获取价格返回信息
//
// swagger:model newbillingCalculateProductPriceResponse
type NewbillingCalculateProductPriceResponse struct {

	// 计费项单价
	ComponentCostMap map[string]NewbillingComponentCost `json:"component_cost_map,omitempty"`

	// 方案
	PlanID string `json:"plan_id,omitempty"`

	// 产品ID。
	ProdID string `json:"prod_id,omitempty"`

	// 折后总价
	TotalAmountPayable float32 `json:"total_amount_payable,omitempty"`

	// 总价
	TotalCost float32 `json:"total_cost,omitempty"`
}

// Validate validates this newbilling calculate product price response
func (m *NewbillingCalculateProductPriceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponentCostMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingCalculateProductPriceResponse) validateComponentCostMap(formats strfmt.Registry) error {
	if swag.IsZero(m.ComponentCostMap) { // not required
		return nil
	}

	for k := range m.ComponentCostMap {

		if err := validate.Required("component_cost_map"+"."+k, "body", m.ComponentCostMap[k]); err != nil {
			return err
		}
		if val, ok := m.ComponentCostMap[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("component_cost_map" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("component_cost_map" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling calculate product price response based on the context it is used
func (m *NewbillingCalculateProductPriceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComponentCostMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingCalculateProductPriceResponse) contextValidateComponentCostMap(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.ComponentCostMap {

		if val, ok := m.ComponentCostMap[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCalculateProductPriceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCalculateProductPriceResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingCalculateProductPriceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}