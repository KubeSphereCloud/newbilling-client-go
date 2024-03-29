// Code generated by go-swagger; DO NOT EDIT.

package charging_manager

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

// NewChargingManagerGetQrcodeParams creates a new ChargingManagerGetQrcodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChargingManagerGetQrcodeParams() *ChargingManagerGetQrcodeParams {
	return &ChargingManagerGetQrcodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChargingManagerGetQrcodeParamsWithTimeout creates a new ChargingManagerGetQrcodeParams object
// with the ability to set a timeout on a request.
func NewChargingManagerGetQrcodeParamsWithTimeout(timeout time.Duration) *ChargingManagerGetQrcodeParams {
	return &ChargingManagerGetQrcodeParams{
		timeout: timeout,
	}
}

// NewChargingManagerGetQrcodeParamsWithContext creates a new ChargingManagerGetQrcodeParams object
// with the ability to set a context for a request.
func NewChargingManagerGetQrcodeParamsWithContext(ctx context.Context) *ChargingManagerGetQrcodeParams {
	return &ChargingManagerGetQrcodeParams{
		Context: ctx,
	}
}

// NewChargingManagerGetQrcodeParamsWithHTTPClient creates a new ChargingManagerGetQrcodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewChargingManagerGetQrcodeParamsWithHTTPClient(client *http.Client) *ChargingManagerGetQrcodeParams {
	return &ChargingManagerGetQrcodeParams{
		HTTPClient: client,
	}
}

/*
ChargingManagerGetQrcodeParams contains all the parameters to send to the API endpoint

	for the charging manager get qrcode operation.

	Typically these are written to a http.Request.
*/
type ChargingManagerGetQrcodeParams struct {

	/* ChargeAmount.

	   支付金额.

	   Format: float
	*/
	ChargeAmount *float32

	/* ChargeChannel.

	   支付渠道- wx ali.
	*/
	ChargeChannel *string

	/* CurrencyCode.

	   支付货币单位.
	*/
	CurrencyCode *string

	/* CustomerID.

	   客户ID.
	*/
	CustomerID *string

	/* OrderID.

	   newbilling的order id，充值时允许为空，支付时不允许为空.
	*/
	OrderID *string

	/* PassbackParams.

	   passback params.
	*/
	PassbackParams *string

	/* PayType.

	   付款类型- charging 支付，recharge 充值.
	*/
	PayType *string

	/* Remarks.

	   备注.
	*/
	Remarks *string

	/* ReturnURL.

	   支付成功后的跳转地址,允许为空，newbilling提供默认跳转地址.
	*/
	ReturnURL *string

	/* Target.

	   二维码展示形式 target: new，新开的支付宝的页面展示二维码，current, 在当前页面展示二维码.
	*/
	Target *string

	/* TradeName.

	   交易商品名称- 比如充值2000元、购买主机.
	*/
	TradeName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the charging manager get qrcode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChargingManagerGetQrcodeParams) WithDefaults() *ChargingManagerGetQrcodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the charging manager get qrcode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChargingManagerGetQrcodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) WithTimeout(timeout time.Duration) *ChargingManagerGetQrcodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) WithContext(ctx context.Context) *ChargingManagerGetQrcodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) WithHTTPClient(client *http.Client) *ChargingManagerGetQrcodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChargeAmount adds the chargeAmount to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) WithChargeAmount(chargeAmount *float32) *ChargingManagerGetQrcodeParams {
	o.SetChargeAmount(chargeAmount)
	return o
}

// SetChargeAmount adds the chargeAmount to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) SetChargeAmount(chargeAmount *float32) {
	o.ChargeAmount = chargeAmount
}

// WithChargeChannel adds the chargeChannel to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) WithChargeChannel(chargeChannel *string) *ChargingManagerGetQrcodeParams {
	o.SetChargeChannel(chargeChannel)
	return o
}

// SetChargeChannel adds the chargeChannel to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) SetChargeChannel(chargeChannel *string) {
	o.ChargeChannel = chargeChannel
}

// WithCurrencyCode adds the currencyCode to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) WithCurrencyCode(currencyCode *string) *ChargingManagerGetQrcodeParams {
	o.SetCurrencyCode(currencyCode)
	return o
}

// SetCurrencyCode adds the currencyCode to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) SetCurrencyCode(currencyCode *string) {
	o.CurrencyCode = currencyCode
}

// WithCustomerID adds the customerID to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) WithCustomerID(customerID *string) *ChargingManagerGetQrcodeParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) SetCustomerID(customerID *string) {
	o.CustomerID = customerID
}

// WithOrderID adds the orderID to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) WithOrderID(orderID *string) *ChargingManagerGetQrcodeParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) SetOrderID(orderID *string) {
	o.OrderID = orderID
}

// WithPassbackParams adds the passbackParams to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) WithPassbackParams(passbackParams *string) *ChargingManagerGetQrcodeParams {
	o.SetPassbackParams(passbackParams)
	return o
}

// SetPassbackParams adds the passbackParams to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) SetPassbackParams(passbackParams *string) {
	o.PassbackParams = passbackParams
}

// WithPayType adds the payType to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) WithPayType(payType *string) *ChargingManagerGetQrcodeParams {
	o.SetPayType(payType)
	return o
}

