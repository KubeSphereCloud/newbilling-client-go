// Code generated by go-swagger; DO NOT EDIT.

package subscription_billing_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSubscriptionBillingManagerDescribeBillingJobsParams creates a new SubscriptionBillingManagerDescribeBillingJobsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubscriptionBillingManagerDescribeBillingJobsParams() *SubscriptionBillingManagerDescribeBillingJobsParams {
	return &SubscriptionBillingManagerDescribeBillingJobsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriptionBillingManagerDescribeBillingJobsParamsWithTimeout creates a new SubscriptionBillingManagerDescribeBillingJobsParams object
// with the ability to set a timeout on a request.
func NewSubscriptionBillingManagerDescribeBillingJobsParamsWithTimeout(timeout time.Duration) *SubscriptionBillingManagerDescribeBillingJobsParams {
	return &SubscriptionBillingManagerDescribeBillingJobsParams{
		timeout: timeout,
	}
}

// NewSubscriptionBillingManagerDescribeBillingJobsParamsWithContext creates a new SubscriptionBillingManagerDescribeBillingJobsParams object
// with the ability to set a context for a request.
func NewSubscriptionBillingManagerDescribeBillingJobsParamsWithContext(ctx context.Context) *SubscriptionBillingManagerDescribeBillingJobsParams {
	return &SubscriptionBillingManagerDescribeBillingJobsParams{
		Context: ctx,
	}
}

// NewSubscriptionBillingManagerDescribeBillingJobsParamsWithHTTPClient creates a new SubscriptionBillingManagerDescribeBillingJobsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubscriptionBillingManagerDescribeBillingJobsParamsWithHTTPClient(client *http.Client) *SubscriptionBillingManagerDescribeBillingJobsParams {
	return &SubscriptionBillingManagerDescribeBillingJobsParams{
		HTTPClient: client,
	}
}

/*
SubscriptionBillingManagerDescribeBillingJobsParams contains all the parameters to send to the API endpoint

	for the subscription billing manager describe billing jobs operation.

	Typically these are written to a http.Request.
*/
type SubscriptionBillingManagerDescribeBillingJobsParams struct {

	/* AccessSysID.

	   接入系统ID.
	*/
	AccessSysID *string

	/* BllJobID.

	   计费任务ID.
	*/
	BllJobID []string

	/* CompID.

	   计费项ID.
	*/
	CompID []string

	/* EndTime.

	   计费任务结束时间.

	   Format: date-time
	*/
	EndTime *strfmt.DateTime

	/* Limit.

	   数据库查询分页大小- 默认 20，最大值 200.

	   Format: uint64
	*/
	Limit *string

	/* Offset.

	   数据库查询偏移位置- 默认 0.

	   Format: uint64
	*/
	Offset *string

	/* PlanID.

	   计费方案ID.
	*/
	PlanID []string

	/* ProdInstID.

	   产品实例ID.
	*/
	ProdInstID []string

	/* Reverse.

	   是否倒序排序- 0：ASC，1：DESC.
	*/
	Reverse *bool

	/* SchedulerJobID.

	   调度任务ID.
	*/
	SchedulerJobID []string

	/* SearchWord.

	   模糊查询关键字- 支持字段：无.
	*/
	SearchWord *string

	/* SortKey.

	   排序字段 - 默认 create_time.
	*/
	SortKey *string

	/* StartTime.

	   计费任务开始时间.

	   Format: date-time
	*/
	StartTime *strfmt.DateTime

	/* Status.

	   计费任务状态.
	*/
	Status []string

	/* SubsID.

	   订阅ID.
	*/
	SubsID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subscription billing manager describe billing jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithDefaults() *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subscription billing manager describe billing jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithTimeout(timeout time.Duration) *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithContext(ctx context.Context) *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithHTTPClient(client *http.Client) *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessSysID adds the accessSysID to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithAccessSysID(accessSysID *string) *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetAccessSysID(accessSysID)
	return o
}

// SetAccessSysID adds the accessSysId to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetAccessSysID(accessSysID *string) {
	o.AccessSysID = accessSysID
}

// WithBllJobID adds the bllJobID to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithBllJobID(bllJobID []string) *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetBllJobID(bllJobID)
	return o
}

// SetBllJobID adds the bllJobId to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetBllJobID(bllJobID []string) {
	o.BllJobID = bllJobID
}

// WithCompID adds the compID to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithCompID(compID []string) *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetCompID(compID)
	return o
}

// SetCompID adds the compId to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetCompID(compID []string) {
	o.CompID = compID
}

// WithEndTime adds the endTime to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithEndTime(endTime *strfmt.DateTime) *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetEndTime(endTime *strfmt.DateTime) {
	o.EndTime = endTime
}

