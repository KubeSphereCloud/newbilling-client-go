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

// NewSubscriptionBillingManagerDescribeBillsParams creates a new SubscriptionBillingManagerDescribeBillsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubscriptionBillingManagerDescribeBillsParams() *SubscriptionBillingManagerDescribeBillsParams {
	return &SubscriptionBillingManagerDescribeBillsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriptionBillingManagerDescribeBillsParamsWithTimeout creates a new SubscriptionBillingManagerDescribeBillsParams object
// with the ability to set a timeout on a request.
func NewSubscriptionBillingManagerDescribeBillsParamsWithTimeout(timeout time.Duration) *SubscriptionBillingManagerDescribeBillsParams {
	return &SubscriptionBillingManagerDescribeBillsParams{
		timeout: timeout,
	}
}

// NewSubscriptionBillingManagerDescribeBillsParamsWithContext creates a new SubscriptionBillingManagerDescribeBillsParams object
// with the ability to set a context for a request.
func NewSubscriptionBillingManagerDescribeBillsParamsWithContext(ctx context.Context) *SubscriptionBillingManagerDescribeBillsParams {
	return &SubscriptionBillingManagerDescribeBillsParams{
		Context: ctx,
	}
}

// NewSubscriptionBillingManagerDescribeBillsParamsWithHTTPClient creates a new SubscriptionBillingManagerDescribeBillsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubscriptionBillingManagerDescribeBillsParamsWithHTTPClient(client *http.Client) *SubscriptionBillingManagerDescribeBillsParams {
	return &SubscriptionBillingManagerDescribeBillsParams{
		HTTPClient: client,
	}
}

