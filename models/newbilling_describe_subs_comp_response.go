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

// NewbillingDescribeSubsCompResponse newbilling describe subs comp response
//
// swagger:model newbillingDescribeSubsCompResponse
type NewbillingDescribeSubsCompResponse struct {

	// subs comps
	SubsComps []*NewbillingSubsCompDetail `json:"subs_comps"`
}

// Validate validates this newbilling describe subs comp response
func (m *NewbillingDescribeSubsCompResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubsComps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeSubsCompResponse) validateSubsComps(formats strfmt.Registry) error {
	if swag.IsZero(m.SubsComps) { // not required
		return nil
	}

	for i := 0; i < len(m.SubsComps); i++ {
		if swag.IsZero(m.SubsComps[i]) { // not required
			continue
		}

		if m.SubsComps[i] != nil {
			if err := m.SubsComps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subs_comps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subs_comps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling describe subs comp response based on the context it is used
func (m *NewbillingDescribeSubsCompResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSubsComps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeSubsCompResponse) contextValidateSubsComps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SubsComps); i++ {

		if m.SubsComps[i] != nil {
			if err := m.SubsComps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subs_comps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subs_comps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDescribeSubsCompResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDescribeSubsCompResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingDescribeSubsCompResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
