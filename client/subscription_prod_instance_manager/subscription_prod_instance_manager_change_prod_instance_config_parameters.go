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

// NewSubscriptionProdInstanceManagerChangeProdInstanceConfigParams creates a new SubscriptionProdInstanceManagerChangeProdInstanceConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubscriptionProdInstanceManagerChangeProdInstanceConfigParams() *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams {
	return &SubscriptionProdInstanceManagerChangeProdInstanceConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriptionProdInstanceManagerChangeProdInstanceConfigParamsWithTimeout creates a new SubscriptionProdInstanceManagerChangeProdInstanceConfigParams object
// with the ability to set a timeout on a request.
func NewSubscriptionProdInstanceManagerChangeProdInstanceConfigParamsWithTimeout(timeout time.Duration) *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams {
	return &SubscriptionProdInstanceManagerChangeProdInstanceConfigParams{
		timeout: timeout,
	}
}

// NewSubscriptionProdInstanceManagerChangeProdInstanceConfigParamsWithContext creates a new SubscriptionProdInstanceManagerChangeProdInstanceConfigParams object
// with the ability to set a context for a request.
func NewSubscriptionProdInstanceManagerChangeProdInstanceConfigParamsWithContext(ctx context.Context) *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams {
	return &SubscriptionProdInstanceManagerChangeProdInstanceConfigParams{
		Context: ctx,
	}
}

// NewSubscriptionProdInstanceManagerChangeProdInstanceConfigParamsWithHTTPClient creates a new SubscriptionProdInstanceManagerChangeProdInstanceConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubscriptionProdInstanceManagerChangeProdInstanceConfigParamsWithHTTPClient(client *http.Client) *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams {
	return &SubscriptionProdInstanceManagerChangeProdInstanceConfigParams{
		HTTPClient: client,
	}
}

/*
SubscriptionProdInstanceManagerChangeProdInstanceConfigParams contains all the parameters to send to the API endpoint

	for the subscription prod instance manager change prod instance config operation.

	Typically these are written to a http.Request.
*/
type SubscriptionProdInstanceManagerChangeProdInstanceConfigParams struct {

	// Body.
	Body *models.NewbillingChangeProdInstanceConfigRequest

	/* ProdInstIDExt.

	   外部产品实例ID
	*/
	ProdInstIDExt string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subscription prod instance manager change prod instance config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams) WithDefaults() *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subscription prod instance manager change prod instance config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subscription prod instance manager change prod instance config params
func (o *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams) WithTimeout(timeout time.Duration) *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscription prod instance manager change prod instance config params
func (o *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscription prod instance manager change prod instance config params
func (o *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams) WithContext(ctx context.Context) *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscription prod instance manager change prod instance config params
func (o *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscription prod instance manager change prod instance config params
func (o *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams) WithHTTPClient(client *http.Client) *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscription prod instance manager change prod instance config params
func (o *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the subscription prod instance manager change prod instance config params
func (o *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams) WithBody(body *models.NewbillingChangeProdInstanceConfigRequest) *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the subscription prod instance manager change prod instance config params
func (o *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams) SetBody(body *models.NewbillingChangeProdInstanceConfigRequest) {
	o.Body = body
}

// WithProdInstIDExt adds the prodInstIDExt to the subscription prod instance manager change prod instance config params
func (o *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams) WithProdInstIDExt(prodInstIDExt string) *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams {
	o.SetProdInstIDExt(prodInstIDExt)
	return o
}

// SetProdInstIDExt adds the prodInstIdExt to the subscription prod instance manager change prod instance config params
func (o *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams) SetProdInstIDExt(prodInstIDExt string) {
	o.ProdInstIDExt = prodInstIDExt
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriptionProdInstanceManagerChangeProdInstanceConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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