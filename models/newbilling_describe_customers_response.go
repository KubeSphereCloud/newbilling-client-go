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

// NewbillingDescribeCustomersResponse 查询操作返回信息
//
// swagger:model newbillingDescribeCustomersResponse
type NewbillingDescribeCustomersResponse struct {

	// 返回操作列表
	Customers []*NewbillingCustomer `json:"customers"`

	// 返回总数
	TotalCount int64 `json:"total_count,omitempty"`
}

// Validate validates this newbilling describe customers response
func (m *NewbillingDescribeCustomersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeCustomersResponse) validateCustomers(formats strfmt.Registry) error {
	if swag.IsZero(m.Customers) { // not required
		return nil
	}

	for i := 0; i < len(m.Customers); i++ {
		if swag.IsZero(m.Customers[i]) { // not required
			continue
		}

		if m.Customers[i] != nil {
			if err := m.Customers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling describe customers response based on the context it is used
func (m *NewbillingDescribeCustomersResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeCustomersResponse) contextValidateCustomers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Customers); i++ {

		if m.Customers[i] != nil {
			if err := m.Customers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDescribeCustomersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDescribeCustomersResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingDescribeCustomersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
