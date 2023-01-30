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

// NewPricingManagerModifyStrategyParams creates a new PricingManagerModifyStrategyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPricingManagerModifyStrategyParams() *PricingManagerModifyStrategyParams {
	return &PricingManagerModifyStrategyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPricingManagerModifyStrategyParamsWithTimeout creates a new PricingManagerModifyStrategyParams object
// with the ability to set a timeout on a request.
func NewPricingManagerModifyStrategyParamsWithTimeout(timeout time.Duration) *PricingManagerModifyStrategyParams {
	return &PricingManagerModifyStrategyParams{
		timeout: timeout,
	}
}

// NewPricingManagerModifyStrategyParamsWithContext creates a new PricingManagerModifyStrategyParams object
// with the ability to set a context for a request.
func NewPricingManagerModifyStrategyParamsWithContext(ctx context.Context) *PricingManagerModifyStrategyParams {
	return &PricingManagerModifyStrategyParams{
		Context: ctx,
	}
}

// NewPricingManagerModifyStrategyParamsWithHTTPClient creates a new PricingManagerModifyStrategyParams object
// with the ability to set a custom HTTPClient for a request.
func NewPricingManagerModifyStrategyParamsWithHTTPClient(client *http.Client) *PricingManagerModifyStrategyParams {
	return &PricingManagerModifyStrategyParams{
		HTTPClient: client,
	}
}

/*
PricingManagerModifyStrategyParams contains all the parameters to send to the API endpoint

	for the pricing manager modify strategy operation.

	Typically these are written to a http.Request.
*/
type PricingManagerModifyStrategyParams struct {

	// Body.
	Body *models.NewbillingModifyStrategyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pricing manager modify strategy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerModifyStrategyParams) WithDefaults() *PricingManagerModifyStrategyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pricing manager modify strategy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerModifyStrategyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pricing manager modify strategy params
func (o *PricingManagerModifyStrategyParams) WithTimeout(timeout time.Duration) *PricingManagerModifyStrategyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pricing manager modify strategy params
func (o *PricingManagerModifyStrategyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pricing manager modify strategy params
func (o *PricingManagerModifyStrategyParams) WithContext(ctx context.Context) *PricingManagerModifyStrategyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pricing manager modify strategy params
func (o *PricingManagerModifyStrategyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pricing manager modify strategy params
func (o *PricingManagerModifyStrategyParams) WithHTTPClient(client *http.Client) *PricingManagerModifyStrategyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pricing manager modify strategy params
func (o *PricingManagerModifyStrategyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pricing manager modify strategy params
func (o *PricingManagerModifyStrategyParams) WithBody(body *models.NewbillingModifyStrategyRequest) *PricingManagerModifyStrategyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pricing manager modify strategy params
func (o *PricingManagerModifyStrategyParams) SetBody(body *models.NewbillingModifyStrategyRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PricingManagerModifyStrategyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
