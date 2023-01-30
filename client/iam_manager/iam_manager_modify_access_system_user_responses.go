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

// IamManagerModifyAccessSystemUserReader is a Reader for the IamManagerModifyAccessSystemUser structure.
type IamManagerModifyAccessSystemUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IamManagerModifyAccessSystemUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIamManagerModifyAccessSystemUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIamManagerModifyAccessSystemUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIamManagerModifyAccessSystemUserOK creates a IamManagerModifyAccessSystemUserOK with default headers values
func NewIamManagerModifyAccessSystemUserOK() *IamManagerModifyAccessSystemUserOK {
	return &IamManagerModifyAccessSystemUserOK{}
}

/*
IamManagerModifyAccessSystemUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type IamManagerModifyAccessSystemUserOK struct {
	Payload *models.NewbillingModifyAccessSystemUserResponse
}

// IsSuccess returns true when this iam manager modify access system user o k response has a 2xx status code
func (o *IamManagerModifyAccessSystemUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iam manager modify access system user o k response has a 3xx status code
func (o *IamManagerModifyAccessSystemUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iam manager modify access system user o k response has a 4xx status code
func (o *IamManagerModifyAccessSystemUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iam manager modify access system user o k response has a 5xx status code
func (o *IamManagerModifyAccessSystemUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iam manager modify access system user o k response a status code equal to that given
func (o *IamManagerModifyAccessSystemUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *IamManagerModifyAccessSystemUserOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/members][%d] iamManagerModifyAccessSystemUserOK  %+v", 200, o.Payload)
}

func (o *IamManagerModifyAccessSystemUserOK) String() string {
	return fmt.Sprintf("[PATCH /v1/members][%d] iamManagerModifyAccessSystemUserOK  %+v", 200, o.Payload)
}

func (o *IamManagerModifyAccessSystemUserOK) GetPayload() *models.NewbillingModifyAccessSystemUserResponse {
	return o.Payload
}

func (o *IamManagerModifyAccessSystemUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingModifyAccessSystemUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIamManagerModifyAccessSystemUserDefault creates a IamManagerModifyAccessSystemUserDefault with default headers values
func NewIamManagerModifyAccessSystemUserDefault(code int) *IamManagerModifyAccessSystemUserDefault {
	return &IamManagerModifyAccessSystemUserDefault{
		_statusCode: code,
	}
}

/*
IamManagerModifyAccessSystemUserDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type IamManagerModifyAccessSystemUserDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the iam manager modify access system user default response
func (o *IamManagerModifyAccessSystemUserDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this iam manager modify access system user default response has a 2xx status code
func (o *IamManagerModifyAccessSystemUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iam manager modify access system user default response has a 3xx status code
func (o *IamManagerModifyAccessSystemUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iam manager modify access system user default response has a 4xx status code
func (o *IamManagerModifyAccessSystemUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iam manager modify access system user default response has a 5xx status code
func (o *IamManagerModifyAccessSystemUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iam manager modify access system user default response a status code equal to that given
func (o *IamManagerModifyAccessSystemUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IamManagerModifyAccessSystemUserDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1/members][%d] IamManager_ModifyAccessSystemUser default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerModifyAccessSystemUserDefault) String() string {
	return fmt.Sprintf("[PATCH /v1/members][%d] IamManager_ModifyAccessSystemUser default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerModifyAccessSystemUserDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *IamManagerModifyAccessSystemUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
