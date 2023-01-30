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

// NewbillingChangeProdInstanceConfigRequest 更改产品实例配置请求信息
//
// swagger:model newbillingChangeProdInstanceConfigRequest
type NewbillingChangeProdInstanceConfigRequest struct {

	// component infos
	ComponentInfos []*NewbillingComponentInfo `json:"component_infos"`

	// 是否只计算差价 true:只计算差价，不会真正的改配 false:默认值，不计算差价，真正改配
	DisparityPrice bool `json:"disparity_price,omitempty"`

	// draw expire time
	DrawExpireTime bool `json:"draw_expire_time,omitempty"`

	// 通知中的事件时间
	// Format: date-time
	EventTime strfmt.DateTime `json:"event_time,omitempty"`

	// extra info
	ExtraInfo string `json:"extra_info,omitempty"`

	// plan id
	PlanID string `json:"plan_id,omitempty"`

	// 外部产品实例ID
	ProdInstIDExt string `json:"prod_inst_id_ext,omitempty"`
}

// Validate validates this newbilling change prod instance config request
func (m *NewbillingChangeProdInstanceConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponentInfos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingChangeProdInstanceConfigRequest) validateComponentInfos(formats strfmt.Registry) error {
	if swag.IsZero(m.ComponentInfos) { // not required
		return nil
	}

	for i := 0; i < len(m.ComponentInfos); i++ {
		if swag.IsZero(m.ComponentInfos[i]) { // not required
			continue
		}

		if m.ComponentInfos[i] != nil {
			if err := m.ComponentInfos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("component_infos" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("component_infos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NewbillingChangeProdInstanceConfigRequest) validateEventTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EventTime) { // not required
		return nil
	}

	if err := validate.FormatOf("event_time", "body", "date-time", m.EventTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this newbilling change prod instance config request based on the context it is used
func (m *NewbillingChangeProdInstanceConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComponentInfos(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingChangeProdInstanceConfigRequest) contextValidateComponentInfos(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ComponentInfos); i++ {

		if m.ComponentInfos[i] != nil {
			if err := m.ComponentInfos[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("component_infos" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("component_infos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingChangeProdInstanceConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingChangeProdInstanceConfigRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingChangeProdInstanceConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
