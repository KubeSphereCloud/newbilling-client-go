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

// PricingManagerCreateProductReader is a Reader for the PricingManagerCreateProduct structure.
type PricingManagerCreateProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingManagerCreateProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingManagerCreateProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPricingManagerCreateProductDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPricingManagerCreateProductOK creates a PricingManagerCreateProductOK with default headers values
func NewPricingManagerCreateProductOK() *PricingManagerCreateProductOK {
	return &PricingManagerCreateProductOK{}
}

/*
PricingManagerCreateProductOK describes a response with status code 200, with default header values.

A successful response.
*/
type PricingManagerCreateProductOK struct {
	Payload *models.NewbillingCreateProductResponse
}

// IsSuccess returns true when this pricing manager create product o k response has a 2xx status code
func (o *PricingManagerCreateProductOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing manager create product o k response has a 3xx status code
func (o *PricingManagerCreateProductOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing manager create product o k response has a 4xx status code
func (o *PricingManagerCreateProductOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing manager create product o k response has a 5xx status code
func (o *PricingManagerCreateProductOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing manager create product o k response a status code equal to that given
func (o *PricingManagerCreateProductOK) IsCode(code int) bool {
	return code == 200
}

func (o *PricingManagerCreateProductOK) Error() string {
	return fmt.Sprintf("[POST /v1/products][%d] pricingManagerCreateProductOK  %+v", 200, o.Payload)
}

func (o *PricingManagerCreateProductOK) String() string {
	return fmt.Sprintf("[POST /v1/products][%d] pricingManagerCreateProductOK  %+v", 200, o.Payload)
}

func (o *PricingManagerCreateProductOK) GetPayload() *models.NewbillingCreateProductResponse {
	return o.Payload
}

func (o *PricingManagerCreateProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingCreateProductResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPricingManagerCreateProductDefault creates a PricingManagerCreateProductDefault with default headers values
func NewPricingManagerCreateProductDefault(code int) *PricingManagerCreateProductDefault {
	return &PricingManagerCreateProductDefault{
		_statusCode: code,
	}
}

/*
PricingManagerCreateProductDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PricingManagerCreateProductDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the pricing manager create product default response
func (o *PricingManagerCreateProductDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this pricing manager create product default response has a 2xx status code
func (o *PricingManagerCreateProductDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing manager create product default response has a 3xx status code
func (o *PricingManagerCreateProductDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing manager create product default response has a 4xx status code
func (o *PricingManagerCreateProductDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing manager create product default response has a 5xx status code
func (o *PricingManagerCreateProductDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing manager create product default response a status code equal to that given
func (o *PricingManagerCreateProductDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PricingManagerCreateProductDefault) Error() string {
	return fmt.Sprintf("[POST /v1/products][%d] PricingManager_CreateProduct default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerCreateProductDefault) String() string {
	return fmt.Sprintf("[POST /v1/products][%d] PricingManager_CreateProduct default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerCreateProductDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PricingManagerCreateProductDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}