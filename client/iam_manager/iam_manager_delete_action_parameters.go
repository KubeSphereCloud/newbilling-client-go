// Code generated by go-swagger; DO NOT EDIT.

package iam_manager

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

// NewIamManagerDeleteActionParams creates a new IamManagerDeleteActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIamManagerDeleteActionParams() *IamManagerDeleteActionParams {
	return &IamManagerDeleteActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIamManagerDeleteActionParamsWithTimeout creates a new IamManagerDeleteActionParams object
// with the ability to set a timeout on a request.
func NewIamManagerDeleteActionParamsWithTimeout(timeout time.Duration) *IamManagerDeleteActionParams {
	return &IamManagerDeleteActionParams{
		timeout: timeout,
	}
}

// NewIamManagerDeleteActionParamsWithContext creates a new IamManagerDeleteActionParams object
// with the ability to set a context for a request.
func NewIamManagerDeleteActionParamsWithContext(ctx context.Context) *IamManagerDeleteActionParams {
	return &IamManagerDeleteActionParams{
		Context: ctx,
	}
}

// NewIamManagerDeleteActionParamsWithHTTPClient creates a new IamManagerDeleteActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewIamManagerDeleteActionParamsWithHTTPClient(client *http.Client) *IamManagerDeleteActionParams {
	return &IamManagerDeleteActionParams{
		HTTPClient: client,
	}
}

/*
IamManagerDeleteActionParams contains all the parameters to send to the API endpoint

	for the iam manager delete action operation.

	Typically these are written to a http.Request.
*/
type IamManagerDeleteActionParams struct {

	// Body.
	Body *models.NewbillingDeleteActionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the iam manager delete action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IamManagerDeleteActionParams) WithDefaults() *IamManagerDeleteActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the iam manager delete action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IamManagerDeleteActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the iam manager delete action params
func (o *IamManagerDeleteActionParams) WithTimeout(timeout time.Duration) *IamManagerDeleteActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the iam manager delete action params
func (o *IamManagerDeleteActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the iam manager delete action params
func (o *IamManagerDeleteActionParams) WithContext(ctx context.Context) *IamManagerDeleteActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the iam manager delete action params
func (o *IamManagerDeleteActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the iam manager delete action params
func (o *IamManagerDeleteActionParams) WithHTTPClient(client *http.Client) *IamManagerDeleteActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the iam manager delete action params
func (o *IamManagerDeleteActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the iam manager delete action params
func (o *IamManagerDeleteActionParams) WithBody(body *models.NewbillingDeleteActionRequest) *IamManagerDeleteActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the iam manager delete action params
func (o *IamManagerDeleteActionParams) SetBody(body *models.NewbillingDeleteActionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IamManagerDeleteActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
