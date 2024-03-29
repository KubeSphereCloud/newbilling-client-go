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

// NewbillingCustomerResourceTransaction newbilling customer resource transaction
//
// swagger:model newbillingCustomerResourceTransaction
type NewbillingCustomerResourceTransaction struct {

	// access sys id
	AccessSysID string `json:"access_sys_id,omitempty"`

	// attr id
	AttrID string `json:"attr_id,omitempty"`

	// attr name
	AttrName string `json:"attr_name,omitempty"`

	// cr id
	CrID string `json:"cr_id,omitempty"`

	// create time
	// Format: date-time
	CreateTime *strfmt.DateTime `json:"create_time,omitempty"`

	// crt id
	CrtID string `json:"crt_id,omitempty"`

	// 客户ID
	CustomerID string `json:"customer_id,omitempty"`

	// order id
	OrderID string `json:"order_id,omitempty"`

	// prod inst id ext
	ProdInstIDExt string `json:"prod_inst_id_ext,omitempty"`

	// trans type
	TransType string `json:"trans_type,omitempty"`

	// trans value
	TransValue string `json:"trans_value,omitempty"`

	// unit
	Unit string `json:"unit,omitempty"`
}

// Validate validates this newbilling customer resource transaction
func (m *NewbillingCustomerResourceTransaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingCustomerResourceTransaction) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling customer resource transaction based on context it is used
func (m *NewbillingCustomerResourceTransaction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCustomerResourceTransaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCustomerResourceTransaction) UnmarshalBinary(b []byte) error {
	var res NewbillingCustomerResourceTransaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
