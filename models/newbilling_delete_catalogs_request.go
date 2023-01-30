// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingDeleteCatalogsRequest 删除产品目录请求信息
//
// swagger:model newbillingDeleteCatalogsRequest
type NewbillingDeleteCatalogsRequest struct {

	// 接入系统ID仅super user有效，接入系统管理员/成员会默认用当前登陆的接入系统
	AccessSysID string `json:"access_sys_id,omitempty"`

	// 产品目录ID列表
	CataID []string `json:"cata_id"`
}

// Validate validates this newbilling delete catalogs request
func (m *NewbillingDeleteCatalogsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling delete catalogs request based on context it is used
func (m *NewbillingDeleteCatalogsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDeleteCatalogsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDeleteCatalogsRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingDeleteCatalogsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}