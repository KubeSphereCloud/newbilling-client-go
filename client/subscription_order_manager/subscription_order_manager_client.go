// Code generated by go-swagger; DO NOT EDIT.

package subscription_order_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new subscription order manager API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for subscription order manager API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SubscriptionOrderManagerCheckExportFile(params *SubscriptionOrderManagerCheckExportFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionOrderManagerCheckExportFileOK, error)

	SubscriptionOrderManagerDescribeConsumeOrders(params *SubscriptionOrderManagerDescribeConsumeOrdersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionOrderManagerDescribeConsumeOrdersOK, error)

	SubscriptionOrderManagerExportConsumeOrders(params *SubscriptionOrderManagerExportConsumeOrdersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionOrderManagerExportConsumeOrdersOK, error)

	SubscriptionOrderManagerGetConsumeOrder(params *SubscriptionOrderManagerGetConsumeOrderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionOrderManagerGetConsumeOrderOK, error)

	SubscriptionOrderManagerGetExportFile(params *SubscriptionOrderManagerGetExportFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionOrderManagerGetExportFileOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
SubscriptionOrderManagerCheckExportFile 检查导出文件是否生成完成s
*/
func (a *Client) SubscriptionOrderManagerCheckExportFile(params *SubscriptionOrderManagerCheckExportFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionOrderManagerCheckExportFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionOrderManagerCheckExportFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SubscriptionOrderManager_CheckExportFile",
		Method:             "GET",
		PathPattern:        "/v1/export/check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SubscriptionOrderManagerCheckExportFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SubscriptionOrderManagerCheckExportFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SubscriptionOrderManagerCheckExportFileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SubscriptionOrderManagerDescribeConsumeOrders 查询子订单列表s
*/
func (a *Client) SubscriptionOrderManagerDescribeConsumeOrders(params *SubscriptionOrderManagerDescribeConsumeOrdersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionOrderManagerDescribeConsumeOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionOrderManagerDescribeConsumeOrdersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SubscriptionOrderManager_DescribeConsumeOrders",
		Method:             "GET",
		PathPattern:        "/v1/orders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SubscriptionOrderManagerDescribeConsumeOrdersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SubscriptionOrderManagerDescribeConsumeOrdersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SubscriptionOrderManagerDescribeConsumeOrdersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SubscriptionOrderManagerExportConsumeOrders 导出订单列表s
*/
func (a *Client) SubscriptionOrderManagerExportConsumeOrders(params *SubscriptionOrderManagerExportConsumeOrdersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionOrderManagerExportConsumeOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionOrderManagerExportConsumeOrdersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SubscriptionOrderManager_ExportConsumeOrders",
		Method:             "GET",
		PathPattern:        "/v1/orders/export",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SubscriptionOrderManagerExportConsumeOrdersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SubscriptionOrderManagerExportConsumeOrdersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SubscriptionOrderManagerExportConsumeOrdersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SubscriptionOrderManagerGetConsumeOrder 查询单个订单详情s
*/
func (a *Client) SubscriptionOrderManagerGetConsumeOrder(params *SubscriptionOrderManagerGetConsumeOrderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionOrderManagerGetConsumeOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionOrderManagerGetConsumeOrderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SubscriptionOrderManager_GetConsumeOrder",
		Method:             "GET",
		PathPattern:        "/v1/orders/{prod_inst_consume_order_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SubscriptionOrderManagerGetConsumeOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SubscriptionOrderManagerGetConsumeOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SubscriptionOrderManagerGetConsumeOrderDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SubscriptionOrderManagerGetExportFile 下载导出文件s
*/
func (a *Client) SubscriptionOrderManagerGetExportFile(params *SubscriptionOrderManagerGetExportFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionOrderManagerGetExportFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionOrderManagerGetExportFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SubscriptionOrderManager_GetExportFile",
		Method:             "GET",
		PathPattern:        "/v1/export/download",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SubscriptionOrderManagerGetExportFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SubscriptionOrderManagerGetExportFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SubscriptionOrderManagerGetExportFileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
