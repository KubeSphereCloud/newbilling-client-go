// Code generated by go-swagger; DO NOT EDIT.

package pricing_manager

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

// NewPricingManagerDescribeComponentsParams creates a new PricingManagerDescribeComponentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPricingManagerDescribeComponentsParams() *PricingManagerDescribeComponentsParams {
	return &PricingManagerDescribeComponentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPricingManagerDescribeComponentsParamsWithTimeout creates a new PricingManagerDescribeComponentsParams object
// with the ability to set a timeout on a request.
func NewPricingManagerDescribeComponentsParamsWithTimeout(timeout time.Duration) *PricingManagerDescribeComponentsParams {
	return &PricingManagerDescribeComponentsParams{
		timeout: timeout,
	}
}

// NewPricingManagerDescribeComponentsParamsWithContext creates a new PricingManagerDescribeComponentsParams object
// with the ability to set a context for a request.
func NewPricingManagerDescribeComponentsParamsWithContext(ctx context.Context) *PricingManagerDescribeComponentsParams {
	return &PricingManagerDescribeComponentsParams{
		Context: ctx,
	}
}

// NewPricingManagerDescribeComponentsParamsWithHTTPClient creates a new PricingManagerDescribeComponentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPricingManagerDescribeComponentsParamsWithHTTPClient(client *http.Client) *PricingManagerDescribeComponentsParams {
	return &PricingManagerDescribeComponentsParams{
		HTTPClient: client,
	}
}

/*
PricingManagerDescribeComponentsParams contains all the parameters to send to the API endpoint

	for the pricing manager describe components operation.

	Typically these are written to a http.Request.
*/
type PricingManagerDescribeComponentsParams struct {

	/* BillingMode.

	   计费模式 - 按包-按资源包；按包-按时间包；按量-按时间使用量；按量-按资源使用量.
	*/
	BillingMode *string

	/* CompCode.

	   comp_code.
	*/
	CompCode *string

	/* CompID.

	   计费项ID.
	*/
	CompID *string

	/* Limit.

	   数据库查询每页大小 - 默认 20, 最大值 200.

	   Format: uint64
	*/
	Limit *string

	/* Name.

	   计费项名称.
	*/
	Name *string

	/* Offset.

	   数据库查询偏移位置 - 默认 0.

	   Format: uint64
	*/
	Offset *string

	/* PlanID.

	   方案ID.
	*/
	PlanID *string

	/* ProdID.

	   产品ID 支持产品code.
	*/
	ProdID *string

	/* Reverse.

	   是否倒序排序 - value = 0 sort ASC, value = 1 sort DESC.
	*/
	Reverse *bool

	/* SearchWord.

	   模糊查询关键字.
	*/
	SearchWord *string

	/* SortKey.

	   排序字段 - 默认 create_time.
	*/
	SortKey *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pricing manager describe components params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerDescribeComponentsParams) WithDefaults() *PricingManagerDescribeComponentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pricing manager describe components params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerDescribeComponentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) WithTimeout(timeout time.Duration) *PricingManagerDescribeComponentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) WithContext(ctx context.Context) *PricingManagerDescribeComponentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) WithHTTPClient(client *http.Client) *PricingManagerDescribeComponentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillingMode adds the billingMode to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) WithBillingMode(billingMode *string) *PricingManagerDescribeComponentsParams {
	o.SetBillingMode(billingMode)
	return o
}

// SetBillingMode adds the billingMode to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) SetBillingMode(billingMode *string) {
	o.BillingMode = billingMode
}

// WithCompCode adds the compCode to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) WithCompCode(compCode *string) *PricingManagerDescribeComponentsParams {
	o.SetCompCode(compCode)
	return o
}

// SetCompCode adds the compCode to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) SetCompCode(compCode *string) {
	o.CompCode = compCode
}

// WithCompID adds the compID to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) WithCompID(compID *string) *PricingManagerDescribeComponentsParams {
	o.SetCompID(compID)
	return o
}

// SetCompID adds the compId to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) SetCompID(compID *string) {
	o.CompID = compID
}

// WithLimit adds the limit to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) WithLimit(limit *string) *PricingManagerDescribeComponentsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithName adds the name to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) WithName(name *string) *PricingManagerDescribeComponentsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) WithOffset(offset *string) *PricingManagerDescribeComponentsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithPlanID adds the planID to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) WithPlanID(planID *string) *PricingManagerDescribeComponentsParams {
	o.SetPlanID(planID)
	return o
}

// SetPlanID adds the planId to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) SetPlanID(planID *string) {
	o.PlanID = planID
}

// WithProdID adds the prodID to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) WithProdID(prodID *string) *PricingManagerDescribeComponentsParams {
	o.SetProdID(prodID)
	return o
}

// SetProdID adds the prodId to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) SetProdID(prodID *string) {
	o.ProdID = prodID
}

// WithReverse adds the reverse to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) WithReverse(reverse *bool) *PricingManagerDescribeComponentsParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSearchWord adds the searchWord to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) WithSearchWord(searchWord *string) *PricingManagerDescribeComponentsParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) WithSortKey(sortKey *string) *PricingManagerDescribeComponentsParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the pricing manager describe components params
func (o *PricingManagerDescribeComponentsParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WriteToRequest writes these params to a swagger request
func (o *PricingManagerDescribeComponentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BillingMode != nil {

		// query param billing_mode
		var qrBillingMode string

		if o.BillingMode != nil {
			qrBillingMode = *o.BillingMode
		}
		qBillingMode := qrBillingMode
		if qBillingMode != "" {

			if err := r.SetQueryParam("billing_mode", qBillingMode); err != nil {
				return err
			}
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

		// query param comp_id
		var qrCompID string

		if o.CompID != nil {
			qrCompID = *o.CompID
		}
		qCompID := qrCompID
		if qCompID != "" {

			if err := r.SetQueryParam("comp_id", qCompID); err != nil {
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

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
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

		// query param plan_id
		var qrPlanID string

		if o.PlanID != nil {
			qrPlanID = *o.PlanID
		}
		qPlanID := qrPlanID
		if qPlanID != "" {

			if err := r.SetQueryParam("plan_id", qPlanID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
