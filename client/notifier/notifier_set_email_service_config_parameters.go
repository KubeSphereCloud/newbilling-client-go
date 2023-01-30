// Code generated by go-swagger; DO NOT EDIT.

package notifier

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

// NewNotifierSetEmailServiceConfigParams creates a new NotifierSetEmailServiceConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNotifierSetEmailServiceConfigParams() *NotifierSetEmailServiceConfigParams {
	return &NotifierSetEmailServiceConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNotifierSetEmailServiceConfigParamsWithTimeout creates a new NotifierSetEmailServiceConfigParams object
// with the ability to set a timeout on a request.
func NewNotifierSetEmailServiceConfigParamsWithTimeout(timeout time.Duration) *NotifierSetEmailServiceConfigParams {
	return &NotifierSetEmailServiceConfigParams{
		timeout: timeout,
	}
}

// NewNotifierSetEmailServiceConfigParamsWithContext creates a new NotifierSetEmailServiceConfigParams object
// with the ability to set a context for a request.
func NewNotifierSetEmailServiceConfigParamsWithContext(ctx context.Context) *NotifierSetEmailServiceConfigParams {
	return &NotifierSetEmailServiceConfigParams{
		Context: ctx,
	}
}

// NewNotifierSetEmailServiceConfigParamsWithHTTPClient creates a new NotifierSetEmailServiceConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewNotifierSetEmailServiceConfigParamsWithHTTPClient(client *http.Client) *NotifierSetEmailServiceConfigParams {
	return &NotifierSetEmailServiceConfigParams{
		HTTPClient: client,
	}
}

/*
NotifierSetEmailServiceConfigParams contains all the parameters to send to the API endpoint

	for the notifier set email service config operation.

	Typically these are written to a http.Request.
*/
type NotifierSetEmailServiceConfigParams struct {

	// Body.
	Body *models.NewbillingSetEmailServiceConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the notifier set email service config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotifierSetEmailServiceConfigParams) WithDefaults() *NotifierSetEmailServiceConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the notifier set email service config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotifierSetEmailServiceConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the notifier set email service config params
func (o *NotifierSetEmailServiceConfigParams) WithTimeout(timeout time.Duration) *NotifierSetEmailServiceConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notifier set email service config params
func (o *NotifierSetEmailServiceConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notifier set email service config params
func (o *NotifierSetEmailServiceConfigParams) WithContext(ctx context.Context) *NotifierSetEmailServiceConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notifier set email service config params
func (o *NotifierSetEmailServiceConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notifier set email service config params
func (o *NotifierSetEmailServiceConfigParams) WithHTTPClient(client *http.Client) *NotifierSetEmailServiceConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notifier set email service config params
func (o *NotifierSetEmailServiceConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the notifier set email service config params
func (o *NotifierSetEmailServiceConfigParams) WithBody(body *models.NewbillingSetEmailServiceConfigRequest) *NotifierSetEmailServiceConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the notifier set email service config params
func (o *NotifierSetEmailServiceConfigParams) SetBody(body *models.NewbillingSetEmailServiceConfigRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *NotifierSetEmailServiceConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
