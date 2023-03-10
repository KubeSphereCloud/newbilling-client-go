// Code generated by go-swagger; DO NOT EDIT.

package notifier

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/newbilling-client-go/models"
)

// NotifierCreateEventReader is a Reader for the NotifierCreateEvent structure.
type NotifierCreateEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotifierCreateEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotifierCreateEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNotifierCreateEventDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNotifierCreateEventOK creates a NotifierCreateEventOK with default headers values
func NewNotifierCreateEventOK() *NotifierCreateEventOK {
	return &NotifierCreateEventOK{}
}

/*
NotifierCreateEventOK describes a response with status code 200, with default header values.

A successful response.
*/
type NotifierCreateEventOK struct {
	Payload *models.NewbillingCreateEventResponse
}

// IsSuccess returns true when this notifier create event o k response has a 2xx status code
func (o *NotifierCreateEventOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notifier create event o k response has a 3xx status code
func (o *NotifierCreateEventOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifier create event o k response has a 4xx status code
func (o *NotifierCreateEventOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this notifier create event o k response has a 5xx status code
func (o *NotifierCreateEventOK) IsServerError() bool {
	return false
}

// IsCode returns true when this notifier create event o k response a status code equal to that given
func (o *NotifierCreateEventOK) IsCode(code int) bool {
	return code == 200
}

func (o *NotifierCreateEventOK) Error() string {
	return fmt.Sprintf("[POST /v1/events][%d] notifierCreateEventOK  %+v", 200, o.Payload)
}

func (o *NotifierCreateEventOK) String() string {
	return fmt.Sprintf("[POST /v1/events][%d] notifierCreateEventOK  %+v", 200, o.Payload)
}

func (o *NotifierCreateEventOK) GetPayload() *models.NewbillingCreateEventResponse {
	return o.Payload
}

func (o *NotifierCreateEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingCreateEventResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotifierCreateEventDefault creates a NotifierCreateEventDefault with default headers values
func NewNotifierCreateEventDefault(code int) *NotifierCreateEventDefault {
	return &NotifierCreateEventDefault{
		_statusCode: code,
	}
}

/*
NotifierCreateEventDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type NotifierCreateEventDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the notifier create event default response
func (o *NotifierCreateEventDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this notifier create event default response has a 2xx status code
func (o *NotifierCreateEventDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this notifier create event default response has a 3xx status code
func (o *NotifierCreateEventDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this notifier create event default response has a 4xx status code
func (o *NotifierCreateEventDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this notifier create event default response has a 5xx status code
func (o *NotifierCreateEventDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this notifier create event default response a status code equal to that given
func (o *NotifierCreateEventDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *NotifierCreateEventDefault) Error() string {
	return fmt.Sprintf("[POST /v1/events][%d] Notifier_CreateEvent default  %+v", o._statusCode, o.Payload)
}

func (o *NotifierCreateEventDefault) String() string {
	return fmt.Sprintf("[POST /v1/events][%d] Notifier_CreateEvent default  %+v", o._statusCode, o.Payload)
}

func (o *NotifierCreateEventDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *NotifierCreateEventDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
