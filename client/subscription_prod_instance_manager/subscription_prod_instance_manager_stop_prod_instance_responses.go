// Code generated by go-swagger; DO NOT EDIT.

package subscription_prod_instance_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/newbilling-client-go/models"
)

// SubscriptionProdInstanceManagerStopProdInstanceReader is a Reader for the SubscriptionProdInstanceManagerStopProdInstance structure.
type SubscriptionProdInstanceManagerStopProdInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriptionProdInstanceManagerStopProdInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscriptionProdInstanceManagerStopProdInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSubscriptionProdInstanceManagerStopProdInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSubscriptionProdInstanceManagerStopProdInstanceOK creates a SubscriptionProdInstanceManagerStopProdInstanceOK with default headers values
func NewSubscriptionProdInstanceManagerStopProdInstanceOK() *SubscriptionProdInstanceManagerStopProdInstanceOK {
	return &SubscriptionProdInstanceManagerStopProdInstanceOK{}
}

/*
SubscriptionProdInstanceManagerStopProdInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type SubscriptionProdInstanceManagerStopProdInstanceOK struct {
	Payload *models.NewbillingStopProdInstanceResponse
}

// IsSuccess returns true when this subscription prod instance manager stop prod instance o k response has a 2xx status code
func (o *SubscriptionProdInstanceManagerStopProdInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this subscription prod instance manager stop prod instance o k response has a 3xx status code
func (o *SubscriptionProdInstanceManagerStopProdInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription prod instance manager stop prod instance o k response has a 4xx status code
func (o *SubscriptionProdInstanceManagerStopProdInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription prod instance manager stop prod instance o k response has a 5xx status code
func (o *SubscriptionProdInstanceManagerStopProdInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription prod instance manager stop prod instance o k response a status code equal to that given
func (o *SubscriptionProdInstanceManagerStopProdInstanceOK) IsCode(code int) bool {
	return code == 200
}

func (o *SubscriptionProdInstanceManagerStopProdInstanceOK) Error() string {
	return fmt.Sprintf("[POST /v1/prodinstances/{prod_inst_id_ext}:stop][%d] subscriptionProdInstanceManagerStopProdInstanceOK  %+v", 200, o.Payload)
}

func (o *SubscriptionProdInstanceManagerStopProdInstanceOK) String() string {
	return fmt.Sprintf("[POST /v1/prodinstances/{prod_inst_id_ext}:stop][%d] subscriptionProdInstanceManagerStopProdInstanceOK  %+v", 200, o.Payload)
}

func (o *SubscriptionProdInstanceManagerStopProdInstanceOK) GetPayload() *models.NewbillingStopProdInstanceResponse {
	return o.Payload
}

func (o *SubscriptionProdInstanceManagerStopProdInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingStopProdInstanceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionProdInstanceManagerStopProdInstanceDefault creates a SubscriptionProdInstanceManagerStopProdInstanceDefault with default headers values
func NewSubscriptionProdInstanceManagerStopProdInstanceDefault(code int) *SubscriptionProdInstanceManagerStopProdInstanceDefault {
	return &SubscriptionProdInstanceManagerStopProdInstanceDefault{
		_statusCode: code,
	}
}

/*
SubscriptionProdInstanceManagerStopProdInstanceDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type SubscriptionProdInstanceManagerStopProdInstanceDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the subscription prod instance manager stop prod instance default response
func (o *SubscriptionProdInstanceManagerStopProdInstanceDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this subscription prod instance manager stop prod instance default response has a 2xx status code
func (o *SubscriptionProdInstanceManagerStopProdInstanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this subscription prod instance manager stop prod instance default response has a 3xx status code
func (o *SubscriptionProdInstanceManagerStopProdInstanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this subscription prod instance manager stop prod instance default response has a 4xx status code
func (o *SubscriptionProdInstanceManagerStopProdInstanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this subscription prod instance manager stop prod instance default response has a 5xx status code
func (o *SubscriptionProdInstanceManagerStopProdInstanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this subscription prod instance manager stop prod instance default response a status code equal to that given
func (o *SubscriptionProdInstanceManagerStopProdInstanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SubscriptionProdInstanceManagerStopProdInstanceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/prodinstances/{prod_inst_id_ext}:stop][%d] SubscriptionProdInstanceManager_StopProdInstance default  %+v", o._statusCode, o.Payload)
}

func (o *SubscriptionProdInstanceManagerStopProdInstanceDefault) String() string {
	return fmt.Sprintf("[POST /v1/prodinstances/{prod_inst_id_ext}:stop][%d] SubscriptionProdInstanceManager_StopProdInstance default  %+v", o._statusCode, o.Payload)
}

func (o *SubscriptionProdInstanceManagerStopProdInstanceDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *SubscriptionProdInstanceManagerStopProdInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}