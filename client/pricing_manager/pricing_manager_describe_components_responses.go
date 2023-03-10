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

// PricingManagerDescribeComponentsReader is a Reader for the PricingManagerDescribeComponents structure.
type PricingManagerDescribeComponentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingManagerDescribeComponentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingManagerDescribeComponentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPricingManagerDescribeComponentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPricingManagerDescribeComponentsOK creates a PricingManagerDescribeComponentsOK with default headers values
func NewPricingManagerDescribeComponentsOK() *PricingManagerDescribeComponentsOK {
	return &PricingManagerDescribeComponentsOK{}
}

/*
PricingManagerDescribeComponentsOK describes a response with status code 200, with default header values.

A successful response.
*/
type PricingManagerDescribeComponentsOK struct {
	Payload *models.NewbillingDescribeComponentsResponse
}

// IsSuccess returns true when this pricing manager describe components o k response has a 2xx status code
func (o *PricingManagerDescribeComponentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing manager describe components o k response has a 3xx status code
func (o *PricingManagerDescribeComponentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing manager describe components o k response has a 4xx status code
func (o *PricingManagerDescribeComponentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing manager describe components o k response has a 5xx status code
func (o *PricingManagerDescribeComponentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing manager describe components o k response a status code equal to that given
func (o *PricingManagerDescribeComponentsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PricingManagerDescribeComponentsOK) Error() string {
	return fmt.Sprintf("[GET /v1/components][%d] pricingManagerDescribeComponentsOK  %+v", 200, o.Payload)
}

func (o *PricingManagerDescribeComponentsOK) String() string {
	return fmt.Sprintf("[GET /v1/components][%d] pricingManagerDescribeComponentsOK  %+v", 200, o.Payload)
}

func (o *PricingManagerDescribeComponentsOK) GetPayload() *models.NewbillingDescribeComponentsResponse {
	return o.Payload
}

func (o *PricingManagerDescribeComponentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingDescribeComponentsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPricingManagerDescribeComponentsDefault creates a PricingManagerDescribeComponentsDefault with default headers values
func NewPricingManagerDescribeComponentsDefault(code int) *PricingManagerDescribeComponentsDefault {
	return &PricingManagerDescribeComponentsDefault{
		_statusCode: code,
	}
}

/*
PricingManagerDescribeComponentsDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PricingManagerDescribeComponentsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the pricing manager describe components default response
func (o *PricingManagerDescribeComponentsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this pricing manager describe components default response has a 2xx status code
func (o *PricingManagerDescribeComponentsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing manager describe components default response has a 3xx status code
func (o *PricingManagerDescribeComponentsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing manager describe components default response has a 4xx status code
func (o *PricingManagerDescribeComponentsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing manager describe components default response has a 5xx status code
func (o *PricingManagerDescribeComponentsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing manager describe components default response a status code equal to that given
func (o *PricingManagerDescribeComponentsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PricingManagerDescribeComponentsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/components][%d] PricingManager_DescribeComponents default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerDescribeComponentsDefault) String() string {
	return fmt.Sprintf("[GET /v1/components][%d] PricingManager_DescribeComponents default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerDescribeComponentsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PricingManagerDescribeComponentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
