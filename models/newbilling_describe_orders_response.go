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

// NewbillingDescribeOrdersResponse newbilling describe orders response
//
// swagger:model newbillingDescribeOrdersResponse
type NewbillingDescribeOrdersResponse struct {

	// 返回结果列表
	OrderSet []*NewbillingOrder `json:"order_set"`

	// 返回总数
	Total int64 `json:"total,omitempty"`
}

// Validate validates this newbilling describe orders response
func (m *NewbillingDescribeOrdersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeOrdersResponse) validateOrderSet(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderSet) { // not required
		return nil
	}

	for i := 0; i < len(m.OrderSet); i++ {
		if swag.IsZero(m.OrderSet[i]) { // not required
			continue
		}

		if m.OrderSet[i] != nil {
			if err := m.OrderSet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("order_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("order_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling describe orders response based on the context it is used
func (m *NewbillingDescribeOrdersResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrderSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeOrdersResponse) contextValidateOrderSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OrderSet); i++ {

		if m.OrderSet[i] != nil {
			if err := m.OrderSet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("order_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("order_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDescribeOrdersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDescribeOrdersResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingDescribeOrdersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
