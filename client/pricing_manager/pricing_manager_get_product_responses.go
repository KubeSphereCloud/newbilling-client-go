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

// PricingManagerGetProductReader is a Reader for the PricingManagerGetProduct structure.
type PricingManagerGetProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingManagerGetProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingManagerGetProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPricingManagerGetProductDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPricingManagerGetProductOK creates a PricingManagerGetProductOK with default headers values
func NewPricingManagerGetProductOK() *PricingManagerGetProductOK {
	return &PricingManagerGetProductOK{}
}

/*
PricingManagerGetProductOK describes a response with status code 200, with default header values.

A successful response.
*/
type PricingManagerGetProductOK struct {
	Payload *models.NewbillingGetProductResponse
}

// IsSuccess returns true when this pricing manager get product o k response has a 2xx status code
func (o *PricingManagerGetProductOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing manager get product o k response has a 3xx status code
func (o *PricingManagerGetProductOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing manager get product o k response has a 4xx status code
func (o *PricingManagerGetProductOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing manager get product o k response has a 5xx status code
func (o *PricingManagerGetProductOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing manager get product o k response a status code equal to that given
func (o *PricingManagerGetProductOK) IsCode(code int) bool {
	return code == 200
}

func (o *PricingManagerGetProductOK) Error() string {
	return fmt.Sprintf("[GET /v1/products/{prod_id}][%d] pricingManagerGetProductOK  %+v", 200, o.Payload)
}

func (o *PricingManagerGetProductOK) String() string {
	return fmt.Sprintf("[GET /v1/products/{prod_id}][%d] pricingManagerGetProductOK  %+v", 200, o.Payload)
}

func (o *PricingManagerGetProductOK) GetPayload() *models.NewbillingGetProductResponse {
	return o.Payload
}

func (o *PricingManagerGetProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingGetProductResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPricingManagerGetProductDefault creates a PricingManagerGetProductDefault with default headers values
func NewPricingManagerGetProductDefault(code int) *PricingManagerGetProductDefault {
	return &PricingManagerGetProductDefault{
		_statusCode: code,
	}
}

/*
PricingManagerGetProductDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PricingManagerGetProductDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the pricing manager get product default response
func (o *PricingManagerGetProductDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this pricing manager get product default response has a 2xx status code
func (o *PricingManagerGetProductDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing manager get product default response has a 3xx status code
func (o *PricingManagerGetProductDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing manager get product default response has a 4xx status code
func (o *PricingManagerGetProductDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing manager get product default response has a 5xx status code
func (o *PricingManagerGetProductDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing manager get product default response a status code equal to that given
func (o *PricingManagerGetProductDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PricingManagerGetProductDefault) Error() string {
	return fmt.Sprintf("[GET /v1/products/{prod_id}][%d] PricingManager_GetProduct default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerGetProductDefault) String() string {
	return fmt.Sprintf("[GET /v1/products/{prod_id}][%d] PricingManager_GetProduct default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerGetProductDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PricingManagerGetProductDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
