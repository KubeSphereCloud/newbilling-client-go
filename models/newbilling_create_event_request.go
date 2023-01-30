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

// NewbillingCreateEventRequest newbilling create event request
//
// swagger:model newbillingCreateEventRequest
type NewbillingCreateEventRequest struct {

	// access sys id
	AccessSysID string `json:"access_sys_id,omitempty"`

	// customer id
	CustomerID string `json:"customer_id,omitempty"`

	// event
	Event string `json:"event,omitempty"`

	// event time
	// Format: date-time
	EventTime strfmt.DateTime `json:"event_time,omitempty"`

	// passback
	Passback string `json:"passback,omitempty"`

	// prod id
	ProdID string `json:"prod_id,omitempty"`

	// prod inst id ext
	ProdInstIDExt string `json:"prod_inst_id_ext,omitempty"`
}

// Validate validates this newbilling create event request
func (m *NewbillingCreateEventRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingCreateEventRequest) validateEventTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EventTime) { // not required
		return nil
	}

	if err := validate.FormatOf("event_time", "body", "date-time", m.EventTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling create event request based on context it is used
func (m *NewbillingCreateEventRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCreateEventRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCreateEventRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingCreateEventRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
