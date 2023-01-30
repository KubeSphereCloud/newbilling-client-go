// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingPricingMigrationComponent newbilling pricing migration component
//
// swagger:model newbillingPricingMigrationComponent
type NewbillingPricingMigrationComponent struct {

	// 接入系统ID
	AccessSysID string `json:"access_sys_id,omitempty"`

	// 计费模式
	BillingMode string `json:"billing_mode,omitempty"`

	// 计费项Code
	CompCode string `json:"comp_code,omitempty"`

	// 计费项ID
	CompID string `json:"comp_id,omitempty"`

	// 计费项说明
	Description string `json:"description,omitempty"`

	// 计费项名称
	Name string `json:"name,omitempty"`

	// 计费方案ID
	PlanID string `json:"plan_id,omitempty"`

	// 产品ID
	ProdID string `json:"prod_id,omitempty"`
}

// Validate validates this newbilling pricing migration component
func (m *NewbillingPricingMigrationComponent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling pricing migration component based on context it is used
func (m *NewbillingPricingMigrationComponent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingPricingMigrationComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingPricingMigrationComponent) UnmarshalBinary(b []byte) error {
	var res NewbillingPricingMigrationComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}