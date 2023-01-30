// Code generated by go-swagger; DO NOT EDIT.

package pricing_access_system_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/newbilling-client-go/models"
)

// PricingAccessSystemManagerDeleteAccessSystemsReader is a Reader for the PricingAccessSystemManagerDeleteAccessSystems structure.
type PricingAccessSystemManagerDeleteAccessSystemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingAccessSystemManagerDeleteAccessSystemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingAccessSystemManagerDeleteAccessSystemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPricingAccessSystemManagerDeleteAccessSystemsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPricingAccessSystemManagerDeleteAccessSystemsOK creates a PricingAccessSystemManagerDeleteAccessSystemsOK with default headers values
func NewPricingAccessSystemManagerDeleteAccessSystemsOK() *PricingAccessSystemManagerDeleteAccessSystemsOK {
	return &PricingAccessSystemManagerDeleteAccessSystemsOK{}
}

/*
PricingAccessSystemManagerDeleteAccessSystemsOK describes a response with status code 200, with default header values.

A successful response.
*/
type PricingAccessSystemManagerDeleteAccessSystemsOK struct {
	Payload *models.NewbillingDeleteAccessSystemsResponse
}

// IsSuccess returns true when this pricing access system manager delete access systems o k response has a 2xx status code
func (o *PricingAccessSystemManagerDeleteAccessSystemsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing access system manager delete access systems o k response has a 3xx status code
func (o *PricingAccessSystemManagerDeleteAccessSystemsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing access system manager delete access systems o k response has a 4xx status code
func (o *PricingAccessSystemManagerDeleteAccessSystemsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing access system manager delete access systems o k response has a 5xx status code
func (o *PricingAccessSystemManagerDeleteAccessSystemsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing access system manager delete access systems o k response a status code equal to that given
func (o *PricingAccessSystemManagerDeleteAccessSystemsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PricingAccessSystemManagerDeleteAccessSystemsOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/accesssystems][%d] pricingAccessSystemManagerDeleteAccessSystemsOK  %+v", 200, o.Payload)
}

func (o *PricingAccessSystemManagerDeleteAccessSystemsOK) String() string {
	return fmt.Sprintf("[DELETE /v1/accesssystems][%d] pricingAccessSystemManagerDeleteAccessSystemsOK  %+v", 200, o.Payload)
}

func (o *PricingAccessSystemManagerDeleteAccessSystemsOK) GetPayload() *models.NewbillingDeleteAccessSystemsResponse {
	return o.Payload
}

func (o *PricingAccessSystemManagerDeleteAccessSystemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingDeleteAccessSystemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPricingAccessSystemManagerDeleteAccessSystemsDefault creates a PricingAccessSystemManagerDeleteAccessSystemsDefault with default headers values
func NewPricingAccessSystemManagerDeleteAccessSystemsDefault(code int) *PricingAccessSystemManagerDeleteAccessSystemsDefault {
	return &PricingAccessSystemManagerDeleteAccessSystemsDefault{
		_statusCode: code,
	}
}

/*
PricingAccessSystemManagerDeleteAccessSystemsDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PricingAccessSystemManagerDeleteAccessSystemsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the pricing access system manager delete access systems default response
func (o *PricingAccessSystemManagerDeleteAccessSystemsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this pricing access system manager delete access systems default response has a 2xx status code
func (o *PricingAccessSystemManagerDeleteAccessSystemsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing access system manager delete access systems default response has a 3xx status code
func (o *PricingAccessSystemManagerDeleteAccessSystemsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing access system manager delete access systems default response has a 4xx status code
func (o *PricingAccessSystemManagerDeleteAccessSystemsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing access system manager delete access systems default response has a 5xx status code
func (o *PricingAccessSystemManagerDeleteAccessSystemsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing access system manager delete access systems default response a status code equal to that given
func (o *PricingAccessSystemManagerDeleteAccessSystemsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PricingAccessSystemManagerDeleteAccessSystemsDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/accesssystems][%d] PricingAccessSystemManager_DeleteAccessSystems default  %+v", o._statusCode, o.Payload)
}

func (o *PricingAccessSystemManagerDeleteAccessSystemsDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/accesssystems][%d] PricingAccessSystemManager_DeleteAccessSystems default  %+v", o._statusCode, o.Payload)
}

func (o *PricingAccessSystemManagerDeleteAccessSystemsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PricingAccessSystemManagerDeleteAccessSystemsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}