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

// NewbillingConsumeOrder 消费订单
//
// swagger:model newbillingConsumeOrder
type NewbillingConsumeOrder struct {

	// access sys id
	AccessSysID string `json:"access_sys_id,omitempty"`

	// amount payable
	AmountPayable float32 `json:"amount_payable,omitempty"`

	// charge status
	ChargeStatus string `json:"charge_status,omitempty"`

	// charge time
	// Format: date-time
	ChargeTime *strfmt.DateTime `json:"charge_time,omitempty"`

	// consume type
	ConsumeType string `json:"consume_type,omitempty"`

	// cost
	Cost float32 `json:"cost,omitempty"`

	// create time
	// Format: date-time
	CreateTime *strfmt.DateTime `json:"create_time,omitempty"`

	// is deleted
	IsDeleted int32 `json:"is_deleted,omitempty"`

	// order user id
	OrderUserID string `json:"order_user_id,omitempty"`

	// order user name
	OrderUserName string `json:"order_user_name,omitempty"`

	// prod id
	ProdID string `json:"prod_id,omitempty"`

	// prod inst consume order id
	ProdInstConsumeOrderID string `json:"prod_inst_consume_order_id,omitempty"`

	// prod inst id
	ProdInstID string `json:"prod_inst_id,omitempty"`

	// prod name
	ProdName string `json:"prod_name,omitempty"`

	// update time
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"update_time,omitempty"`
}

// Validate validates this newbilling consume order
func (m *NewbillingConsumeOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChargeTime(formats); err != nil {
		res = append(res, err)
	}

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

func (m *NewbillingConsumeOrder) validateChargeTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ChargeTime) { // not required
		return nil
	}

	if err := validate.FormatOf("charge_time", "body", "date-time", m.ChargeTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingConsumeOrder) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingConsumeOrder) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling consume order based on context it is used
func (m *NewbillingConsumeOrder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingConsumeOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingConsumeOrder) UnmarshalBinary(b []byte) error {
	var res NewbillingConsumeOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
