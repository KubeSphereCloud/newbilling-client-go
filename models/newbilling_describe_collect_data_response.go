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

// NewbillingDescribeCollectDataResponse newbilling describe collect data response
//
// swagger:model newbillingDescribeCollectDataResponse
type NewbillingDescribeCollectDataResponse struct {

	// 返回结果列表
	CollectDataSet []*NewbillingCollectData `json:"collect_data_set"`

	// 返回总数
	Total int64 `json:"total,omitempty"`
}

// Validate validates this newbilling describe collect data response
func (m *NewbillingDescribeCollectDataResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollectDataSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeCollectDataResponse) validateCollectDataSet(formats strfmt.Registry) error {
	if swag.IsZero(m.CollectDataSet) { // not required
		return nil
	}

	for i := 0; i < len(m.CollectDataSet); i++ {
		if swag.IsZero(m.CollectDataSet[i]) { // not required
			continue
		}

		if m.CollectDataSet[i] != nil {
			if err := m.CollectDataSet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("collect_data_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("collect_data_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling describe collect data response based on the context it is used
func (m *NewbillingDescribeCollectDataResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCollectDataSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeCollectDataResponse) contextValidateCollectDataSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CollectDataSet); i++ {

		if m.CollectDataSet[i] != nil {
			if err := m.CollectDataSet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("collect_data_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("collect_data_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDescribeCollectDataResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDescribeCollectDataResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingDescribeCollectDataResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
