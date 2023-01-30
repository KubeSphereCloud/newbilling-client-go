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

// PricingManagerModifyStrategyReader is a Reader for the PricingManagerModifyStrategy structure.
type PricingManagerModifyStrategyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingManagerModifyStrategyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingManagerModifyStrategyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPricingManagerModifyStrategyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPricingManagerModifyStrategyOK creates a PricingManagerModifyStrategyOK with default headers values
func NewPricingManagerModifyStrategyOK() *PricingManagerModifyStrategyOK {
	return &PricingManagerModifyStrategyOK{}
}

/*
PricingManagerModifyStrategyOK describes a response with status code 200, with default header values.

A successful response.
*/
type PricingManagerModifyStrategyOK struct {
	Payload *models.NewbillingModifyStrategyResponse
}

// IsSuccess returns true when this pricing manager modify strategy o k response has a 2xx status code
func (o *PricingManagerModifyStrategyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing manager modify strategy o k response has a 3xx status code
func (o *PricingManagerModifyStrategyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing manager modify strategy o k response has a 4xx status code
func (o *PricingManagerModifyStrategyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing manager modify strategy o k response has a 5xx status code
func (o *PricingManagerModifyStrategyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing manager modify strategy o k response a status code equal to that given
func (o *PricingManagerModifyStrategyOK) IsCode(code int) bool {
	return code == 200
}

func (o *PricingManagerModifyStrategyOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/strategies][%d] pricingManagerModifyStrategyOK  %+v", 200, o.Payload)
}

func (o *PricingManagerModifyStrategyOK) String() string {
	return fmt.Sprintf("[PATCH /v1/strategies][%d] pricingManagerModifyStrategyOK  %+v", 200, o.Payload)
}

func (o *PricingManagerModifyStrategyOK) GetPayload() *models.NewbillingModifyStrategyResponse {
	return o.Payload
}

func (o *PricingManagerModifyStrategyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingModifyStrategyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPricingManagerModifyStrategyDefault creates a PricingManagerModifyStrategyDefault with default headers values
func NewPricingManagerModifyStrategyDefault(code int) *PricingManagerModifyStrategyDefault {
	return &PricingManagerModifyStrategyDefault{
		_statusCode: code,
	}
}

/*
PricingManagerModifyStrategyDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PricingManagerModifyStrategyDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the pricing manager modify strategy default response
func (o *PricingManagerModifyStrategyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this pricing manager modify strategy default response has a 2xx status code
func (o *PricingManagerModifyStrategyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing manager modify strategy default response has a 3xx status code
func (o *PricingManagerModifyStrategyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing manager modify strategy default response has a 4xx status code
func (o *PricingManagerModifyStrategyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing manager modify strategy default response has a 5xx status code
func (o *PricingManagerModifyStrategyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing manager modify strategy default response a status code equal to that given
func (o *PricingManagerModifyStrategyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PricingManagerModifyStrategyDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1/strategies][%d] PricingManager_ModifyStrategy default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerModifyStrategyDefault) String() string {
	return fmt.Sprintf("[PATCH /v1/strategies][%d] PricingManager_ModifyStrategy default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerModifyStrategyDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PricingManagerModifyStrategyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}