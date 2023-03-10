// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingStrategyPricing newbilling strategy pricing
//
// swagger:model newbillingStrategyPricing
type NewbillingStrategyPricing struct {

	// strategy
	Strategy *NewbillingStrategy `json:"strategy,omitempty"`

	// strategy final price
	StrategyFinalPrice float32 `json:"strategy_final_price,omitempty"`

	// strategy id
	StrategyID string `json:"strategy_id,omitempty"`

	// strategy price
	StrategyPrice float32 `json:"strategy_price,omitempty"`
}

// Validate validates this newbilling strategy pricing
func (m *NewbillingStrategyPricing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingStrategyPricing) validateStrategy(formats strfmt.Registry) error {
	if swag.IsZero(m.Strategy) { // not required
		return nil
	}

	if m.Strategy != nil {
		if err := m.Strategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("strategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("strategy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this newbilling strategy pricing based on the context it is used
func (m *NewbillingStrategyPricing) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStrategy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingStrategyPricing) contextValidateStrategy(ctx context.Context, formats strfmt.Registry) error {

	if m.Strategy != nil {
		if err := m.Strategy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("strategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("strategy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingStrategyPricing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingStrategyPricing) UnmarshalBinary(b []byte) error {
	var res NewbillingStrategyPricing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
