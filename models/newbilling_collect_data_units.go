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
)

// NewbillingCollectDataUnits newbilling collect data units
//
// swagger:model newbillingCollectDataUnits
type NewbillingCollectDataUnits struct {

	// data set
	DataSet []*NewbillingCollectDataUnit `json:"data_set"`
}

// Validate validates this newbilling collect data units
func (m *NewbillingCollectDataUnits) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingCollectDataUnits) validateDataSet(formats strfmt.Registry) error {
	if swag.IsZero(m.DataSet) { // not required
		return nil
	}

	for i := 0; i < len(m.DataSet); i++ {
		if swag.IsZero(m.DataSet[i]) { // not required
			continue
		}

		if m.DataSet[i] != nil {
			if err := m.DataSet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling collect data units based on the context it is used
func (m *NewbillingCollectDataUnits) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDataSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingCollectDataUnits) contextValidateDataSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DataSet); i++ {

		if m.DataSet[i] != nil {
			if err := m.DataSet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCollectDataUnits) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCollectDataUnits) UnmarshalBinary(b []byte) error {
	var res NewbillingCollectDataUnits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
