// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingCancelPrdOrderResponse 取消主订单返回信息
//
// swagger:model newbillingCancelPrdOrderResponse
type NewbillingCancelPrdOrderResponse struct {

	// 主订单ID
	OrderID string `json:"order_id,omitempty"`
}

// Validate validates this newbilling cancel prd order response
func (m *NewbillingCancelPrdOrderResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling cancel prd order response based on context it is used
func (m *NewbillingCancelPrdOrderResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCancelPrdOrderResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCancelPrdOrderResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingCancelPrdOrderResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
