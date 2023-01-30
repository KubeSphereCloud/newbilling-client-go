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

// NewbillingDescribeServicesResponse 查询对接服务返回
//
// swagger:model newbillingDescribeServicesResponse
type NewbillingDescribeServicesResponse struct {

	// 返回结果列表
	ServiceSet []*NewbillingService `json:"service_set"`

	// 返回总数
	Total int64 `json:"total,omitempty"`
}

// Validate validates this newbilling describe services response
func (m *NewbillingDescribeServicesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeServicesResponse) validateServiceSet(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceSet) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceSet); i++ {
		if swag.IsZero(m.ServiceSet[i]) { // not required
			continue
		}

		if m.ServiceSet[i] != nil {
			if err := m.ServiceSet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("service_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling describe services response based on the context it is used
func (m *NewbillingDescribeServicesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeServicesResponse) contextValidateServiceSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServiceSet); i++ {

		if m.ServiceSet[i] != nil {
			if err := m.ServiceSet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("service_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDescribeServicesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDescribeServicesResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingDescribeServicesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
