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

// NewbillingStopProdInstanceRequest 停止产品实例计费请求信息
//
// swagger:model newbillingStopProdInstanceRequest
type NewbillingStopProdInstanceRequest struct {

	// disparity price
	DisparityPrice bool `json:"disparity_price,omitempty"`

	// 通知中的事件时间
	// Format: date-time
	EventTime strfmt.DateTime `json:"event_time,omitempty"`

	// not refund
	NotRefund bool `json:"not_refund,omitempty"`

	// 产品实例ID
	ProdInstIDExt string `json:"prod_inst_id_ext,omitempty"`

	// refund all
	RefundAll bool `json:"refund_all,omitempty"`
}

// Validate validates this newbilling stop prod instance request
func (m *NewbillingStopProdInstanceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingStopProdInstanceRequest) validateEventTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EventTime) { // not required
		return nil
	}

	if err := validate.FormatOf("event_time", "body", "date-time", m.EventTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling stop prod instance request based on context it is used
func (m *NewbillingStopProdInstanceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingStopProdInstanceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingStopProdInstanceRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingStopProdInstanceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
