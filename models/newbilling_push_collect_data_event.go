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

// NewbillingPushCollectDataEvent newbilling push collect data event
//
// swagger:model newbillingPushCollectDataEvent
type NewbillingPushCollectDataEvent struct {

	// 计量属性值
	CollectDataSet []*NewbillingPushCollectDataEventCollectData `json:"collect_data_set"`

	// 采集事件ID
	CollectEventID string `json:"collect_event_id,omitempty"`

	// 结束时间
	// Format: date-time
	EndTime *strfmt.DateTime `json:"end_time,omitempty"`

	// 产品ID
	ProdID string `json:"prod_id,omitempty"`

	// 资源ID
	ProdInstIDExt string `json:"prod_inst_id_ext,omitempty"`

	// 事件触发扣费结果code，0=触发扣费成功 其它code=触发code失败
	ResultCode string `json:"result_code,omitempty"`

	// 非0code 失败原因
	ResultMsg string `json:"result_msg,omitempty"`

	// 开始时间
	// Format: date-time
	StartTime *strfmt.DateTime `json:"start_time,omitempty"`
}

// Validate validates this newbilling push collect data event
func (m *NewbillingPushCollectDataEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollectDataSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingPushCollectDataEvent) validateCollectDataSet(formats strfmt.Registry) error {
	if swag.IsZero(m.CollectDataSet) { // not required
		return nil
	}

	for i := 0; i < len(m.CollectDataSet); i++ {
		if swag.IsZero(m.CollectDataSet[i]) { // not required
			continue
		}

		if m.CollectDataSet[i] != nil {
			if err := m.CollectDataSet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("collect_data_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("collect_data_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NewbillingPushCollectDataEvent) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingPushCollectDataEvent) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this newbilling push collect data event based on the context it is used
func (m *NewbillingPushCollectDataEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCollectDataSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingPushCollectDataEvent) contextValidateCollectDataSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CollectDataSet); i++ {

		if m.CollectDataSet[i] != nil {
			if err := m.CollectDataSet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("collect_data_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("collect_data_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingPushCollectDataEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingPushCollectDataEvent) UnmarshalBinary(b []byte) error {
	var res NewbillingPushCollectDataEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
