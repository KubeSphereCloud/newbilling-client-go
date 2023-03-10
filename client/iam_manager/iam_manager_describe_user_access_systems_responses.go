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

// IamManagerDescribeUserAccessSystemsReader is a Reader for the IamManagerDescribeUserAccessSystems structure.
type IamManagerDescribeUserAccessSystemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IamManagerDescribeUserAccessSystemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIamManagerDescribeUserAccessSystemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIamManagerDescribeUserAccessSystemsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIamManagerDescribeUserAccessSystemsOK creates a IamManagerDescribeUserAccessSystemsOK with default headers values
func NewIamManagerDescribeUserAccessSystemsOK() *IamManagerDescribeUserAccessSystemsOK {
	return &IamManagerDescribeUserAccessSystemsOK{}
}

/*
IamManagerDescribeUserAccessSystemsOK describes a response with status code 200, with default header values.

A successful response.
*/
type IamManagerDescribeUserAccessSystemsOK struct {
	Payload *models.NewbillingDescribeUserAccessSystemsResponse
}

// IsSuccess returns true when this iam manager describe user access systems o k response has a 2xx status code
func (o *IamManagerDescribeUserAccessSystemsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iam manager describe user access systems o k response has a 3xx status code
func (o *IamManagerDescribeUserAccessSystemsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iam manager describe user access systems o k response has a 4xx status code
func (o *IamManagerDescribeUserAccessSystemsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iam manager describe user access systems o k response has a 5xx status code
func (o *IamManagerDescribeUserAccessSystemsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iam manager describe user access systems o k response a status code equal to that given
func (o *IamManagerDescribeUserAccessSystemsOK) IsCode(code int) bool {
	return code == 200
}

func (o *IamManagerDescribeUserAccessSystemsOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/accesssystems][%d] iamManagerDescribeUserAccessSystemsOK  %+v", 200, o.Payload)
}

func (o *IamManagerDescribeUserAccessSystemsOK) String() string {
	return fmt.Sprintf("[GET /v1/users/accesssystems][%d] iamManagerDescribeUserAccessSystemsOK  %+v", 200, o.Payload)
}

func (o *IamManagerDescribeUserAccessSystemsOK) GetPayload() *models.NewbillingDescribeUserAccessSystemsResponse {
	return o.Payload
}

func (o *IamManagerDescribeUserAccessSystemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingDescribeUserAccessSystemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIamManagerDescribeUserAccessSystemsDefault creates a IamManagerDescribeUserAccessSystemsDefault with default headers values
func NewIamManagerDescribeUserAccessSystemsDefault(code int) *IamManagerDescribeUserAccessSystemsDefault {
	return &IamManagerDescribeUserAccessSystemsDefault{
		_statusCode: code,
	}
}

/*
IamManagerDescribeUserAccessSystemsDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type IamManagerDescribeUserAccessSystemsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the iam manager describe user access systems default response
func (o *IamManagerDescribeUserAccessSystemsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this iam manager describe user access systems default response has a 2xx status code
func (o *IamManagerDescribeUserAccessSystemsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iam manager describe user access systems default response has a 3xx status code
func (o *IamManagerDescribeUserAccessSystemsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iam manager describe user access systems default response has a 4xx status code
func (o *IamManagerDescribeUserAccessSystemsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iam manager describe user access systems default response has a 5xx status code
func (o *IamManagerDescribeUserAccessSystemsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iam manager describe user access systems default response a status code equal to that given
func (o *IamManagerDescribeUserAccessSystemsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IamManagerDescribeUserAccessSystemsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/users/accesssystems][%d] IamManager_DescribeUserAccessSystems default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerDescribeUserAccessSystemsDefault) String() string {
	return fmt.Sprintf("[GET /v1/users/accesssystems][%d] IamManager_DescribeUserAccessSystems default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerDescribeUserAccessSystemsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *IamManagerDescribeUserAccessSystemsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
