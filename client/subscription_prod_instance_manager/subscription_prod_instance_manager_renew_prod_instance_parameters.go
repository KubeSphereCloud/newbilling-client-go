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

// NewSubscriptionProdInstanceManagerRenewProdInstanceParams creates a new SubscriptionProdInstanceManagerRenewProdInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubscriptionProdInstanceManagerRenewProdInstanceParams() *SubscriptionProdInstanceManagerRenewProdInstanceParams {
	return &SubscriptionProdInstanceManagerRenewProdInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriptionProdInstanceManagerRenewProdInstanceParamsWithTimeout creates a new SubscriptionProdInstanceManagerRenewProdInstanceParams object
// with the ability to set a timeout on a request.
func NewSubscriptionProdInstanceManagerRenewProdInstanceParamsWithTimeout(timeout time.Duration) *SubscriptionProdInstanceManagerRenewProdInstanceParams {
	return &SubscriptionProdInstanceManagerRenewProdInstanceParams{
		timeout: timeout,
	}
}

// NewSubscriptionProdInstanceManagerRenewProdInstanceParamsWithContext creates a new SubscriptionProdInstanceManagerRenewProdInstanceParams object
// with the ability to set a context for a request.
func NewSubscriptionProdInstanceManagerRenewProdInstanceParamsWithContext(ctx context.Context) *SubscriptionProdInstanceManagerRenewProdInstanceParams {
	return &SubscriptionProdInstanceManagerRenewProdInstanceParams{
		Context: ctx,
	}
}

// NewSubscriptionProdInstanceManagerRenewProdInstanceParamsWithHTTPClient creates a new SubscriptionProdInstanceManagerRenewProdInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubscriptionProdInstanceManagerRenewProdInstanceParamsWithHTTPClient(client *http.Client) *SubscriptionProdInstanceManagerRenewProdInstanceParams {
	return &SubscriptionProdInstanceManagerRenewProdInstanceParams{
		HTTPClient: client,
	}
}

/*
SubscriptionProdInstanceManagerRenewProdInstanceParams contains all the parameters to send to the API endpoint

	for the subscription prod instance manager renew prod instance operation.

	Typically these are written to a http.Request.
*/
type SubscriptionProdInstanceManagerRenewProdInstanceParams struct {

	/* Body.

	   charge_channel: instant 实时支付，即第三方支付,为""时是指用余额支付
	*/
	Body *models.NewbillingRenewProdInstanceRequestRenewParams

	/* ProdInstIDExt.

	   产品实例ID
	*/
	ProdInstIDExt string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subscription prod instance manager renew prod instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionProdInstanceManagerRenewProdInstanceParams) WithDefaults() *SubscriptionProdInstanceManagerRenewProdInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subscription prod instance manager renew prod instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionProdInstanceManagerRenewProdInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subscription prod instance manager renew prod instance params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceParams) WithTimeout(timeout time.Duration) *SubscriptionProdInstanceManagerRenewProdInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscription prod instance manager renew prod instance params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscription prod instance manager renew prod instance params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceParams) WithContext(ctx context.Context) *SubscriptionProdInstanceManagerRenewProdInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscription prod instance manager renew prod instance params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscription prod instance manager renew prod instance params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceParams) WithHTTPClient(client *http.Client) *SubscriptionProdInstanceManagerRenewProdInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscription prod instance manager renew prod instance params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the subscription prod instance manager renew prod instance params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceParams) WithBody(body *models.NewbillingRenewProdInstanceRequestRenewParams) *SubscriptionProdInstanceManagerRenewProdInstanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the subscription prod instance manager renew prod instance params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceParams) SetBody(body *models.NewbillingRenewProdInstanceRequestRenewParams) {
	o.Body = body
}

// WithProdInstIDExt adds the prodInstIDExt to the subscription prod instance manager renew prod instance params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceParams) WithProdInstIDExt(prodInstIDExt string) *SubscriptionProdInstanceManagerRenewProdInstanceParams {
	o.SetProdInstIDExt(prodInstIDExt)
	return o
}

// SetProdInstIDExt adds the prodInstIdExt to the subscription prod instance manager renew prod instance params
func (o *SubscriptionProdInstanceManagerRenewProdInstanceParams) SetProdInstIDExt(prodInstIDExt string) {
	o.ProdInstIDExt = prodInstIDExt
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriptionProdInstanceManagerRenewProdInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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