/*
SubscriptionBillingManagerDescribeBillsParams contains all the parameters to send to the API endpoint

	for the subscription billing manager describe bills operation.

	Typically these are written to a http.Request.
*/
type SubscriptionBillingManagerDescribeBillsParams struct {

	/* AccessSysID.

	   接入系统ID.
	*/
	AccessSysID []string

	/* AccountPeriod.

	   账期.
	*/
	AccountPeriod []string

	/* BillID.

	   计费账单ID.
	*/
	BillID []string

	/* BillingMode.

	   计费模式.
	*/
	BillingMode []string

	/* BllJobID.

	   计费任务ID.
	*/
	BllJobID []string

	// CompCode.
	CompCode *string

	/* CompID.

	   计费项ID.
	*/
	CompID []string

	/* CompName.

	   计费项名称.
	*/
	CompName *string

	/* EndTime.

	   计费结束时间.

	   Format: date-time
	*/
	EndTime *strfmt.DateTime

	// ExcludeZeroBill.
	ExcludeZeroBill *bool

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

	/* OrderUser.

	   订单用户ID或名称- 模糊查询.
	*/
	OrderUser *string

	/* Prod.

	   产品ID或名称- 模糊查询.
	*/
	Prod *string

	/* ProdCode.

	   产品ode.
	*/
	ProdCode *string

	/* ProdID.

	   产品ID.
	*/
	ProdID *string

	/* ProdInstID.

	   产品实例ID.
	*/
	ProdInstID *string

	/* ProdInstIDExt.

	   外部产品实例ID.
	*/
	ProdInstIDExt []string

	/* ProdName.

	   产品名称.
	*/
	ProdName *string

	/* Reverse.

	   是否倒序排序- 0：ASC，1：DESC.
	*/
	Reverse *bool

	/* SearchWord.

	   模糊查询关键字- 支持字段：order_user_name、prod_name、comp_name.
	*/
	SearchWord *string

	/* SortKey.

	   排序字段- 默认 create_time.
	*/
	SortKey *string

	/* StartTime.

	   计费开始时间.

	   Format: date-time
	*/
	StartTime *strfmt.DateTime

	/* Status.

	   状态.
	*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subscription billing manager describe bills params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionBillingManagerDescribeBillsParams) WithDefaults() *SubscriptionBillingManagerDescribeBillsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subscription billing manager describe bills params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionBillingManagerDescribeBillsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithTimeout(timeout time.Duration) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithContext(ctx context.Context) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithHTTPClient(client *http.Client) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessSysID adds the accessSysID to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithAccessSysID(accessSysID []string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetAccessSysID(accessSysID)
	return o
}

// SetAccessSysID adds the accessSysId to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetAccessSysID(accessSysID []string) {
	o.AccessSysID = accessSysID
}

// WithAccountPeriod adds the accountPeriod to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithAccountPeriod(accountPeriod []string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetAccountPeriod(accountPeriod)
	return o
}

// SetAccountPeriod adds the accountPeriod to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetAccountPeriod(accountPeriod []string) {
	o.AccountPeriod = accountPeriod
}

// WithBillID adds the billID to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithBillID(billID []string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetBillID(billID)
	return o
}

// SetBillID adds the billId to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetBillID(billID []string) {
	o.BillID = billID
}

// WithBillingMode adds the billingMode to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithBillingMode(billingMode []string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetBillingMode(billingMode)
	return o
}

// SetBillingMode adds the billingMode to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetBillingMode(billingMode []string) {
	o.BillingMode = billingMode
}

// WithBllJobID adds the bllJobID to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithBllJobID(bllJobID []string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetBllJobID(bllJobID)
	return o
}

// SetBllJobID adds the bllJobId to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetBllJobID(bllJobID []string) {
	o.BllJobID = bllJobID
}

// WithCompCode adds the compCode to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithCompCode(compCode *string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetCompCode(compCode)
	return o
}

// SetCompCode adds the compCode to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetCompCode(compCode *string) {
	o.CompCode = compCode
}

// WithCompID adds the compID to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithCompID(compID []string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetCompID(compID)
	return o
}

// SetCompID adds the compId to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetCompID(compID []string) {
	o.CompID = compID
}

// WithCompName adds the compName to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithCompName(compName *string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetCompName(compName)
	return o
}

// SetCompName adds the compName to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetCompName(compName *string) {
	o.CompName = compName
}

// WithEndTime adds the endTime to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithEndTime(endTime *strfmt.DateTime) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetEndTime(endTime *strfmt.DateTime) {
	o.EndTime = endTime
}

// WithExcludeZeroBill adds the excludeZeroBill to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithExcludeZeroBill(excludeZeroBill *bool) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetExcludeZeroBill(excludeZeroBill)
	return o
}

// SetExcludeZeroBill adds the excludeZeroBill to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetExcludeZeroBill(excludeZeroBill *bool) {
	o.ExcludeZeroBill = excludeZeroBill
}

// WithLimit adds the limit to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithLimit(limit *string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithOffset(offset *string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithOrderUser adds the orderUser to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithOrderUser(orderUser *string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetOrderUser(orderUser)
	return o
}

// SetOrderUser adds the orderUser to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetOrderUser(orderUser *string) {
	o.OrderUser = orderUser
}

// WithProd adds the prod to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithProd(prod *string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetProd(prod)
	return o
}

// SetProd adds the prod to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetProd(prod *string) {
	o.Prod = prod
}

// WithProdCode adds the prodCode to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithProdCode(prodCode *string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetProdCode(prodCode)
	return o
}

// SetProdCode adds the prodCode to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetProdCode(prodCode *string) {
	o.ProdCode = prodCode
}

// WithProdID adds the prodID to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithProdID(prodID *string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetProdID(prodID)
	return o
}

// SetProdID adds the prodId to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetProdID(prodID *string) {
	o.ProdID = prodID
}

// WithProdInstID adds the prodInstID to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithProdInstID(prodInstID *string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetProdInstID(prodInstID)
	return o
}

// SetProdInstID adds the prodInstId to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetProdInstID(prodInstID *string) {
	o.ProdInstID = prodInstID
}

// WithProdInstIDExt adds the prodInstIDExt to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithProdInstIDExt(prodInstIDExt []string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetProdInstIDExt(prodInstIDExt)
	return o
}

// SetProdInstIDExt adds the prodInstIdExt to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetProdInstIDExt(prodInstIDExt []string) {
	o.ProdInstIDExt = prodInstIDExt
}

// WithProdName adds the prodName to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithProdName(prodName *string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetProdName(prodName)
	return o
}

// SetProdName adds the prodName to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetProdName(prodName *string) {
	o.ProdName = prodName
}

// WithReverse adds the reverse to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithReverse(reverse *bool) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSearchWord adds the searchWord to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithSearchWord(searchWord *string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithSortKey(sortKey *string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStartTime adds the startTime to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithStartTime(startTime *strfmt.DateTime) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetStartTime(startTime *strfmt.DateTime) {
	o.StartTime = startTime
}

// WithStatus adds the status to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) WithStatus(status []string) *SubscriptionBillingManagerDescribeBillsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the subscription billing manager describe bills params
func (o *SubscriptionBillingManagerDescribeBillsParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriptionBillingManagerDescribeBillsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessSysID != nil {

		// binding items for access_sys_id
		joinedAccessSysID := o.bindParamAccessSysID(reg)

		// query array param access_sys_id
		if err := r.SetQueryParam("access_sys_id", joinedAccessSysID...); err != nil {
			return err
		}
	}

	if o.AccountPeriod != nil {

		// binding items for account_period
		joinedAccountPeriod := o.bindParamAccountPeriod(reg)

		// query array param account_period
		if err := r.SetQueryParam("account_period", joinedAccountPeriod...); err != nil {
			return err
		}
	}

	if o.BillID != nil {

		// binding items for bill_id
		joinedBillID := o.bindParamBillID(reg)

		// query array param bill_id
		if err := r.SetQueryParam("bill_id", joinedBillID...); err != nil {
			return err
		}
	}

	if o.BillingMode != nil {

		// binding items for billing_mode
		joinedBillingMode := o.bindParamBillingMode(reg)

		// query array param billing_mode
		if err := r.SetQueryParam("billing_mode", joinedBillingMode...); err != nil {
			return err
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

	if o.CompCode != nil {

		// query param comp_code
		var qrCompCode string

		if o.CompCode != nil {
			qrCompCode = *o.CompCode
		}
		qCompCode := qrCompCode
		if qCompCode != "" {

			if err := r.SetQueryParam("comp_code", qCompCode); err != nil {
				return err
			}
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

	if o.CompName != nil {

		// query param comp_name
		var qrCompName string

		if o.CompName != nil {
			qrCompName = *o.CompName
		}
		qCompName := qrCompName
		if qCompName != "" {

			if err := r.SetQueryParam("comp_name", qCompName); err != nil {
				return err
			}
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

	if o.ExcludeZeroBill != nil {

		// query param exclude_zero_bill
		var qrExcludeZeroBill bool

		if o.ExcludeZeroBill != nil {
			qrExcludeZeroBill = *o.ExcludeZeroBill
		}
		qExcludeZeroBill := swag.FormatBool(qrExcludeZeroBill)
		if qExcludeZeroBill != "" {

			if err := r.SetQueryParam("exclude_zero_bill", qExcludeZeroBill); err != nil {
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

	if o.OrderUser != nil {

		// query param order_user
		var qrOrderUser string

		if o.OrderUser != nil {
			qrOrderUser = *o.OrderUser
		}
		qOrderUser := qrOrderUser
		if qOrderUser != "" {

			if err := r.SetQueryParam("order_user", qOrderUser); err != nil {
				return err
			}
		}
	}

	if o.Prod != nil {

		// query param prod
		var qrProd string

		if o.Prod != nil {
			qrProd = *o.Prod
		}
		qProd := qrProd
		if qProd != "" {

			if err := r.SetQueryParam("prod", qProd); err != nil {
				return err
			}
		}
	}

	if o.ProdCode != nil {

		// query param prod_code
		var qrProdCode string

		if o.ProdCode != nil {
			qrProdCode = *o.ProdCode
		}
		qProdCode := qrProdCode
		if qProdCode != "" {

			if err := r.SetQueryParam("prod_code", qProdCode); err != nil {
				return err
			}
		}
	}

	if o.ProdID != nil {

		// query param prod_id
		var qrProdID string

		if o.ProdID != nil {
			qrProdID = *o.ProdID
		}
		qProdID := qrProdID
		if qProdID != "" {

			if err := r.SetQueryParam("prod_id", qProdID); err != nil {
				return err
			}
		}
	}

	if o.ProdInstID != nil {

		// query param prod_inst_id
		var qrProdInstID string

		if o.ProdInstID != nil {
			qrProdInstID = *o.ProdInstID
		}
		qProdInstID := qrProdInstID
		if qProdInstID != "" {

			if err := r.SetQueryParam("prod_inst_id", qProdInstID); err != nil {
				return err
			}
		}
	}

	if o.ProdInstIDExt != nil {

		// binding items for prod_inst_id_ext
		joinedProdInstIDExt := o.bindParamProdInstIDExt(reg)

		// query array param prod_inst_id_ext
		if err := r.SetQueryParam("prod_inst_id_ext", joinedProdInstIDExt...); err != nil {
			return err
		}
	}

	if o.ProdName != nil {

		// query param prod_name
		var qrProdName string

		if o.ProdName != nil {
			qrProdName = *o.ProdName
		}
		qProdName := qrProdName
		if qProdName != "" {

			if err := r.SetQueryParam("prod_name", qProdName); err != nil {
				return err
			}
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSubscriptionBillingManagerDescribeBills binds the parameter access_sys_id
func (o *SubscriptionBillingManagerDescribeBillsParams) bindParamAccessSysID(formats strfmt.Registry) []string {
	accessSysIDIR := o.AccessSysID

	var accessSysIDIC []string
	for _, accessSysIDIIR := range accessSysIDIR { // explode []string

		accessSysIDIIV := accessSysIDIIR // string as string
		accessSysIDIC = append(accessSysIDIC, accessSysIDIIV)
	}

	// items.CollectionFormat: "multi"
	accessSysIDIS := swag.JoinByFormat(accessSysIDIC, "multi")

	return accessSysIDIS
}

// bindParamSubscriptionBillingManagerDescribeBills binds the parameter account_period
func (o *SubscriptionBillingManagerDescribeBillsParams) bindParamAccountPeriod(formats strfmt.Registry) []string {
	accountPeriodIR := o.AccountPeriod

	var accountPeriodIC []string
	for _, accountPeriodIIR := range accountPeriodIR { // explode []string

		accountPeriodIIV := accountPeriodIIR // string as string
		accountPeriodIC = append(accountPeriodIC, accountPeriodIIV)
	}

	// items.CollectionFormat: "multi"
	accountPeriodIS := swag.JoinByFormat(accountPeriodIC, "multi")

	return accountPeriodIS
}

// bindParamSubscriptionBillingManagerDescribeBills binds the parameter bill_id
func (o *SubscriptionBillingManagerDescribeBillsParams) bindParamBillID(formats strfmt.Registry) []string {
	billIDIR := o.BillID

	var billIDIC []string
	for _, billIDIIR := range billIDIR { // explode []string

		billIDIIV := billIDIIR // string as string
		billIDIC = append(billIDIC, billIDIIV)
	}

	// items.CollectionFormat: "multi"
	billIDIS := swag.JoinByFormat(billIDIC, "multi")

	return billIDIS
}

// bindParamSubscriptionBillingManagerDescribeBills binds the parameter billing_mode
func (o *SubscriptionBillingManagerDescribeBillsParams) bindParamBillingMode(formats strfmt.Registry) []string {
	billingModeIR := o.BillingMode

	var billingModeIC []string
	for _, billingModeIIR := range billingModeIR { // explode []string

		billingModeIIV := billingModeIIR // string as string
		billingModeIC = append(billingModeIC, billingModeIIV)
	}

	// items.CollectionFormat: "multi"
	billingModeIS := swag.JoinByFormat(billingModeIC, "multi")

	return billingModeIS
}

// bindParamSubscriptionBillingManagerDescribeBills binds the parameter bll_job_id
func (o *SubscriptionBillingManagerDescribeBillsParams) bindParamBllJobID(formats strfmt.Registry) []string {
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

// bindParamSubscriptionBillingManagerDescribeBills binds the parameter comp_id
func (o *SubscriptionBillingManagerDescribeBillsParams) bindParamCompID(formats strfmt.Registry) []string {
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

// bindParamSubscriptionBillingManagerDescribeBills binds the parameter prod_inst_id_ext
func (o *SubscriptionBillingManagerDescribeBillsParams) bindParamProdInstIDExt(formats strfmt.Registry) []string {
	prodInstIDExtIR := o.ProdInstIDExt

	var prodInstIDExtIC []string
	for _, prodInstIDExtIIR := range prodInstIDExtIR { // explode []string

		prodInstIDExtIIV := prodInstIDExtIIR // string as string
		prodInstIDExtIC = append(prodInstIDExtIC, prodInstIDExtIIV)
	}

	// items.CollectionFormat: "multi"
	prodInstIDExtIS := swag.JoinByFormat(prodInstIDExtIC, "multi")

	return prodInstIDExtIS
}

// bindParamSubscriptionBillingManagerDescribeBills binds the parameter status
func (o *SubscriptionBillingManagerDescribeBillsParams) bindParamStatus(formats strfmt.Registry) []string {
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
