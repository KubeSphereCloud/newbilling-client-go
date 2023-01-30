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

// IamManagerInviteUserReader is a Reader for the IamManagerInviteUser structure.
type IamManagerInviteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IamManagerInviteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIamManagerInviteUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIamManagerInviteUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIamManagerInviteUserOK creates a IamManagerInviteUserOK with default headers values
func NewIamManagerInviteUserOK() *IamManagerInviteUserOK {
	return &IamManagerInviteUserOK{}
}

/*
IamManagerInviteUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type IamManagerInviteUserOK struct {
	Payload *models.NewbillingInviteUserResponse
}

// IsSuccess returns true when this iam manager invite user o k response has a 2xx status code
func (o *IamManagerInviteUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iam manager invite user o k response has a 3xx status code
func (o *IamManagerInviteUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iam manager invite user o k response has a 4xx status code
func (o *IamManagerInviteUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iam manager invite user o k response has a 5xx status code
func (o *IamManagerInviteUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iam manager invite user o k response a status code equal to that given
func (o *IamManagerInviteUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *IamManagerInviteUserOK) Error() string {
	return fmt.Sprintf("[POST /v1/users:invite][%d] iamManagerInviteUserOK  %+v", 200, o.Payload)
}

func (o *IamManagerInviteUserOK) String() string {
	return fmt.Sprintf("[POST /v1/users:invite][%d] iamManagerInviteUserOK  %+v", 200, o.Payload)
}

func (o *IamManagerInviteUserOK) GetPayload() *models.NewbillingInviteUserResponse {
	return o.Payload
}

func (o *IamManagerInviteUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingInviteUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIamManagerInviteUserDefault creates a IamManagerInviteUserDefault with default headers values
func NewIamManagerInviteUserDefault(code int) *IamManagerInviteUserDefault {
	return &IamManagerInviteUserDefault{
		_statusCode: code,
	}
}

/*
IamManagerInviteUserDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type IamManagerInviteUserDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the iam manager invite user default response
func (o *IamManagerInviteUserDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this iam manager invite user default response has a 2xx status code
func (o *IamManagerInviteUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iam manager invite user default response has a 3xx status code
func (o *IamManagerInviteUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iam manager invite user default response has a 4xx status code
func (o *IamManagerInviteUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iam manager invite user default response has a 5xx status code
func (o *IamManagerInviteUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iam manager invite user default response a status code equal to that given
func (o *IamManagerInviteUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IamManagerInviteUserDefault) Error() string {
	return fmt.Sprintf("[POST /v1/users:invite][%d] IamManager_InviteUser default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerInviteUserDefault) String() string {
	return fmt.Sprintf("[POST /v1/users:invite][%d] IamManager_InviteUser default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerInviteUserDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *IamManagerInviteUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
