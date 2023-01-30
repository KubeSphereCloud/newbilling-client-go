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

// PaymentPrepayAlipayAppReader is a Reader for the PaymentPrepayAlipayApp structure.
type PaymentPrepayAlipayAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentPrepayAlipayAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentPrepayAlipayAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPaymentPrepayAlipayAppDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPaymentPrepayAlipayAppOK creates a PaymentPrepayAlipayAppOK with default headers values
func NewPaymentPrepayAlipayAppOK() *PaymentPrepayAlipayAppOK {
	return &PaymentPrepayAlipayAppOK{}
}

/*
PaymentPrepayAlipayAppOK describes a response with status code 200, with default header values.

A successful response.
*/
type PaymentPrepayAlipayAppOK struct {
	Payload *models.NewbillingPrepayAlipayAppResponse
}

// IsSuccess returns true when this payment prepay alipay app o k response has a 2xx status code
func (o *PaymentPrepayAlipayAppOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment prepay alipay app o k response has a 3xx status code
func (o *PaymentPrepayAlipayAppOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment prepay alipay app o k response has a 4xx status code
func (o *PaymentPrepayAlipayAppOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment prepay alipay app o k response has a 5xx status code
func (o *PaymentPrepayAlipayAppOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment prepay alipay app o k response a status code equal to that given
func (o *PaymentPrepayAlipayAppOK) IsCode(code int) bool {
	return code == 200
}

func (o *PaymentPrepayAlipayAppOK) Error() string {
	return fmt.Sprintf("[POST /v1/prepay/alipay/app][%d] paymentPrepayAlipayAppOK  %+v", 200, o.Payload)
}

func (o *PaymentPrepayAlipayAppOK) String() string {
	return fmt.Sprintf("[POST /v1/prepay/alipay/app][%d] paymentPrepayAlipayAppOK  %+v", 200, o.Payload)
}

func (o *PaymentPrepayAlipayAppOK) GetPayload() *models.NewbillingPrepayAlipayAppResponse {
	return o.Payload
}

func (o *PaymentPrepayAlipayAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingPrepayAlipayAppResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentPrepayAlipayAppDefault creates a PaymentPrepayAlipayAppDefault with default headers values
func NewPaymentPrepayAlipayAppDefault(code int) *PaymentPrepayAlipayAppDefault {
	return &PaymentPrepayAlipayAppDefault{
		_statusCode: code,
	}
}

/*
PaymentPrepayAlipayAppDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PaymentPrepayAlipayAppDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the payment prepay alipay app default response
func (o *PaymentPrepayAlipayAppDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this payment prepay alipay app default response has a 2xx status code
func (o *PaymentPrepayAlipayAppDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this payment prepay alipay app default response has a 3xx status code
func (o *PaymentPrepayAlipayAppDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this payment prepay alipay app default response has a 4xx status code
func (o *PaymentPrepayAlipayAppDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this payment prepay alipay app default response has a 5xx status code
func (o *PaymentPrepayAlipayAppDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this payment prepay alipay app default response a status code equal to that given
func (o *PaymentPrepayAlipayAppDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PaymentPrepayAlipayAppDefault) Error() string {
	return fmt.Sprintf("[POST /v1/prepay/alipay/app][%d] Payment_PrepayAlipayApp default  %+v", o._statusCode, o.Payload)
}

func (o *PaymentPrepayAlipayAppDefault) String() string {
	return fmt.Sprintf("[POST /v1/prepay/alipay/app][%d] Payment_PrepayAlipayApp default  %+v", o._statusCode, o.Payload)
}

func (o *PaymentPrepayAlipayAppDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PaymentPrepayAlipayAppDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
