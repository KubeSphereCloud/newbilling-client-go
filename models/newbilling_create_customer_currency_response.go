// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingCreateCustomerCurrencyResponse newbilling create customer currency response
//
// swagger:model newbillingCreateCustomerCurrencyResponse
type NewbillingCreateCustomerCurrencyResponse struct {

	// ok
	Ok bool `json:"ok,omitempty"`
}

// Validate validates this newbilling create customer currency response
func (m *NewbillingCreateCustomerCurrencyResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling create customer currency response based on context it is used
func (m *NewbillingCreateCustomerCurrencyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCreateCustomerCurrencyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCreateCustomerCurrencyResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingCreateCustomerCurrencyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
