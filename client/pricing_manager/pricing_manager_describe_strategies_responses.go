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

// PricingManagerDescribeStrategiesReader is a Reader for the PricingManagerDescribeStrategies structure.
type PricingManagerDescribeStrategiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingManagerDescribeStrategiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingManagerDescribeStrategiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPricingManagerDescribeStrategiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPricingManagerDescribeStrategiesOK creates a PricingManagerDescribeStrategiesOK with default headers values
func NewPricingManagerDescribeStrategiesOK() *PricingManagerDescribeStrategiesOK {
	return &PricingManagerDescribeStrategiesOK{}
}

/*
PricingManagerDescribeStrategiesOK describes a response with status code 200, with default header values.

A successful response.
*/
type PricingManagerDescribeStrategiesOK struct {
	Payload *models.NewbillingDescribeStrategiesResponse
}

// IsSuccess returns true when this pricing manager describe strategies o k response has a 2xx status code
func (o *PricingManagerDescribeStrategiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing manager describe strategies o k response has a 3xx status code
func (o *PricingManagerDescribeStrategiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing manager describe strategies o k response has a 4xx status code
func (o *PricingManagerDescribeStrategiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing manager describe strategies o k response has a 5xx status code
func (o *PricingManagerDescribeStrategiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing manager describe strategies o k response a status code equal to that given
func (o *PricingManagerDescribeStrategiesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PricingManagerDescribeStrategiesOK) Error() string {
	return fmt.Sprintf("[GET /v1/strategies][%d] pricingManagerDescribeStrategiesOK  %+v", 200, o.Payload)
}

func (o *PricingManagerDescribeStrategiesOK) String() string {
	return fmt.Sprintf("[GET /v1/strategies][%d] pricingManagerDescribeStrategiesOK  %+v", 200, o.Payload)
}

func (o *PricingManagerDescribeStrategiesOK) GetPayload() *models.NewbillingDescribeStrategiesResponse {
	return o.Payload
}

func (o *PricingManagerDescribeStrategiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingDescribeStrategiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPricingManagerDescribeStrategiesDefault creates a PricingManagerDescribeStrategiesDefault with default headers values
func NewPricingManagerDescribeStrategiesDefault(code int) *PricingManagerDescribeStrategiesDefault {
	return &PricingManagerDescribeStrategiesDefault{
		_statusCode: code,
	}
}

/*
PricingManagerDescribeStrategiesDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PricingManagerDescribeStrategiesDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the pricing manager describe strategies default response
func (o *PricingManagerDescribeStrategiesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this pricing manager describe strategies default response has a 2xx status code
func (o *PricingManagerDescribeStrategiesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing manager describe strategies default response has a 3xx status code
func (o *PricingManagerDescribeStrategiesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing manager describe strategies default response has a 4xx status code
func (o *PricingManagerDescribeStrategiesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing manager describe strategies default response has a 5xx status code
func (o *PricingManagerDescribeStrategiesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing manager describe strategies default response a status code equal to that given
func (o *PricingManagerDescribeStrategiesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PricingManagerDescribeStrategiesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/strategies][%d] PricingManager_DescribeStrategies default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerDescribeStrategiesDefault) String() string {
	return fmt.Sprintf("[GET /v1/strategies][%d] PricingManager_DescribeStrategies default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerDescribeStrategiesDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PricingManagerDescribeStrategiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
