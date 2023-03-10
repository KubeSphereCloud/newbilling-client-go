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

// NewbillingOrderProductV2 [包含具体配置的一类产品，例如：2核4G的虚拟主机]
//
// swagger:model newbillingOrderProductV2
type NewbillingOrderProductV2 struct {

	// 计费方案
	Plan *NewbillingPlanWithAttrsV2 `json:"plan,omitempty"`

	// 产品ID
	ProdID string `json:"prod_id,omitempty"`

	// 外部产品实例ID，多个代表多个实例
	ProdInstanceExtArray []*NewbillingProdInstanceExt `json:"prod_instance_ext_array"`

	// 产品名称
	ProdName string `json:"prod_name,omitempty"`
}

// Validate validates this newbilling order product v2
func (m *NewbillingOrderProductV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProdInstanceExtArray(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingOrderProductV2) validatePlan(formats strfmt.Registry) error {
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

func (m *NewbillingOrderProductV2) validateProdInstanceExtArray(formats strfmt.Registry) error {
	if swag.IsZero(m.ProdInstanceExtArray) { // not required
		return nil
	}

	for i := 0; i < len(m.ProdInstanceExtArray); i++ {
		if swag.IsZero(m.ProdInstanceExtArray[i]) { // not required
			continue
		}

		if m.ProdInstanceExtArray[i] != nil {
			if err := m.ProdInstanceExtArray[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prod_instance_ext_array" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("prod_instance_ext_array" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling order product v2 based on the context it is used
func (m *NewbillingOrderProductV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProdInstanceExtArray(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingOrderProductV2) contextValidatePlan(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NewbillingOrderProductV2) contextValidateProdInstanceExtArray(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProdInstanceExtArray); i++ {

		if m.ProdInstanceExtArray[i] != nil {
			if err := m.ProdInstanceExtArray[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prod_instance_ext_array" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("prod_instance_ext_array" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingOrderProductV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingOrderProductV2) UnmarshalBinary(b []byte) error {
	var res NewbillingOrderProductV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
