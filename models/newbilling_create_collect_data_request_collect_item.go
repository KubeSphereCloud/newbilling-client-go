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

// NewbillingCreateCollectDataRequestCollectItem newbilling create collect data request collect item
//
// swagger:model newbillingCreateCollectDataRequestCollectItem
type NewbillingCreateCollectDataRequestCollectItem struct {

	// prod inst id ext
	ProdInstIDExt string `json:"prod_inst_id_ext,omitempty"`

	// target map
	TargetMap map[string]NewbillingCollectDataUnits `json:"target_map,omitempty"`
}

// Validate validates this newbilling create collect data request collect item
func (m *NewbillingCreateCollectDataRequestCollectItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTargetMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingCreateCollectDataRequestCollectItem) validateTargetMap(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetMap) { // not required
		return nil
	}

	for k := range m.TargetMap {

		if err := validate.Required("target_map"+"."+k, "body", m.TargetMap[k]); err != nil {
			return err
		}
		if val, ok := m.TargetMap[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("target_map" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("target_map" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling create collect data request collect item based on the context it is used
func (m *NewbillingCreateCollectDataRequestCollectItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTargetMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingCreateCollectDataRequestCollectItem) contextValidateTargetMap(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.TargetMap {

		if val, ok := m.TargetMap[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCreateCollectDataRequestCollectItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCreateCollectDataRequestCollectItem) UnmarshalBinary(b []byte) error {
	var res NewbillingCreateCollectDataRequestCollectItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
