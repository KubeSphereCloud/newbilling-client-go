// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingRenewProdInstanceResponse 续订产品实例返回信息
//
// swagger:model newbillingRenewProdInstanceResponse
type NewbillingRenewProdInstanceResponse struct {

	// 订单应付金额
	AmountPayable float32 `json:"amount_payable,omitempty"`

	// 订单原价
	Cost float32 `json:"cost,omitempty"`

	// order id
	OrderID string `json:"order_id,omitempty"`

	// 产品实例ID
	ProdInstIDExt string `json:"prod_inst_id_ext,omitempty"`
}

// Validate validates this newbilling renew prod instance response
func (m *NewbillingRenewProdInstanceResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling renew prod instance response based on context it is used
func (m *NewbillingRenewProdInstanceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingRenewProdInstanceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingRenewProdInstanceResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingRenewProdInstanceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}