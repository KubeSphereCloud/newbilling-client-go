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

// NewbillingDescribeBillingJobsResponse 查询计费任务返回信息
//
// swagger:model newbillingDescribeBillingJobsResponse
type NewbillingDescribeBillingJobsResponse struct {

	// 返回结果列表
	BillingJobSet []*NewbillingBillingJob `json:"billing_job_set"`

	// 匹配条件的结果总数
	Total int64 `json:"total,omitempty"`
}

// Validate validates this newbilling describe billing jobs response
func (m *NewbillingDescribeBillingJobsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingJobSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeBillingJobsResponse) validateBillingJobSet(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingJobSet) { // not required
		return nil
	}

	for i := 0; i < len(m.BillingJobSet); i++ {
		if swag.IsZero(m.BillingJobSet[i]) { // not required
			continue
		}

		if m.BillingJobSet[i] != nil {
			if err := m.BillingJobSet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("billing_job_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("billing_job_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling describe billing jobs response based on the context it is used
func (m *NewbillingDescribeBillingJobsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBillingJobSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeBillingJobsResponse) contextValidateBillingJobSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BillingJobSet); i++ {

		if m.BillingJobSet[i] != nil {
			if err := m.BillingJobSet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("billing_job_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("billing_job_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDescribeBillingJobsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDescribeBillingJobsResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingDescribeBillingJobsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}