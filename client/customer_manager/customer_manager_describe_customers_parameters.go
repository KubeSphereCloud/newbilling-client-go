// Code generated by go-swagger; DO NOT EDIT.

package customer_manager

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

// NewCustomerManagerDescribeCustomersParams creates a new CustomerManagerDescribeCustomersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerManagerDescribeCustomersParams() *CustomerManagerDescribeCustomersParams {
	return &CustomerManagerDescribeCustomersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerManagerDescribeCustomersParamsWithTimeout creates a new CustomerManagerDescribeCustomersParams object
// with the ability to set a timeout on a request.
func NewCustomerManagerDescribeCustomersParamsWithTimeout(timeout time.Duration) *CustomerManagerDescribeCustomersParams {
	return &CustomerManagerDescribeCustomersParams{
		timeout: timeout,
	}
}

// NewCustomerManagerDescribeCustomersParamsWithContext creates a new CustomerManagerDescribeCustomersParams object
// with the ability to set a context for a request.
func NewCustomerManagerDescribeCustomersParamsWithContext(ctx context.Context) *CustomerManagerDescribeCustomersParams {
	return &CustomerManagerDescribeCustomersParams{
		Context: ctx,
	}
}

// NewCustomerManagerDescribeCustomersParamsWithHTTPClient creates a new CustomerManagerDescribeCustomersParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerManagerDescribeCustomersParamsWithHTTPClient(client *http.Client) *CustomerManagerDescribeCustomersParams {
	return &CustomerManagerDescribeCustomersParams{
		HTTPClient: client,
	}
}

/*
CustomerManagerDescribeCustomersParams contains all the parameters to send to the API endpoint

	for the customer manager describe customers operation.

	Typically these are written to a http.Request.
*/
type CustomerManagerDescribeCustomersParams struct {

	/* CustomerID.

	   客户ID.
	*/
	CustomerID []string

	/* IsParent.

	   是否是主帐户.
	*/
	IsParent *bool

	/* Limit.

	   数据库查询每页大小 - 默认 20, 最大值 200.

	   Format: uint64
	*/
	Limit *string

	/* Name.

	   客户姓名.
	*/
	Name []string

	/* Offset.

	   数据库查询偏移位置 - 默认 0.

	   Format: uint64
	*/
	Offset *string

	/* ParentCustomerID.

	   父客户Id.
	*/
	ParentCustomerID []string

	/* Reverse.

	   是否倒序排序 - value = 0 sort ASC, value = 1 sort DESC.
	*/
	Reverse *bool

	/* SearchWord.

	   模糊查询关键字 - 支持字段：customer_id/customer_source/access_sys_id/parent_customer_id/name.
	*/
	SearchWord *string

	/* SortKey.

	   排序字段 - 默认 create_time.
	*/
	SortKey *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer manager describe customers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerManagerDescribeCustomersParams) WithDefaults() *CustomerManagerDescribeCustomersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer manager describe customers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerManagerDescribeCustomersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) WithTimeout(timeout time.Duration) *CustomerManagerDescribeCustomersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) WithContext(ctx context.Context) *CustomerManagerDescribeCustomersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) WithHTTPClient(client *http.Client) *CustomerManagerDescribeCustomersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) WithCustomerID(customerID []string) *CustomerManagerDescribeCustomersParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) SetCustomerID(customerID []string) {
	o.CustomerID = customerID
}

// WithIsParent adds the isParent to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) WithIsParent(isParent *bool) *CustomerManagerDescribeCustomersParams {
	o.SetIsParent(isParent)
	return o
}

// SetIsParent adds the isParent to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) SetIsParent(isParent *bool) {
	o.IsParent = isParent
}

// WithLimit adds the limit to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) WithLimit(limit *string) *CustomerManagerDescribeCustomersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithName adds the name to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) WithName(name []string) *CustomerManagerDescribeCustomersParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) SetName(name []string) {
	o.Name = name
}

