// Code generated by go-swagger; DO NOT EDIT.

package charging_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/newbilling-client-go/models"
)

// ChargingManagerChargeAlipayReturnReader is a Reader for the ChargingManagerChargeAlipayReturn structure.
type ChargingManagerChargeAlipayReturnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChargingManagerChargeAlipayReturnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChargingManagerChargeAlipayReturnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChargingManagerChargeAlipayReturnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChargingManagerChargeAlipayReturnOK creates a ChargingManagerChargeAlipayReturnOK with default headers values
func NewChargingManagerChargeAlipayReturnOK() *ChargingManagerChargeAlipayReturnOK {
	return &ChargingManagerChargeAlipayReturnOK{}
}

/*
ChargingManagerChargeAlipayReturnOK describes a response with status code 200, with default header values.

A successful response.
*/
type ChargingManagerChargeAlipayReturnOK struct {
	Payload *models.NewbillingChargeAlipayReturnResponse
}

// IsSuccess returns true when this charging manager charge alipay return o k response has a 2xx status code
func (o *ChargingManagerChargeAlipayReturnOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this charging manager charge alipay return o k response has a 3xx status code
func (o *ChargingManagerChargeAlipayReturnOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this charging manager charge alipay return o k response has a 4xx status code
func (o *ChargingManagerChargeAlipayReturnOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this charging manager charge alipay return o k response has a 5xx status code
func (o *ChargingManagerChargeAlipayReturnOK) IsServerError() bool {
	return false
}

// IsCode returns true when this charging manager charge alipay return o k response a status code equal to that given
func (o *ChargingManagerChargeAlipayReturnOK) IsCode(code int) bool {
	return code == 200
}

func (o *ChargingManagerChargeAlipayReturnOK) Error() string {
	return fmt.Sprintf("[POST /v1/charge_returns/alipay][%d] chargingManagerChargeAlipayReturnOK  %+v", 200, o.Payload)
}

func (o *ChargingManagerChargeAlipayReturnOK) String() string {
	return fmt.Sprintf("[POST /v1/charge_returns/alipay][%d] chargingManagerChargeAlipayReturnOK  %+v", 200, o.Payload)
}

func (o *ChargingManagerChargeAlipayReturnOK) GetPayload() *models.NewbillingChargeAlipayReturnResponse {
	return o.Payload
}

func (o *ChargingManagerChargeAlipayReturnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingChargeAlipayReturnResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChargingManagerChargeAlipayReturnDefault creates a ChargingManagerChargeAlipayReturnDefault with default headers values
func NewChargingManagerChargeAlipayReturnDefault(code int) *ChargingManagerChargeAlipayReturnDefault {
	return &ChargingManagerChargeAlipayReturnDefault{
		_statusCode: code,
	}
}

/*
ChargingManagerChargeAlipayReturnDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type ChargingManagerChargeAlipayReturnDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the charging manager charge alipay return default response
func (o *ChargingManagerChargeAlipayReturnDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this charging manager charge alipay return default response has a 2xx status code
func (o *ChargingManagerChargeAlipayReturnDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this charging manager charge alipay return default response has a 3xx status code
func (o *ChargingManagerChargeAlipayReturnDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this charging manager charge alipay return default response has a 4xx status code
func (o *ChargingManagerChargeAlipayReturnDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this charging manager charge alipay return default response has a 5xx status code
func (o *ChargingManagerChargeAlipayReturnDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this charging manager charge alipay return default response a status code equal to that given
func (o *ChargingManagerChargeAlipayReturnDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ChargingManagerChargeAlipayReturnDefault) Error() string {
	return fmt.Sprintf("[POST /v1/charge_returns/alipay][%d] ChargingManager_ChargeAlipayReturn default  %+v", o._statusCode, o.Payload)
}

func (o *ChargingManagerChargeAlipayReturnDefault) String() string {
	return fmt.Sprintf("[POST /v1/charge_returns/alipay][%d] ChargingManager_ChargeAlipayReturn default  %+v", o._statusCode, o.Payload)
}

func (o *ChargingManagerChargeAlipayReturnDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ChargingManagerChargeAlipayReturnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
