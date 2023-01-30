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

// NewIamManagerDeleteRoleParams creates a new IamManagerDeleteRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIamManagerDeleteRoleParams() *IamManagerDeleteRoleParams {
	return &IamManagerDeleteRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIamManagerDeleteRoleParamsWithTimeout creates a new IamManagerDeleteRoleParams object
// with the ability to set a timeout on a request.
func NewIamManagerDeleteRoleParamsWithTimeout(timeout time.Duration) *IamManagerDeleteRoleParams {
	return &IamManagerDeleteRoleParams{
		timeout: timeout,
	}
}

// NewIamManagerDeleteRoleParamsWithContext creates a new IamManagerDeleteRoleParams object
// with the ability to set a context for a request.
func NewIamManagerDeleteRoleParamsWithContext(ctx context.Context) *IamManagerDeleteRoleParams {
	return &IamManagerDeleteRoleParams{
		Context: ctx,
	}
}

// NewIamManagerDeleteRoleParamsWithHTTPClient creates a new IamManagerDeleteRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewIamManagerDeleteRoleParamsWithHTTPClient(client *http.Client) *IamManagerDeleteRoleParams {
	return &IamManagerDeleteRoleParams{
		HTTPClient: client,
	}
}

/*
IamManagerDeleteRoleParams contains all the parameters to send to the API endpoint

	for the iam manager delete role operation.

	Typically these are written to a http.Request.
*/
type IamManagerDeleteRoleParams struct {

	// Body.
	Body *models.NewbillingDeleteRoleRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the iam manager delete role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IamManagerDeleteRoleParams) WithDefaults() *IamManagerDeleteRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the iam manager delete role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IamManagerDeleteRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the iam manager delete role params
func (o *IamManagerDeleteRoleParams) WithTimeout(timeout time.Duration) *IamManagerDeleteRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the iam manager delete role params
func (o *IamManagerDeleteRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the iam manager delete role params
func (o *IamManagerDeleteRoleParams) WithContext(ctx context.Context) *IamManagerDeleteRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the iam manager delete role params
func (o *IamManagerDeleteRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the iam manager delete role params
func (o *IamManagerDeleteRoleParams) WithHTTPClient(client *http.Client) *IamManagerDeleteRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the iam manager delete role params
func (o *IamManagerDeleteRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the iam manager delete role params
func (o *IamManagerDeleteRoleParams) WithBody(body *models.NewbillingDeleteRoleRequest) *IamManagerDeleteRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the iam manager delete role params
func (o *IamManagerDeleteRoleParams) SetBody(body *models.NewbillingDeleteRoleRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IamManagerDeleteRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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