// Code generated by go-swagger; DO NOT EDIT.

package payment

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

	"github.com/KubeSphereCloud/newbilling-client-go/models"
)

// NewPaymentPrepayAlipayAppParams creates a new PaymentPrepayAlipayAppParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPaymentPrepayAlipayAppParams() *PaymentPrepayAlipayAppParams {
	return &PaymentPrepayAlipayAppParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPaymentPrepayAlipayAppParamsWithTimeout creates a new PaymentPrepayAlipayAppParams object
// with the ability to set a timeout on a request.
func NewPaymentPrepayAlipayAppParamsWithTimeout(timeout time.Duration) *PaymentPrepayAlipayAppParams {
	return &PaymentPrepayAlipayAppParams{
		timeout: timeout,
	}
}

// NewPaymentPrepayAlipayAppParamsWithContext creates a new PaymentPrepayAlipayAppParams object
// with the ability to set a context for a request.
func NewPaymentPrepayAlipayAppParamsWithContext(ctx context.Context) *PaymentPrepayAlipayAppParams {
	return &PaymentPrepayAlipayAppParams{
		Context: ctx,
	}
}

// NewPaymentPrepayAlipayAppParamsWithHTTPClient creates a new PaymentPrepayAlipayAppParams object
// with the ability to set a custom HTTPClient for a request.
func NewPaymentPrepayAlipayAppParamsWithHTTPClient(client *http.Client) *PaymentPrepayAlipayAppParams {
	return &PaymentPrepayAlipayAppParams{
		HTTPClient: client,
	}
}

/*
PaymentPrepayAlipayAppParams contains all the parameters to send to the API endpoint

	for the payment prepay alipay app operation.

	Typically these are written to a http.Request.
*/
type PaymentPrepayAlipayAppParams struct {

	// Body.
	Body *models.NewbillingPrepayAlipayAppRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the payment prepay alipay app params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentPrepayAlipayAppParams) WithDefaults() *PaymentPrepayAlipayAppParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the payment prepay alipay app params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentPrepayAlipayAppParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the payment prepay alipay app params
func (o *PaymentPrepayAlipayAppParams) WithTimeout(timeout time.Duration) *PaymentPrepayAlipayAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the payment prepay alipay app params
func (o *PaymentPrepayAlipayAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the payment prepay alipay app params
func (o *PaymentPrepayAlipayAppParams) WithContext(ctx context.Context) *PaymentPrepayAlipayAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the payment prepay alipay app params
func (o *PaymentPrepayAlipayAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the payment prepay alipay app params
func (o *PaymentPrepayAlipayAppParams) WithHTTPClient(client *http.Client) *PaymentPrepayAlipayAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the payment prepay alipay app params
func (o *PaymentPrepayAlipayAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the payment prepay alipay app params
func (o *PaymentPrepayAlipayAppParams) WithBody(body *models.NewbillingPrepayAlipayAppRequest) *PaymentPrepayAlipayAppParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the payment prepay alipay app params
func (o *PaymentPrepayAlipayAppParams) SetBody(body *models.NewbillingPrepayAlipayAppRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PaymentPrepayAlipayAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}