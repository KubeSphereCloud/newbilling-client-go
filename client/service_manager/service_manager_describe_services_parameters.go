// Code generated by go-swagger; DO NOT EDIT.

package service_manager

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

// NewServiceManagerDescribeServicesParams creates a new ServiceManagerDescribeServicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServiceManagerDescribeServicesParams() *ServiceManagerDescribeServicesParams {
	return &ServiceManagerDescribeServicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServiceManagerDescribeServicesParamsWithTimeout creates a new ServiceManagerDescribeServicesParams object
// with the ability to set a timeout on a request.
func NewServiceManagerDescribeServicesParamsWithTimeout(timeout time.Duration) *ServiceManagerDescribeServicesParams {
	return &ServiceManagerDescribeServicesParams{
		timeout: timeout,
	}
}

// NewServiceManagerDescribeServicesParamsWithContext creates a new ServiceManagerDescribeServicesParams object
// with the ability to set a context for a request.
func NewServiceManagerDescribeServicesParamsWithContext(ctx context.Context) *ServiceManagerDescribeServicesParams {
	return &ServiceManagerDescribeServicesParams{
		Context: ctx,
	}
}

// NewServiceManagerDescribeServicesParamsWithHTTPClient creates a new ServiceManagerDescribeServicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewServiceManagerDescribeServicesParamsWithHTTPClient(client *http.Client) *ServiceManagerDescribeServicesParams {
	return &ServiceManagerDescribeServicesParams{
		HTTPClient: client,
	}
}

/*
ServiceManagerDescribeServicesParams contains all the parameters to send to the API endpoint

	for the service manager describe services operation.

	Typically these are written to a http.Request.
*/
type ServiceManagerDescribeServicesParams struct {

	/* AccessSysID.

	   接入系统ID.
	*/
	AccessSysID *string

	/* Channel.

	   对接服务渠道.
	*/
	Channel []string

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

	/* ScopeID.

	   对接系统服务配置对象id.
	*/
	ScopeID []string

	/* ScopeType.

	   对接系统服务配置对象.
	*/
	ScopeType []string

	/* SearchWord.

	   模糊查询关键字.
	*/
	SearchWord *string

	/* Service.

	   对接服务类型，目前支持 account/payment/metering/notifier.
	*/
	Service []string

	/* ServiceID.

	   对接服务ID.
	*/
	ServiceID []string

	/* SortKey.

	   排序字段 - 默认 create_time.
	*/
	SortKey *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service manager describe services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceManagerDescribeServicesParams) WithDefaults() *ServiceManagerDescribeServicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service manager describe services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceManagerDescribeServicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) WithTimeout(timeout time.Duration) *ServiceManagerDescribeServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) WithContext(ctx context.Context) *ServiceManagerDescribeServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) WithHTTPClient(client *http.Client) *ServiceManagerDescribeServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessSysID adds the accessSysID to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) WithAccessSysID(accessSysID *string) *ServiceManagerDescribeServicesParams {
	o.SetAccessSysID(accessSysID)
	return o
}

// SetAccessSysID adds the accessSysId to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) SetAccessSysID(accessSysID *string) {
	o.AccessSysID = accessSysID
}

// WithChannel adds the channel to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) WithChannel(channel []string) *ServiceManagerDescribeServicesParams {
	o.SetChannel(channel)
	return o
}

// SetChannel adds the channel to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) SetChannel(channel []string) {
	o.Channel = channel
}

// WithLimit adds the limit to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) WithLimit(limit *string) *ServiceManagerDescribeServicesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) WithOffset(offset *string) *ServiceManagerDescribeServicesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithReverse adds the reverse to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) WithReverse(reverse *bool) *ServiceManagerDescribeServicesParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithScopeID adds the scopeID to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) WithScopeID(scopeID []string) *ServiceManagerDescribeServicesParams {
	o.SetScopeID(scopeID)
	return o
}

// SetScopeID adds the scopeId to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) SetScopeID(scopeID []string) {
	o.ScopeID = scopeID
}

// WithScopeType adds the scopeType to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) WithScopeType(scopeType []string) *ServiceManagerDescribeServicesParams {
	o.SetScopeType(scopeType)
	return o
}

// SetScopeType adds the scopeType to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) SetScopeType(scopeType []string) {
	o.ScopeType = scopeType
}

// WithSearchWord adds the searchWord to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) WithSearchWord(searchWord *string) *ServiceManagerDescribeServicesParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithService adds the service to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) WithService(service []string) *ServiceManagerDescribeServicesParams {
	o.SetService(service)
	return o
}

