// Code generated by go-swagger; DO NOT EDIT.

package pricing_manager

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

// NewPricingManagerDeleteProductsParams creates a new PricingManagerDeleteProductsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPricingManagerDeleteProductsParams() *PricingManagerDeleteProductsParams {
	return &PricingManagerDeleteProductsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPricingManagerDeleteProductsParamsWithTimeout creates a new PricingManagerDeleteProductsParams object
// with the ability to set a timeout on a request.
func NewPricingManagerDeleteProductsParamsWithTimeout(timeout time.Duration) *PricingManagerDeleteProductsParams {
	return &PricingManagerDeleteProductsParams{
		timeout: timeout,
	}
}

// NewPricingManagerDeleteProductsParamsWithContext creates a new PricingManagerDeleteProductsParams object
// with the ability to set a context for a request.
func NewPricingManagerDeleteProductsParamsWithContext(ctx context.Context) *PricingManagerDeleteProductsParams {
	return &PricingManagerDeleteProductsParams{
		Context: ctx,
	}
}

// NewPricingManagerDeleteProductsParamsWithHTTPClient creates a new PricingManagerDeleteProductsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPricingManagerDeleteProductsParamsWithHTTPClient(client *http.Client) *PricingManagerDeleteProductsParams {
	return &PricingManagerDeleteProductsParams{
		HTTPClient: client,
	}
}

/*
PricingManagerDeleteProductsParams contains all the parameters to send to the API endpoint

	for the pricing manager delete products operation.

	Typically these are written to a http.Request.
*/
type PricingManagerDeleteProductsParams struct {

	// Body.
	Body *models.NewbillingDeleteProductsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pricing manager delete products params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerDeleteProductsParams) WithDefaults() *PricingManagerDeleteProductsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pricing manager delete products params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerDeleteProductsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pricing manager delete products params
func (o *PricingManagerDeleteProductsParams) WithTimeout(timeout time.Duration) *PricingManagerDeleteProductsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pricing manager delete products params
func (o *PricingManagerDeleteProductsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pricing manager delete products params
func (o *PricingManagerDeleteProductsParams) WithContext(ctx context.Context) *PricingManagerDeleteProductsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pricing manager delete products params
func (o *PricingManagerDeleteProductsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pricing manager delete products params
func (o *PricingManagerDeleteProductsParams) WithHTTPClient(client *http.Client) *PricingManagerDeleteProductsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pricing manager delete products params
func (o *PricingManagerDeleteProductsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pricing manager delete products params
func (o *PricingManagerDeleteProductsParams) WithBody(body *models.NewbillingDeleteProductsRequest) *PricingManagerDeleteProductsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pricing manager delete products params
func (o *PricingManagerDeleteProductsParams) SetBody(body *models.NewbillingDeleteProductsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PricingManagerDeleteProductsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
