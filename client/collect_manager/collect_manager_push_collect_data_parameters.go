// Code generated by go-swagger; DO NOT EDIT.

package collect_manager

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

// NewCollectManagerPushCollectDataParams creates a new CollectManagerPushCollectDataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCollectManagerPushCollectDataParams() *CollectManagerPushCollectDataParams {
	return &CollectManagerPushCollectDataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCollectManagerPushCollectDataParamsWithTimeout creates a new CollectManagerPushCollectDataParams object
// with the ability to set a timeout on a request.
func NewCollectManagerPushCollectDataParamsWithTimeout(timeout time.Duration) *CollectManagerPushCollectDataParams {
	return &CollectManagerPushCollectDataParams{
		timeout: timeout,
	}
}

// NewCollectManagerPushCollectDataParamsWithContext creates a new CollectManagerPushCollectDataParams object
// with the ability to set a context for a request.
func NewCollectManagerPushCollectDataParamsWithContext(ctx context.Context) *CollectManagerPushCollectDataParams {
	return &CollectManagerPushCollectDataParams{
		Context: ctx,
	}
}

// NewCollectManagerPushCollectDataParamsWithHTTPClient creates a new CollectManagerPushCollectDataParams object
// with the ability to set a custom HTTPClient for a request.
func NewCollectManagerPushCollectDataParamsWithHTTPClient(client *http.Client) *CollectManagerPushCollectDataParams {
	return &CollectManagerPushCollectDataParams{
		HTTPClient: client,
	}
}

/*
CollectManagerPushCollectDataParams contains all the parameters to send to the API endpoint

	for the collect manager push collect data operation.

	Typically these are written to a http.Request.
*/
type CollectManagerPushCollectDataParams struct {

	// Body.
	Body *models.NewbillingPushCollectDataRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the collect manager push collect data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CollectManagerPushCollectDataParams) WithDefaults() *CollectManagerPushCollectDataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the collect manager push collect data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CollectManagerPushCollectDataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the collect manager push collect data params
func (o *CollectManagerPushCollectDataParams) WithTimeout(timeout time.Duration) *CollectManagerPushCollectDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the collect manager push collect data params
func (o *CollectManagerPushCollectDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the collect manager push collect data params
func (o *CollectManagerPushCollectDataParams) WithContext(ctx context.Context) *CollectManagerPushCollectDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the collect manager push collect data params
func (o *CollectManagerPushCollectDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the collect manager push collect data params
func (o *CollectManagerPushCollectDataParams) WithHTTPClient(client *http.Client) *CollectManagerPushCollectDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the collect manager push collect data params
func (o *CollectManagerPushCollectDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the collect manager push collect data params
func (o *CollectManagerPushCollectDataParams) WithBody(body *models.NewbillingPushCollectDataRequest) *CollectManagerPushCollectDataParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the collect manager push collect data params
func (o *CollectManagerPushCollectDataParams) SetBody(body *models.NewbillingPushCollectDataRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CollectManagerPushCollectDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
