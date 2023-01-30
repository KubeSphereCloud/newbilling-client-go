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

// IamManagerModifyUserReader is a Reader for the IamManagerModifyUser structure.
type IamManagerModifyUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IamManagerModifyUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIamManagerModifyUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIamManagerModifyUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIamManagerModifyUserOK creates a IamManagerModifyUserOK with default headers values
func NewIamManagerModifyUserOK() *IamManagerModifyUserOK {
	return &IamManagerModifyUserOK{}
}

/*
IamManagerModifyUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type IamManagerModifyUserOK struct {
	Payload *models.NewbillingModifyUserResponse
}

// IsSuccess returns true when this iam manager modify user o k response has a 2xx status code
func (o *IamManagerModifyUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iam manager modify user o k response has a 3xx status code
func (o *IamManagerModifyUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iam manager modify user o k response has a 4xx status code
func (o *IamManagerModifyUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iam manager modify user o k response has a 5xx status code
func (o *IamManagerModifyUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iam manager modify user o k response a status code equal to that given
func (o *IamManagerModifyUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *IamManagerModifyUserOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/users][%d] iamManagerModifyUserOK  %+v", 200, o.Payload)
}

func (o *IamManagerModifyUserOK) String() string {
	return fmt.Sprintf("[PATCH /v1/users][%d] iamManagerModifyUserOK  %+v", 200, o.Payload)
}

func (o *IamManagerModifyUserOK) GetPayload() *models.NewbillingModifyUserResponse {
	return o.Payload
}

func (o *IamManagerModifyUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingModifyUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIamManagerModifyUserDefault creates a IamManagerModifyUserDefault with default headers values
func NewIamManagerModifyUserDefault(code int) *IamManagerModifyUserDefault {
	return &IamManagerModifyUserDefault{
		_statusCode: code,
	}
}

/*
IamManagerModifyUserDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type IamManagerModifyUserDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the iam manager modify user default response
func (o *IamManagerModifyUserDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this iam manager modify user default response has a 2xx status code
func (o *IamManagerModifyUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iam manager modify user default response has a 3xx status code
func (o *IamManagerModifyUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iam manager modify user default response has a 4xx status code
func (o *IamManagerModifyUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iam manager modify user default response has a 5xx status code
func (o *IamManagerModifyUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iam manager modify user default response a status code equal to that given
func (o *IamManagerModifyUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IamManagerModifyUserDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1/users][%d] IamManager_ModifyUser default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerModifyUserDefault) String() string {
	return fmt.Sprintf("[PATCH /v1/users][%d] IamManager_ModifyUser default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerModifyUserDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *IamManagerModifyUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
