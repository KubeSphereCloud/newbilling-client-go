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

// CustomerManagerDescribeCustomerResourcesReader is a Reader for the CustomerManagerDescribeCustomerResources structure.
type CustomerManagerDescribeCustomerResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerManagerDescribeCustomerResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerManagerDescribeCustomerResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerManagerDescribeCustomerResourcesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerManagerDescribeCustomerResourcesOK creates a CustomerManagerDescribeCustomerResourcesOK with default headers values
func NewCustomerManagerDescribeCustomerResourcesOK() *CustomerManagerDescribeCustomerResourcesOK {
	return &CustomerManagerDescribeCustomerResourcesOK{}
}

/*
CustomerManagerDescribeCustomerResourcesOK describes a response with status code 200, with default header values.

A successful response.
*/
type CustomerManagerDescribeCustomerResourcesOK struct {
	Payload *models.NewbillingDescribeCustomerResourcesResponse
}

// IsSuccess returns true when this customer manager describe customer resources o k response has a 2xx status code
func (o *CustomerManagerDescribeCustomerResourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer manager describe customer resources o k response has a 3xx status code
func (o *CustomerManagerDescribeCustomerResourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer manager describe customer resources o k response has a 4xx status code
func (o *CustomerManagerDescribeCustomerResourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer manager describe customer resources o k response has a 5xx status code
func (o *CustomerManagerDescribeCustomerResourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer manager describe customer resources o k response a status code equal to that given
func (o *CustomerManagerDescribeCustomerResourcesOK) IsCode(code int) bool {
	return code == 200
}

func (o *CustomerManagerDescribeCustomerResourcesOK) Error() string {
	return fmt.Sprintf("[GET /v1/customer/resources][%d] customerManagerDescribeCustomerResourcesOK  %+v", 200, o.Payload)
}

func (o *CustomerManagerDescribeCustomerResourcesOK) String() string {
	return fmt.Sprintf("[GET /v1/customer/resources][%d] customerManagerDescribeCustomerResourcesOK  %+v", 200, o.Payload)
}

func (o *CustomerManagerDescribeCustomerResourcesOK) GetPayload() *models.NewbillingDescribeCustomerResourcesResponse {
	return o.Payload
}

func (o *CustomerManagerDescribeCustomerResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingDescribeCustomerResourcesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerManagerDescribeCustomerResourcesDefault creates a CustomerManagerDescribeCustomerResourcesDefault with default headers values
func NewCustomerManagerDescribeCustomerResourcesDefault(code int) *CustomerManagerDescribeCustomerResourcesDefault {
	return &CustomerManagerDescribeCustomerResourcesDefault{
		_statusCode: code,
	}
}

/*
CustomerManagerDescribeCustomerResourcesDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type CustomerManagerDescribeCustomerResourcesDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the customer manager describe customer resources default response
func (o *CustomerManagerDescribeCustomerResourcesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this customer manager describe customer resources default response has a 2xx status code
func (o *CustomerManagerDescribeCustomerResourcesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer manager describe customer resources default response has a 3xx status code
func (o *CustomerManagerDescribeCustomerResourcesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer manager describe customer resources default response has a 4xx status code
func (o *CustomerManagerDescribeCustomerResourcesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer manager describe customer resources default response has a 5xx status code
func (o *CustomerManagerDescribeCustomerResourcesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer manager describe customer resources default response a status code equal to that given
func (o *CustomerManagerDescribeCustomerResourcesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CustomerManagerDescribeCustomerResourcesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/customer/resources][%d] CustomerManager_DescribeCustomerResources default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerManagerDescribeCustomerResourcesDefault) String() string {
	return fmt.Sprintf("[GET /v1/customer/resources][%d] CustomerManager_DescribeCustomerResources default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerManagerDescribeCustomerResourcesDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *CustomerManagerDescribeCustomerResourcesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