// WithLimit adds the limit to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithLimit(limit *string) *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithOffset(offset *string) *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithPlanID adds the planID to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithPlanID(planID []string) *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetPlanID(planID)
	return o
}

// SetPlanID adds the planId to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetPlanID(planID []string) {
	o.PlanID = planID
}

// WithProdInstID adds the prodInstID to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithProdInstID(prodInstID []string) *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetProdInstID(prodInstID)
	return o
}

// SetProdInstID adds the prodInstId to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetProdInstID(prodInstID []string) {
	o.ProdInstID = prodInstID
}

// WithReverse adds the reverse to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithReverse(reverse *bool) *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSchedulerJobID adds the schedulerJobID to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithSchedulerJobID(schedulerJobID []string) *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetSchedulerJobID(schedulerJobID)
	return o
}

// SetSchedulerJobID adds the schedulerJobId to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetSchedulerJobID(schedulerJobID []string) {
	o.SchedulerJobID = schedulerJobID
}

// WithSearchWord adds the searchWord to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithSearchWord(searchWord *string) *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithSortKey(sortKey *string) *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStartTime adds the startTime to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithStartTime(startTime *strfmt.DateTime) *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetStartTime(startTime *strfmt.DateTime) {
	o.StartTime = startTime
}

// WithStatus adds the status to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithStatus(status []string) *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetStatus(status []string) {
	o.Status = status
}

// WithSubsID adds the subsID to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WithSubsID(subsID []string) *SubscriptionBillingManagerDescribeBillingJobsParams {
	o.SetSubsID(subsID)
	return o
}

