// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingDeleteAccessSystemsRequest 删除接入系统请求信息
//
// swagger:model newbillingDeleteAccessSystemsRequest
type NewbillingDeleteAccessSystemsRequest struct {

	// 接入系统ID列表 @gotags: valid:"Required"
	AccessSysID []string `json:"access_sys_id"`
}

// Validate validates this newbilling delete access systems request
func (m *NewbillingDeleteAccessSystemsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling delete access systems request based on context it is used
func (m *NewbillingDeleteAccessSystemsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDeleteAccessSystemsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDeleteAccessSystemsRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingDeleteAccessSystemsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
