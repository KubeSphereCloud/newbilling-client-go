// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingDeleteActionResponse 删除操作返回信息
//
// swagger:model newbillingDeleteActionResponse
type NewbillingDeleteActionResponse struct {

	// 是否删除成功
	Ok bool `json:"ok,omitempty"`
}

// Validate validates this newbilling delete action response
func (m *NewbillingDeleteActionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling delete action response based on context it is used
func (m *NewbillingDeleteActionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDeleteActionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDeleteActionResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingDeleteActionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
