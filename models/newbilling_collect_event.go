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

// NewbillingCollectEvent newbilling collect event
//
// swagger:model newbillingCollectEvent
type NewbillingCollectEvent struct {

	// 接入系统ID
	AccessSysID string `json:"access_sys_id,omitempty"`

	// 采集事件ID
	CollectEventID string `json:"collect_event_id,omitempty"`

	// 创建时间
	// Format: date-time
	CreateTime *strfmt.DateTime `json:"create_time,omitempty"`

	// 结束时间
	// Format: date-time
	EndTime *strfmt.DateTime `json:"end_time,omitempty"`

	// 产品ID
	ProdID string `json:"prod_id,omitempty"`

	// 资源ID
	ProdInstIDExt string `json:"prod_inst_id_ext,omitempty"`

	// 重试次数
	RetryTimes int32 `json:"retry_times,omitempty"`

	// 开始时间
	// Format: date-time
	StartTime *strfmt.DateTime `json:"start_time,omitempty"`

	// 事件状态
	Status string `json:"status,omitempty"`

	// 状态变更时间
	// Format: date-time
	StatusTime *strfmt.DateTime `json:"status_time,omitempty"`

	// 更新时间
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"update_time,omitempty"`
}

// Validate validates this newbilling collect event
func (m *NewbillingCollectEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusTime(formats); err != nil {
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

func (m *NewbillingCollectEvent) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingCollectEvent) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingCollectEvent) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingCollectEvent) validateStatusTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusTime) { // not required
		return nil
	}

	if err := validate.FormatOf("status_time", "body", "date-time", m.StatusTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingCollectEvent) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling collect event based on context it is used
func (m *NewbillingCollectEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCollectEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCollectEvent) UnmarshalBinary(b []byte) error {
	var res NewbillingCollectEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
