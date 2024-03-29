// Code generated by go-swagger; DO NOT EDIT.

package subscription_package_order_manager

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

// NewSubscriptionPackageOrderManagerCancelPrdOrderParams creates a new SubscriptionPackageOrderManagerCancelPrdOrderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubscriptionPackageOrderManagerCancelPrdOrderParams() *SubscriptionPackageOrderManagerCancelPrdOrderParams {
	return &SubscriptionPackageOrderManagerCancelPrdOrderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriptionPackageOrderManagerCancelPrdOrderParamsWithTimeout creates a new SubscriptionPackageOrderManagerCancelPrdOrderParams object
// with the ability to set a timeout on a request.
func NewSubscriptionPackageOrderManagerCancelPrdOrderParamsWithTimeout(timeout time.Duration) *SubscriptionPackageOrderManagerCancelPrdOrderParams {
	return &SubscriptionPackageOrderManagerCancelPrdOrderParams{
		timeout: timeout,
	}
}

// NewSubscriptionPackageOrderManagerCancelPrdOrderParamsWithContext creates a new SubscriptionPackageOrderManagerCancelPrdOrderParams object
// with the ability to set a context for a request.
func NewSubscriptionPackageOrderManagerCancelPrdOrderParamsWithContext(ctx context.Context) *SubscriptionPackageOrderManagerCancelPrdOrderParams {
	return &SubscriptionPackageOrderManagerCancelPrdOrderParams{
		Context: ctx,
	}
}

// NewSubscriptionPackageOrderManagerCancelPrdOrderParamsWithHTTPClient creates a new SubscriptionPackageOrderManagerCancelPrdOrderParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubscriptionPackageOrderManagerCancelPrdOrderParamsWithHTTPClient(client *http.Client) *SubscriptionPackageOrderManagerCancelPrdOrderParams {
	return &SubscriptionPackageOrderManagerCancelPrdOrderParams{
		HTTPClient: client,
	}
}

/*
SubscriptionPackageOrderManagerCancelPrdOrderParams contains all the parameters to send to the API endpoint

	for the subscription package order manager cancel prd order operation.

	Typically these are written to a http.Request.
*/
type SubscriptionPackageOrderManagerCancelPrdOrderParams struct {

	// Body.
	Body *models.NewbillingCancelPrdOrderRequest

	/* OrderID.

	   主订单ID
	*/
	OrderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subscription package order manager cancel prd order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionPackageOrderManagerCancelPrdOrderParams) WithDefaults() *SubscriptionPackageOrderManagerCancelPrdOrderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subscription package order manager cancel prd order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionPackageOrderManagerCancelPrdOrderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subscription package order manager cancel prd order params
func (o *SubscriptionPackageOrderManagerCancelPrdOrderParams) WithTimeout(timeout time.Duration) *SubscriptionPackageOrderManagerCancelPrdOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscription package order manager cancel prd order params
func (o *SubscriptionPackageOrderManagerCancelPrdOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscription package order manager cancel prd order params
func (o *SubscriptionPackageOrderManagerCancelPrdOrderParams) WithContext(ctx context.Context) *SubscriptionPackageOrderManagerCancelPrdOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscription package order manager cancel prd order params
func (o *SubscriptionPackageOrderManagerCancelPrdOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscription package order manager cancel prd order params
func (o *SubscriptionPackageOrderManagerCancelPrdOrderParams) WithHTTPClient(client *http.Client) *SubscriptionPackageOrderManagerCancelPrdOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscription package order manager cancel prd order params
func (o *SubscriptionPackageOrderManagerCancelPrdOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the subscription package order manager cancel prd order params
func (o *SubscriptionPackageOrderManagerCancelPrdOrderParams) WithBody(body *models.NewbillingCancelPrdOrderRequest) *SubscriptionPackageOrderManagerCancelPrdOrderParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the subscription package order manager cancel prd order params
func (o *SubscriptionPackageOrderManagerCancelPrdOrderParams) SetBody(body *models.NewbillingCancelPrdOrderRequest) {
	o.Body = body
}

// WithOrderID adds the orderID to the subscription package order manager cancel prd order params
func (o *SubscriptionPackageOrderManagerCancelPrdOrderParams) WithOrderID(orderID string) *SubscriptionPackageOrderManagerCancelPrdOrderParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the subscription package order manager cancel prd order params
func (o *SubscriptionPackageOrderManagerCancelPrdOrderParams) SetOrderID(orderID string) {
	o.OrderID = orderID
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriptionPackageOrderManagerCancelPrdOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param order_id
	if err := r.SetPathParam("order_id", o.OrderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
