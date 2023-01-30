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

// SubscriptionPackageOrderManagerCreatePrdOrderV2Reader is a Reader for the SubscriptionPackageOrderManagerCreatePrdOrderV2 structure.
type SubscriptionPackageOrderManagerCreatePrdOrderV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscriptionPackageOrderManagerCreatePrdOrderV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSubscriptionPackageOrderManagerCreatePrdOrderV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSubscriptionPackageOrderManagerCreatePrdOrderV2OK creates a SubscriptionPackageOrderManagerCreatePrdOrderV2OK with default headers values
func NewSubscriptionPackageOrderManagerCreatePrdOrderV2OK() *SubscriptionPackageOrderManagerCreatePrdOrderV2OK {
	return &SubscriptionPackageOrderManagerCreatePrdOrderV2OK{}
}

/*
SubscriptionPackageOrderManagerCreatePrdOrderV2OK describes a response with status code 200, with default header values.

A successful response.
*/
type SubscriptionPackageOrderManagerCreatePrdOrderV2OK struct {
	Payload *models.NewbillingCreatePrdOrderResponse
}

// IsSuccess returns true when this subscription package order manager create prd order v2 o k response has a 2xx status code
func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this subscription package order manager create prd order v2 o k response has a 3xx status code
func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription package order manager create prd order v2 o k response has a 4xx status code
func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription package order manager create prd order v2 o k response has a 5xx status code
func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription package order manager create prd order v2 o k response a status code equal to that given
func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2OK) IsCode(code int) bool {
	return code == 200
}

func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2OK) Error() string {
	return fmt.Sprintf("[POST /v2/packageorders][%d] subscriptionPackageOrderManagerCreatePrdOrderV2OK  %+v", 200, o.Payload)
}

func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2OK) String() string {
	return fmt.Sprintf("[POST /v2/packageorders][%d] subscriptionPackageOrderManagerCreatePrdOrderV2OK  %+v", 200, o.Payload)
}

func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2OK) GetPayload() *models.NewbillingCreatePrdOrderResponse {
	return o.Payload
}

func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingCreatePrdOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionPackageOrderManagerCreatePrdOrderV2Default creates a SubscriptionPackageOrderManagerCreatePrdOrderV2Default with default headers values
func NewSubscriptionPackageOrderManagerCreatePrdOrderV2Default(code int) *SubscriptionPackageOrderManagerCreatePrdOrderV2Default {
	return &SubscriptionPackageOrderManagerCreatePrdOrderV2Default{
		_statusCode: code,
	}
}

/*
SubscriptionPackageOrderManagerCreatePrdOrderV2Default describes a response with status code -1, with default header values.

An unexpected error response
*/
type SubscriptionPackageOrderManagerCreatePrdOrderV2Default struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the subscription package order manager create prd order v2 default response
func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this subscription package order manager create prd order v2 default response has a 2xx status code
func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this subscription package order manager create prd order v2 default response has a 3xx status code
func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this subscription package order manager create prd order v2 default response has a 4xx status code
func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this subscription package order manager create prd order v2 default response has a 5xx status code
func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this subscription package order manager create prd order v2 default response a status code equal to that given
func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2Default) Error() string {
	return fmt.Sprintf("[POST /v2/packageorders][%d] SubscriptionPackageOrderManager_CreatePrdOrderV2 default  %+v", o._statusCode, o.Payload)
}

func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2Default) String() string {
	return fmt.Sprintf("[POST /v2/packageorders][%d] SubscriptionPackageOrderManager_CreatePrdOrderV2 default  %+v", o._statusCode, o.Payload)
}

func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2Default) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *SubscriptionPackageOrderManagerCreatePrdOrderV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}