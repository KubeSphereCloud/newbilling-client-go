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

// NewPricingManagerModifyFilterNameParams creates a new PricingManagerModifyFilterNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPricingManagerModifyFilterNameParams() *PricingManagerModifyFilterNameParams {
	return &PricingManagerModifyFilterNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPricingManagerModifyFilterNameParamsWithTimeout creates a new PricingManagerModifyFilterNameParams object
// with the ability to set a timeout on a request.
func NewPricingManagerModifyFilterNameParamsWithTimeout(timeout time.Duration) *PricingManagerModifyFilterNameParams {
	return &PricingManagerModifyFilterNameParams{
		timeout: timeout,
	}
}

// NewPricingManagerModifyFilterNameParamsWithContext creates a new PricingManagerModifyFilterNameParams object
// with the ability to set a context for a request.
func NewPricingManagerModifyFilterNameParamsWithContext(ctx context.Context) *PricingManagerModifyFilterNameParams {
	return &PricingManagerModifyFilterNameParams{
		Context: ctx,
	}
}

// NewPricingManagerModifyFilterNameParamsWithHTTPClient creates a new PricingManagerModifyFilterNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPricingManagerModifyFilterNameParamsWithHTTPClient(client *http.Client) *PricingManagerModifyFilterNameParams {
	return &PricingManagerModifyFilterNameParams{
		HTTPClient: client,
	}
}

/*
PricingManagerModifyFilterNameParams contains all the parameters to send to the API endpoint

	for the pricing manager modify filter name operation.

	Typically these are written to a http.Request.
*/
type PricingManagerModifyFilterNameParams struct {

	// Body.
	Body *models.NewbillingModifyFilterNameRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pricing manager modify filter name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerModifyFilterNameParams) WithDefaults() *PricingManagerModifyFilterNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pricing manager modify filter name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerModifyFilterNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pricing manager modify filter name params
func (o *PricingManagerModifyFilterNameParams) WithTimeout(timeout time.Duration) *PricingManagerModifyFilterNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pricing manager modify filter name params
func (o *PricingManagerModifyFilterNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pricing manager modify filter name params
func (o *PricingManagerModifyFilterNameParams) WithContext(ctx context.Context) *PricingManagerModifyFilterNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pricing manager modify filter name params
func (o *PricingManagerModifyFilterNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pricing manager modify filter name params
func (o *PricingManagerModifyFilterNameParams) WithHTTPClient(client *http.Client) *PricingManagerModifyFilterNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pricing manager modify filter name params
func (o *PricingManagerModifyFilterNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pricing manager modify filter name params
func (o *PricingManagerModifyFilterNameParams) WithBody(body *models.NewbillingModifyFilterNameRequest) *PricingManagerModifyFilterNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pricing manager modify filter name params
func (o *PricingManagerModifyFilterNameParams) SetBody(body *models.NewbillingModifyFilterNameRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PricingManagerModifyFilterNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
