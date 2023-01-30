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
)

// NewSubscriptionProdInstanceManagerResumeProdInstanceParams creates a new SubscriptionProdInstanceManagerResumeProdInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubscriptionProdInstanceManagerResumeProdInstanceParams() *SubscriptionProdInstanceManagerResumeProdInstanceParams {
	return &SubscriptionProdInstanceManagerResumeProdInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriptionProdInstanceManagerResumeProdInstanceParamsWithTimeout creates a new SubscriptionProdInstanceManagerResumeProdInstanceParams object
// with the ability to set a timeout on a request.
func NewSubscriptionProdInstanceManagerResumeProdInstanceParamsWithTimeout(timeout time.Duration) *SubscriptionProdInstanceManagerResumeProdInstanceParams {
	return &SubscriptionProdInstanceManagerResumeProdInstanceParams{
		timeout: timeout,
	}
}

// NewSubscriptionProdInstanceManagerResumeProdInstanceParamsWithContext creates a new SubscriptionProdInstanceManagerResumeProdInstanceParams object
// with the ability to set a context for a request.
func NewSubscriptionProdInstanceManagerResumeProdInstanceParamsWithContext(ctx context.Context) *SubscriptionProdInstanceManagerResumeProdInstanceParams {
	return &SubscriptionProdInstanceManagerResumeProdInstanceParams{
		Context: ctx,
	}
}

// NewSubscriptionProdInstanceManagerResumeProdInstanceParamsWithHTTPClient creates a new SubscriptionProdInstanceManagerResumeProdInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubscriptionProdInstanceManagerResumeProdInstanceParamsWithHTTPClient(client *http.Client) *SubscriptionProdInstanceManagerResumeProdInstanceParams {
	return &SubscriptionProdInstanceManagerResumeProdInstanceParams{
		HTTPClient: client,
	}
}

/*
SubscriptionProdInstanceManagerResumeProdInstanceParams contains all the parameters to send to the API endpoint

	for the subscription prod instance manager resume prod instance operation.

	Typically these are written to a http.Request.
*/
type SubscriptionProdInstanceManagerResumeProdInstanceParams struct {

	/* ProdInstIDExt.

	   【接入系统产品实例ID】
	*/
	ProdInstIDExt string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subscription prod instance manager resume prod instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionProdInstanceManagerResumeProdInstanceParams) WithDefaults() *SubscriptionProdInstanceManagerResumeProdInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subscription prod instance manager resume prod instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionProdInstanceManagerResumeProdInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subscription prod instance manager resume prod instance params
func (o *SubscriptionProdInstanceManagerResumeProdInstanceParams) WithTimeout(timeout time.Duration) *SubscriptionProdInstanceManagerResumeProdInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscription prod instance manager resume prod instance params
func (o *SubscriptionProdInstanceManagerResumeProdInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscription prod instance manager resume prod instance params
func (o *SubscriptionProdInstanceManagerResumeProdInstanceParams) WithContext(ctx context.Context) *SubscriptionProdInstanceManagerResumeProdInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscription prod instance manager resume prod instance params
func (o *SubscriptionProdInstanceManagerResumeProdInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscription prod instance manager resume prod instance params
func (o *SubscriptionProdInstanceManagerResumeProdInstanceParams) WithHTTPClient(client *http.Client) *SubscriptionProdInstanceManagerResumeProdInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscription prod instance manager resume prod instance params
func (o *SubscriptionProdInstanceManagerResumeProdInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProdInstIDExt adds the prodInstIDExt to the subscription prod instance manager resume prod instance params
func (o *SubscriptionProdInstanceManagerResumeProdInstanceParams) WithProdInstIDExt(prodInstIDExt string) *SubscriptionProdInstanceManagerResumeProdInstanceParams {
	o.SetProdInstIDExt(prodInstIDExt)
	return o
}

// SetProdInstIDExt adds the prodInstIdExt to the subscription prod instance manager resume prod instance params
func (o *SubscriptionProdInstanceManagerResumeProdInstanceParams) SetProdInstIDExt(prodInstIDExt string) {
	o.ProdInstIDExt = prodInstIDExt
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriptionProdInstanceManagerResumeProdInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param prod_inst_id_ext
	if err := r.SetPathParam("prod_inst_id_ext", o.ProdInstIDExt); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}