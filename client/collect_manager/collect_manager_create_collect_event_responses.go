// Code generated by go-swagger; DO NOT EDIT.

package collect_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/newbilling-client-go/models"
)

// CollectManagerCreateCollectEventReader is a Reader for the CollectManagerCreateCollectEvent structure.
type CollectManagerCreateCollectEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CollectManagerCreateCollectEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCollectManagerCreateCollectEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCollectManagerCreateCollectEventDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCollectManagerCreateCollectEventOK creates a CollectManagerCreateCollectEventOK with default headers values
func NewCollectManagerCreateCollectEventOK() *CollectManagerCreateCollectEventOK {
	return &CollectManagerCreateCollectEventOK{}
}

/*
CollectManagerCreateCollectEventOK describes a response with status code 200, with default header values.

A successful response.
*/
type CollectManagerCreateCollectEventOK struct {
	Payload *models.NewbillingCreateCollectEventResponse
}

// IsSuccess returns true when this collect manager create collect event o k response has a 2xx status code
func (o *CollectManagerCreateCollectEventOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this collect manager create collect event o k response has a 3xx status code
func (o *CollectManagerCreateCollectEventOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this collect manager create collect event o k response has a 4xx status code
func (o *CollectManagerCreateCollectEventOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this collect manager create collect event o k response has a 5xx status code
func (o *CollectManagerCreateCollectEventOK) IsServerError() bool {
	return false
}

// IsCode returns true when this collect manager create collect event o k response a status code equal to that given
func (o *CollectManagerCreateCollectEventOK) IsCode(code int) bool {
	return code == 200
}

func (o *CollectManagerCreateCollectEventOK) Error() string {
	return fmt.Sprintf("[POST /v1/collect/events][%d] collectManagerCreateCollectEventOK  %+v", 200, o.Payload)
}

func (o *CollectManagerCreateCollectEventOK) String() string {
	return fmt.Sprintf("[POST /v1/collect/events][%d] collectManagerCreateCollectEventOK  %+v", 200, o.Payload)
}

func (o *CollectManagerCreateCollectEventOK) GetPayload() *models.NewbillingCreateCollectEventResponse {
	return o.Payload
}

func (o *CollectManagerCreateCollectEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingCreateCollectEventResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCollectManagerCreateCollectEventDefault creates a CollectManagerCreateCollectEventDefault with default headers values
func NewCollectManagerCreateCollectEventDefault(code int) *CollectManagerCreateCollectEventDefault {
	return &CollectManagerCreateCollectEventDefault{
		_statusCode: code,
	}
}

/*
CollectManagerCreateCollectEventDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type CollectManagerCreateCollectEventDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the collect manager create collect event default response
func (o *CollectManagerCreateCollectEventDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this collect manager create collect event default response has a 2xx status code
func (o *CollectManagerCreateCollectEventDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this collect manager create collect event default response has a 3xx status code
func (o *CollectManagerCreateCollectEventDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this collect manager create collect event default response has a 4xx status code
func (o *CollectManagerCreateCollectEventDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this collect manager create collect event default response has a 5xx status code
func (o *CollectManagerCreateCollectEventDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this collect manager create collect event default response a status code equal to that given
func (o *CollectManagerCreateCollectEventDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CollectManagerCreateCollectEventDefault) Error() string {
	return fmt.Sprintf("[POST /v1/collect/events][%d] CollectManager_CreateCollectEvent default  %+v", o._statusCode, o.Payload)
}

func (o *CollectManagerCreateCollectEventDefault) String() string {
	return fmt.Sprintf("[POST /v1/collect/events][%d] CollectManager_CreateCollectEvent default  %+v", o._statusCode, o.Payload)
}

func (o *CollectManagerCreateCollectEventDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *CollectManagerCreateCollectEventDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
