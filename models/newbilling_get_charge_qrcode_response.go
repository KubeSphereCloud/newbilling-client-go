// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingGetChargeQrcodeResponse 创建支付记录返回信息
//
// swagger:model newbillingGetChargeQrcodeResponse
type NewbillingGetChargeQrcodeResponse struct {

	// 支付记录ID
	ChargeID string `json:"charge_id,omitempty"`

	// 支付宝支付的url
	ChargeURL string `json:"charge_url,omitempty"`
}

// Validate validates this newbilling get charge qrcode response
func (m *NewbillingGetChargeQrcodeResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling get charge qrcode response based on context it is used
func (m *NewbillingGetChargeQrcodeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingGetChargeQrcodeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingGetChargeQrcodeResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingGetChargeQrcodeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
