// Code generated by go-swagger; DO NOT EDIT.

package subscription_package_order_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/newbilling-client-go/models"
)

// SubscriptionPackageOrderManagerChargePrdOrderReader is a Reader for the SubscriptionPackageOrderManagerChargePrdOrder structure.
type SubscriptionPackageOrderManagerChargePrdOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriptionPackageOrderManagerChargePrdOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscriptionPackageOrderManagerChargePrdOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSubscriptionPackageOrderManagerChargePrdOrderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSubscriptionPackageOrderManagerChargePrdOrderOK creates a SubscriptionPackageOrderManagerChargePrdOrderOK with default headers values
func NewSubscriptionPackageOrderManagerChargePrdOrderOK() *SubscriptionPackageOrderManagerChargePrdOrderOK {
	return &SubscriptionPackageOrderManagerChargePrdOrderOK{}
}

/*
SubscriptionPackageOrderManagerChargePrdOrderOK describes a response with status code 200, with default header values.

A successful response.
*/
type SubscriptionPackageOrderManagerChargePrdOrderOK struct {
	Payload *models.NewbillingChargePrdOrderResponse
}

// IsSuccess returns true when this subscription package order manager charge prd order o k response has a 2xx status code
func (o *SubscriptionPackageOrderManagerChargePrdOrderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this subscription package order manager charge prd order o k response has a 3xx status code
func (o *SubscriptionPackageOrderManagerChargePrdOrderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription package order manager charge prd order o k response has a 4xx status code
func (o *SubscriptionPackageOrderManagerChargePrdOrderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription package order manager charge prd order o k response has a 5xx status code
func (o *SubscriptionPackageOrderManagerChargePrdOrderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription package order manager charge prd order o k response a status code equal to that given
func (o *SubscriptionPackageOrderManagerChargePrdOrderOK) IsCode(code int) bool {
	return code == 200
}

func (o *SubscriptionPackageOrderManagerChargePrdOrderOK) Error() string {
	return fmt.Sprintf("[POST /v1/packageorders/{order_id}:pay][%d] subscriptionPackageOrderManagerChargePrdOrderOK  %+v", 200, o.Payload)
}

func (o *SubscriptionPackageOrderManagerChargePrdOrderOK) String() string {
	return fmt.Sprintf("[POST /v1/packageorders/{order_id}:pay][%d] subscriptionPackageOrderManagerChargePrdOrderOK  %+v", 200, o.Payload)
}

func (o *SubscriptionPackageOrderManagerChargePrdOrderOK) GetPayload() *models.NewbillingChargePrdOrderResponse {
	return o.Payload
}

func (o *SubscriptionPackageOrderManagerChargePrdOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingChargePrdOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionPackageOrderManagerChargePrdOrderDefault creates a SubscriptionPackageOrderManagerChargePrdOrderDefault with default headers values
func NewSubscriptionPackageOrderManagerChargePrdOrderDefault(code int) *SubscriptionPackageOrderManagerChargePrdOrderDefault {
	return &SubscriptionPackageOrderManagerChargePrdOrderDefault{
		_statusCode: code,
	}
}

/*
SubscriptionPackageOrderManagerChargePrdOrderDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type SubscriptionPackageOrderManagerChargePrdOrderDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the subscription package order manager charge prd order default response
func (o *SubscriptionPackageOrderManagerChargePrdOrderDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this subscription package order manager charge prd order default response has a 2xx status code
func (o *SubscriptionPackageOrderManagerChargePrdOrderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this subscription package order manager charge prd order default response has a 3xx status code
func (o *SubscriptionPackageOrderManagerChargePrdOrderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this subscription package order manager charge prd order default response has a 4xx status code
func (o *SubscriptionPackageOrderManagerChargePrdOrderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this subscription package order manager charge prd order default response has a 5xx status code
func (o *SubscriptionPackageOrderManagerChargePrdOrderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this subscription package order manager charge prd order default response a status code equal to that given
func (o *SubscriptionPackageOrderManagerChargePrdOrderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SubscriptionPackageOrderManagerChargePrdOrderDefault) Error() string {
	return fmt.Sprintf("[POST /v1/packageorders/{order_id}:pay][%d] SubscriptionPackageOrderManager_ChargePrdOrder default  %+v", o._statusCode, o.Payload)
}

func (o *SubscriptionPackageOrderManagerChargePrdOrderDefault) String() string {
	return fmt.Sprintf("[POST /v1/packageorders/{order_id}:pay][%d] SubscriptionPackageOrderManager_ChargePrdOrder default  %+v", o._statusCode, o.Payload)
}

func (o *SubscriptionPackageOrderManagerChargePrdOrderDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *SubscriptionPackageOrderManagerChargePrdOrderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}