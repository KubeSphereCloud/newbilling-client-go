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

// NewbillingDescribeCollectEventsResponse newbilling describe collect events response
//
// swagger:model newbillingDescribeCollectEventsResponse
type NewbillingDescribeCollectEventsResponse struct {

	// 返回结果列表
	CollectEventSet []*NewbillingCollectEvent `json:"collect_event_set"`

	// 返回总数
	Total int64 `json:"total,omitempty"`
}

// Validate validates this newbilling describe collect events response
func (m *NewbillingDescribeCollectEventsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollectEventSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeCollectEventsResponse) validateCollectEventSet(formats strfmt.Registry) error {
	if swag.IsZero(m.CollectEventSet) { // not required
		return nil
	}

	for i := 0; i < len(m.CollectEventSet); i++ {
		if swag.IsZero(m.CollectEventSet[i]) { // not required
			continue
		}

		if m.CollectEventSet[i] != nil {
			if err := m.CollectEventSet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("collect_event_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("collect_event_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling describe collect events response based on the context it is used
func (m *NewbillingDescribeCollectEventsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCollectEventSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeCollectEventsResponse) contextValidateCollectEventSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CollectEventSet); i++ {

		if m.CollectEventSet[i] != nil {
			if err := m.CollectEventSet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("collect_event_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("collect_event_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDescribeCollectEventsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDescribeCollectEventsResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingDescribeCollectEventsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
