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

// NewbillingPlanIDAndCode newbilling plan Id and code
//
// swagger:model newbillingPlanIdAndCode
type NewbillingPlanIDAndCode struct {

	// code
	Code string `json:"code,omitempty"`

	// comps
	Comps []*NewbillingCompIDAndCode `json:"comps"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this newbilling plan Id and code
func (m *NewbillingPlanIDAndCode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingPlanIDAndCode) validateComps(formats strfmt.Registry) error {
	if swag.IsZero(m.Comps) { // not required
		return nil
	}

	for i := 0; i < len(m.Comps); i++ {
		if swag.IsZero(m.Comps[i]) { // not required
			continue
		}

		if m.Comps[i] != nil {
			if err := m.Comps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("comps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("comps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling plan Id and code based on the context it is used
func (m *NewbillingPlanIDAndCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingPlanIDAndCode) contextValidateComps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Comps); i++ {

		if m.Comps[i] != nil {
			if err := m.Comps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("comps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("comps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingPlanIDAndCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingPlanIDAndCode) UnmarshalBinary(b []byte) error {
	var res NewbillingPlanIDAndCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
