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

// NewbillingDescribePlansResponse 查询方案返回信息
//
// swagger:model newbillingDescribePlansResponse
type NewbillingDescribePlansResponse struct {

	// 返回结果列表
	PlanSet []*NewbillingPlan `json:"plan_set"`

	// 返回总数
	Total int64 `json:"total,omitempty"`
}

// Validate validates this newbilling describe plans response
func (m *NewbillingDescribePlansResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlanSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribePlansResponse) validatePlanSet(formats strfmt.Registry) error {
	if swag.IsZero(m.PlanSet) { // not required
		return nil
	}

	for i := 0; i < len(m.PlanSet); i++ {
		if swag.IsZero(m.PlanSet[i]) { // not required
			continue
		}

		if m.PlanSet[i] != nil {
			if err := m.PlanSet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("plan_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("plan_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling describe plans response based on the context it is used
func (m *NewbillingDescribePlansResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlanSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribePlansResponse) contextValidatePlanSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PlanSet); i++ {

		if m.PlanSet[i] != nil {
			if err := m.PlanSet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("plan_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("plan_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDescribePlansResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDescribePlansResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingDescribePlansResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
