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

// PricingAccessSystemManagerCreateAccessSystemReader is a Reader for the PricingAccessSystemManagerCreateAccessSystem structure.
type PricingAccessSystemManagerCreateAccessSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingAccessSystemManagerCreateAccessSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingAccessSystemManagerCreateAccessSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPricingAccessSystemManagerCreateAccessSystemDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPricingAccessSystemManagerCreateAccessSystemOK creates a PricingAccessSystemManagerCreateAccessSystemOK with default headers values
func NewPricingAccessSystemManagerCreateAccessSystemOK() *PricingAccessSystemManagerCreateAccessSystemOK {
	return &PricingAccessSystemManagerCreateAccessSystemOK{}
}

/*
PricingAccessSystemManagerCreateAccessSystemOK describes a response with status code 200, with default header values.

A successful response.
*/
type PricingAccessSystemManagerCreateAccessSystemOK struct {
	Payload *models.NewbillingCreateAccessSystemResponse
}

// IsSuccess returns true when this pricing access system manager create access system o k response has a 2xx status code
func (o *PricingAccessSystemManagerCreateAccessSystemOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing access system manager create access system o k response has a 3xx status code
func (o *PricingAccessSystemManagerCreateAccessSystemOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing access system manager create access system o k response has a 4xx status code
func (o *PricingAccessSystemManagerCreateAccessSystemOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing access system manager create access system o k response has a 5xx status code
func (o *PricingAccessSystemManagerCreateAccessSystemOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing access system manager create access system o k response a status code equal to that given
func (o *PricingAccessSystemManagerCreateAccessSystemOK) IsCode(code int) bool {
	return code == 200
}

func (o *PricingAccessSystemManagerCreateAccessSystemOK) Error() string {
	return fmt.Sprintf("[POST /v1/accesssystemcatalogs][%d] pricingAccessSystemManagerCreateAccessSystemOK  %+v", 200, o.Payload)
}

func (o *PricingAccessSystemManagerCreateAccessSystemOK) String() string {
	return fmt.Sprintf("[POST /v1/accesssystemcatalogs][%d] pricingAccessSystemManagerCreateAccessSystemOK  %+v", 200, o.Payload)
}

func (o *PricingAccessSystemManagerCreateAccessSystemOK) GetPayload() *models.NewbillingCreateAccessSystemResponse {
	return o.Payload
}

func (o *PricingAccessSystemManagerCreateAccessSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingCreateAccessSystemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPricingAccessSystemManagerCreateAccessSystemDefault creates a PricingAccessSystemManagerCreateAccessSystemDefault with default headers values
func NewPricingAccessSystemManagerCreateAccessSystemDefault(code int) *PricingAccessSystemManagerCreateAccessSystemDefault {
	return &PricingAccessSystemManagerCreateAccessSystemDefault{
		_statusCode: code,
	}
}

/*
PricingAccessSystemManagerCreateAccessSystemDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PricingAccessSystemManagerCreateAccessSystemDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the pricing access system manager create access system default response
func (o *PricingAccessSystemManagerCreateAccessSystemDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this pricing access system manager create access system default response has a 2xx status code
func (o *PricingAccessSystemManagerCreateAccessSystemDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing access system manager create access system default response has a 3xx status code
func (o *PricingAccessSystemManagerCreateAccessSystemDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing access system manager create access system default response has a 4xx status code
func (o *PricingAccessSystemManagerCreateAccessSystemDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing access system manager create access system default response has a 5xx status code
func (o *PricingAccessSystemManagerCreateAccessSystemDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing access system manager create access system default response a status code equal to that given
func (o *PricingAccessSystemManagerCreateAccessSystemDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PricingAccessSystemManagerCreateAccessSystemDefault) Error() string {
	return fmt.Sprintf("[POST /v1/accesssystemcatalogs][%d] PricingAccessSystemManager_CreateAccessSystem default  %+v", o._statusCode, o.Payload)
}

func (o *PricingAccessSystemManagerCreateAccessSystemDefault) String() string {
	return fmt.Sprintf("[POST /v1/accesssystemcatalogs][%d] PricingAccessSystemManager_CreateAccessSystem default  %+v", o._statusCode, o.Payload)
}

func (o *PricingAccessSystemManagerCreateAccessSystemDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PricingAccessSystemManagerCreateAccessSystemDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}