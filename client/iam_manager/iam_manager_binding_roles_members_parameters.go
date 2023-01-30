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

// NewIamManagerBindingRolesMembersParams creates a new IamManagerBindingRolesMembersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIamManagerBindingRolesMembersParams() *IamManagerBindingRolesMembersParams {
	return &IamManagerBindingRolesMembersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIamManagerBindingRolesMembersParamsWithTimeout creates a new IamManagerBindingRolesMembersParams object
// with the ability to set a timeout on a request.
func NewIamManagerBindingRolesMembersParamsWithTimeout(timeout time.Duration) *IamManagerBindingRolesMembersParams {
	return &IamManagerBindingRolesMembersParams{
		timeout: timeout,
	}
}

// NewIamManagerBindingRolesMembersParamsWithContext creates a new IamManagerBindingRolesMembersParams object
// with the ability to set a context for a request.
func NewIamManagerBindingRolesMembersParamsWithContext(ctx context.Context) *IamManagerBindingRolesMembersParams {
	return &IamManagerBindingRolesMembersParams{
		Context: ctx,
	}
}

// NewIamManagerBindingRolesMembersParamsWithHTTPClient creates a new IamManagerBindingRolesMembersParams object
// with the ability to set a custom HTTPClient for a request.
func NewIamManagerBindingRolesMembersParamsWithHTTPClient(client *http.Client) *IamManagerBindingRolesMembersParams {
	return &IamManagerBindingRolesMembersParams{
		HTTPClient: client,
	}
}

/*
IamManagerBindingRolesMembersParams contains all the parameters to send to the API endpoint

	for the iam manager binding roles members operation.

	Typically these are written to a http.Request.
*/
type IamManagerBindingRolesMembersParams struct {

	// Body.
	Body *models.NewbillingBindingMembersRolesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the iam manager binding roles members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IamManagerBindingRolesMembersParams) WithDefaults() *IamManagerBindingRolesMembersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the iam manager binding roles members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IamManagerBindingRolesMembersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the iam manager binding roles members params
func (o *IamManagerBindingRolesMembersParams) WithTimeout(timeout time.Duration) *IamManagerBindingRolesMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the iam manager binding roles members params
func (o *IamManagerBindingRolesMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the iam manager binding roles members params
func (o *IamManagerBindingRolesMembersParams) WithContext(ctx context.Context) *IamManagerBindingRolesMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the iam manager binding roles members params
func (o *IamManagerBindingRolesMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the iam manager binding roles members params
func (o *IamManagerBindingRolesMembersParams) WithHTTPClient(client *http.Client) *IamManagerBindingRolesMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the iam manager binding roles members params
func (o *IamManagerBindingRolesMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the iam manager binding roles members params
func (o *IamManagerBindingRolesMembersParams) WithBody(body *models.NewbillingBindingMembersRolesRequest) *IamManagerBindingRolesMembersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the iam manager binding roles members params
func (o *IamManagerBindingRolesMembersParams) SetBody(body *models.NewbillingBindingMembersRolesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IamManagerBindingRolesMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
