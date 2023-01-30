// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingSendMailsResponse newbilling send mails response
//
// swagger:model newbillingSendMailsResponse
type NewbillingSendMailsResponse struct {

	// 邮件发送是否成功
	IsSucc bool `json:"is_succ,omitempty"`
}

// Validate validates this newbilling send mails response
func (m *NewbillingSendMailsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling send mails response based on context it is used
func (m *NewbillingSendMailsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingSendMailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingSendMailsResponse) UnmarshalBinary(b []byte) error {
	var res NewbillingSendMailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}