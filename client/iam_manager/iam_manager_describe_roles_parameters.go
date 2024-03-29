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
	"github.com/go-openapi/swag"
)

// NewIamManagerDescribeRolesParams creates a new IamManagerDescribeRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIamManagerDescribeRolesParams() *IamManagerDescribeRolesParams {
	return &IamManagerDescribeRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIamManagerDescribeRolesParamsWithTimeout creates a new IamManagerDescribeRolesParams object
// with the ability to set a timeout on a request.
func NewIamManagerDescribeRolesParamsWithTimeout(timeout time.Duration) *IamManagerDescribeRolesParams {
	return &IamManagerDescribeRolesParams{
		timeout: timeout,
	}
}

// NewIamManagerDescribeRolesParamsWithContext creates a new IamManagerDescribeRolesParams object
// with the ability to set a context for a request.
func NewIamManagerDescribeRolesParamsWithContext(ctx context.Context) *IamManagerDescribeRolesParams {
	return &IamManagerDescribeRolesParams{
		Context: ctx,
	}
}

// NewIamManagerDescribeRolesParamsWithHTTPClient creates a new IamManagerDescribeRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewIamManagerDescribeRolesParamsWithHTTPClient(client *http.Client) *IamManagerDescribeRolesParams {
	return &IamManagerDescribeRolesParams{
		HTTPClient: client,
	}
}

/*
IamManagerDescribeRolesParams contains all the parameters to send to the API endpoint

	for the iam manager describe roles operation.

	Typically these are written to a http.Request.
*/
type IamManagerDescribeRolesParams struct {

	/* Limit.

	   数据库查询每页大小 - 默认 20, 最大值 200.

	   Format: uint64
	*/
	Limit *string

	/* Offset.

	   数据库查询偏移位置 - 默认 0.

	   Format: uint64
	*/
	Offset *string

	/* Reverse.

	   是否倒序排序 - value = 0 sort ASC, value = 1 sort DESC.
	*/
	Reverse *bool

	/* RoleID.

	   角色ID.
	*/
	RoleID []string

	/* RoleName.

	   角色名称.
	*/
	RoleName []string

	/* SearchWord.

	   模糊查询关键字 - 支持字段：role_id/role_name.
	*/
	SearchWord *string

	/* SortKey.

	   排序字段 - 默认 create_time.
	*/
	SortKey *string

	/* Status.

	   角色状态- value = 1 可用, value = 2 禁用.
	*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the iam manager describe roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IamManagerDescribeRolesParams) WithDefaults() *IamManagerDescribeRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the iam manager describe roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IamManagerDescribeRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) WithTimeout(timeout time.Duration) *IamManagerDescribeRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) WithContext(ctx context.Context) *IamManagerDescribeRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) WithHTTPClient(client *http.Client) *IamManagerDescribeRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) WithLimit(limit *string) *IamManagerDescribeRolesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) WithOffset(offset *string) *IamManagerDescribeRolesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithReverse adds the reverse to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) WithReverse(reverse *bool) *IamManagerDescribeRolesParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithRoleID adds the roleID to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) WithRoleID(roleID []string) *IamManagerDescribeRolesParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) SetRoleID(roleID []string) {
	o.RoleID = roleID
}

// WithRoleName adds the roleName to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) WithRoleName(roleName []string) *IamManagerDescribeRolesParams {
	o.SetRoleName(roleName)
	return o
}

// SetRoleName adds the roleName to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) SetRoleName(roleName []string) {
	o.RoleName = roleName
}

// WithSearchWord adds the searchWord to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) WithSearchWord(searchWord *string) *IamManagerDescribeRolesParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) WithSortKey(sortKey *string) *IamManagerDescribeRolesParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStatus adds the status to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) WithStatus(status []string) *IamManagerDescribeRolesParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the iam manager describe roles params
func (o *IamManagerDescribeRolesParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *IamManagerDescribeRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit string

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Reverse != nil {

		// query param reverse
		var qrReverse bool

		if o.Reverse != nil {
			qrReverse = *o.Reverse
		}
		qReverse := swag.FormatBool(qrReverse)
		if qReverse != "" {

			if err := r.SetQueryParam("reverse", qReverse); err != nil {
				return err
			}
		}
	}

	if o.RoleID != nil {

		// binding items for role_id
		joinedRoleID := o.bindParamRoleID(reg)

		// query array param role_id
		if err := r.SetQueryParam("role_id", joinedRoleID...); err != nil {
			return err
		}
	}

	if o.RoleName != nil {

		// binding items for role_name
		joinedRoleName := o.bindParamRoleName(reg)

		// query array param role_name
		if err := r.SetQueryParam("role_name", joinedRoleName...); err != nil {
			return err
		}
	}

	if o.SearchWord != nil {

		// query param search_word
		var qrSearchWord string

		if o.SearchWord != nil {
			qrSearchWord = *o.SearchWord
		}
		qSearchWord := qrSearchWord
		if qSearchWord != "" {

			if err := r.SetQueryParam("search_word", qSearchWord); err != nil {
				return err
			}
		}
	}

	if o.SortKey != nil {

		// query param sort_key
		var qrSortKey string

		if o.SortKey != nil {
			qrSortKey = *o.SortKey
		}
		qSortKey := qrSortKey
		if qSortKey != "" {

			if err := r.SetQueryParam("sort_key", qSortKey); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// binding items for status
		joinedStatus := o.bindParamStatus(reg)

		// query array param status
		if err := r.SetQueryParam("status", joinedStatus...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamIamManagerDescribeRoles binds the parameter role_id
func (o *IamManagerDescribeRolesParams) bindParamRoleID(formats strfmt.Registry) []string {
	roleIDIR := o.RoleID

	var roleIDIC []string
	for _, roleIDIIR := range roleIDIR { // explode []string

		roleIDIIV := roleIDIIR // string as string
		roleIDIC = append(roleIDIC, roleIDIIV)
	}

	// items.CollectionFormat: "multi"
	roleIDIS := swag.JoinByFormat(roleIDIC, "multi")

	return roleIDIS
}

// bindParamIamManagerDescribeRoles binds the parameter role_name
func (o *IamManagerDescribeRolesParams) bindParamRoleName(formats strfmt.Registry) []string {
	roleNameIR := o.RoleName

	var roleNameIC []string
	for _, roleNameIIR := range roleNameIR { // explode []string

		roleNameIIV := roleNameIIR // string as string
		roleNameIC = append(roleNameIC, roleNameIIV)
	}

	// items.CollectionFormat: "multi"
	roleNameIS := swag.JoinByFormat(roleNameIC, "multi")

	return roleNameIS
}

// bindParamIamManagerDescribeRoles binds the parameter status
func (o *IamManagerDescribeRolesParams) bindParamStatus(formats strfmt.Registry) []string {
	statusIR := o.Status

	var statusIC []string
	for _, statusIIR := range statusIR { // explode []string

		statusIIV := statusIIR // string as string
		statusIC = append(statusIC, statusIIV)
	}

	// items.CollectionFormat: "multi"
	statusIS := swag.JoinByFormat(statusIC, "multi")

	return statusIS
}
