// Code generated by go-swagger; DO NOT EDIT.

package collect_manager

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

// NewCollectManagerDescribeCollectEventsParams creates a new CollectManagerDescribeCollectEventsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCollectManagerDescribeCollectEventsParams() *CollectManagerDescribeCollectEventsParams {
	return &CollectManagerDescribeCollectEventsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCollectManagerDescribeCollectEventsParamsWithTimeout creates a new CollectManagerDescribeCollectEventsParams object
// with the ability to set a timeout on a request.
func NewCollectManagerDescribeCollectEventsParamsWithTimeout(timeout time.Duration) *CollectManagerDescribeCollectEventsParams {
	return &CollectManagerDescribeCollectEventsParams{
		timeout: timeout,
	}
}

// NewCollectManagerDescribeCollectEventsParamsWithContext creates a new CollectManagerDescribeCollectEventsParams object
// with the ability to set a context for a request.
func NewCollectManagerDescribeCollectEventsParamsWithContext(ctx context.Context) *CollectManagerDescribeCollectEventsParams {
	return &CollectManagerDescribeCollectEventsParams{
		Context: ctx,
	}
}

// NewCollectManagerDescribeCollectEventsParamsWithHTTPClient creates a new CollectManagerDescribeCollectEventsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCollectManagerDescribeCollectEventsParamsWithHTTPClient(client *http.Client) *CollectManagerDescribeCollectEventsParams {
	return &CollectManagerDescribeCollectEventsParams{
		HTTPClient: client,
	}
}

/*
CollectManagerDescribeCollectEventsParams contains all the parameters to send to the API endpoint

	for the collect manager describe collect events operation.

	Typically these are written to a http.Request.
*/
type CollectManagerDescribeCollectEventsParams struct {

	/* AccessSysID.

	   接入系统ID.
	*/
	AccessSysID *string

	/* CollectEventID.

	   采集事件ID.
	*/
	CollectEventID []string

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

	/* ProdID.

	   产品ID.
	*/
	ProdID []string

	/* ProdInstIDExt.

	   外部产品实例ID.
	*/
	ProdInstIDExt []string

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

	   采集事件状态.
	*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the collect manager describe collect events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CollectManagerDescribeCollectEventsParams) WithDefaults() *CollectManagerDescribeCollectEventsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the collect manager describe collect events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CollectManagerDescribeCollectEventsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) WithTimeout(timeout time.Duration) *CollectManagerDescribeCollectEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) WithContext(ctx context.Context) *CollectManagerDescribeCollectEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) WithHTTPClient(client *http.Client) *CollectManagerDescribeCollectEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessSysID adds the accessSysID to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) WithAccessSysID(accessSysID *string) *CollectManagerDescribeCollectEventsParams {
	o.SetAccessSysID(accessSysID)
	return o
}

// SetAccessSysID adds the accessSysId to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) SetAccessSysID(accessSysID *string) {
	o.AccessSysID = accessSysID
}

// WithCollectEventID adds the collectEventID to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) WithCollectEventID(collectEventID []string) *CollectManagerDescribeCollectEventsParams {
	o.SetCollectEventID(collectEventID)
	return o
}

// SetCollectEventID adds the collectEventId to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) SetCollectEventID(collectEventID []string) {
	o.CollectEventID = collectEventID
}

// WithLimit adds the limit to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) WithLimit(limit *string) *CollectManagerDescribeCollectEventsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) WithOffset(offset *string) *CollectManagerDescribeCollectEventsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithProdID adds the prodID to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) WithProdID(prodID []string) *CollectManagerDescribeCollectEventsParams {
	o.SetProdID(prodID)
	return o
}

// SetProdID adds the prodId to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) SetProdID(prodID []string) {
	o.ProdID = prodID
}

// WithProdInstIDExt adds the prodInstIDExt to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) WithProdInstIDExt(prodInstIDExt []string) *CollectManagerDescribeCollectEventsParams {
	o.SetProdInstIDExt(prodInstIDExt)
	return o
}

// SetProdInstIDExt adds the prodInstIdExt to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) SetProdInstIDExt(prodInstIDExt []string) {
	o.ProdInstIDExt = prodInstIDExt
}

// WithReverse adds the reverse to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) WithReverse(reverse *bool) *CollectManagerDescribeCollectEventsParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSearchWord adds the searchWord to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) WithSearchWord(searchWord *string) *CollectManagerDescribeCollectEventsParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) WithSortKey(sortKey *string) *CollectManagerDescribeCollectEventsParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStatus adds the status to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) WithStatus(status []string) *CollectManagerDescribeCollectEventsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the collect manager describe collect events params
func (o *CollectManagerDescribeCollectEventsParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *CollectManagerDescribeCollectEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.CollectEventID != nil {

		// binding items for collect_event_id
		joinedCollectEventID := o.bindParamCollectEventID(reg)

		// query array param collect_event_id
		if err := r.SetQueryParam("collect_event_id", joinedCollectEventID...); err != nil {
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

	if o.ProdID != nil {

		// binding items for prod_id
		joinedProdID := o.bindParamProdID(reg)

		// query array param prod_id
		if err := r.SetQueryParam("prod_id", joinedProdID...); err != nil {
			return err
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

// bindParamCollectManagerDescribeCollectEvents binds the parameter collect_event_id
func (o *CollectManagerDescribeCollectEventsParams) bindParamCollectEventID(formats strfmt.Registry) []string {
	collectEventIDIR := o.CollectEventID

	var collectEventIDIC []string
	for _, collectEventIDIIR := range collectEventIDIR { // explode []string

		collectEventIDIIV := collectEventIDIIR // string as string
		collectEventIDIC = append(collectEventIDIC, collectEventIDIIV)
	}

	// items.CollectionFormat: "multi"
	collectEventIDIS := swag.JoinByFormat(collectEventIDIC, "multi")

	return collectEventIDIS
}

// bindParamCollectManagerDescribeCollectEvents binds the parameter prod_id
func (o *CollectManagerDescribeCollectEventsParams) bindParamProdID(formats strfmt.Registry) []string {
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

// bindParamCollectManagerDescribeCollectEvents binds the parameter prod_inst_id_ext
func (o *CollectManagerDescribeCollectEventsParams) bindParamProdInstIDExt(formats strfmt.Registry) []string {
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

// bindParamCollectManagerDescribeCollectEvents binds the parameter status
func (o *CollectManagerDescribeCollectEventsParams) bindParamStatus(formats strfmt.Registry) []string {
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