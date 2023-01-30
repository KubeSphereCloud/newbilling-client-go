// Code generated by go-swagger; DO NOT EDIT.

package pricing_manager

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

// NewPricingManagerDescribeCatalogsParams creates a new PricingManagerDescribeCatalogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPricingManagerDescribeCatalogsParams() *PricingManagerDescribeCatalogsParams {
	return &PricingManagerDescribeCatalogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPricingManagerDescribeCatalogsParamsWithTimeout creates a new PricingManagerDescribeCatalogsParams object
// with the ability to set a timeout on a request.
func NewPricingManagerDescribeCatalogsParamsWithTimeout(timeout time.Duration) *PricingManagerDescribeCatalogsParams {
	return &PricingManagerDescribeCatalogsParams{
		timeout: timeout,
	}
}

// NewPricingManagerDescribeCatalogsParamsWithContext creates a new PricingManagerDescribeCatalogsParams object
// with the ability to set a context for a request.
func NewPricingManagerDescribeCatalogsParamsWithContext(ctx context.Context) *PricingManagerDescribeCatalogsParams {
	return &PricingManagerDescribeCatalogsParams{
		Context: ctx,
	}
}

// NewPricingManagerDescribeCatalogsParamsWithHTTPClient creates a new PricingManagerDescribeCatalogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPricingManagerDescribeCatalogsParamsWithHTTPClient(client *http.Client) *PricingManagerDescribeCatalogsParams {
	return &PricingManagerDescribeCatalogsParams{
		HTTPClient: client,
	}
}

