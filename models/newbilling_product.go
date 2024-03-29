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

// NewbillingProduct 产品
//
// swagger:model newbillingProduct
type NewbillingProduct struct {

	// 接入系统ID
	AccessSysID string `json:"access_sys_id,omitempty"`

	// 产品目录ID
	CataID string `json:"cata_id,omitempty"`

	// 创建时间
	// Format: date-time
	CreateTime *strfmt.DateTime `json:"create_time,omitempty"`

	// 产品描述
	Description string `json:"description,omitempty"`

	// 生效时间
	// Format: date-time
	EffectiveTime *strfmt.DateTime `json:"effective_time,omitempty"`

	// 失效时间
	// Format: date-time
	ExpirationTime *strfmt.DateTime `json:"expiration_time,omitempty"`

	// 产品名称
	Name string `json:"name,omitempty"`

	// 产品编码
	ProdCode string `json:"prod_code,omitempty"`

	// 产品ID
	ProdID string `json:"prod_id,omitempty"`

	// 备注
	Remark string `json:"remark,omitempty"`

	// 更新时间
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"update_time,omitempty"`
}

// Validate validates this newbilling product
func (m *NewbillingProduct) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationTime(formats); err != nil {
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

func (m *NewbillingProduct) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingProduct) validateEffectiveTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EffectiveTime) { // not required
		return nil
	}

	if err := validate.FormatOf("effective_time", "body", "date-time", m.EffectiveTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingProduct) validateExpirationTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpirationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("expiration_time", "body", "date-time", m.ExpirationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingProduct) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling product based on context it is used
func (m *NewbillingProduct) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingProduct) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingProduct) UnmarshalBinary(b []byte) error {
	var res NewbillingProduct
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
