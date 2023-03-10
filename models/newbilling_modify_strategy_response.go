// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingModifyStrategyResponse 修改策略返回信息
//
// swagger:model newbillingModifyStrategyResponse
type NewbillingModifyStrategyResponse struct {

	// 策略ID
	StrategyID string `json:"strategy_id,omitempty"`
}

// Validate validates this newbilling modify strategy response
func (m *NewbillingModifyStrategyResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling modify strategy response based on context it is used
func (m *NewbillingModifyStrategyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingModifyStrategyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingModifyStrategyResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingModifyStrategyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
