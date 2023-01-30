// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingCreateActionResponse 创建操作返回信息
//
// swagger:model newbillingCreateActionResponse
type NewbillingCreateActionResponse struct {

	// 操作信息ID
	ActionID string `json:"action_id,omitempty"`
}

// Validate validates this newbilling create action response
func (m *NewbillingCreateActionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling create action response based on context it is used
func (m *NewbillingCreateActionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCreateActionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCreateActionResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingCreateActionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}