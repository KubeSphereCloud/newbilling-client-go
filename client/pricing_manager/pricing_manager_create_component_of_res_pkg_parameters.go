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

// NewPricingManagerCreateComponentOfResPkgParams creates a new PricingManagerCreateComponentOfResPkgParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPricingManagerCreateComponentOfResPkgParams() *PricingManagerCreateComponentOfResPkgParams {
	return &PricingManagerCreateComponentOfResPkgParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPricingManagerCreateComponentOfResPkgParamsWithTimeout creates a new PricingManagerCreateComponentOfResPkgParams object
// with the ability to set a timeout on a request.
func NewPricingManagerCreateComponentOfResPkgParamsWithTimeout(timeout time.Duration) *PricingManagerCreateComponentOfResPkgParams {
	return &PricingManagerCreateComponentOfResPkgParams{
		timeout: timeout,
	}
}

// NewPricingManagerCreateComponentOfResPkgParamsWithContext creates a new PricingManagerCreateComponentOfResPkgParams object
// with the ability to set a context for a request.
func NewPricingManagerCreateComponentOfResPkgParamsWithContext(ctx context.Context) *PricingManagerCreateComponentOfResPkgParams {
	return &PricingManagerCreateComponentOfResPkgParams{
		Context: ctx,
	}
}

// NewPricingManagerCreateComponentOfResPkgParamsWithHTTPClient creates a new PricingManagerCreateComponentOfResPkgParams object
// with the ability to set a custom HTTPClient for a request.
func NewPricingManagerCreateComponentOfResPkgParamsWithHTTPClient(client *http.Client) *PricingManagerCreateComponentOfResPkgParams {
	return &PricingManagerCreateComponentOfResPkgParams{
		HTTPClient: client,
	}
}

/*
PricingManagerCreateComponentOfResPkgParams contains all the parameters to send to the API endpoint

	for the pricing manager create component of res pkg operation.

	Typically these are written to a http.Request.
*/
type PricingManagerCreateComponentOfResPkgParams struct {

	// Body.
	Body *models.NewbillingComponentOfPricingSimpleRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pricing manager create component of res pkg params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerCreateComponentOfResPkgParams) WithDefaults() *PricingManagerCreateComponentOfResPkgParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pricing manager create component of res pkg params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerCreateComponentOfResPkgParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pricing manager create component of res pkg params
func (o *PricingManagerCreateComponentOfResPkgParams) WithTimeout(timeout time.Duration) *PricingManagerCreateComponentOfResPkgParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pricing manager create component of res pkg params
func (o *PricingManagerCreateComponentOfResPkgParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pricing manager create component of res pkg params
func (o *PricingManagerCreateComponentOfResPkgParams) WithContext(ctx context.Context) *PricingManagerCreateComponentOfResPkgParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pricing manager create component of res pkg params
func (o *PricingManagerCreateComponentOfResPkgParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pricing manager create component of res pkg params
func (o *PricingManagerCreateComponentOfResPkgParams) WithHTTPClient(client *http.Client) *PricingManagerCreateComponentOfResPkgParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pricing manager create component of res pkg params
func (o *PricingManagerCreateComponentOfResPkgParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pricing manager create component of res pkg params
func (o *PricingManagerCreateComponentOfResPkgParams) WithBody(body *models.NewbillingComponentOfPricingSimpleRequest) *PricingManagerCreateComponentOfResPkgParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pricing manager create component of res pkg params
func (o *PricingManagerCreateComponentOfResPkgParams) SetBody(body *models.NewbillingComponentOfPricingSimpleRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PricingManagerCreateComponentOfResPkgParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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