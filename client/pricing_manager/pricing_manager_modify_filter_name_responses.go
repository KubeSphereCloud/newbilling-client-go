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

// PricingManagerModifyFilterNameReader is a Reader for the PricingManagerModifyFilterName structure.
type PricingManagerModifyFilterNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingManagerModifyFilterNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingManagerModifyFilterNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPricingManagerModifyFilterNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPricingManagerModifyFilterNameOK creates a PricingManagerModifyFilterNameOK with default headers values
func NewPricingManagerModifyFilterNameOK() *PricingManagerModifyFilterNameOK {
	return &PricingManagerModifyFilterNameOK{}
}

/*
PricingManagerModifyFilterNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type PricingManagerModifyFilterNameOK struct {
	Payload *models.NewbillingModifyFilterNameResponse
}

// IsSuccess returns true when this pricing manager modify filter name o k response has a 2xx status code
func (o *PricingManagerModifyFilterNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing manager modify filter name o k response has a 3xx status code
func (o *PricingManagerModifyFilterNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing manager modify filter name o k response has a 4xx status code
func (o *PricingManagerModifyFilterNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing manager modify filter name o k response has a 5xx status code
func (o *PricingManagerModifyFilterNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing manager modify filter name o k response a status code equal to that given
func (o *PricingManagerModifyFilterNameOK) IsCode(code int) bool {
	return code == 200
}

func (o *PricingManagerModifyFilterNameOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/filters/name][%d] pricingManagerModifyFilterNameOK  %+v", 200, o.Payload)
}

func (o *PricingManagerModifyFilterNameOK) String() string {
	return fmt.Sprintf("[PATCH /v1/filters/name][%d] pricingManagerModifyFilterNameOK  %+v", 200, o.Payload)
}

func (o *PricingManagerModifyFilterNameOK) GetPayload() *models.NewbillingModifyFilterNameResponse {
	return o.Payload
}

func (o *PricingManagerModifyFilterNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingModifyFilterNameResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPricingManagerModifyFilterNameDefault creates a PricingManagerModifyFilterNameDefault with default headers values
func NewPricingManagerModifyFilterNameDefault(code int) *PricingManagerModifyFilterNameDefault {
	return &PricingManagerModifyFilterNameDefault{
		_statusCode: code,
	}
}

/*
PricingManagerModifyFilterNameDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PricingManagerModifyFilterNameDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the pricing manager modify filter name default response
func (o *PricingManagerModifyFilterNameDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this pricing manager modify filter name default response has a 2xx status code
func (o *PricingManagerModifyFilterNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing manager modify filter name default response has a 3xx status code
func (o *PricingManagerModifyFilterNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing manager modify filter name default response has a 4xx status code
func (o *PricingManagerModifyFilterNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing manager modify filter name default response has a 5xx status code
func (o *PricingManagerModifyFilterNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing manager modify filter name default response a status code equal to that given
func (o *PricingManagerModifyFilterNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PricingManagerModifyFilterNameDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1/filters/name][%d] PricingManager_ModifyFilterName default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerModifyFilterNameDefault) String() string {
	return fmt.Sprintf("[PATCH /v1/filters/name][%d] PricingManager_ModifyFilterName default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerModifyFilterNameDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PricingManagerModifyFilterNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
