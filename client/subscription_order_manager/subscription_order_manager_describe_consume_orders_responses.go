// Code generated by go-swagger; DO NOT EDIT.

package subscription_order_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/newbilling-client-go/models"
)

// SubscriptionOrderManagerDescribeConsumeOrdersReader is a Reader for the SubscriptionOrderManagerDescribeConsumeOrders structure.
type SubscriptionOrderManagerDescribeConsumeOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriptionOrderManagerDescribeConsumeOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscriptionOrderManagerDescribeConsumeOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSubscriptionOrderManagerDescribeConsumeOrdersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSubscriptionOrderManagerDescribeConsumeOrdersOK creates a SubscriptionOrderManagerDescribeConsumeOrdersOK with default headers values
func NewSubscriptionOrderManagerDescribeConsumeOrdersOK() *SubscriptionOrderManagerDescribeConsumeOrdersOK {
	return &SubscriptionOrderManagerDescribeConsumeOrdersOK{}
}

/*
SubscriptionOrderManagerDescribeConsumeOrdersOK describes a response with status code 200, with default header values.

A successful response.
*/
type SubscriptionOrderManagerDescribeConsumeOrdersOK struct {
	Payload *models.NewbillingDescribeConsumeOrdersResponse
}

// IsSuccess returns true when this subscription order manager describe consume orders o k response has a 2xx status code
func (o *SubscriptionOrderManagerDescribeConsumeOrdersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this subscription order manager describe consume orders o k response has a 3xx status code
func (o *SubscriptionOrderManagerDescribeConsumeOrdersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription order manager describe consume orders o k response has a 4xx status code
func (o *SubscriptionOrderManagerDescribeConsumeOrdersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription order manager describe consume orders o k response has a 5xx status code
func (o *SubscriptionOrderManagerDescribeConsumeOrdersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription order manager describe consume orders o k response a status code equal to that given
func (o *SubscriptionOrderManagerDescribeConsumeOrdersOK) IsCode(code int) bool {
	return code == 200
}

func (o *SubscriptionOrderManagerDescribeConsumeOrdersOK) Error() string {
	return fmt.Sprintf("[GET /v1/orders][%d] subscriptionOrderManagerDescribeConsumeOrdersOK  %+v", 200, o.Payload)
}

func (o *SubscriptionOrderManagerDescribeConsumeOrdersOK) String() string {
	return fmt.Sprintf("[GET /v1/orders][%d] subscriptionOrderManagerDescribeConsumeOrdersOK  %+v", 200, o.Payload)
}

func (o *SubscriptionOrderManagerDescribeConsumeOrdersOK) GetPayload() *models.NewbillingDescribeConsumeOrdersResponse {
	return o.Payload
}

func (o *SubscriptionOrderManagerDescribeConsumeOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingDescribeConsumeOrdersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionOrderManagerDescribeConsumeOrdersDefault creates a SubscriptionOrderManagerDescribeConsumeOrdersDefault with default headers values
func NewSubscriptionOrderManagerDescribeConsumeOrdersDefault(code int) *SubscriptionOrderManagerDescribeConsumeOrdersDefault {
	return &SubscriptionOrderManagerDescribeConsumeOrdersDefault{
		_statusCode: code,
	}
}

/*
SubscriptionOrderManagerDescribeConsumeOrdersDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type SubscriptionOrderManagerDescribeConsumeOrdersDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the subscription order manager describe consume orders default response
func (o *SubscriptionOrderManagerDescribeConsumeOrdersDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this subscription order manager describe consume orders default response has a 2xx status code
func (o *SubscriptionOrderManagerDescribeConsumeOrdersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this subscription order manager describe consume orders default response has a 3xx status code
func (o *SubscriptionOrderManagerDescribeConsumeOrdersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this subscription order manager describe consume orders default response has a 4xx status code
func (o *SubscriptionOrderManagerDescribeConsumeOrdersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this subscription order manager describe consume orders default response has a 5xx status code
func (o *SubscriptionOrderManagerDescribeConsumeOrdersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this subscription order manager describe consume orders default response a status code equal to that given
func (o *SubscriptionOrderManagerDescribeConsumeOrdersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SubscriptionOrderManagerDescribeConsumeOrdersDefault) Error() string {
	return fmt.Sprintf("[GET /v1/orders][%d] SubscriptionOrderManager_DescribeConsumeOrders default  %+v", o._statusCode, o.Payload)
}

func (o *SubscriptionOrderManagerDescribeConsumeOrdersDefault) String() string {
	return fmt.Sprintf("[GET /v1/orders][%d] SubscriptionOrderManager_DescribeConsumeOrders default  %+v", o._statusCode, o.Payload)
}

func (o *SubscriptionOrderManagerDescribeConsumeOrdersDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *SubscriptionOrderManagerDescribeConsumeOrdersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}