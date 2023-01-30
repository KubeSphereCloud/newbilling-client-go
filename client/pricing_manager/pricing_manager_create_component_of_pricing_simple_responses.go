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

// PricingManagerCreateComponentOfPricingSimpleReader is a Reader for the PricingManagerCreateComponentOfPricingSimple structure.
type PricingManagerCreateComponentOfPricingSimpleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingManagerCreateComponentOfPricingSimpleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingManagerCreateComponentOfPricingSimpleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPricingManagerCreateComponentOfPricingSimpleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPricingManagerCreateComponentOfPricingSimpleOK creates a PricingManagerCreateComponentOfPricingSimpleOK with default headers values
func NewPricingManagerCreateComponentOfPricingSimpleOK() *PricingManagerCreateComponentOfPricingSimpleOK {
	return &PricingManagerCreateComponentOfPricingSimpleOK{}
}

/*
PricingManagerCreateComponentOfPricingSimpleOK describes a response with status code 200, with default header values.

A successful response.
*/
type PricingManagerCreateComponentOfPricingSimpleOK struct {
	Payload *models.NewbillingComponentOfPricingSimpleResponse
}

// IsSuccess returns true when this pricing manager create component of pricing simple o k response has a 2xx status code
func (o *PricingManagerCreateComponentOfPricingSimpleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing manager create component of pricing simple o k response has a 3xx status code
func (o *PricingManagerCreateComponentOfPricingSimpleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing manager create component of pricing simple o k response has a 4xx status code
func (o *PricingManagerCreateComponentOfPricingSimpleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing manager create component of pricing simple o k response has a 5xx status code
func (o *PricingManagerCreateComponentOfPricingSimpleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing manager create component of pricing simple o k response a status code equal to that given
func (o *PricingManagerCreateComponentOfPricingSimpleOK) IsCode(code int) bool {
	return code == 200
}

func (o *PricingManagerCreateComponentOfPricingSimpleOK) Error() string {
	return fmt.Sprintf("[POST /v1/components:pricing_simple][%d] pricingManagerCreateComponentOfPricingSimpleOK  %+v", 200, o.Payload)
}

func (o *PricingManagerCreateComponentOfPricingSimpleOK) String() string {
	return fmt.Sprintf("[POST /v1/components:pricing_simple][%d] pricingManagerCreateComponentOfPricingSimpleOK  %+v", 200, o.Payload)
}

func (o *PricingManagerCreateComponentOfPricingSimpleOK) GetPayload() *models.NewbillingComponentOfPricingSimpleResponse {
	return o.Payload
}

func (o *PricingManagerCreateComponentOfPricingSimpleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingComponentOfPricingSimpleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPricingManagerCreateComponentOfPricingSimpleDefault creates a PricingManagerCreateComponentOfPricingSimpleDefault with default headers values
func NewPricingManagerCreateComponentOfPricingSimpleDefault(code int) *PricingManagerCreateComponentOfPricingSimpleDefault {
	return &PricingManagerCreateComponentOfPricingSimpleDefault{
		_statusCode: code,
	}
}

/*
PricingManagerCreateComponentOfPricingSimpleDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PricingManagerCreateComponentOfPricingSimpleDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the pricing manager create component of pricing simple default response
func (o *PricingManagerCreateComponentOfPricingSimpleDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this pricing manager create component of pricing simple default response has a 2xx status code
func (o *PricingManagerCreateComponentOfPricingSimpleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing manager create component of pricing simple default response has a 3xx status code
func (o *PricingManagerCreateComponentOfPricingSimpleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing manager create component of pricing simple default response has a 4xx status code
func (o *PricingManagerCreateComponentOfPricingSimpleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing manager create component of pricing simple default response has a 5xx status code
func (o *PricingManagerCreateComponentOfPricingSimpleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing manager create component of pricing simple default response a status code equal to that given
func (o *PricingManagerCreateComponentOfPricingSimpleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PricingManagerCreateComponentOfPricingSimpleDefault) Error() string {
	return fmt.Sprintf("[POST /v1/components:pricing_simple][%d] PricingManager_CreateComponentOfPricingSimple default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerCreateComponentOfPricingSimpleDefault) String() string {
	return fmt.Sprintf("[POST /v1/components:pricing_simple][%d] PricingManager_CreateComponentOfPricingSimple default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerCreateComponentOfPricingSimpleDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PricingManagerCreateComponentOfPricingSimpleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}