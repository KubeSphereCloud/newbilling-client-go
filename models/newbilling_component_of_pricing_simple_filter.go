// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingComponentOfPricingSimpleFilter newbilling component of pricing simple filter
//
// swagger:model newbillingComponentOfPricingSimpleFilter
type NewbillingComponentOfPricingSimpleFilter struct {

	// 条件ID 编辑时不能为空
	FilterID string `json:"filter_id,omitempty"`

	// 条件左值，规格=（计费项关联的属性的值attr_id）, 有效期=(patr_PackageMonth:月|patr_PackageDay:天）
	LeftValue string `json:"left_value,omitempty"`

	// 操作符
	Operator string `json:"operator,omitempty"`

	// 条件右值
	RightValue string `json:"right_value,omitempty"`
}

// Validate validates this newbilling component of pricing simple filter
func (m *NewbillingComponentOfPricingSimpleFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling component of pricing simple filter based on context it is used
func (m *NewbillingComponentOfPricingSimpleFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingComponentOfPricingSimpleFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingComponentOfPricingSimpleFilter) UnmarshalBinary(b []byte) error {
	var res NewbillingComponentOfPricingSimpleFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}