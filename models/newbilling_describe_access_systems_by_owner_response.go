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

// NewbillingDescribeAccessSystemsByOwnerResponse 查询用户返回信息
//
// swagger:model newbillingDescribeAccessSystemsByOwnerResponse
type NewbillingDescribeAccessSystemsByOwnerResponse struct {

	// 返回用户列表
	OwnerSet []*NewbillingAccessSystemOwner `json:"owner_set"`

	// 返回总数
	TotalCount int64 `json:"total_count,omitempty"`
}

// Validate validates this newbilling describe access systems by owner response
func (m *NewbillingDescribeAccessSystemsByOwnerResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOwnerSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeAccessSystemsByOwnerResponse) validateOwnerSet(formats strfmt.Registry) error {
	if swag.IsZero(m.OwnerSet) { // not required
		return nil
	}

	for i := 0; i < len(m.OwnerSet); i++ {
		if swag.IsZero(m.OwnerSet[i]) { // not required
			continue
		}

		if m.OwnerSet[i] != nil {
			if err := m.OwnerSet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("owner_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("owner_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling describe access systems by owner response based on the context it is used
func (m *NewbillingDescribeAccessSystemsByOwnerResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOwnerSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeAccessSystemsByOwnerResponse) contextValidateOwnerSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OwnerSet); i++ {

		if m.OwnerSet[i] != nil {
			if err := m.OwnerSet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("owner_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("owner_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDescribeAccessSystemsByOwnerResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDescribeAccessSystemsByOwnerResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingDescribeAccessSystemsByOwnerResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
