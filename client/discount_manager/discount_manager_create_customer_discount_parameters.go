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

// NewDiscountManagerCreateCustomerDiscountParams creates a new DiscountManagerCreateCustomerDiscountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDiscountManagerCreateCustomerDiscountParams() *DiscountManagerCreateCustomerDiscountParams {
	return &DiscountManagerCreateCustomerDiscountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDiscountManagerCreateCustomerDiscountParamsWithTimeout creates a new DiscountManagerCreateCustomerDiscountParams object
// with the ability to set a timeout on a request.
func NewDiscountManagerCreateCustomerDiscountParamsWithTimeout(timeout time.Duration) *DiscountManagerCreateCustomerDiscountParams {
	return &DiscountManagerCreateCustomerDiscountParams{
		timeout: timeout,
	}
}

// NewDiscountManagerCreateCustomerDiscountParamsWithContext creates a new DiscountManagerCreateCustomerDiscountParams object
// with the ability to set a context for a request.
func NewDiscountManagerCreateCustomerDiscountParamsWithContext(ctx context.Context) *DiscountManagerCreateCustomerDiscountParams {
	return &DiscountManagerCreateCustomerDiscountParams{
		Context: ctx,
	}
}

// NewDiscountManagerCreateCustomerDiscountParamsWithHTTPClient creates a new DiscountManagerCreateCustomerDiscountParams object
// with the ability to set a custom HTTPClient for a request.
func NewDiscountManagerCreateCustomerDiscountParamsWithHTTPClient(client *http.Client) *DiscountManagerCreateCustomerDiscountParams {
	return &DiscountManagerCreateCustomerDiscountParams{
		HTTPClient: client,
	}
}

/*
DiscountManagerCreateCustomerDiscountParams contains all the parameters to send to the API endpoint

	for the discount manager create customer discount operation.

	Typically these are written to a http.Request.
*/
type DiscountManagerCreateCustomerDiscountParams struct {

	// Body.
	Body *models.NewbillingCreateCustomerDiscountRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the discount manager create customer discount params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiscountManagerCreateCustomerDiscountParams) WithDefaults() *DiscountManagerCreateCustomerDiscountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the discount manager create customer discount params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiscountManagerCreateCustomerDiscountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the discount manager create customer discount params
func (o *DiscountManagerCreateCustomerDiscountParams) WithTimeout(timeout time.Duration) *DiscountManagerCreateCustomerDiscountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the discount manager create customer discount params
func (o *DiscountManagerCreateCustomerDiscountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the discount manager create customer discount params
func (o *DiscountManagerCreateCustomerDiscountParams) WithContext(ctx context.Context) *DiscountManagerCreateCustomerDiscountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the discount manager create customer discount params
func (o *DiscountManagerCreateCustomerDiscountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the discount manager create customer discount params
func (o *DiscountManagerCreateCustomerDiscountParams) WithHTTPClient(client *http.Client) *DiscountManagerCreateCustomerDiscountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the discount manager create customer discount params
func (o *DiscountManagerCreateCustomerDiscountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the discount manager create customer discount params
func (o *DiscountManagerCreateCustomerDiscountParams) WithBody(body *models.NewbillingCreateCustomerDiscountRequest) *DiscountManagerCreateCustomerDiscountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the discount manager create customer discount params
func (o *DiscountManagerCreateCustomerDiscountParams) SetBody(body *models.NewbillingCreateCustomerDiscountRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DiscountManagerCreateCustomerDiscountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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