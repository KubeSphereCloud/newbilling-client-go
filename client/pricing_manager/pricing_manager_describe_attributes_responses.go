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

// PricingManagerDescribeAttributesReader is a Reader for the PricingManagerDescribeAttributes structure.
type PricingManagerDescribeAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingManagerDescribeAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingManagerDescribeAttributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPricingManagerDescribeAttributesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPricingManagerDescribeAttributesOK creates a PricingManagerDescribeAttributesOK with default headers values
func NewPricingManagerDescribeAttributesOK() *PricingManagerDescribeAttributesOK {
	return &PricingManagerDescribeAttributesOK{}
}

/*
PricingManagerDescribeAttributesOK describes a response with status code 200, with default header values.

A successful response.
*/
type PricingManagerDescribeAttributesOK struct {
	Payload *models.NewbillingDescribeAttributesResponse
}

// IsSuccess returns true when this pricing manager describe attributes o k response has a 2xx status code
func (o *PricingManagerDescribeAttributesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing manager describe attributes o k response has a 3xx status code
func (o *PricingManagerDescribeAttributesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing manager describe attributes o k response has a 4xx status code
func (o *PricingManagerDescribeAttributesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing manager describe attributes o k response has a 5xx status code
func (o *PricingManagerDescribeAttributesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing manager describe attributes o k response a status code equal to that given
func (o *PricingManagerDescribeAttributesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PricingManagerDescribeAttributesOK) Error() string {
	return fmt.Sprintf("[GET /v1/attributes][%d] pricingManagerDescribeAttributesOK  %+v", 200, o.Payload)
}

func (o *PricingManagerDescribeAttributesOK) String() string {
	return fmt.Sprintf("[GET /v1/attributes][%d] pricingManagerDescribeAttributesOK  %+v", 200, o.Payload)
}

func (o *PricingManagerDescribeAttributesOK) GetPayload() *models.NewbillingDescribeAttributesResponse {
	return o.Payload
}

func (o *PricingManagerDescribeAttributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingDescribeAttributesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPricingManagerDescribeAttributesDefault creates a PricingManagerDescribeAttributesDefault with default headers values
func NewPricingManagerDescribeAttributesDefault(code int) *PricingManagerDescribeAttributesDefault {
	return &PricingManagerDescribeAttributesDefault{
		_statusCode: code,
	}
}

/*
PricingManagerDescribeAttributesDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PricingManagerDescribeAttributesDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the pricing manager describe attributes default response
func (o *PricingManagerDescribeAttributesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this pricing manager describe attributes default response has a 2xx status code
func (o *PricingManagerDescribeAttributesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing manager describe attributes default response has a 3xx status code
func (o *PricingManagerDescribeAttributesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing manager describe attributes default response has a 4xx status code
func (o *PricingManagerDescribeAttributesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing manager describe attributes default response has a 5xx status code
func (o *PricingManagerDescribeAttributesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing manager describe attributes default response a status code equal to that given
func (o *PricingManagerDescribeAttributesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PricingManagerDescribeAttributesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/attributes][%d] PricingManager_DescribeAttributes default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerDescribeAttributesDefault) String() string {
	return fmt.Sprintf("[GET /v1/attributes][%d] PricingManager_DescribeAttributes default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerDescribeAttributesDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PricingManagerDescribeAttributesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
