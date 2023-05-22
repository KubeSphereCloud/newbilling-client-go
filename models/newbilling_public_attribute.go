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

// NewbillingPublicAttribute newbilling public attribute
//
// swagger:model newbillingPublicAttribute
type NewbillingPublicAttribute struct {

	// billing mode
	BillingMode string `json:"billing_mode,omitempty"`

	// create time
	// Format: date-time
	CreateTime *strfmt.DateTime `json:"create_time,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// public attr id
	PublicAttrID string `json:"public_attr_id,omitempty"`

	// unit
	Unit string `json:"unit,omitempty"`

	// update time
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"update_time,omitempty"`

	// value type
	ValueType string `json:"value_type,omitempty"`
}

// Validate validates this newbilling public attribute
func (m *NewbillingPublicAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingPublicAttribute) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingPublicAttribute) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling public attribute based on context it is used
func (m *NewbillingPublicAttribute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingPublicAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingPublicAttribute) UnmarshalBinary(b []byte) error {
	var res NewbillingPublicAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
