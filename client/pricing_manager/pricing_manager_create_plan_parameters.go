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

// NewPricingManagerCreatePlanParams creates a new PricingManagerCreatePlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPricingManagerCreatePlanParams() *PricingManagerCreatePlanParams {
	return &PricingManagerCreatePlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPricingManagerCreatePlanParamsWithTimeout creates a new PricingManagerCreatePlanParams object
// with the ability to set a timeout on a request.
func NewPricingManagerCreatePlanParamsWithTimeout(timeout time.Duration) *PricingManagerCreatePlanParams {
	return &PricingManagerCreatePlanParams{
		timeout: timeout,
	}
}

// NewPricingManagerCreatePlanParamsWithContext creates a new PricingManagerCreatePlanParams object
// with the ability to set a context for a request.
func NewPricingManagerCreatePlanParamsWithContext(ctx context.Context) *PricingManagerCreatePlanParams {
	return &PricingManagerCreatePlanParams{
		Context: ctx,
	}
}

// NewPricingManagerCreatePlanParamsWithHTTPClient creates a new PricingManagerCreatePlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewPricingManagerCreatePlanParamsWithHTTPClient(client *http.Client) *PricingManagerCreatePlanParams {
	return &PricingManagerCreatePlanParams{
		HTTPClient: client,
	}
}

/*
PricingManagerCreatePlanParams contains all the parameters to send to the API endpoint

	for the pricing manager create plan operation.

	Typically these are written to a http.Request.
*/
type PricingManagerCreatePlanParams struct {

	// Body.
	Body *models.NewbillingCreatePlanRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pricing manager create plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerCreatePlanParams) WithDefaults() *PricingManagerCreatePlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pricing manager create plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerCreatePlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pricing manager create plan params
func (o *PricingManagerCreatePlanParams) WithTimeout(timeout time.Duration) *PricingManagerCreatePlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pricing manager create plan params
func (o *PricingManagerCreatePlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pricing manager create plan params
func (o *PricingManagerCreatePlanParams) WithContext(ctx context.Context) *PricingManagerCreatePlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pricing manager create plan params
func (o *PricingManagerCreatePlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pricing manager create plan params
func (o *PricingManagerCreatePlanParams) WithHTTPClient(client *http.Client) *PricingManagerCreatePlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pricing manager create plan params
func (o *PricingManagerCreatePlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pricing manager create plan params
func (o *PricingManagerCreatePlanParams) WithBody(body *models.NewbillingCreatePlanRequest) *PricingManagerCreatePlanParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pricing manager create plan params
func (o *PricingManagerCreatePlanParams) SetBody(body *models.NewbillingCreatePlanRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PricingManagerCreatePlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
