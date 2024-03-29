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

// NewbillingBillingJob 计费任务
//
// swagger:model newbillingBillingJob
type NewbillingBillingJob struct {

	// 计费任务ID
	BllJobID string `json:"bll_job_id,omitempty"`

	// 计费项ID
	CompID string `json:"comp_id,omitempty"`

	// 计费项名称
	CompName string `json:"comp_name,omitempty"`

	// 计费任务创建时间
	// Format: date-time
	CreateTime *strfmt.DateTime `json:"create_time,omitempty"`

	// 调度周期表达式
	CronExpr string `json:"cron_expr,omitempty"`

	// 计费任务结束时间
	// Format: date-time
	EndTime *strfmt.DateTime `json:"end_time,omitempty"`

	// 计费方案ID
	PlanID string `json:"plan_id,omitempty"`

	// 产品实例ID
	ProdInstID string `json:"prod_inst_id,omitempty"`

	// 调度任务ID
	SchedulerJobID string `json:"scheduler_job_id,omitempty"`

	// 计费任务开始时间
	// Format: date-time
	StartTime *strfmt.DateTime `json:"start_time,omitempty"`

	// 任务状态
	Status string `json:"status,omitempty"`

	// 订阅ID
	SubsID string `json:"subs_id,omitempty"`

	// 计费任务更新时间
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"update_time,omitempty"`
}

// Validate validates this newbilling billing job
func (m *NewbillingBillingJob) Validate(formats strfmt.Registry) error {
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

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingBillingJob) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingBillingJob) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingBillingJob) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewbillingBillingJob) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this newbilling billing job based on context it is used
func (m *NewbillingBillingJob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingBillingJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingBillingJob) UnmarshalBinary(b []byte) error {
	var res NewbillingBillingJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
