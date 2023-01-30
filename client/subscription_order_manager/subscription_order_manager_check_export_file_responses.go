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

// SubscriptionOrderManagerCheckExportFileReader is a Reader for the SubscriptionOrderManagerCheckExportFile structure.
type SubscriptionOrderManagerCheckExportFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriptionOrderManagerCheckExportFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscriptionOrderManagerCheckExportFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSubscriptionOrderManagerCheckExportFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSubscriptionOrderManagerCheckExportFileOK creates a SubscriptionOrderManagerCheckExportFileOK with default headers values
func NewSubscriptionOrderManagerCheckExportFileOK() *SubscriptionOrderManagerCheckExportFileOK {
	return &SubscriptionOrderManagerCheckExportFileOK{}
}

/*
SubscriptionOrderManagerCheckExportFileOK describes a response with status code 200, with default header values.

A successful response.
*/
type SubscriptionOrderManagerCheckExportFileOK struct {
	Payload *models.NewbillingCheckExportFileResponse
}

// IsSuccess returns true when this subscription order manager check export file o k response has a 2xx status code
func (o *SubscriptionOrderManagerCheckExportFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this subscription order manager check export file o k response has a 3xx status code
func (o *SubscriptionOrderManagerCheckExportFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription order manager check export file o k response has a 4xx status code
func (o *SubscriptionOrderManagerCheckExportFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription order manager check export file o k response has a 5xx status code
func (o *SubscriptionOrderManagerCheckExportFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription order manager check export file o k response a status code equal to that given
func (o *SubscriptionOrderManagerCheckExportFileOK) IsCode(code int) bool {
	return code == 200
}

func (o *SubscriptionOrderManagerCheckExportFileOK) Error() string {
	return fmt.Sprintf("[GET /v1/export/check][%d] subscriptionOrderManagerCheckExportFileOK  %+v", 200, o.Payload)
}

func (o *SubscriptionOrderManagerCheckExportFileOK) String() string {
	return fmt.Sprintf("[GET /v1/export/check][%d] subscriptionOrderManagerCheckExportFileOK  %+v", 200, o.Payload)
}

func (o *SubscriptionOrderManagerCheckExportFileOK) GetPayload() *models.NewbillingCheckExportFileResponse {
	return o.Payload
}

func (o *SubscriptionOrderManagerCheckExportFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingCheckExportFileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionOrderManagerCheckExportFileDefault creates a SubscriptionOrderManagerCheckExportFileDefault with default headers values
func NewSubscriptionOrderManagerCheckExportFileDefault(code int) *SubscriptionOrderManagerCheckExportFileDefault {
	return &SubscriptionOrderManagerCheckExportFileDefault{
		_statusCode: code,
	}
}

/*
SubscriptionOrderManagerCheckExportFileDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type SubscriptionOrderManagerCheckExportFileDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the subscription order manager check export file default response
func (o *SubscriptionOrderManagerCheckExportFileDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this subscription order manager check export file default response has a 2xx status code
func (o *SubscriptionOrderManagerCheckExportFileDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this subscription order manager check export file default response has a 3xx status code
func (o *SubscriptionOrderManagerCheckExportFileDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this subscription order manager check export file default response has a 4xx status code
func (o *SubscriptionOrderManagerCheckExportFileDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this subscription order manager check export file default response has a 5xx status code
func (o *SubscriptionOrderManagerCheckExportFileDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this subscription order manager check export file default response a status code equal to that given
func (o *SubscriptionOrderManagerCheckExportFileDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SubscriptionOrderManagerCheckExportFileDefault) Error() string {
	return fmt.Sprintf("[GET /v1/export/check][%d] SubscriptionOrderManager_CheckExportFile default  %+v", o._statusCode, o.Payload)
}

func (o *SubscriptionOrderManagerCheckExportFileDefault) String() string {
	return fmt.Sprintf("[GET /v1/export/check][%d] SubscriptionOrderManager_CheckExportFile default  %+v", o._statusCode, o.Payload)
}

func (o *SubscriptionOrderManagerCheckExportFileDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *SubscriptionOrderManagerCheckExportFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
