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

// PricingManagerModifyComponentReader is a Reader for the PricingManagerModifyComponent structure.
type PricingManagerModifyComponentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingManagerModifyComponentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingManagerModifyComponentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPricingManagerModifyComponentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPricingManagerModifyComponentOK creates a PricingManagerModifyComponentOK with default headers values
func NewPricingManagerModifyComponentOK() *PricingManagerModifyComponentOK {
	return &PricingManagerModifyComponentOK{}
}

/*
PricingManagerModifyComponentOK describes a response with status code 200, with default header values.

A successful response.
*/
type PricingManagerModifyComponentOK struct {
	Payload *models.NewbillingModifyComponentResponse
}

// IsSuccess returns true when this pricing manager modify component o k response has a 2xx status code
func (o *PricingManagerModifyComponentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing manager modify component o k response has a 3xx status code
func (o *PricingManagerModifyComponentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing manager modify component o k response has a 4xx status code
func (o *PricingManagerModifyComponentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing manager modify component o k response has a 5xx status code
func (o *PricingManagerModifyComponentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing manager modify component o k response a status code equal to that given
func (o *PricingManagerModifyComponentOK) IsCode(code int) bool {
	return code == 200
}

func (o *PricingManagerModifyComponentOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/components][%d] pricingManagerModifyComponentOK  %+v", 200, o.Payload)
}

func (o *PricingManagerModifyComponentOK) String() string {
	return fmt.Sprintf("[PATCH /v1/components][%d] pricingManagerModifyComponentOK  %+v", 200, o.Payload)
}

func (o *PricingManagerModifyComponentOK) GetPayload() *models.NewbillingModifyComponentResponse {
	return o.Payload
}

func (o *PricingManagerModifyComponentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingModifyComponentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPricingManagerModifyComponentDefault creates a PricingManagerModifyComponentDefault with default headers values
func NewPricingManagerModifyComponentDefault(code int) *PricingManagerModifyComponentDefault {
	return &PricingManagerModifyComponentDefault{
		_statusCode: code,
	}
}

/*
PricingManagerModifyComponentDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PricingManagerModifyComponentDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the pricing manager modify component default response
func (o *PricingManagerModifyComponentDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this pricing manager modify component default response has a 2xx status code
func (o *PricingManagerModifyComponentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing manager modify component default response has a 3xx status code
func (o *PricingManagerModifyComponentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing manager modify component default response has a 4xx status code
func (o *PricingManagerModifyComponentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing manager modify component default response has a 5xx status code
func (o *PricingManagerModifyComponentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing manager modify component default response a status code equal to that given
func (o *PricingManagerModifyComponentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PricingManagerModifyComponentDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1/components][%d] PricingManager_ModifyComponent default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerModifyComponentDefault) String() string {
	return fmt.Sprintf("[PATCH /v1/components][%d] PricingManager_ModifyComponent default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerModifyComponentDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PricingManagerModifyComponentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}