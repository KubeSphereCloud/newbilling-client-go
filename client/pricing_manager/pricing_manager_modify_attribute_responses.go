// Code generated by go-swagger; DO NOT EDIT.

package pricing_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/newbilling-client-go/models"
)

// PricingManagerModifyAttributeReader is a Reader for the PricingManagerModifyAttribute structure.
type PricingManagerModifyAttributeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingManagerModifyAttributeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingManagerModifyAttributeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPricingManagerModifyAttributeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPricingManagerModifyAttributeOK creates a PricingManagerModifyAttributeOK with default headers values
func NewPricingManagerModifyAttributeOK() *PricingManagerModifyAttributeOK {
	return &PricingManagerModifyAttributeOK{}
}

/*
PricingManagerModifyAttributeOK describes a response with status code 200, with default header values.

A successful response.
*/
type PricingManagerModifyAttributeOK struct {
	Payload *models.NewbillingModifyAttributeResponse
}

// IsSuccess returns true when this pricing manager modify attribute o k response has a 2xx status code
func (o *PricingManagerModifyAttributeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing manager modify attribute o k response has a 3xx status code
func (o *PricingManagerModifyAttributeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing manager modify attribute o k response has a 4xx status code
func (o *PricingManagerModifyAttributeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing manager modify attribute o k response has a 5xx status code
func (o *PricingManagerModifyAttributeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing manager modify attribute o k response a status code equal to that given
func (o *PricingManagerModifyAttributeOK) IsCode(code int) bool {
	return code == 200
}

func (o *PricingManagerModifyAttributeOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/attributes][%d] pricingManagerModifyAttributeOK  %+v", 200, o.Payload)
}

func (o *PricingManagerModifyAttributeOK) String() string {
	return fmt.Sprintf("[PATCH /v1/attributes][%d] pricingManagerModifyAttributeOK  %+v", 200, o.Payload)
}

func (o *PricingManagerModifyAttributeOK) GetPayload() *models.NewbillingModifyAttributeResponse {
	return o.Payload
}

func (o *PricingManagerModifyAttributeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingModifyAttributeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPricingManagerModifyAttributeDefault creates a PricingManagerModifyAttributeDefault with default headers values
func NewPricingManagerModifyAttributeDefault(code int) *PricingManagerModifyAttributeDefault {
	return &PricingManagerModifyAttributeDefault{
		_statusCode: code,
	}
}

/*
PricingManagerModifyAttributeDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PricingManagerModifyAttributeDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the pricing manager modify attribute default response
func (o *PricingManagerModifyAttributeDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this pricing manager modify attribute default response has a 2xx status code
func (o *PricingManagerModifyAttributeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing manager modify attribute default response has a 3xx status code
func (o *PricingManagerModifyAttributeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing manager modify attribute default response has a 4xx status code
func (o *PricingManagerModifyAttributeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing manager modify attribute default response has a 5xx status code
func (o *PricingManagerModifyAttributeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing manager modify attribute default response a status code equal to that given
func (o *PricingManagerModifyAttributeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PricingManagerModifyAttributeDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1/attributes][%d] PricingManager_ModifyAttribute default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerModifyAttributeDefault) String() string {
	return fmt.Sprintf("[PATCH /v1/attributes][%d] PricingManager_ModifyAttribute default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerModifyAttributeDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PricingManagerModifyAttributeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}