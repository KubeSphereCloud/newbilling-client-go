// Code generated by go-swagger; DO NOT EDIT.

package subscription_prod_instance_manager

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

// NewSubscriptionProdInstanceManagerDescribeProdInstancesParams creates a new SubscriptionProdInstanceManagerDescribeProdInstancesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubscriptionProdInstanceManagerDescribeProdInstancesParams() *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	return &SubscriptionProdInstanceManagerDescribeProdInstancesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriptionProdInstanceManagerDescribeProdInstancesParamsWithTimeout creates a new SubscriptionProdInstanceManagerDescribeProdInstancesParams object
// with the ability to set a timeout on a request.
func NewSubscriptionProdInstanceManagerDescribeProdInstancesParamsWithTimeout(timeout time.Duration) *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	return &SubscriptionProdInstanceManagerDescribeProdInstancesParams{
		timeout: timeout,
	}
}

// NewSubscriptionProdInstanceManagerDescribeProdInstancesParamsWithContext creates a new SubscriptionProdInstanceManagerDescribeProdInstancesParams object
// with the ability to set a context for a request.
func NewSubscriptionProdInstanceManagerDescribeProdInstancesParamsWithContext(ctx context.Context) *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	return &SubscriptionProdInstanceManagerDescribeProdInstancesParams{
		Context: ctx,
	}
}

// NewSubscriptionProdInstanceManagerDescribeProdInstancesParamsWithHTTPClient creates a new SubscriptionProdInstanceManagerDescribeProdInstancesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubscriptionProdInstanceManagerDescribeProdInstancesParamsWithHTTPClient(client *http.Client) *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	return &SubscriptionProdInstanceManagerDescribeProdInstancesParams{
		HTTPClient: client,
	}
}

