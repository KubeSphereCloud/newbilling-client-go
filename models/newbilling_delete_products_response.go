// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingDeleteProductsResponse 删除产品返回信息
//
// swagger:model newbillingDeleteProductsResponse
type NewbillingDeleteProductsResponse struct {

	// 产品ID列表
	ProdID []string `json:"prod_id"`
}

// Validate validates this newbilling delete products response
func (m *NewbillingDeleteProductsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling delete products response based on context it is used
func (m *NewbillingDeleteProductsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDeleteProductsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDeleteProductsResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingDeleteProductsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
