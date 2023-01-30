// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingCreateEventResponse newbilling create event response
//
// swagger:model newbillingCreateEventResponse
type NewbillingCreateEventResponse struct {

	// event id
	EventID string `json:"event_id,omitempty"`
}

// Validate validates this newbilling create event response
func (m *NewbillingCreateEventResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling create event response based on context it is used
func (m *NewbillingCreateEventResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCreateEventResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCreateEventResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingCreateEventResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}