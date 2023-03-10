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

// IamManagerDescribeTokenReader is a Reader for the IamManagerDescribeToken structure.
type IamManagerDescribeTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IamManagerDescribeTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIamManagerDescribeTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIamManagerDescribeTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIamManagerDescribeTokenOK creates a IamManagerDescribeTokenOK with default headers values
func NewIamManagerDescribeTokenOK() *IamManagerDescribeTokenOK {
	return &IamManagerDescribeTokenOK{}
}

/*
IamManagerDescribeTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type IamManagerDescribeTokenOK struct {
	Payload *models.NewbillingDescribeTokenResponse
}

// IsSuccess returns true when this iam manager describe token o k response has a 2xx status code
func (o *IamManagerDescribeTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iam manager describe token o k response has a 3xx status code
func (o *IamManagerDescribeTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iam manager describe token o k response has a 4xx status code
func (o *IamManagerDescribeTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iam manager describe token o k response has a 5xx status code
func (o *IamManagerDescribeTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iam manager describe token o k response a status code equal to that given
func (o *IamManagerDescribeTokenOK) IsCode(code int) bool {
	return code == 200
}

func (o *IamManagerDescribeTokenOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/token][%d] iamManagerDescribeTokenOK  %+v", 200, o.Payload)
}

func (o *IamManagerDescribeTokenOK) String() string {
	return fmt.Sprintf("[GET /v1/users/token][%d] iamManagerDescribeTokenOK  %+v", 200, o.Payload)
}

func (o *IamManagerDescribeTokenOK) GetPayload() *models.NewbillingDescribeTokenResponse {
	return o.Payload
}

func (o *IamManagerDescribeTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingDescribeTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIamManagerDescribeTokenDefault creates a IamManagerDescribeTokenDefault with default headers values
func NewIamManagerDescribeTokenDefault(code int) *IamManagerDescribeTokenDefault {
	return &IamManagerDescribeTokenDefault{
		_statusCode: code,
	}
}

/*
IamManagerDescribeTokenDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type IamManagerDescribeTokenDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the iam manager describe token default response
func (o *IamManagerDescribeTokenDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this iam manager describe token default response has a 2xx status code
func (o *IamManagerDescribeTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iam manager describe token default response has a 3xx status code
func (o *IamManagerDescribeTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iam manager describe token default response has a 4xx status code
func (o *IamManagerDescribeTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iam manager describe token default response has a 5xx status code
func (o *IamManagerDescribeTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iam manager describe token default response a status code equal to that given
func (o *IamManagerDescribeTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IamManagerDescribeTokenDefault) Error() string {
	return fmt.Sprintf("[GET /v1/users/token][%d] IamManager_DescribeToken default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerDescribeTokenDefault) String() string {
	return fmt.Sprintf("[GET /v1/users/token][%d] IamManager_DescribeToken default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerDescribeTokenDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *IamManagerDescribeTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
