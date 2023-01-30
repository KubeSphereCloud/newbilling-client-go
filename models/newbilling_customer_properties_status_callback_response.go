// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingCustomerPropertiesStatusCallbackResponse newbilling customer properties status callback response
//
// swagger:model newbillingCustomerPropertiesStatusCallbackResponse
type NewbillingCustomerPropertiesStatusCallbackResponse struct {

	// ok
	Ok bool `json:"ok,omitempty"`
}

// Validate validates this newbilling customer properties status callback response
func (m *NewbillingCustomerPropertiesStatusCallbackResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling customer properties status callback response based on context it is used
func (m *NewbillingCustomerPropertiesStatusCallbackResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCustomerPropertiesStatusCallbackResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCustomerPropertiesStatusCallbackResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingCustomerPropertiesStatusCallbackResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
