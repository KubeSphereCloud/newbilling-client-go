// Code generated by go-swagger; DO NOT EDIT.

package charging_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new charging manager API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for charging manager API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ChargingManagerChargeAlipayReturn(params *ChargingManagerChargeAlipayReturnParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChargingManagerChargeAlipayReturnOK, error)

	ChargingManagerCreateCharge(params *ChargingManagerCreateChargeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChargingManagerCreateChargeOK, error)

	ChargingManagerCreateRecharge(params *ChargingManagerCreateRechargeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChargingManagerCreateRechargeOK, error)

	ChargingManagerGetChargeQrcode(params *ChargingManagerGetChargeQrcodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChargingManagerGetChargeQrcodeOK, error)

	ChargingManagerGetQrcode(params *ChargingManagerGetQrcodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChargingManagerGetQrcodeOK, error)

	ChargingManagerListCharges(params *ChargingManagerListChargesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChargingManagerListChargesOK, error)

	ChargingManagerPaymentCallback(params *ChargingManagerPaymentCallbackParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChargingManagerPaymentCallbackOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ChargingManagerChargeAlipayReturn 接收支付宝回调s 为anybox包装的接口
*/
func (a *Client) ChargingManagerChargeAlipayReturn(params *ChargingManagerChargeAlipayReturnParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChargingManagerChargeAlipayReturnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChargingManagerChargeAlipayReturnParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ChargingManager_ChargeAlipayReturn",
		Method:             "POST",
		PathPattern:        "/v1/charge_returns/alipay",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChargingManagerChargeAlipayReturnReader{formats: a.formats},
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
	success, ok := result.(*ChargingManagerChargeAlipayReturnOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChargingManagerChargeAlipayReturnDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ChargingManagerCreateCharge 扣费记录s 创建
*/
func (a *Client) ChargingManagerCreateCharge(params *ChargingManagerCreateChargeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChargingManagerCreateChargeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChargingManagerCreateChargeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ChargingManager_CreateCharge",
		Method:             "POST",
		PathPattern:        "/v1/charges",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChargingManagerCreateChargeReader{formats: a.formats},
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
	success, ok := result.(*ChargingManagerCreateChargeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChargingManagerCreateChargeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ChargingManagerCreateRecharge ns b内部接口 pitrixpay回调 充值记录 创建
*/
func (a *Client) ChargingManagerCreateRecharge(params *ChargingManagerCreateRechargeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChargingManagerCreateRechargeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChargingManagerCreateRechargeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ChargingManager_CreateRecharge",
		Method:             "POST",
		PathPattern:        "/v1/recharges",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChargingManagerCreateRechargeReader{formats: a.formats},
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
	success, ok := result.(*ChargingManagerCreateRechargeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChargingManagerCreateRechargeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ChargingManagerGetChargeQrcode 获取支付宝支付二维码s 为anybox包装的接口
*/
func (a *Client) ChargingManagerGetChargeQrcode(params *ChargingManagerGetChargeQrcodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChargingManagerGetChargeQrcodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChargingManagerGetChargeQrcodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ChargingManager_GetChargeQrcode",
		Method:             "GET",
		PathPattern:        "/v1/charge/qrcode",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChargingManagerGetChargeQrcodeReader{formats: a.formats},
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
	success, ok := result.(*ChargingManagerGetChargeQrcodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChargingManagerGetChargeQrcodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ChargingManagerGetQrcode 获取支付二维码s
*/
func (a *Client) ChargingManagerGetQrcode(params *ChargingManagerGetQrcodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChargingManagerGetQrcodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChargingManagerGetQrcodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ChargingManager_GetQrcode",
		Method:             "GET",
		PathPattern:        "/v1/qrcode",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChargingManagerGetQrcodeReader{formats: a.formats},
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
	success, ok := result.(*ChargingManagerGetQrcodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChargingManagerGetQrcodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ChargingManagerListCharges 查询支付列表s
*/
func (a *Client) ChargingManagerListCharges(params *ChargingManagerListChargesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChargingManagerListChargesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChargingManagerListChargesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ChargingManager_ListCharges",
		Method:             "GET",
		PathPattern:        "/v1/charges",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChargingManagerListChargesReader{formats: a.formats},
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
	success, ok := result.(*ChargingManagerListChargesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChargingManagerListChargesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ChargingManagerPaymentCallback 支付回调s
*/
func (a *Client) ChargingManagerPaymentCallback(params *ChargingManagerPaymentCallbackParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChargingManagerPaymentCallbackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChargingManagerPaymentCallbackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ChargingManager_PaymentCallback",
		Method:             "POST",
		PathPattern:        "/v1/payment:callback",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChargingManagerPaymentCallbackReader{formats: a.formats},
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
	success, ok := result.(*ChargingManagerPaymentCallbackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChargingManagerPaymentCallbackDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}