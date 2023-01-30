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

// NewbillingResourcePackageOrderProduct newbilling resource package order product
//
// swagger:model newbillingResourcePackageOrderProduct
type NewbillingResourcePackageOrderProduct struct {

	// 计费方案
	Plan *NewbillingPlanWithAttrsV2 `json:"plan,omitempty"`

	// 产品ID
	ProdID string `json:"prod_id,omitempty"`

	// 产品名称
	ProdName string `json:"prod_name,omitempty"`
}

// Validate validates this newbilling resource package order product
func (m *NewbillingResourcePackageOrderProduct) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingResourcePackageOrderProduct) validatePlan(formats strfmt.Registry) error {
	if swag.IsZero(m.Plan) { // not required
		return nil
	}

	if m.Plan != nil {
		if err := m.Plan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plan")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this newbilling resource package order product based on the context it is used
func (m *NewbillingResourcePackageOrderProduct) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingResourcePackageOrderProduct) contextValidatePlan(ctx context.Context, formats strfmt.Registry) error {

	if m.Plan != nil {
		if err := m.Plan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingResourcePackageOrderProduct) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingResourcePackageOrderProduct) UnmarshalBinary(b []byte) error {
	var res NewbillingResourcePackageOrderProduct
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
