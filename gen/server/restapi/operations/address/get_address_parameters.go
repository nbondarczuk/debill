// Code generated by go-swagger; DO NOT EDIT.

package address

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetAddressParams creates a new GetAddressParams object
//
// There are no default values defined in the spec.
func NewGetAddressParams() GetAddressParams {

	return GetAddressParams{}
}

// GetAddressParams contains all the bound params for the get address operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAddress
type GetAddressParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Unique ID of an Address
	  Required: true
	  In: query
	*/
	Address string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAddressParams() beforehand.
func (o *GetAddressParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAddress, qhkAddress, _ := qs.GetOK("Address")
	if err := o.bindAddress(qAddress, qhkAddress, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAddress binds and validates parameter Address from query.
func (o *GetAddressParams) bindAddress(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("Address", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("Address", "query", raw); err != nil {
		return err
	}
	o.Address = raw

	return nil
}
