// Code generated by go-swagger; DO NOT EDIT.

package pricing_cost_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/newbilling-client-go/models"
)

// PricingCostManagerCalculateComponentsPriceReader is a Reader for the PricingCostManagerCalculateComponentsPrice structure.
type PricingCostManagerCalculateComponentsPriceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingCostManagerCalculateComponentsPriceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingCostManagerCalculateComponentsPriceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPricingCostManagerCalculateComponentsPriceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPricingCostManagerCalculateComponentsPriceOK creates a PricingCostManagerCalculateComponentsPriceOK with default headers values
func NewPricingCostManagerCalculateComponentsPriceOK() *PricingCostManagerCalculateComponentsPriceOK {
	return &PricingCostManagerCalculateComponentsPriceOK{}
}

/*
PricingCostManagerCalculateComponentsPriceOK describes a response with status code 200, with default header values.

A successful response.
*/
type PricingCostManagerCalculateComponentsPriceOK struct {
	Payload *models.NewbillingCalculatePriceResponse
}

// IsSuccess returns true when this pricing cost manager calculate components price o k response has a 2xx status code
func (o *PricingCostManagerCalculateComponentsPriceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing cost manager calculate components price o k response has a 3xx status code
func (o *PricingCostManagerCalculateComponentsPriceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing cost manager calculate components price o k response has a 4xx status code
func (o *PricingCostManagerCalculateComponentsPriceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing cost manager calculate components price o k response has a 5xx status code
func (o *PricingCostManagerCalculateComponentsPriceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing cost manager calculate components price o k response a status code equal to that given
func (o *PricingCostManagerCalculateComponentsPriceOK) IsCode(code int) bool {
	return code == 200
}

func (o *PricingCostManagerCalculateComponentsPriceOK) Error() string {
	return fmt.Sprintf("[POST /v1/components/cost][%d] pricingCostManagerCalculateComponentsPriceOK  %+v", 200, o.Payload)
}

func (o *PricingCostManagerCalculateComponentsPriceOK) String() string {
	return fmt.Sprintf("[POST /v1/components/cost][%d] pricingCostManagerCalculateComponentsPriceOK  %+v", 200, o.Payload)
}

func (o *PricingCostManagerCalculateComponentsPriceOK) GetPayload() *models.NewbillingCalculatePriceResponse {
	return o.Payload
}

func (o *PricingCostManagerCalculateComponentsPriceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingCalculatePriceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPricingCostManagerCalculateComponentsPriceDefault creates a PricingCostManagerCalculateComponentsPriceDefault with default headers values
func NewPricingCostManagerCalculateComponentsPriceDefault(code int) *PricingCostManagerCalculateComponentsPriceDefault {
	return &PricingCostManagerCalculateComponentsPriceDefault{
		_statusCode: code,
	}
}

/*
PricingCostManagerCalculateComponentsPriceDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PricingCostManagerCalculateComponentsPriceDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the pricing cost manager calculate components price default response
func (o *PricingCostManagerCalculateComponentsPriceDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this pricing cost manager calculate components price default response has a 2xx status code
func (o *PricingCostManagerCalculateComponentsPriceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing cost manager calculate components price default response has a 3xx status code
func (o *PricingCostManagerCalculateComponentsPriceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing cost manager calculate components price default response has a 4xx status code
func (o *PricingCostManagerCalculateComponentsPriceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing cost manager calculate components price default response has a 5xx status code
func (o *PricingCostManagerCalculateComponentsPriceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing cost manager calculate components price default response a status code equal to that given
func (o *PricingCostManagerCalculateComponentsPriceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PricingCostManagerCalculateComponentsPriceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/components/cost][%d] PricingCostManager_CalculateComponentsPrice default  %+v", o._statusCode, o.Payload)
}

func (o *PricingCostManagerCalculateComponentsPriceDefault) String() string {
	return fmt.Sprintf("[POST /v1/components/cost][%d] PricingCostManager_CalculateComponentsPrice default  %+v", o._statusCode, o.Payload)
}

func (o *PricingCostManagerCalculateComponentsPriceDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PricingCostManagerCalculateComponentsPriceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
