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

// NewbillingDescribeCustomerResourcesResponse 查询操作返回信息
//
// swagger:model newbillingDescribeCustomerResourcesResponse
type NewbillingDescribeCustomerResourcesResponse struct {

	// 返回操作列表
	Crs []*NewbillingCustomerResource `json:"crs"`

	// 返回总数
	TotalCount int64 `json:"total_count,omitempty"`
}

// Validate validates this newbilling describe customer resources response
func (m *NewbillingDescribeCustomerResourcesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCrs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeCustomerResourcesResponse) validateCrs(formats strfmt.Registry) error {
	if swag.IsZero(m.Crs) { // not required
		return nil
	}

	for i := 0; i < len(m.Crs); i++ {
		if swag.IsZero(m.Crs[i]) { // not required
			continue
		}

		if m.Crs[i] != nil {
			if err := m.Crs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("crs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("crs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling describe customer resources response based on the context it is used
func (m *NewbillingDescribeCustomerResourcesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCrs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeCustomerResourcesResponse) contextValidateCrs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Crs); i++ {

		if m.Crs[i] != nil {
			if err := m.Crs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("crs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("crs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDescribeCustomerResourcesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDescribeCustomerResourcesResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingDescribeCustomerResourcesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}