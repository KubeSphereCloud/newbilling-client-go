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

// NewbillingChangeProdInstanceConfigResponse newbilling change prod instance config response
//
// swagger:model newbillingChangeProdInstanceConfigResponse
type NewbillingChangeProdInstanceConfigResponse struct {

	// 主订单应付金额
	AmountPayable float32 `json:"amount_payable,omitempty"`

	// 主订单原价
	Cost float32 `json:"cost,omitempty"`

	// 计算差价后各计费项的差价明细,当disparity_price为true时才生效有数据
	DisparityPriceDetails []*NewbillingDisparityPriceDetail `json:"disparity_price_details"`

	// 主订单ID
	OrderID string `json:"order_id,omitempty"`

	// 主订单状态- 待支付、已支付、作废
	OrderStatus string `json:"order_status,omitempty"`
}

// Validate validates this newbilling change prod instance config response
func (m *NewbillingChangeProdInstanceConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisparityPriceDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingChangeProdInstanceConfigResponse) validateDisparityPriceDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.DisparityPriceDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.DisparityPriceDetails); i++ {
		if swag.IsZero(m.DisparityPriceDetails[i]) { // not required
			continue
		}

		if m.DisparityPriceDetails[i] != nil {
			if err := m.DisparityPriceDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disparity_price_details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disparity_price_details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling change prod instance config response based on the context it is used
func (m *NewbillingChangeProdInstanceConfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisparityPriceDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingChangeProdInstanceConfigResponse) contextValidateDisparityPriceDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DisparityPriceDetails); i++ {

		if m.DisparityPriceDetails[i] != nil {
			if err := m.DisparityPriceDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disparity_price_details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disparity_price_details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingChangeProdInstanceConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingChangeProdInstanceConfigResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingChangeProdInstanceConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
