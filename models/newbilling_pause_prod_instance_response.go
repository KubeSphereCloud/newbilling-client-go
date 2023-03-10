// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingPauseProdInstanceResponse 【暂停产品实例计费返回信息】
//
// swagger:model newbillingPauseProdInstanceResponse
type NewbillingPauseProdInstanceResponse struct {

	// 【接入系统产品实例ID】
	ProdInstIDExt string `json:"prod_inst_id_ext,omitempty"`
}

// Validate validates this newbilling pause prod instance response
func (m *NewbillingPauseProdInstanceResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling pause prod instance response based on context it is used
func (m *NewbillingPauseProdInstanceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingPauseProdInstanceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingPauseProdInstanceResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingPauseProdInstanceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
