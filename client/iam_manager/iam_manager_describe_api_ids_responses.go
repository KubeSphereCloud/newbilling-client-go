// Code generated by go-swagger; DO NOT EDIT.

package iam_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/newbilling-client-go/models"
)

// IamManagerDescribeAPIIdsReader is a Reader for the IamManagerDescribeAPIIds structure.
type IamManagerDescribeAPIIdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IamManagerDescribeAPIIdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIamManagerDescribeAPIIdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIamManagerDescribeAPIIdsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIamManagerDescribeAPIIdsOK creates a IamManagerDescribeAPIIdsOK with default headers values
func NewIamManagerDescribeAPIIdsOK() *IamManagerDescribeAPIIdsOK {
	return &IamManagerDescribeAPIIdsOK{}
}

/*
IamManagerDescribeAPIIdsOK describes a response with status code 200, with default header values.

A successful response.
*/
type IamManagerDescribeAPIIdsOK struct {
	Payload *models.NewbillingDescribeAPIIdsResponse
}

// IsSuccess returns true when this iam manager describe Api ids o k response has a 2xx status code
func (o *IamManagerDescribeAPIIdsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iam manager describe Api ids o k response has a 3xx status code
func (o *IamManagerDescribeAPIIdsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iam manager describe Api ids o k response has a 4xx status code
func (o *IamManagerDescribeAPIIdsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iam manager describe Api ids o k response has a 5xx status code
func (o *IamManagerDescribeAPIIdsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iam manager describe Api ids o k response a status code equal to that given
func (o *IamManagerDescribeAPIIdsOK) IsCode(code int) bool {
	return code == 200
}

func (o *IamManagerDescribeAPIIdsOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/actions][%d] iamManagerDescribeApiIdsOK  %+v", 200, o.Payload)
}

func (o *IamManagerDescribeAPIIdsOK) String() string {
	return fmt.Sprintf("[GET /v1/users/actions][%d] iamManagerDescribeApiIdsOK  %+v", 200, o.Payload)
}

func (o *IamManagerDescribeAPIIdsOK) GetPayload() *models.NewbillingDescribeAPIIdsResponse {
	return o.Payload
}

func (o *IamManagerDescribeAPIIdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingDescribeAPIIdsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIamManagerDescribeAPIIdsDefault creates a IamManagerDescribeAPIIdsDefault with default headers values
func NewIamManagerDescribeAPIIdsDefault(code int) *IamManagerDescribeAPIIdsDefault {
	return &IamManagerDescribeAPIIdsDefault{
		_statusCode: code,
	}
}

/*
IamManagerDescribeAPIIdsDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type IamManagerDescribeAPIIdsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the iam manager describe Api ids default response
func (o *IamManagerDescribeAPIIdsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this iam manager describe Api ids default response has a 2xx status code
func (o *IamManagerDescribeAPIIdsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iam manager describe Api ids default response has a 3xx status code
func (o *IamManagerDescribeAPIIdsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iam manager describe Api ids default response has a 4xx status code
func (o *IamManagerDescribeAPIIdsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iam manager describe Api ids default response has a 5xx status code
func (o *IamManagerDescribeAPIIdsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iam manager describe Api ids default response a status code equal to that given
func (o *IamManagerDescribeAPIIdsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IamManagerDescribeAPIIdsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/users/actions][%d] IamManager_DescribeApiIds default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerDescribeAPIIdsDefault) String() string {
	return fmt.Sprintf("[GET /v1/users/actions][%d] IamManager_DescribeApiIds default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerDescribeAPIIdsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *IamManagerDescribeAPIIdsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
