// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingCreateCatalogResponse 创建产品目录返回信息
//
// swagger:model newbillingCreateCatalogResponse
type NewbillingCreateCatalogResponse struct {

	// 产品目录ID
	CataID string `json:"cata_id,omitempty"`
}

// Validate validates this newbilling create catalog response
func (m *NewbillingCreateCatalogResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling create catalog response based on context it is used
func (m *NewbillingCreateCatalogResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCreateCatalogResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCreateCatalogResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingCreateCatalogResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}