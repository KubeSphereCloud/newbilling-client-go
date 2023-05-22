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

// NewbillingStopCollectRequest newbilling stop collect request
//
// swagger:model newbillingStopCollectRequest
type NewbillingStopCollectRequest struct {

	// 接入系统ID
	AccessSysID string `json:"access_sys_id,omitempty"`

	// 资源ID
	ProdInstIDExt string `json:"prod_inst_id_ext,omitempty"`

	// 实例停止时间
	// Format: date-time
	StopTime *strfmt.DateTime `json:"stop_time,omitempty"`
}

// Validate validates this newbilling stop collect request
func (m *NewbillingStopCollectRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStopTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingStopCollectRequest) validateStopTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StopTime) { // not required
		return nil
	}

	if err := validate.FormatOf("stop_time", "body", "date-time", m.StopTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling stop collect request based on context it is used
func (m *NewbillingStopCollectRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingStopCollectRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingStopCollectRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingStopCollectRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
