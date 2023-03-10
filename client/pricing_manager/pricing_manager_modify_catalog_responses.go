// Code generated by go-swagger; DO NOT EDIT.

package pricing_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/newbilling-client-go/models"
)

// PricingManagerModifyCatalogReader is a Reader for the PricingManagerModifyCatalog structure.
type PricingManagerModifyCatalogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingManagerModifyCatalogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingManagerModifyCatalogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPricingManagerModifyCatalogDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPricingManagerModifyCatalogOK creates a PricingManagerModifyCatalogOK with default headers values
func NewPricingManagerModifyCatalogOK() *PricingManagerModifyCatalogOK {
	return &PricingManagerModifyCatalogOK{}
}

/*
PricingManagerModifyCatalogOK describes a response with status code 200, with default header values.

A successful response.
*/
type PricingManagerModifyCatalogOK struct {
	Payload *models.NewbillingModifyCatalogResponse
}

// IsSuccess returns true when this pricing manager modify catalog o k response has a 2xx status code
func (o *PricingManagerModifyCatalogOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing manager modify catalog o k response has a 3xx status code
func (o *PricingManagerModifyCatalogOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing manager modify catalog o k response has a 4xx status code
func (o *PricingManagerModifyCatalogOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing manager modify catalog o k response has a 5xx status code
func (o *PricingManagerModifyCatalogOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing manager modify catalog o k response a status code equal to that given
func (o *PricingManagerModifyCatalogOK) IsCode(code int) bool {
	return code == 200
}

func (o *PricingManagerModifyCatalogOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/catalogs][%d] pricingManagerModifyCatalogOK  %+v", 200, o.Payload)
}

func (o *PricingManagerModifyCatalogOK) String() string {
	return fmt.Sprintf("[PATCH /v1/catalogs][%d] pricingManagerModifyCatalogOK  %+v", 200, o.Payload)
}

func (o *PricingManagerModifyCatalogOK) GetPayload() *models.NewbillingModifyCatalogResponse {
	return o.Payload
}

func (o *PricingManagerModifyCatalogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingModifyCatalogResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPricingManagerModifyCatalogDefault creates a PricingManagerModifyCatalogDefault with default headers values
func NewPricingManagerModifyCatalogDefault(code int) *PricingManagerModifyCatalogDefault {
	return &PricingManagerModifyCatalogDefault{
		_statusCode: code,
	}
}

/*
PricingManagerModifyCatalogDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PricingManagerModifyCatalogDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the pricing manager modify catalog default response
func (o *PricingManagerModifyCatalogDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this pricing manager modify catalog default response has a 2xx status code
func (o *PricingManagerModifyCatalogDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing manager modify catalog default response has a 3xx status code
func (o *PricingManagerModifyCatalogDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing manager modify catalog default response has a 4xx status code
func (o *PricingManagerModifyCatalogDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing manager modify catalog default response has a 5xx status code
func (o *PricingManagerModifyCatalogDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing manager modify catalog default response a status code equal to that given
func (o *PricingManagerModifyCatalogDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PricingManagerModifyCatalogDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1/catalogs][%d] PricingManager_ModifyCatalog default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerModifyCatalogDefault) String() string {
	return fmt.Sprintf("[PATCH /v1/catalogs][%d] PricingManager_ModifyCatalog default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerModifyCatalogDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PricingManagerModifyCatalogDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
