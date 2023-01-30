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

// PricingManagerDeletePlansReader is a Reader for the PricingManagerDeletePlans structure.
type PricingManagerDeletePlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingManagerDeletePlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingManagerDeletePlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPricingManagerDeletePlansDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPricingManagerDeletePlansOK creates a PricingManagerDeletePlansOK with default headers values
func NewPricingManagerDeletePlansOK() *PricingManagerDeletePlansOK {
	return &PricingManagerDeletePlansOK{}
}

/*
PricingManagerDeletePlansOK describes a response with status code 200, with default header values.

A successful response.
*/
type PricingManagerDeletePlansOK struct {
	Payload *models.NewbillingDeletePlansResponse
}

// IsSuccess returns true when this pricing manager delete plans o k response has a 2xx status code
func (o *PricingManagerDeletePlansOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing manager delete plans o k response has a 3xx status code
func (o *PricingManagerDeletePlansOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing manager delete plans o k response has a 4xx status code
func (o *PricingManagerDeletePlansOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing manager delete plans o k response has a 5xx status code
func (o *PricingManagerDeletePlansOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing manager delete plans o k response a status code equal to that given
func (o *PricingManagerDeletePlansOK) IsCode(code int) bool {
	return code == 200
}

func (o *PricingManagerDeletePlansOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/plans][%d] pricingManagerDeletePlansOK  %+v", 200, o.Payload)
}

func (o *PricingManagerDeletePlansOK) String() string {
	return fmt.Sprintf("[DELETE /v1/plans][%d] pricingManagerDeletePlansOK  %+v", 200, o.Payload)
}

func (o *PricingManagerDeletePlansOK) GetPayload() *models.NewbillingDeletePlansResponse {
	return o.Payload
}

func (o *PricingManagerDeletePlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingDeletePlansResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPricingManagerDeletePlansDefault creates a PricingManagerDeletePlansDefault with default headers values
func NewPricingManagerDeletePlansDefault(code int) *PricingManagerDeletePlansDefault {
	return &PricingManagerDeletePlansDefault{
		_statusCode: code,
	}
}

/*
PricingManagerDeletePlansDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PricingManagerDeletePlansDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the pricing manager delete plans default response
func (o *PricingManagerDeletePlansDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this pricing manager delete plans default response has a 2xx status code
func (o *PricingManagerDeletePlansDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing manager delete plans default response has a 3xx status code
func (o *PricingManagerDeletePlansDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing manager delete plans default response has a 4xx status code
func (o *PricingManagerDeletePlansDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing manager delete plans default response has a 5xx status code
func (o *PricingManagerDeletePlansDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing manager delete plans default response a status code equal to that given
func (o *PricingManagerDeletePlansDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PricingManagerDeletePlansDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/plans][%d] PricingManager_DeletePlans default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerDeletePlansDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/plans][%d] PricingManager_DeletePlans default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerDeletePlansDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PricingManagerDeletePlansDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}