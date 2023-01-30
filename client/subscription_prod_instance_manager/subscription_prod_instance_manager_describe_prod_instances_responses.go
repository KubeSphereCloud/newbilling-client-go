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

// SubscriptionProdInstanceManagerDescribeProdInstancesReader is a Reader for the SubscriptionProdInstanceManagerDescribeProdInstances structure.
type SubscriptionProdInstanceManagerDescribeProdInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscriptionProdInstanceManagerDescribeProdInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSubscriptionProdInstanceManagerDescribeProdInstancesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSubscriptionProdInstanceManagerDescribeProdInstancesOK creates a SubscriptionProdInstanceManagerDescribeProdInstancesOK with default headers values
func NewSubscriptionProdInstanceManagerDescribeProdInstancesOK() *SubscriptionProdInstanceManagerDescribeProdInstancesOK {
	return &SubscriptionProdInstanceManagerDescribeProdInstancesOK{}
}

/*
SubscriptionProdInstanceManagerDescribeProdInstancesOK describes a response with status code 200, with default header values.

A successful response.
*/
type SubscriptionProdInstanceManagerDescribeProdInstancesOK struct {
	Payload *models.NewbillingDescribeProdInstancesResponse
}

// IsSuccess returns true when this subscription prod instance manager describe prod instances o k response has a 2xx status code
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this subscription prod instance manager describe prod instances o k response has a 3xx status code
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription prod instance manager describe prod instances o k response has a 4xx status code
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription prod instance manager describe prod instances o k response has a 5xx status code
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription prod instance manager describe prod instances o k response a status code equal to that given
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesOK) IsCode(code int) bool {
	return code == 200
}

func (o *SubscriptionProdInstanceManagerDescribeProdInstancesOK) Error() string {
	return fmt.Sprintf("[GET /v1/prodinstances][%d] subscriptionProdInstanceManagerDescribeProdInstancesOK  %+v", 200, o.Payload)
}

func (o *SubscriptionProdInstanceManagerDescribeProdInstancesOK) String() string {
	return fmt.Sprintf("[GET /v1/prodinstances][%d] subscriptionProdInstanceManagerDescribeProdInstancesOK  %+v", 200, o.Payload)
}

func (o *SubscriptionProdInstanceManagerDescribeProdInstancesOK) GetPayload() *models.NewbillingDescribeProdInstancesResponse {
	return o.Payload
}

func (o *SubscriptionProdInstanceManagerDescribeProdInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingDescribeProdInstancesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionProdInstanceManagerDescribeProdInstancesDefault creates a SubscriptionProdInstanceManagerDescribeProdInstancesDefault with default headers values
func NewSubscriptionProdInstanceManagerDescribeProdInstancesDefault(code int) *SubscriptionProdInstanceManagerDescribeProdInstancesDefault {
	return &SubscriptionProdInstanceManagerDescribeProdInstancesDefault{
		_statusCode: code,
	}
}

/*
SubscriptionProdInstanceManagerDescribeProdInstancesDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type SubscriptionProdInstanceManagerDescribeProdInstancesDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the subscription prod instance manager describe prod instances default response
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this subscription prod instance manager describe prod instances default response has a 2xx status code
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this subscription prod instance manager describe prod instances default response has a 3xx status code
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this subscription prod instance manager describe prod instances default response has a 4xx status code
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this subscription prod instance manager describe prod instances default response has a 5xx status code
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this subscription prod instance manager describe prod instances default response a status code equal to that given
func (o *SubscriptionProdInstanceManagerDescribeProdInstancesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SubscriptionProdInstanceManagerDescribeProdInstancesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/prodinstances][%d] SubscriptionProdInstanceManager_DescribeProdInstances default  %+v", o._statusCode, o.Payload)
}

func (o *SubscriptionProdInstanceManagerDescribeProdInstancesDefault) String() string {
	return fmt.Sprintf("[GET /v1/prodinstances][%d] SubscriptionProdInstanceManager_DescribeProdInstances default  %+v", o._statusCode, o.Payload)
}

func (o *SubscriptionProdInstanceManagerDescribeProdInstancesDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *SubscriptionProdInstanceManagerDescribeProdInstancesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}