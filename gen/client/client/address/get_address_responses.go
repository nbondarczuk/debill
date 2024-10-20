// Code generated by go-swagger; DO NOT EDIT.

package address

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAddressReader is a Reader for the GetAddress structure.
type GetAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /api/address] getAddress", response, response.Code())
	}
}

// NewGetAddressOK creates a GetAddressOK with default headers values
func NewGetAddressOK() *GetAddressOK {
	return &GetAddressOK{}
}

/*
GetAddressOK describes a response with status code 200, with default header values.

GetAddressOK get address o k
*/
type GetAddressOK struct {
}

// IsSuccess returns true when this get address o k response has a 2xx status code
func (o *GetAddressOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get address o k response has a 3xx status code
func (o *GetAddressOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get address o k response has a 4xx status code
func (o *GetAddressOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get address o k response has a 5xx status code
func (o *GetAddressOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get address o k response a status code equal to that given
func (o *GetAddressOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get address o k response
func (o *GetAddressOK) Code() int {
	return 200
}

func (o *GetAddressOK) Error() string {
	return fmt.Sprintf("[GET /api/address][%d] getAddressOK", 200)
}

func (o *GetAddressOK) String() string {
	return fmt.Sprintf("[GET /api/address][%d] getAddressOK", 200)
}

func (o *GetAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
