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
)

// NewPricingManagerGetPricingParams creates a new PricingManagerGetPricingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPricingManagerGetPricingParams() *PricingManagerGetPricingParams {
	return &PricingManagerGetPricingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPricingManagerGetPricingParamsWithTimeout creates a new PricingManagerGetPricingParams object
// with the ability to set a timeout on a request.
func NewPricingManagerGetPricingParamsWithTimeout(timeout time.Duration) *PricingManagerGetPricingParams {
	return &PricingManagerGetPricingParams{
		timeout: timeout,
	}
}

// NewPricingManagerGetPricingParamsWithContext creates a new PricingManagerGetPricingParams object
// with the ability to set a context for a request.
func NewPricingManagerGetPricingParamsWithContext(ctx context.Context) *PricingManagerGetPricingParams {
	return &PricingManagerGetPricingParams{
		Context: ctx,
	}
}

// NewPricingManagerGetPricingParamsWithHTTPClient creates a new PricingManagerGetPricingParams object
// with the ability to set a custom HTTPClient for a request.
func NewPricingManagerGetPricingParamsWithHTTPClient(client *http.Client) *PricingManagerGetPricingParams {
	return &PricingManagerGetPricingParams{
		HTTPClient: client,
	}
}

/*
PricingManagerGetPricingParams contains all the parameters to send to the API endpoint

	for the pricing manager get pricing operation.

	Typically these are written to a http.Request.
*/
type PricingManagerGetPricingParams struct {

	// CataID.
	CataID *string

	// PlanID.
	PlanID *string

	// ProdID.
	ProdID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pricing manager get pricing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerGetPricingParams) WithDefaults() *PricingManagerGetPricingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pricing manager get pricing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerGetPricingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pricing manager get pricing params
func (o *PricingManagerGetPricingParams) WithTimeout(timeout time.Duration) *PricingManagerGetPricingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pricing manager get pricing params
func (o *PricingManagerGetPricingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pricing manager get pricing params
func (o *PricingManagerGetPricingParams) WithContext(ctx context.Context) *PricingManagerGetPricingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pricing manager get pricing params
func (o *PricingManagerGetPricingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pricing manager get pricing params
func (o *PricingManagerGetPricingParams) WithHTTPClient(client *http.Client) *PricingManagerGetPricingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pricing manager get pricing params
func (o *PricingManagerGetPricingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCataID adds the cataID to the pricing manager get pricing params
func (o *PricingManagerGetPricingParams) WithCataID(cataID *string) *PricingManagerGetPricingParams {
	o.SetCataID(cataID)
	return o
}

// SetCataID adds the cataId to the pricing manager get pricing params
func (o *PricingManagerGetPricingParams) SetCataID(cataID *string) {
	o.CataID = cataID
}

// WithPlanID adds the planID to the pricing manager get pricing params
func (o *PricingManagerGetPricingParams) WithPlanID(planID *string) *PricingManagerGetPricingParams {
	o.SetPlanID(planID)
	return o
}

// SetPlanID adds the planId to the pricing manager get pricing params
func (o *PricingManagerGetPricingParams) SetPlanID(planID *string) {
	o.PlanID = planID
}

// WithProdID adds the prodID to the pricing manager get pricing params
func (o *PricingManagerGetPricingParams) WithProdID(prodID *string) *PricingManagerGetPricingParams {
	o.SetProdID(prodID)
	return o
}

// SetProdID adds the prodId to the pricing manager get pricing params
func (o *PricingManagerGetPricingParams) SetProdID(prodID *string) {
	o.ProdID = prodID
}

// WriteToRequest writes these params to a swagger request
func (o *PricingManagerGetPricingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CataID != nil {

		// query param cata_id
		var qrCataID string

		if o.CataID != nil {
			qrCataID = *o.CataID
		}
		qCataID := qrCataID
		if qCataID != "" {

			if err := r.SetQueryParam("cata_id", qCataID); err != nil {
				return err
			}
		}
	}

	if o.PlanID != nil {

		// query param plan_id
		var qrPlanID string

		if o.PlanID != nil {
			qrPlanID = *o.PlanID
		}
		qPlanID := qrPlanID
		if qPlanID != "" {

			if err := r.SetQueryParam("plan_id", qPlanID); err != nil {
				return err
			}
		}
	}

	if o.ProdID != nil {

		// query param prod_id
		var qrProdID string

		if o.ProdID != nil {
			qrProdID = *o.ProdID
		}
		qProdID := qrProdID
		if qProdID != "" {

			if err := r.SetQueryParam("prod_id", qProdID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
