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

// NewbillingTransferCompToTimePackageRequest 产品实例将时间量转为时间包请求信息
//
// swagger:model newbillingTransferCompToTimePackageRequest
type NewbillingTransferCompToTimePackageRequest struct {

	// 到期是否自动续约
	CompAutoRenew int32 `json:"comp_auto_renew,omitempty"`

	// 计费项时长
	CompDuration string `json:"comp_duration,omitempty"`

	// 到期续约时长
	CompRenewDuration string `json:"comp_renew_duration,omitempty"`

	// 通知中的事件时间
	// Format: date-time
	EventTime *strfmt.DateTime `json:"event_time,omitempty"`

	// 外部产品实例ID
	ProdInstIDExt string `json:"prod_inst_id_ext,omitempty"`
}

// Validate validates this newbilling transfer comp to time package request
func (m *NewbillingTransferCompToTimePackageRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingTransferCompToTimePackageRequest) validateEventTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EventTime) { // not required
		return nil
	}

	if err := validate.FormatOf("event_time", "body", "date-time", m.EventTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling transfer comp to time package request based on context it is used
func (m *NewbillingTransferCompToTimePackageRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingTransferCompToTimePackageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingTransferCompToTimePackageRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingTransferCompToTimePackageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
