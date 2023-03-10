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

// IamManagerModifyRoleReader is a Reader for the IamManagerModifyRole structure.
type IamManagerModifyRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IamManagerModifyRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIamManagerModifyRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIamManagerModifyRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIamManagerModifyRoleOK creates a IamManagerModifyRoleOK with default headers values
func NewIamManagerModifyRoleOK() *IamManagerModifyRoleOK {
	return &IamManagerModifyRoleOK{}
}

/*
IamManagerModifyRoleOK describes a response with status code 200, with default header values.

A successful response.
*/
type IamManagerModifyRoleOK struct {
	Payload *models.NewbillingModifyRoleResponse
}

// IsSuccess returns true when this iam manager modify role o k response has a 2xx status code
func (o *IamManagerModifyRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iam manager modify role o k response has a 3xx status code
func (o *IamManagerModifyRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iam manager modify role o k response has a 4xx status code
func (o *IamManagerModifyRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iam manager modify role o k response has a 5xx status code
func (o *IamManagerModifyRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iam manager modify role o k response a status code equal to that given
func (o *IamManagerModifyRoleOK) IsCode(code int) bool {
	return code == 200
}

func (o *IamManagerModifyRoleOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/roles][%d] iamManagerModifyRoleOK  %+v", 200, o.Payload)
}

func (o *IamManagerModifyRoleOK) String() string {
	return fmt.Sprintf("[PATCH /v1/roles][%d] iamManagerModifyRoleOK  %+v", 200, o.Payload)
}

func (o *IamManagerModifyRoleOK) GetPayload() *models.NewbillingModifyRoleResponse {
	return o.Payload
}

func (o *IamManagerModifyRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingModifyRoleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIamManagerModifyRoleDefault creates a IamManagerModifyRoleDefault with default headers values
func NewIamManagerModifyRoleDefault(code int) *IamManagerModifyRoleDefault {
	return &IamManagerModifyRoleDefault{
		_statusCode: code,
	}
}

/*
IamManagerModifyRoleDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type IamManagerModifyRoleDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the iam manager modify role default response
func (o *IamManagerModifyRoleDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this iam manager modify role default response has a 2xx status code
func (o *IamManagerModifyRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iam manager modify role default response has a 3xx status code
func (o *IamManagerModifyRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iam manager modify role default response has a 4xx status code
func (o *IamManagerModifyRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iam manager modify role default response has a 5xx status code
func (o *IamManagerModifyRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iam manager modify role default response a status code equal to that given
func (o *IamManagerModifyRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IamManagerModifyRoleDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1/roles][%d] IamManager_ModifyRole default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerModifyRoleDefault) String() string {
	return fmt.Sprintf("[PATCH /v1/roles][%d] IamManager_ModifyRole default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerModifyRoleDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *IamManagerModifyRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