// SetPayType adds the payType to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) SetPayType(payType *string) {
	o.PayType = payType
}

// WithRemarks adds the remarks to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) WithRemarks(remarks *string) *ChargingManagerGetQrcodeParams {
	o.SetRemarks(remarks)
	return o
}

// SetRemarks adds the remarks to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) SetRemarks(remarks *string) {
	o.Remarks = remarks
}

// WithReturnURL adds the returnURL to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) WithReturnURL(returnURL *string) *ChargingManagerGetQrcodeParams {
	o.SetReturnURL(returnURL)
	return o
}

// SetReturnURL adds the returnUrl to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) SetReturnURL(returnURL *string) {
	o.ReturnURL = returnURL
}

// WithTarget adds the target to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) WithTarget(target *string) *ChargingManagerGetQrcodeParams {
	o.SetTarget(target)
	return o
}

// SetTarget adds the target to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) SetTarget(target *string) {
	o.Target = target
}

// WithTradeName adds the tradeName to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) WithTradeName(tradeName *string) *ChargingManagerGetQrcodeParams {
	o.SetTradeName(tradeName)
	return o
}

// SetTradeName adds the tradeName to the charging manager get qrcode params
func (o *ChargingManagerGetQrcodeParams) SetTradeName(tradeName *string) {
	o.TradeName = tradeName
}

// WriteToRequest writes these params to a swagger request
func (o *ChargingManagerGetQrcodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ChargeAmount != nil {

		// query param charge_amount
		var qrChargeAmount float32

		if o.ChargeAmount != nil {
			qrChargeAmount = *o.ChargeAmount
		}
		qChargeAmount := swag.FormatFloat32(qrChargeAmount)
		if qChargeAmount != "" {

			if err := r.SetQueryParam("charge_amount", qChargeAmount); err != nil {
				return err
			}
		}
	}

	if o.ChargeChannel != nil {

		// query param charge_channel
		var qrChargeChannel string

		if o.ChargeChannel != nil {
			qrChargeChannel = *o.ChargeChannel
		}
		qChargeChannel := qrChargeChannel
		if qChargeChannel != "" {

			if err := r.SetQueryParam("charge_channel", qChargeChannel); err != nil {
				return err
			}
		}
	}

	if o.CurrencyCode != nil {

		// query param currency_code
		var qrCurrencyCode string

		if o.CurrencyCode != nil {
			qrCurrencyCode = *o.CurrencyCode
		}
		qCurrencyCode := qrCurrencyCode
		if qCurrencyCode != "" {

			if err := r.SetQueryParam("currency_code", qCurrencyCode); err != nil {
				return err
			}
		}
	}

	if o.CustomerID != nil {

		// query param customer_id
		var qrCustomerID string

		if o.CustomerID != nil {
			qrCustomerID = *o.CustomerID
		}
		qCustomerID := qrCustomerID
		if qCustomerID != "" {

			if err := r.SetQueryParam("customer_id", qCustomerID); err != nil {
				return err
			}
		}
	}

	if o.OrderID != nil {

		// query param order_id
		var qrOrderID string

		if o.OrderID != nil {
			qrOrderID = *o.OrderID
		}
		qOrderID := qrOrderID
		if qOrderID != "" {

			if err := r.SetQueryParam("order_id", qOrderID); err != nil {
				return err
			}
		}
	}

	if o.PassbackParams != nil {

		// query param passback_params
		var qrPassbackParams string

		if o.PassbackParams != nil {
			qrPassbackParams = *o.PassbackParams
		}
		qPassbackParams := qrPassbackParams
		if qPassbackParams != "" {

			if err := r.SetQueryParam("passback_params", qPassbackParams); err != nil {
				return err
			}
		}
	}

	if o.PayType != nil {

		// query param pay_type
		var qrPayType string

		if o.PayType != nil {
			qrPayType = *o.PayType
		}
		qPayType := qrPayType
		if qPayType != "" {

			if err := r.SetQueryParam("pay_type", qPayType); err != nil {
				return err
			}
		}
	}

	if o.Remarks != nil {

		// query param remarks
		var qrRemarks string

		if o.Remarks != nil {
			qrRemarks = *o.Remarks
		}
		qRemarks := qrRemarks
		if qRemarks != "" {

			if err := r.SetQueryParam("remarks", qRemarks); err != nil {
				return err
			}
		}
	}

	if o.ReturnURL != nil {

		// query param return_url
		var qrReturnURL string

		if o.ReturnURL != nil {
			qrReturnURL = *o.ReturnURL
		}
		qReturnURL := qrReturnURL
		if qReturnURL != "" {

			if err := r.SetQueryParam("return_url", qReturnURL); err != nil {
				return err
			}
		}
	}

	if o.Target != nil {

		// query param target
		var qrTarget string

		if o.Target != nil {
			qrTarget = *o.Target
		}
		qTarget := qrTarget
		if qTarget != "" {

			if err := r.SetQueryParam("target", qTarget); err != nil {
				return err
			}
		}
	}

	if o.TradeName != nil {

		// query param trade_name
		var qrTradeName string

		if o.TradeName != nil {
			qrTradeName = *o.TradeName
		}
		qTradeName := qrTradeName
		if qTradeName != "" {

			if err := r.SetQueryParam("trade_name", qTradeName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