/*
SubscriptionProdInstanceManagerDescribeProdInstancesParams contains all the parameters to send to the API endpoint

	for the subscription prod instance manager describe prod instances operation.

	Typically these are written to a http.Request.
*/
type SubscriptionProdInstanceManagerDescribeProdInstancesParams struct {

	/* IsDeleted.

	   是否删除- 0：否，1：是.
	*/
	IsDeleted []int32

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

	/* OrderID.

	   订单ID.
	*/
	OrderID []string

	/* OrderUserID.

	   订单用户ID.
	*/
	OrderUserID []string

	/* ProdID.

	   产品ID.
	*/
	ProdID []string

	/* ProdInstID.

	   产品实例ID.
	*/
	ProdInstID []string

	/* ProdInstIDExt.

	   外部产品实例ID.
	*/
	ProdInstIDExt []string

	/* Reverse.

	   是否倒序排序- 0：ASC，1：DESC.
	*/
	Reverse *bool

	/* SearchWord.

	   模糊查询关键字- 支持字段：prod_name.
	*/
	SearchWord *string

	/* SortKey.

	   排序字段- 默认 create_time.
	*/
	SortKey *string

	/* Status.

	   产品实例状态.
	*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subscription prod instance manager describe prod instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) WithDefaults() *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subscription prod instance manager describe prod instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) WithTimeout(timeout time.Duration) *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) WithContext(ctx context.Context) *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) WithHTTPClient(client *http.Client) *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIsDeleted adds the isDeleted to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) WithIsDeleted(isDeleted []int32) *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	o.SetIsDeleted(isDeleted)
	return o
}

// SetIsDeleted adds the isDeleted to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) SetIsDeleted(isDeleted []int32) {
	o.IsDeleted = isDeleted
}

// WithLimit adds the limit to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) WithLimit(limit *string) *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) WithOffset(offset *string) *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithOrderID adds the orderID to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) WithOrderID(orderID []string) *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) SetOrderID(orderID []string) {
	o.OrderID = orderID
}

// WithOrderUserID adds the orderUserID to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) WithOrderUserID(orderUserID []string) *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	o.SetOrderUserID(orderUserID)
	return o
}

// SetOrderUserID adds the orderUserId to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) SetOrderUserID(orderUserID []string) {
	o.OrderUserID = orderUserID
}

// WithProdID adds the prodID to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) WithProdID(prodID []string) *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	o.SetProdID(prodID)
	return o
}

// SetProdID adds the prodId to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) SetProdID(prodID []string) {
	o.ProdID = prodID
}

// WithProdInstID adds the prodInstID to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) WithProdInstID(prodInstID []string) *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	o.SetProdInstID(prodInstID)
	return o
}

// SetProdInstID adds the prodInstId to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) SetProdInstID(prodInstID []string) {
	o.ProdInstID = prodInstID
}

// WithProdInstIDExt adds the prodInstIDExt to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) WithProdInstIDExt(prodInstIDExt []string) *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	o.SetProdInstIDExt(prodInstIDExt)
	return o
}

// SetProdInstIDExt adds the prodInstIdExt to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) SetProdInstIDExt(prodInstIDExt []string) {
	o.ProdInstIDExt = prodInstIDExt
}

// WithReverse adds the reverse to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) WithReverse(reverse *bool) *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSearchWord adds the searchWord to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) WithSearchWord(searchWord *string) *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) WithSortKey(sortKey *string) *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStatus adds the status to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) WithStatus(status []string) *SubscriptionProdInstanceManagerDescribeProdInstancesParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the subscription prod instance manager describe prod instances params
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IsDeleted != nil {

		// binding items for is_deleted
		joinedIsDeleted := o.bindParamIsDeleted(reg)

		// query array param is_deleted
		if err := r.SetQueryParam("is_deleted", joinedIsDeleted...); err != nil {
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

	if o.OrderID != nil {

		// binding items for order_id
		joinedOrderID := o.bindParamOrderID(reg)

		// query array param order_id
		if err := r.SetQueryParam("order_id", joinedOrderID...); err != nil {
			return err
		}
	}

	if o.OrderUserID != nil {

		// binding items for order_user_id
		joinedOrderUserID := o.bindParamOrderUserID(reg)

		// query array param order_user_id
		if err := r.SetQueryParam("order_user_id", joinedOrderUserID...); err != nil {
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

	if o.ProdInstID != nil {

		// binding items for prod_inst_id
		joinedProdInstID := o.bindParamProdInstID(reg)

		// query array param prod_inst_id
		if err := r.SetQueryParam("prod_inst_id", joinedProdInstID...); err != nil {
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

// bindParamSubscriptionProdInstanceManagerDescribeProdInstances binds the parameter is_deleted
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) bindParamIsDeleted(formats strfmt.Registry) []string {
	isDeletedIR := o.IsDeleted

	var isDeletedIC []string
	for _, isDeletedIIR := range isDeletedIR { // explode []int32

		isDeletedIIV := swag.FormatInt32(isDeletedIIR) // int32 as string
		isDeletedIC = append(isDeletedIC, isDeletedIIV)
	}

	// items.CollectionFormat: "multi"
	isDeletedIS := swag.JoinByFormat(isDeletedIC, "multi")

	return isDeletedIS
}

// bindParamSubscriptionProdInstanceManagerDescribeProdInstances binds the parameter order_id
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) bindParamOrderID(formats strfmt.Registry) []string {
	orderIDIR := o.OrderID

	var orderIDIC []string
	for _, orderIDIIR := range orderIDIR { // explode []string

		orderIDIIV := orderIDIIR // string as string
		orderIDIC = append(orderIDIC, orderIDIIV)
	}

	// items.CollectionFormat: "multi"
	orderIDIS := swag.JoinByFormat(orderIDIC, "multi")

	return orderIDIS
}

// bindParamSubscriptionProdInstanceManagerDescribeProdInstances binds the parameter order_user_id
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) bindParamOrderUserID(formats strfmt.Registry) []string {
	orderUserIDIR := o.OrderUserID

	var orderUserIDIC []string
	for _, orderUserIDIIR := range orderUserIDIR { // explode []string

		orderUserIDIIV := orderUserIDIIR // string as string
		orderUserIDIC = append(orderUserIDIC, orderUserIDIIV)
	}

	// items.CollectionFormat: "multi"
	orderUserIDIS := swag.JoinByFormat(orderUserIDIC, "multi")

	return orderUserIDIS
}

// bindParamSubscriptionProdInstanceManagerDescribeProdInstances binds the parameter prod_id
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) bindParamProdID(formats strfmt.Registry) []string {
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

// bindParamSubscriptionProdInstanceManagerDescribeProdInstances binds the parameter prod_inst_id
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) bindParamProdInstID(formats strfmt.Registry) []string {
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

// bindParamSubscriptionProdInstanceManagerDescribeProdInstances binds the parameter prod_inst_id_ext
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) bindParamProdInstIDExt(formats strfmt.Registry) []string {
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

// bindParamSubscriptionProdInstanceManagerDescribeProdInstances binds the parameter status
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesParams) bindParamStatus(formats strfmt.Registry) []string {
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
