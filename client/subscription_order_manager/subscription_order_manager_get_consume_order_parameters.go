// Code generated by go-swagger; DO NOT EDIT.

package subscription_order_manager

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

// NewSubscriptionOrderManagerGetConsumeOrderParams creates a new SubscriptionOrderManagerGetConsumeOrderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubscriptionOrderManagerGetConsumeOrderParams() *SubscriptionOrderManagerGetConsumeOrderParams {
	return &SubscriptionOrderManagerGetConsumeOrderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriptionOrderManagerGetConsumeOrderParamsWithTimeout creates a new SubscriptionOrderManagerGetConsumeOrderParams object
// with the ability to set a timeout on a request.
func NewSubscriptionOrderManagerGetConsumeOrderParamsWithTimeout(timeout time.Duration) *SubscriptionOrderManagerGetConsumeOrderParams {
	return &SubscriptionOrderManagerGetConsumeOrderParams{
		timeout: timeout,
	}
}

// NewSubscriptionOrderManagerGetConsumeOrderParamsWithContext creates a new SubscriptionOrderManagerGetConsumeOrderParams object
// with the ability to set a context for a request.
func NewSubscriptionOrderManagerGetConsumeOrderParamsWithContext(ctx context.Context) *SubscriptionOrderManagerGetConsumeOrderParams {
	return &SubscriptionOrderManagerGetConsumeOrderParams{
		Context: ctx,
	}
}

// NewSubscriptionOrderManagerGetConsumeOrderParamsWithHTTPClient creates a new SubscriptionOrderManagerGetConsumeOrderParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubscriptionOrderManagerGetConsumeOrderParamsWithHTTPClient(client *http.Client) *SubscriptionOrderManagerGetConsumeOrderParams {
	return &SubscriptionOrderManagerGetConsumeOrderParams{
		HTTPClient: client,
	}
}

/*
SubscriptionOrderManagerGetConsumeOrderParams contains all the parameters to send to the API endpoint

	for the subscription order manager get consume order operation.

	Typically these are written to a http.Request.
*/
type SubscriptionOrderManagerGetConsumeOrderParams struct {

	// ProdInstConsumeOrderID.
	ProdInstConsumeOrderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subscription order manager get consume order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionOrderManagerGetConsumeOrderParams) WithDefaults() *SubscriptionOrderManagerGetConsumeOrderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subscription order manager get consume order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionOrderManagerGetConsumeOrderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subscription order manager get consume order params
func (o *SubscriptionOrderManagerGetConsumeOrderParams) WithTimeout(timeout time.Duration) *SubscriptionOrderManagerGetConsumeOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscription order manager get consume order params
func (o *SubscriptionOrderManagerGetConsumeOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscription order manager get consume order params
func (o *SubscriptionOrderManagerGetConsumeOrderParams) WithContext(ctx context.Context) *SubscriptionOrderManagerGetConsumeOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscription order manager get consume order params
func (o *SubscriptionOrderManagerGetConsumeOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscription order manager get consume order params
func (o *SubscriptionOrderManagerGetConsumeOrderParams) WithHTTPClient(client *http.Client) *SubscriptionOrderManagerGetConsumeOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscription order manager get consume order params
func (o *SubscriptionOrderManagerGetConsumeOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProdInstConsumeOrderID adds the prodInstConsumeOrderID to the subscription order manager get consume order params
func (o *SubscriptionOrderManagerGetConsumeOrderParams) WithProdInstConsumeOrderID(prodInstConsumeOrderID string) *SubscriptionOrderManagerGetConsumeOrderParams {
	o.SetProdInstConsumeOrderID(prodInstConsumeOrderID)
	return o
}

// SetProdInstConsumeOrderID adds the prodInstConsumeOrderId to the subscription order manager get consume order params
func (o *SubscriptionOrderManagerGetConsumeOrderParams) SetProdInstConsumeOrderID(prodInstConsumeOrderID string) {
	o.ProdInstConsumeOrderID = prodInstConsumeOrderID
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriptionOrderManagerGetConsumeOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param prod_inst_consume_order_id
	if err := r.SetPathParam("prod_inst_consume_order_id", o.ProdInstConsumeOrderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}