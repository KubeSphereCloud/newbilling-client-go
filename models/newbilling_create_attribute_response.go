// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingCreateAttributeResponse 创建属性返回信息
//
// swagger:model newbillingCreateAttributeResponse
type NewbillingCreateAttributeResponse struct {

	// 属性ID
	AttrID string `json:"attr_id,omitempty"`
}

// Validate validates this newbilling create attribute response
func (m *NewbillingCreateAttributeResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling create attribute response based on context it is used
func (m *NewbillingCreateAttributeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCreateAttributeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCreateAttributeResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingCreateAttributeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
