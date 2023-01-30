// Code generated by go-swagger; DO NOT EDIT.

package iam_manager4_appcenter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/newbilling-client-go/models"
)

// IamManager4AppcenterRefreshOpenTokenReader is a Reader for the IamManager4AppcenterRefreshOpenToken structure.
type IamManager4AppcenterRefreshOpenTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IamManager4AppcenterRefreshOpenTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIamManager4AppcenterRefreshOpenTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIamManager4AppcenterRefreshOpenTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIamManager4AppcenterRefreshOpenTokenOK creates a IamManager4AppcenterRefreshOpenTokenOK with default headers values
func NewIamManager4AppcenterRefreshOpenTokenOK() *IamManager4AppcenterRefreshOpenTokenOK {
	return &IamManager4AppcenterRefreshOpenTokenOK{}
}

/*
IamManager4AppcenterRefreshOpenTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type IamManager4AppcenterRefreshOpenTokenOK struct {
	Payload *models.NewbillingRefreshOpenTokenResponse
}

// IsSuccess returns true when this iam manager4 appcenter refresh open token o k response has a 2xx status code
func (o *IamManager4AppcenterRefreshOpenTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iam manager4 appcenter refresh open token o k response has a 3xx status code
func (o *IamManager4AppcenterRefreshOpenTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iam manager4 appcenter refresh open token o k response has a 4xx status code
func (o *IamManager4AppcenterRefreshOpenTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iam manager4 appcenter refresh open token o k response has a 5xx status code
func (o *IamManager4AppcenterRefreshOpenTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iam manager4 appcenter refresh open token o k response a status code equal to that given
func (o *IamManager4AppcenterRefreshOpenTokenOK) IsCode(code int) bool {
	return code == 200
}

func (o *IamManager4AppcenterRefreshOpenTokenOK) Error() string {
	return fmt.Sprintf("[POST /v1/open/token:refresh][%d] iamManager4AppcenterRefreshOpenTokenOK  %+v", 200, o.Payload)
}

func (o *IamManager4AppcenterRefreshOpenTokenOK) String() string {
	return fmt.Sprintf("[POST /v1/open/token:refresh][%d] iamManager4AppcenterRefreshOpenTokenOK  %+v", 200, o.Payload)
}

func (o *IamManager4AppcenterRefreshOpenTokenOK) GetPayload() *models.NewbillingRefreshOpenTokenResponse {
	return o.Payload
}

func (o *IamManager4AppcenterRefreshOpenTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingRefreshOpenTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIamManager4AppcenterRefreshOpenTokenDefault creates a IamManager4AppcenterRefreshOpenTokenDefault with default headers values
func NewIamManager4AppcenterRefreshOpenTokenDefault(code int) *IamManager4AppcenterRefreshOpenTokenDefault {
	return &IamManager4AppcenterRefreshOpenTokenDefault{
		_statusCode: code,
	}
}

/*
IamManager4AppcenterRefreshOpenTokenDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type IamManager4AppcenterRefreshOpenTokenDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the iam manager4 appcenter refresh open token default response
func (o *IamManager4AppcenterRefreshOpenTokenDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this iam manager4 appcenter refresh open token default response has a 2xx status code
func (o *IamManager4AppcenterRefreshOpenTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iam manager4 appcenter refresh open token default response has a 3xx status code
func (o *IamManager4AppcenterRefreshOpenTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iam manager4 appcenter refresh open token default response has a 4xx status code
func (o *IamManager4AppcenterRefreshOpenTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iam manager4 appcenter refresh open token default response has a 5xx status code
func (o *IamManager4AppcenterRefreshOpenTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iam manager4 appcenter refresh open token default response a status code equal to that given
func (o *IamManager4AppcenterRefreshOpenTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IamManager4AppcenterRefreshOpenTokenDefault) Error() string {
	return fmt.Sprintf("[POST /v1/open/token:refresh][%d] IamManager4Appcenter_RefreshOpenToken default  %+v", o._statusCode, o.Payload)
}

func (o *IamManager4AppcenterRefreshOpenTokenDefault) String() string {
	return fmt.Sprintf("[POST /v1/open/token:refresh][%d] IamManager4Appcenter_RefreshOpenToken default  %+v", o._statusCode, o.Payload)
}

func (o *IamManager4AppcenterRefreshOpenTokenDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *IamManager4AppcenterRefreshOpenTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
