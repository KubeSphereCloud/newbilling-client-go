// Code generated by go-swagger; DO NOT EDIT.

package pricing_access_system_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/newbilling-client-go/models"
)

// PricingAccessSystemManagerModifyAccessSystemReader is a Reader for the PricingAccessSystemManagerModifyAccessSystem structure.
type PricingAccessSystemManagerModifyAccessSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingAccessSystemManagerModifyAccessSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingAccessSystemManagerModifyAccessSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPricingAccessSystemManagerModifyAccessSystemDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPricingAccessSystemManagerModifyAccessSystemOK creates a PricingAccessSystemManagerModifyAccessSystemOK with default headers values
func NewPricingAccessSystemManagerModifyAccessSystemOK() *PricingAccessSystemManagerModifyAccessSystemOK {
	return &PricingAccessSystemManagerModifyAccessSystemOK{}
}

/*
PricingAccessSystemManagerModifyAccessSystemOK describes a response with status code 200, with default header values.

A successful response.
*/
type PricingAccessSystemManagerModifyAccessSystemOK struct {
	Payload *models.NewbillingModifyAccessSystemResponse
}

// IsSuccess returns true when this pricing access system manager modify access system o k response has a 2xx status code
func (o *PricingAccessSystemManagerModifyAccessSystemOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing access system manager modify access system o k response has a 3xx status code
func (o *PricingAccessSystemManagerModifyAccessSystemOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing access system manager modify access system o k response has a 4xx status code
func (o *PricingAccessSystemManagerModifyAccessSystemOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing access system manager modify access system o k response has a 5xx status code
func (o *PricingAccessSystemManagerModifyAccessSystemOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing access system manager modify access system o k response a status code equal to that given
func (o *PricingAccessSystemManagerModifyAccessSystemOK) IsCode(code int) bool {
	return code == 200
}

func (o *PricingAccessSystemManagerModifyAccessSystemOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/accesssystems][%d] pricingAccessSystemManagerModifyAccessSystemOK  %+v", 200, o.Payload)
}

func (o *PricingAccessSystemManagerModifyAccessSystemOK) String() string {
	return fmt.Sprintf("[PATCH /v1/accesssystems][%d] pricingAccessSystemManagerModifyAccessSystemOK  %+v", 200, o.Payload)
}

func (o *PricingAccessSystemManagerModifyAccessSystemOK) GetPayload() *models.NewbillingModifyAccessSystemResponse {
	return o.Payload
}

func (o *PricingAccessSystemManagerModifyAccessSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingModifyAccessSystemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPricingAccessSystemManagerModifyAccessSystemDefault creates a PricingAccessSystemManagerModifyAccessSystemDefault with default headers values
func NewPricingAccessSystemManagerModifyAccessSystemDefault(code int) *PricingAccessSystemManagerModifyAccessSystemDefault {
	return &PricingAccessSystemManagerModifyAccessSystemDefault{
		_statusCode: code,
	}
}

/*
PricingAccessSystemManagerModifyAccessSystemDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PricingAccessSystemManagerModifyAccessSystemDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the pricing access system manager modify access system default response
func (o *PricingAccessSystemManagerModifyAccessSystemDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this pricing access system manager modify access system default response has a 2xx status code
func (o *PricingAccessSystemManagerModifyAccessSystemDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing access system manager modify access system default response has a 3xx status code
func (o *PricingAccessSystemManagerModifyAccessSystemDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing access system manager modify access system default response has a 4xx status code
func (o *PricingAccessSystemManagerModifyAccessSystemDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing access system manager modify access system default response has a 5xx status code
func (o *PricingAccessSystemManagerModifyAccessSystemDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing access system manager modify access system default response a status code equal to that given
func (o *PricingAccessSystemManagerModifyAccessSystemDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PricingAccessSystemManagerModifyAccessSystemDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1/accesssystems][%d] PricingAccessSystemManager_ModifyAccessSystem default  %+v", o._statusCode, o.Payload)
}

func (o *PricingAccessSystemManagerModifyAccessSystemDefault) String() string {
	return fmt.Sprintf("[PATCH /v1/accesssystems][%d] PricingAccessSystemManager_ModifyAccessSystem default  %+v", o._statusCode, o.Payload)
}

func (o *PricingAccessSystemManagerModifyAccessSystemDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PricingAccessSystemManagerModifyAccessSystemDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
