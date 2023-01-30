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

// NewbillingDescribeAccessSystemsResponse 查询接入系统返回信息
//
// swagger:model newbillingDescribeAccessSystemsResponse
type NewbillingDescribeAccessSystemsResponse struct {

	// 返回结果列表
	AccessSystemSet []*NewbillingAccessSystem `json:"access_system_set"`

	// 返回总数
	Total int64 `json:"total,omitempty"`
}

// Validate validates this newbilling describe access systems response
func (m *NewbillingDescribeAccessSystemsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessSystemSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeAccessSystemsResponse) validateAccessSystemSet(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessSystemSet) { // not required
		return nil
	}

	for i := 0; i < len(m.AccessSystemSet); i++ {
		if swag.IsZero(m.AccessSystemSet[i]) { // not required
			continue
		}

		if m.AccessSystemSet[i] != nil {
			if err := m.AccessSystemSet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("access_system_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("access_system_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling describe access systems response based on the context it is used
func (m *NewbillingDescribeAccessSystemsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessSystemSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeAccessSystemsResponse) contextValidateAccessSystemSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AccessSystemSet); i++ {

		if m.AccessSystemSet[i] != nil {
			if err := m.AccessSystemSet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("access_system_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("access_system_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDescribeAccessSystemsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDescribeAccessSystemsResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingDescribeAccessSystemsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
