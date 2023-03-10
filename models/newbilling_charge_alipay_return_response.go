// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingChargeAlipayReturnResponse 支付宝回调返回信息
//
// swagger:model newbillingChargeAlipayReturnResponse
type NewbillingChargeAlipayReturnResponse struct {

	// charge id
	ChargeID string `json:"charge_id,omitempty"`
}

// Validate validates this newbilling charge alipay return response
func (m *NewbillingChargeAlipayReturnResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling charge alipay return response based on context it is used
func (m *NewbillingChargeAlipayReturnResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingChargeAlipayReturnResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingChargeAlipayReturnResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingChargeAlipayReturnResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
