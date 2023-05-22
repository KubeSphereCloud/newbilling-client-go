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

// NewbillingPlan 方案
//
// swagger:model newbillingPlan
type NewbillingPlan struct {

	// 创建时间
	// Format: date-time
	CreateTime *strfmt.DateTime `json:"create_time,omitempty"`

	// 方案描述
	Description string `json:"description,omitempty"`

	// 方案编码
	PlanCode string `json:"plan_code,omitempty"`

	// 方案ID
	PlanID string `json:"plan_id,omitempty"`

	// 方案名称
	PlanName string `json:"plan_name,omitempty"`

	// 定价方式
	PricingMethod string `json:"pricing_method,omitempty"`

	// 产品ID
	ProdID string `json:"prod_id,omitempty"`

	// 状态
	Status string `json:"status,omitempty"`

	// 状态更新时间
	// Format: date-time
	StatusTime *strfmt.DateTime `json:"status_time,omitempty"`

	// 更新ID
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"update_time,omitempty"`
}

// Validate validates this newbilling plan
func (m *NewbillingPlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusTime(formats); err != nil {
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

func (m *NewbillingPlan) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingPlan) validateStatusTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusTime) { // not required
		return nil
	}

	if err := validate.FormatOf("status_time", "body", "date-time", m.StatusTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingPlan) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling plan based on context it is used
func (m *NewbillingPlan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingPlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingPlan) UnmarshalBinary(b []byte) error {
	var res NewbillingPlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
