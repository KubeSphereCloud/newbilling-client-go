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

// SubscriptionProdInstanceManagerTransferCompToTimePackageReader is a Reader for the SubscriptionProdInstanceManagerTransferCompToTimePackage structure.
type SubscriptionProdInstanceManagerTransferCompToTimePackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscriptionProdInstanceManagerTransferCompToTimePackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSubscriptionProdInstanceManagerTransferCompToTimePackageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSubscriptionProdInstanceManagerTransferCompToTimePackageOK creates a SubscriptionProdInstanceManagerTransferCompToTimePackageOK with default headers values
func NewSubscriptionProdInstanceManagerTransferCompToTimePackageOK() *SubscriptionProdInstanceManagerTransferCompToTimePackageOK {
	return &SubscriptionProdInstanceManagerTransferCompToTimePackageOK{}
}

/*
SubscriptionProdInstanceManagerTransferCompToTimePackageOK describes a response with status code 200, with default header values.

A successful response.
*/
type SubscriptionProdInstanceManagerTransferCompToTimePackageOK struct {
	Payload *models.NewbillingTransferCompToTimePackageResponse
}

// IsSuccess returns true when this subscription prod instance manager transfer comp to time package o k response has a 2xx status code
func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this subscription prod instance manager transfer comp to time package o k response has a 3xx status code
func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription prod instance manager transfer comp to time package o k response has a 4xx status code
func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription prod instance manager transfer comp to time package o k response has a 5xx status code
func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription prod instance manager transfer comp to time package o k response a status code equal to that given
func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageOK) IsCode(code int) bool {
	return code == 200
}

func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageOK) Error() string {
	return fmt.Sprintf("[POST /v1/prodinstances/{prod_inst_id_ext}:transfer2timepackage][%d] subscriptionProdInstanceManagerTransferCompToTimePackageOK  %+v", 200, o.Payload)
}

func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageOK) String() string {
	return fmt.Sprintf("[POST /v1/prodinstances/{prod_inst_id_ext}:transfer2timepackage][%d] subscriptionProdInstanceManagerTransferCompToTimePackageOK  %+v", 200, o.Payload)
}

func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageOK) GetPayload() *models.NewbillingTransferCompToTimePackageResponse {
	return o.Payload
}

func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingTransferCompToTimePackageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionProdInstanceManagerTransferCompToTimePackageDefault creates a SubscriptionProdInstanceManagerTransferCompToTimePackageDefault with default headers values
func NewSubscriptionProdInstanceManagerTransferCompToTimePackageDefault(code int) *SubscriptionProdInstanceManagerTransferCompToTimePackageDefault {
	return &SubscriptionProdInstanceManagerTransferCompToTimePackageDefault{
		_statusCode: code,
	}
}

/*
SubscriptionProdInstanceManagerTransferCompToTimePackageDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type SubscriptionProdInstanceManagerTransferCompToTimePackageDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the subscription prod instance manager transfer comp to time package default response
func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this subscription prod instance manager transfer comp to time package default response has a 2xx status code
func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this subscription prod instance manager transfer comp to time package default response has a 3xx status code
func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this subscription prod instance manager transfer comp to time package default response has a 4xx status code
func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this subscription prod instance manager transfer comp to time package default response has a 5xx status code
func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this subscription prod instance manager transfer comp to time package default response a status code equal to that given
func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageDefault) Error() string {
	return fmt.Sprintf("[POST /v1/prodinstances/{prod_inst_id_ext}:transfer2timepackage][%d] SubscriptionProdInstanceManager_TransferCompToTimePackage default  %+v", o._statusCode, o.Payload)
}

func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageDefault) String() string {
	return fmt.Sprintf("[POST /v1/prodinstances/{prod_inst_id_ext}:transfer2timepackage][%d] SubscriptionProdInstanceManager_TransferCompToTimePackage default  %+v", o._statusCode, o.Payload)
}

func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *SubscriptionProdInstanceManagerTransferCompToTimePackageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
