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

// NewCollectManagerStartCollectParams creates a new CollectManagerStartCollectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCollectManagerStartCollectParams() *CollectManagerStartCollectParams {
	return &CollectManagerStartCollectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCollectManagerStartCollectParamsWithTimeout creates a new CollectManagerStartCollectParams object
// with the ability to set a timeout on a request.
func NewCollectManagerStartCollectParamsWithTimeout(timeout time.Duration) *CollectManagerStartCollectParams {
	return &CollectManagerStartCollectParams{
		timeout: timeout,
	}
}

// NewCollectManagerStartCollectParamsWithContext creates a new CollectManagerStartCollectParams object
// with the ability to set a context for a request.
func NewCollectManagerStartCollectParamsWithContext(ctx context.Context) *CollectManagerStartCollectParams {
	return &CollectManagerStartCollectParams{
		Context: ctx,
	}
}

// NewCollectManagerStartCollectParamsWithHTTPClient creates a new CollectManagerStartCollectParams object
// with the ability to set a custom HTTPClient for a request.
func NewCollectManagerStartCollectParamsWithHTTPClient(client *http.Client) *CollectManagerStartCollectParams {
	return &CollectManagerStartCollectParams{
		HTTPClient: client,
	}
}

/*
CollectManagerStartCollectParams contains all the parameters to send to the API endpoint

	for the collect manager start collect operation.

	Typically these are written to a http.Request.
*/
type CollectManagerStartCollectParams struct {

	// Body.
	Body *models.NewbillingStartCollectRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the collect manager start collect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CollectManagerStartCollectParams) WithDefaults() *CollectManagerStartCollectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the collect manager start collect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CollectManagerStartCollectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the collect manager start collect params
func (o *CollectManagerStartCollectParams) WithTimeout(timeout time.Duration) *CollectManagerStartCollectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the collect manager start collect params
func (o *CollectManagerStartCollectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the collect manager start collect params
func (o *CollectManagerStartCollectParams) WithContext(ctx context.Context) *CollectManagerStartCollectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the collect manager start collect params
func (o *CollectManagerStartCollectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the collect manager start collect params
func (o *CollectManagerStartCollectParams) WithHTTPClient(client *http.Client) *CollectManagerStartCollectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the collect manager start collect params
func (o *CollectManagerStartCollectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the collect manager start collect params
func (o *CollectManagerStartCollectParams) WithBody(body *models.NewbillingStartCollectRequest) *CollectManagerStartCollectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the collect manager start collect params
func (o *CollectManagerStartCollectParams) SetBody(body *models.NewbillingStartCollectRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CollectManagerStartCollectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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