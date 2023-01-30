// Code generated by go-swagger; DO NOT EDIT.

package discount_manager

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

// NewDiscountManagerDescribeCustomerDiscountsParams creates a new DiscountManagerDescribeCustomerDiscountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDiscountManagerDescribeCustomerDiscountsParams() *DiscountManagerDescribeCustomerDiscountsParams {
	return &DiscountManagerDescribeCustomerDiscountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDiscountManagerDescribeCustomerDiscountsParamsWithTimeout creates a new DiscountManagerDescribeCustomerDiscountsParams object
// with the ability to set a timeout on a request.
func NewDiscountManagerDescribeCustomerDiscountsParamsWithTimeout(timeout time.Duration) *DiscountManagerDescribeCustomerDiscountsParams {
	return &DiscountManagerDescribeCustomerDiscountsParams{
		timeout: timeout,
	}
}

// NewDiscountManagerDescribeCustomerDiscountsParamsWithContext creates a new DiscountManagerDescribeCustomerDiscountsParams object
// with the ability to set a context for a request.
func NewDiscountManagerDescribeCustomerDiscountsParamsWithContext(ctx context.Context) *DiscountManagerDescribeCustomerDiscountsParams {
	return &DiscountManagerDescribeCustomerDiscountsParams{
		Context: ctx,
	}
}

// NewDiscountManagerDescribeCustomerDiscountsParamsWithHTTPClient creates a new DiscountManagerDescribeCustomerDiscountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDiscountManagerDescribeCustomerDiscountsParamsWithHTTPClient(client *http.Client) *DiscountManagerDescribeCustomerDiscountsParams {
	return &DiscountManagerDescribeCustomerDiscountsParams{
		HTTPClient: client,
	}
}

/*
DiscountManagerDescribeCustomerDiscountsParams contains all the parameters to send to the API endpoint

	for the discount manager describe customer discounts operation.

	Typically these are written to a http.Request.
*/
type DiscountManagerDescribeCustomerDiscountsParams struct {

	/* AccessSysID.

	   接入系统ID.
	*/
	AccessSysID []string

	/* CustomerDiscountName.

	   产品折扣名称.
	*/
	CustomerDiscountName []string

	/* CustomerID.

	   客户ID.
	*/
	CustomerID []string

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

	/* Reverse.

	   是否倒序排序 - value = 0 sort ASC, value = 1 sort DESC.
	*/
	Reverse *bool

	/* SearchWord.

	   模糊查询关键字 - 支持字段：toadd.
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

// WithDefaults hydrates default values in the discount manager describe customer discounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiscountManagerDescribeCustomerDiscountsParams) WithDefaults() *DiscountManagerDescribeCustomerDiscountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the discount manager describe customer discounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiscountManagerDescribeCustomerDiscountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) WithTimeout(timeout time.Duration) *DiscountManagerDescribeCustomerDiscountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) WithContext(ctx context.Context) *DiscountManagerDescribeCustomerDiscountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) WithHTTPClient(client *http.Client) *DiscountManagerDescribeCustomerDiscountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessSysID adds the accessSysID to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) WithAccessSysID(accessSysID []string) *DiscountManagerDescribeCustomerDiscountsParams {
	o.SetAccessSysID(accessSysID)
	return o
}

// SetAccessSysID adds the accessSysId to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) SetAccessSysID(accessSysID []string) {
	o.AccessSysID = accessSysID
}

// WithCustomerDiscountName adds the customerDiscountName to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) WithCustomerDiscountName(customerDiscountName []string) *DiscountManagerDescribeCustomerDiscountsParams {
	o.SetCustomerDiscountName(customerDiscountName)
	return o
}

// SetCustomerDiscountName adds the customerDiscountName to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) SetCustomerDiscountName(customerDiscountName []string) {
	o.CustomerDiscountName = customerDiscountName
}

// WithCustomerID adds the customerID to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) WithCustomerID(customerID []string) *DiscountManagerDescribeCustomerDiscountsParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) SetCustomerID(customerID []string) {
	o.CustomerID = customerID
}

// WithLimit adds the limit to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) WithLimit(limit *string) *DiscountManagerDescribeCustomerDiscountsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) WithOffset(offset *string) *DiscountManagerDescribeCustomerDiscountsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithReverse adds the reverse to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) WithReverse(reverse *bool) *DiscountManagerDescribeCustomerDiscountsParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSearchWord adds the searchWord to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) WithSearchWord(searchWord *string) *DiscountManagerDescribeCustomerDiscountsParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) WithSortKey(sortKey *string) *DiscountManagerDescribeCustomerDiscountsParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the discount manager describe customer discounts params
func (o *DiscountManagerDescribeCustomerDiscountsParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WriteToRequest writes these params to a swagger request
func (o *DiscountManagerDescribeCustomerDiscountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.CustomerDiscountName != nil {

		// binding items for customer_discount_name
		joinedCustomerDiscountName := o.bindParamCustomerDiscountName(reg)

		// query array param customer_discount_name
		if err := r.SetQueryParam("customer_discount_name", joinedCustomerDiscountName...); err != nil {
			return err
		}
	}

	if o.CustomerID != nil {

		// binding items for customer_id
		joinedCustomerID := o.bindParamCustomerID(reg)

		// query array param customer_id
		if err := r.SetQueryParam("customer_id", joinedCustomerID...); err != nil {
			return err
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

// bindParamDiscountManagerDescribeCustomerDiscounts binds the parameter access_sys_id
func (o *DiscountManagerDescribeCustomerDiscountsParams) bindParamAccessSysID(formats strfmt.Registry) []string {
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

// bindParamDiscountManagerDescribeCustomerDiscounts binds the parameter customer_discount_name
func (o *DiscountManagerDescribeCustomerDiscountsParams) bindParamCustomerDiscountName(formats strfmt.Registry) []string {
	customerDiscountNameIR := o.CustomerDiscountName

	var customerDiscountNameIC []string
	for _, customerDiscountNameIIR := range customerDiscountNameIR { // explode []string

		customerDiscountNameIIV := customerDiscountNameIIR // string as string
		customerDiscountNameIC = append(customerDiscountNameIC, customerDiscountNameIIV)
	}

	// items.CollectionFormat: "multi"
	customerDiscountNameIS := swag.JoinByFormat(customerDiscountNameIC, "multi")

	return customerDiscountNameIS
}

// bindParamDiscountManagerDescribeCustomerDiscounts binds the parameter customer_id
func (o *DiscountManagerDescribeCustomerDiscountsParams) bindParamCustomerID(formats strfmt.Registry) []string {
	customerIDIR := o.CustomerID

	var customerIDIC []string
	for _, customerIDIIR := range customerIDIR { // explode []string

		customerIDIIV := customerIDIIR // string as string
		customerIDIC = append(customerIDIC, customerIDIIV)
	}

	// items.CollectionFormat: "multi"
	customerIDIS := swag.JoinByFormat(customerIDIC, "multi")

	return customerIDIS
}
