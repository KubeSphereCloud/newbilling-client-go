// Code generated by go-swagger; DO NOT EDIT.

package charging_manager

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

// NewChargingManagerCreateChargeParams creates a new ChargingManagerCreateChargeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChargingManagerCreateChargeParams() *ChargingManagerCreateChargeParams {
	return &ChargingManagerCreateChargeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChargingManagerCreateChargeParamsWithTimeout creates a new ChargingManagerCreateChargeParams object
// with the ability to set a timeout on a request.
func NewChargingManagerCreateChargeParamsWithTimeout(timeout time.Duration) *ChargingManagerCreateChargeParams {
	return &ChargingManagerCreateChargeParams{
		timeout: timeout,
	}
}

// NewChargingManagerCreateChargeParamsWithContext creates a new ChargingManagerCreateChargeParams object
// with the ability to set a context for a request.
func NewChargingManagerCreateChargeParamsWithContext(ctx context.Context) *ChargingManagerCreateChargeParams {
	return &ChargingManagerCreateChargeParams{
		Context: ctx,
	}
}

// NewChargingManagerCreateChargeParamsWithHTTPClient creates a new ChargingManagerCreateChargeParams object
// with the ability to set a custom HTTPClient for a request.
func NewChargingManagerCreateChargeParamsWithHTTPClient(client *http.Client) *ChargingManagerCreateChargeParams {
	return &ChargingManagerCreateChargeParams{
		HTTPClient: client,
	}
}

/*
ChargingManagerCreateChargeParams contains all the parameters to send to the API endpoint

	for the charging manager create charge operation.

	Typically these are written to a http.Request.
*/
type ChargingManagerCreateChargeParams struct {

	// Body.
	Body *models.NewbillingCreateChargeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the charging manager create charge params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChargingManagerCreateChargeParams) WithDefaults() *ChargingManagerCreateChargeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the charging manager create charge params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChargingManagerCreateChargeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the charging manager create charge params
func (o *ChargingManagerCreateChargeParams) WithTimeout(timeout time.Duration) *ChargingManagerCreateChargeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the charging manager create charge params
func (o *ChargingManagerCreateChargeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the charging manager create charge params
func (o *ChargingManagerCreateChargeParams) WithContext(ctx context.Context) *ChargingManagerCreateChargeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the charging manager create charge params
func (o *ChargingManagerCreateChargeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the charging manager create charge params
func (o *ChargingManagerCreateChargeParams) WithHTTPClient(client *http.Client) *ChargingManagerCreateChargeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the charging manager create charge params
func (o *ChargingManagerCreateChargeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the charging manager create charge params
func (o *ChargingManagerCreateChargeParams) WithBody(body *models.NewbillingCreateChargeRequest) *ChargingManagerCreateChargeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the charging manager create charge params
func (o *ChargingManagerCreateChargeParams) SetBody(body *models.NewbillingCreateChargeRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChargingManagerCreateChargeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
