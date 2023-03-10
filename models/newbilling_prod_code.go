// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingProdCode newbilling prod code
//
// swagger:model newbillingProdCode
type NewbillingProdCode struct {

	// plan codes
	PlanCodes []*NewbillingPlanCode `json:"plan_codes"`

	// prod code
	ProdCode string `json:"prod_code,omitempty"`
}

// Validate validates this newbilling prod code
func (m *NewbillingProdCode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlanCodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingProdCode) validatePlanCodes(formats strfmt.Registry) error {
	if swag.IsZero(m.PlanCodes) { // not required
		return nil
	}

	for i := 0; i < len(m.PlanCodes); i++ {
		if swag.IsZero(m.PlanCodes[i]) { // not required
			continue
		}

		if m.PlanCodes[i] != nil {
			if err := m.PlanCodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("plan_codes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("plan_codes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling prod code based on the context it is used
func (m *NewbillingProdCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlanCodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingProdCode) contextValidatePlanCodes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PlanCodes); i++ {

		if m.PlanCodes[i] != nil {
			if err := m.PlanCodes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("plan_codes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("plan_codes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingProdCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingProdCode) UnmarshalBinary(b []byte) error {
	var res NewbillingProdCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
