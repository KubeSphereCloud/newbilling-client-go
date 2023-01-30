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

// NewbillingAttribute 属性 - 不带value字段，用于属性的增删改查，配置产品的时候使用。
//
// swagger:model newbillingAttribute
type NewbillingAttribute struct {

	// 属性编码
	AttrCode string `json:"attr_code,omitempty"`

	// 属性ID
	AttrID string `json:"attr_id,omitempty"`

	// 创建时间
	// Format: date-time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// 属性描述
	Description string `json:"description,omitempty"`

	// 是否需要计量
	IsNeedMeter int64 `json:"is_need_meter,omitempty"`

	// 属性名称
	Name string `json:"name,omitempty"`

	// 产品ID
	ProdID string `json:"prod_id,omitempty"`

	// 属性值单位
	Unit string `json:"unit,omitempty"`

	// 更新时间
	// Format: date-time
	UpdateTime strfmt.DateTime `json:"update_time,omitempty"`

	// value
	Value string `json:"value,omitempty"`

	// 属性值类型
	ValueType string `json:"value_type,omitempty"`
}

// Validate validates this newbilling attribute
func (m *NewbillingAttribute) Validate(formats strfmt.Registry) error {
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

func (m *NewbillingAttribute) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingAttribute) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling attribute based on context it is used
func (m *NewbillingAttribute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingAttribute) UnmarshalBinary(b []byte) error {
	var res NewbillingAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}