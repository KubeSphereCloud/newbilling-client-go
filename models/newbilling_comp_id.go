// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingCompID newbilling comp Id
//
// swagger:model newbillingCompId
type NewbillingCompID struct {

	// comp id
	CompID string `json:"comp_id,omitempty"`

	// strategy ids
	StrategyIds []string `json:"strategy_ids"`
}

// Validate validates this newbilling comp Id
func (m *NewbillingCompID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling comp Id based on context it is used
func (m *NewbillingCompID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCompID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCompID) UnmarshalBinary(b []byte) error {
	var res NewbillingCompID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
