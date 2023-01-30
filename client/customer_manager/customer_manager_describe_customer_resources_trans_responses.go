// Code generated by go-swagger; DO NOT EDIT.

package customer_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/newbilling-client-go/models"
)

// CustomerManagerDescribeCustomerResourcesTransReader is a Reader for the CustomerManagerDescribeCustomerResourcesTrans structure.
type CustomerManagerDescribeCustomerResourcesTransReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerManagerDescribeCustomerResourcesTransReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerManagerDescribeCustomerResourcesTransOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerManagerDescribeCustomerResourcesTransDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerManagerDescribeCustomerResourcesTransOK creates a CustomerManagerDescribeCustomerResourcesTransOK with default headers values
func NewCustomerManagerDescribeCustomerResourcesTransOK() *CustomerManagerDescribeCustomerResourcesTransOK {
	return &CustomerManagerDescribeCustomerResourcesTransOK{}
}

/*
CustomerManagerDescribeCustomerResourcesTransOK describes a response with status code 200, with default header values.

A successful response.
*/
type CustomerManagerDescribeCustomerResourcesTransOK struct {
	Payload *models.NewbillingDescribeCustomerResourcesTransResponse
}

// IsSuccess returns true when this customer manager describe customer resources trans o k response has a 2xx status code
func (o *CustomerManagerDescribeCustomerResourcesTransOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer manager describe customer resources trans o k response has a 3xx status code
func (o *CustomerManagerDescribeCustomerResourcesTransOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer manager describe customer resources trans o k response has a 4xx status code
func (o *CustomerManagerDescribeCustomerResourcesTransOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer manager describe customer resources trans o k response has a 5xx status code
func (o *CustomerManagerDescribeCustomerResourcesTransOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer manager describe customer resources trans o k response a status code equal to that given
func (o *CustomerManagerDescribeCustomerResourcesTransOK) IsCode(code int) bool {
	return code == 200
}

func (o *CustomerManagerDescribeCustomerResourcesTransOK) Error() string {
	return fmt.Sprintf("[GET /v1/customer/resources/trans][%d] customerManagerDescribeCustomerResourcesTransOK  %+v", 200, o.Payload)
}

func (o *CustomerManagerDescribeCustomerResourcesTransOK) String() string {
	return fmt.Sprintf("[GET /v1/customer/resources/trans][%d] customerManagerDescribeCustomerResourcesTransOK  %+v", 200, o.Payload)
}

func (o *CustomerManagerDescribeCustomerResourcesTransOK) GetPayload() *models.NewbillingDescribeCustomerResourcesTransResponse {
	return o.Payload
}

func (o *CustomerManagerDescribeCustomerResourcesTransOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingDescribeCustomerResourcesTransResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerManagerDescribeCustomerResourcesTransDefault creates a CustomerManagerDescribeCustomerResourcesTransDefault with default headers values
func NewCustomerManagerDescribeCustomerResourcesTransDefault(code int) *CustomerManagerDescribeCustomerResourcesTransDefault {
	return &CustomerManagerDescribeCustomerResourcesTransDefault{
		_statusCode: code,
	}
}

/*
CustomerManagerDescribeCustomerResourcesTransDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type CustomerManagerDescribeCustomerResourcesTransDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the customer manager describe customer resources trans default response
func (o *CustomerManagerDescribeCustomerResourcesTransDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this customer manager describe customer resources trans default response has a 2xx status code
func (o *CustomerManagerDescribeCustomerResourcesTransDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer manager describe customer resources trans default response has a 3xx status code
func (o *CustomerManagerDescribeCustomerResourcesTransDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer manager describe customer resources trans default response has a 4xx status code
func (o *CustomerManagerDescribeCustomerResourcesTransDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer manager describe customer resources trans default response has a 5xx status code
func (o *CustomerManagerDescribeCustomerResourcesTransDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer manager describe customer resources trans default response a status code equal to that given
func (o *CustomerManagerDescribeCustomerResourcesTransDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CustomerManagerDescribeCustomerResourcesTransDefault) Error() string {
	return fmt.Sprintf("[GET /v1/customer/resources/trans][%d] CustomerManager_DescribeCustomerResourcesTrans default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerManagerDescribeCustomerResourcesTransDefault) String() string {
	return fmt.Sprintf("[GET /v1/customer/resources/trans][%d] CustomerManager_DescribeCustomerResourcesTrans default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerManagerDescribeCustomerResourcesTransDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *CustomerManagerDescribeCustomerResourcesTransDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
