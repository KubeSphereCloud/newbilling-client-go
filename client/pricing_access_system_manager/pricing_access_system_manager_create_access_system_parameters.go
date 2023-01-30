// Code generated by go-swagger; DO NOT EDIT.

package pricing_access_system_manager

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

// NewPricingAccessSystemManagerCreateAccessSystemParams creates a new PricingAccessSystemManagerCreateAccessSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPricingAccessSystemManagerCreateAccessSystemParams() *PricingAccessSystemManagerCreateAccessSystemParams {
	return &PricingAccessSystemManagerCreateAccessSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPricingAccessSystemManagerCreateAccessSystemParamsWithTimeout creates a new PricingAccessSystemManagerCreateAccessSystemParams object
// with the ability to set a timeout on a request.
func NewPricingAccessSystemManagerCreateAccessSystemParamsWithTimeout(timeout time.Duration) *PricingAccessSystemManagerCreateAccessSystemParams {
	return &PricingAccessSystemManagerCreateAccessSystemParams{
		timeout: timeout,
	}
}

// NewPricingAccessSystemManagerCreateAccessSystemParamsWithContext creates a new PricingAccessSystemManagerCreateAccessSystemParams object
// with the ability to set a context for a request.
func NewPricingAccessSystemManagerCreateAccessSystemParamsWithContext(ctx context.Context) *PricingAccessSystemManagerCreateAccessSystemParams {
	return &PricingAccessSystemManagerCreateAccessSystemParams{
		Context: ctx,
	}
}

// NewPricingAccessSystemManagerCreateAccessSystemParamsWithHTTPClient creates a new PricingAccessSystemManagerCreateAccessSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewPricingAccessSystemManagerCreateAccessSystemParamsWithHTTPClient(client *http.Client) *PricingAccessSystemManagerCreateAccessSystemParams {
	return &PricingAccessSystemManagerCreateAccessSystemParams{
		HTTPClient: client,
	}
}

/*
PricingAccessSystemManagerCreateAccessSystemParams contains all the parameters to send to the API endpoint

	for the pricing access system manager create access system operation.

	Typically these are written to a http.Request.
*/
type PricingAccessSystemManagerCreateAccessSystemParams struct {

	// Body.
	Body *models.NewbillingCreateAccessSystemRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pricing access system manager create access system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingAccessSystemManagerCreateAccessSystemParams) WithDefaults() *PricingAccessSystemManagerCreateAccessSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pricing access system manager create access system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingAccessSystemManagerCreateAccessSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pricing access system manager create access system params
func (o *PricingAccessSystemManagerCreateAccessSystemParams) WithTimeout(timeout time.Duration) *PricingAccessSystemManagerCreateAccessSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pricing access system manager create access system params
func (o *PricingAccessSystemManagerCreateAccessSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pricing access system manager create access system params
func (o *PricingAccessSystemManagerCreateAccessSystemParams) WithContext(ctx context.Context) *PricingAccessSystemManagerCreateAccessSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pricing access system manager create access system params
func (o *PricingAccessSystemManagerCreateAccessSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pricing access system manager create access system params
func (o *PricingAccessSystemManagerCreateAccessSystemParams) WithHTTPClient(client *http.Client) *PricingAccessSystemManagerCreateAccessSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pricing access system manager create access system params
func (o *PricingAccessSystemManagerCreateAccessSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pricing access system manager create access system params
func (o *PricingAccessSystemManagerCreateAccessSystemParams) WithBody(body *models.NewbillingCreateAccessSystemRequest) *PricingAccessSystemManagerCreateAccessSystemParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pricing access system manager create access system params
func (o *PricingAccessSystemManagerCreateAccessSystemParams) SetBody(body *models.NewbillingCreateAccessSystemRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PricingAccessSystemManagerCreateAccessSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
