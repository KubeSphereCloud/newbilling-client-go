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

// PricingManagerDescribeCatalogsReader is a Reader for the PricingManagerDescribeCatalogs structure.
type PricingManagerDescribeCatalogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingManagerDescribeCatalogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingManagerDescribeCatalogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPricingManagerDescribeCatalogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPricingManagerDescribeCatalogsOK creates a PricingManagerDescribeCatalogsOK with default headers values
func NewPricingManagerDescribeCatalogsOK() *PricingManagerDescribeCatalogsOK {
	return &PricingManagerDescribeCatalogsOK{}
}

/*
PricingManagerDescribeCatalogsOK describes a response with status code 200, with default header values.

A successful response.
*/
type PricingManagerDescribeCatalogsOK struct {
	Payload *models.NewbillingDescribeCatalogsResponse
}

// IsSuccess returns true when this pricing manager describe catalogs o k response has a 2xx status code
func (o *PricingManagerDescribeCatalogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing manager describe catalogs o k response has a 3xx status code
func (o *PricingManagerDescribeCatalogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing manager describe catalogs o k response has a 4xx status code
func (o *PricingManagerDescribeCatalogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing manager describe catalogs o k response has a 5xx status code
func (o *PricingManagerDescribeCatalogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing manager describe catalogs o k response a status code equal to that given
func (o *PricingManagerDescribeCatalogsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PricingManagerDescribeCatalogsOK) Error() string {
	return fmt.Sprintf("[GET /v1/catalogs][%d] pricingManagerDescribeCatalogsOK  %+v", 200, o.Payload)
}

func (o *PricingManagerDescribeCatalogsOK) String() string {
	return fmt.Sprintf("[GET /v1/catalogs][%d] pricingManagerDescribeCatalogsOK  %+v", 200, o.Payload)
}

func (o *PricingManagerDescribeCatalogsOK) GetPayload() *models.NewbillingDescribeCatalogsResponse {
	return o.Payload
}

func (o *PricingManagerDescribeCatalogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewbillingDescribeCatalogsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPricingManagerDescribeCatalogsDefault creates a PricingManagerDescribeCatalogsDefault with default headers values
func NewPricingManagerDescribeCatalogsDefault(code int) *PricingManagerDescribeCatalogsDefault {
	return &PricingManagerDescribeCatalogsDefault{
		_statusCode: code,
	}
}

/*
PricingManagerDescribeCatalogsDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type PricingManagerDescribeCatalogsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the pricing manager describe catalogs default response
func (o *PricingManagerDescribeCatalogsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this pricing manager describe catalogs default response has a 2xx status code
func (o *PricingManagerDescribeCatalogsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing manager describe catalogs default response has a 3xx status code
func (o *PricingManagerDescribeCatalogsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing manager describe catalogs default response has a 4xx status code
func (o *PricingManagerDescribeCatalogsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing manager describe catalogs default response has a 5xx status code
func (o *PricingManagerDescribeCatalogsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing manager describe catalogs default response a status code equal to that given
func (o *PricingManagerDescribeCatalogsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PricingManagerDescribeCatalogsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/catalogs][%d] PricingManager_DescribeCatalogs default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerDescribeCatalogsDefault) String() string {
	return fmt.Sprintf("[GET /v1/catalogs][%d] PricingManager_DescribeCatalogs default  %+v", o._statusCode, o.Payload)
}

func (o *PricingManagerDescribeCatalogsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *PricingManagerDescribeCatalogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}