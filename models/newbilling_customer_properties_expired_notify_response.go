// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingCustomerPropertiesExpiredNotifyResponse newbilling customer properties expired notify response
//
// swagger:model newbillingCustomerPropertiesExpiredNotifyResponse
type NewbillingCustomerPropertiesExpiredNotifyResponse struct {

	// ok
	Ok bool `json:"ok,omitempty"`
}

// Validate validates this newbilling customer properties expired notify response
func (m *NewbillingCustomerPropertiesExpiredNotifyResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling customer properties expired notify response based on context it is used
func (m *NewbillingCustomerPropertiesExpiredNotifyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCustomerPropertiesExpiredNotifyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCustomerPropertiesExpiredNotifyResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingCustomerPropertiesExpiredNotifyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
