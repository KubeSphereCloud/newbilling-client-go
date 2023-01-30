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

// NewbillingGetPricingResponse newbilling get pricing response
//
// swagger:model newbillingGetPricingResponse
type NewbillingGetPricingResponse struct {

	// attrs
	Attrs []*NewbillingAttrInfo `json:"attrs"`

	// catalog
	Catalog *NewbillingCreateCatalogRequest `json:"catalog,omitempty"`

	// components
	Components []*NewbillingCompInfo `json:"components"`

	// plan
	Plan *NewbillingCreatePlanRequest `json:"plan,omitempty"`

	// product
	Product *NewbillingCreateProductRequest `json:"product,omitempty"`

	// pub attr filters
	PubAttrFilters []*NewbillingFilterInfo `json:"pub_attr_filters"`
}

// Validate validates this newbilling get pricing response
func (m *NewbillingGetPricingResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttrs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCatalog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePubAttrFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingGetPricingResponse) validateAttrs(formats strfmt.Registry) error {
	if swag.IsZero(m.Attrs) { // not required
		return nil
	}

	for i := 0; i < len(m.Attrs); i++ {
		if swag.IsZero(m.Attrs[i]) { // not required
			continue
		}

		if m.Attrs[i] != nil {
			if err := m.Attrs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attrs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attrs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NewbillingGetPricingResponse) validateCatalog(formats strfmt.Registry) error {
	if swag.IsZero(m.Catalog) { // not required
		return nil
	}

	if m.Catalog != nil {
		if err := m.Catalog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("catalog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("catalog")
			}
			return err
		}
	}

	return nil
}

func (m *NewbillingGetPricingResponse) validateComponents(formats strfmt.Registry) error {
	if swag.IsZero(m.Components) { // not required
		return nil
	}

	for i := 0; i < len(m.Components); i++ {
		if swag.IsZero(m.Components[i]) { // not required
			continue
		}

		if m.Components[i] != nil {
			if err := m.Components[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NewbillingGetPricingResponse) validatePlan(formats strfmt.Registry) error {
	if swag.IsZero(m.Plan) { // not required
		return nil
	}

	if m.Plan != nil {
		if err := m.Plan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plan")
			}
			return err
		}
	}

	return nil
}

func (m *NewbillingGetPricingResponse) validateProduct(formats strfmt.Registry) error {
	if swag.IsZero(m.Product) { // not required
		return nil
	}

	if m.Product != nil {
		if err := m.Product.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("product")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("product")
			}
			return err
		}
	}

	return nil
}

func (m *NewbillingGetPricingResponse) validatePubAttrFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.PubAttrFilters) { // not required
		return nil
	}

	for i := 0; i < len(m.PubAttrFilters); i++ {
		if swag.IsZero(m.PubAttrFilters[i]) { // not required
			continue
		}

		if m.PubAttrFilters[i] != nil {
			if err := m.PubAttrFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pub_attr_filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pub_attr_filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this newbilling get pricing response based on the context it is used
func (m *NewbillingGetPricingResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttrs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCatalog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProduct(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePubAttrFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewbillingGetPricingResponse) contextValidateAttrs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attrs); i++ {

		if m.Attrs[i] != nil {
			if err := m.Attrs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attrs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attrs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NewbillingGetPricingResponse) contextValidateCatalog(ctx context.Context, formats strfmt.Registry) error {

	if m.Catalog != nil {
		if err := m.Catalog.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("catalog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("catalog")
			}
			return err
		}
	}

	return nil
}

func (m *NewbillingGetPricingResponse) contextValidateComponents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Components); i++ {

		if m.Components[i] != nil {
			if err := m.Components[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NewbillingGetPricingResponse) contextValidatePlan(ctx context.Context, formats strfmt.Registry) error {

	if m.Plan != nil {
		if err := m.Plan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plan")
			}
			return err
		}
	}

	return nil
}

func (m *NewbillingGetPricingResponse) contextValidateProduct(ctx context.Context, formats strfmt.Registry) error {

	if m.Product != nil {
		if err := m.Product.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("product")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("product")
			}
			return err
		}
	}

	return nil
}

func (m *NewbillingGetPricingResponse) contextValidatePubAttrFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PubAttrFilters); i++ {

		if m.PubAttrFilters[i] != nil {
			if err := m.PubAttrFilters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pub_attr_filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pub_attr_filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingGetPricingResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingGetPricingResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingGetPricingResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
