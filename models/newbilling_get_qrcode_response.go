// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingGetQrcodeResponse 支付二维码返回信息
//
// swagger:model newbillingGetQrcodeResponse
type NewbillingGetQrcodeResponse struct {

	// 支付记录ID
	ChargeID string `json:"charge_id,omitempty"`

	// 支付宝支付的url
	ChargeURL string `json:"charge_url,omitempty"`
}

// Validate validates this newbilling get qrcode response
func (m *NewbillingGetQrcodeResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling get qrcode response based on context it is used
func (m *NewbillingGetQrcodeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingGetQrcodeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingGetQrcodeResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingGetQrcodeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
