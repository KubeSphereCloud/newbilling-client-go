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

// IamManagerDescribeRolesReader is a Reader for the IamManagerDescribeRoles structure.
type IamManagerDescribeRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IamManagerDescribeRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIamManagerDescribeRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIamManagerDescribeRolesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIamManagerDescribeRolesOK creates a IamManagerDescribeRolesOK with default headers values
func NewIamManagerDescribeRolesOK() *IamManagerDescribeRolesOK {
	return &IamManagerDescribeRolesOK{}
}

/*
IamManagerDescribeRolesOK describes a response with status code 200, with default header values.

A successful response.
*/
type IamManagerDescribeRolesOK struct {
	Payload *models.NewbillingDescribeRolesResponse
}

// IsSuccess returns true when this iam manager describe roles o k response has a 2xx status code
func (o *IamManagerDescribeRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iam manager describe roles o k response has a 3xx status code
func (o *IamManagerDescribeRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iam manager describe roles o k response has a 4xx status code
func (o *IamManagerDescribeRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iam manager describe roles o k response has a 5xx status code
func (o *IamManagerDescribeRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iam manager describe roles o k response a status code equal to that given
func (o *IamManagerDescribeRolesOK) IsCode(code int) bool {
	return code == 200
}

func (o *IamManagerDescribeRolesOK) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] iamManagerDescribeRolesOK  %+v", 200, o.Payload)
}

func (o *IamManagerDescribeRolesOK) String() string {
	return fmt.Sprintf("[GET /v1/roles][%d] iamManagerDescribeRolesOK  %+v", 200, o.Payload)
}

func (o *IamManagerDescribeRolesOK) GetPayload() *models.NewbillingDescribeRolesResponse {
	return o.Payload
}

func (o *IamManagerDescribeRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingDescribeRolesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIamManagerDescribeRolesDefault creates a IamManagerDescribeRolesDefault with default headers values
func NewIamManagerDescribeRolesDefault(code int) *IamManagerDescribeRolesDefault {
	return &IamManagerDescribeRolesDefault{
		_statusCode: code,
	}
}

/*
IamManagerDescribeRolesDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type IamManagerDescribeRolesDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the iam manager describe roles default response
func (o *IamManagerDescribeRolesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this iam manager describe roles default response has a 2xx status code
func (o *IamManagerDescribeRolesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iam manager describe roles default response has a 3xx status code
func (o *IamManagerDescribeRolesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iam manager describe roles default response has a 4xx status code
func (o *IamManagerDescribeRolesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iam manager describe roles default response has a 5xx status code
func (o *IamManagerDescribeRolesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iam manager describe roles default response a status code equal to that given
func (o *IamManagerDescribeRolesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IamManagerDescribeRolesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] IamManager_DescribeRoles default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerDescribeRolesDefault) String() string {
	return fmt.Sprintf("[GET /v1/roles][%d] IamManager_DescribeRoles default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerDescribeRolesDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *IamManagerDescribeRolesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}