// SetService adds the service to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) SetService(service []string) {
	o.Service = service
}

// WithServiceID adds the serviceID to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) WithServiceID(serviceID []string) *ServiceManagerDescribeServicesParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) SetServiceID(serviceID []string) {
	o.ServiceID = serviceID
}

// WithSortKey adds the sortKey to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) WithSortKey(sortKey *string) *ServiceManagerDescribeServicesParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the service manager describe services params
func (o *ServiceManagerDescribeServicesParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceManagerDescribeServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Channel != nil {

		// binding items for channel
		joinedChannel := o.bindParamChannel(reg)

		// query array param channel
		if err := r.SetQueryParam("channel", joinedChannel...); err != nil {
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

	if o.ScopeID != nil {

		// binding items for scope_id
		joinedScopeID := o.bindParamScopeID(reg)

		// query array param scope_id
		if err := r.SetQueryParam("scope_id", joinedScopeID...); err != nil {
			return err
		}
	}

	if o.ScopeType != nil {

		// binding items for scope_type
		joinedScopeType := o.bindParamScopeType(reg)

		// query array param scope_type
		if err := r.SetQueryParam("scope_type", joinedScopeType...); err != nil {
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

	if o.Service != nil {

		// binding items for service
		joinedService := o.bindParamService(reg)

		// query array param service
		if err := r.SetQueryParam("service", joinedService...); err != nil {
			return err
		}
	}

	if o.ServiceID != nil {

		// binding items for service_id
		joinedServiceID := o.bindParamServiceID(reg)

		// query array param service_id
		if err := r.SetQueryParam("service_id", joinedServiceID...); err != nil {
			return err
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

// bindParamServiceManagerDescribeServices binds the parameter channel
func (o *ServiceManagerDescribeServicesParams) bindParamChannel(formats strfmt.Registry) []string {
	channelIR := o.Channel

	var channelIC []string
	for _, channelIIR := range channelIR { // explode []string

		channelIIV := channelIIR // string as string
		channelIC = append(channelIC, channelIIV)
	}

	// items.CollectionFormat: "multi"
	channelIS := swag.JoinByFormat(channelIC, "multi")

	return channelIS
}

// bindParamServiceManagerDescribeServices binds the parameter scope_id
func (o *ServiceManagerDescribeServicesParams) bindParamScopeID(formats strfmt.Registry) []string {
	scopeIDIR := o.ScopeID

	var scopeIDIC []string
	for _, scopeIDIIR := range scopeIDIR { // explode []string

		scopeIDIIV := scopeIDIIR // string as string
		scopeIDIC = append(scopeIDIC, scopeIDIIV)
	}

	// items.CollectionFormat: "multi"
	scopeIDIS := swag.JoinByFormat(scopeIDIC, "multi")

	return scopeIDIS
}

// bindParamServiceManagerDescribeServices binds the parameter scope_type
func (o *ServiceManagerDescribeServicesParams) bindParamScopeType(formats strfmt.Registry) []string {
	scopeTypeIR := o.ScopeType

	var scopeTypeIC []string
	for _, scopeTypeIIR := range scopeTypeIR { // explode []string

		scopeTypeIIV := scopeTypeIIR // string as string
		scopeTypeIC = append(scopeTypeIC, scopeTypeIIV)
	}

	// items.CollectionFormat: "multi"
	scopeTypeIS := swag.JoinByFormat(scopeTypeIC, "multi")

	return scopeTypeIS
}

// bindParamServiceManagerDescribeServices binds the parameter service
func (o *ServiceManagerDescribeServicesParams) bindParamService(formats strfmt.Registry) []string {
	serviceIR := o.Service

	var serviceIC []string
	for _, serviceIIR := range serviceIR { // explode []string

		serviceIIV := serviceIIR // string as string
		serviceIC = append(serviceIC, serviceIIV)
	}

	// items.CollectionFormat: "multi"
	serviceIS := swag.JoinByFormat(serviceIC, "multi")

	return serviceIS
}

// bindParamServiceManagerDescribeServices binds the parameter service_id
func (o *ServiceManagerDescribeServicesParams) bindParamServiceID(formats strfmt.Registry) []string {
	serviceIDIR := o.ServiceID

	var serviceIDIC []string
	for _, serviceIDIIR := range serviceIDIR { // explode []string

		serviceIDIIV := serviceIDIIR // string as string
		serviceIDIC = append(serviceIDIC, serviceIDIIV)
	}

	// items.CollectionFormat: "multi"
	serviceIDIS := swag.JoinByFormat(serviceIDIC, "multi")

	return serviceIDIS
}
