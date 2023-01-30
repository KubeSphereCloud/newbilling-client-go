// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingUpdateProdInstanceRequest 改变产品实例自动续约时长请求信息
//
// swagger:model newbillingUpdateProdInstanceRequest
type NewbillingUpdateProdInstanceRequest struct {

	// 产品实例是否自动续约
	IsAutoRenew int32 `json:"is_auto_renew,omitempty"`

	// 产品实例ID
	ProdInstIDExt string `json:"prod_inst_id_ext,omitempty"`

	// 产品实例自动续约时长
	RenewDuration string `json:"renew_duration,omitempty"`
}

// Validate validates this newbilling update prod instance request
func (m *NewbillingUpdateProdInstanceRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling update prod instance request based on context it is used
func (m *NewbillingUpdateProdInstanceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingUpdateProdInstanceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingUpdateProdInstanceRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingUpdateProdInstanceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
