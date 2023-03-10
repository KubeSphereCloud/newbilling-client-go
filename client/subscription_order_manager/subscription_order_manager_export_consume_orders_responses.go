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

// SubscriptionOrderManagerExportConsumeOrdersReader is a Reader for the SubscriptionOrderManagerExportConsumeOrders structure.
type SubscriptionOrderManagerExportConsumeOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriptionOrderManagerExportConsumeOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscriptionOrderManagerExportConsumeOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSubscriptionOrderManagerExportConsumeOrdersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSubscriptionOrderManagerExportConsumeOrdersOK creates a SubscriptionOrderManagerExportConsumeOrdersOK with default headers values
func NewSubscriptionOrderManagerExportConsumeOrdersOK() *SubscriptionOrderManagerExportConsumeOrdersOK {
	return &SubscriptionOrderManagerExportConsumeOrdersOK{}
}

/*
SubscriptionOrderManagerExportConsumeOrdersOK describes a response with status code 200, with default header values.

A successful response.
*/
type SubscriptionOrderManagerExportConsumeOrdersOK struct {
	Payload *models.NewbillingExportConsumeOrdersResponse
}

// IsSuccess returns true when this subscription order manager export consume orders o k response has a 2xx status code
func (o *SubscriptionOrderManagerExportConsumeOrdersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this subscription order manager export consume orders o k response has a 3xx status code
func (o *SubscriptionOrderManagerExportConsumeOrdersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription order manager export consume orders o k response has a 4xx status code
func (o *SubscriptionOrderManagerExportConsumeOrdersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription order manager export consume orders o k response has a 5xx status code
func (o *SubscriptionOrderManagerExportConsumeOrdersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription order manager export consume orders o k response a status code equal to that given
func (o *SubscriptionOrderManagerExportConsumeOrdersOK) IsCode(code int) bool {
	return code == 200
}

func (o *SubscriptionOrderManagerExportConsumeOrdersOK) Error() string {
	return fmt.Sprintf("[GET /v1/orders/export][%d] subscriptionOrderManagerExportConsumeOrdersOK  %+v", 200, o.Payload)
}

func (o *SubscriptionOrderManagerExportConsumeOrdersOK) String() string {
	return fmt.Sprintf("[GET /v1/orders/export][%d] subscriptionOrderManagerExportConsumeOrdersOK  %+v", 200, o.Payload)
}

func (o *SubscriptionOrderManagerExportConsumeOrdersOK) GetPayload() *models.NewbillingExportConsumeOrdersResponse {
	return o.Payload
}

func (o *SubscriptionOrderManagerExportConsumeOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingExportConsumeOrdersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionOrderManagerExportConsumeOrdersDefault creates a SubscriptionOrderManagerExportConsumeOrdersDefault with default headers values
func NewSubscriptionOrderManagerExportConsumeOrdersDefault(code int) *SubscriptionOrderManagerExportConsumeOrdersDefault {
	return &SubscriptionOrderManagerExportConsumeOrdersDefault{
		_statusCode: code,
	}
}

/*
SubscriptionOrderManagerExportConsumeOrdersDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type SubscriptionOrderManagerExportConsumeOrdersDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the subscription order manager export consume orders default response
func (o *SubscriptionOrderManagerExportConsumeOrdersDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this subscription order manager export consume orders default response has a 2xx status code
func (o *SubscriptionOrderManagerExportConsumeOrdersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this subscription order manager export consume orders default response has a 3xx status code
func (o *SubscriptionOrderManagerExportConsumeOrdersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this subscription order manager export consume orders default response has a 4xx status code
func (o *SubscriptionOrderManagerExportConsumeOrdersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this subscription order manager export consume orders default response has a 5xx status code
func (o *SubscriptionOrderManagerExportConsumeOrdersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this subscription order manager export consume orders default response a status code equal to that given
func (o *SubscriptionOrderManagerExportConsumeOrdersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SubscriptionOrderManagerExportConsumeOrdersDefault) Error() string {
	return fmt.Sprintf("[GET /v1/orders/export][%d] SubscriptionOrderManager_ExportConsumeOrders default  %+v", o._statusCode, o.Payload)
}

func (o *SubscriptionOrderManagerExportConsumeOrdersDefault) String() string {
	return fmt.Sprintf("[GET /v1/orders/export][%d] SubscriptionOrderManager_ExportConsumeOrders default  %+v", o._statusCode, o.Payload)
}

func (o *SubscriptionOrderManagerExportConsumeOrdersDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *SubscriptionOrderManagerExportConsumeOrdersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
