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

// NewPricingManagerDescribeAttributesParams creates a new PricingManagerDescribeAttributesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPricingManagerDescribeAttributesParams() *PricingManagerDescribeAttributesParams {
	return &PricingManagerDescribeAttributesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPricingManagerDescribeAttributesParamsWithTimeout creates a new PricingManagerDescribeAttributesParams object
// with the ability to set a timeout on a request.
func NewPricingManagerDescribeAttributesParamsWithTimeout(timeout time.Duration) *PricingManagerDescribeAttributesParams {
	return &PricingManagerDescribeAttributesParams{
		timeout: timeout,
	}
}

// NewPricingManagerDescribeAttributesParamsWithContext creates a new PricingManagerDescribeAttributesParams object
// with the ability to set a context for a request.
func NewPricingManagerDescribeAttributesParamsWithContext(ctx context.Context) *PricingManagerDescribeAttributesParams {
	return &PricingManagerDescribeAttributesParams{
		Context: ctx,
	}
}

// NewPricingManagerDescribeAttributesParamsWithHTTPClient creates a new PricingManagerDescribeAttributesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPricingManagerDescribeAttributesParamsWithHTTPClient(client *http.Client) *PricingManagerDescribeAttributesParams {
	return &PricingManagerDescribeAttributesParams{
		HTTPClient: client,
	}
}

/*
PricingManagerDescribeAttributesParams contains all the parameters to send to the API endpoint

	for the pricing manager describe attributes operation.

	Typically these are written to a http.Request.
*/
type PricingManagerDescribeAttributesParams struct {

	/* AttrID.

	   属性ID.
	*/
	AttrID *string

	/* IsNeedMeter.

	   是否需要计量 gotags:valid:"OneOf(0,1)".

	   Format: int64
	*/
	IsNeedMeter *int64

	/* Limit.

	   数据库查询每页大小 - 默认 20, 最大值 200.

	   Format: uint64
	*/
	Limit *string

	/* Name.

	   属性名称.
	*/
	Name *string

	/* Offset.

	   数据库查询偏移位置 - 默认 0.

	   Format: uint64
	*/
	Offset *string

	/* ProdID.

	   产品ID.
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

	/* Unit.

	   属性值单位.
	*/
	Unit *string

	/* ValueType.

	   属性值类型.
	*/
	ValueType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pricing manager describe attributes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerDescribeAttributesParams) WithDefaults() *PricingManagerDescribeAttributesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pricing manager describe attributes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerDescribeAttributesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) WithTimeout(timeout time.Duration) *PricingManagerDescribeAttributesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) WithContext(ctx context.Context) *PricingManagerDescribeAttributesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) WithHTTPClient(client *http.Client) *PricingManagerDescribeAttributesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttrID adds the attrID to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) WithAttrID(attrID *string) *PricingManagerDescribeAttributesParams {
	o.SetAttrID(attrID)
	return o
}

// SetAttrID adds the attrId to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) SetAttrID(attrID *string) {
	o.AttrID = attrID
}

// WithIsNeedMeter adds the isNeedMeter to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) WithIsNeedMeter(isNeedMeter *int64) *PricingManagerDescribeAttributesParams {
	o.SetIsNeedMeter(isNeedMeter)
	return o
}

// SetIsNeedMeter adds the isNeedMeter to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) SetIsNeedMeter(isNeedMeter *int64) {
	o.IsNeedMeter = isNeedMeter
}

// WithLimit adds the limit to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) WithLimit(limit *string) *PricingManagerDescribeAttributesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithName adds the name to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) WithName(name *string) *PricingManagerDescribeAttributesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) WithOffset(offset *string) *PricingManagerDescribeAttributesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithProdID adds the prodID to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) WithProdID(prodID *string) *PricingManagerDescribeAttributesParams {
	o.SetProdID(prodID)
	return o
}

// SetProdID adds the prodId to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) SetProdID(prodID *string) {
	o.ProdID = prodID
}

// WithReverse adds the reverse to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) WithReverse(reverse *bool) *PricingManagerDescribeAttributesParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSearchWord adds the searchWord to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) WithSearchWord(searchWord *string) *PricingManagerDescribeAttributesParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) WithSortKey(sortKey *string) *PricingManagerDescribeAttributesParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithUnit adds the unit to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) WithUnit(unit *string) *PricingManagerDescribeAttributesParams {
	o.SetUnit(unit)
	return o
}

// SetUnit adds the unit to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) SetUnit(unit *string) {
	o.Unit = unit
}

// WithValueType adds the valueType to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) WithValueType(valueType *string) *PricingManagerDescribeAttributesParams {
	o.SetValueType(valueType)
	return o
}

// SetValueType adds the valueType to the pricing manager describe attributes params
func (o *PricingManagerDescribeAttributesParams) SetValueType(valueType *string) {
	o.ValueType = valueType
}

// WriteToRequest writes these params to a swagger request
func (o *PricingManagerDescribeAttributesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AttrID != nil {

		// query param attr_id
		var qrAttrID string

		if o.AttrID != nil {
			qrAttrID = *o.AttrID
		}
		qAttrID := qrAttrID
		if qAttrID != "" {

			if err := r.SetQueryParam("attr_id", qAttrID); err != nil {
				return err
			}
		}
	}

	if o.IsNeedMeter != nil {

		// query param is_need_meter
		var qrIsNeedMeter int64

		if o.IsNeedMeter != nil {
			qrIsNeedMeter = *o.IsNeedMeter
		}
		qIsNeedMeter := swag.FormatInt64(qrIsNeedMeter)
		if qIsNeedMeter != "" {

			if err := r.SetQueryParam("is_need_meter", qIsNeedMeter); err != nil {
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

	if o.Unit != nil {

		// query param unit
		var qrUnit string

		if o.Unit != nil {
			qrUnit = *o.Unit
		}
		qUnit := qrUnit
		if qUnit != "" {

			if err := r.SetQueryParam("unit", qUnit); err != nil {
				return err
			}
		}
	}

	if o.ValueType != nil {

		// query param value_type
		var qrValueType string

		if o.ValueType != nil {
			qrValueType = *o.ValueType
		}
		qValueType := qrValueType
		if qValueType != "" {

			if err := r.SetQueryParam("value_type", qValueType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
