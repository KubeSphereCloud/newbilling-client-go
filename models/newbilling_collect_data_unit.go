// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewbillingCollectDataUnit newbilling collect data unit
//
// swagger:model newbillingCollectDataUnit
type NewbillingCollectDataUnit struct {

	// 采集值
	Amount float32 `json:"amount,omitempty"`

	// 结束时间
	// Format: date-time
	EndTime strfmt.DateTime `json:"end_time,omitempty"`

	// 开始时间
	// Format: date-time
	StartTime strfmt.DateTime `json:"start_time,omitempty"`
}

// Validate validates this newbilling collect data unit
func (m *NewbillingCollectDataUnit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingCollectDataUnit) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingCollectDataUnit) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling collect data unit based on context it is used
func (m *NewbillingCollectDataUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCollectDataUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCollectDataUnit) UnmarshalBinary(b []byte) error {
	var res NewbillingCollectDataUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
