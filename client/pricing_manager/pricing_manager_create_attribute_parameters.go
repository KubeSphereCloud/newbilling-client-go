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

// NewPricingManagerCreateAttributeParams creates a new PricingManagerCreateAttributeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPricingManagerCreateAttributeParams() *PricingManagerCreateAttributeParams {
	return &PricingManagerCreateAttributeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPricingManagerCreateAttributeParamsWithTimeout creates a new PricingManagerCreateAttributeParams object
// with the ability to set a timeout on a request.
func NewPricingManagerCreateAttributeParamsWithTimeout(timeout time.Duration) *PricingManagerCreateAttributeParams {
	return &PricingManagerCreateAttributeParams{
		timeout: timeout,
	}
}

// NewPricingManagerCreateAttributeParamsWithContext creates a new PricingManagerCreateAttributeParams object
// with the ability to set a context for a request.
func NewPricingManagerCreateAttributeParamsWithContext(ctx context.Context) *PricingManagerCreateAttributeParams {
	return &PricingManagerCreateAttributeParams{
		Context: ctx,
	}
}

// NewPricingManagerCreateAttributeParamsWithHTTPClient creates a new PricingManagerCreateAttributeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPricingManagerCreateAttributeParamsWithHTTPClient(client *http.Client) *PricingManagerCreateAttributeParams {
	return &PricingManagerCreateAttributeParams{
		HTTPClient: client,
	}
}

/*
PricingManagerCreateAttributeParams contains all the parameters to send to the API endpoint

	for the pricing manager create attribute operation.

	Typically these are written to a http.Request.
*/
type PricingManagerCreateAttributeParams struct {

	// Body.
	Body *models.NewbillingCreateAttributeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pricing manager create attribute params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerCreateAttributeParams) WithDefaults() *PricingManagerCreateAttributeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pricing manager create attribute params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerCreateAttributeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pricing manager create attribute params
func (o *PricingManagerCreateAttributeParams) WithTimeout(timeout time.Duration) *PricingManagerCreateAttributeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pricing manager create attribute params
func (o *PricingManagerCreateAttributeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pricing manager create attribute params
func (o *PricingManagerCreateAttributeParams) WithContext(ctx context.Context) *PricingManagerCreateAttributeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pricing manager create attribute params
func (o *PricingManagerCreateAttributeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pricing manager create attribute params
func (o *PricingManagerCreateAttributeParams) WithHTTPClient(client *http.Client) *PricingManagerCreateAttributeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pricing manager create attribute params
func (o *PricingManagerCreateAttributeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pricing manager create attribute params
func (o *PricingManagerCreateAttributeParams) WithBody(body *models.NewbillingCreateAttributeRequest) *PricingManagerCreateAttributeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pricing manager create attribute params
func (o *PricingManagerCreateAttributeParams) SetBody(body *models.NewbillingCreateAttributeRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PricingManagerCreateAttributeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
