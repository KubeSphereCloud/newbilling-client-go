// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/newbilling-client-go/models"
)

// PaymentGetOrderReader is a Reader for the PaymentGetOrder structure.
type PaymentGetOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentGetOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentGetOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPaymentGetOrderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPaymentGetOrderOK creates a PaymentGetOrderOK with default headers values
func NewPaymentGetOrderOK() *PaymentGetOrderOK {
	return &PaymentGetOrderOK{}
}

/*
PaymentGetOrderOK describes a response with status code 200, with default header values.

A successful response.
*/
type PaymentGetOrderOK struct {
	Payload *models.NewbillingGetOrderResponse
}

// IsSuccess returns true when this payment get order o k response has a 2xx status code
func (o *PaymentGetOrderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment get order o k response has a 3xx status code
func (o *PaymentGetOrderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get order o k response has a 4xx status code
func (o *PaymentGetOrderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment get order o k response has a 5xx status code
func (o *PaymentGetOrderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get order o k response a status code equal to that given
func (o *PaymentGetOrderOK) IsCode(code int) bool {
	return code == 200
}

func (o *PaymentGetOrderOK) Error() string {
	return fmt.Sprintf("[GET /v1/payment/orders/{query_no}][%d] paymentGetOrderOK  %+v", 200, o.Payload)
}

func (o *PaymentGetOrderOK) String() string {
	return fmt.Sprintf("[GET /v1/payment/orders/{query_no}][%d] paymentGetOrderOK  %+v", 200, o.Payload)
}

func (o *PaymentGetOrderOK) GetPayload() *models.NewbillingGetOrderResponse {
	return o.Payload
}

func (o *PaymentGetOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingGetOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetOrderDefault creates a PaymentGetOrderDefault with default headers values
func NewPaymentGetOrderDefault(code int) *PaymentGetOrderDefault {
	return &PaymentGetOrderDefault{
		_statusCode: code,
	}
}

/*
PaymentGetOrderDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PaymentGetOrderDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the payment get order default response
func (o *PaymentGetOrderDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this payment get order default response has a 2xx status code
func (o *PaymentGetOrderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this payment get order default response has a 3xx status code
func (o *PaymentGetOrderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this payment get order default response has a 4xx status code
func (o *PaymentGetOrderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this payment get order default response has a 5xx status code
func (o *PaymentGetOrderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this payment get order default response a status code equal to that given
func (o *PaymentGetOrderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PaymentGetOrderDefault) Error() string {
	return fmt.Sprintf("[GET /v1/payment/orders/{query_no}][%d] Payment_GetOrder default  %+v", o._statusCode, o.Payload)
}

func (o *PaymentGetOrderDefault) String() string {
	return fmt.Sprintf("[GET /v1/payment/orders/{query_no}][%d] Payment_GetOrder default  %+v", o._statusCode, o.Payload)
}

func (o *PaymentGetOrderDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PaymentGetOrderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
