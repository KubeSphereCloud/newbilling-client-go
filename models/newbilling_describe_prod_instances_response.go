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

// NewbillingDescribeProdInstancesResponse 查询产品实例返回信息
//
// swagger:model newbillingDescribeProdInstancesResponse
type NewbillingDescribeProdInstancesResponse struct {

	// 返回结果列表
	ProdInstanceSet []*NewbillingProdInstance `json:"prod_instance_set"`

	// 匹配条件的结果总数
	Total int64 `json:"total,omitempty"`
}

// Validate validates this newbilling describe prod instances response
func (m *NewbillingDescribeProdInstancesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProdInstanceSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeProdInstancesResponse) validateProdInstanceSet(formats strfmt.Registry) error {
	if swag.IsZero(m.ProdInstanceSet) { // not required
		return nil
	}

	for i := 0; i < len(m.ProdInstanceSet); i++ {
		if swag.IsZero(m.ProdInstanceSet[i]) { // not required
			continue
		}

		if m.ProdInstanceSet[i] != nil {
			if err := m.ProdInstanceSet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prod_instance_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("prod_instance_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling describe prod instances response based on the context it is used
func (m *NewbillingDescribeProdInstancesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProdInstanceSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeProdInstancesResponse) contextValidateProdInstanceSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProdInstanceSet); i++ {

		if m.ProdInstanceSet[i] != nil {
			if err := m.ProdInstanceSet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prod_instance_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("prod_instance_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDescribeProdInstancesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDescribeProdInstancesResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingDescribeProdInstancesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}