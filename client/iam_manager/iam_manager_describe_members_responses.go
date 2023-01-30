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

// IamManagerDescribeMembersReader is a Reader for the IamManagerDescribeMembers structure.
type IamManagerDescribeMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IamManagerDescribeMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIamManagerDescribeMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIamManagerDescribeMembersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIamManagerDescribeMembersOK creates a IamManagerDescribeMembersOK with default headers values
func NewIamManagerDescribeMembersOK() *IamManagerDescribeMembersOK {
	return &IamManagerDescribeMembersOK{}
}

/*
IamManagerDescribeMembersOK describes a response with status code 200, with default header values.

A successful response.
*/
type IamManagerDescribeMembersOK struct {
	Payload *models.NewbillingDescribeUsersResponse
}

// IsSuccess returns true when this iam manager describe members o k response has a 2xx status code
func (o *IamManagerDescribeMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iam manager describe members o k response has a 3xx status code
func (o *IamManagerDescribeMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iam manager describe members o k response has a 4xx status code
func (o *IamManagerDescribeMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iam manager describe members o k response has a 5xx status code
func (o *IamManagerDescribeMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iam manager describe members o k response a status code equal to that given
func (o *IamManagerDescribeMembersOK) IsCode(code int) bool {
	return code == 200
}

func (o *IamManagerDescribeMembersOK) Error() string {
	return fmt.Sprintf("[GET /v1/accesssystems/members][%d] iamManagerDescribeMembersOK  %+v", 200, o.Payload)
}

func (o *IamManagerDescribeMembersOK) String() string {
	return fmt.Sprintf("[GET /v1/accesssystems/members][%d] iamManagerDescribeMembersOK  %+v", 200, o.Payload)
}

func (o *IamManagerDescribeMembersOK) GetPayload() *models.NewbillingDescribeUsersResponse {
	return o.Payload
}

func (o *IamManagerDescribeMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingDescribeUsersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIamManagerDescribeMembersDefault creates a IamManagerDescribeMembersDefault with default headers values
func NewIamManagerDescribeMembersDefault(code int) *IamManagerDescribeMembersDefault {
	return &IamManagerDescribeMembersDefault{
		_statusCode: code,
	}
}

/*
IamManagerDescribeMembersDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type IamManagerDescribeMembersDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the iam manager describe members default response
func (o *IamManagerDescribeMembersDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this iam manager describe members default response has a 2xx status code
func (o *IamManagerDescribeMembersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iam manager describe members default response has a 3xx status code
func (o *IamManagerDescribeMembersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iam manager describe members default response has a 4xx status code
func (o *IamManagerDescribeMembersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iam manager describe members default response has a 5xx status code
func (o *IamManagerDescribeMembersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iam manager describe members default response a status code equal to that given
func (o *IamManagerDescribeMembersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IamManagerDescribeMembersDefault) Error() string {
	return fmt.Sprintf("[GET /v1/accesssystems/members][%d] IamManager_DescribeMembers default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerDescribeMembersDefault) String() string {
	return fmt.Sprintf("[GET /v1/accesssystems/members][%d] IamManager_DescribeMembers default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerDescribeMembersDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *IamManagerDescribeMembersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
