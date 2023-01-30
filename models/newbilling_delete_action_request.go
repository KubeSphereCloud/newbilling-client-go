// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingDeleteActionRequest 删除操作请求信息
//
// swagger:model newbillingDeleteActionRequest
type NewbillingDeleteActionRequest struct {

	// 操作信息ID
	ActionID string `json:"action_id,omitempty"`
}

// Validate validates this newbilling delete action request
func (m *NewbillingDeleteActionRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling delete action request based on context it is used
func (m *NewbillingDeleteActionRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDeleteActionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDeleteActionRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingDeleteActionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
