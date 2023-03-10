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

// CollectManagerCreateCollectDataReader is a Reader for the CollectManagerCreateCollectData structure.
type CollectManagerCreateCollectDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CollectManagerCreateCollectDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCollectManagerCreateCollectDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCollectManagerCreateCollectDataDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCollectManagerCreateCollectDataOK creates a CollectManagerCreateCollectDataOK with default headers values
func NewCollectManagerCreateCollectDataOK() *CollectManagerCreateCollectDataOK {
	return &CollectManagerCreateCollectDataOK{}
}

/*
CollectManagerCreateCollectDataOK describes a response with status code 200, with default header values.

A successful response.
*/
type CollectManagerCreateCollectDataOK struct {
	Payload models.NewbillingCreateCollectDataResponse
}

// IsSuccess returns true when this collect manager create collect data o k response has a 2xx status code
func (o *CollectManagerCreateCollectDataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this collect manager create collect data o k response has a 3xx status code
func (o *CollectManagerCreateCollectDataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this collect manager create collect data o k response has a 4xx status code
func (o *CollectManagerCreateCollectDataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this collect manager create collect data o k response has a 5xx status code
func (o *CollectManagerCreateCollectDataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this collect manager create collect data o k response a status code equal to that given
func (o *CollectManagerCreateCollectDataOK) IsCode(code int) bool {
	return code == 200
}

func (o *CollectManagerCreateCollectDataOK) Error() string {
	return fmt.Sprintf("[POST /v1/collect/data][%d] collectManagerCreateCollectDataOK  %+v", 200, o.Payload)
}

func (o *CollectManagerCreateCollectDataOK) String() string {
	return fmt.Sprintf("[POST /v1/collect/data][%d] collectManagerCreateCollectDataOK  %+v", 200, o.Payload)
}

func (o *CollectManagerCreateCollectDataOK) GetPayload() models.NewbillingCreateCollectDataResponse {
	return o.Payload
}

func (o *CollectManagerCreateCollectDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCollectManagerCreateCollectDataDefault creates a CollectManagerCreateCollectDataDefault with default headers values
func NewCollectManagerCreateCollectDataDefault(code int) *CollectManagerCreateCollectDataDefault {
	return &CollectManagerCreateCollectDataDefault{
		_statusCode: code,
	}
}

/*
CollectManagerCreateCollectDataDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type CollectManagerCreateCollectDataDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the collect manager create collect data default response
func (o *CollectManagerCreateCollectDataDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this collect manager create collect data default response has a 2xx status code
func (o *CollectManagerCreateCollectDataDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this collect manager create collect data default response has a 3xx status code
func (o *CollectManagerCreateCollectDataDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this collect manager create collect data default response has a 4xx status code
func (o *CollectManagerCreateCollectDataDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this collect manager create collect data default response has a 5xx status code
func (o *CollectManagerCreateCollectDataDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this collect manager create collect data default response a status code equal to that given
func (o *CollectManagerCreateCollectDataDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CollectManagerCreateCollectDataDefault) Error() string {
	return fmt.Sprintf("[POST /v1/collect/data][%d] CollectManager_CreateCollectData default  %+v", o._statusCode, o.Payload)
}

func (o *CollectManagerCreateCollectDataDefault) String() string {
	return fmt.Sprintf("[POST /v1/collect/data][%d] CollectManager_CreateCollectData default  %+v", o._statusCode, o.Payload)
}

func (o *CollectManagerCreateCollectDataDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *CollectManagerCreateCollectDataDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