// WithOffset adds the offset to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) WithOffset(offset *string) *CustomerManagerDescribeCustomersParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithParentCustomerID adds the parentCustomerID to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) WithParentCustomerID(parentCustomerID []string) *CustomerManagerDescribeCustomersParams {
	o.SetParentCustomerID(parentCustomerID)
	return o
}

// SetParentCustomerID adds the parentCustomerId to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) SetParentCustomerID(parentCustomerID []string) {
	o.ParentCustomerID = parentCustomerID
}

// WithReverse adds the reverse to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) WithReverse(reverse *bool) *CustomerManagerDescribeCustomersParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSearchWord adds the searchWord to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) WithSearchWord(searchWord *string) *CustomerManagerDescribeCustomersParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) WithSortKey(sortKey *string) *CustomerManagerDescribeCustomersParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the customer manager describe customers params
func (o *CustomerManagerDescribeCustomersParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerManagerDescribeCustomersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CustomerID != nil {

		// binding items for customer_id
		joinedCustomerID := o.bindParamCustomerID(reg)

		// query array param customer_id
		if err := r.SetQueryParam("customer_id", joinedCustomerID...); err != nil {
			return err
		}
	}

	if o.IsParent != nil {

		// query param is_parent
		var qrIsParent bool

		if o.IsParent != nil {
			qrIsParent = *o.IsParent
		}
		qIsParent := swag.FormatBool(qrIsParent)
		if qIsParent != "" {

			if err := r.SetQueryParam("is_parent", qIsParent); err != nil {
				return err
			}
		}
	}

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

	if o.Name != nil {

		// binding items for name
		joinedName := o.bindParamName(reg)

		// query array param name
		if err := r.SetQueryParam("name", joinedName...); err != nil {
			return err
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

	if o.ParentCustomerID != nil {

		// binding items for parent_customer_id
		joinedParentCustomerID := o.bindParamParentCustomerID(reg)

		// query array param parent_customer_id
		if err := r.SetQueryParam("parent_customer_id", joinedParentCustomerID...); err != nil {
			return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCustomerManagerDescribeCustomers binds the parameter customer_id
func (o *CustomerManagerDescribeCustomersParams) bindParamCustomerID(formats strfmt.Registry) []string {
	customerIDIR := o.CustomerID

	var customerIDIC []string
	for _, customerIDIIR := range customerIDIR { // explode []string

		customerIDIIV := customerIDIIR // string as string
		customerIDIC = append(customerIDIC, customerIDIIV)
	}

	// items.CollectionFormat: "multi"
	customerIDIS := swag.JoinByFormat(customerIDIC, "multi")

	return customerIDIS
}

// bindParamCustomerManagerDescribeCustomers binds the parameter name
func (o *CustomerManagerDescribeCustomersParams) bindParamName(formats strfmt.Registry) []string {
	nameIR := o.Name

	var nameIC []string
	for _, nameIIR := range nameIR { // explode []string

		nameIIV := nameIIR // string as string
		nameIC = append(nameIC, nameIIV)
	}

	// items.CollectionFormat: "multi"
	nameIS := swag.JoinByFormat(nameIC, "multi")

	return nameIS
}

// bindParamCustomerManagerDescribeCustomers binds the parameter parent_customer_id
func (o *CustomerManagerDescribeCustomersParams) bindParamParentCustomerID(formats strfmt.Registry) []string {
	parentCustomerIDIR := o.ParentCustomerID

	var parentCustomerIDIC []string
	for _, parentCustomerIDIIR := range parentCustomerIDIR { // explode []string

		parentCustomerIDIIV := parentCustomerIDIIR // string as string
		parentCustomerIDIC = append(parentCustomerIDIC, parentCustomerIDIIV)
	}

	// items.CollectionFormat: "multi"
	parentCustomerIDIS := swag.JoinByFormat(parentCustomerIDIC, "multi")

	return parentCustomerIDIS
}
