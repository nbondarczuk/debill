// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetAccountOKCode is the HTTP code returned for type GetAccountOK
const GetAccountOKCode int = 200

/*
GetAccountOK get account o k

swagger:response getAccountOK
*/
type GetAccountOK struct {
}

// NewGetAccountOK creates GetAccountOK with default headers values
func NewGetAccountOK() *GetAccountOK {

	return &GetAccountOK{}
}

// WriteResponse to the client
func (o *GetAccountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
