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

// NewPricingAccessSystemManagerModifyAccessSystemParams creates a new PricingAccessSystemManagerModifyAccessSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPricingAccessSystemManagerModifyAccessSystemParams() *PricingAccessSystemManagerModifyAccessSystemParams {
	return &PricingAccessSystemManagerModifyAccessSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPricingAccessSystemManagerModifyAccessSystemParamsWithTimeout creates a new PricingAccessSystemManagerModifyAccessSystemParams object
// with the ability to set a timeout on a request.
func NewPricingAccessSystemManagerModifyAccessSystemParamsWithTimeout(timeout time.Duration) *PricingAccessSystemManagerModifyAccessSystemParams {
	return &PricingAccessSystemManagerModifyAccessSystemParams{
		timeout: timeout,
	}
}

// NewPricingAccessSystemManagerModifyAccessSystemParamsWithContext creates a new PricingAccessSystemManagerModifyAccessSystemParams object
// with the ability to set a context for a request.
func NewPricingAccessSystemManagerModifyAccessSystemParamsWithContext(ctx context.Context) *PricingAccessSystemManagerModifyAccessSystemParams {
	return &PricingAccessSystemManagerModifyAccessSystemParams{
		Context: ctx,
	}
}

// NewPricingAccessSystemManagerModifyAccessSystemParamsWithHTTPClient creates a new PricingAccessSystemManagerModifyAccessSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewPricingAccessSystemManagerModifyAccessSystemParamsWithHTTPClient(client *http.Client) *PricingAccessSystemManagerModifyAccessSystemParams {
	return &PricingAccessSystemManagerModifyAccessSystemParams{
		HTTPClient: client,
	}
}

/*
PricingAccessSystemManagerModifyAccessSystemParams contains all the parameters to send to the API endpoint

	for the pricing access system manager modify access system operation.

	Typically these are written to a http.Request.
*/
type PricingAccessSystemManagerModifyAccessSystemParams struct {

	// Body.
	Body *models.NewbillingModifyAccessSystemRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pricing access system manager modify access system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingAccessSystemManagerModifyAccessSystemParams) WithDefaults() *PricingAccessSystemManagerModifyAccessSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pricing access system manager modify access system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingAccessSystemManagerModifyAccessSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pricing access system manager modify access system params
func (o *PricingAccessSystemManagerModifyAccessSystemParams) WithTimeout(timeout time.Duration) *PricingAccessSystemManagerModifyAccessSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pricing access system manager modify access system params
func (o *PricingAccessSystemManagerModifyAccessSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pricing access system manager modify access system params
func (o *PricingAccessSystemManagerModifyAccessSystemParams) WithContext(ctx context.Context) *PricingAccessSystemManagerModifyAccessSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pricing access system manager modify access system params
func (o *PricingAccessSystemManagerModifyAccessSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pricing access system manager modify access system params
func (o *PricingAccessSystemManagerModifyAccessSystemParams) WithHTTPClient(client *http.Client) *PricingAccessSystemManagerModifyAccessSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pricing access system manager modify access system params
func (o *PricingAccessSystemManagerModifyAccessSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pricing access system manager modify access system params
func (o *PricingAccessSystemManagerModifyAccessSystemParams) WithBody(body *models.NewbillingModifyAccessSystemRequest) *PricingAccessSystemManagerModifyAccessSystemParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pricing access system manager modify access system params
func (o *PricingAccessSystemManagerModifyAccessSystemParams) SetBody(body *models.NewbillingModifyAccessSystemRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PricingAccessSystemManagerModifyAccessSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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