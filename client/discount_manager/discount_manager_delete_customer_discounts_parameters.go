// Code generated by go-swagger; DO NOT EDIT.

package discount_manager

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

// NewDiscountManagerDeleteCustomerDiscountsParams creates a new DiscountManagerDeleteCustomerDiscountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDiscountManagerDeleteCustomerDiscountsParams() *DiscountManagerDeleteCustomerDiscountsParams {
	return &DiscountManagerDeleteCustomerDiscountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDiscountManagerDeleteCustomerDiscountsParamsWithTimeout creates a new DiscountManagerDeleteCustomerDiscountsParams object
// with the ability to set a timeout on a request.
func NewDiscountManagerDeleteCustomerDiscountsParamsWithTimeout(timeout time.Duration) *DiscountManagerDeleteCustomerDiscountsParams {
	return &DiscountManagerDeleteCustomerDiscountsParams{
		timeout: timeout,
	}
}

// NewDiscountManagerDeleteCustomerDiscountsParamsWithContext creates a new DiscountManagerDeleteCustomerDiscountsParams object
// with the ability to set a context for a request.
func NewDiscountManagerDeleteCustomerDiscountsParamsWithContext(ctx context.Context) *DiscountManagerDeleteCustomerDiscountsParams {
	return &DiscountManagerDeleteCustomerDiscountsParams{
		Context: ctx,
	}
}

// NewDiscountManagerDeleteCustomerDiscountsParamsWithHTTPClient creates a new DiscountManagerDeleteCustomerDiscountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDiscountManagerDeleteCustomerDiscountsParamsWithHTTPClient(client *http.Client) *DiscountManagerDeleteCustomerDiscountsParams {
	return &DiscountManagerDeleteCustomerDiscountsParams{
		HTTPClient: client,
	}
}

/*
DiscountManagerDeleteCustomerDiscountsParams contains all the parameters to send to the API endpoint

	for the discount manager delete customer discounts operation.

	Typically these are written to a http.Request.
*/
type DiscountManagerDeleteCustomerDiscountsParams struct {

	// Body.
	Body *models.NewbillingDeleteCustomerDiscountsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the discount manager delete customer discounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiscountManagerDeleteCustomerDiscountsParams) WithDefaults() *DiscountManagerDeleteCustomerDiscountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the discount manager delete customer discounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiscountManagerDeleteCustomerDiscountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the discount manager delete customer discounts params
func (o *DiscountManagerDeleteCustomerDiscountsParams) WithTimeout(timeout time.Duration) *DiscountManagerDeleteCustomerDiscountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the discount manager delete customer discounts params
func (o *DiscountManagerDeleteCustomerDiscountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the discount manager delete customer discounts params
func (o *DiscountManagerDeleteCustomerDiscountsParams) WithContext(ctx context.Context) *DiscountManagerDeleteCustomerDiscountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the discount manager delete customer discounts params
func (o *DiscountManagerDeleteCustomerDiscountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the discount manager delete customer discounts params
func (o *DiscountManagerDeleteCustomerDiscountsParams) WithHTTPClient(client *http.Client) *DiscountManagerDeleteCustomerDiscountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the discount manager delete customer discounts params
func (o *DiscountManagerDeleteCustomerDiscountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the discount manager delete customer discounts params
func (o *DiscountManagerDeleteCustomerDiscountsParams) WithBody(body *models.NewbillingDeleteCustomerDiscountsRequest) *DiscountManagerDeleteCustomerDiscountsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the discount manager delete customer discounts params
func (o *DiscountManagerDeleteCustomerDiscountsParams) SetBody(body *models.NewbillingDeleteCustomerDiscountsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DiscountManagerDeleteCustomerDiscountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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