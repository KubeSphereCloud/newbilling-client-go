// Code generated by go-swagger; DO NOT EDIT.

package pricing_access_system_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pricing access system manager API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pricing access system manager API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PricingAccessSystemManagerCreateAccessSystem(params *PricingAccessSystemManagerCreateAccessSystemParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PricingAccessSystemManagerCreateAccessSystemOK, error)

	PricingAccessSystemManagerCreateAccessSystemForOpen(params *PricingAccessSystemManagerCreateAccessSystemForOpenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PricingAccessSystemManagerCreateAccessSystemForOpenOK, error)

	PricingAccessSystemManagerDeleteAccessSystems(params *PricingAccessSystemManagerDeleteAccessSystemsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PricingAccessSystemManagerDeleteAccessSystemsOK, error)

	PricingAccessSystemManagerDescribeAccessSystems(params *PricingAccessSystemManagerDescribeAccessSystemsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PricingAccessSystemManagerDescribeAccessSystemsOK, error)

	PricingAccessSystemManagerModifyAccessSystem(params *PricingAccessSystemManagerModifyAccessSystemParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PricingAccessSystemManagerModifyAccessSystemOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PricingAccessSystemManagerCreateAccessSystem 接入系统s 创建
*/
func (a *Client) PricingAccessSystemManagerCreateAccessSystem(params *PricingAccessSystemManagerCreateAccessSystemParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PricingAccessSystemManagerCreateAccessSystemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPricingAccessSystemManagerCreateAccessSystemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PricingAccessSystemManager_CreateAccessSystem",
		Method:             "POST",
		PathPattern:        "/v1/accesssystemcatalogs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PricingAccessSystemManagerCreateAccessSystemReader{formats: a.formats},
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
	success, ok := result.(*PricingAccessSystemManagerCreateAccessSystemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PricingAccessSystemManagerCreateAccessSystemDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PricingAccessSystemManagerCreateAccessSystemForOpen appcenter专用接口s 创建接入系统
*/
func (a *Client) PricingAccessSystemManagerCreateAccessSystemForOpen(params *PricingAccessSystemManagerCreateAccessSystemForOpenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PricingAccessSystemManagerCreateAccessSystemForOpenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPricingAccessSystemManagerCreateAccessSystemForOpenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PricingAccessSystemManager_CreateAccessSystemForOpen",
		Method:             "POST",
		PathPattern:        "/v1/open/accesssystemcatalogs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PricingAccessSystemManagerCreateAccessSystemForOpenReader{formats: a.formats},
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
	success, ok := result.(*PricingAccessSystemManagerCreateAccessSystemForOpenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PricingAccessSystemManagerCreateAccessSystemForOpenDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PricingAccessSystemManagerDeleteAccessSystems 接入系统s 删除
*/
func (a *Client) PricingAccessSystemManagerDeleteAccessSystems(params *PricingAccessSystemManagerDeleteAccessSystemsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PricingAccessSystemManagerDeleteAccessSystemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPricingAccessSystemManagerDeleteAccessSystemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PricingAccessSystemManager_DeleteAccessSystems",
		Method:             "DELETE",
		PathPattern:        "/v1/accesssystems",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PricingAccessSystemManagerDeleteAccessSystemsReader{formats: a.formats},
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
	success, ok := result.(*PricingAccessSystemManagerDeleteAccessSystemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PricingAccessSystemManagerDeleteAccessSystemsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PricingAccessSystemManagerDescribeAccessSystems 接入系统s 查询
*/
func (a *Client) PricingAccessSystemManagerDescribeAccessSystems(params *PricingAccessSystemManagerDescribeAccessSystemsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PricingAccessSystemManagerDescribeAccessSystemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPricingAccessSystemManagerDescribeAccessSystemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PricingAccessSystemManager_DescribeAccessSystems",
		Method:             "GET",
		PathPattern:        "/v1/accesssystems",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PricingAccessSystemManagerDescribeAccessSystemsReader{formats: a.formats},
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
	success, ok := result.(*PricingAccessSystemManagerDescribeAccessSystemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PricingAccessSystemManagerDescribeAccessSystemsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PricingAccessSystemManagerModifyAccessSystem 接入系统s 基础信息修改
*/
func (a *Client) PricingAccessSystemManagerModifyAccessSystem(params *PricingAccessSystemManagerModifyAccessSystemParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PricingAccessSystemManagerModifyAccessSystemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPricingAccessSystemManagerModifyAccessSystemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PricingAccessSystemManager_ModifyAccessSystem",
		Method:             "PATCH",
		PathPattern:        "/v1/accesssystems",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PricingAccessSystemManagerModifyAccessSystemReader{formats: a.formats},
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
	success, ok := result.(*PricingAccessSystemManagerModifyAccessSystemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PricingAccessSystemManagerModifyAccessSystemDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
