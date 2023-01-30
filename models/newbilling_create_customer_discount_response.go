// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingCreateCustomerDiscountResponse 创建折扣返回信息
//
// swagger:model newbillingCreateCustomerDiscountResponse
type NewbillingCreateCustomerDiscountResponse struct {

	// customer discount id
	CustomerDiscountID string `json:"customer_discount_id,omitempty"`
}

// Validate validates this newbilling create customer discount response
func (m *NewbillingCreateCustomerDiscountResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling create customer discount response based on context it is used
func (m *NewbillingCreateCustomerDiscountResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCreateCustomerDiscountResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCreateCustomerDiscountResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingCreateCustomerDiscountResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}