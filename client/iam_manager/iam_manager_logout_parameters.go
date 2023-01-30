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

// NewIamManagerLogoutParams creates a new IamManagerLogoutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIamManagerLogoutParams() *IamManagerLogoutParams {
	return &IamManagerLogoutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIamManagerLogoutParamsWithTimeout creates a new IamManagerLogoutParams object
// with the ability to set a timeout on a request.
func NewIamManagerLogoutParamsWithTimeout(timeout time.Duration) *IamManagerLogoutParams {
	return &IamManagerLogoutParams{
		timeout: timeout,
	}
}

// NewIamManagerLogoutParamsWithContext creates a new IamManagerLogoutParams object
// with the ability to set a context for a request.
func NewIamManagerLogoutParamsWithContext(ctx context.Context) *IamManagerLogoutParams {
	return &IamManagerLogoutParams{
		Context: ctx,
	}
}

// NewIamManagerLogoutParamsWithHTTPClient creates a new IamManagerLogoutParams object
// with the ability to set a custom HTTPClient for a request.
func NewIamManagerLogoutParamsWithHTTPClient(client *http.Client) *IamManagerLogoutParams {
	return &IamManagerLogoutParams{
		HTTPClient: client,
	}
}

/*
IamManagerLogoutParams contains all the parameters to send to the API endpoint

	for the iam manager logout operation.

	Typically these are written to a http.Request.
*/
type IamManagerLogoutParams struct {

	// Body.
	Body models.NewbillingLogoutRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the iam manager logout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IamManagerLogoutParams) WithDefaults() *IamManagerLogoutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the iam manager logout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IamManagerLogoutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the iam manager logout params
func (o *IamManagerLogoutParams) WithTimeout(timeout time.Duration) *IamManagerLogoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the iam manager logout params
func (o *IamManagerLogoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the iam manager logout params
func (o *IamManagerLogoutParams) WithContext(ctx context.Context) *IamManagerLogoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the iam manager logout params
func (o *IamManagerLogoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the iam manager logout params
func (o *IamManagerLogoutParams) WithHTTPClient(client *http.Client) *IamManagerLogoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the iam manager logout params
func (o *IamManagerLogoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the iam manager logout params
func (o *IamManagerLogoutParams) WithBody(body models.NewbillingLogoutRequest) *IamManagerLogoutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the iam manager logout params
func (o *IamManagerLogoutParams) SetBody(body models.NewbillingLogoutRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IamManagerLogoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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