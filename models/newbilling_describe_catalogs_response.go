// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingDescribeCatalogsResponse 查询产品目录返回信息
//
// swagger:model newbillingDescribeCatalogsResponse
type NewbillingDescribeCatalogsResponse struct {

	// 返回结果列表
	CatalogSet []*NewbillingCatalog `json:"catalog_set"`

	// 返回总数
	Total int64 `json:"total,omitempty"`
}

// Validate validates this newbilling describe catalogs response
func (m *NewbillingDescribeCatalogsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCatalogSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeCatalogsResponse) validateCatalogSet(formats strfmt.Registry) error {
	if swag.IsZero(m.CatalogSet) { // not required
		return nil
	}

	for i := 0; i < len(m.CatalogSet); i++ {
		if swag.IsZero(m.CatalogSet[i]) { // not required
			continue
		}

		if m.CatalogSet[i] != nil {
			if err := m.CatalogSet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("catalog_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("catalog_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling describe catalogs response based on the context it is used
func (m *NewbillingDescribeCatalogsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCatalogSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingDescribeCatalogsResponse) contextValidateCatalogSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CatalogSet); i++ {

		if m.CatalogSet[i] != nil {
			if err := m.CatalogSet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("catalog_set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("catalog_set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingDescribeCatalogsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingDescribeCatalogsResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingDescribeCatalogsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}