/*
PricingManagerDescribeCatalogsParams contains all the parameters to send to the API endpoint

	for the pricing manager describe catalogs operation.

	Typically these are written to a http.Request.
*/
type PricingManagerDescribeCatalogsParams struct {

	/* AccessSysID.

	   接入系统ID.
	*/
	AccessSysID []string

	/* CataCode.

	   产品目录编码.
	*/
	CataCode []string

	/* CataID.

	   产品目录ID.
	*/
	CataID []string

	/* Limit.

	   数据库查询每页大小 - 默认 20, 最大值 200.

	   Format: uint64
	*/
	Limit *string

	/* Name.

	   产品目录名称.
	*/
	Name []string

	/* Offset.

	   数据库查询偏移位置 - 默认 0.

	   Format: uint64
	*/
	Offset *string

	/* ParentCataID.

	   产品目录的父目录ID.
	*/
	ParentCataID []string

	/* Reverse.

	   是否倒序排序 - value = 0 sort ASC, value = 1 sort DESC.
	*/
	Reverse *bool

	/* SearchWord.

	   模糊查询关键字.
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

// WithDefaults hydrates default values in the pricing manager describe catalogs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerDescribeCatalogsParams) WithDefaults() *PricingManagerDescribeCatalogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pricing manager describe catalogs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PricingManagerDescribeCatalogsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) WithTimeout(timeout time.Duration) *PricingManagerDescribeCatalogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) WithContext(ctx context.Context) *PricingManagerDescribeCatalogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) WithHTTPClient(client *http.Client) *PricingManagerDescribeCatalogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessSysID adds the accessSysID to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) WithAccessSysID(accessSysID []string) *PricingManagerDescribeCatalogsParams {
	o.SetAccessSysID(accessSysID)
	return o
}

// SetAccessSysID adds the accessSysId to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) SetAccessSysID(accessSysID []string) {
	o.AccessSysID = accessSysID
}

// WithCataCode adds the cataCode to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) WithCataCode(cataCode []string) *PricingManagerDescribeCatalogsParams {
	o.SetCataCode(cataCode)
	return o
}

// SetCataCode adds the cataCode to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) SetCataCode(cataCode []string) {
	o.CataCode = cataCode
}

// WithCataID adds the cataID to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) WithCataID(cataID []string) *PricingManagerDescribeCatalogsParams {
	o.SetCataID(cataID)
	return o
}

// SetCataID adds the cataId to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) SetCataID(cataID []string) {
	o.CataID = cataID
}

// WithLimit adds the limit to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) WithLimit(limit *string) *PricingManagerDescribeCatalogsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithName adds the name to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) WithName(name []string) *PricingManagerDescribeCatalogsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) SetName(name []string) {
	o.Name = name
}

// WithOffset adds the offset to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) WithOffset(offset *string) *PricingManagerDescribeCatalogsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithParentCataID adds the parentCataID to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) WithParentCataID(parentCataID []string) *PricingManagerDescribeCatalogsParams {
	o.SetParentCataID(parentCataID)
	return o
}

// SetParentCataID adds the parentCataId to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) SetParentCataID(parentCataID []string) {
	o.ParentCataID = parentCataID
}

// WithReverse adds the reverse to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) WithReverse(reverse *bool) *PricingManagerDescribeCatalogsParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSearchWord adds the searchWord to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) WithSearchWord(searchWord *string) *PricingManagerDescribeCatalogsParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) WithSortKey(sortKey *string) *PricingManagerDescribeCatalogsParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the pricing manager describe catalogs params
func (o *PricingManagerDescribeCatalogsParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WriteToRequest writes these params to a swagger request
func (o *PricingManagerDescribeCatalogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessSysID != nil {

		// binding items for access_sys_id
		joinedAccessSysID := o.bindParamAccessSysID(reg)

		// query array param access_sys_id
		if err := r.SetQueryParam("access_sys_id", joinedAccessSysID...); err != nil {
			return err
		}
	}

	if o.CataCode != nil {

		// binding items for cata_code
		joinedCataCode := o.bindParamCataCode(reg)

		// query array param cata_code
		if err := r.SetQueryParam("cata_code", joinedCataCode...); err != nil {
			return err
		}
	}

	if o.CataID != nil {

		// binding items for cata_id
		joinedCataID := o.bindParamCataID(reg)

		// query array param cata_id
		if err := r.SetQueryParam("cata_id", joinedCataID...); err != nil {
			return err
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

	if o.ParentCataID != nil {

		// binding items for parent_cata_id
		joinedParentCataID := o.bindParamParentCataID(reg)

		// query array param parent_cata_id
		if err := r.SetQueryParam("parent_cata_id", joinedParentCataID...); err != nil {
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

// bindParamPricingManagerDescribeCatalogs binds the parameter access_sys_id
func (o *PricingManagerDescribeCatalogsParams) bindParamAccessSysID(formats strfmt.Registry) []string {
	accessSysIDIR := o.AccessSysID

	var accessSysIDIC []string
	for _, accessSysIDIIR := range accessSysIDIR { // explode []string

		accessSysIDIIV := accessSysIDIIR // string as string
		accessSysIDIC = append(accessSysIDIC, accessSysIDIIV)
	}

	// items.CollectionFormat: "multi"
	accessSysIDIS := swag.JoinByFormat(accessSysIDIC, "multi")

	return accessSysIDIS
}

// bindParamPricingManagerDescribeCatalogs binds the parameter cata_code
func (o *PricingManagerDescribeCatalogsParams) bindParamCataCode(formats strfmt.Registry) []string {
	cataCodeIR := o.CataCode

	var cataCodeIC []string
	for _, cataCodeIIR := range cataCodeIR { // explode []string

		cataCodeIIV := cataCodeIIR // string as string
		cataCodeIC = append(cataCodeIC, cataCodeIIV)
	}

	// items.CollectionFormat: "multi"
	cataCodeIS := swag.JoinByFormat(cataCodeIC, "multi")

	return cataCodeIS
}

// bindParamPricingManagerDescribeCatalogs binds the parameter cata_id
func (o *PricingManagerDescribeCatalogsParams) bindParamCataID(formats strfmt.Registry) []string {
	cataIDIR := o.CataID

	var cataIDIC []string
	for _, cataIDIIR := range cataIDIR { // explode []string

		cataIDIIV := cataIDIIR // string as string
		cataIDIC = append(cataIDIC, cataIDIIV)
	}

	// items.CollectionFormat: "multi"
	cataIDIS := swag.JoinByFormat(cataIDIC, "multi")

	return cataIDIS
}

// bindParamPricingManagerDescribeCatalogs binds the parameter name
func (o *PricingManagerDescribeCatalogsParams) bindParamName(formats strfmt.Registry) []string {
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

// bindParamPricingManagerDescribeCatalogs binds the parameter parent_cata_id
func (o *PricingManagerDescribeCatalogsParams) bindParamParentCataID(formats strfmt.Registry) []string {
	parentCataIDIR := o.ParentCataID

	var parentCataIDIC []string
	for _, parentCataIDIIR := range parentCataIDIR { // explode []string

		parentCataIDIIV := parentCataIDIIR // string as string
		parentCataIDIC = append(parentCataIDIC, parentCataIDIIV)
	}

	// items.CollectionFormat: "multi"
	parentCataIDIS := swag.JoinByFormat(parentCataIDIC, "multi")

	return parentCataIDIS
}
