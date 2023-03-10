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

// IamManagerLogoutReader is a Reader for the IamManagerLogout structure.
type IamManagerLogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IamManagerLogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIamManagerLogoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIamManagerLogoutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIamManagerLogoutOK creates a IamManagerLogoutOK with default headers values
func NewIamManagerLogoutOK() *IamManagerLogoutOK {
	return &IamManagerLogoutOK{}
}

/*
IamManagerLogoutOK describes a response with status code 200, with default header values.

A successful response.
*/
type IamManagerLogoutOK struct {
	Payload *models.NewbillingLogoutResponse
}

// IsSuccess returns true when this iam manager logout o k response has a 2xx status code
func (o *IamManagerLogoutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iam manager logout o k response has a 3xx status code
func (o *IamManagerLogoutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iam manager logout o k response has a 4xx status code
func (o *IamManagerLogoutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iam manager logout o k response has a 5xx status code
func (o *IamManagerLogoutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iam manager logout o k response a status code equal to that given
func (o *IamManagerLogoutOK) IsCode(code int) bool {
	return code == 200
}

func (o *IamManagerLogoutOK) Error() string {
	return fmt.Sprintf("[POST /v1/logout][%d] iamManagerLogoutOK  %+v", 200, o.Payload)
}

func (o *IamManagerLogoutOK) String() string {
	return fmt.Sprintf("[POST /v1/logout][%d] iamManagerLogoutOK  %+v", 200, o.Payload)
}

func (o *IamManagerLogoutOK) GetPayload() *models.NewbillingLogoutResponse {
	return o.Payload
}

func (o *IamManagerLogoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingLogoutResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIamManagerLogoutDefault creates a IamManagerLogoutDefault with default headers values
func NewIamManagerLogoutDefault(code int) *IamManagerLogoutDefault {
	return &IamManagerLogoutDefault{
		_statusCode: code,
	}
}

/*
IamManagerLogoutDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type IamManagerLogoutDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the iam manager logout default response
func (o *IamManagerLogoutDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this iam manager logout default response has a 2xx status code
func (o *IamManagerLogoutDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iam manager logout default response has a 3xx status code
func (o *IamManagerLogoutDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iam manager logout default response has a 4xx status code
func (o *IamManagerLogoutDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iam manager logout default response has a 5xx status code
func (o *IamManagerLogoutDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iam manager logout default response a status code equal to that given
func (o *IamManagerLogoutDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IamManagerLogoutDefault) Error() string {
	return fmt.Sprintf("[POST /v1/logout][%d] IamManager_Logout default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerLogoutDefault) String() string {
	return fmt.Sprintf("[POST /v1/logout][%d] IamManager_Logout default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerLogoutDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *IamManagerLogoutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
