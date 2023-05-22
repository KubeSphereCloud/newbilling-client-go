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

// NewbillingStopProdInstanceResponse 停止产品实例返回信息
//
// swagger:model newbillingStopProdInstanceResponse
type NewbillingStopProdInstanceResponse struct {

	// 实际退费金额
	AmountPayable float32 `json:"amount_payable,omitempty"`

	// 退费的原价
	Cost float32 `json:"cost,omitempty"`

	// 产品实例ID
	ProdInstIDExt string `json:"prod_inst_id_ext,omitempty"`

	// stop time
	// Format: date-time
	StopTime *strfmt.DateTime `json:"stop_time,omitempty"`
}

// Validate validates this newbilling stop prod instance response
func (m *NewbillingStopProdInstanceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStopTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingStopProdInstanceResponse) validateStopTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StopTime) { // not required
		return nil
	}

	if err := validate.FormatOf("stop_time", "body", "date-time", m.StopTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling stop prod instance response based on context it is used
func (m *NewbillingStopProdInstanceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingStopProdInstanceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingStopProdInstanceResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingStopProdInstanceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