// SetSubsID adds the subsId to the subscription billing manager describe billing jobs params
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) SetSubsID(subsID []string) {
	o.SubsID = subsID
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessSysID != nil {

		// query param access_sys_id
		var qrAccessSysID string

		if o.AccessSysID != nil {
			qrAccessSysID = *o.AccessSysID
		}
		qAccessSysID := qrAccessSysID
		if qAccessSysID != "" {

			if err := r.SetQueryParam("access_sys_id", qAccessSysID); err != nil {
				return err
			}
		}
	}

	if o.BllJobID != nil {

		// binding items for bll_job_id
		joinedBllJobID := o.bindParamBllJobID(reg)

		// query array param bll_job_id
		if err := r.SetQueryParam("bll_job_id", joinedBllJobID...); err != nil {
			return err
		}
	}

	if o.CompID != nil {

		// binding items for comp_id
		joinedCompID := o.bindParamCompID(reg)

		// query array param comp_id
		if err := r.SetQueryParam("comp_id", joinedCompID...); err != nil {
			return err
		}
	}

	if o.EndTime != nil {

		// query param end_time
		var qrEndTime strfmt.DateTime

		if o.EndTime != nil {
			qrEndTime = *o.EndTime
		}
		qEndTime := qrEndTime.String()
		if qEndTime != "" {

			if err := r.SetQueryParam("end_time", qEndTime); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit string

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.PlanID != nil {

		// binding items for plan_id
		joinedPlanID := o.bindParamPlanID(reg)

		// query array param plan_id
		if err := r.SetQueryParam("plan_id", joinedPlanID...); err != nil {
			return err
		}
	}

	if o.ProdInstID != nil {

		// binding items for prod_inst_id
		joinedProdInstID := o.bindParamProdInstID(reg)

		// query array param prod_inst_id
		if err := r.SetQueryParam("prod_inst_id", joinedProdInstID...); err != nil {
			return err
		}
	}

	if o.Reverse != nil {

		// query param reverse
		var qrReverse bool

		if o.Reverse != nil {
			qrReverse = *o.Reverse
		}
		qReverse := swag.FormatBool(qrReverse)
		if qReverse != "" {

			if err := r.SetQueryParam("reverse", qReverse); err != nil {
				return err
			}
		}
	}

	if o.SchedulerJobID != nil {

		// binding items for scheduler_job_id
		joinedSchedulerJobID := o.bindParamSchedulerJobID(reg)

		// query array param scheduler_job_id
		if err := r.SetQueryParam("scheduler_job_id", joinedSchedulerJobID...); err != nil {
			return err
		}
	}

	if o.SearchWord != nil {

		// query param search_word
		var qrSearchWord string

		if o.SearchWord != nil {
			qrSearchWord = *o.SearchWord
		}
		qSearchWord := qrSearchWord
		if qSearchWord != "" {

			if err := r.SetQueryParam("search_word", qSearchWord); err != nil {
				return err
			}
		}
	}

	if o.SortKey != nil {

		// query param sort_key
		var qrSortKey string

		if o.SortKey != nil {
			qrSortKey = *o.SortKey
		}
		qSortKey := qrSortKey
		if qSortKey != "" {

			if err := r.SetQueryParam("sort_key", qSortKey); err != nil {
				return err
			}
		}
	}

	if o.StartTime != nil {

		// query param start_time
		var qrStartTime strfmt.DateTime

		if o.StartTime != nil {
			qrStartTime = *o.StartTime
		}
		qStartTime := qrStartTime.String()
		if qStartTime != "" {

			if err := r.SetQueryParam("start_time", qStartTime); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// binding items for status
		joinedStatus := o.bindParamStatus(reg)

		// query array param status
		if err := r.SetQueryParam("status", joinedStatus...); err != nil {
			return err
		}
	}

	if o.SubsID != nil {

		// binding items for subs_id
		joinedSubsID := o.bindParamSubsID(reg)

		// query array param subs_id
		if err := r.SetQueryParam("subs_id", joinedSubsID...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSubscriptionBillingManagerDescribeBillingJobs binds the parameter bll_job_id
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) bindParamBllJobID(formats strfmt.Registry) []string {
	bllJobIDIR := o.BllJobID

	var bllJobIDIC []string
	for _, bllJobIDIIR := range bllJobIDIR { // explode []string

		bllJobIDIIV := bllJobIDIIR // string as string
		bllJobIDIC = append(bllJobIDIC, bllJobIDIIV)
	}

	// items.CollectionFormat: "multi"
	bllJobIDIS := swag.JoinByFormat(bllJobIDIC, "multi")

	return bllJobIDIS
}

// bindParamSubscriptionBillingManagerDescribeBillingJobs binds the parameter comp_id
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) bindParamCompID(formats strfmt.Registry) []string {
	compIDIR := o.CompID

	var compIDIC []string
	for _, compIDIIR := range compIDIR { // explode []string

		compIDIIV := compIDIIR // string as string
		compIDIC = append(compIDIC, compIDIIV)
	}

	// items.CollectionFormat: "multi"
	compIDIS := swag.JoinByFormat(compIDIC, "multi")

	return compIDIS
}

// bindParamSubscriptionBillingManagerDescribeBillingJobs binds the parameter plan_id
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) bindParamPlanID(formats strfmt.Registry) []string {
	planIDIR := o.PlanID

	var planIDIC []string
	for _, planIDIIR := range planIDIR { // explode []string

		planIDIIV := planIDIIR // string as string
		planIDIC = append(planIDIC, planIDIIV)
	}

	// items.CollectionFormat: "multi"
	planIDIS := swag.JoinByFormat(planIDIC, "multi")

	return planIDIS
}

// bindParamSubscriptionBillingManagerDescribeBillingJobs binds the parameter prod_inst_id
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) bindParamProdInstID(formats strfmt.Registry) []string {
	prodInstIDIR := o.ProdInstID

	var prodInstIDIC []string
	for _, prodInstIDIIR := range prodInstIDIR { // explode []string

		prodInstIDIIV := prodInstIDIIR // string as string
		prodInstIDIC = append(prodInstIDIC, prodInstIDIIV)
	}

	// items.CollectionFormat: "multi"
	prodInstIDIS := swag.JoinByFormat(prodInstIDIC, "multi")

	return prodInstIDIS
}

// bindParamSubscriptionBillingManagerDescribeBillingJobs binds the parameter scheduler_job_id
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) bindParamSchedulerJobID(formats strfmt.Registry) []string {
	schedulerJobIDIR := o.SchedulerJobID

	var schedulerJobIDIC []string
	for _, schedulerJobIDIIR := range schedulerJobIDIR { // explode []string

		schedulerJobIDIIV := schedulerJobIDIIR // string as string
		schedulerJobIDIC = append(schedulerJobIDIC, schedulerJobIDIIV)
	}

	// items.CollectionFormat: "multi"
	schedulerJobIDIS := swag.JoinByFormat(schedulerJobIDIC, "multi")

	return schedulerJobIDIS
}

// bindParamSubscriptionBillingManagerDescribeBillingJobs binds the parameter status
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) bindParamStatus(formats strfmt.Registry) []string {
	statusIR := o.Status

	var statusIC []string
	for _, statusIIR := range statusIR { // explode []string

		statusIIV := statusIIR // string as string
		statusIC = append(statusIC, statusIIV)
	}

	// items.CollectionFormat: "multi"
	statusIS := swag.JoinByFormat(statusIC, "multi")

	return statusIS
}

// bindParamSubscriptionBillingManagerDescribeBillingJobs binds the parameter subs_id
func (o *SubscriptionBillingManagerDescribeBillingJobsParams) bindParamSubsID(formats strfmt.Registry) []string {
	subsIDIR := o.SubsID

	var subsIDIC []string
	for _, subsIDIIR := range subsIDIR { // explode []string

		subsIDIIV := subsIDIIR // string as string
		subsIDIC = append(subsIDIC, subsIDIIV)
	}

	// items.CollectionFormat: "multi"
	subsIDIS := swag.JoinByFormat(subsIDIC, "multi")

	return subsIDIS
}
