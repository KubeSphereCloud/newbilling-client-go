// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingMappingComponentResponse 设置计费项转换返回信息
//
// swagger:model newbillingMappingComponentResponse
type NewbillingMappingComponentResponse struct {

	// 需要转换的计费项ID
	CompID string `json:"comp_id,omitempty"`
}

// Validate validates this newbilling mapping component response
func (m *NewbillingMappingComponentResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling mapping component response based on context it is used
func (m *NewbillingMappingComponentResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingMappingComponentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingMappingComponentResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingMappingComponentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
