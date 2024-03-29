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

// NewPricingManagerDescribePlansParams creates a new PricingManagerDescribePlansParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPricingManagerDescribePlansParams() *PricingManagerDescribePlansParams {
	return &PricingManagerDescribePlansParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPricingManagerDescribePlansParamsWithTimeout creates a new PricingManagerDescribePlansParams object
// with the ability to set a timeout on a request.
func NewPricingManagerDescribePlansParamsWithTimeout(timeout time.Duration) *PricingManagerDescribePlansParams {
	return &PricingManagerDescribePlansParams{
		timeout: timeout,
	}
}

// NewPricingManagerDescribePlansParamsWithContext creates a new PricingManagerDescribePlansParams object
// with the ability to set a context for a request.
func NewPricingManagerDescribePlansParamsWithContext(ctx context.Context) *PricingManagerDescribePlansParams {
	return &PricingManagerDescribePlansParams{
		Context: ctx,
	}
}

// NewPricingManagerDescribePlansParamsWithHTTPClient creates a new PricingManagerDescribePlansParams object
// with the ability to set a custom HTTPClient for a request.
func NewPricingManagerDescribePlansParamsWithHTTPClient(client *http.Client) *PricingManagerDescribePlansParams {
	return &PricingManagerDescribePlansParams{
		HTTPClient: client,
	}
}

/*
PricingManagerDescribePlansParams contains all the parameters to send to the API endpoint

	for the pricing manager describe plans operation.

	Typically these are written to a http.Request.
*/
type PricingManagerDescribePlansParams struct {

	/* Limit.

	   数据库查询每页大小 - 默认 20, 最大值 200.

	   Format: uint64
	*/
	Limit *string

	/* Offset.

	   数据库查询偏移位置 - 默认 0.

	   Format: uint64
	*/
	Offset *string

	/* PlanCode.

	   方案编码.
	*/
	PlanCode []string

	/* PlanID.

	   方案ID.
	*/
	PlanID []string

	/* ProdID.

	   产品ID.
	*/
	ProdID []string

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

	/* Status.

	   状态editing/published todo.
	*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pricing manager describe plans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerDescribePlansParams) WithDefaults() *PricingManagerDescribePlansParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pricing manager describe plans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerDescribePlansParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) WithTimeout(timeout time.Duration) *PricingManagerDescribePlansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) WithContext(ctx context.Context) *PricingManagerDescribePlansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) WithHTTPClient(client *http.Client) *PricingManagerDescribePlansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) WithLimit(limit *string) *PricingManagerDescribePlansParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) WithOffset(offset *string) *PricingManagerDescribePlansParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithPlanCode adds the planCode to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) WithPlanCode(planCode []string) *PricingManagerDescribePlansParams {
	o.SetPlanCode(planCode)
	return o
}

// SetPlanCode adds the planCode to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) SetPlanCode(planCode []string) {
	o.PlanCode = planCode
}

// WithPlanID adds the planID to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) WithPlanID(planID []string) *PricingManagerDescribePlansParams {
	o.SetPlanID(planID)
	return o
}

// SetPlanID adds the planId to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) SetPlanID(planID []string) {
	o.PlanID = planID
}

// WithProdID adds the prodID to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) WithProdID(prodID []string) *PricingManagerDescribePlansParams {
	o.SetProdID(prodID)
	return o
}

// SetProdID adds the prodId to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) SetProdID(prodID []string) {
	o.ProdID = prodID
}

// WithReverse adds the reverse to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) WithReverse(reverse *bool) *PricingManagerDescribePlansParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSearchWord adds the searchWord to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) WithSearchWord(searchWord *string) *PricingManagerDescribePlansParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) WithSortKey(sortKey *string) *PricingManagerDescribePlansParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStatus adds the status to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) WithStatus(status []string) *PricingManagerDescribePlansParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the pricing manager describe plans params
func (o *PricingManagerDescribePlansParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *PricingManagerDescribePlansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.PlanCode != nil {

		// binding items for plan_code
		joinedPlanCode := o.bindParamPlanCode(reg)

		// query array param plan_code
		if err := r.SetQueryParam("plan_code", joinedPlanCode...); err != nil {
			return err
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

	if o.ProdID != nil {

		// binding items for prod_id
		joinedProdID := o.bindParamProdID(reg)

		// query array param prod_id
		if err := r.SetQueryParam("prod_id", joinedProdID...); err != nil {
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

// bindParamPricingManagerDescribePlans binds the parameter plan_code
func (o *PricingManagerDescribePlansParams) bindParamPlanCode(formats strfmt.Registry) []string {
	planCodeIR := o.PlanCode

	var planCodeIC []string
	for _, planCodeIIR := range planCodeIR { // explode []string

		planCodeIIV := planCodeIIR // string as string
		planCodeIC = append(planCodeIC, planCodeIIV)
	}

	// items.CollectionFormat: "multi"
	planCodeIS := swag.JoinByFormat(planCodeIC, "multi")

	return planCodeIS
}

// bindParamPricingManagerDescribePlans binds the parameter plan_id
func (o *PricingManagerDescribePlansParams) bindParamPlanID(formats strfmt.Registry) []string {
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

// bindParamPricingManagerDescribePlans binds the parameter prod_id
func (o *PricingManagerDescribePlansParams) bindParamProdID(formats strfmt.Registry) []string {
	prodIDIR := o.ProdID

	var prodIDIC []string
	for _, prodIDIIR := range prodIDIR { // explode []string

		prodIDIIV := prodIDIIR // string as string
		prodIDIC = append(prodIDIC, prodIDIIV)
	}

	// items.CollectionFormat: "multi"
	prodIDIS := swag.JoinByFormat(prodIDIC, "multi")

	return prodIDIS
}

// bindParamPricingManagerDescribePlans binds the parameter status
func (o *PricingManagerDescribePlansParams) bindParamStatus(formats strfmt.Registry) []string {
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
