// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewbillingSubscriptionComponent newbilling subscription component
//
// swagger:model newbillingSubscriptionComponent
type NewbillingSubscriptionComponent struct {

	// amount payable
	AmountPayable float32 `json:"amount_payable,omitempty"`

	// attrs
	Attrs []*NewbillingComponentAttribute `json:"attrs"`

	// bill id
	BillID string `json:"bill_id,omitempty"`

	// billing mode
	BillingMode string `json:"billing_mode,omitempty"`

	// comp id
	CompID string `json:"comp_id,omitempty"`

	// comp name
	CompName string `json:"comp_name,omitempty"`

	// cost
	Cost float32 `json:"cost,omitempty"`

	// create time
	// Format: date-time
	CreateTime *strfmt.DateTime `json:"create_time,omitempty"`

	// discount
	Discount float32 `json:"discount,omitempty"`

	// duration
	Duration string `json:"duration,omitempty"`

	// expire time
	// Format: date-time
	ExpireTime *strfmt.DateTime `json:"expire_time,omitempty"`

	// is deleted
	IsDeleted int32 `json:"is_deleted,omitempty"`

	// plan id
	PlanID string `json:"plan_id,omitempty"`

	// prod inst id
	ProdInstID string `json:"prod_inst_id,omitempty"`

	// prod inst id ext
	ProdInstIDExt string `json:"prod_inst_id_ext,omitempty"`

	// start time
	// Format: date-time
	StartTime *strfmt.DateTime `json:"start_time,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// subs component id
	SubsComponentID string `json:"subs_component_id,omitempty"`

	// subs id
	SubsID string `json:"subs_id,omitempty"`

	// unit price
	UnitPrice float32 `json:"unit_price,omitempty"`

	// update time
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"update_time,omitempty"`
}

// Validate validates this newbilling subscription component
func (m *NewbillingSubscriptionComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttrs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpireTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
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

func (m *NewbillingSubscriptionComponent) validateAttrs(formats strfmt.Registry) error {
	if swag.IsZero(m.Attrs) { // not required
		return nil
	}

	for i := 0; i < len(m.Attrs); i++ {
		if swag.IsZero(m.Attrs[i]) { // not required
			continue
		}

		if m.Attrs[i] != nil {
			if err := m.Attrs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attrs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attrs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NewbillingSubscriptionComponent) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingSubscriptionComponent) validateExpireTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpireTime) { // not required
		return nil
	}

	if err := validate.FormatOf("expire_time", "body", "date-time", m.ExpireTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingSubscriptionComponent) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingSubscriptionComponent) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this newbilling subscription component based on the context it is used
func (m *NewbillingSubscriptionComponent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttrs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingSubscriptionComponent) contextValidateAttrs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attrs); i++ {

		if m.Attrs[i] != nil {
			if err := m.Attrs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attrs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attrs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingSubscriptionComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingSubscriptionComponent) UnmarshalBinary(b []byte) error {
	var res NewbillingSubscriptionComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
