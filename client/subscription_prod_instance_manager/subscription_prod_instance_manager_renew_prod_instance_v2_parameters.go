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

	"github.com/KubeSphereCloud/newbilling-client-go/models"
)

// NewSubscriptionProdInstanceManagerRenewProdInstanceV2Params creates a new SubscriptionProdInstanceManagerRenewProdInstanceV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubscriptionProdInstanceManagerRenewProdInstanceV2Params() *SubscriptionProdInstanceManagerRenewProdInstanceV2Params {
	return &SubscriptionProdInstanceManagerRenewProdInstanceV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriptionProdInstanceManagerRenewProdInstanceV2ParamsWithTimeout creates a new SubscriptionProdInstanceManagerRenewProdInstanceV2Params object
// with the ability to set a timeout on a request.
func NewSubscriptionProdInstanceManagerRenewProdInstanceV2ParamsWithTimeout(timeout time.Duration) *SubscriptionProdInstanceManagerRenewProdInstanceV2Params {
	return &SubscriptionProdInstanceManagerRenewProdInstanceV2Params{
		timeout: timeout,
	}
}

// NewSubscriptionProdInstanceManagerRenewProdInstanceV2ParamsWithContext creates a new SubscriptionProdInstanceManagerRenewProdInstanceV2Params object
// with the ability to set a context for a request.
func NewSubscriptionProdInstanceManagerRenewProdInstanceV2ParamsWithContext(ctx context.Context) *SubscriptionProdInstanceManagerRenewProdInstanceV2Params {
	return &SubscriptionProdInstanceManagerRenewProdInstanceV2Params{
		Context: ctx,
	}
}

// NewSubscriptionProdInstanceManagerRenewProdInstanceV2ParamsWithHTTPClient creates a new SubscriptionProdInstanceManagerRenewProdInstanceV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewSubscriptionProdInstanceManagerRenewProdInstanceV2ParamsWithHTTPClient(client *http.Client) *SubscriptionProdInstanceManagerRenewProdInstanceV2Params {
	return &SubscriptionProdInstanceManagerRenewProdInstanceV2Params{
		HTTPClient: client,
	}
}

/*
SubscriptionProdInstanceManagerRenewProdInstanceV2Params contains all the parameters to send to the API endpoint

	for the subscription prod instance manager renew prod instance v2 operation.

	Typically these are written to a http.Request.
*/
type SubscriptionProdInstanceManagerRenewProdInstanceV2Params struct {

	/* Body.

	   charge_channel: instant 实时支付，即第三方支付,为""时是指用余额支付
	*/
	Body *models.NewbillingRenewProdInstanceV2RequestRenewParams

	/* ProdInstIDExt.

	   产品实例ID
	*/
	ProdInstIDExt string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subscription prod instance manager renew prod instance v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionProdInstanceManagerRenewProdInstanceV2Params) WithDefaults() *SubscriptionProdInstanceManagerRenewProdInstanceV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subscription prod instance manager renew prod instance v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionProdInstanceManagerRenewProdInstanceV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subscription prod instance manager renew prod instance v2 params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceV2Params) WithTimeout(timeout time.Duration) *SubscriptionProdInstanceManagerRenewProdInstanceV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscription prod instance manager renew prod instance v2 params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscription prod instance manager renew prod instance v2 params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceV2Params) WithContext(ctx context.Context) *SubscriptionProdInstanceManagerRenewProdInstanceV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscription prod instance manager renew prod instance v2 params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscription prod instance manager renew prod instance v2 params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceV2Params) WithHTTPClient(client *http.Client) *SubscriptionProdInstanceManagerRenewProdInstanceV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscription prod instance manager renew prod instance v2 params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the subscription prod instance manager renew prod instance v2 params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceV2Params) WithBody(body *models.NewbillingRenewProdInstanceV2RequestRenewParams) *SubscriptionProdInstanceManagerRenewProdInstanceV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the subscription prod instance manager renew prod instance v2 params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceV2Params) SetBody(body *models.NewbillingRenewProdInstanceV2RequestRenewParams) {
	o.Body = body
}

// WithProdInstIDExt adds the prodInstIDExt to the subscription prod instance manager renew prod instance v2 params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceV2Params) WithProdInstIDExt(prodInstIDExt string) *SubscriptionProdInstanceManagerRenewProdInstanceV2Params {
	o.SetProdInstIDExt(prodInstIDExt)
	return o
}

// SetProdInstIDExt adds the prodInstIdExt to the subscription prod instance manager renew prod instance v2 params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceV2Params) SetProdInstIDExt(prodInstIDExt string) {
	o.ProdInstIDExt = prodInstIDExt
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriptionProdInstanceManagerRenewProdInstanceV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param prod_inst_id_ext
	if err := r.SetPathParam("prod_inst_id_ext", o.ProdInstIDExt); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}