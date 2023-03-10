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

// IamManagerRefreshTokenReader is a Reader for the IamManagerRefreshToken structure.
type IamManagerRefreshTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IamManagerRefreshTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIamManagerRefreshTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIamManagerRefreshTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIamManagerRefreshTokenOK creates a IamManagerRefreshTokenOK with default headers values
func NewIamManagerRefreshTokenOK() *IamManagerRefreshTokenOK {
	return &IamManagerRefreshTokenOK{}
}

/*
IamManagerRefreshTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type IamManagerRefreshTokenOK struct {
	Payload *models.NewbillingRefreshTokenResponse
}

// IsSuccess returns true when this iam manager refresh token o k response has a 2xx status code
func (o *IamManagerRefreshTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iam manager refresh token o k response has a 3xx status code
func (o *IamManagerRefreshTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iam manager refresh token o k response has a 4xx status code
func (o *IamManagerRefreshTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iam manager refresh token o k response has a 5xx status code
func (o *IamManagerRefreshTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iam manager refresh token o k response a status code equal to that given
func (o *IamManagerRefreshTokenOK) IsCode(code int) bool {
	return code == 200
}

func (o *IamManagerRefreshTokenOK) Error() string {
	return fmt.Sprintf("[POST /v1/users/token][%d] iamManagerRefreshTokenOK  %+v", 200, o.Payload)
}

func (o *IamManagerRefreshTokenOK) String() string {
	return fmt.Sprintf("[POST /v1/users/token][%d] iamManagerRefreshTokenOK  %+v", 200, o.Payload)
}

func (o *IamManagerRefreshTokenOK) GetPayload() *models.NewbillingRefreshTokenResponse {
	return o.Payload
}

func (o *IamManagerRefreshTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingRefreshTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIamManagerRefreshTokenDefault creates a IamManagerRefreshTokenDefault with default headers values
func NewIamManagerRefreshTokenDefault(code int) *IamManagerRefreshTokenDefault {
	return &IamManagerRefreshTokenDefault{
		_statusCode: code,
	}
}

/*
IamManagerRefreshTokenDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type IamManagerRefreshTokenDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the iam manager refresh token default response
func (o *IamManagerRefreshTokenDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this iam manager refresh token default response has a 2xx status code
func (o *IamManagerRefreshTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iam manager refresh token default response has a 3xx status code
func (o *IamManagerRefreshTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iam manager refresh token default response has a 4xx status code
func (o *IamManagerRefreshTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iam manager refresh token default response has a 5xx status code
func (o *IamManagerRefreshTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iam manager refresh token default response a status code equal to that given
func (o *IamManagerRefreshTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IamManagerRefreshTokenDefault) Error() string {
	return fmt.Sprintf("[POST /v1/users/token][%d] IamManager_RefreshToken default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerRefreshTokenDefault) String() string {
	return fmt.Sprintf("[POST /v1/users/token][%d] IamManager_RefreshToken default  %+v", o._statusCode, o.Payload)
}

func (o *IamManagerRefreshTokenDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *IamManagerRefreshTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
