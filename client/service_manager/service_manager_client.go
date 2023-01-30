// Code generated by go-swagger; DO NOT EDIT.

package service_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new service manager API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service manager API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ServiceManagerDescribeServices(params *ServiceManagerDescribeServicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceManagerDescribeServicesOK, error)

	ServiceManagerModifyService(params *ServiceManagerModifyServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceManagerModifyServiceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ServiceManagerDescribeServices 对接服务s 获取
*/
func (a *Client) ServiceManagerDescribeServices(params *ServiceManagerDescribeServicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceManagerDescribeServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceManagerDescribeServicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ServiceManager_DescribeServices",
		Method:             "GET",
		PathPattern:        "/v1/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServiceManagerDescribeServicesReader{formats: a.formats},
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
	success, ok := result.(*ServiceManagerDescribeServicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServiceManagerDescribeServicesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ServiceManagerModifyService 对接服务s 修改
*/
func (a *Client) ServiceManagerModifyService(params *ServiceManagerModifyServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceManagerModifyServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceManagerModifyServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ServiceManager_ModifyService",
		Method:             "PATCH",
		PathPattern:        "/v1/service",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServiceManagerModifyServiceReader{formats: a.formats},
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
	success, ok := result.(*ServiceManagerModifyServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServiceManagerModifyServiceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
