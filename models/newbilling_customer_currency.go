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

// NewbillingCustomerCurrency newbilling customer currency
//
// swagger:model newbillingCustomerCurrency
type NewbillingCustomerCurrency struct {

	// access sys id
	AccessSysID string `json:"access_sys_id,omitempty"`

	// cc id
	CcID string `json:"cc_id,omitempty"`

	// currency type
	CurrencyType string `json:"currency_type,omitempty"`

	// currency value
	CurrencyValue string `json:"currency_value,omitempty"`

	// customer id
	CustomerID string `json:"customer_id,omitempty"`

	// duration
	Duration string `json:"duration,omitempty"`

	// expire time
	// Format: date-time
	ExpireTime strfmt.DateTime `json:"expire_time,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// start time
	// Format: date-time
	StartTime strfmt.DateTime `json:"start_time,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// used value
	UsedValue string `json:"used_value,omitempty"`
}

// Validate validates this newbilling customer currency
func (m *NewbillingCustomerCurrency) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpireTime(formats); err != nil {
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

func (m *NewbillingCustomerCurrency) validateExpireTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpireTime) { // not required
		return nil
	}

	if err := validate.FormatOf("expire_time", "body", "date-time", m.ExpireTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingCustomerCurrency) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling customer currency based on context it is used
func (m *NewbillingCustomerCurrency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCustomerCurrency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCustomerCurrency) UnmarshalBinary(b []byte) error {
	var res NewbillingCustomerCurrency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}