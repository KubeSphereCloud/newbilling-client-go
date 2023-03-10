// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingMappingComponentRequest 设置计费项转换请求信息
//
// swagger:model newbillingMappingComponentRequest
type NewbillingMappingComponentRequest struct {

	// 需要转换的计费项ID   @gotags: valid:"Required"
	CompID string `json:"comp_id,omitempty"`

	// 转换后的计费项ID   @gotags: valid:"Required"
	MappingCompID string `json:"mapping_comp_id,omitempty"`
}

// Validate validates this newbilling mapping component request
func (m *NewbillingMappingComponentRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling mapping component request based on context it is used
func (m *NewbillingMappingComponentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingMappingComponentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingMappingComponentRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingMappingComponentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
